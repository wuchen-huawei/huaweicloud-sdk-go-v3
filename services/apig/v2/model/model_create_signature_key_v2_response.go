/*
 * APIG
 *
 * API网关（API Gateway）是为开发者、合作伙伴提供的高性能、高可用、高安全的API托管服务，帮助用户轻松构建、管理和发布任意规模的API。
 *
 */

package model

import (
	"encoding/json"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"strings"
)

// Response Object
type CreateSignatureKeyV2Response struct {
	// 签名密钥的密钥
	SignSecret string `json:"sign_secret,omitempty"`
	// 更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`
	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`
	// 签名密钥的名称
	Name string `json:"name,omitempty"`
	// 签名密钥的编号
	Id string `json:"id,omitempty"`
	// 签名密钥的key
	SignKey string `json:"sign_key,omitempty"`
	// 签名密钥类型。
	SignType CreateSignatureKeyV2ResponseSignType `json:"sign_type,omitempty"`
}

func (o CreateSignatureKeyV2Response) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateSignatureKeyV2Response", string(data)}, " ")
}

type CreateSignatureKeyV2ResponseSignType struct {
	value string
}

type CreateSignatureKeyV2ResponseSignTypeEnum struct {
	HMAC  CreateSignatureKeyV2ResponseSignType
	BASIC CreateSignatureKeyV2ResponseSignType
}

func GetCreateSignatureKeyV2ResponseSignTypeEnum() CreateSignatureKeyV2ResponseSignTypeEnum {
	return CreateSignatureKeyV2ResponseSignTypeEnum{
		HMAC: CreateSignatureKeyV2ResponseSignType{
			value: "hmac",
		},
		BASIC: CreateSignatureKeyV2ResponseSignType{
			value: "basic",
		},
	}
}

func (c CreateSignatureKeyV2ResponseSignType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *CreateSignatureKeyV2ResponseSignType) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}
