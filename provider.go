package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/j4ng5y/terraform-provider-tenableio/resources/resources/platform"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"access_key": {
				Type:        schema.TypeString,
				Required:    true,
				Default:     "",
				Description: descriptions["access_key"],
			},
			"secret_key": {
				Type:        schema.TypeString,
				Required:    true,
				Default:     "",
				Description: descriptions["secret_key"],
			},
		},
		DataSourcesMap: map[string]*schema.Resource{},
		ResourcesMap: map[string]*schema.Resource{
			"tenableio_user": resources.ResourceUser(),
		},
	}
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"access_key": "The access key for tenable.io API operations. You can retrieve this\n" +
			"from the 'User' section of the tenable.io console.",
		"secret_key": "The secret key for tenable.io API operations. You can retrieve this\n" +
			"from the 'User' section of the tenable.io console.",
	}
}
