/*
 * AS
 *
 * 弹性伸缩API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteScalingInstanceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteScalingInstanceResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteScalingInstanceResponse", string(data)}, " ")
}
