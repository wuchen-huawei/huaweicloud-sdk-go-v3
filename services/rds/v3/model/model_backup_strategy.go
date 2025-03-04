/*
 * RDS
 *
 * API v3
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// 自动备份策略。
type BackupStrategy struct {
	// 备份时间段。自动备份将在该时间段内触发。  取值范围：非空，格式必须为hh:mm-HH:MM且有效，当前时间指UTC时间。  HH取值必须比hh大1。 mm和MM取值必须相同，且取值必须为00、15、30或45。
	StartTime string `json:"start_time"`
	// 指定备份文件的可保存天数。  取值范围：0～732。该参数缺省，或取值为0，表示关闭自动备份策略。如果需要延长保留时间请联系客服人员申请，自动备份最长可以保留2562天。  说明：SQL Server的HA实例不支持关闭自动备份策略。
	KeepDays *int32 `json:"keep_days,omitempty"`
}

func (o BackupStrategy) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BackupStrategy", string(data)}, " ")
}
