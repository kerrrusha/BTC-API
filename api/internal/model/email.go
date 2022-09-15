package model

import "net/mail"

type Email struct {
	Email string `json:"email"`
}

func (email *Email) IsValid() bool {
	_, err := mail.ParseAddress(email.Email)
	return err == nil
}

type Emails struct {
	Emails []string `json:"emails"`
}
