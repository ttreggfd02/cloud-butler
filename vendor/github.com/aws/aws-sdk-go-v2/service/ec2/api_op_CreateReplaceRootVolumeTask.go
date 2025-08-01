// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Replaces the EBS-backed root volume for a running instance with a new volume
// that is restored to the original root volume's launch state, that is restored to
// a specific snapshot taken from the original root volume, or that is restored
// from an AMI that has the same key characteristics as that of the instance.
//
// For more information, see [Replace a root volume] in the Amazon EC2 User Guide.
//
// [Replace a root volume]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/replace-root.html
func (c *Client) CreateReplaceRootVolumeTask(ctx context.Context, params *CreateReplaceRootVolumeTaskInput, optFns ...func(*Options)) (*CreateReplaceRootVolumeTaskOutput, error) {
	if params == nil {
		params = &CreateReplaceRootVolumeTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateReplaceRootVolumeTask", params, optFns, c.addOperationCreateReplaceRootVolumeTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateReplaceRootVolumeTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateReplaceRootVolumeTaskInput struct {

	// The ID of the instance for which to replace the root volume.
	//
	// This member is required.
	InstanceId *string

	// Unique, case-sensitive identifier you provide to ensure the idempotency of the
	// request. If you do not specify a client token, a randomly generated token is
	// used for the request to ensure idempotency. For more information, see [Ensuring idempotency].
	//
	// [Ensuring idempotency]: https://docs.aws.amazon.com/ec2/latest/devguide/ec2-api-idempotency.html
	ClientToken *string

	// Indicates whether to automatically delete the original root volume after the
	// root volume replacement task completes. To delete the original root volume,
	// specify true . If you choose to keep the original root volume after the
	// replacement task completes, you must manually delete it when you no longer need
	// it.
	DeleteReplacedRootVolume *bool

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The ID of the AMI to use to restore the root volume. The specified AMI must
	// have the same product code, billing information, architecture type, and
	// virtualization type as that of the instance.
	//
	// If you want to restore the replacement volume from a specific snapshot, or if
	// you want to restore it to its launch state, omit this parameter.
	ImageId *string

	// The ID of the snapshot from which to restore the replacement root volume. The
	// specified snapshot must be a snapshot that you previously created from the
	// original root volume.
	//
	// If you want to restore the replacement root volume to the initial launch state,
	// or if you want to restore the replacement root volume from an AMI, omit this
	// parameter.
	SnapshotId *string

	// The tags to apply to the root volume replacement task.
	TagSpecifications []types.TagSpecification

	// Specifies the Amazon EBS Provisioned Rate for Volume Initialization (volume
	// initialization rate), in MiB/s, at which to download the snapshot blocks from
	// Amazon S3 to the replacement root volume. This is also known as volume
	// initialization. Specifying a volume initialization rate ensures that the volume
	// is initialized at a predictable and consistent rate after creation.
	//
	// Omit this parameter if:
	//
	//   - You want to create the volume using fast snapshot restore. You must specify
	//   a snapshot that is enabled for fast snapshot restore. In this case, the volume
	//   is fully initialized at creation.
	//
	// If you specify a snapshot that is enabled for fast snapshot restore and a
	//   volume initialization rate, the volume will be initialized at the specified rate
	//   instead of fast snapshot restore.
	//
	//   - You want to create a volume that is initialized at the default rate.
	//
	// For more information, see [Initialize Amazon EBS volumes] in the Amazon EC2 User Guide.
	//
	// Valid range: 100 - 300 MiB/s
	//
	// [Initialize Amazon EBS volumes]: https://docs.aws.amazon.com/ebs/latest/userguide/initalize-volume.html
	VolumeInitializationRate *int64

	noSmithyDocumentSerde
}

type CreateReplaceRootVolumeTaskOutput struct {

	// Information about the root volume replacement task.
	ReplaceRootVolumeTask *types.ReplaceRootVolumeTask

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateReplaceRootVolumeTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateReplaceRootVolumeTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateReplaceRootVolumeTask{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateReplaceRootVolumeTask"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addSpanRetryLoop(stack, options); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addCredentialSource(stack, options); err != nil {
		return err
	}
	if err = addIdempotencyToken_opCreateReplaceRootVolumeTaskMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateReplaceRootVolumeTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateReplaceRootVolumeTask(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addInterceptBeforeRetryLoop(stack, options); err != nil {
		return err
	}
	if err = addInterceptAttempt(stack, options); err != nil {
		return err
	}
	if err = addInterceptExecution(stack, options); err != nil {
		return err
	}
	if err = addInterceptBeforeSerialization(stack, options); err != nil {
		return err
	}
	if err = addInterceptAfterSerialization(stack, options); err != nil {
		return err
	}
	if err = addInterceptBeforeSigning(stack, options); err != nil {
		return err
	}
	if err = addInterceptAfterSigning(stack, options); err != nil {
		return err
	}
	if err = addInterceptTransmit(stack, options); err != nil {
		return err
	}
	if err = addInterceptBeforeDeserialization(stack, options); err != nil {
		return err
	}
	if err = addInterceptAfterDeserialization(stack, options); err != nil {
		return err
	}
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

type idempotencyToken_initializeOpCreateReplaceRootVolumeTask struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateReplaceRootVolumeTask) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateReplaceRootVolumeTask) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateReplaceRootVolumeTaskInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateReplaceRootVolumeTaskInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateReplaceRootVolumeTaskMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateReplaceRootVolumeTask{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateReplaceRootVolumeTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateReplaceRootVolumeTask",
	}
}
