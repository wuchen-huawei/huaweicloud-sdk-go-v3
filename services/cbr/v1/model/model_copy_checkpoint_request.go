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

// Request Object
type CopyCheckpointRequest struct {
	Body *CheckpointReplicateReq `json:"body,omitempty"`
}

func (o CopyCheckpointRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CopyCheckpointRequest", string(data)}, " ")
}
