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

// 云服务器规格。
type Flavor struct {
	// 云服务器规格的ID。
	Id string `json:"id"`
	// 云服务器规格的名称。
	Name string `json:"name"`
	// 云服务器规格对应的CPU核数。
	Vcpus string `json:"vcpus"`
	// 云服务器规格对应的内存大小，单位为MB。
	Ram int32 `json:"ram"`
	// 云服务器规格对应要求的系统盘大小。  当前未使用该参数，缺省值为0。
	Disk string `json:"disk"`
	// 云服务器规格对应要求的交换分区大小。  当前未使用该参数，缺省值为\"\"。
	Swap string `json:"swap"`
	// 扩展属性，临时盘大小。  当前未使用该参数，缺省值为0
	OSFLVEXTDATAephemeral int32 `json:"OS-FLV-EXT-DATA:ephemeral"`
	// 扩展属性，该云服务器规格是否禁用。  当前未使用该参数，缺省值为false。
	OSFLVDISABLEDdisabled bool `json:"OS-FLV-DISABLED:disabled"`
	// 云服务器可使用网络带宽与网络硬件带宽的比例。  当前未使用该参数，缺省值为1.0。
	RxtxFactor float32 `json:"rxtx_factor"`
	// 云服务器可使用网络带宽的软限制。  当前未使用该参数，缺省值为null。
	RxtxQuota string `json:"rxtx_quota"`
	//   云服务器可使用网络带宽的硬限制。  当前未使用该参数，缺省值为null。
	RxtxCap string `json:"rxtx_cap"`
	// 扩展属性，flavor是否给所有租户使用。  - true：表示给所有租户使用。 - false：表示给指定租户使用。  缺省值为true。
	OsFlavorAccessisPublic bool `json:"os-flavor-access:is_public"`
	// 规格相关快捷链接地址。
	Links              []FlavorLink              `json:"links"`
	OsExtraSpecs       *FlavorExtraSpec          `json:"os_extra_specs"`
	AttachableQuantity *ServerAttachableQuantity `json:"attachableQuantity,omitempty"`
}

func (o Flavor) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"Flavor", string(data)}, " ")
}
