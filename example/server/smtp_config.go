package server

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/gzuidhof/myrtle/example/server/smtpclient"
)

const defaultSMTPConfigPath = "example/server/smtp.config.json"

type smtpSettings struct {
	Client    *smtpclient.Client
	DefaultTo string
}

type smtpFileConfig struct {
	Host        string `json:"host"`
	Port        int    `json:"port"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	FromName    string `json:"from_name"`
	FromAddress string `json:"from_address"`
	DefaultTo   string `json:"default_to"`
}

func loadSMTPSettings() (*smtpSettings, error) {
	configPath := strings.TrimSpace(os.Getenv("MYRTLE_SMTP_CONFIG"))
	if configPath == "" {
		configPath = defaultSMTPConfigPath
	}

	content, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}

		return nil, err
	}

	var config smtpFileConfig
	if err := json.Unmarshal(content, &config); err != nil {
		return nil, fmt.Errorf("invalid smtp config json: %w", err)
	}

	client, err := smtpclient.New(smtpclient.Config{
		Host:        config.Host,
		Port:        config.Port,
		Username:    config.Username,
		Password:    config.Password,
		FromName:    config.FromName,
		FromAddress: config.FromAddress,
	})
	if err != nil {
		return nil, fmt.Errorf("invalid smtp config: %w", err)
	}

	return &smtpSettings{
		Client:    client,
		DefaultTo: strings.TrimSpace(config.DefaultTo),
	}, nil
}
