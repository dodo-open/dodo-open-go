package client

import (
	"errors"
	"time"
)

type OptionHandler func(*Config) error

// WithBaseApi Customize DoDoBot Api base host
func WithBaseApi(baseApi string) OptionHandler {
	return func(config *Config) error {
		if baseApi == "" {
			return errors.New("invalid BaseApi (empty string detected)")
		}
		config.BaseApi = baseApi
		return nil
	}
}

// WithTimeout Customize RestyClient request timeout
func WithTimeout(duration time.Duration) OptionHandler {
	return func(config *Config) error {
		config.Timeout = duration
		return nil
	}
}

// WithDebugMode Toggle debug mode
func WithDebugMode(flag bool) OptionHandler {
	return func(config *Config) error {
		config.IsDebug = flag
		return nil
	}
}
