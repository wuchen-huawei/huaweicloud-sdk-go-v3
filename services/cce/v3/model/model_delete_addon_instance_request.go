/*
 * CCE
 *
 * CCE开放API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteAddonInstanceRequest struct {
	ContentType string `json:"Content-Type"`
	Id          string `json:"id"`
	ClusterId   string `json:"cluster_id"`
}

func (o DeleteAddonInstanceRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteAddonInstanceRequest", string(data)}, " ")
}
