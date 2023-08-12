package csp

type Role struct {
	Name                    string   `json:"name"`
	DisplayName             string   `json:"displayName"`
	Description             string   `json:"description"`
	OnAccess                bool     `json:"onAccess"`
	Visible                 bool     `json:"visible"`
	Type                    string   `json:"type"`
	Composable              bool     `json:"composable"`
	DisallowedResourceTypes []string `json:"disallowedResourceTypes"`
	Bundled                 bool     `json:"bundled"`
	Permissions             []string `json:"permissions"`
}
