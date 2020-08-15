/*
 * IMS
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// Response Object
type UpdateImageResponse struct {
	// 备份ID。如果是备份创建的镜像，则填写为备份的ID，否则为空
	BackupId string `json:"__backup_id,omitempty"`
	// 镜像来源。公共镜像为空
	DataOrigin string `json:"__data_origin,omitempty"`
	// 镜像描述信息。 支持字母、数字、中文等，不支持回车、<、 >，长度不能超过1024个字符。
	Description string `json:"__description,omitempty"`
	// 镜像的存储位置
	ImageLocation string `json:"__image_location,omitempty"`
	// 镜像文件的大小，单位为字节
	ImageSize string `json:"__image_size,omitempty"`
	// 镜像后端存储类型，目前只支持uds
	ImageSourceType UpdateImageResponseImageSourceType `json:"__image_source_type,omitempty"`
	// 镜像类型，目前支持以下类型： 公共镜像：gold 私有镜像：private 共享镜像：shared
	Imagetype UpdateImageResponseImagetype `json:"__imagetype,omitempty"`
	// 是否完成了初始化配置。取值为true或false。如果用户确定完成了初始化配置，则可以设置为true，否            则设置为false。默认为false。
	IsConfigInit UpdateImageResponseIsConfigInit `json:"__is_config_init,omitempty"`
	// 是否是注册过的镜像，取值为“true”或者“false”
	Isregistered UpdateImageResponseIsregistered `json:"__isregistered,omitempty"`
	// 父镜像ID。公共镜像或通过文件创建的私有镜像，取值为空
	Originalimagename string `json:"__originalimagename,omitempty"`
	// 操作系统位数，一般取值为“32”或者“64”
	OsBit UpdateImageResponseOsBit `json:"__os_bit,omitempty"`
	// 操作系统类型，目前取值Linux， Windows，Other
	OsType UpdateImageResponseOsType `json:"__os_type,omitempty"`
	// 操作系统具体版本
	OsVersion string `json:"__os_version,omitempty"`
	// 镜像平台分类
	Platform UpdateImageResponsePlatform `json:"__platform,omitempty"`
	// 市场镜像的产品ID
	Productcode string `json:"__productcode,omitempty"`
	// 镜像来源表示该镜像支持密集存储。如果镜像支持密集存储性能，则值为true，否则无需增加该属性。
	SupportDiskintensive string `json:"__support_diskintensive,omitempty"`
	// 表示该镜像支持高计算性能。如果镜像支持高计算性能，则值为true，否则无需增加该属性。
	SupportHighperformance string `json:"__support_highperformance,omitempty"`
	// 如果镜像支持KVM，取值为true，否则无需增加该属性。
	SupportKvm string `json:"__support_kvm,omitempty"`
	// 表示该镜像是支持KVM虚拟化平台下的GPU类型，如果不支持KVM虚拟机下GPU类型，无需添加该属性。该属性与“__support_xen”和“__support_kvm”属性不共存。
	SupportKvmGpuType string `json:"__support_kvm_gpu_type,omitempty"`
	// 如果镜像支持KVM虚拟化下Infiniband网卡类型，取值为true。否则，无需添加该属性。该属性与“__support_xen”属性不共存。
	SupportKvmInfiniband string `json:"__support_kvm_infiniband,omitempty"`
	// 表示该镜像支持超大内存。如果镜像支持超大内存，取值为true，否则无需增加该属性。
	SupportLargememory string `json:"__support_largememory,omitempty"`
	// 如果镜像支持XEN，取值为true，否则无需增加该属性。
	SupportXen string `json:"__support_xen,omitempty"`
	// 表示该镜像是支持XEN虚拟化平台下的GPU优化类型，如果不支持XEN虚拟化下GPU类型，无需添加该属性            。该属性与“__support_xen”和“__support_kvm”属性不共存。
	SupportXenGpuType string `json:"__support_xen_gpu_type,omitempty"`
	// 如果镜像支持XEN虚拟化下HANA类型，取值为true。否则，无需添加该属性。该属性与“__support_xen”             和“__support_kvm”属性不共存。
	SupportXenHana string `json:"__support_xen_hana,omitempty"`
	// 表示当前镜像是否支持发布为市场镜像,true表示支持,false 表示不支持
	SystemSupportMarket bool `json:"__system_support_market,omitempty"`
	// 目前暂时不使用
	Checksum string `json:"checksum,omitempty"`
	// 容器类型
	ContainerFormat string `json:"container_format,omitempty"`
	// 创建时间。格式为UTC时间
	CreatedAt string `json:"created_at,omitempty"`
	// 是否是删除的镜像，取值为true或者false
	Deleted bool `json:"deleted,omitempty"`
	// 删除时间。格式为UTC时间
	DeletedAt string `json:"deleted_at,omitempty"`
	// 镜像的格式，目前支持vhd，zvhd、raw，qcow2。默认值是vhd
	DiskFormat string `json:"disk_format,omitempty"`
	// 表示当前镜像所属的企业项目。取值为0或无该值，表示属于default企业项目，取值为UUID，表示属于             该UUID对应的企业项目。
	EnterpriseProjectId string `json:"enterprise_project_id,omitempty"`
	// 镜像文件下载和上传链接
	File string `json:"file,omitempty"`
	// 镜像ID
	Id string `json:"id,omitempty"`
	// 镜像运行需要的最小磁盘容量，单位为GB。取值为40～1024GB。
	MinDisk int32 `json:"min_disk,omitempty"`
	// 镜像运行需要的最小内存，单位为MB。参数取值依据弹性云服务器的规格限制，默认设置为0
	MinRam int32 `json:"min_ram,omitempty"`
	// 镜像名称。 名称的首尾字母不能为空格。 名称的长度至为1～128位。 名称包含以下4种字符： 大写字母 小写字母 数字 特殊字符包含-、.、_、空格和中文。
	Name string `json:"name,omitempty"`
	// 镜像属于哪个租户
	Owner string `json:"owner,omitempty"`
	// 是否是受保护的，受保护的镜像不允许删除。取值为true或false
	Protected bool `json:"protected,omitempty"`
	// 镜像视图
	Schema string `json:"schema,omitempty"`
	// 镜像链接信息
	Self string `json:"self,omitempty"`
	// 目前暂时不使用
	Size int32 `json:"size,omitempty"`
	// 镜像状态。取值如下：queued：表示镜像元数据已经创建成功，等待 上传镜像文件。saving：表示镜像 正在上传文件到后端存储。deleted：表示镜像已经删除。killed：表示镜像上传错误。active：表示镜像可以正常使用
	Status UpdateImageResponseStatus `json:"status,omitempty"`
	// 镜像标签列表
	Tags []string `json:"tags,omitempty"`
	// 更新时间。格式为UTC时间
	UpdatedAt string `json:"updated_at,omitempty"`
	// 镜像使用环境类型：FusionCompute，Ironic，DataImage。如果弹性云服务器镜像，则取值为FusionCompute，如果是数据卷镜像则取Dat            aImage，如果是裸金属服务器镜像，则取值是Ironic
	VirtualEnvType UpdateImageResponseVirtualEnvType `json:"virtual_env_type,omitempty"`
	// 目前暂时不使用
	VirtualSize int32 `json:"virtual_size,omitempty"`
	// 是否被其他租户可见，取值为private或public
	Visibility UpdateImageResponseVisibility `json:"visibility,omitempty"`
	// 镜像架构类型。取值包括： x86 arm
	Architecture UpdateImageResponseArchitecture `json:"architecture,omitempty"`
	// 表示当前镜像支持CloudInit密码/密钥注入方式，建议设置为\"true\"或者\"false\"。 如果取值为\"true\"，表示该镜像不支持CloudInit注入密码/密钥，其他取值时表示支持CloudInit注入密钥/密码。
	SupportFcInject UpdateImageResponseSupportFcInject `json:"__support_fc_inject,omitempty"`
	// 云服务器的启动方式。目前支持： bios：表示bios引导启动。 uefi：表示uefi引导启动。
	HwFirmwareType UpdateImageResponseHwFirmwareType `json:"hw_firmware_type,omitempty"`
	// 是否是ARM架构类型的镜像，取值为“true”或者“false”。
	SupportArm UpdateImageResponseSupportArm `json:"__support_arm,omitempty"`
}

func (o UpdateImageResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateImageResponse", string(data)}, " ")
}

type UpdateImageResponseImageSourceType struct {
	value string
}

type UpdateImageResponseImageSourceTypeEnum struct {
	UDS   UpdateImageResponseImageSourceType
	SWIFT UpdateImageResponseImageSourceType
}

func GetUpdateImageResponseImageSourceTypeEnum() UpdateImageResponseImageSourceTypeEnum {
	return UpdateImageResponseImageSourceTypeEnum{
		UDS: UpdateImageResponseImageSourceType{
			value: "uds",
		},
		SWIFT: UpdateImageResponseImageSourceType{
			value: "swift",
		},
	}
}

func (c UpdateImageResponseImageSourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *UpdateImageResponseImageSourceType) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}

type UpdateImageResponseImagetype struct {
	value string
}

type UpdateImageResponseImagetypeEnum struct {
	GOLD    UpdateImageResponseImagetype
	PRIVATE UpdateImageResponseImagetype
	SHARED  UpdateImageResponseImagetype
}

func GetUpdateImageResponseImagetypeEnum() UpdateImageResponseImagetypeEnum {
	return UpdateImageResponseImagetypeEnum{
		GOLD: UpdateImageResponseImagetype{
			value: "gold",
		},
		PRIVATE: UpdateImageResponseImagetype{
			value: "private",
		},
		SHARED: UpdateImageResponseImagetype{
			value: "shared",
		},
	}
}

func (c UpdateImageResponseImagetype) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *UpdateImageResponseImagetype) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}

type UpdateImageResponseIsConfigInit struct {
	value string
}

type UpdateImageResponseIsConfigInitEnum struct {
	TRUE  UpdateImageResponseIsConfigInit
	FALSE UpdateImageResponseIsConfigInit
}

func GetUpdateImageResponseIsConfigInitEnum() UpdateImageResponseIsConfigInitEnum {
	return UpdateImageResponseIsConfigInitEnum{
		TRUE: UpdateImageResponseIsConfigInit{
			value: "true",
		},
		FALSE: UpdateImageResponseIsConfigInit{
			value: "false",
		},
	}
}

func (c UpdateImageResponseIsConfigInit) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *UpdateImageResponseIsConfigInit) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}

type UpdateImageResponseIsregistered struct {
	value string
}

type UpdateImageResponseIsregisteredEnum struct {
	TRUE  UpdateImageResponseIsregistered
	FALSE UpdateImageResponseIsregistered
}

func GetUpdateImageResponseIsregisteredEnum() UpdateImageResponseIsregisteredEnum {
	return UpdateImageResponseIsregisteredEnum{
		TRUE: UpdateImageResponseIsregistered{
			value: "true",
		},
		FALSE: UpdateImageResponseIsregistered{
			value: "false",
		},
	}
}

func (c UpdateImageResponseIsregistered) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *UpdateImageResponseIsregistered) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}

type UpdateImageResponseOsBit struct {
	value string
}

type UpdateImageResponseOsBitEnum struct {
	_32 UpdateImageResponseOsBit
	_64 UpdateImageResponseOsBit
}

func GetUpdateImageResponseOsBitEnum() UpdateImageResponseOsBitEnum {
	return UpdateImageResponseOsBitEnum{
		_32: UpdateImageResponseOsBit{
			value: "32",
		},
		_64: UpdateImageResponseOsBit{
			value: "64",
		},
	}
}

func (c UpdateImageResponseOsBit) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *UpdateImageResponseOsBit) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}

type UpdateImageResponseOsType struct {
	value string
}

type UpdateImageResponseOsTypeEnum struct {
	LINUX   UpdateImageResponseOsType
	WINDOWS UpdateImageResponseOsType
	OTHER   UpdateImageResponseOsType
}

func GetUpdateImageResponseOsTypeEnum() UpdateImageResponseOsTypeEnum {
	return UpdateImageResponseOsTypeEnum{
		LINUX: UpdateImageResponseOsType{
			value: "Linux",
		},
		WINDOWS: UpdateImageResponseOsType{
			value: "Windows",
		},
		OTHER: UpdateImageResponseOsType{
			value: "Other",
		},
	}
}

func (c UpdateImageResponseOsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *UpdateImageResponseOsType) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}

type UpdateImageResponsePlatform struct {
	value string
}

type UpdateImageResponsePlatformEnum struct {
	WINDOWS      UpdateImageResponsePlatform
	UBUNTU       UpdateImageResponsePlatform
	RED_HAT      UpdateImageResponsePlatform
	SUSE         UpdateImageResponsePlatform
	CENT_OS      UpdateImageResponsePlatform
	DEBIAN       UpdateImageResponsePlatform
	OPEN_SUSE    UpdateImageResponsePlatform
	ORACLE_LINUX UpdateImageResponsePlatform
	FEDORA       UpdateImageResponsePlatform
	OTHER        UpdateImageResponsePlatform
	CORE_OS      UpdateImageResponsePlatform
	EULE_OS      UpdateImageResponsePlatform
}

func GetUpdateImageResponsePlatformEnum() UpdateImageResponsePlatformEnum {
	return UpdateImageResponsePlatformEnum{
		WINDOWS: UpdateImageResponsePlatform{
			value: "Windows",
		},
		UBUNTU: UpdateImageResponsePlatform{
			value: "Ubuntu",
		},
		RED_HAT: UpdateImageResponsePlatform{
			value: "RedHat",
		},
		SUSE: UpdateImageResponsePlatform{
			value: "SUSE",
		},
		CENT_OS: UpdateImageResponsePlatform{
			value: "CentOS",
		},
		DEBIAN: UpdateImageResponsePlatform{
			value: "Debian",
		},
		OPEN_SUSE: UpdateImageResponsePlatform{
			value: "OpenSUSE",
		},
		ORACLE_LINUX: UpdateImageResponsePlatform{
			value: "Oracle Linux",
		},
		FEDORA: UpdateImageResponsePlatform{
			value: "Fedora",
		},
		OTHER: UpdateImageResponsePlatform{
			value: "Other",
		},
		CORE_OS: UpdateImageResponsePlatform{
			value: "CoreOS",
		},
		EULE_OS: UpdateImageResponsePlatform{
			value: "EuleOS",
		},
	}
}

func (c UpdateImageResponsePlatform) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *UpdateImageResponsePlatform) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}

type UpdateImageResponseStatus struct {
	value string
}

type UpdateImageResponseStatusEnum struct {
	QUEUED  UpdateImageResponseStatus
	SAVING  UpdateImageResponseStatus
	DELETED UpdateImageResponseStatus
	KILLED  UpdateImageResponseStatus
	ACTIVE  UpdateImageResponseStatus
}

func GetUpdateImageResponseStatusEnum() UpdateImageResponseStatusEnum {
	return UpdateImageResponseStatusEnum{
		QUEUED: UpdateImageResponseStatus{
			value: "queued",
		},
		SAVING: UpdateImageResponseStatus{
			value: "saving",
		},
		DELETED: UpdateImageResponseStatus{
			value: "deleted",
		},
		KILLED: UpdateImageResponseStatus{
			value: "killed",
		},
		ACTIVE: UpdateImageResponseStatus{
			value: "active",
		},
	}
}

func (c UpdateImageResponseStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *UpdateImageResponseStatus) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}

type UpdateImageResponseVirtualEnvType struct {
	value string
}

type UpdateImageResponseVirtualEnvTypeEnum struct {
	FUSION_COMPUTE UpdateImageResponseVirtualEnvType
	IRONIC         UpdateImageResponseVirtualEnvType
	DATA_IMAGE     UpdateImageResponseVirtualEnvType
}

func GetUpdateImageResponseVirtualEnvTypeEnum() UpdateImageResponseVirtualEnvTypeEnum {
	return UpdateImageResponseVirtualEnvTypeEnum{
		FUSION_COMPUTE: UpdateImageResponseVirtualEnvType{
			value: "FusionCompute",
		},
		IRONIC: UpdateImageResponseVirtualEnvType{
			value: "Ironic",
		},
		DATA_IMAGE: UpdateImageResponseVirtualEnvType{
			value: "DataImage",
		},
	}
}

func (c UpdateImageResponseVirtualEnvType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *UpdateImageResponseVirtualEnvType) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}

type UpdateImageResponseVisibility struct {
	value string
}

type UpdateImageResponseVisibilityEnum struct {
	PRIVATE UpdateImageResponseVisibility
	PUBLIC  UpdateImageResponseVisibility
}

func GetUpdateImageResponseVisibilityEnum() UpdateImageResponseVisibilityEnum {
	return UpdateImageResponseVisibilityEnum{
		PRIVATE: UpdateImageResponseVisibility{
			value: "private",
		},
		PUBLIC: UpdateImageResponseVisibility{
			value: "public",
		},
	}
}

func (c UpdateImageResponseVisibility) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *UpdateImageResponseVisibility) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}

type UpdateImageResponseArchitecture struct {
	value string
}

type UpdateImageResponseArchitectureEnum struct {
	X86 UpdateImageResponseArchitecture
	ARM UpdateImageResponseArchitecture
}

func GetUpdateImageResponseArchitectureEnum() UpdateImageResponseArchitectureEnum {
	return UpdateImageResponseArchitectureEnum{
		X86: UpdateImageResponseArchitecture{
			value: "x86",
		},
		ARM: UpdateImageResponseArchitecture{
			value: "arm",
		},
	}
}

func (c UpdateImageResponseArchitecture) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *UpdateImageResponseArchitecture) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}

type UpdateImageResponseSupportFcInject struct {
	value string
}

type UpdateImageResponseSupportFcInjectEnum struct {
	TRUE  UpdateImageResponseSupportFcInject
	FALSE UpdateImageResponseSupportFcInject
}

func GetUpdateImageResponseSupportFcInjectEnum() UpdateImageResponseSupportFcInjectEnum {
	return UpdateImageResponseSupportFcInjectEnum{
		TRUE: UpdateImageResponseSupportFcInject{
			value: "true",
		},
		FALSE: UpdateImageResponseSupportFcInject{
			value: "false",
		},
	}
}

func (c UpdateImageResponseSupportFcInject) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *UpdateImageResponseSupportFcInject) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}

type UpdateImageResponseHwFirmwareType struct {
	value string
}

type UpdateImageResponseHwFirmwareTypeEnum struct {
	BIOS UpdateImageResponseHwFirmwareType
	UEFI UpdateImageResponseHwFirmwareType
}

func GetUpdateImageResponseHwFirmwareTypeEnum() UpdateImageResponseHwFirmwareTypeEnum {
	return UpdateImageResponseHwFirmwareTypeEnum{
		BIOS: UpdateImageResponseHwFirmwareType{
			value: "bios",
		},
		UEFI: UpdateImageResponseHwFirmwareType{
			value: "uefi",
		},
	}
}

func (c UpdateImageResponseHwFirmwareType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *UpdateImageResponseHwFirmwareType) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}

type UpdateImageResponseSupportArm struct {
	value string
}

type UpdateImageResponseSupportArmEnum struct {
	TRUE  UpdateImageResponseSupportArm
	FALSE UpdateImageResponseSupportArm
}

func GetUpdateImageResponseSupportArmEnum() UpdateImageResponseSupportArmEnum {
	return UpdateImageResponseSupportArmEnum{
		TRUE: UpdateImageResponseSupportArm{
			value: "true",
		},
		FALSE: UpdateImageResponseSupportArm{
			value: "false",
		},
	}
}

func (c UpdateImageResponseSupportArm) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *UpdateImageResponseSupportArm) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}
