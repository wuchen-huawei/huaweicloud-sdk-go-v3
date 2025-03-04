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

type UpdateSecurityGroupRequestBody struct {
	// 新的安全组ID。
	SecurityGroupId string `json:"security_group_id"`
}

func (o UpdateSecurityGroupRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateSecurityGroupRequestBody", string(data)}, " ")
}
