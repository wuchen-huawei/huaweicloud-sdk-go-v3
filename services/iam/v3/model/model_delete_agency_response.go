/*
 * IAM
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// Response Object
type DeleteAgencyResponse struct {
}

func (o DeleteAgencyResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteAgencyResponse", string(data)}, " ")
}
