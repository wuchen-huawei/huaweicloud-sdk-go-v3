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

// Response Object
type ResizePostPaidServerResponse struct {
	// 提交任务成功后返回的任务ID，用户可以使用该ID对任务执行情况进行查询。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResizePostPaidServerResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ResizePostPaidServerResponse", string(data)}, " ")
}
