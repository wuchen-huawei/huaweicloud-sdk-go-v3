/*
 * EPS
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// 停用启用企业项目操作
type EnableAction struct {
	// 启用操作
	Action EnableActionAction `json:"action"`
}

func (o EnableAction) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"EnableAction", string(data)}, " ")
}

type EnableActionAction struct {
	value string
}

type EnableActionActionEnum struct {
	ENABLE EnableActionAction
}

func GetEnableActionActionEnum() EnableActionActionEnum {
	return EnableActionActionEnum{
		ENABLE: EnableActionAction{
			value: "enable",
		},
	}
}

func (c EnableActionAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *EnableActionAction) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}
