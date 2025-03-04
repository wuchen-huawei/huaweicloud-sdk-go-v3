/*
 * BSS
 *
 * Business Support System API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListConversionsRequest struct {
	XLanguage   *string `json:"X-Language,omitempty"`
	MeasureType *int32  `json:"measure_type,omitempty"`
}

func (o ListConversionsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListConversionsRequest", string(data)}, " ")
}
