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
type DeleteServerMetadataResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteServerMetadataResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteServerMetadataResponse", string(data)}, " ")
}
