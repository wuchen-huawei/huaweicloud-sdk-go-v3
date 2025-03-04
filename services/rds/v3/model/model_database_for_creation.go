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

// 数据库信息。
type DatabaseForCreation struct {
	// 数据库名称。 数据库名称长度可在1～64个字符之间，由字母、数字、中划线、下划线或$组成，$累计总长度小于等于10个字符，（MySQL 8.0不可包含$）。
	Name string `json:"name"`
	// 数据库使用的字符集。
	CharacterSet string `json:"character_set"`
}

func (o DatabaseForCreation) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DatabaseForCreation", string(data)}, " ")
}
