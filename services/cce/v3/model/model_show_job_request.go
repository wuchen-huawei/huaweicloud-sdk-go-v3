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
type ShowJobRequest struct {
	JobId       string `json:"job_id"`
	ContentType string `json:"Content-Type"`
}

func (o ShowJobRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowJobRequest", string(data)}, " ")
}
