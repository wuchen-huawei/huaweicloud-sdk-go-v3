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

type PwdResetRequest struct {
	// 数据库密码
	DbUserPwd string `json:"db_user_pwd"`
}

func (o PwdResetRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"PwdResetRequest", string(data)}, " ")
}
