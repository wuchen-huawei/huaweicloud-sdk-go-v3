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
type ShowVaultResponse struct {
	Vault *VaultGet `json:"vault,omitempty"`
}

func (o ShowVaultResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowVaultResponse", string(data)}, " ")
}
