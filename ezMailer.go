package ezMailer

import (
	"bytes"
	"net/smtp"
	"text/template"
	"time"
)

// Builder //
type Builder struct {
	To       string
	Username string
	Secret   string
	Server   string
	Template string
	Data     interface{}
}

type data struct {
	To     string
	Secret string
}

// SendMail //
func (a Builder) SendMail() error {
	t := time.Now()

	// auth to the servermail
	auth := smtp.PlainAuth("", a.Username, a.Secret, a.Server)

	bodyhtml, err := parseTemplate(a.Template, a.Data)
	if err != nil {
		return err
	}

	// Setup Headers / body
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n"
	header := "From: noreply@mail.fr\nDate: " + t.String() + "\n\n"
	subject := "Subject: Réinitialisation du mot de passe\n"
	msg := []byte(subject + mime + header + "\n" + bodyhtml)

	// Send mail
	err = smtp.SendMail(a.Server+":587", auth, a.Username, []string{a.To}, msg)

	return err
}

// parseTemplate used for read .tpl
func parseTemplate(templateName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateName)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
