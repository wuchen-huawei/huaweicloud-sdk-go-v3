/*
 * Cbr
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
type CopyBackupResponse struct {
	Replication *BackupReplicateRespBody `json:"replication,omitempty"`
}

func (o CopyBackupResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CopyBackupResponse", string(data)}, " ")
}
