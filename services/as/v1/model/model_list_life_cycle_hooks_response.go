/*
 * AS
 *
 * 弹性伸缩API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListLifeCycleHooksResponse struct {
	// 生命周期挂钩列表。
	LifecycleHooks *[]LifecycleHookList `json:"lifecycle_hooks,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListLifeCycleHooksResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListLifeCycleHooksResponse", string(data)}, " ")
}
