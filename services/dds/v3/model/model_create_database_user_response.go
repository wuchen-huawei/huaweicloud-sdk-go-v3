/*
 * DDS
 *
 * API v3
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateDatabaseUserResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateDatabaseUserResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateDatabaseUserResponse", string(data)}, " ")
}
