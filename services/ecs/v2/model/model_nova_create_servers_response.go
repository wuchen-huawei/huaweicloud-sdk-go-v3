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
type NovaCreateServersResponse struct {
	Server         *NovaCreateServersResult `json:"server,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o NovaCreateServersResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"NovaCreateServersResponse", string(data)}, " ")
}
