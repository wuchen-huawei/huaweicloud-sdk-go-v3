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

// 数据库用户信息。
type UserForList struct {
	// 数据库用户名称。
	Name string `json:"name"`
}

func (o UserForList) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UserForList", string(data)}, " ")
}
