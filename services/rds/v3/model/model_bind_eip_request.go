/*
 * RDS
 *
 * API v3
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

type BindEipRequest struct {
	// 待绑定的弹性公网IP地址，仅允许使用标准的IP地址格式。is_bind为true时必选
	PublicIp *string `json:"public_ip,omitempty"`
	// 弹性公网IP对应的ID，仅允许使用标准的UUID格式。is_bind为true时必选
	PublicIpId *string `json:"public_ip_id,omitempty"`
	// - true，绑定弹性公网IP。 - false，解绑弹性公网IP。
	IsBind bool `json:"is_bind"`
}

func (o BindEipRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BindEipRequest", string(data)}, " ")
}
