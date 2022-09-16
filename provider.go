package main

import (
	"github.com/dtcookie/terraform-provider-sample/resources/records"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"sample_record": records.Resource(),
		},
	}
}
