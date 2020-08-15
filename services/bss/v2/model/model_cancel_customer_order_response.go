/*
 * Bss
 *
 * Business Support System API
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// Response Object
type CancelCustomerOrderResponse struct {
}

func (o CancelCustomerOrderResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CancelCustomerOrderResponse", string(data)}, " ")
}
