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

// Creates a NAT gateway in the specified subnet. This action creates a network
// interface in the specified subnet with a private IP address from the IP address
// range of the subnet. You can create either a public NAT gateway or a private NAT
// gateway.
//
// With a public NAT gateway, internet-bound traffic from a private subnet can be
// routed to the NAT gateway, so that instances in a private subnet can connect to
// the internet.
//
// With a private NAT gateway, private communication is routed across VPCs and
// on-premises networks through a transit gateway or virtual private gateway.
// Common use cases include running large workloads behind a small pool of
// allowlisted IPv4 addresses, preserving private IPv4 addresses, and communicating
// between overlapping networks.
//
// For more information, see [NAT gateways] in the Amazon VPC User Guide.
//
// When you create a public NAT gateway and assign it an EIP or secondary EIPs,
// the network border group of the EIPs must match the network border group of the
// Availability Zone (AZ) that the public NAT gateway is in. If it's not the same,
// the NAT gateway will fail to launch. You can see the network border group for
// the subnet's AZ by viewing the details of the subnet. Similarly, you can view
// the network border group of an EIP by viewing the details of the EIP address.
// For more information about network border groups and EIPs, see [Allocate an Elastic IP address]in the Amazon
// VPC User Guide.
//
// [NAT gateways]: https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html
// [Allocate an Elastic IP address]: https://docs.aws.amazon.com/vpc/latest/userguide/WorkWithEIPs.html
func (c *Client) CreateNatGateway(ctx context.Context, params *CreateNatGatewayInput, optFns ...func(*Options)) (*CreateNatGatewayOutput, error) {
	if params == nil {
		params = &CreateNatGatewayInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateNatGateway", params, optFns, c.addOperationCreateNatGatewayMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateNatGatewayOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateNatGatewayInput struct {

	// The ID of the subnet in which to create the NAT gateway.
	//
	// This member is required.
	SubnetId *string

	// [Public NAT gateways only] The allocation ID of an Elastic IP address to
	// associate with the NAT gateway. You cannot specify an Elastic IP address with a
	// private NAT gateway. If the Elastic IP address is associated with another
	// resource, you must first disassociate it.
	AllocationId *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. For more information, see [Ensuring idempotency].
	//
	// Constraint: Maximum 64 ASCII characters.
	//
	// [Ensuring idempotency]: https://docs.aws.amazon.com/ec2/latest/devguide/ec2-api-idempotency.html
	ClientToken *string

	// Indicates whether the NAT gateway supports public or private connectivity. The
	// default is public connectivity.
	ConnectivityType types.ConnectivityType

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The private IPv4 address to assign to the NAT gateway. If you don't provide an
	// address, a private IPv4 address will be automatically assigned.
	PrivateIpAddress *string

	// Secondary EIP allocation IDs. For more information, see [Create a NAT gateway] in the Amazon VPC User
	// Guide.
	//
	// [Create a NAT gateway]: https://docs.aws.amazon.com/vpc/latest/userguide/nat-gateway-working-with.html
	SecondaryAllocationIds []string

	// [Private NAT gateway only] The number of secondary private IPv4 addresses you
	// want to assign to the NAT gateway. For more information about secondary
	// addresses, see [Create a NAT gateway]in the Amazon VPC User Guide.
	//
	// [Create a NAT gateway]: https://docs.aws.amazon.com/vpc/latest/userguide/nat-gateway-working-with.html
	SecondaryPrivateIpAddressCount *int32

	// Secondary private IPv4 addresses. For more information about secondary
	// addresses, see [Create a NAT gateway]in the Amazon VPC User Guide.
	//
	// [Create a NAT gateway]: https://docs.aws.amazon.com/vpc/latest/userguide/nat-gateway-working-with.html
	SecondaryPrivateIpAddresses []string

	// The tags to assign to the NAT gateway.
	TagSpecifications []types.TagSpecification

	noSmithyDocumentSerde
}

type CreateNatGatewayOutput struct {

	// Unique, case-sensitive identifier to ensure the idempotency of the request.
	// Only returned if a client token was provided in the request.
	ClientToken *string

	// Information about the NAT gateway.
	NatGateway *types.NatGateway

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateNatGatewayMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateNatGateway{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateNatGateway{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateNatGateway"); err != nil {
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
	if err = addIdempotencyToken_opCreateNatGatewayMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateNatGatewayValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateNatGateway(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateNatGateway struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateNatGateway) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateNatGateway) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateNatGatewayInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateNatGatewayInput ")
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
func addIdempotencyToken_opCreateNatGatewayMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateNatGateway{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateNatGateway(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateNatGateway",
	}
}
