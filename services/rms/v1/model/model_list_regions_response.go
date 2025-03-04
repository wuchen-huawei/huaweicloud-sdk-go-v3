/*
 * RMS
 *
 * Resource Manager Api
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListRegionsResponse struct {
	// 区域信息项列表
	Value          *[]Region `json:"value,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListRegionsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListRegionsResponse", string(data)}, " ")
}
