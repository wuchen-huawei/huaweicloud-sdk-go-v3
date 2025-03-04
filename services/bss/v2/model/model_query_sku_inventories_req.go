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

type QuerySkuInventoriesReq struct {
	// |参数名称：待查询库存项| |参数约束以及描述：待查询库存项|
	SkuItems []SkuItem `json:"sku_items"`
}

func (o QuerySkuInventoriesReq) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"QuerySkuInventoriesReq", string(data)}, " ")
}
