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
type ListProvincesRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`
	Offset    *int32  `json:"offset,omitempty"`
	Limit     *int32  `json:"limit,omitempty"`
}

func (o ListProvincesRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListProvincesRequest", string(data)}, " ")
}
