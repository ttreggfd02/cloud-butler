// pkg/aws_session/session.go

package aws_session

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/spf13/viper"
)

// NewConfig 建立並回傳一個新的 AWS 組態設定
func NewConfig(ctx context.Context) (aws.Config, error) {
	// 從 viper 讀取 AWS 相關設定
	region := viper.GetString("aws.region")
	accessKeyID := viper.GetString("aws.accessKeyID")
	secretAccessKey := viper.GetString("aws.secretAccessKey")

	// 使用 config.LoadDefaultConfig 來載入預設設定
	// 並客製化憑證和區域
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(region),
		config.WithCredentialsProvider(aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
			return aws.Credentials{
				AccessKeyID:     accessKeyID,
				SecretAccessKey: secretAccessKey,
			}, nil
		})),
	)

	if err != nil {
		return aws.Config{}, err
	}

	return cfg, nil
}
// NewConfigWithRegion 建立並回傳一個指定區域的 AWS 組態設定
func NewConfigWithRegion(ctx context.Context, region string) (aws.Config, error) {
	accessKeyID := viper.GetString("aws.accessKeyID")
	secretAccessKey := viper.GetString("aws.secretAccessKey")

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(region), // <-- 使用傳入的 region
		config.WithCredentialsProvider(aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
			return aws.Credentials{
				AccessKeyID:     accessKeyID,
				SecretAccessKey: secretAccessKey,
			}, nil
		})),
	)

	if err != nil {
		return aws.Config{}, err
	}

	return cfg, nil
}
