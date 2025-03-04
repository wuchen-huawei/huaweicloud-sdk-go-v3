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

// 执行任务的实例信息。
type GetTaskDetailListRspInstance struct {
	// 实例ID。
	Id string `json:"id"`
	// 实例名称。
	Name string `json:"name"`
}

func (o GetTaskDetailListRspInstance) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"GetTaskDetailListRspInstance", string(data)}, " ")
}
