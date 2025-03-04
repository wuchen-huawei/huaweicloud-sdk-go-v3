/*
 * AS
 *
 * 弹性伸缩API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowResourceQuotaResponse struct {
	Quotas         *AllQuotas `json:"quotas,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ShowResourceQuotaResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowResourceQuotaResponse", string(data)}, " ")
}
