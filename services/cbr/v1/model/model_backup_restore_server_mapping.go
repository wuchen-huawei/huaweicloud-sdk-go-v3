/*
 * Cbr
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

type BackupRestoreServerMapping struct {
	// 卷备份ID，可以通过控制台或者“查询指定备份”接口获取。
	BackupId string `json:"backup_id"`
	// 待恢复目标卷ID
	VolumeId string `json:"volume_id"`
}

func (o BackupRestoreServerMapping) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BackupRestoreServerMapping", string(data)}, " ")
}
