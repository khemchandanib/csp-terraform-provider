// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package csp

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCspRole() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRoleCreate,
		ReadContext:   resourceRoleRead,
		UpdateContext: resourceRoleUpdate,
		DeleteContext: resourceRoleDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"on_access": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"visible": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"composable": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"bundled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"disallowed_resource_types": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"permissions": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
		},
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceRoleCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*CspClient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	cspRole := Role{}
	cspRole.Name = d.Get("name").(string)
	cspRole.DisplayName = d.Get("display_name").(string)
	cspRole.Description = d.Get("description").(string)
	cspRole.Type = d.Get("type").(string)
	cspRole.Visible = d.Get("visible").(bool)
	cspRole.OnAccess = d.Get("on_access").(bool)
	cspRole.Composable = d.Get("composable").(bool)
	// Get the disallowed_resource_types attribute
	disallowedResourceTypesSet := d.Get("disallowed_resource_types").(*schema.Set)
	disallowedResourceTypes := make([]string, disallowedResourceTypesSet.Len())
	for i, v := range disallowedResourceTypesSet.List() {
		disallowedResourceTypes[i] = v.(string)
	}
	cspRole.DisallowedResourceTypes = disallowedResourceTypes

	o, err := c.CreateRole(cspRole)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(o.Name)

	resourceRoleRead(ctx, d, m)

	return diags
}

func resourceRoleRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*CspClient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	roleName := d.Id()

	role, err := c.GetRole(roleName)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("name", role.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("display_name", role.DisplayName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", role.Description); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("visible", role.Visible); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("type", role.Type); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("on_access", role.OnAccess); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("composable", role.Composable); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("disallowed_resource_types", role.DisallowedResourceTypes); err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceRoleUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*CspClient)

	roleName := d.Id()

	if d.HasChanges("name", "display_name", "description", "visible", "type", "on_access", "composable", "disallowed_resource_types") {
		cspRole := Role{}
		cspRole.Name = d.Get("name").(string)
		cspRole.DisplayName = d.Get("display_name").(string)
		cspRole.Description = d.Get("description").(string)
		cspRole.Type = d.Get("type").(string)
		cspRole.Visible = d.Get("visible").(bool)
		cspRole.OnAccess = d.Get("on_access").(bool)
		cspRole.Composable = d.Get("composable").(bool)
		// Get the disallowed_resource_types attribute
		disallowedResourceTypesSet := d.Get("disallowed_resource_types").(*schema.Set)
		disallowedResourceTypes := make([]string, disallowedResourceTypesSet.Len())
		for i, v := range disallowedResourceTypesSet.List() {
			disallowedResourceTypes[i] = v.(string)
		}
		cspRole.DisallowedResourceTypes = disallowedResourceTypes

		_, err := c.UpdateRole(roleName, cspRole)
		if err != nil {
			return diag.FromErr(err)
		}

		d.Set("last_updated", time.Now().Format(time.RFC850))
	}

	return resourceRoleRead(ctx, d, m)
}

func resourceRoleDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*CspClient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	roleName := d.Id()

	err := c.DeleteRole(roleName)
	if err != nil {
		return diag.FromErr(err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

/*func flattenRole(role Role) []interface{} {
	c := make(map[string]interface{})
	c["name"] = role.Name
	c["displayName"] = role.DisplayName
	c["description"] = role.Description
	c["visible"] = role.Visible
	c["type"] = role.Type
	c["onaccess"] = role.OnAccess
	c["composable"] = role.Composable

	return []interface{}{c}
}*/

//"name","displayName","description","visible","type","onaccess","composable"
