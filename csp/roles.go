package csp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetAllOrders - Returns all user's order
func (c *CspClient) GetAllRoles() (*[]Role, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/csp/gateway/iam-roles-mgmt/api/services/%s/roles", c.HostURL, c.ServiceDefinitionID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	roles := []Role{}
	err = json.Unmarshal(body, &roles)
	if err != nil {
		return nil, err
	}

	return &roles, nil
}

// GetOrder - Returns a specifc order
func (c *CspClient) GetRole(roleName string) (*Role, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/csp/gateway/iam-roles-mgmt/api/services/%s/roles/%s", c.HostURL, c.ServiceDefinitionID, roleName), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	role := Role{}
	err = json.Unmarshal(body, &role)
	if err != nil {
		return nil, err
	}

	return &role, nil
}

// CreateOrder - Create new order
func (c *CspClient) CreateRole(role Role) (*Role, error) {
	rb, err := json.Marshal(role)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/csp/gateway/iam-roles-mgmt/api/services/%s/roles", c.HostURL, c.ServiceDefinitionID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	newRole := Role{}
	err = json.Unmarshal(body, &newRole)
	if err != nil {
		return nil, err
	}

	return &newRole, nil
}

// UpdateOrder - Updates an order
func (c *CspClient) UpdateRole(roleName string, role Role) (*Role, error) {
	rb, err := json.Marshal(role)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/csp/gateway/iam-roles-mgmt/api/services/%s/roles/%s", c.HostURL, c.ServiceDefinitionID, roleName), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	updatedRole := Role{}
	err = json.Unmarshal(body, &updatedRole)
	if err != nil {
		return nil, err
	}

	return &updatedRole, nil
}

// DeleteOrder - Deletes an order
func (c *CspClient) DeleteRole(roleName string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/csp/gateway/iam-roles-mgmt/api/services/%s/roles/%s", c.HostURL, c.ServiceDefinitionID, roleName), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	/*if string(body) != "Deleted Role" {
		return errors.New(string(body))
	}*/

	return nil
}
