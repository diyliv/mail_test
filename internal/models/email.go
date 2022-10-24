package models

import (
	"context"
	"time"

	"github.com/lib/pq"
)

type Email struct {
	SenderEmail          string         `json:"sender_email"`
	SenderMsg            string         `json:"sender_msg"`
	SenderPassword       string         `json:"sender_password"`
	SenderHashedPassword string         `json:"sender_hashed_password"`
	RecipientsEmail      pq.StringArray `json:"recipients_email"`
	EmailStatus          string         `json:"email_status"`
	SentAt               time.Time      `json:"sent_at"`
}

type EmailInterface interface {
	SendEmail(ctx context.Context, email Email) error
	DelayEmail(ctx context.Context, email Email) error
}
