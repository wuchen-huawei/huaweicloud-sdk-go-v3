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

// Response Object
type ShowAuditlogDownloadLinkResponse struct {
	// 审计日志下载链接列表。
	Links          *[]string `json:"links,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowAuditlogDownloadLinkResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowAuditlogDownloadLinkResponse", string(data)}, " ")
}
