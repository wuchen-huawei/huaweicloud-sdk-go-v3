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

type DeleteStatus struct {
	// 集群删除时更新的资源记录总数
	Added *int32 `json:"added,omitempty"`
	// 基于当前集群资源记录信息，生成实际最新资源记录总数
	CurrentTotal *int32 `json:"current_total,omitempty"`
	// 集群删除时删除的资源记录总数
	Deleted *int32 `json:"deleted,omitempty"`
	// 集群删除时已经存在的集群资源记录总数
	PreviousTotal *int32 `json:"previous_total,omitempty"`
	// 集群删除时更新的资源记录总数
	Updated *int32 `json:"updated,omitempty"`
}

func (o DeleteStatus) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteStatus", string(data)}, " ")
}
