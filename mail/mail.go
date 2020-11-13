package mail

import (
	"log"
	"net/smtp"
	"os"
	"strings"
)

type database struct {
	username string
	password string
	host     string
	port     string
	db       string
}

type mailer struct {
	email    string
	siteName string
}

var db1 database

//SendMail ...
func SendMail() {

	host := os.Getenv("host")
	hostPort := os.Getenv("hostPort")
	emailSender := os.Getenv("emailSender")
	password := os.Getenv("password")
	emailReceiver := os.Getenv("emailReceiver")

	//receiver
	to := []string{emailReceiver}

	// fmt.Println("to", to, len(mailList))

	toHeadder := strings.Join(to, ",")

	// fmt.Println(getEmails())

	if len(to) > 0 {
		// Choose auth method and set it up
		auth := smtp.PlainAuth("", emailSender, password, host)

		// Here we do it all: connect to our server, set up a message and send it

		msg := []byte("To:" + toHeadder + "\r\n" +
			"Subject: Starting_mail_sender \r\n" +
			"\r\n" +
			"Bonjour, \n\n" +
			" ceci est un test" +
			"\r\n")

		err := smtp.SendMail(host+":"+hostPort, auth, emailSender, to, msg)
		if err != nil {
			log.Fatal(err)
		}

		// fmt.Println("email sent")
	}

}
