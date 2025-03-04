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
type CreateConfigurationResponse struct {
	Configuration  *ConfigurationSummaryForCreate `json:"configuration,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o CreateConfigurationResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateConfigurationResponse", string(data)}, " ")
}
