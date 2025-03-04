/*
 * SMN
 *
 * SMN Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListResourceTagsResponse struct {
	// 资源标签列表。
	Tags           *[]ResourceTag `json:"tags,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListResourceTagsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListResourceTagsResponse", string(data)}, " ")
}
