// internal/scanner/ebs.go

package scanner

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

// UnattachedVolumeResult 用來存放掃描結果
type UnattachedVolumeResult struct {
	VolumeID string
	Size     int32
	Created  string
}

// ScanUnattachedVolumes 掃描指定區域中所有未掛載的 EBS 磁碟區
func ScanUnattachedVolumes(ctx context.Context, cfg aws.Config) ([]UnattachedVolumeResult, error) {
	// 建立 EC2 客戶端
	client := ec2.NewFromConfig(cfg)

	var results []UnattachedVolumeResult

	fmt.Println("正在呼叫 AWS API 取得 EBS 磁碟區資訊...")

	// AWS 的列表 API 通常有分頁，我們需要使用 Paginator 來處理
	paginator := ec2.NewDescribeVolumesPaginator(client, &ec2.DescribeVolumesInput{})

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("無法取得 EBS 磁碟區頁面: %w", err)
		}

		for _, volume := range page.Volumes {
			// 核心邏輯：如果磁碟區的狀態是 'available'，代表它沒有被掛載
			if volume.State == types.VolumeStateAvailable {
				results = append(results, UnattachedVolumeResult{
					VolumeID: *volume.VolumeId,
					Size:     *volume.Size,
					Created:  volume.CreateTime.Format("2006-01-02"),
				})
			}
		}
	}

	fmt.Printf("API 呼叫完成，發現 %d 個未掛載的磁碟區。\n", len(results))
	return results, nil
}
