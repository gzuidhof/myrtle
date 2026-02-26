package smtpclient

import (
	"fmt"
	"net/mail"
	"net/smtp"
	"strings"

	"github.com/jordan-wright/email"
)

type Config struct {
	Host        string
	Port        int
	Username    string
	Password    string
	FromName    string
	FromAddress string
}

type Client struct {
	host string
	port int
	auth smtp.Auth
	from string
}

func New(config Config) (*Client, error) {
	host := strings.TrimSpace(config.Host)
	if host == "" {
		return nil, fmt.Errorf("smtp host is required")
	}

	if config.Port <= 0 {
		return nil, fmt.Errorf("smtp port must be greater than zero")
	}

	fromAddress := strings.TrimSpace(config.FromAddress)
	if fromAddress == "" {
		return nil, fmt.Errorf("smtp from_address is required")
	}

	if _, err := mail.ParseAddress(fromAddress); err != nil {
		return nil, fmt.Errorf("invalid smtp from_address: %w", err)
	}

	from := fromAddress
	if fromName := strings.TrimSpace(config.FromName); fromName != "" {
		from = (&mail.Address{Name: fromName, Address: fromAddress}).String()
	}

	var auth smtp.Auth
	if strings.TrimSpace(config.Username) != "" {
		auth = smtp.PlainAuth("", config.Username, config.Password, host)
	}

	return &Client{
		host: host,
		port: config.Port,
		auth: auth,
		from: from,
	}, nil
}

func (client *Client) Send(to, subject, htmlBody, markdownBody string) error {
	if client == nil {
		return fmt.Errorf("smtp client is nil")
	}

	recipient := strings.TrimSpace(to)
	if recipient == "" {
		return fmt.Errorf("recipient is required")
	}

	if _, err := mail.ParseAddress(recipient); err != nil {
		return fmt.Errorf("invalid recipient address: %w", err)
	}

	message := email.NewEmail()
	message.From = client.from
	message.To = []string{recipient}
	message.Subject = strings.TrimSpace(subject)
	message.HTML = []byte(htmlBody)
	message.Text = []byte(markdownBody)

	return message.Send(fmt.Sprintf("%s:%d", client.host, client.port), client.auth)
}
