/*
 * CCE
 *
 * CCE开放API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

type Login struct {
	// 选择密钥对方式登录时的密钥对名称。密钥对和密码登录方式二者必选其一。
	SshKey       *string       `json:"sshKey,omitempty"`
	UserPassword *UserPassword `json:"userPassword,omitempty"`
}

func (o Login) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"Login", string(data)}, " ")
}
