// Package email provides a flexible and modular email sending system
// that supports multiple providers (SMTP, Zoho, Gmail, etc.)
package email

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
)

// Provider defines the interface that all email providers must implement
type Provider interface {
	// GetAuth returns the authentication mechanism for the provider
	GetAuth() smtp.Auth
	// GetAddress returns the server address (host:port)
	GetAddress() string
	// GetFrom returns the sender email address
	GetFrom() string
}

// Client is the main email client that works with any Provider
type Client struct {
	provider Provider
}

// NewClient creates a new email client with the specified provider
func NewClient(provider Provider) *Client {
	return &Client{
		provider: provider,
	}
}

// EmailMessage represents an email to be sent
type EmailMessage struct {
	To      []string
	Subject string
	Body    string
}

// SendEmail sends an email using the configured provider
func (c *Client) SendEmail(msg *EmailMessage) error {
	auth := c.provider.GetAuth()
	addr := c.provider.GetAddress()
	from := c.provider.GetFrom()

	// Prepare the email message with proper headers
	mime := "MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\r\n\r\n"
	message := []byte(fmt.Sprintf("Subject: %s\r\n%s%s", msg.Subject, mime, msg.Body))

	return smtp.SendMail(addr, auth, from, msg.To, message)
}

// SendTemplateEmail sends an email using a template
func (c *Client) SendTemplateEmail(to []string, subject string, tmpl *template.Template, data interface{}) error {
	var buffer bytes.Buffer

	// Execute the template with the provided data
	if err := tmpl.Execute(&buffer, data); err != nil {
		return fmt.Errorf("template execution failed: %w", err)
	}

	msg := &EmailMessage{
		To:      to,
		Subject: subject,
		Body:    buffer.String(),
	}

	return c.SendEmail(msg)
}

// SendSingleEmail is a convenience method for sending to a single recipient
func (c *Client) SendSingleEmail(to string, subject string, tmpl *template.Template, data interface{}) error {
	return c.SendTemplateEmail([]string{to}, subject, tmpl, data)
}

// ===== Provider Implementations =====

// SMTPConfig holds common SMTP configuration
type SMTPConfig struct {
	Host     string
	Port     string
	Username string
	Password string
}

// Validate checks if the configuration is valid
func (c *SMTPConfig) Validate() error {
	if c.Host == "" {
		return fmt.Errorf("host is required")
	}
	if c.Port == "" {
		return fmt.Errorf("port is required")
	}
	if c.Username == "" {
		return fmt.Errorf("username is required")
	}
	if c.Password == "" {
		return fmt.Errorf("password is required")
	}
	return nil
}

// GenericSMTPProvider is a generic SMTP provider implementation
type GenericSMTPProvider struct {
	config SMTPConfig
}

// NewGenericSMTPProvider creates a new generic SMTP provider
func NewGenericSMTPProvider(config SMTPConfig) (*GenericSMTPProvider, error) {
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}
	return &GenericSMTPProvider{config: config}, nil
}

func (p *GenericSMTPProvider) GetAuth() smtp.Auth {
	return smtp.PlainAuth("", p.config.Username, p.config.Password, p.config.Host)
}

func (p *GenericSMTPProvider) GetAddress() string {
	return p.config.Host + ":" + p.config.Port
}

func (p *GenericSMTPProvider) GetFrom() string {
	return p.config.Username
}

// GmailProvider is a specialized provider for Gmail
type GmailProvider struct {
	*GenericSMTPProvider
}

// NewGmailProvider creates a new Gmail provider with default settings
func NewGmailProvider(username, password string) (*GmailProvider, error) {
	config := SMTPConfig{
		Host:     "smtp.gmail.com",
		Port:     "587",
		Username: username,
		Password: password,
	}
	
	generic, err := NewGenericSMTPProvider(config)
	if err != nil {
		return nil, err
	}
	
	return &GmailProvider{GenericSMTPProvider: generic}, nil
}

// ZohoProvider is a specialized provider for Zoho Mail
type ZohoProvider struct {
	*GenericSMTPProvider
}

// NewZohoProvider creates a new Zoho provider with default settings
func NewZohoProvider(username, password string) (*ZohoProvider, error) {
	config := SMTPConfig{
		Host:     "smtp.zoho.com",
		Port:     "587",
		Username: username,
		Password: password,
	}
	
	generic, err := NewGenericSMTPProvider(config)
	if err != nil {
		return nil, err
	}
	
	return &ZohoProvider{GenericSMTPProvider: generic}, nil
}

// ===== Example: Custom Provider Implementation =====

// OutlookProvider is an example of adding a new provider
type OutlookProvider struct {
	*GenericSMTPProvider
}

// NewOutlookProvider creates a new Outlook provider
func NewOutlookProvider(username, password string) (*OutlookProvider, error) {
	config := SMTPConfig{
		Host:     "smtp.office365.com",
		Port:     "587",
		Username: username,
		Password: password,
	}
	
	generic, err := NewGenericSMTPProvider(config)
	if err != nil {
		return nil, err
	}
	
	return &OutlookProvider{GenericSMTPProvider: generic}, nil
}
