package device42

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func dataSourceDevices() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDevicesRead,
		Schema: map[string]*schema.Schema{
			"type": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{
						"physical",
						"virtual",
						"blade",
						"other",
						"cluster",
						"unknown",
					}, false),
				},
			},
			"limit": {
				Type:     schema.TypeInt,
				Default:  10,
				Optional: true,
			},
			"offset": {
				Type:     schema.TypeInt,
				Default:  0,
				Optional: true,
			},
		},
	}
}

func dataSourceDevicesRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)

}
