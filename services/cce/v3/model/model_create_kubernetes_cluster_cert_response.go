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

// Response Object
type CreateKubernetesClusterCertResponse struct {
	// API版本，固定值“v1”。
	ApiVersion *string `json:"apiVersion,omitempty"`
	// 集群列表。
	Clusters *[]Clusters `json:"clusters,omitempty"`
	// 上下文列表。
	Contexts *[]Contexts `json:"contexts,omitempty"`
	// 当前上下文，若存在publicIp（虚拟机弹性IP）时为 external; 若不存在publicIp为 internal。
	CurrentContext *string `json:"current-context,omitempty"`
	// API类型，固定值“Config”，该值不可修改。
	Kind *string `json:"kind,omitempty"`
	// 当前未使用该字段，当前默认为空。
	Preferences *interface{} `json:"preferences,omitempty"`
	// 存放了指定用户的一些证书信息和ClientKey信息。
	Users          *[]Users `json:"users,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o CreateKubernetesClusterCertResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateKubernetesClusterCertResponse", string(data)}, " ")
}
