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

// Request Object
type DeleteManualBackupRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`
	BackupId  string  `json:"backup_id"`
}

func (o DeleteManualBackupRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteManualBackupRequest", string(data)}, " ")
}
