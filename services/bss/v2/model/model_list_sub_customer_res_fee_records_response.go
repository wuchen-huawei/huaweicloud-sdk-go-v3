/*
 * BSS
 *
 * Business Support System API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListSubCustomerResFeeRecordsResponse struct {
	// |参数名称：资源费用记录数据。具体请参见表 ResFeeRecordV2。| |参数约束以及描述：资源费用记录数据。具体请参见表 ResFeeRecordV2。|
	FeeRecords *[]SubCustomerResFeeRecordV2 `json:"fee_records,omitempty"`
	// |参数名称：结果集数量，只有成功才返回这个参数。| |参数的约束及描述：结果集数量，只有成功才返回这个参数。|
	Count *int32 `json:"count,omitempty"`
	// |参数名称：货币单位代码：CNY：人民币USD：美元| |参数约束及描述：货币单位代码：CNY：人民币USD：美元|
	Currency       *string `json:"currency,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListSubCustomerResFeeRecordsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListSubCustomerResFeeRecordsResponse", string(data)}, " ")
}
