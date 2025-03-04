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

type CustomerInfoV2 struct {
	// |参数名称：客户的客户ID。| |参数约束及描述：客户的客户ID。|
	CustomerId string `json:"customer_id"`
}

func (o CustomerInfoV2) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CustomerInfoV2", string(data)}, " ")
}
