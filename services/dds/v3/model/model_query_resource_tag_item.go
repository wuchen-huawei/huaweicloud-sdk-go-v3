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

type QueryResourceTagItem struct {
	// 标签键。最大长度36个unicode字符，key不能为空。 字符集：0-9，A-Z，a-z，“_”，“-”，中文。
	Key *string `json:"key,omitempty"`
	// 标签值。最大长度43个unicode字符，可以为空字符串。 字符集：0-9，A-Z，a-z，“_”，“.”，“-”，中文。
	Value *string `json:"value,omitempty"`
}

func (o QueryResourceTagItem) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"QueryResourceTagItem", string(data)}, " ")
}
