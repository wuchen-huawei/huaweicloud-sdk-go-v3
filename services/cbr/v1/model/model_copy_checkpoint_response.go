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

// Response Object
type CopyCheckpointResponse struct {
	Replication    *CheckpointReplicateRespBody `json:"replication,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o CopyCheckpointResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CopyCheckpointResponse", string(data)}, " ")
}
