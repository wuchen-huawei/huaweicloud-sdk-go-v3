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

type SlowlogDownloadInfo struct {
	// 任务ID
	WorkflowId string `json:"workflow_id"`
	// 生成的下载文件名
	FileName string `json:"file_name"`
	// 生成链接的生成状态
	Status string `json:"status"`
	// 文件大小
	FileSize string `json:"file_size"`
	// 下载链接
	FileLink string `json:"file_link"`
	// 生成时间
	CreateAt int64 `json:"create_at"`
	// 更新时间
	UpdateAt int64 `json:"update_at"`
}

func (o SlowlogDownloadInfo) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"SlowlogDownloadInfo", string(data)}, " ")
}
