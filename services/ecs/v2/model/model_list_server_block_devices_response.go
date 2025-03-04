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
type ListServerBlockDevicesResponse struct {
	AttachableQuantity *BlockDeviceAttachableQuantity `json:"attachableQuantity,omitempty"`
	// 云服务器挂载信息列表。
	VolumeAttachments *[]ServerBlockDevice `json:"volumeAttachments,omitempty"`
	HttpStatusCode    int                  `json:"-"`
}

func (o ListServerBlockDevicesResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListServerBlockDevicesResponse", string(data)}, " ")
}
