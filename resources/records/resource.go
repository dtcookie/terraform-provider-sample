package records

import (
	"context"

	"github.com/dtcookie/terraform-provider-sample/endpoints/records"
	"github.com/dtcookie/terraform-provider-sample/endpoints/repo"
	"github.com/dtcookie/terraform-provider-sample/tfh"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Resource() *schema.Resource {
	return &schema.Resource{
		Schema:        new(records.Record).Schema(),
		CreateContext: Create,
		UpdateContext: Update,
		ReadContext:   Read,
		DeleteContext: Delete,
		Importer:      &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},
	}
}

var storage = repo.New()

func Create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	id := uuid.New().String()
	d.SetId(id)
	record := new(records.Record)
	if err := record.UnmarshalHCL(ctx, tfh.New(d)); err != nil {
		return diag.FromErr(err)
	}
	if err := storage.Put(id, record); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)
	return Read(ctx, d, m)
}

func Update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	record := new(records.Record)
	if err := record.UnmarshalHCL(ctx, tfh.New(d)); err != nil {
		return diag.FromErr(err)
	}
	if err := storage.Put(d.Id(), record); err != nil {
		return diag.FromErr(err)
	}
	return Read(ctx, d, m)
}

func Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	id := d.Id()
	record := new(records.Record)
	if err := storage.Get(id, record); err != nil {
		return diag.FromErr(err)
	}
	marshalled, err := record.MarshalHCL(ctx)
	if err != nil {
		return diag.FromErr(err)
	}
	for k, v := range marshalled {
		d.Set(k, v)
	}
	return diag.Diagnostics{}
}

func Delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	storage.Delete(d.Id())
	return diag.Diagnostics{}
}
