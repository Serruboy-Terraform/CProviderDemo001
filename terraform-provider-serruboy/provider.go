package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"public_ip": public_ip(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"public_data": dataSource(),
		},
	}
}
