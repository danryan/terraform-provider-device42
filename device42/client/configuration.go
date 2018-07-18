package client

import "net/http"

type contextKey string

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().
func (c contextKey) String() string {
	return "auth " + string(c)
}

var (
	// ContextBasicAuth key
	ContextBasicAuth = contextKey("basic")
)

// BasicAuth type
type BasicAuth struct {
	Username string
	Password string
}

// Configuration for the client
type Configuration struct {
	BasePath      string
	Host          string
	Scheme        string
	DefaultHeader map[string]string
	UserAgent     string
	HTTPClient    *http.Client
}

func NewConfiguration() *Configuration {
	cfg := &Configuration{
		BasePath:      "http://api.device42.com/v1",
		DefaultHeader: make(map[string]string),
		UserAgent:     "Device42/1.0.0/go",
	}
	return cfg
}

func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}
