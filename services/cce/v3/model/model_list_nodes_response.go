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
type ListNodesResponse struct {
	// API版本，固定值“v3”
	ApiVersion *string `json:"apiVersion,omitempty"`
	// 节点对象列表，包含了当前集群下所有节点的详细信息。可通过items.metadata.name下的值来找到对应的节点。
	Items *[]V3Node `json:"items,omitempty"`
	// API类型，固定值“List”
	Kind           *string `json:"kind,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListNodesResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListNodesResponse", string(data)}, " ")
}
