/*
 * EPS
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// Request Object
type DisableEpRequest struct {
	EnterpriseProjectId string         `json:"enterprise_project_id"`
	Body                *DisableAction `json:"body,omitempty"`
}

func (o DisableEpRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DisableEpRequest", string(data)}, " ")
}
