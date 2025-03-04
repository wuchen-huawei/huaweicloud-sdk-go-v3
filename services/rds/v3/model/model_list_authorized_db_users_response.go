/*
 * RDS
 *
 * API v3
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListAuthorizedDbUsersResponse struct {
	// 用户及相关权限。
	Users *[]UserWithPrivilege `json:"users,omitempty"`
	// 总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAuthorizedDbUsersResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListAuthorizedDbUsersResponse", string(data)}, " ")
}
