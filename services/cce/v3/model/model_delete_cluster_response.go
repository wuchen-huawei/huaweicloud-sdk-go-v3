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
type DeleteClusterResponse struct {
	// API版本，固定值“v3”，该值不可修改。
	ApiVersion *string `json:"apiVersion,omitempty"`
	// API类型，固定值“Cluster”或“cluster”，该值不可修改。
	Kind           *string          `json:"kind,omitempty"`
	Metadata       *ClusterMetadata `json:"metadata,omitempty"`
	Spec           *V3ClusterSpec   `json:"spec,omitempty"`
	Status         *ClusterStatus   `json:"status,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o DeleteClusterResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteClusterResponse", string(data)}, " ")
}
