package repository

import (
	"context"
	"database/sql"

	"github.com/diyliv/mailganer/internal/models"
	"go.uber.org/zap"
)

type postgres struct {
	logger *zap.Logger
	psql   *sql.DB
}

func NewPostgres(logger *zap.Logger, psql *sql.DB) *postgres {
	return &postgres{
		logger: logger,
		psql:   psql,
	}
}
func (p *postgres) SendEmail(ctx context.Context, email models.Email) error {
	_, err := p.psql.Exec("INSERT INTO mail (sender_email, sender_msg, sender_password, sender_hashed_password, recipients_email, email_status) VALUES($1, $2, $3, $4, $5, $6)",
		email.SenderEmail,
		email.SenderMsg,
		email.SenderPassword,
		email.SenderHashedPassword,
		email.RecipientsEmail,
		email.EmailStatus)
	if err != nil {
		p.logger.Error("Error while inserting data: " + err.Error())
		return err
	}

	return nil
}

func (p *postgres) DelayEmail(ctx context.Context, email models.Email) error {
	_, err := p.psql.Exec("INSERT INTO mail (sender_email, sender_msg, sender_password, sender_hashed_password, recipients_email, email_status) VALUES($1, $2, $3, $4, $5, $6)",
		email.SenderEmail,
		email.SenderMsg,
		email.SenderPassword,
		email.SenderHashedPassword,
		email.RecipientsEmail,
		email.EmailStatus)
	if err != nil {
		p.logger.Error("Error while inserting data: " + err.Error())
		return err
	}

	return nil
}
