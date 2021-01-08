package main

import (
	"os"
	"strconv"

	"github.com/astaxie/beego/logs"
	"github.com/joho/godotenv"
	"gopkg.in/mail.v2"
	gomail "gopkg.in/mail.v2"
)

func sendEmail() {
	err := godotenv.Load()
	if err != nil {
		logs.Error("Error loading .env file")
	}

	SMTP := os.Getenv("SMTP")
	SMTPPORT, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	SMTPID := os.Getenv("SMTPID")
	SMTPPASS := os.Getenv("SMTPPASS")

	m := gomail.NewMessage()

	m.SetHeader("From", "register@closersonline.com")
	m.SetHeader("To", "kalusstone@gmail.com")
	m.SetHeader("Subject", "[NA] This is a reduce rebounce rate mail")
	m.SetBody("text/html", "This is a reduce rebounce rate mail from a cron batch job")

	d := gomail.NewDialer(SMTP, SMTPPORT, SMTPID, SMTPPASS)
	d.StartTLSPolicy = mail.MandatoryStartTLS

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		logs.Error("send email error: ", err)
	} else {
		logs.Info("success send email")
	}

}

func main() {
	//logging
	logs.SetLogger(logs.AdapterFile, `{"filename":"./logs/project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":14,"color":true}`)
	sendEmail()
}
