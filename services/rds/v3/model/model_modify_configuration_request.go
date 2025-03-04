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
type ModifyConfigurationRequest struct {
	XLanguage *string                 `json:"X-Language,omitempty"`
	ConfigId  string                  `json:"config_id"`
	Body      *ConfigurationForUpdate `json:"body,omitempty"`
}

func (o ModifyConfigurationRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ModifyConfigurationRequest", string(data)}, " ")
}
