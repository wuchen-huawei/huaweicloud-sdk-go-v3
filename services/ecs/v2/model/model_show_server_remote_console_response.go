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
type ShowServerRemoteConsoleResponse struct {
	RemoteConsole  *ServerRemoteConsole `json:"remote_console,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowServerRemoteConsoleResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowServerRemoteConsoleResponse", string(data)}, " ")
}
