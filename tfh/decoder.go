package tfh

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ResourceData interface {
	GetOk(key string) (interface{}, bool)
}

type decoder struct {
	d      *schema.ResourceData
	prefix string
}

func (me *decoder) GetOk(key string) (interface{}, bool) {
	if len(me.prefix) == 0 {
		tflog.Info(context.Background(), fmt.Sprintf("GetOk(%s)", key))
		return me.d.GetOk(key)
	}
	tflog.Info(context.Background(), fmt.Sprintf("GetOk(%s.%s)", me.prefix, key))
	return me.d.GetOk(fmt.Sprintf("%s.%s", me.prefix, key))
}

func New(d ResourceData, keys ...string) ResourceData {
	if d == nil {
		panic("parent decoder must not be null")
	}
	switch t := d.(type) {
	case *schema.ResourceData:
		return &decoder{t, strings.Join(keys, ".")}
	case *decoder:
		if len(t.prefix) == 0 {
			return &decoder{t.d, strings.Join(keys, ".")}
		} else if len(keys) == 0 {
			return &decoder{t.d, t.prefix}
		}
		return New(t.d, append([]string{t.prefix}, keys...)...)
	default:
		panic(fmt.Sprintf("unsupported decoder of type %T", d))
	}
}
