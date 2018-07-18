package main

import (
	"github.com/danryan/terraform-provider-device42/device42"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return device42.Provider()
		},
	})
}
