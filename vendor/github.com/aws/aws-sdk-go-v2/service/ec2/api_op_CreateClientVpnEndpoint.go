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

// Creates a Client VPN endpoint. A Client VPN endpoint is the resource you create
// and configure to enable and manage client VPN sessions. It is the destination
// endpoint at which all client VPN sessions are terminated.
func (c *Client) CreateClientVpnEndpoint(ctx context.Context, params *CreateClientVpnEndpointInput, optFns ...func(*Options)) (*CreateClientVpnEndpointOutput, error) {
	if params == nil {
		params = &CreateClientVpnEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateClientVpnEndpoint", params, optFns, c.addOperationCreateClientVpnEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateClientVpnEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateClientVpnEndpointInput struct {

	// Information about the authentication method to be used to authenticate clients.
	//
	// This member is required.
	AuthenticationOptions []types.ClientVpnAuthenticationRequest

	// The IPv4 address range, in CIDR notation, from which to assign client IP
	// addresses. The address range cannot overlap with the local CIDR of the VPC in
	// which the associated subnet is located, or the routes that you add manually. The
	// address range cannot be changed after the Client VPN endpoint has been created.
	// Client CIDR range must have a size of at least /22 and must not be greater than
	// /12.
	//
	// This member is required.
	ClientCidrBlock *string

	// Information about the client connection logging options.
	//
	// If you enable client connection logging, data about client connections is sent
	// to a Cloudwatch Logs log stream. The following information is logged:
	//
	//   - Client connection requests
	//
	//   - Client connection results (successful and unsuccessful)
	//
	//   - Reasons for unsuccessful client connection requests
	//
	//   - Client connection termination time
	//
	// This member is required.
	ConnectionLogOptions *types.ConnectionLogOptions

	// The ARN of the server certificate. For more information, see the [Certificate Manager User Guide].
	//
	// [Certificate Manager User Guide]: https://docs.aws.amazon.com/acm/latest/userguide/
	//
	// This member is required.
	ServerCertificateArn *string

	// The options for managing connection authorization for new client connections.
	ClientConnectOptions *types.ClientConnectOptions

	// Options for enabling a customizable text banner that will be displayed on
	// Amazon Web Services provided clients when a VPN session is established.
	ClientLoginBannerOptions *types.ClientLoginBannerOptions

	// Client route enforcement is a feature of the Client VPN service that helps
	// enforce administrator defined routes on devices connected through the VPN. T his
	// feature helps improve your security posture by ensuring that network traffic
	// originating from a connected client is not inadvertently sent outside the VPN
	// tunnel.
	//
	// Client route enforcement works by monitoring the route table of a connected
	// device for routing policy changes to the VPN connection. If the feature detects
	// any VPN routing policy modifications, it will automatically force an update to
	// the route table, reverting it back to the expected route configurations.
	ClientRouteEnforcementOptions *types.ClientRouteEnforcementOptions

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. For more information, see [Ensuring idempotency].
	//
	// [Ensuring idempotency]: https://docs.aws.amazon.com/ec2/latest/devguide/ec2-api-idempotency.html
	ClientToken *string

	// A brief description of the Client VPN endpoint.
	Description *string

	// Indicates whether the client VPN session is disconnected after the maximum
	// timeout specified in SessionTimeoutHours is reached. If true , users are
	// prompted to reconnect client VPN. If false , client VPN attempts to reconnect
	// automatically. The default value is true .
	DisconnectOnSessionTimeout *bool

	// Information about the DNS servers to be used for DNS resolution. A Client VPN
	// endpoint can have up to two DNS servers. If no DNS server is specified, the DNS
	// address configured on the device is used for the DNS server.
	DnsServers []string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The IDs of one or more security groups to apply to the target network. You must
	// also specify the ID of the VPC that contains the security groups.
	SecurityGroupIds []string

	// Specify whether to enable the self-service portal for the Client VPN endpoint.
	//
	// Default Value: enabled
	SelfServicePortal types.SelfServicePortal

	// The maximum VPN session duration time in hours.
	//
	// Valid values: 8 | 10 | 12 | 24
	//
	// Default value: 24
	SessionTimeoutHours *int32

	// Indicates whether split-tunnel is enabled on the Client VPN endpoint.
	//
	// By default, split-tunnel on a VPN endpoint is disabled.
	//
	// For information about split-tunnel VPN endpoints, see [Split-tunnel Client VPN endpoint] in the Client VPN
	// Administrator Guide.
	//
	// [Split-tunnel Client VPN endpoint]: https://docs.aws.amazon.com/vpn/latest/clientvpn-admin/split-tunnel-vpn.html
	SplitTunnel *bool

	// The tags to apply to the Client VPN endpoint during creation.
	TagSpecifications []types.TagSpecification

	// The transport protocol to be used by the VPN session.
	//
	// Default value: udp
	TransportProtocol types.TransportProtocol

	// The ID of the VPC to associate with the Client VPN endpoint. If no security
	// group IDs are specified in the request, the default security group for the VPC
	// is applied.
	VpcId *string

	// The port number to assign to the Client VPN endpoint for TCP and UDP traffic.
	//
	// Valid Values: 443 | 1194
	//
	// Default Value: 443
	VpnPort *int32

	noSmithyDocumentSerde
}

type CreateClientVpnEndpointOutput struct {

	// The ID of the Client VPN endpoint.
	ClientVpnEndpointId *string

	// The DNS name to be used by clients when establishing their VPN session.
	DnsName *string

	// The current state of the Client VPN endpoint.
	Status *types.ClientVpnEndpointStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateClientVpnEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateClientVpnEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateClientVpnEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateClientVpnEndpoint"); err != nil {
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
	if err = addIdempotencyToken_opCreateClientVpnEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateClientVpnEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateClientVpnEndpoint(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateClientVpnEndpoint struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateClientVpnEndpoint) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateClientVpnEndpoint) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateClientVpnEndpointInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateClientVpnEndpointInput ")
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
func addIdempotencyToken_opCreateClientVpnEndpointMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateClientVpnEndpoint{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateClientVpnEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateClientVpnEndpoint",
	}
}
