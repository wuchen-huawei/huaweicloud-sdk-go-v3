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

// 资源标签结构体。
type ResourceTag struct {
	// 键，表示要匹配的字段。  当前key的参数值只能取“resource_name”，此时value的参数值为云服务器名称。  - key不能重复，value为匹配的值。  - 此字段为固定字典值。  - 不允许为空字符串。
	Key string `json:"key"`
	// 值。  当前key的参数值只能取“resource_name”，此时value的参数值为云服务器名称。  - 每个值最大长度255个unicode字符。  - 不可以为空。
	Value string `json:"value"`
}

func (o ResourceTag) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ResourceTag", string(data)}, " ")
}
