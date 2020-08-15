/*
 * TMS
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// 创建预定义标签请求
type ReqCreatePredefineTag struct {
	// 操作标识（区分大小写）：create（创建）
	Action ReqCreatePredefineTagAction `json:"action"`
	// 标签列表
	Tags []PredefineTagRequest `json:"tags"`
}

func (o ReqCreatePredefineTag) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ReqCreatePredefineTag", string(data)}, " ")
}

type ReqCreatePredefineTagAction struct {
	value string
}

type ReqCreatePredefineTagActionEnum struct {
	CREATE ReqCreatePredefineTagAction
}

func GetReqCreatePredefineTagActionEnum() ReqCreatePredefineTagActionEnum {
	return ReqCreatePredefineTagActionEnum{
		CREATE: ReqCreatePredefineTagAction{
			value: "create",
		},
	}
}

func (c ReqCreatePredefineTagAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ReqCreatePredefineTagAction) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}
