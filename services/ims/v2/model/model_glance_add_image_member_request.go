/*
 * IMS
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// Request Object
type GlanceAddImageMemberRequest struct {
	ImageId string                           `json:"image_id"`
	Body    *GlanceAddImageMemberRequestBody `json:"body,omitempty"`
}

func (o GlanceAddImageMemberRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"GlanceAddImageMemberRequest", string(data)}, " ")
}
