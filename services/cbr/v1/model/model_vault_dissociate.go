/*
 * CBR
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

type VaultDissociate struct {
	// 策略ID
	PolicyId string `json:"policy_id"`
}

func (o VaultDissociate) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"VaultDissociate", string(data)}, " ")
}
