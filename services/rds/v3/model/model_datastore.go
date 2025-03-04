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

type Datastore struct {
	// 数据库引擎，不区分大小写： - MySQL - PostgreSQL - SQLServer
	Type DatastoreType `json:"type"`
	// 数据库版本。  - MySQL引擎支持5.6、5.7、8.0版本。取值示例：5.7。具有相应权限的用户才可使用8.0，您可联系华为云客服人员申请。 - PostgreSQL引擎支持9.5、9.6、10、11版本。取值示例：9.6。 - Microsoft SQL Server：仅支持2017 企业版、2017 标准版、2017 web版、2014 标准版、2014 企业版、2016 标准版、2016 企业版、2012 企业版、2012 标准版、2012 web版、2008 R2 企业版、2008 R2 web版、2014 web版、2016 web版。取值示例2014_SE。
	Version string `json:"version"`
}

func (o Datastore) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"Datastore", string(data)}, " ")
}

type DatastoreType struct {
	value string
}

type DatastoreTypeEnum struct {
	MY_SQL      DatastoreType
	POSTGRE_SQL DatastoreType
	SQL_SERVER  DatastoreType
}

func GetDatastoreTypeEnum() DatastoreTypeEnum {
	return DatastoreTypeEnum{
		MY_SQL: DatastoreType{
			value: "MySQL",
		},
		POSTGRE_SQL: DatastoreType{
			value: "PostgreSQL",
		},
		SQL_SERVER: DatastoreType{
			value: "SQLServer",
		},
	}
}

func (c DatastoreType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *DatastoreType) UnmarshalJSON(b []byte) error {
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
