package model

import (
	"github.com/crazy-max/diun/pkg/registry"
)

// NotifEntry represents a notification entry
type NotifEntry struct {
	Status   ImageStatus       `json:"status,omitempty"`
	Provider string            `json:"provider,omitempty"`
	Image    registry.Image    `json:"image,omitempty"`
	Manifest registry.Manifest `json:"manifest,omitempty"`
}

// Notif holds data necessary for notification configuration
type Notif struct {
	Mail     NotifMail     `yaml:"mail,omitempty"`
	Slack    NotifSlack    `yaml:"slack,omitempty"`
	Telegram NotifTelegram `yaml:"telegram,omitempty"`
	Webhook  NotifWebhook  `yaml:"webhook,omitempty"`
}

// NotifMail holds mail notification configuration details
type NotifMail struct {
	Enable             bool   `yaml:"enable,omitempty"`
	Host               string `yaml:"host,omitempty"`
	Port               int    `yaml:"port,omitempty"`
	SSL                bool   `yaml:"ssl,omitempty"`
	InsecureSkipVerify bool   `yaml:"insecure_skip_verify,omitempty"`
	Username           string `yaml:"username,omitempty"`
	UsernameFile       string `yaml:"username_file,omitempty"`
	Password           string `yaml:"password,omitempty"`
	PasswordFile       string `yaml:"password_file,omitempty"`
	From               string `yaml:"from,omitempty"`
	To                 string `yaml:"to,omitempty"`
}

// NotifSlack holds slack notification configuration details
type NotifSlack struct {
	Enable     bool   `yaml:"enable,omitempty"`
	WebhookURL string `yaml:"webhook_url,omitempty"`
}

// NotifTelegram holds Telegram notification configuration details
type NotifTelegram struct {
	Enable   bool    `yaml:"enable,omitempty"`
	BotToken string  `yaml:"token,omitempty"`
	ChatIDs  []int64 `yaml:"chat_ids,omitempty"`
}

// NotifWebhook holds webhook notification configuration details
type NotifWebhook struct {
	Enable   bool              `yaml:"enable,omitempty"`
	Endpoint string            `yaml:"endpoint,omitempty"`
	Method   string            `yaml:"method,omitempty"`
	Headers  map[string]string `yaml:"headers,omitempty"`
	Timeout  int               `yaml:"timeout,omitempty"`
}
