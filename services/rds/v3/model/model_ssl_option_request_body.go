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

type SslOptionRequestBody struct {
	// - true, 打开ssl开关。 - false, 关闭ssl开关。
	SslOption bool `json:"ssl_option"`
}

func (o SslOptionRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"SslOptionRequestBody", string(data)}, " ")
}
