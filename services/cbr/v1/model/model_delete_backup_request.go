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

// Request Object
type DeleteBackupRequest struct {
	BackupId string `json:"backup_id"`
}

func (o DeleteBackupRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteBackupRequest", string(data)}, " ")
}
