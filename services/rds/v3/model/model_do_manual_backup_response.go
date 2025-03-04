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
type DoManualBackupResponse struct {
	Backup         *BackupInfo `json:"backup,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o DoManualBackupResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DoManualBackupResponse", string(data)}, " ")
}
