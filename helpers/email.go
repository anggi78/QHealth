package helpers

import (
	"log"
	"os"
	"strconv"

	"github.com/go-gomail/gomail"
)

func SendEmail(email string, subject string, code string) error {

	smtpServer := os.Getenv("SMTP_SERVER")
	smtpPortStr := os.Getenv("SMTP_PORT")
	smtpUsername := os.Getenv("SMTP_USERNAME")
	smtpPassword := os.Getenv("SMTP_PASSWORD")

	smtpPort, err := strconv.Atoi(smtpPortStr)
	if err != nil {
		return err
	}

	verificationEmailHTML := `
	<html>
	<head>
		<meta charset="UTF-8">
		<title>Forgot Password - Code Verification</title>
		<style>
			body {
				font-family: Arial, sans-serif;
			}
			.container {
				width: 80%;
				margin: 0 auto;
				border: 1px solid #ccc;
				padding: 20px;
			}
			.header {
				background-color: #f0f0f0;
				padding: 10px;
			}
			.header h2 {
				margin: 0;
				color: #333;
			}
			.content {
				margin-top: 20px;
			}
			.content p {
				margin: 0;
				color: #333;
			}
			.footer {
				margin-top: 20px;
				text-align: center;
				color: #777;
			}
			.footer p {
				margin: 0;
			}
		</style>
	</head>
	<body>
		<div class="container">
			<div class="header">
				<h2>Forgot Password - Code Verification</h2>
			</div>
			<div class="content">
				<p>We received a request to reset your password. Please use the following code to verify your identity:</p>
				<p><b>Verification Code: ` + code + `</b></p>
				<p>If you did not request a password reset, please ignore this email.</p>
			</div>
			<div class="footer">
				<p>All rights reserved &copy; 2023 Your Company</p>
			</div>
		</div>
	</body>
	</html>`

	m := gomail.NewMessage()
	m.SetHeader("From", smtpUsername)
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", verificationEmailHTML)

	d := gomail.NewDialer(smtpServer, smtpPort, smtpUsername, smtpPassword)

	if err := d.DialAndSend(m); err != nil {
		log.Printf("Failed to send email to %s: %v", email, err)
		return err
	}	

	return nil
}