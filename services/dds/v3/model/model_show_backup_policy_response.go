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

// Response Object
type ShowBackupPolicyResponse struct {
	BackupPolicy   *BackupPolicyItem `json:"backup_policy,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowBackupPolicyResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowBackupPolicyResponse", string(data)}, " ")
}
