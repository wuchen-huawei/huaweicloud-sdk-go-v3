/*
 * CBR
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowVaultResourceInstancesResponse struct {
	// 符合查询条件的资源列表（action为count时无此参数）。
	Resources *[]TagResource `json:"resources,omitempty"`
	// 符合查询条件的资源总个数
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowVaultResourceInstancesResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowVaultResourceInstancesResponse", string(data)}, " ")
}
