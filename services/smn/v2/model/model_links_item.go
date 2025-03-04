/*
 * SMN
 *
 * SMN Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

type LinksItem struct {
	// 对应快捷链接。
	Href string `json:"href"`
	// 快捷链接标记名称。
	Rel string `json:"rel"`
}

func (o LinksItem) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"LinksItem", string(data)}, " ")
}
