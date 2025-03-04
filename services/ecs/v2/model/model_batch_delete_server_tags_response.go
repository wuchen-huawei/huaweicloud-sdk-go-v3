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
type BatchDeleteServerTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteServerTagsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchDeleteServerTagsResponse", string(data)}, " ")
}
