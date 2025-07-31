// internal/scanner/ebs.go

package scanner

import (
	"context"
	"fmt"
        "time"
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
// OldSnapshotResult 用來存放過舊快照的掃描結果
type OldSnapshotResult struct {
	SnapshotID string
	VolumeID   string
	StartTime  string
	VolumeSize int32
}

// ScanOldSnapshots 掃描過舊的 EBS 快照
func ScanOldSnapshots(ctx context.Context, cfg aws.Config, minDaysOld int) ([]OldSnapshotResult, error) {
	client := ec2.NewFromConfig(cfg)
	var results []OldSnapshotResult

	fmt.Println("正在呼叫 AWS API 取得 EBS 快照資訊...")

	// 計算截止時間
	cutoffTime := time.Now().AddDate(0, 0, -minDaysOld)

	// 重要：我們只掃描自己帳號下的快照，所以要設定 OwnerIds 為 "self"
	// 否則 API 會試圖列出所有公開快照，非常慢且量大
	input := &ec2.DescribeSnapshotsInput{
		OwnerIds: []string{"self"},
	}

	paginator := ec2.NewDescribeSnapshotsPaginator(client, input)

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("無法取得 EBS 快照頁面: %w", err)
		}

		for _, snapshot := range page.Snapshots {
			// 核心邏輯：如果快照的建立時間早於我們的截止時間，就將其視為過舊
			if snapshot.StartTime.Before(cutoffTime) {
				results = append(results, OldSnapshotResult{
					SnapshotID: *snapshot.SnapshotId,
					VolumeID:   aws.ToString(snapshot.VolumeId), // VolumeId 可能是 nil，用 aws.ToString 較安全
					StartTime:  snapshot.StartTime.Format("2006-01-02"),
					VolumeSize: *snapshot.VolumeSize,
				})
			}
		}
	}

	fmt.Printf("API 呼叫完成，發現 %d 個過舊的快照。\n", len(results))
	return results, nil
}
func DeleteUnattachedVolume(ctx context.Context, cfg aws.Config, volumeID string) error {
	client := ec2.NewFromConfig(cfg)
	_, err := client.DeleteVolume(ctx, &ec2.DeleteVolumeInput{
		VolumeId: aws.String(volumeID),
	})
	if err != nil {
		return fmt.Errorf("無法刪除 volume %s: %w", volumeID, err)
	}
	return nil
}

// DeleteOldSnapshot 刪除指定的 EBS 快照
func DeleteOldSnapshot(ctx context.Context, cfg aws.Config, snapshotID string) error {
	client := ec2.NewFromConfig(cfg)
	_, err := client.DeleteSnapshot(ctx, &ec2.DeleteSnapshotInput{
		SnapshotId: aws.String(snapshotID),
	})
	if err != nil {
		return fmt.Errorf("無法刪除 snapshot %s: %w", snapshotID, err)
	}
	return nil
}
