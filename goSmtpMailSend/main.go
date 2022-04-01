package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/smtp"
	"os"
)

var (
	// smtp setting
	gEmailId  string = "" // smtp sender email
	gEmailPw  string = "" // smtp sender password
	gSmtpHost string = "" // smtp host (smtp.gmail.com)
	gSmtpPort string = "" // smtp port (587)

	// email setting
	gMailTitle     string   = ""         // email title
	gMailContent   string   = ""         // email content
	gMailRecipient []string = []string{} // email recipient
)

type (
	// smtp setting struct
	SmtpConf struct {
		GmailId  string `json:"g_mail_id"`
		GmailPw  string `json:"g_mail_pw"`
		SmtpHost string `json:"g_smtp_host"`
		SmtpPort string `json:"g_smtp_port"`
	}

	// mail setting struct
	MailConf struct {
		Title     string
		Content   string
		Recipient []string
	}
)

func init() {
	var (
		sSmtpConf SmtpConf = SmtpConf{}
		sMailConf MailConf = MailConf{}
		sErr      error    = nil
		sConf     *os.File
	)

	// smtp setting
	sConf, sErr = os.Open("config.conf")
	if sErr != nil {
		fmt.Println(sErr)
	}

	sByte, _ := ioutil.ReadAll(sConf)
	json.Unmarshal(sByte, &sSmtpConf)

	// mail setting
	sConf, sErr = os.Open("mailSet.json")
	if sErr != nil {
		fmt.Println(sErr)
	}

	sByte, _ = ioutil.ReadAll(sConf)
	json.Unmarshal(sByte, &sMailConf)

	gEmailId = sSmtpConf.GmailId
	gEmailPw = sSmtpConf.GmailPw
	gSmtpHost = sSmtpConf.SmtpHost
	gSmtpPort = sSmtpConf.SmtpPort

	gMailTitle = sMailConf.Title
	gMailContent = sMailConf.Content
	gMailRecipient = sMailConf.Recipient

} // end func init()

func main() {
	var (
		sAuth smtp.Auth
		sFrom string   // sender
		sTo   []string // recipient list
		sErr  error

		sHeaderSubject string
		sHeaderBlank   string
		sBody          string
		sMsg           []byte

		sSumTxt []byte
	)

	// Set login information on the mail server.
	sAuth = smtp.PlainAuth("", gEmailId, gEmailPw, gSmtpHost)

	sFrom = gEmailId
	sTo = gMailRecipient

	// Compose email
	sHeaderSubject = "Subject: " + gMailTitle + "\r\n"
	sHeaderBlank = "\r\n"
	sBody = string(sSumTxt) + "\r\n"
	sMsg = []byte(sHeaderSubject + sHeaderBlank + sBody)

	// Send email
	sErr = smtp.SendMail(gSmtpHost+":"+gSmtpPort, sAuth, sFrom, sTo, sMsg)
	if sErr != nil {
		fmt.Println("sErr : ", sErr)
		panic(sErr)
	}

	fmt.Println("Mail send complete")

} // end func main()
