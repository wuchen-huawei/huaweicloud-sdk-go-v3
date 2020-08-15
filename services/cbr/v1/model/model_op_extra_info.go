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

type OpExtraInfo struct {
	Backup          *OpExtendInfoBckup           `json:"backup,omitempty"`
	Common          *OpExtendInfoCommon          `json:"common"`
	Delete          *OpExtendInfoDelete          `json:"delete,omitempty"`
	Sync            *OpExtendInfoSync            `json:"sync,omitempty"`
	RemoveResources *OpExtendInfoRemoveResources `json:"remove_resources,omitempty"`
	Replication     *OpExtendInfoReplication     `json:"replication,omitempty"`
	Resource        *Resource                    `json:"resource"`
	Restore         *OpExtendInfoRestore         `json:"restore,omitempty"`
	VaultDelete     *OpExtendInfoVaultDelete     `json:"vault_delete,omitempty"`
}

func (o OpExtraInfo) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"OpExtraInfo", string(data)}, " ")
}
