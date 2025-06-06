package helpers

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-gomail/gomail"
)

type Mail struct {
	Host     string
	Port     string
	Username string
	Password string
}

func SendMessageNotification(email string) error {
	mail := Mail{
		Host:     os.Getenv("SMTP_SERVER"),
		Port:     os.Getenv("SMTP_PORT"),
		Username: os.Getenv("SMTP_USERNAME"),
		Password: os.Getenv("SMTP_PASSWORD"),
	}

	body := `
		<html>
		<head>
			<meta charset="UTF-8">
			<title>New Message Notification</title>
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
					<h2>New Message Received</h2>
				</div>
				<div class="content">
					<p>Hello,</p>
					<p>You have received a new message. Please check your application to view it.</p>
				</div>
				<div class="footer">
					<p>All rights reserved &copy; 2023 Your Company</p>
				</div>
			</div>
		</body>
		</html>`

	to := []string{email}

	m := gomail.NewMessage()
	m.SetHeader("From", mail.Username)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", "QHealth Notification")
	m.SetBody("text/html", body)

	Port, _ := strconv.Atoi(mail.Port)

	dialer := gomail.NewDialer(
		mail.Host,
		Port,
		mail.Username,
		mail.Password,
	)

	if err := dialer.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func SendQueueNotification(email string) error {
    if email == "" {
        return fmt.Errorf("email tidak tersedia")
    }

    mail := Mail{
        Host:     os.Getenv("SMTP_SERVER"),
        Port:     os.Getenv("SMTP_PORT"),
        Username: os.Getenv("SMTP_USERNAME"),
        Password: os.Getenv("SMTP_PASSWORD"),
    }

    // Format konten email
    body := `
    <html>
    <head>
        <meta charset="UTF-8">
        <title>Queue Notification</title>
    </head>
    <body>
        <h3>Halo,</h3>
        <p>Posisi antrean Anda telah berubah. Silakan cek aplikasi untuk informasi lebih lanjut.</p>
        <p>Terima kasih telah menggunakan layanan kami!</p>
        <br>
        <p>Salam,</p>
        <p>Tim QHealth</p>
    </body>
    </html>
    `

    to := []string{email}

    // Buat pesan email
    m := gomail.NewMessage()
    m.SetHeader("From", mail.Username)
    m.SetHeader("To", to...)
    m.SetHeader("Subject", "QHealth Notification: Perubahan Posisi Antrean")
    m.SetBody("text/html", body)

    // Konversi port
    port, err := strconv.Atoi(mail.Port)
    if err != nil {
        return fmt.Errorf("invalid SMTP port: %v", err)
    }

    // Dialer SMTP
    dialer := gomail.NewDialer(
        mail.Host,
        port,
        mail.Username,
        mail.Password,
    )

    // Kirim email
    if err := dialer.DialAndSend(m); err != nil {
        return fmt.Errorf("failed to send email: %v", err)
    }

    return nil
}


func SendOTP(email string, otp string) error {
	mail := Mail{
		Host:     os.Getenv("SMTP_SERVER"),
		Port:     os.Getenv("SMTP_PORT"),
		Username: os.Getenv("SMTP_USERNAME"),
		Password: os.Getenv("SMTP_PASSWORD"),
	}

	body := `<html xmlns:v="urn:schemas-microsoft-com:vml" xmlns:o="urn:schemas-microsoft-com:office:office" lang="en"> <head> <title></title> <meta http-equiv="Content-Type" content="text/html; charset=utf-8"> <meta name="viewport" content="width=device-width, initial-scale=1.0"> <!--[if mso]><xml><o:OfficeDocumentSettings><o:PixelsPerInch>96</o:PixelsPerInch><o:AllowPNG/></o:OfficeDocumentSettings></xml><![endif]--> <style> * { box-sizing: border-box; } body { margin: 0; padding: 0; } a[x-apple-data-detectors] { color: inherit !important; text-decoration: inherit !important; } #MessageViewBody a { color: inherit; text-decoration: none; } p { line-height: inherit } .desktop_hide, .desktop_hide table { mso-hide: all; display: none; max-height: 0px; overflow: hidden; } .image_block img+div { display: none; } @media (max-width:660px) { .desktop_hide table.icons-inner, .social_block.desktop_hide .social-table { display: inline-block !important; } .icons-inner { text-align: center; } .icons-inner td { margin: 0 auto; } .mobile_hide { display: none; } .row-content { width: 100% !important; } .stack .column { width: 100%; display: block; } .mobile_hide { min-height: 0; max-height: 0; max-width: 0; overflow: hidden; font-size: 0px; } .desktop_hide, .desktop_hide table { display: table !important; max-height: none !important; } } </style> </head> <body style="background-color: #f1f4f8; margin: 0; padding: 0; -webkit-text-size-adjust: none; text-size-adjust: none;"> <table class="nl-container" width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; background-color: #f1f4f8;"> <tbody> <tr> <td> <table class="row row-1" align="center" width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;"> <tbody> <tr> <td> <table class="row-content stack" align="center" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; color: #000; width: 640px; margin: 0 auto;" width="640"> <tbody> <tr> <td class="column column-1" width="100%" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; font-weight: 400; text-align: left; vertical-align: top; border-top: 0px; border-right: 0px; border-bottom: 0px; border-left: 0px;"> <table class="divider_block block-1 mobile_hide" width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;"> <tr> <td class="pad" style="padding-left:10px;padding-right:10px;padding-top:30px;"> <div class="alignment" align="center"> <table border="0" cellpadding="0" cellspacing="0" role="presentation" width="100%" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;"> <tr> <td class="divider_inner" style="font-size: 1px; line-height: 1px; border-top: 0px solid #BBBBBB;"> <span>&#8202;</span></td> </tr> </table> </div> </td> </tr> </table> </td> </tr> </tbody> </table> </td> </tr> </tbody> </table> <table class="row row-2" align="center" width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;"> <tbody> <tr> <td> <table class="row-content stack" align="center" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; background-color: #fff; color: #000; width: 640px; margin: 0 auto;" width="640"> <tbody> <tr> <td class="column column-1" width="100%" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; font-weight: 400; text-align: left; vertical-align: top; border-top: 0px; border-right: 0px; border-bottom: 0px; border-left: 0px;"> <table class="divider_block block-1" width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;"> <tr> <td class="pad" style="padding-left:10px;padding-right:10px;padding-top:15px;"> <div class="alignment" align="center"> <table border="0" cellpadding="0" cellspacing="0" role="presentation" width="100%" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;"> <tr> <td class="divider_inner" style="font-size: 1px; line-height: 1px; border-top: 0px solid #BBBBBB;"> <span>&#8202;</span></td> </tr> </table> </div> </td> </tr> </table> <table class="image_block block-2" width="100%" border="0" cellpadding="20" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;"> <tr> <td class="pad"> <div class="alignment" align="center" style="line-height:10px"><img src="https://d1oco4z2z1fhwp.cloudfront.net/templates/default/891/Logo.png" style="display: block; height: auto; border: 0; max-width: 147px; width: 100%;" width="147" alt="Image" title="Image"></div> </td> </tr> </table> <table class="divider_block block-3" width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;"> <tr> <td class="pad" style="padding-left:10px;padding-right:10px;padding-top:18px;"> <div class="alignment" align="center"> <table border="0" cellpadding="0" cellspacing="0" role="presentation" width="100%" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;"> <tr> <td class="divider_inner" style="font-size: 1px; line-height: 1px; border-top: 0px solid #BBBBBB;"> <span>&#8202;</span></td> </tr> </table> </div> </td> </tr> </table> </td> </tr> </tbody> </table> </td> </tr> </tbody> </table> <table class="row row-3" align="center" width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;"> <tbody> <tr> <td> <table class="row-content stack" align="center" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; background-color: #fff; color: #000; width: 640px; margin: 0 auto;" width="640"> <tbody> <tr> <td class="column column-1" width="100%" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; font-weight: 400; text-align: left; vertical-align: top; border-top: 0px; border-right: 0px; border-bottom: 0px; border-left: 0px;"> <table class="image_block block-1" width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;"> <tr> <td class="pad" style="width:100%;"> <div class="alignment" align="center" style="line-height:10px"><img src="https://d1oco4z2z1fhwp.cloudfront.net/templates/default/891/Invite_Illustration.png" style="display: block; height: auto; border: 0; max-width: 640px; width: 100%;" width="640" alt="I'm an image" title="I'm an image"></div> </td> </tr> </table> <table class="paragraph_block block-2" width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; word-break: break-word;"> <tr> <td class="pad" style="padding-bottom:10px;padding-left:40px;padding-right:40px;padding-top:20px;"> <div style="color:#555555;font-family:Trebuchet MS, Lucida Grande, Lucida Sans Unicode, Lucida Sans, Tahoma, sans-serif;font-size:46px;line-height:120%;text-align:center;mso-line-height-alt:55.199999999999996px;"> <p style="margin: 0; word-break: break-word;"><span style="color: #003188;"><strong>OTP Reset Password</strong></span></p> </div> </td> </tr> </table> <table class="paragraph_block block-3" width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; word-break: break-word;"> <tr> <td class="pad" style="padding-bottom:10px;padding-left:40px;padding-right:40px;padding-top:10px;"> <div style="color:#555555;font-family:Trebuchet MS, Lucida Grande, Lucida Sans Unicode, Lucida Sans, Tahoma, sans-serif;font-size:16px;line-height:150%;text-align:center;mso-line-height-alt:24px;"> <p style="margin: 0; word-break: break-word;">We have received a request to reset the password for your account. Below is the OTP code to verify your identity</p> </div> </td> </tr> </table> <table class="divider_block block-4" width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;"> <tr> <td class="pad" style="padding-left:10px;padding-right:10px;padding-top:44px;"> <div class="alignment" align="center"> <table border="0" cellpadding="0" cellspacing="0" role="presentation" width="100%" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;"> <tr> <td class="divider_inner" style="font-size: 1px; line-height: 1px; border-top: 0px solid #BBBBBB;"> <span>&#8202;</span></td> </tr> </table> </div> </td> </tr> </table> <table class="image_block block-5" width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;"> <tr> <td class="pad" style="width:100%;"> <div class="alignment" align="center" style="line-height:10px"><img src="https://d1oco4z2z1fhwp.cloudfront.net/templates/default/891/Divisor_Lines.png" style="display: block; height: auto; border: 0; max-width: 158px; width: 100%;" width="158" alt="Image" title="Image"></div> </td> </tr> </table> <table class="divider_block block-6" width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;"> <tr> <td class="pad" style="padding-left:10px;padding-right:10px;padding-top:35px;"> <div class="alignment" align="center"> <table border="0" cellpadding="0" cellspacing="0" role="presentation" width="100%" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;"> <tr> <td class="divider_inner" style="font-size: 1px; line-height: 1px; border-top: 0px solid #BBBBBB;"> <span>&#8202;</span></td> </tr> </table> </div> </td> </tr> </table> <table class="paragraph_block block-7" width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; word-break: break-word;"> <tr> <td class="pad" style="padding-bottom:10px;padding-left:40px;padding-right:40px;padding-top:20px;"> <div style="color:#555555;font-family:Trebuchet MS, Lucida Grande, Lucida Sans Unicode, Lucida Sans, Tahoma, sans-serif;font-size:24px;line-height:120%;text-align:center;mso-line-height-alt:28.799999999999997px;"> <p style="margin: 0; word-break: break-word;"><span style="color: #003188;"><strong>Your OTP</strong></span></p> </div> </td> </tr> </table> <table class="paragraph_block block-8" width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; word-break: break-word;"> <tr> <td class="pad" style="padding-bottom:10px;padding-left:40px;padding-right:40px;padding-top:20px;"> <div style="color:#555555;font-family:Trebuchet MS, Lucida Grande, Lucida Sans Unicode, Lucida Sans, Tahoma, sans-serif;font-size:45px;line-height:120%;text-align:center;mso-line-height-alt:54px;"> <p style="margin: 0; word-break: break-word;"> <strong>` + otp + `</strong></p> </div> </td> </tr> </table> <table class="paragraph_block block-9" width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; word-break: break-word;"> <tr> <td class="pad" style="padding-bottom:10px;padding-left:40px;padding-right:40px;padding-top:10px;"> <div style="color:#555555;font-family:Trebuchet MS, Lucida Grande, Lucida Sans Unicode, Lucida Sans, Tahoma, sans-serif;font-size:16px;line-height:150%;text-align:center;mso-line-height-alt:24px;"> <p style="margin: 0; word-break: break-word;">Please use this code within 15 minutes. Do not share this code with anyone, including our team. If you did not initiate this request, kindly ignore this message.</p> </div> </td> </tr> </table> <table class="divider_block block-10" width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;"> <tr> <td class="pad" style="padding-left:10px;padding-right:10px;padding-top:40px;"> <div class="alignment" align="center"> <table border="0" cellpadding="0" cellspacing="0" role="presentation" width="100%" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;"> <tr> <td class="divider_inner" style="font-size: 1px; line-height: 1px; border-top: 0px dotted #BBBBBB;"> <span>&#8202;</span></td> </tr> </table> </div> </td> </tr> </table> <table class="paragraph_block block-11" width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; word-break: break-word;"> <tr> <td class="pad" style="padding-bottom:10px;padding-left:40px;padding-right:40px;padding-top:10px;"> <div style="color:#555555;font-family:Trebuchet MS, Lucida Grande, Lucida Sans Unicode, Lucida Sans, Tahoma, sans-serif;font-size:16px;line-height:150%;text-align:center;mso-line-height-alt:24px;"> <p style="margin: 0; word-break: break-word;">Thank you,</p> <p style="margin: 0; word-break: break-word;"> E-Complaint Team</p> </div> </td> </tr> </table> <table class="divider_block block-12" width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;"> <tr> <td class="pad" style="padding-left:10px;padding-right:10px;padding-top:40px;"> <div class="alignment" align="center"> <table border="0" cellpadding="0" cellspacing="0" role="presentation" width="100%" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;"> <tr> <td class="divider_inner" style="font-size: 1px; line-height: 1px; border-top: 0px dotted #BBBBBB;"> <span>&#8202;</span></td> </tr> </table> </div> </td> </tr> </table> </td> </tr> </tbody> </table> </td> </tr> </tbody> </table> <table class="row row-4" align="center" width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;"> <tbody> <tr> <td> <table class="row-content stack" align="center" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; background-color: #fff; color: #000; width: 640px; margin: 0 auto;" width="640"> <tbody> <tr> <td class="column column-1" width="100%" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; font-weight: 400; text-align: left; border-top: 1px solid #E5EAF3; vertical-align: top; border-right: 0px; border-bottom: 0px; border-left: 0px;"> <table class="divider_block block-1" width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;"> <tr> <td class="pad" style="padding-left:10px;padding-right:10px;padding-top:35px;"> <div class="alignment" align="center"> <table border="0" cellpadding="0" cellspacing="0" role="presentation" width="100%" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;"> <tr> <td class="divider_inner" style="font-size: 1px; line-height: 1px; border-top: 0px solid #BBBBBB;"> <span>&#8202;</span></td> </tr> </table> </div> </td> </tr> </table> <table class="image_block block-2" width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;"> <tr> <td class="pad" style="padding-bottom:15px;padding-left:20px;padding-right:20px;padding-top:20px;width:100%;"> <div class="alignment" align="center" style="line-height:10px"><img src="https://d1oco4z2z1fhwp.cloudfront.net/templates/default/891/Logo.png" style="display: block; height: auto; border: 0; max-width: 147px; width: 100%;" width="147" alt="Image" title="Image"></div> </td> </tr> </table> <table class="social_block block-3" width="100%" border="0" cellpadding="10" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;"> <tr> <td class="pad"> <div class="alignment" align="center"> <table class="social-table" width="184px" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; display: inline-block;"> <tr> <td style="padding:0 7px 0 7px;"><a href="https://www.facebook.com" target="_blank"><img src="https://app-rsrc.getbee.io/public/resources/social-networks-icon-sets/circle-blue/facebook@2x.png" width="32" height="32" alt="Facebook" title="Facebook" style="display: block; height: auto; border: 0;"></a> </td> <td style="padding:0 7px 0 7px;"><a href="https://www.twitter.com" target="_blank"><img src="https://app-rsrc.getbee.io/public/resources/social-networks-icon-sets/circle-blue/twitter@2x.png" width="32" height="32" alt="Twitter" title="Twitter" style="display: block; height: auto; border: 0;"></a> </td> <td style="padding:0 7px 0 7px;"><a href="https://www.instagram.com" target="_blank"><img src="https://app-rsrc.getbee.io/public/resources/social-networks-icon-sets/circle-blue/instagram@2x.png" width="32" height="32" alt="Instagram" title="Instagram" style="display: block; height: auto; border: 0;"></a> </td> <td style="padding:0 7px 0 7px;"><a href="https://www.linkedin.com" target="_blank"><img src="https://app-rsrc.getbee.io/public/resources/social-networks-icon-sets/circle-blue/linkedin@2x.png" width="32" height="32" alt="LinkedIn" title="LinkedIn" style="display: block; height: auto; border: 0;"></a> </td> </tr> </table> </div> </td> </tr> </table> <table class="divider_block block-4" width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;"> <tr> <td class="pad" style="padding-left:10px;padding-right:10px;padding-top:40px;"> <div class="alignment" align="center"> <table border="0" cellpadding="0" cellspacing="0" role="presentation" width="100%" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;"> <tr> <td class="divider_inner" style="font-size: 1px; line-height: 1px; border-top: 0px solid #BBBBBB;"> <span>&#8202;</span></td> </tr> </table> </div> </td> </tr> </table> </td> </tr> </tbody> </table> </td> </tr> </tbody> </table> <table class="row row-5" align="center" width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;"> <tbody> <tr> <td> <table class="row-content stack" align="center" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; color: #000; width: 640px; margin: 0 auto;" width="640"> <tbody> <tr> <td class="column column-1" width="100%" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; font-weight: 400; text-align: left; vertical-align: top; border-top: 0px; border-right: 0px; border-bottom: 0px; border-left: 0px;"> <table class="divider_block block-1 mobile_hide" width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;"> <tr> <td class="pad" style="padding-left:10px;padding-right:10px;padding-top:30px;"> <div class="alignment" align="center"> <table border="0" cellpadding="0" cellspacing="0" role="presentation" width="100%" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;"> <tr> <td class="divider_inner" style="font-size: 1px; line-height: 1px; border-top: 0px solid #BBBBBB;"> <span>&#8202;</span></td> </tr> </table> </div> </td> </tr> </table> </td> </tr> </tbody> </table> </td> </tr> </tbody> </table> <table class="row row-6" align="center" width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; background-color: #ffffff;"> <tbody> <tr> <td> <table class="row-content stack" align="center" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; background-color: #fff; color: #000; width: 640px; margin: 0 auto;" width="640"> <tbody> <tr> <td class="column column-1" width="100%" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; font-weight: 400; text-align: left; padding-bottom: 5px; padding-top: 5px; vertical-align: top; border-top: 0px; border-right: 0px; border-bottom: 0px; border-left: 0px;"> <table class="icons_block block-1" width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;"> <tr> <td class="pad" style="vertical-align: middle; color: #1e0e4b; font-family: 'Inter', sans-serif; font-size: 15px; padding-bottom: 5px; padding-top: 5px; text-align: center;"> <table width="100%" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;"> <tr> <td class="alignment" style="vertical-align: middle; text-align: center;"> <!--[if vml]><table align="left" cellpadding="0" cellspacing="0" role="presentation" style="display:inline-block;padding-left:0px;padding-right:0px;mso-table-lspace: 0pt;mso-table-rspace: 0pt;"><![endif]--> <!--[if !vml]><!--> <table class="icons-inner" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; display: inline-block; margin-right: -4px; padding-left: 0px; padding-right: 0px;" cellpadding="0" cellspacing="0" role="presentation"><!--<![endif]--> <tr> <td style="vertical-align: middle; text-align: center; padding-top: 5px; padding-bottom: 5px; padding-left: 5px; padding-right: 6px;"> <a href="http://designedwithbeefree.com/" target="_blank" style="text-decoration: none;"><img class="icon" alt="Beefree Logo" src="https://d1oco4z2z1fhwp.cloudfront.net/assets/Beefree-logo.png" height="32" width="34" align="center" style="display: block; height: auto; margin: 0 auto; border: 0;"></a> </td> <td style="font-family: 'Inter', sans-serif; font-size: 15px; color: #1e0e4b; vertical-align: middle; letter-spacing: undefined; text-align: center;"> <a href="http://designedwithbeefree.com/" target="_blank" style="color: #1e0e4b; text-decoration: none;">Designed with Beefree</a></td> </tr> </table> </td> </tr> </table> </td> </tr> </table> </td> </tr> </tbody> </table> </td> </tr> </tbody> </table> </td> </tr> </tbody> </table><!-- End --> </body> </html>`

	to := []string{email}

	m := gomail.NewMessage()
	m.SetHeader("From", mail.Username)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", "GovComplaint Notification")
	m.SetBody("text/html", body)

	Port, _ := strconv.Atoi(mail.Port)

	dialer := gomail.NewDialer(
		mail.Host,
		Port,
		mail.Username,
		mail.Password,
	)

	if err := dialer.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
