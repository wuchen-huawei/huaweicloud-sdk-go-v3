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

// Request Object
type ShowAuditlogDownloadLinkRequest struct {
	XLanguage  *string                              `json:"X-Language,omitempty"`
	InstanceId string                               `json:"instance_id"`
	Body       *GenerateAuditlogDownloadLinkRequest `json:"body,omitempty"`
}

func (o ShowAuditlogDownloadLinkRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowAuditlogDownloadLinkRequest", string(data)}, " ")
}
