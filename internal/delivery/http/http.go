package http

import (
	"context"
	"encoding/json"
	"net/http"
	"net/smtp"
	"time"

	"github.com/diyliv/mailganer/config"
	"github.com/diyliv/mailganer/internal/models"
	"github.com/diyliv/mailganer/pkg/utils/hashpass"
	"go.uber.org/zap"
)

type httpservice struct {
	cfg               *config.Config
	logger            *zap.Logger
	postgresInterface models.EmailInterface
}

func NewHttpService(cfg *config.Config, logger *zap.Logger, postgresInterface models.EmailInterface) *httpservice {
	return &httpservice{
		cfg:               cfg,
		logger:            logger,
		postgresInterface: postgresInterface,
	}
}

func (h *httpservice) SendEmail(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var emailInfo models.Email

		if err := json.NewDecoder(r.Body).Decode(&emailInfo); err != nil {
			h.logger.Error("Error while decoding body: " + err.Error())
		}

		emailInfo.SenderPassword = h.cfg.Email.SenderPassword
		hashedPass := hashpass.HashPass(emailInfo.SenderPassword)

		if err := h.sendEmail("", "", emailInfo); err != nil {
			h.logger.Error("Error while sending email: " + err.Error())
		}

		if err := h.postgresInterface.SendEmail(ctx, models.Email{
			SenderEmail:          emailInfo.SenderEmail,
			SenderPassword:       emailInfo.SenderPassword,
			SenderMsg:            emailInfo.SenderMsg,
			SenderHashedPassword: string(hashedPass),
			RecipientsEmail:      emailInfo.RecipientsEmail,
			EmailStatus:          "sent",
		}); err != nil {
			h.logger.Error("Error while calling db method: " + err.Error())
		}

		if err := h.writeResponse(w, http.StatusOK,
			emailInfo.SenderEmail,
			emailInfo.SenderMsg,
			emailInfo.RecipientsEmail); err != nil {
			h.logger.Error("Error while writing response: " + err.Error())
		}
	}
}

func (h *httpservice) DelayEmail(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var emailInfo models.Email

		if err := json.NewDecoder(r.Body).Decode(&emailInfo); err != nil {
			h.logger.Error("Error while decoding body: " + err.Error())
		}

		emailInfo.SenderPassword = h.cfg.Email.SenderPassword
		hashedPass := hashpass.HashPass(emailInfo.SenderPassword)

		if err := h.postgresInterface.SendEmail(ctx, models.Email{
			SenderEmail:          emailInfo.SenderEmail,
			SenderPassword:       emailInfo.SenderPassword,
			SenderMsg:            emailInfo.SenderMsg,
			SenderHashedPassword: string(hashedPass),
			RecipientsEmail:      emailInfo.RecipientsEmail,
			EmailStatus:          "sent",
		}); err != nil {
			h.logger.Error("Error while calling db method: " + err.Error())
		}

		h.logger.Info("Got delay message...")
		go func() {
			timer := time.NewTicker(time.Second * 5)
			for {
				select {
				case <-timer.C:
					h.logger.Info("Sending delay message...")
					if err := h.sendEmail("", "", emailInfo); err != nil {
						h.logger.Error("Error while sending delay message: " + err.Error())
						if err := h.writeResponse(w, http.StatusBadRequest, "Something went wrong"); err != nil {
							h.logger.Error("Error while writing response: " + err.Error())
						}
					}
					if err := h.writeResponse(w, http.StatusOK,
						emailInfo.SenderEmail,
						emailInfo.SenderMsg,
						emailInfo.RecipientsEmail); err != nil {
						h.logger.Error("Error while writing response: " + err.Error())
					}
				case <-ctx.Done():
					break
				}
			}
		}()
	}
}

func (h *httpservice) sendEmail(addr, host string, email models.Email) error {
	auth := smtp.PlainAuth("", email.SenderEmail, email.SenderPassword, host)

	if err := smtp.SendMail(
		addr,
		auth,
		email.SenderEmail,
		email.RecipientsEmail,
		[]byte(email.SenderMsg)); err != nil {
		h.logger.Error("Error while sending email: " + err.Error())
		return err
	}

	return nil
}

func (h *httpservice) writeResponse(w http.ResponseWriter, code int, data ...interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(&data); err != nil {
		h.logger.Error("Error while encoding data: " + err.Error())
		return err
	}
	return nil
}
