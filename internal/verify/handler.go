package verify

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/smtp"
	"net/url"
	"os"

	"api-email-auch/internal/config"

	"github.com/jordan-wright/email"
)

const dataFilePath = "data/verification.json"

type Handler struct {
	cfg *config.Config
}

type VerificationData struct {
	Email string `json:"email"`
	Hash  string `json:"hash"`
}

func NewHandler(cfg *config.Config) *Handler {
	return &Handler{cfg: cfg}
}

func (h *Handler) SendHandler(w http.ResponseWriter, r *http.Request) {
	data, err := loadVerificationData()
	if err != nil {
		http.Error(w, "failed to read verification data", http.StatusInternalServerError)
		return
	}

	if data.Email == "" || data.Hash == "" {
		http.Error(w, "verification data is incomplete", http.StatusInternalServerError)
		return
	}

	if h.cfg.Email == "" || h.cfg.Password == "" || h.cfg.Address == "" {
		http.Error(w, "SMTP config is incomplete", http.StatusInternalServerError)
		return
	}

	smtpHost, _, err := net.SplitHostPort(h.cfg.Address)
	if err != nil {
		http.Error(w, "invalid SMTP address", http.StatusInternalServerError)
		return
	}

	verifyURL := "http://localhost:8080/verify/" + url.PathEscape(data.Hash)

	msg := email.NewEmail()
	msg.From = h.cfg.Email
	msg.To = []string{data.Email}
	msg.Subject = "Email verification"
	msg.Text = []byte("Verify your email: " + verifyURL)

	auth := smtp.PlainAuth("", h.cfg.Email, h.cfg.Password, smtpHost)

	if err := msg.Send(h.cfg.Address, auth); err != nil {
		http.Error(w, "failed to send verification email", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "verification email sent")
}

func (h *Handler) VerifyHandler(w http.ResponseWriter, r *http.Request) {
	hash := r.PathValue("hash")
	if hash == "" {
		http.Error(w, "hash is required", http.StatusBadRequest)
		return
	}

	data, err := loadVerificationData()
	if err != nil {
		http.Error(w, "failed to read verification data", http.StatusInternalServerError)
		return
	}

	if hash != data.Hash {
		http.Error(w, "invalid verification hash", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "email verified")
}

func loadVerificationData() (VerificationData, error) {
	file, err := os.Open(dataFilePath)
	if err != nil {
		return VerificationData{}, err
	}
	defer file.Close()

	var data VerificationData
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return VerificationData{}, err
	}

	return data, nil
}
