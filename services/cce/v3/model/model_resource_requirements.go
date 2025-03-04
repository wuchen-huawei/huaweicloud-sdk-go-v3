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

type ResourceRequirements struct {
	// 资源限制，创建时指定无效
	Limits map[string]string `json:"limits,omitempty"`
	// 资源需求，创建时指定无效
	Requests map[string]string `json:"requests,omitempty"`
}

func (o ResourceRequirements) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ResourceRequirements", string(data)}, " ")
}
