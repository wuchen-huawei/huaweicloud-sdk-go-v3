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

// Response Object
type ShowConfigurationResponse struct {
	// 参数组ID。
	Id *string `json:"id,omitempty"`
	// 参数组名称。
	Name *string `json:"name,omitempty"`
	// 参数组描述。
	Description *string `json:"description,omitempty"`
	// 引擎版本。
	DatastoreVersionName *string `json:"datastore_version_name,omitempty"`
	// 引擎名。
	DatastoreName *ShowConfigurationResponseDatastoreName `json:"datastore_name,omitempty"`
	// 创建时间，格式为\"yyyy-MM-ddTHH:mm:ssZ\"。  其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	Created *string `json:"created,omitempty"`
	// 更新时间，格式为\"yyyy-MM-ddTHH:mm:ssZ\"。  其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	Updated *string `json:"updated,omitempty"`
	// 参数对象，用户基于默认参数模板自定义的参数配置。
	ConfigurationParameters *[]ConfigurationParameter `json:"configuration_parameters,omitempty"`
	HttpStatusCode          int                       `json:"-"`
}

func (o ShowConfigurationResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowConfigurationResponse", string(data)}, " ")
}

type ShowConfigurationResponseDatastoreName struct {
	value string
}

type ShowConfigurationResponseDatastoreNameEnum struct {
	MYSQL      ShowConfigurationResponseDatastoreName
	POSTGRESQL ShowConfigurationResponseDatastoreName
	SQLSERVER  ShowConfigurationResponseDatastoreName
}

func GetShowConfigurationResponseDatastoreNameEnum() ShowConfigurationResponseDatastoreNameEnum {
	return ShowConfigurationResponseDatastoreNameEnum{
		MYSQL: ShowConfigurationResponseDatastoreName{
			value: "mysql",
		},
		POSTGRESQL: ShowConfigurationResponseDatastoreName{
			value: "postgresql",
		},
		SQLSERVER: ShowConfigurationResponseDatastoreName{
			value: "sqlserver",
		},
	}
}

func (c ShowConfigurationResponseDatastoreName) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ShowConfigurationResponseDatastoreName) UnmarshalJSON(b []byte) error {
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
