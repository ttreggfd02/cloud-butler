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

// Describes an Elastic IP address transfer. For more information, see [Transfer Elastic IP addresses] in the
// Amazon VPC User Guide.
//
// When you transfer an Elastic IP address, there is a two-step handshake between
// the source and transfer Amazon Web Services accounts. When the source account
// starts the transfer, the transfer account has seven days to accept the Elastic
// IP address transfer. During those seven days, the source account can view the
// pending transfer by using this action. After seven days, the transfer expires
// and ownership of the Elastic IP address returns to the source account. Accepted
// transfers are visible to the source account for 14 days after the transfers have
// been accepted.
//
// [Transfer Elastic IP addresses]: https://docs.aws.amazon.com/vpc/latest/userguide/vpc-eips.html#transfer-EIPs-intro
func (c *Client) DescribeAddressTransfers(ctx context.Context, params *DescribeAddressTransfersInput, optFns ...func(*Options)) (*DescribeAddressTransfersOutput, error) {
	if params == nil {
		params = &DescribeAddressTransfersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAddressTransfers", params, optFns, c.addOperationDescribeAddressTransfersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAddressTransfersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAddressTransfersInput struct {

	// The allocation IDs of Elastic IP addresses.
	AllocationIds []string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The maximum number of address transfers to return in one page of results.
	MaxResults *int32

	// Specify the pagination token from a previous request to retrieve the next page
	// of results.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeAddressTransfersOutput struct {

	// The Elastic IP address transfer.
	AddressTransfers []types.AddressTransfer

	// Specify the pagination token from a previous request to retrieve the next page
	// of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAddressTransfersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeAddressTransfers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeAddressTransfers{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeAddressTransfers"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAddressTransfers(options.Region), middleware.Before); err != nil {
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

// DescribeAddressTransfersPaginatorOptions is the paginator options for
// DescribeAddressTransfers
type DescribeAddressTransfersPaginatorOptions struct {
	// The maximum number of address transfers to return in one page of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeAddressTransfersPaginator is a paginator for DescribeAddressTransfers
type DescribeAddressTransfersPaginator struct {
	options   DescribeAddressTransfersPaginatorOptions
	client    DescribeAddressTransfersAPIClient
	params    *DescribeAddressTransfersInput
	nextToken *string
	firstPage bool
}

// NewDescribeAddressTransfersPaginator returns a new
// DescribeAddressTransfersPaginator
func NewDescribeAddressTransfersPaginator(client DescribeAddressTransfersAPIClient, params *DescribeAddressTransfersInput, optFns ...func(*DescribeAddressTransfersPaginatorOptions)) *DescribeAddressTransfersPaginator {
	if params == nil {
		params = &DescribeAddressTransfersInput{}
	}

	options := DescribeAddressTransfersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeAddressTransfersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeAddressTransfersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeAddressTransfers page.
func (p *DescribeAddressTransfersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeAddressTransfersOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.DescribeAddressTransfers(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// DescribeAddressTransfersAPIClient is a client that implements the
// DescribeAddressTransfers operation.
type DescribeAddressTransfersAPIClient interface {
	DescribeAddressTransfers(context.Context, *DescribeAddressTransfersInput, ...func(*Options)) (*DescribeAddressTransfersOutput, error)
}

var _ DescribeAddressTransfersAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeAddressTransfers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeAddressTransfers",
	}
}
