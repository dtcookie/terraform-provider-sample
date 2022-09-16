package records

import (
	"context"

	"github.com/dtcookie/terraform-provider-sample/tfh"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Record struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (me *Record) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Description: "The name of the record",
			Required:    true,
		},
		"value": {
			Type:        schema.TypeString,
			Description: "The value of the record",
			Required:    true,
		},
	}
}

func (me Record) MarshalHCL(ctx context.Context) (map[string]interface{}, error) {
	return map[string]interface{}{
		"name":  me.Name,
		"value": me.Value,
	}, nil
}

func (me *Record) UnmarshalHCL(ctx context.Context, d tfh.ResourceData) error {
	if name, ok := d.GetOk("name"); ok {
		me.Name = name.(string)
	}
	if value, ok := d.GetOk("value"); ok {
		me.Value = value.(string)
	}
	return nil
}
