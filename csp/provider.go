// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package csp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"cspurl": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CSP_URL", nil),
			},
			"token": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("TOKEN", nil),
			},
			"servicedefinitionid": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SVC_DEF_ID", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"csp_role": resourceCspRole(),
		},
		/*DataSourcesMap: map[string]*schema.Resource{
			"csp_order":       dataSourceCspRole(),
		},*/
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	token := d.Get("token").(string)
	serviceDefinitionId := d.Get("servicedefinitionid").(string)

	var host *string

	hVal, ok := d.GetOk("cspurl")
	if ok {
		tempHost := hVal.(string)
		host = &tempHost
	}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if token != "" && serviceDefinitionId != "" {
		c, err := NewCspClient(host, &token, &serviceDefinitionId)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to create CSP client",
				Detail:   "Unable to authenticate user for authenticated CSP client",
			})

			return nil, diags
		}

		return c, diags
	}

	c, err := NewCspClient(host, nil, nil)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create CSP client",
			Detail:   "Unable to create anonymous CSP client",
		})
		return nil, diags
	}

	return c, diags
}
