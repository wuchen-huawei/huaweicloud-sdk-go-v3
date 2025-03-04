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

// Response Object
type DownloadSlowlogResponse struct {
	// 慢日志下载链接列表
	List *[]SlowlogDownloadInfo `json:"list,omitempty"`
	// - 慢日志下载链接生成状态。FINISH，表示下载链接已经生成完成。CREATING，表示正在生成文件，准备下载链接。FAILED，表示存在日志文件准备失败。
	Status *string `json:"status,omitempty"`
	// - 慢日志链接数量。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o DownloadSlowlogResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DownloadSlowlogResponse", string(data)}, " ")
}
