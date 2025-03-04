/*
 * RDS
 *
 * API v3
 *
 */

package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

type ConfigurationParameter struct {
	// 参数名称。
	Name string `json:"name"`
	// 参数值。
	Value string `json:"value"`
	// 修改该参数是否需要重启实例。
	RestartRequired bool `json:"restart_required"`
	// 该参数是否只读。
	Readonly bool `json:"readonly"`
	// 参数取值范围。
	ValueRange string `json:"value_range"`
	// 参数类型，取值为“string”、“integer”、“boolean”、“list”或“float”之一。
	Type ConfigurationParameterType `json:"type"`
	// 参数描述。
	Description string `json:"description"`
}

func (o ConfigurationParameter) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ConfigurationParameter", string(data)}, " ")
}

type ConfigurationParameterType struct {
	value string
}

type ConfigurationParameterTypeEnum struct {
	STRING  ConfigurationParameterType
	INTEGER ConfigurationParameterType
	BOOLEAN ConfigurationParameterType
	LIST    ConfigurationParameterType
	FLOAT   ConfigurationParameterType
}

func GetConfigurationParameterTypeEnum() ConfigurationParameterTypeEnum {
	return ConfigurationParameterTypeEnum{
		STRING: ConfigurationParameterType{
			value: "string",
		},
		INTEGER: ConfigurationParameterType{
			value: "integer",
		},
		BOOLEAN: ConfigurationParameterType{
			value: "boolean",
		},
		LIST: ConfigurationParameterType{
			value: "list",
		},
		FLOAT: ConfigurationParameterType{
			value: "float",
		},
	}
}

func (c ConfigurationParameterType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ConfigurationParameterType) UnmarshalJSON(b []byte) error {
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
