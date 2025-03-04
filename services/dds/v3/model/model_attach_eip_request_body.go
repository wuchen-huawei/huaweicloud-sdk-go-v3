/*
 * DDS
 *
 * API v3
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

type AttachEipRequestBody struct {
	// 公网IP的ID。
	PublicIpId string `json:"public_ip_id"`
	// 公网IP。
	PublicIp string `json:"public_ip"`
}

func (o AttachEipRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"AttachEipRequestBody", string(data)}, " ")
}
