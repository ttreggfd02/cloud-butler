package scanner

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/ttreggfd02/cloud-butler/pkg/aws_session"
)

// PublicBucketResult 用來存放公開 S3 儲存桶的掃描結果
type PublicBucketResult struct {
	BucketName string
	Region     string
	Issue      string
}

// ScanPublicS3Buckets 掃描公開存取的 S3 儲存桶
func ScanPublicS3Buckets(ctx context.Context, cfg aws.Config) ([]PublicBucketResult, error) {
	client := s3.NewFromConfig(cfg)
	var results []PublicBucketResult

	fmt.Println("正在呼叫 AWS API 取得 S3 儲存桶列表...")
	listBucketsOutput, err := client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		return nil, fmt.Errorf("無法取得 S3 儲存桶列表: %w", err)
	}

	fmt.Printf("API 呼叫完成，將檢查 %d 個儲存桶的權限...\n", len(listBucketsOutput.Buckets))

	for _, bucket := range listBucketsOutput.Buckets {
		bucketName := aws.ToString(bucket.Name)

		locationOutput, err := client.GetBucketLocation(ctx, &s3.GetBucketLocationInput{
			Bucket: bucket.Name,
		})
		if err != nil {
			fmt.Printf("警告：無法取得儲存桶 %s 的位置，已跳過: %v\n", bucketName, err)
			continue
		}
		region := string(locationOutput.LocationConstraint)
		if region == "" {
			region = "us-east-1"
		}

		regionalCfg, err := aws_session.NewConfigWithRegion(ctx, region)
		if err != nil {
			fmt.Printf("警告：無法為區域 %s 建立 S3 客戶端，已跳過儲存桶 %s: %v\n", region, bucketName, err)
			continue
		}
		regionalClient := s3.NewFromConfig(regionalCfg)

		pabOutput, err := regionalClient.GetPublicAccessBlock(ctx, &s3.GetPublicAccessBlockInput{
			Bucket: bucket.Name,
		})

		var nspab *types.NoSuchPublicAccessBlockConfiguration
		if errors.As(err, &nspab) {
			results = append(results, PublicBucketResult{
				BucketName: bucketName,
				Region:     region,
				Issue:      "未設定公開存取區塊 (Public Access is not blocked)",
			})
			continue
		} else if err != nil {
			fmt.Printf("警告：檢查儲存桶 %s 的權限時發生非預期錯誤，已跳過: %v\n", bucketName, err)
			continue
		}

		config := pabOutput.PublicAccessBlockConfiguration
		if !aws.ToBool(config.BlockPublicAcls) || !aws.ToBool(config.BlockPublicPolicy) || !aws.ToBool(config.IgnorePublicAcls) || !aws.ToBool(config.RestrictPublicBuckets) {
			results = append(results, PublicBucketResult{
				BucketName: bucketName,
				Region:     region,
				Issue:      "部分或所有公開存取未被封鎖 (Public access block is not fully enabled)",
			})
		}
	}

	return results, nil
}
