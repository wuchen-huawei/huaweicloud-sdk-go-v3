/*
 * CBR
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

// Request Object
type ShowProtectableRequest struct {
	InstanceId      string                                `json:"instance_id"`
	ProtectableType ShowProtectableRequestProtectableType `json:"protectable_type"`
}

func (o ShowProtectableRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowProtectableRequest", string(data)}, " ")
}

type ShowProtectableRequestProtectableType struct {
	value string
}

type ShowProtectableRequestProtectableTypeEnum struct {
	SERVER ShowProtectableRequestProtectableType
	DISK   ShowProtectableRequestProtectableType
}

func GetShowProtectableRequestProtectableTypeEnum() ShowProtectableRequestProtectableTypeEnum {
	return ShowProtectableRequestProtectableTypeEnum{
		SERVER: ShowProtectableRequestProtectableType{
			value: "server",
		},
		DISK: ShowProtectableRequestProtectableType{
			value: "disk",
		},
	}
}

func (c ShowProtectableRequestProtectableType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ShowProtectableRequestProtectableType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
