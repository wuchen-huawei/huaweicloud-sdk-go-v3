/*
 * BSS
 *
 * Business Support System API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateEnterpriseRealnameAuthenticationResponse struct {
	// |参数名称：是否需要转人工审核，只有状态码为200才返回该参数：0：不需要1：需要| |参数的约束及描述：是否需要转人工审核，只有状态码为200才返回该参数：0：不需要1：需要|
	IsReview       *int32 `json:"is_review,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateEnterpriseRealnameAuthenticationResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateEnterpriseRealnameAuthenticationResponse", string(data)}, " ")
}
