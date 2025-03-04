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
type ListErrorLogsResponse struct {
	ErrorLogList *[]ErrorLog `json:"error_log_list,omitempty"`
	// 总记录数。
	TotalRecord    *int32 `json:"total_record,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListErrorLogsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListErrorLogsResponse", string(data)}, " ")
}
