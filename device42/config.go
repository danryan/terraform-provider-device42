package device42

import (
	apiclient "github.com/danryan/terraform-provider-device42/device42/client"
)

// Config defines the configuration options for the Device42 client
type Config struct {
	// The Device42 API endpoint
	APIEndpoint string

	// The Device42 Username
	Username string

	// The Device42 Password
	Password string
}

type Client struct {
	config Config

	client *apiclient.APIClient
}

// NewClient returns a new client config
func (c *Config) NewClient() (*Client, error) {
	cfg := &apiclient.Configuration{
		BasePath: c.APIEndpoint,
	}

	apiclient := apiclient.NewAPIClient(cfg)

	client := Client{
		config: *c,
		client: apiclient,
	}

	return &client, nil
}

func (c *Client) Client() *apiclient.APIClient {
	return c.client
}
