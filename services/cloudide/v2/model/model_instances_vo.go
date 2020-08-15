/*
 * cloudide
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

type InstancesVo struct {
	// 角色权限列表
	ActionList []RoleAction `json:"action_list,omitempty"`
	// cpu架构 x86|arm
	Arch       InstancesVoArch `json:"arch,omitempty"`
	Attributes *Attributes     `json:"attributes,omitempty"`
	// cpu规格.arm架构支持4U8G，x86架构支持1U1G,2U4G,2U8G 与技术栈配置的规格对应，可通过技术栈管理ListStacksByTag接口获取。如果标签不为空，以标签配置的技术栈规格为准。 quantum技术栈，x86架构cpu规格为2U8G;其他技术栈，x86架构cpu规格为1U1G,2U4G
	CpuMemory InstancesVoCpuMemory `json:"cpu_memory,omitempty"`
	// 创建时间
	CreatedTime string `json:"created_time,omitempty"`
	// 描述
	Description string `json:"description,omitempty"`
	// 实例名。 可以输入中文、数字、字母、下划线、点、破折号。长度介于3-100之间
	DisplayName string `json:"display_name,omitempty"`
	// 组织名
	DomainName string `json:"domain_name,omitempty"`
	// id
	Id string `json:"id,omitempty"`
	// 是否临时实例。 false页面会显示
	IsTemporary bool `json:"is_temporary,omitempty"`
	// 标签
	Label string `json:"label,omitempty"`
	// 链接
	Link string `json:"link,omitempty"`
	// 名称
	Name string `json:"name,omitempty"`
	// 组织id（对应华为云账号的domainId）
	OrganizationId string `json:"organization_id,omitempty"`
	// 用户id
	OwnerId string `json:"owner_id,omitempty"`
	// 用户名
	OwnerName string `json:"owner_name,omitempty"`
	// 平台ID
	PlatformId int64 `json:"platform_id,omitempty"`
	// 是否私有平台
	Private bool `json:"private,omitempty"`
	// PVC规格 5GB|10GB|20GB
	PvcQuantity InstancesVoPvcQuantity `json:"pvc_quantity,omitempty"`
	// 实例的生命周期 arm架构,生命周期只能设置成30，60。x86架构可取值为30，60，240，1440和-1。除-1外，其它值的单位为“分钟”。实例在到达生命周期后，将会被暂停（已保存的数据不会被删除）。-1表示实例不会自动停止。
	RefreshInterval int64 `json:"refresh_interval,omitempty"`
	// 区域
	Region string `json:"region,omitempty"`
	Role   *Role  `json:"role,omitempty"`
	// 角色id
	RoleId string `json:"role_id,omitempty"`
	// server
	ServerMap map[string]string `json:"server_map,omitempty"`
	// 服务链接
	ServerUrl string `json:"server_url,omitempty"`
	// 技术栈ID 目前可取值all，java，go，python，cpp，nodejs，quantum，blockchain，dcn，vue，ruby。
	StackId string `json:"stack_id,omitempty"`
	// 实例状态
	Status InstancesVoStatus `json:"status,omitempty"`
	// 子组织
	SubOrg string `json:"sub_org,omitempty"`
	// 更新时间
	UpdatedTime string `json:"updated_time,omitempty"`
}

func (o InstancesVo) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"InstancesVo", string(data)}, " ")
}

type InstancesVoArch struct {
	value string
}

type InstancesVoArchEnum struct {
	X86 InstancesVoArch
	ARM InstancesVoArch
}

func GetInstancesVoArchEnum() InstancesVoArchEnum {
	return InstancesVoArchEnum{
		X86: InstancesVoArch{
			value: "x86",
		},
		ARM: InstancesVoArch{
			value: "arm",
		},
	}
}

func (c InstancesVoArch) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *InstancesVoArch) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}

type InstancesVoCpuMemory struct {
	value string
}

type InstancesVoCpuMemoryEnum struct {
	_1_U1_G InstancesVoCpuMemory
	_2_U4_G InstancesVoCpuMemory
	_2_U8_G InstancesVoCpuMemory
	_4_U8_G InstancesVoCpuMemory
}

func GetInstancesVoCpuMemoryEnum() InstancesVoCpuMemoryEnum {
	return InstancesVoCpuMemoryEnum{
		_1_U1_G: InstancesVoCpuMemory{
			value: "1U1G",
		},
		_2_U4_G: InstancesVoCpuMemory{
			value: "2U4G",
		},
		_2_U8_G: InstancesVoCpuMemory{
			value: "2U8G",
		},
		_4_U8_G: InstancesVoCpuMemory{
			value: "4U8G",
		},
	}
}

func (c InstancesVoCpuMemory) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *InstancesVoCpuMemory) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}

type InstancesVoPvcQuantity struct {
	value string
}

type InstancesVoPvcQuantityEnum struct {
	_5_GB  InstancesVoPvcQuantity
	_10_GB InstancesVoPvcQuantity
	_20_GB InstancesVoPvcQuantity
}

func GetInstancesVoPvcQuantityEnum() InstancesVoPvcQuantityEnum {
	return InstancesVoPvcQuantityEnum{
		_5_GB: InstancesVoPvcQuantity{
			value: "5GB",
		},
		_10_GB: InstancesVoPvcQuantity{
			value: "10GB",
		},
		_20_GB: InstancesVoPvcQuantity{
			value: "20GB",
		},
	}
}

func (c InstancesVoPvcQuantity) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *InstancesVoPvcQuantity) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}

type InstancesVoStatus struct {
	value string
}

type InstancesVoStatusEnum struct {
	INIT          InstancesVoStatus
	STARTING      InstancesVoStatus
	RUNNING       InstancesVoStatus
	STOPPING      InstancesVoStatus
	STOPPED       InstancesVoStatus
	DELETING      InstancesVoStatus
	DELETED       InstancesVoStatus
	DELETE_FAILED InstancesVoStatus
}

func GetInstancesVoStatusEnum() InstancesVoStatusEnum {
	return InstancesVoStatusEnum{
		INIT: InstancesVoStatus{
			value: "INIT",
		},
		STARTING: InstancesVoStatus{
			value: "STARTING",
		},
		RUNNING: InstancesVoStatus{
			value: "RUNNING",
		},
		STOPPING: InstancesVoStatus{
			value: "STOPPING",
		},
		STOPPED: InstancesVoStatus{
			value: "STOPPED",
		},
		DELETING: InstancesVoStatus{
			value: "DELETING",
		},
		DELETED: InstancesVoStatus{
			value: "DELETED",
		},
		DELETE_FAILED: InstancesVoStatus{
			value: "DELETE_FAILED",
		},
	}
}

func (c InstancesVoStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *InstancesVoStatus) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}
