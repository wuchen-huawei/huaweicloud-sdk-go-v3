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

type VaultBackupReq struct {
	Checkpoint *VaultBackup `json:"checkpoint"`
}

func (o VaultBackupReq) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"VaultBackupReq", string(data)}, " ")
}
