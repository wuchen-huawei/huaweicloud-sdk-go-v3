/*
 * EVS
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// Request Object
type ListVolumeTagsRequest struct {
}

func (o ListVolumeTagsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListVolumeTagsRequest", string(data)}, " ")
}
