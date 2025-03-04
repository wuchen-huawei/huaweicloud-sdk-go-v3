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

// 主网卡的描述信息。
type NicSpec struct {
	// 主网卡的IP将通过fixedIps指定，数量不得大于创建的节点数。fixedIps或ipBlock同时只能指定一个。
	FixedIps *[]string `json:"fixedIps,omitempty"`
	// IP段的CIDR格式，创建的节点IP将属于该IP段内。fixedIps或ipBlock同时只能指定一个。
	IpBlock *string `json:"ipBlock,omitempty"`
	// 网卡所在子网的ID。
	SubnetId *string `json:"subnetId,omitempty"`
}

func (o NicSpec) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"NicSpec", string(data)}, " ")
}
