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
type ShowScalingPolicyResponse struct {
	ScalingPolicy  *ScalingPolicyDetail `json:"scaling_policy,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowScalingPolicyResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowScalingPolicyResponse", string(data)}, " ")
}
