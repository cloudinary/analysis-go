package cldconfig

import (
	"errors"
	"net/url"
	"os"
)

type Configuration struct {
	Cloud Cloud
}

type Cloud struct {
	CloudName string
	APIKey    string
	APISecret string
}

func New() (*Configuration, error) {
	cldURLStr := os.Getenv("CLOUDINARY_URL")
	if cldURLStr == "" {
		return &Configuration{}, nil
	}
	parsed, err := url.Parse(cldURLStr)
	if err != nil {
		return nil, err
	}
	if parsed.Host == "" {
		return nil, errors.New("invalid CLOUDINARY_URL: missing cloud name")
	}
	if parsed.User == nil {
		return nil, errors.New("invalid CLOUDINARY_URL: missing credentials")
	}
	apiKey := parsed.User.Username()
	apiSecret, hasSecret := parsed.User.Password()
	if apiKey == "" || !hasSecret || apiSecret == "" {
		return nil, errors.New("invalid CLOUDINARY_URL: incomplete API credentials")
	}
	return &Configuration{
		Cloud: Cloud{
			CloudName: parsed.Host,
			APIKey:    apiKey,
			APISecret: apiSecret,
		},
	}, nil
}
