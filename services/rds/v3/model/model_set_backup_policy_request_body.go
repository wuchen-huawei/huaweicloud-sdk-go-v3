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

type SetBackupPolicyRequestBody struct {
	BackupPolicy *BackupPolicy `json:"backup_policy"`
	// 仅关闭备份策略时有效。  - true（默认），表示保留自动备份和差异备份。 - false，表示关闭备份策略的同时，删除已有的自动备份和差异备份。
	ReserveBackups *bool `json:"reserve_backups,omitempty"`
}

func (o SetBackupPolicyRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"SetBackupPolicyRequestBody", string(data)}, " ")
}
