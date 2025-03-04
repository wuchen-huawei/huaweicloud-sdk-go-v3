/*
 * DDS
 *
 * API v3
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListDatabaseRolesRequest struct {
	InstanceId string  `json:"instance_id"`
	RoleName   *string `json:"role_name,omitempty"`
	DbName     *string `json:"db_name,omitempty"`
	Offset     *int32  `json:"offset,omitempty"`
	Limit      *int32  `json:"limit,omitempty"`
}

func (o ListDatabaseRolesRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListDatabaseRolesRequest", string(data)}, " ")
}
