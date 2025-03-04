/*
 * BSS
 *
 * Business Support System API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListEnterpriseMultiAccountRequest struct {
	SubCustomerId string `json:"sub_customer_id"`
	BalanceType   string `json:"balance_type"`
	Offset        *int32 `json:"offset,omitempty"`
	Limit         *int32 `json:"limit,omitempty"`
}

func (o ListEnterpriseMultiAccountRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListEnterpriseMultiAccountRequest", string(data)}, " ")
}
