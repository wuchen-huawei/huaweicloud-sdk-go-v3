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
type DeleteServerGroupMemberResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteServerGroupMemberResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteServerGroupMemberResponse", string(data)}, " ")
}
