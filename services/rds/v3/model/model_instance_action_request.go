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

type InstanceActionRequest struct {
	ResizeFlavor  *ResizeFlavorRequest          `json:"resize_flavor,omitempty"`
	EnlargeVolume *EnlargeVolume                `json:"enlarge_volume,omitempty"`
	Restart       *InstanceActionRequestRestart `json:"restart,omitempty"`
	SingleToHa    *Single2Ha                    `json:"single_to_ha,omitempty"`
}

func (o InstanceActionRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"InstanceActionRequest", string(data)}, " ")
}
