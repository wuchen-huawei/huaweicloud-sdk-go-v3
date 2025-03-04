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
type CreateServerGroupResponse struct {
	ServerGroup    *CreateServerGroupResult `json:"server_group,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o CreateServerGroupResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateServerGroupResponse", string(data)}, " ")
}
