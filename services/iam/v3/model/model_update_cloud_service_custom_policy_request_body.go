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

//
type UpdateCloudServiceCustomPolicyRequestBody struct {
	Role *ServicePolicyRoleOption `json:"role"`
}

func (o UpdateCloudServiceCustomPolicyRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateCloudServiceCustomPolicyRequestBody", string(data)}, " ")
}
