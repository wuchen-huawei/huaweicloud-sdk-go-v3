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

// Request Object
type CreateManualBackupRequest struct {
	Body *CreateManualBackupRequestBody `json:"body,omitempty"`
}

func (o CreateManualBackupRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateManualBackupRequest", string(data)}, " ")
}
