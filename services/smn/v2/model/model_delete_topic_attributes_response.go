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
type DeleteTopicAttributesResponse struct {
	// 请求的唯一标识ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteTopicAttributesResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteTopicAttributesResponse", string(data)}, " ")
}
