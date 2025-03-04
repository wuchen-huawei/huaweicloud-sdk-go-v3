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

//
type ResizePrePaidServerOption struct {
	// 变更后的云服务器规格ID。
	FlavorRef string `json:"flavorRef"`
	// 新专属主机ID（仅适用于专属主机上的弹性云服务器）。
	DedicatedHostId *string                  `json:"dedicated_host_id,omitempty"`
	Extendparam     *ResizeServerExtendParam `json:"extendparam,omitempty"`
	// 取值为withStopServer ，支持开机状态下变更规格。  mode取值为withStopServer时，对开机状态的云服务器执行变更规格操作，系统自动对云服务器先执行关机，再变更规格，变更成功后再执行开机。
	Mode *string `json:"mode,omitempty"`
}

func (o ResizePrePaidServerOption) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ResizePrePaidServerOption", string(data)}, " ")
}
