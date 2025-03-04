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

// 节点网卡的描述信息。
type NodeNicSpec struct {
	// 扩展网卡
	ExtNics    *[]NicSpec `json:"extNics,omitempty"`
	PrimaryNic *NicSpec   `json:"primaryNic,omitempty"`
}

func (o NodeNicSpec) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"NodeNicSpec", string(data)}, " ")
}
