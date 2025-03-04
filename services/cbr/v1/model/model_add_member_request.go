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

// Request Object
type AddMemberRequest struct {
	BackupId string         `json:"backup_id"`
	Body     *AddMembersReq `json:"body,omitempty"`
}

func (o AddMemberRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"AddMemberRequest", string(data)}, " ")
}
