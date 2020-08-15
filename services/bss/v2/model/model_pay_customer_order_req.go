/*
 * Bss
 *
 * Business Support System API
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

type PayCustomerOrderReq struct {
	// |参数名称：字段预留。优惠券列表，目前仅支持传递一个优惠券ID。请从“1.3-查询订单可用优惠券”接口的响应参数中获取。| |参数约束以及描述：字段预留。优惠券列表，目前仅支持传递一个优惠券ID。请从“1.3-查询订单可用优惠券”接口的响应参数中获取。|
	CouponInfos []CouponSimpleInfoOrderPay `json:"coupon_infos,omitempty"`
	// |参数名称：订单ID。| |参数约束及描述：订单ID。|
	OrderId string `json:"order_id"`
}

func (o PayCustomerOrderReq) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"PayCustomerOrderReq", string(data)}, " ")
}
