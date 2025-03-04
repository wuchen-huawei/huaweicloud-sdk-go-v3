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

// Request Object
type ListMessageTemplateDetailsRequest struct {
	MessageTemplateId string `json:"message_template_id"`
}

func (o ListMessageTemplateDetailsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListMessageTemplateDetailsRequest", string(data)}, " ")
}
