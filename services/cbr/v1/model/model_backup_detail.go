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
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"strings"
)

type BackupDetail struct {
	// 还原点ID
	CheckpointId string `json:"checkpoint_id"`
	// 创建时间，例如:\"2020-02-05T10:38:34.209782\"
	CreatedAt *sdktime.SdkTime `json:"created_at"`
	// 备份描述
	Description string `json:"description"`
	// 过期时间，例如:\"2020-02-05T10:38:34.209782\"
	ExpiredAt  *sdktime.SdkTime  `json:"expired_at"`
	ExtendInfo *BackupExtendInfo `json:"extend_info"`
	// 备份ID
	Id string `json:"id"`
	// 备份类型
	ImageType BackupDetailImageType `json:"image_type"`
	// 备份名称
	Name string `json:"name"`
	// 父备份ID
	ParentId string `json:"parent_id"`
	// 项目ID
	ProjectId string `json:"project_id"`
	// 备份时间
	ProtectedAt string `json:"protected_at"`
	// 资源可用区
	ResourceAz string `json:"resource_az"`
	// 资源ID
	ResourceId string `json:"resource_id"`
	// 资源名称
	ResourceName string `json:"resource_name"`
	// 资源大小，单位为GB
	ResourceSize int32 `json:"resource_size"`
	// 资源类型
	ResourceType BackupDetailResourceType `json:"resource_type"`
	// 备份状态
	Status BackupDetailStatus `json:"status"`
	// 更新时间，例如:\"2020-02-05T10:38:34.209782\"
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`
	// 存储库ID
	VaultId string `json:"vault_id"`
	// 复制记录
	ReplicationRecords []ReplicationRecordGet `json:"replication_records"`
	// 企业项目id,默认为‘0’。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	//
	Children []BackupResp `json:"children"`
}

func (o BackupDetail) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BackupDetail", string(data)}, " ")
}

type BackupDetailImageType struct {
	value string
}

type BackupDetailImageTypeEnum struct {
	BACKUP      BackupDetailImageType
	REPLICATION BackupDetailImageType
}

func GetBackupDetailImageTypeEnum() BackupDetailImageTypeEnum {
	return BackupDetailImageTypeEnum{
		BACKUP: BackupDetailImageType{
			value: "backup",
		},
		REPLICATION: BackupDetailImageType{
			value: "replication",
		},
	}
}

func (c BackupDetailImageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *BackupDetailImageType) UnmarshalJSON(b []byte) error {
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

type BackupDetailResourceType struct {
	value string
}

type BackupDetailResourceTypeEnum struct {
	OSNOVASERVER   BackupDetailResourceType
	OSCINDERVOLUME BackupDetailResourceType
}

func GetBackupDetailResourceTypeEnum() BackupDetailResourceTypeEnum {
	return BackupDetailResourceTypeEnum{
		OSNOVASERVER: BackupDetailResourceType{
			value: "OS::Nova::Server",
		},
		OSCINDERVOLUME: BackupDetailResourceType{
			value: "OS::Cinder::Volume",
		},
	}
}

func (c BackupDetailResourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *BackupDetailResourceType) UnmarshalJSON(b []byte) error {
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

type BackupDetailStatus struct {
	value string
}

type BackupDetailStatusEnum struct {
	AVAILABLE       BackupDetailStatus
	PROTECTING      BackupDetailStatus
	DELETING        BackupDetailStatus
	RESTORING       BackupDetailStatus
	ERROR           BackupDetailStatus
	WAITING_PROTECT BackupDetailStatus
	WAITING_DELETE  BackupDetailStatus
	WAITING_RESTORE BackupDetailStatus
}

func GetBackupDetailStatusEnum() BackupDetailStatusEnum {
	return BackupDetailStatusEnum{
		AVAILABLE: BackupDetailStatus{
			value: "available",
		},
		PROTECTING: BackupDetailStatus{
			value: "protecting",
		},
		DELETING: BackupDetailStatus{
			value: "deleting",
		},
		RESTORING: BackupDetailStatus{
			value: "restoring",
		},
		ERROR: BackupDetailStatus{
			value: "error",
		},
		WAITING_PROTECT: BackupDetailStatus{
			value: "waiting_protect",
		},
		WAITING_DELETE: BackupDetailStatus{
			value: "waiting_delete",
		},
		WAITING_RESTORE: BackupDetailStatus{
			value: "waiting_restore",
		},
	}
}

func (c BackupDetailStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *BackupDetailStatus) UnmarshalJSON(b []byte) error {
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
