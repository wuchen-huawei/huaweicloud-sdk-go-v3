/*
 * ECS
 *
 * ECS Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type NovaListServersDetailsResponse struct {
	// 查询云服务器信息列表。
	Servers *[]NovaServer `json:"servers,omitempty"`
	// 分页查询时，查询下一页数据链接。
	ServersLinks   *[]PageLink `json:"servers_links,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o NovaListServersDetailsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"NovaListServersDetailsResponse", string(data)}, " ")
}
