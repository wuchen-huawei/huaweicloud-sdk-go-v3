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

// 参数模板信息。
type ConfigurationSummaryForCreate struct {
	// 参数组ID。
	Id string `json:"id"`
	// 参数组名称。
	Name string `json:"name"`
	// 参数组描述。
	Description *string `json:"description,omitempty"`
	// 引擎版本。
	DatastoreVersionName string `json:"datastore_version_name"`
	// 引擎名。
	DatastoreName ConfigurationSummaryForCreateDatastoreName `json:"datastore_name"`
	// 创建时间，格式为\"yyyy-MM-ddTHH:mm:ssZ\"。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	Created string `json:"created"`
	// 更新时间，格式为\"yyyy-MM-ddTHH:mm:ssZ\"。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	Updated string `json:"updated"`
}

func (o ConfigurationSummaryForCreate) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ConfigurationSummaryForCreate", string(data)}, " ")
}

type ConfigurationSummaryForCreateDatastoreName struct {
	value string
}

type ConfigurationSummaryForCreateDatastoreNameEnum struct {
	MYSQL      ConfigurationSummaryForCreateDatastoreName
	POSTGRESQL ConfigurationSummaryForCreateDatastoreName
	SQLSERVER  ConfigurationSummaryForCreateDatastoreName
}

func GetConfigurationSummaryForCreateDatastoreNameEnum() ConfigurationSummaryForCreateDatastoreNameEnum {
	return ConfigurationSummaryForCreateDatastoreNameEnum{
		MYSQL: ConfigurationSummaryForCreateDatastoreName{
			value: "mysql",
		},
		POSTGRESQL: ConfigurationSummaryForCreateDatastoreName{
			value: "postgresql",
		},
		SQLSERVER: ConfigurationSummaryForCreateDatastoreName{
			value: "sqlserver",
		},
	}
}

func (c ConfigurationSummaryForCreateDatastoreName) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ConfigurationSummaryForCreateDatastoreName) UnmarshalJSON(b []byte) error {
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
