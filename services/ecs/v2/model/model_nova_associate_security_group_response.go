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
type NovaAssociateSecurityGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o NovaAssociateSecurityGroupResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"NovaAssociateSecurityGroupResponse", string(data)}, " ")
}
