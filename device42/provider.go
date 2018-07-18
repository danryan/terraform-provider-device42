package device42

import (
	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a schema.Provider for Device42
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_endpoint": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("DEVICE42_API_ENDPOINT", ""),
				Description: "The API endpoint of the Device42 server.",
			},
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("DEVICE42_USERNAME", ""),
				Description: "The username for Device42 server authentication",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("DEVICE42_PASSWORD", ""),
				Description: "The password for Device42 server authentication",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"device42_devices": dataSourceDevices(),
		},

		ResourcesMap: map[string]*schema.Resource{},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(data *schema.ResourceData) (interface{}, error) {
	config := Config{
		APIEndpoint: data.Get("api_endpoint").(string),
		Username:    data.Get("username").(string),
		Password:    data.Get("password").(string),
	}

	client, err := config.NewClient()
	if err != nil {
		return nil, errwrap.Wrapf("error initializing Device42 client: {{err}}", err)
	}

	return client, nil
}
