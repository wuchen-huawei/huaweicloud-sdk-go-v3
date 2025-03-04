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

type CheckpointCreate struct {
	// 创建时间，例如:\"2020-02-05T10:38:34.209782\"
	CreatedAt string `json:"created_at"`
	// 还原点ID
	Id string `json:"id"`
	// 项目ID
	ProjectId string `json:"project_id"`
	// 状态
	Status    CheckpointCreateStatus   `json:"status"`
	Vault     *CheckpointPlanCreate    `json:"vault,omitempty"`
	ExtraInfo *CheckpointExtraInfoResp `json:"extra_info,omitempty"`
}

func (o CheckpointCreate) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CheckpointCreate", string(data)}, " ")
}

type CheckpointCreateStatus struct {
	value string
}

type CheckpointCreateStatusEnum struct {
	AVAILABLE CheckpointCreateStatus
	DELETING  CheckpointCreateStatus
	CREATING  CheckpointCreateStatus
	RESTORING CheckpointCreateStatus
	ERROR     CheckpointCreateStatus
}

func GetCheckpointCreateStatusEnum() CheckpointCreateStatusEnum {
	return CheckpointCreateStatusEnum{
		AVAILABLE: CheckpointCreateStatus{
			value: "available",
		},
		DELETING: CheckpointCreateStatus{
			value: "deleting",
		},
		CREATING: CheckpointCreateStatus{
			value: "creating",
		},
		RESTORING: CheckpointCreateStatus{
			value: "restoring",
		},
		ERROR: CheckpointCreateStatus{
			value: "error",
		},
	}
}

func (c CheckpointCreateStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *CheckpointCreateStatus) UnmarshalJSON(b []byte) error {
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
