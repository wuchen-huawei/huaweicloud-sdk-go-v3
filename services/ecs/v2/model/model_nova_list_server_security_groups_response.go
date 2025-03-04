/*
 * ECS
 *
 * ECS Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type NovaListServerSecurityGroupsResponse struct {
	// security_group列表
	SecurityGroups *[]NovaSecurityGroup `json:"security_groups,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o NovaListServerSecurityGroupsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"NovaListServerSecurityGroupsResponse", string(data)}, " ")
}
