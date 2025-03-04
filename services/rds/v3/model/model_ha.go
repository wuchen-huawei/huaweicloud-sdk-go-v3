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

// HA配置参数，创建HA实例时使用。
type Ha struct {
	// 实例主备模式，取值：Ha（主备），不区分大小写。
	Mode HaMode `json:"mode"`
	// 备机同步参数。实例主备模式为Ha时有效。 取值： - MySQL为“async”或“semisync”。 - PostgreSQL为“async”或“sync”。 - Microsoft SQL Server为“sync”。
	ReplicationMode HaReplicationMode `json:"replication_mode"`
}

func (o Ha) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"Ha", string(data)}, " ")
}

type HaMode struct {
	value string
}

type HaModeEnum struct {
	HA     HaMode
	SINGLE HaMode
}

func GetHaModeEnum() HaModeEnum {
	return HaModeEnum{
		HA: HaMode{
			value: "Ha",
		},
		SINGLE: HaMode{
			value: "Single",
		},
	}
}

func (c HaMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *HaMode) UnmarshalJSON(b []byte) error {
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

type HaReplicationMode struct {
	value string
}

type HaReplicationModeEnum struct {
	ASYNC    HaReplicationMode
	SEMISYNC HaReplicationMode
	SYNC     HaReplicationMode
}

func GetHaReplicationModeEnum() HaReplicationModeEnum {
	return HaReplicationModeEnum{
		ASYNC: HaReplicationMode{
			value: "async",
		},
		SEMISYNC: HaReplicationMode{
			value: "semisync",
		},
		SYNC: HaReplicationMode{
			value: "sync",
		},
	}
}

func (c HaReplicationMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *HaReplicationMode) UnmarshalJSON(b []byte) error {
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
