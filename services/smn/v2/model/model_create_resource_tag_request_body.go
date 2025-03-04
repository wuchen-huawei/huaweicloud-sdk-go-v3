/*
 * SMN
 *
 * SMN Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

type CreateResourceTagRequestBody struct {
	Tag *CreateResourceTagRequestBodyTag `json:"tag"`
}

func (o CreateResourceTagRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateResourceTagRequestBody", string(data)}, " ")
}
