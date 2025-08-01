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

// Modifies the specified network interface attribute. You can specify only one
// attribute at a time. You can use this action to attach and detach security
// groups from an existing EC2 instance.
func (c *Client) ModifyNetworkInterfaceAttribute(ctx context.Context, params *ModifyNetworkInterfaceAttributeInput, optFns ...func(*Options)) (*ModifyNetworkInterfaceAttributeOutput, error) {
	if params == nil {
		params = &ModifyNetworkInterfaceAttributeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyNetworkInterfaceAttribute", params, optFns, c.addOperationModifyNetworkInterfaceAttributeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyNetworkInterfaceAttributeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for ModifyNetworkInterfaceAttribute.
type ModifyNetworkInterfaceAttributeInput struct {

	// The ID of the network interface.
	//
	// This member is required.
	NetworkInterfaceId *string

	// Indicates whether to assign a public IPv4 address to a network interface. This
	// option can be enabled for any network interface but will only apply to the
	// primary network interface (eth0).
	AssociatePublicIpAddress *bool

	// A list of subnet IDs to associate with the network interface.
	AssociatedSubnetIds []string

	// Information about the interface attachment. If modifying the delete on
	// termination attribute, you must specify the ID of the interface attachment.
	Attachment *types.NetworkInterfaceAttachmentChanges

	// A connection tracking specification.
	ConnectionTrackingSpecification *types.ConnectionTrackingSpecificationRequest

	// A description for the network interface.
	Description *types.AttributeValue

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// Updates the ENA Express configuration for the network interface that’s attached
	// to the instance.
	EnaSrdSpecification *types.EnaSrdSpecification

	// If you’re modifying a network interface in a dual-stack or IPv6-only subnet,
	// you have the option to assign a primary IPv6 IP address. A primary IPv6 address
	// is an IPv6 GUA address associated with an ENI that you have enabled to use a
	// primary IPv6 address. Use this option if the instance that this ENI will be
	// attached to relies on its IPv6 address not changing. Amazon Web Services will
	// automatically assign an IPv6 address associated with the ENI attached to your
	// instance to be the primary IPv6 address. Once you enable an IPv6 GUA address to
	// be a primary IPv6, you cannot disable it. When you enable an IPv6 GUA address to
	// be a primary IPv6, the first IPv6 GUA will be made the primary IPv6 address
	// until the instance is terminated or the network interface is detached. If you
	// have multiple IPv6 addresses associated with an ENI attached to your instance
	// and you enable a primary IPv6 address, the first IPv6 GUA address associated
	// with the ENI becomes the primary IPv6 address.
	EnablePrimaryIpv6 *bool

	// Changes the security groups for the network interface. The new set of groups
	// you specify replaces the current set. You must specify at least one group, even
	// if it's just the default security group in the VPC. You must specify the ID of
	// the security group, not the name.
	Groups []string

	// Enable or disable source/destination checks, which ensure that the instance is
	// either the source or the destination of any traffic that it receives. If the
	// value is true , source/destination checks are enabled; otherwise, they are
	// disabled. The default value is true . You must disable source/destination checks
	// if the instance runs services such as network address translation, routing, or
	// firewalls.
	SourceDestCheck *types.AttributeBooleanValue

	noSmithyDocumentSerde
}

type ModifyNetworkInterfaceAttributeOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyNetworkInterfaceAttributeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyNetworkInterfaceAttribute{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyNetworkInterfaceAttribute{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ModifyNetworkInterfaceAttribute"); err != nil {
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
	if err = addOpModifyNetworkInterfaceAttributeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyNetworkInterfaceAttribute(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyNetworkInterfaceAttribute(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ModifyNetworkInterfaceAttribute",
	}
}
