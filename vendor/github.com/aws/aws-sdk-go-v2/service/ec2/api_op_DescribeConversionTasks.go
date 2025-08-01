// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"time"
)

// Describes the specified conversion tasks or all your conversion tasks. For more
// information, see the [VM Import/Export User Guide].
//
// For information about the import manifest referenced by this API action, see [VM Import Manifest].
//
// [VM Import Manifest]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/manifest.html
// [VM Import/Export User Guide]: https://docs.aws.amazon.com/vm-import/latest/userguide/
func (c *Client) DescribeConversionTasks(ctx context.Context, params *DescribeConversionTasksInput, optFns ...func(*Options)) (*DescribeConversionTasksOutput, error) {
	if params == nil {
		params = &DescribeConversionTasksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeConversionTasks", params, optFns, c.addOperationDescribeConversionTasksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeConversionTasksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeConversionTasksInput struct {

	// The conversion task IDs.
	ConversionTaskIds []string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	noSmithyDocumentSerde
}

type DescribeConversionTasksOutput struct {

	// Information about the conversion tasks.
	ConversionTasks []types.ConversionTask

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeConversionTasksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeConversionTasks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeConversionTasks{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeConversionTasks"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeConversionTasks(options.Region), middleware.Before); err != nil {
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

// ConversionTaskCancelledWaiterOptions are waiter options for
// ConversionTaskCancelledWaiter
type ConversionTaskCancelledWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	//
	// Passing options here is functionally equivalent to passing values to this
	// config's ClientOptions field that extend the inner client's APIOptions directly.
	APIOptions []func(*middleware.Stack) error

	// Functional options to be passed to all operations invoked by this client.
	//
	// Function values that modify the inner APIOptions are applied after the waiter
	// config's own APIOptions modifiers.
	ClientOptions []func(*Options)

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// ConversionTaskCancelledWaiter will use default minimum delay of 15 seconds. Note
	// that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, ConversionTaskCancelledWaiter will use default max delay of 120
	// seconds. Note that MaxDelay must resolve to value greater than or equal to the
	// MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state.
	//
	// By default service-modeled logic will populate this option. This option can
	// thus be used to define a custom waiter state with fall-back to service-modeled
	// waiter state mutators.The function returns an error in case of a failure state.
	// In case of retry state, this function returns a bool value of true and nil
	// error, while in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeConversionTasksInput, *DescribeConversionTasksOutput, error) (bool, error)
}

// ConversionTaskCancelledWaiter defines the waiters for ConversionTaskCancelled
type ConversionTaskCancelledWaiter struct {
	client DescribeConversionTasksAPIClient

	options ConversionTaskCancelledWaiterOptions
}

// NewConversionTaskCancelledWaiter constructs a ConversionTaskCancelledWaiter.
func NewConversionTaskCancelledWaiter(client DescribeConversionTasksAPIClient, optFns ...func(*ConversionTaskCancelledWaiterOptions)) *ConversionTaskCancelledWaiter {
	options := ConversionTaskCancelledWaiterOptions{}
	options.MinDelay = 15 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = conversionTaskCancelledStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &ConversionTaskCancelledWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for ConversionTaskCancelled waiter. The
// maxWaitDur is the maximum wait duration the waiter will wait. The maxWaitDur is
// required and must be greater than zero.
func (w *ConversionTaskCancelledWaiter) Wait(ctx context.Context, params *DescribeConversionTasksInput, maxWaitDur time.Duration, optFns ...func(*ConversionTaskCancelledWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for ConversionTaskCancelled waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *ConversionTaskCancelledWaiter) WaitForOutput(ctx context.Context, params *DescribeConversionTasksInput, maxWaitDur time.Duration, optFns ...func(*ConversionTaskCancelledWaiterOptions)) (*DescribeConversionTasksOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.DescribeConversionTasks(ctx, params, func(o *Options) {
			baseOpts := []func(*Options){
				addIsWaiterUserAgent,
			}
			o.APIOptions = append(o.APIOptions, apiOptions...)
			for _, opt := range baseOpts {
				opt(o)
			}
			for _, opt := range options.ClientOptions {
				opt(o)
			}
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for ConversionTaskCancelled waiter")
}

func conversionTaskCancelledStateRetryable(ctx context.Context, input *DescribeConversionTasksInput, output *DescribeConversionTasksOutput, err error) (bool, error) {

	if err == nil {
		v1 := output.ConversionTasks
		var v2 []types.ConversionTaskState
		for _, v := range v1 {
			v3 := v.State
			v2 = append(v2, v3)
		}
		expectedValue := "cancelled"
		match := len(v2) > 0
		for _, v := range v2 {
			if string(v) != expectedValue {
				match = false
				break
			}
		}

		if match {
			return false, nil
		}
	}

	if err != nil {
		return false, err
	}
	return true, nil
}

// ConversionTaskCompletedWaiterOptions are waiter options for
// ConversionTaskCompletedWaiter
type ConversionTaskCompletedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	//
	// Passing options here is functionally equivalent to passing values to this
	// config's ClientOptions field that extend the inner client's APIOptions directly.
	APIOptions []func(*middleware.Stack) error

	// Functional options to be passed to all operations invoked by this client.
	//
	// Function values that modify the inner APIOptions are applied after the waiter
	// config's own APIOptions modifiers.
	ClientOptions []func(*Options)

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// ConversionTaskCompletedWaiter will use default minimum delay of 15 seconds. Note
	// that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, ConversionTaskCompletedWaiter will use default max delay of 120
	// seconds. Note that MaxDelay must resolve to value greater than or equal to the
	// MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state.
	//
	// By default service-modeled logic will populate this option. This option can
	// thus be used to define a custom waiter state with fall-back to service-modeled
	// waiter state mutators.The function returns an error in case of a failure state.
	// In case of retry state, this function returns a bool value of true and nil
	// error, while in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeConversionTasksInput, *DescribeConversionTasksOutput, error) (bool, error)
}

// ConversionTaskCompletedWaiter defines the waiters for ConversionTaskCompleted
type ConversionTaskCompletedWaiter struct {
	client DescribeConversionTasksAPIClient

	options ConversionTaskCompletedWaiterOptions
}

// NewConversionTaskCompletedWaiter constructs a ConversionTaskCompletedWaiter.
func NewConversionTaskCompletedWaiter(client DescribeConversionTasksAPIClient, optFns ...func(*ConversionTaskCompletedWaiterOptions)) *ConversionTaskCompletedWaiter {
	options := ConversionTaskCompletedWaiterOptions{}
	options.MinDelay = 15 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = conversionTaskCompletedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &ConversionTaskCompletedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for ConversionTaskCompleted waiter. The
// maxWaitDur is the maximum wait duration the waiter will wait. The maxWaitDur is
// required and must be greater than zero.
func (w *ConversionTaskCompletedWaiter) Wait(ctx context.Context, params *DescribeConversionTasksInput, maxWaitDur time.Duration, optFns ...func(*ConversionTaskCompletedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for ConversionTaskCompleted waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *ConversionTaskCompletedWaiter) WaitForOutput(ctx context.Context, params *DescribeConversionTasksInput, maxWaitDur time.Duration, optFns ...func(*ConversionTaskCompletedWaiterOptions)) (*DescribeConversionTasksOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.DescribeConversionTasks(ctx, params, func(o *Options) {
			baseOpts := []func(*Options){
				addIsWaiterUserAgent,
			}
			o.APIOptions = append(o.APIOptions, apiOptions...)
			for _, opt := range baseOpts {
				opt(o)
			}
			for _, opt := range options.ClientOptions {
				opt(o)
			}
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for ConversionTaskCompleted waiter")
}

func conversionTaskCompletedStateRetryable(ctx context.Context, input *DescribeConversionTasksInput, output *DescribeConversionTasksOutput, err error) (bool, error) {

	if err == nil {
		v1 := output.ConversionTasks
		var v2 []types.ConversionTaskState
		for _, v := range v1 {
			v3 := v.State
			v2 = append(v2, v3)
		}
		expectedValue := "completed"
		match := len(v2) > 0
		for _, v := range v2 {
			if string(v) != expectedValue {
				match = false
				break
			}
		}

		if match {
			return false, nil
		}
	}

	if err == nil {
		v1 := output.ConversionTasks
		var v2 []types.ConversionTaskState
		for _, v := range v1 {
			v3 := v.State
			v2 = append(v2, v3)
		}
		expectedValue := "cancelled"
		var match bool
		for _, v := range v2 {
			if string(v) == expectedValue {
				match = true
				break
			}
		}

		if match {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err == nil {
		v1 := output.ConversionTasks
		var v2 []types.ConversionTaskState
		for _, v := range v1 {
			v3 := v.State
			v2 = append(v2, v3)
		}
		expectedValue := "cancelling"
		var match bool
		for _, v := range v2 {
			if string(v) == expectedValue {
				match = true
				break
			}
		}

		if match {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err != nil {
		return false, err
	}
	return true, nil
}

// ConversionTaskDeletedWaiterOptions are waiter options for
// ConversionTaskDeletedWaiter
type ConversionTaskDeletedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	//
	// Passing options here is functionally equivalent to passing values to this
	// config's ClientOptions field that extend the inner client's APIOptions directly.
	APIOptions []func(*middleware.Stack) error

	// Functional options to be passed to all operations invoked by this client.
	//
	// Function values that modify the inner APIOptions are applied after the waiter
	// config's own APIOptions modifiers.
	ClientOptions []func(*Options)

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// ConversionTaskDeletedWaiter will use default minimum delay of 15 seconds. Note
	// that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, ConversionTaskDeletedWaiter will use default max delay of 120
	// seconds. Note that MaxDelay must resolve to value greater than or equal to the
	// MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state.
	//
	// By default service-modeled logic will populate this option. This option can
	// thus be used to define a custom waiter state with fall-back to service-modeled
	// waiter state mutators.The function returns an error in case of a failure state.
	// In case of retry state, this function returns a bool value of true and nil
	// error, while in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeConversionTasksInput, *DescribeConversionTasksOutput, error) (bool, error)
}

// ConversionTaskDeletedWaiter defines the waiters for ConversionTaskDeleted
type ConversionTaskDeletedWaiter struct {
	client DescribeConversionTasksAPIClient

	options ConversionTaskDeletedWaiterOptions
}

// NewConversionTaskDeletedWaiter constructs a ConversionTaskDeletedWaiter.
func NewConversionTaskDeletedWaiter(client DescribeConversionTasksAPIClient, optFns ...func(*ConversionTaskDeletedWaiterOptions)) *ConversionTaskDeletedWaiter {
	options := ConversionTaskDeletedWaiterOptions{}
	options.MinDelay = 15 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = conversionTaskDeletedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &ConversionTaskDeletedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for ConversionTaskDeleted waiter. The maxWaitDur
// is the maximum wait duration the waiter will wait. The maxWaitDur is required
// and must be greater than zero.
func (w *ConversionTaskDeletedWaiter) Wait(ctx context.Context, params *DescribeConversionTasksInput, maxWaitDur time.Duration, optFns ...func(*ConversionTaskDeletedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for ConversionTaskDeleted waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *ConversionTaskDeletedWaiter) WaitForOutput(ctx context.Context, params *DescribeConversionTasksInput, maxWaitDur time.Duration, optFns ...func(*ConversionTaskDeletedWaiterOptions)) (*DescribeConversionTasksOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.DescribeConversionTasks(ctx, params, func(o *Options) {
			baseOpts := []func(*Options){
				addIsWaiterUserAgent,
			}
			o.APIOptions = append(o.APIOptions, apiOptions...)
			for _, opt := range baseOpts {
				opt(o)
			}
			for _, opt := range options.ClientOptions {
				opt(o)
			}
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for ConversionTaskDeleted waiter")
}

func conversionTaskDeletedStateRetryable(ctx context.Context, input *DescribeConversionTasksInput, output *DescribeConversionTasksOutput, err error) (bool, error) {

	if err == nil {
		v1 := output.ConversionTasks
		var v2 []types.ConversionTaskState
		for _, v := range v1 {
			v3 := v.State
			v2 = append(v2, v3)
		}
		expectedValue := "deleted"
		match := len(v2) > 0
		for _, v := range v2 {
			if string(v) != expectedValue {
				match = false
				break
			}
		}

		if match {
			return false, nil
		}
	}

	if err != nil {
		return false, err
	}
	return true, nil
}

// DescribeConversionTasksAPIClient is a client that implements the
// DescribeConversionTasks operation.
type DescribeConversionTasksAPIClient interface {
	DescribeConversionTasks(context.Context, *DescribeConversionTasksInput, ...func(*Options)) (*DescribeConversionTasksOutput, error)
}

var _ DescribeConversionTasksAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeConversionTasks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeConversionTasks",
	}
}
