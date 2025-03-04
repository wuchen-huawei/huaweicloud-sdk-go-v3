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
type SetAuditlogPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetAuditlogPolicyResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"SetAuditlogPolicyResponse", string(data)}, " ")
}
