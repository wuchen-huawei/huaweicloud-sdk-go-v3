/*
 * SMN
 *
 * SMN Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

type ListApplicationEndpointAttributesResponseBodyAttributes struct {
	// 设备是否可用。
	Enabled string `json:"enabled"`
	// 设备token。
	Token string `json:"token"`
	// 用户数据。
	UserData string `json:"user_data"`
}

func (o ListApplicationEndpointAttributesResponseBodyAttributes) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListApplicationEndpointAttributesResponseBodyAttributes", string(data)}, " ")
}
