/*
 * EVS
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
type RollbackSnapshotResponse struct {
	Rollback *RollbackInfo `json:"rollback,omitempty"`
}

func (o RollbackSnapshotResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"RollbackSnapshotResponse", string(data)}, " ")
}
