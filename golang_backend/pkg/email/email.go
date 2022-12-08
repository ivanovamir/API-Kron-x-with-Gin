package email

import (
	"bytes"
	"fmt"
	"html/template"
	"kron-x/internal/dto"
	"net/smtp"
)

var _ Email = &email{}

type email struct {
	MimeHeaders       string
	SenderEmail       string
	SenderAppPassword string
	SmtpHost          string
	SmtpPort          string
	DirectorEmail     string
}

func (e *email) SendEmailCode(email, code string) error {
	var body bytes.Buffer

	auth := smtp.PlainAuth("", e.SenderEmail, e.SenderAppPassword, e.SmtpHost)

	address := e.SmtpHost + ":" + e.SmtpPort

	body.Write([]byte(fmt.Sprintf("Subject:«Kron-x» \n%s\n\n", e.MimeHeaders)))

	t, err := template.ParseFiles("templates/code.html")
	if err != nil {
		return err
	} else {
		if err := t.Execute(&body, struct {
			Code string
		}{
			Code: code,
		}); err != nil {
			return err
		} else {
			// Sending email.
			err = smtp.SendMail(address, auth, e.SenderEmail, []string{email}, body.Bytes())
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (e *email) SendEmailPassword(email, userPassword string) error {
	var body bytes.Buffer

	auth := smtp.PlainAuth("", e.SenderEmail, e.SenderAppPassword, e.SmtpHost)

	address := e.SmtpHost + ":" + e.SmtpPort

	body.Write([]byte(fmt.Sprintf("Subject:«Kron-x» \n%s\n\n", e.MimeHeaders)))

	t, err := template.ParseFiles("templates/password.html")
	if err != nil {
		return err
	} else {
		if err := t.Execute(&body, struct {
			Password string
		}{
			Password: userPassword,
		}); err != nil {
			return err
		} else {
			// Sending email.
			err = smtp.SendMail(address, auth, e.SenderEmail, []string{email}, body.Bytes())
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (e *email) SenEmailChangePassword(email, newPassword string) error {

	var body bytes.Buffer

	auth := smtp.PlainAuth("", e.SenderEmail, e.SenderAppPassword, e.SmtpHost)

	address := e.SmtpHost + ":" + e.SmtpPort

	body.Write([]byte(fmt.Sprintf("Subject:«Kron-x» \n%s\n\n", e.MimeHeaders)))

	t, err := template.ParseFiles("templates/forgotpassword.html")
	if err != nil {
		return err
	} else {
		if err := t.Execute(&body, struct {
			NewPassword string
		}{
			NewPassword: newPassword,
		}); err != nil {
			return err
		} else {
			// Sending email.
			err = smtp.SendMail(address, auth, e.SenderEmail, []string{email}, body.Bytes())
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (e *email) SendOrderUr(emailData *dto.SendOrderEmailUr) error {
	var body bytes.Buffer

	// Authentication.
	auth := smtp.PlainAuth("", e.SenderEmail, e.SenderAppPassword, e.SmtpHost)

	address := e.SmtpHost + ":" + e.SmtpPort

	body.Write([]byte(fmt.Sprintf("Subject:Ваш заказ с сайта: «Kron-x» \n%s\n\n", e.MimeHeaders)))

	t, err := template.ParseFiles("templates/index_jurik.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	if err := t.Execute(&body, emailData); err != nil {
		return err
	} else {
		// Sending email.
		err = smtp.SendMail(address, auth, e.SenderEmail, []string{emailData.Email, e.DirectorEmail}, body.Bytes())
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *email) SendOrderPhyz(emailData *dto.SendOrderEmailPhyz) error {
	var body bytes.Buffer

	// Authentication.
	auth := smtp.PlainAuth("", e.SenderEmail, e.SenderAppPassword, e.SmtpHost)

	address := e.SmtpHost + ":" + e.SmtpPort

	body.Write([]byte(fmt.Sprintf("Subject:Ваш заказ с сайта: «Kron-x» \n%s\n\n", e.MimeHeaders)))

	t, err := template.ParseFiles("templates/index_fizik.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	if err := t.Execute(&body, emailData); err != nil {
		return err
	} else {
		// Sending email.
		err = smtp.SendMail(address, auth, e.SenderEmail, []string{e.DirectorEmail}, body.Bytes())
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *email) SendFeedbackCall(emailData *dto.FeedbackCall) error {
	var body bytes.Buffer

	// Authentication.
	auth := smtp.PlainAuth("", e.SenderEmail, e.SenderAppPassword, e.SmtpHost)

	address := e.SmtpHost + ":" + e.SmtpPort

	body.Write([]byte(fmt.Sprintf("Subject:«Kron-x» \n%s\n\n", e.MimeHeaders)))

	t, err := template.ParseFiles("templates/feedback_call.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	if err := t.Execute(&body, emailData); err != nil {
		return err
	} else {
		// Sending email.
		err = smtp.SendMail(address, auth, e.SenderEmail, []string{e.DirectorEmail}, body.Bytes())
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *email) SendFeedbackSelection(emailData *dto.FeedbackSelection) error {
	var body bytes.Buffer

	// Authentication.
	auth := smtp.PlainAuth("", e.SenderEmail, e.SenderAppPassword, e.SmtpHost)

	address := e.SmtpHost + ":" + e.SmtpPort

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject:«Kron-x» \n%s\n\n", mimeHeaders)))

	t, err := template.ParseFiles("templates/feedback_selection.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	if err := t.Execute(&body, emailData); err != nil {
		return err
	} else {
		// Sending email.
		err = smtp.SendMail(address, auth, e.SenderEmail, []string{e.DirectorEmail}, body.Bytes())
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *email) SendSupport(emailData *dto.Support) error {
	var body bytes.Buffer

	// Authentication.
	auth := smtp.PlainAuth("", e.SenderEmail, e.SenderAppPassword, e.SmtpHost)

	address := e.SmtpHost + ":" + e.SmtpPort

	body.Write([]byte(fmt.Sprintf("Subject:«Kron-x» \n%s\n\n", e.MimeHeaders)))

	t, err := template.ParseFiles("templates/support.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	if err := t.Execute(&body, emailData); err != nil {
		return err
	} else {
		// Sending email.
		err = smtp.SendMail(address, auth, e.SenderEmail, []string{e.DirectorEmail}, body.Bytes())
		if err != nil {
			return err
		}
	}
	return nil
}
