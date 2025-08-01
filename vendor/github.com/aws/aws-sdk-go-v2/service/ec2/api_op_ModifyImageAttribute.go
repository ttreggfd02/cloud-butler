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

// Modifies the specified attribute of the specified AMI. You can specify only one
// attribute at a time.
//
// To specify the attribute, you can use the Attribute parameter, or one of the
// following parameters: Description , ImdsSupport , or LaunchPermission .
//
// Images with an Amazon Web Services Marketplace product code cannot be made
// public.
//
// To enable the SriovNetSupport enhanced networking attribute of an image, enable
// SriovNetSupport on an instance and create an AMI from the instance.
func (c *Client) ModifyImageAttribute(ctx context.Context, params *ModifyImageAttributeInput, optFns ...func(*Options)) (*ModifyImageAttributeOutput, error) {
	if params == nil {
		params = &ModifyImageAttributeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyImageAttribute", params, optFns, c.addOperationModifyImageAttributeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyImageAttributeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for ModifyImageAttribute.
type ModifyImageAttributeInput struct {

	// The ID of the AMI.
	//
	// This member is required.
	ImageId *string

	// The name of the attribute to modify.
	//
	// Valid values: description | imdsSupport | launchPermission
	Attribute *string

	// A new description for the AMI.
	Description *types.AttributeValue

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// Set to v2.0 to indicate that IMDSv2 is specified in the AMI. Instances launched
	// from this AMI will have HttpTokens automatically set to required so that, by
	// default, the instance requires that IMDSv2 is used when requesting instance
	// metadata. In addition, HttpPutResponseHopLimit is set to 2 . For more
	// information, see [Configure the AMI]in the Amazon EC2 User Guide.
	//
	// Do not use this parameter unless your AMI software supports IMDSv2. After you
	// set the value to v2.0 , you can't undo it. The only way to “reset” your AMI is
	// to create a new AMI from the underlying snapshot.
	//
	// [Configure the AMI]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/configuring-IMDS-new-instances.html#configure-IMDS-new-instances-ami-configuration
	ImdsSupport *types.AttributeValue

	// A new launch permission for the AMI.
	LaunchPermission *types.LaunchPermissionModifications

	// The operation type. This parameter can be used only when the Attribute
	// parameter is launchPermission .
	OperationType types.OperationType

	// The Amazon Resource Name (ARN) of an organization. This parameter can be used
	// only when the Attribute parameter is launchPermission .
	OrganizationArns []string

	// The Amazon Resource Name (ARN) of an organizational unit (OU). This parameter
	// can be used only when the Attribute parameter is launchPermission .
	OrganizationalUnitArns []string

	// Not supported.
	ProductCodes []string

	// The user groups. This parameter can be used only when the Attribute parameter
	// is launchPermission .
	UserGroups []string

	// The Amazon Web Services account IDs. This parameter can be used only when the
	// Attribute parameter is launchPermission .
	UserIds []string

	// The value of the attribute being modified. This parameter can be used only when
	// the Attribute parameter is description or imdsSupport .
	Value *string

	noSmithyDocumentSerde
}

type ModifyImageAttributeOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyImageAttributeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyImageAttribute{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyImageAttribute{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ModifyImageAttribute"); err != nil {
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
	if err = addOpModifyImageAttributeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyImageAttribute(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyImageAttribute(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ModifyImageAttribute",
	}
}
