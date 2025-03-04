/*
 * CCE
 *
 * CCE开放API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteCloudPersistentVolumeClaimsResponse struct {
	// API版本，固定值**v1**
	ApiVersion *string `json:"apiVersion,omitempty"`
	// API类型，固定值**PersistentVolumeClaim**
	Kind           *string                        `json:"kind,omitempty"`
	Metadata       *PersistentVolumeClaimMetadata `json:"metadata,omitempty"`
	Spec           *PersistentVolumeClaimSpec     `json:"spec,omitempty"`
	Status         *PersistentVolumeClaimStatus   `json:"status,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o DeleteCloudPersistentVolumeClaimsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteCloudPersistentVolumeClaimsResponse", string(data)}, " ")
}
