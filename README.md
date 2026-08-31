# API Email Verification

Готовый учебный проект под домашнее задание.

## Что реализовано

- SMTP-конфигурация из `.env`;
- `POST /send` — читает локальные данные из `data/verification.json` и отправляет письмо;
- `GET /verify/{hash}` — сравнивает hash из URL с hash из локального JSON;
- кастомный `http.ServeMux`;
- `github.com/jordan-wright/email`.

## Структура

```text
api-email-auch-final/
├── cmd/
│   └── main.go
├── data/
│   └── verification.json
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── server/
│   │   └── mux.go
│   └── verify/
│       └── handler.go
├── .env.example
├── .gitignore
├── go.mod
└── README.md
```

## Запуск

1. Создай `.env` из примера:

```bash
cp .env.example .env
```

2. Заполни:

```env
EMAIL=your_email@gmail.com
PASSWORD=your_app_password
ADDRESS=smtp.gmail.com:587
```

Для Gmail обычно нужен App Password.

3. Укажи получателя и hash в `data/verification.json`:

```json
{
  "email": "recipient@example.com",
  "hash": "demo-hash-123"
}
```

4. Подтяни зависимости:

```bash
go mod tidy
```

5. Запускай из корня проекта:

```bash
go run ./cmd
```

## Проверка

Отправить письмо:

```bash
curl -X POST http://localhost:8080/send
```

Проверить hash:

```bash
curl http://localhost:8080/verify/demo-hash-123
```

Ожидаемый ответ:

```text
email verified
```

## Важно

`.env` не коммить в Git.

В исходном задании не указан точный формат локального JSON, поэтому здесь используется минимальный вариант с двумя полями: `email` и `hash`.
