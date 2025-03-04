/*
 * CCE
 *
 * CCE开放API
 *
 */

package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

type NodePoolSpec struct {
	Autoscaling *NodePoolNodeAutoscaling `json:"autoscaling,omitempty"`
	// 节点池初始化节点个数。
	InitialNodeCount *int32          `json:"initialNodeCount,omitempty"`
	NodeManagement   *NodeManagement `json:"nodeManagement,omitempty"`
	NodeTemplate     *V3NodeSpec     `json:"nodeTemplate"`
	// 节点池类型。不填写时默认为vm。  - vm：弹性云服务器 - ElasticBMS：C6型弹性裸金属通用计算增强型云服务器，规格示例：c6.22xlarge.2.physical
	Type *NodePoolSpecType `json:"type,omitempty"`
}

func (o NodePoolSpec) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"NodePoolSpec", string(data)}, " ")
}

type NodePoolSpecType struct {
	value string
}

type NodePoolSpecTypeEnum struct {
	VM          NodePoolSpecType
	ELASTIC_BMS NodePoolSpecType
}

func GetNodePoolSpecTypeEnum() NodePoolSpecTypeEnum {
	return NodePoolSpecTypeEnum{
		VM: NodePoolSpecType{
			value: "vm",
		},
		ELASTIC_BMS: NodePoolSpecType{
			value: "ElasticBMS",
		},
	}
}

func (c NodePoolSpecType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *NodePoolSpecType) UnmarshalJSON(b []byte) error {
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
