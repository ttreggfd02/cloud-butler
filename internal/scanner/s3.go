package scanner

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/smithy-go" // Import the smithy-go package
	"github.com/ttreggfd02/cloud-butler/pkg/aws_session"
)

// PublicBucketResult stores the results of the S3 public bucket scan.
type PublicBucketResult struct {
	BucketName string
	Region     string
	Issue      string
}

// ScanPublicS3Buckets scans for publicly accessible S3 buckets.
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
		// The location constraint for us-east-1 is an empty string.
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

		if err != nil {
			var apiErr smithy.APIError
			// Check if the error is an APIError and if its code matches the specific error we're looking for.
			if errors.As(err, &apiErr) && apiErr.ErrorCode() == "NoSuchPublicAccessBlockConfiguration" {
				results = append(results, PublicBucketResult{
					BucketName: bucketName,
					Region:     region,
					Issue:      "未設定公開存取區塊 (Public Access is not blocked)",
				})
			} else {
				// Handle other potential errors during GetPublicAccessBlock
				fmt.Printf("警告：檢查儲存桶 %s 的權限時發生非預期錯誤，已跳過: %v\n", bucketName, err)
			}
			continue
		}

		// If there was no error, pabOutput is valid and we can check the configuration.
		if pabOutput.PublicAccessBlockConfiguration != nil {
			config := pabOutput.PublicAccessBlockConfiguration
			if !aws.ToBool(config.BlockPublicAcls) || !aws.ToBool(config.BlockPublicPolicy) || !aws.ToBool(config.IgnorePublicAcls) || !aws.ToBool(config.RestrictPublicBuckets) {
				results = append(results, PublicBucketResult{
					BucketName: bucketName,
					Region:     region,
					Issue:      "部分或所有公開存取未被封鎖 (Public access block is not fully enabled)",
				})
			}
		}
	}

	return results, nil
}
