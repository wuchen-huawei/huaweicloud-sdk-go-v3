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

type CancelCustomerOrderReq struct {
	// |参数名称：订单ID。| |参数约束及描述：订单ID。|
	OrderId string `json:"order_id"`
}

func (o CancelCustomerOrderReq) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CancelCustomerOrderReq", string(data)}, " ")
}
