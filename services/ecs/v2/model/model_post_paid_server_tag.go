/*
 * ECS
 *
 * ECS Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// 弹性云服务器的标签。
type PostPaidServerTag struct {
	// 键。  最大长度36个unicode字符。key不能为空。不能包含非打印字符ASCII(0-31)，\"=\", \"*\",“<”,“>”,“\\”,“,”,“|”,“/”。  同一资源的key值不能重复。
	Key string `json:"key"`
	//   值。  每个值最大长度43个unicode字符，可以为空字符串。 不能包含非打印字符ASCII(0-31)，“=”,“*”,“<”,“>”,“\\”,“,”,“|”,“/”。
	Value string `json:"value"`
}

func (o PostPaidServerTag) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"PostPaidServerTag", string(data)}, " ")
}
