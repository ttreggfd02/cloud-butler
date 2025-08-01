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

// Creates a target for your Traffic Mirror session.
//
// A Traffic Mirror target is the destination for mirrored traffic. The Traffic
// Mirror source and the Traffic Mirror target (monitoring appliances) can be in
// the same VPC, or in different VPCs connected via VPC peering or a transit
// gateway.
//
// A Traffic Mirror target can be a network interface, a Network Load Balancer, or
// a Gateway Load Balancer endpoint.
//
// To use the target in a Traffic Mirror session, use [CreateTrafficMirrorSession].
//
// [CreateTrafficMirrorSession]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateTrafficMirrorSession.htm
func (c *Client) CreateTrafficMirrorTarget(ctx context.Context, params *CreateTrafficMirrorTargetInput, optFns ...func(*Options)) (*CreateTrafficMirrorTargetOutput, error) {
	if params == nil {
		params = &CreateTrafficMirrorTargetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTrafficMirrorTarget", params, optFns, c.addOperationCreateTrafficMirrorTargetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTrafficMirrorTargetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTrafficMirrorTargetInput struct {

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. For more information, see [How to ensure idempotency].
	//
	// [How to ensure idempotency]: https://docs.aws.amazon.com/ec2/latest/devguide/ec2-api-idempotency.html
	ClientToken *string

	// The description of the Traffic Mirror target.
	Description *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The ID of the Gateway Load Balancer endpoint.
	GatewayLoadBalancerEndpointId *string

	// The network interface ID that is associated with the target.
	NetworkInterfaceId *string

	// The Amazon Resource Name (ARN) of the Network Load Balancer that is associated
	// with the target.
	NetworkLoadBalancerArn *string

	// The tags to assign to the Traffic Mirror target.
	TagSpecifications []types.TagSpecification

	noSmithyDocumentSerde
}

type CreateTrafficMirrorTargetOutput struct {

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. For more information, see [How to ensure idempotency].
	//
	// [How to ensure idempotency]: https://docs.aws.amazon.com/ec2/latest/devguide/ec2-api-idempotency.html
	ClientToken *string

	// Information about the Traffic Mirror target.
	TrafficMirrorTarget *types.TrafficMirrorTarget

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateTrafficMirrorTargetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateTrafficMirrorTarget{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateTrafficMirrorTarget{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateTrafficMirrorTarget"); err != nil {
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
	if err = addIdempotencyToken_opCreateTrafficMirrorTargetMiddleware(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTrafficMirrorTarget(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateTrafficMirrorTarget struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateTrafficMirrorTarget) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateTrafficMirrorTarget) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateTrafficMirrorTargetInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateTrafficMirrorTargetInput ")
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
func addIdempotencyToken_opCreateTrafficMirrorTargetMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateTrafficMirrorTarget{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateTrafficMirrorTarget(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateTrafficMirrorTarget",
	}
}
