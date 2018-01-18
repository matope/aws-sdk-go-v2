// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package costexplorer

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

const opGetCostAndUsage = "GetCostAndUsage"

// GetCostAndUsageRequest is a API request type for the GetCostAndUsage API operation.
type GetCostAndUsageRequest struct {
	*aws.Request
	Input *GetCostAndUsageInput
}

// Send marshals and sends the GetCostAndUsage API request.
func (r GetCostAndUsageRequest) Send() (*GetCostAndUsageOutput, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*GetCostAndUsageOutput), nil
}

// GetCostAndUsageRequest returns a request value for making API operation for
// AWS Cost Explorer Service.
//
// Retrieve cost and usage metrics for your account. You can specify which cost
// and usage-related metric, such as BlendedCosts or UsageQuantity, that you
// want the request to return. You can also filter and group your data by various
// dimensions, such as AWS Service or AvailabilityZone, in a specific time range.
// See the GetDimensionValues action for a complete list of the valid dimensions.
// Master accounts in an organization have access to all member accounts.
//
//    // Example sending a request using the GetCostAndUsageRequest method.
//    req := client.GetCostAndUsageRequest(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ce-2017-10-25/GetCostAndUsage
func (c *CostExplorer) GetCostAndUsageRequest(input *GetCostAndUsageInput) GetCostAndUsageRequest {
	op := &aws.Operation{
		Name:       opGetCostAndUsage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetCostAndUsageInput{}
	}

	output := &GetCostAndUsageOutput{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return GetCostAndUsageRequest{Request: req, Input: input}
}

const opGetDimensionValues = "GetDimensionValues"

// GetDimensionValuesRequest is a API request type for the GetDimensionValues API operation.
type GetDimensionValuesRequest struct {
	*aws.Request
	Input *GetDimensionValuesInput
}

// Send marshals and sends the GetDimensionValues API request.
func (r GetDimensionValuesRequest) Send() (*GetDimensionValuesOutput, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*GetDimensionValuesOutput), nil
}

// GetDimensionValuesRequest returns a request value for making API operation for
// AWS Cost Explorer Service.
//
// You can use GetDimensionValues to retrieve all available filter values for
// a specific filter over a period of time. You can search the dimension values
// for an arbitrary string.
//
//    // Example sending a request using the GetDimensionValuesRequest method.
//    req := client.GetDimensionValuesRequest(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ce-2017-10-25/GetDimensionValues
func (c *CostExplorer) GetDimensionValuesRequest(input *GetDimensionValuesInput) GetDimensionValuesRequest {
	op := &aws.Operation{
		Name:       opGetDimensionValues,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetDimensionValuesInput{}
	}

	output := &GetDimensionValuesOutput{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return GetDimensionValuesRequest{Request: req, Input: input}
}

const opGetReservationUtilization = "GetReservationUtilization"

// GetReservationUtilizationRequest is a API request type for the GetReservationUtilization API operation.
type GetReservationUtilizationRequest struct {
	*aws.Request
	Input *GetReservationUtilizationInput
}

// Send marshals and sends the GetReservationUtilization API request.
func (r GetReservationUtilizationRequest) Send() (*GetReservationUtilizationOutput, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*GetReservationUtilizationOutput), nil
}

// GetReservationUtilizationRequest returns a request value for making API operation for
// AWS Cost Explorer Service.
//
// You can retrieve the Reservation utilization for your account. Master accounts
// in an organization have access to their associated member accounts. You can
// filter data by dimensions in a time period. You can use GetDimensionValues
// to determine the possible dimension values. Currently, you can group only
// by SUBSCRIPTION_ID.
//
//    // Example sending a request using the GetReservationUtilizationRequest method.
//    req := client.GetReservationUtilizationRequest(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ce-2017-10-25/GetReservationUtilization
func (c *CostExplorer) GetReservationUtilizationRequest(input *GetReservationUtilizationInput) GetReservationUtilizationRequest {
	op := &aws.Operation{
		Name:       opGetReservationUtilization,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetReservationUtilizationInput{}
	}

	output := &GetReservationUtilizationOutput{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return GetReservationUtilizationRequest{Request: req, Input: input}
}

const opGetTags = "GetTags"

// GetTagsRequest is a API request type for the GetTags API operation.
type GetTagsRequest struct {
	*aws.Request
	Input *GetTagsInput
}

// Send marshals and sends the GetTags API request.
func (r GetTagsRequest) Send() (*GetTagsOutput, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*GetTagsOutput), nil
}

// GetTagsRequest returns a request value for making API operation for
// AWS Cost Explorer Service.
//
// You can query for available tag keys and tag values for a specified period.
// You can search the tag values for an arbitrary string.
//
//    // Example sending a request using the GetTagsRequest method.
//    req := client.GetTagsRequest(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ce-2017-10-25/GetTags
func (c *CostExplorer) GetTagsRequest(input *GetTagsInput) GetTagsRequest {
	op := &aws.Operation{
		Name:       opGetTags,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetTagsInput{}
	}

	output := &GetTagsOutput{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return GetTagsRequest{Request: req, Input: input}
}

// The time period that you want the usage and costs for.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ce-2017-10-25/DateInterval
type DateInterval struct {
	_ struct{} `type:"structure"`

	// The end of the time period that you want the usage and costs for. The end
	// date is exclusive. For example, if the end is 2017-05-01, then the cost and
	// usage data is retrieved from the start date but not including 2017-05-01.
	//
	// End is a required field
	End *string `type:"string" required:"true"`

	// The beginning of the time period that you want the usage and costs for. The
	// start date is inclusive. For example, if start is 2017-01-01, then the cost
	// and usage data is retrieved starting at 2017-01-01 up to the end date.
	//
	// Start is a required field
	Start *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DateInterval) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DateInterval) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DateInterval) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DateInterval"}

	if s.End == nil {
		invalidParams.Add(aws.NewErrParamRequired("End"))
	}

	if s.Start == nil {
		invalidParams.Add(aws.NewErrParamRequired("Start"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The metadata that you can use to filter and group your results. You can use
// GetDimensionValues to find specific values.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ce-2017-10-25/DimensionValues
type DimensionValues struct {
	_ struct{} `type:"structure"`

	// The names of the metadata types that you can use to filter and group your
	// results. For example, AZ returns a list of Availability Zones.
	Key Dimension `type:"string" enum:"true"`

	// The metadata values that you can use to filter and group your results. You
	// can use GetDimensionValues to find specific values.
	Values []string `type:"list"`
}

// String returns the string representation
func (s DimensionValues) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DimensionValues) GoString() string {
	return s.String()
}

// The metadata of a specific type that you can use to filter and group your
// results. You can use GetDimensionValues to find specific values.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ce-2017-10-25/DimensionValuesWithAttributes
type DimensionValuesWithAttributes struct {
	_ struct{} `type:"structure"`

	// The attribute that applies to a specific Dimension.
	Attributes map[string]string `type:"map"`

	// The value of a dimension with a specific attribute.
	Value *string `type:"string"`
}

// String returns the string representation
func (s DimensionValuesWithAttributes) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DimensionValuesWithAttributes) GoString() string {
	return s.String()
}

// Use Expression to filter by cost or by usage. There are two patterns:
//
//    * Simple dimension values - You can set the dimension name and values
//    for the filters that you plan to use. For example, you can filter for
//    InstanceType==m4.xlarge OR InstanceType==c4.large. The Expression for
//    that looks like this.
//
// { "Dimensions": { "Key": "InstanceType", "Values": [ "m4.xlarge", “c4.large”
//    ] } }
//
// The list of dimension values are OR'd together to retrieve cost or usage
//    data. You can create Expression and DimensionValues objects using either
//    with* methods or set* methods in multiple lines.
//
//    * Compound dimension values with logical operations - You can use multiple
//    Expression types and the logical operators AND/OR/NOT to create a list
//    of one or more Expression objects. This allows you to filter on more advanced
//    options. For example, you can filter on ((InstanceType == m4.large OR
//    InstanceType == m3.large) OR (Tag.Type == Type1)) AND (UsageType != DataTransfer).
//    The Expression for that looks like this.
//
// { "And": [ {"Or": [ {"Dimensions": { "Key": "InstanceType", "Values": [ "m4.x.large",
//    "c4.large" ] }}, {"Tag": { "Key": "TagName", "Values": ["Value1"] } }
//    ]}, {"Not": {"dimensions": { "Key": "UsageType", "Values": ["DataTransfer"]
//    }}} ] }
//
// Because each Expression can have only one operator, the service returns an
//    error if more than one is specified. The following example shows an Expression
//    object that will create an error.
//
//  { "And": [ ... ], "DimensionValues": { "Dimension": "UsageType", "Values":
//    [ "DataTransfer" ] } }
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ce-2017-10-25/Expression
type Expression struct {
	_ struct{} `type:"structure"`

	// Return results that match both Dimension objects.
	And []Expression `type:"list"`

	// The specific Dimension to use for Expression.
	Dimensions *DimensionValues `type:"structure"`

	// Return results that don't match Dimension.
	Not *Expression `type:"structure"`

	// Return results that match either Dimension.
	Or []Expression `type:"list"`

	// The specific Tag to use for Expression.
	Tags *TagValues `type:"structure"`
}

// String returns the string representation
func (s Expression) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Expression) GoString() string {
	return s.String()
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ce-2017-10-25/GetCostAndUsageRequest
type GetCostAndUsageInput struct {
	_ struct{} `type:"structure"`

	// Filters AWS costs by different dimensions. For example, you can specify Service
	// and Linked Account and get the costs associated with that account's usage
	// of that service. You can nest Expression objects to define any combination
	// of dimension filters. For more information, see the Expression object or
	// More Examples.
	Filter *Expression `type:"structure"`

	// Sets the AWS cost granularity to MONTHLY or DAILY.
	Granularity Granularity `type:"string" enum:"true"`

	// You can group AWS costs using up to two different groups, either dimensions,
	// tag keys, or both.
	//
	// When you group by tag key, you get all tag values, including empty strings.
	//
	// Valid values are: AZ, INSTANCE_TYPE, LINKED_ACCCOUNT, OPERATION, PURCHASE_TYPE,
	// SERVICE, USAGE_TYPE, TAGS, and PLATFORM.
	GroupBy []GroupDefinition `type:"list"`

	// Which metrics are returned in the query. For more information about blended
	// and unblended rates, see https://aws.amazon.com/premiumsupport/knowledge-center/blended-rates-intro/.
	//
	// Valid values are BlendedCost, UnblendedCost, and UsageQuantity.
	//
	// If you return the UsageQuantity metric, the service aggregates all usage
	// numbers without taking into account the units. For example, if you aggregate
	// usageQuantity across all of EC2, the results aren't meaningful because EC2
	// compute hours and data transfer are measured in different units (for example,
	// hours vs. GB). To get more meaningful UsageQuantity metrics, filter by UsageType
	// or UsageTypeGroups.
	Metrics []string `type:"list"`

	// The token to retrieve the next set of results. AWS provides the token when
	// the response from a previous call has more results than the maximum page
	// size.
	NextPageToken *string `type:"string"`

	// Sets the start and end dates for retrieving AWS costs. The start date is
	// inclusive, but the end date is exclusive. For example, if start is 2017-01-01
	// and end is 2017-05-01, then the cost and usage data is retrieved from 2017-01-01
	// up to and including 2017-04-30 but not including 2017-05-01.
	TimePeriod *DateInterval `type:"structure"`
}

// String returns the string representation
func (s GetCostAndUsageInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetCostAndUsageInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetCostAndUsageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetCostAndUsageInput"}
	if s.TimePeriod != nil {
		if err := s.TimePeriod.Validate(); err != nil {
			invalidParams.AddNested("TimePeriod", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ce-2017-10-25/GetCostAndUsageResponse
type GetCostAndUsageOutput struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response

	// The groups specified by the the Filter or GroupBy parameters in the request.
	GroupDefinitions []GroupDefinition `type:"list"`

	// The token for the next set of retrievable results. AWS provides the token
	// when the response from a previous call has more results than the maximum
	// page size.
	NextPageToken *string `type:"string"`

	// The time period covered by the results in the response.
	ResultsByTime []ResultByTime `type:"list"`
}

// String returns the string representation
func (s GetCostAndUsageOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetCostAndUsageOutput) GoString() string {
	return s.String()
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s GetCostAndUsageOutput) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ce-2017-10-25/GetDimensionValuesRequest
type GetDimensionValuesInput struct {
	_ struct{} `type:"structure"`

	// The context for the call to GetDimensionValues. This can be RESERVED_INSTANCE
	// or COST_AND_USAGE. The default value is COST_AND_USAGE. If the context is
	// set to RESERVED_INSTANCE, the resulting dimension values can be used in the
	// GetReservationUtilization action. If the context is set to COST_AND_USAGE,
	// , the resulting dimension values can be used in the GetCostAndUsage operation.
	//
	// If you set the context to CostAndUsage, you can use the following dimensions
	// for searching:
	//
	//    * AZ - The Availability Zone. An example is us-east-1a.
	//
	//    * InstanceType - The type of EC2 instance. An example is m4.xlarge.
	//
	//    * LinkedAccount - The description in the attribute map that includes the
	//    full name of the member account. The value field contains the AWS ID of
	//    the member account
	//
	//    * Operation - The action performed. Examples include RunInstance and CreateBucket.
	//
	//    * PurchaseType - The reservation type of the purchase to which this usage
	//    is related. Examples include: On Demand Instances and Standard Reserved
	//    Instances
	//
	//    * Service - The AWS service such as DynamoDB.
	//
	//    * UsageType -The type of usage. An example is DataTransfer-In-Bytes. The
	//    response for the GetDimensionValues action includes a unit attribute,
	//    examples of which include GB and Hrs.
	//
	//    * UsageTypeGroup - The grouping of common usage types. An example is EC2:
	//    CloudWatch – Alarms. The response for this action includes a unit attribute.
	//
	//    * RecordType - The different types of charges such as RI fees, usage costs,
	//    tax refunds, and credits
	//
	// If you set the context to ReservedInstance, you can use the following dimensions
	// for searching:
	//
	//    * AZ - The Availability Zone. An example is us-east-1a.
	//
	//    * InstanceType - The type of EC2 instance. An example is m4.xlarge.
	//
	//    * LinkedAccount - The description in the attribute map that includes the
	//    full name of the member account. The value field contains the AWS ID of
	//    the member account
	//
	//    * Platform - The operating system. Examples are Windows or Linux.
	//
	//    * Region - The AWS region.
	//
	//    * Scope - The scope of a reserved instance (RI). Values are regional or
	//    a single availability zone.
	//
	//    * Tenancy - The tenancy of a resource. Examples are shared or dedicated.
	Context Context `type:"string" enum:"true"`

	// The name of the dimension. Different Dimensionsare available for different
	// Contexts. For more information, see Context.
	//
	// Dimension is a required field
	Dimension Dimension `type:"string" required:"true" enum:"true"`

	// The token to retrieve the next set of results. AWS provides the token when
	// the response from a previous call has more results than the maximum page
	// size.
	NextPageToken *string `type:"string"`

	// The value that you want to search the filter values for.
	SearchString *string `type:"string"`

	// The start and end dates for retrieving the dimension values. The start date
	// is inclusive, but the end date is exclusive. For example, if start is 2017-01-01
	// and end is 2017-05-01, then the cost and usage data is retrieved from 2017-01-01
	// up to and including 2017-04-30 but not including 2017-05-01.
	//
	// TimePeriod is a required field
	TimePeriod *DateInterval `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetDimensionValuesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetDimensionValuesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDimensionValuesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDimensionValuesInput"}
	if len(s.Dimension) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Dimension"))
	}

	if s.TimePeriod == nil {
		invalidParams.Add(aws.NewErrParamRequired("TimePeriod"))
	}
	if s.TimePeriod != nil {
		if err := s.TimePeriod.Validate(); err != nil {
			invalidParams.AddNested("TimePeriod", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ce-2017-10-25/GetDimensionValuesResponse
type GetDimensionValuesOutput struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response

	// The filters that you used to filter your request. Some dimensions are available
	// only for a specific context:
	//
	// If you set the context to CostAndUsage, you can use the following dimensions
	// for searching:
	//
	//    * AZ - The Availability Zone. An example is us-east-1a.
	//
	//    * InstanceType - The type of EC2 instance. An example is m4.xlarge.
	//
	//    * LinkedAccount - The description in the attribute map that includes the
	//    full name of the member account. The value field contains the AWS ID of
	//    the member account
	//
	//    * Operation - The action performed. Examples include RunInstance and CreateBucket.
	//
	//    * PurchaseType - The reservation type of the purchase to which this usage
	//    is related. Examples include: On Demand Instances and Standard Reserved
	//    Instances
	//
	//    * Service - The AWS service such as DynamoDB.
	//
	//    * UsageType -The type of usage. An example is DataTransfer-In-Bytes. The
	//    response for the GetDimensionValues action includes a unit attribute,
	//    examples of which include GB and Hrs.
	//
	//    * UsageTypeGroup - The grouping of common usage types. An example is EC2:
	//    CloudWatch – Alarms. The response for this action includes a unit attribute.
	//
	//    * RecordType - The different types of charges such as RI fees, usage costs,
	//    tax refunds, and credits
	//
	// If you set the context to ReservedInstance, you can use the following dimensions
	// for searching:
	//
	//    * AZ - The Availability Zone. An example is us-east-1a.
	//
	//    * InstanceType - The type of EC2 instance. An example is m4.xlarge.
	//
	//    * LinkedAccount - The description in the attribute map that includes the
	//    full name of the member account. The value field contains the AWS ID of
	//    the member account
	//
	//    * Platform - The operating system. Examples are Windows or Linux.
	//
	//    * Region - The AWS region.
	//
	//    * Scope - The scope of a reserved instance (RI). Values are regional or
	//    a single availability zone.
	//
	//    * Tenancy - The tenancy of a resource. Examples are shared or dedicated.
	//
	// DimensionValues is a required field
	DimensionValues []DimensionValuesWithAttributes `type:"list" required:"true"`

	// The token for the next set of retrievable results. AWS provides the token
	// when the response from a previous call has more results than the maximum
	// page size.
	NextPageToken *string `type:"string"`

	// The number of results that AWS returned at one time.
	//
	// ReturnSize is a required field
	ReturnSize *int64 `type:"integer" required:"true"`

	// The total number of search results.
	//
	// TotalSize is a required field
	TotalSize *int64 `type:"integer" required:"true"`
}

// String returns the string representation
func (s GetDimensionValuesOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetDimensionValuesOutput) GoString() string {
	return s.String()
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s GetDimensionValuesOutput) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ce-2017-10-25/GetReservationUtilizationRequest
type GetReservationUtilizationInput struct {
	_ struct{} `type:"structure"`

	// Filters utilization data by using different dimensions. GetReservationUtilization
	// uses the same Expression object as the other operations, but only AND is
	// supported among each dimension, and nesting is supported up to only one level
	// deep. If there are multiple values for a dimension, they are OR'd together.
	Filter *Expression `type:"structure"`

	// Sets the AWS cost granularity to MONTHLY or DAILY. If both GroupBy and granularity
	// are not set, GetReservationUtilization defaults to DAILY. If GroupBy is set,
	// Granularity can't be set, and the response object doesn't include MONTHLY
	// or DAILY granularity.
	Granularity Granularity `type:"string" enum:"true"`

	// Groups only by SubscriptionId. Metadata is included.
	GroupBy []GroupDefinition `type:"list"`

	// The token to retrieve the next set of results. AWS provides the token when
	// the response from a previous call has more results than the maximum page
	// size.
	NextPageToken *string `type:"string"`

	// Sets the start and end dates for retrieving reserve instance (RI) utilization.
	// The start date is inclusive, but the end date is exclusive. For example,
	// if start is 2017-01-01 and end is 2017-05-01, then the cost and usage data
	// is retrieved from 2017-01-01 up to and including 2017-04-30 but not including
	// 2017-05-01.
	//
	// TimePeriod is a required field
	TimePeriod *DateInterval `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetReservationUtilizationInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetReservationUtilizationInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetReservationUtilizationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetReservationUtilizationInput"}

	if s.TimePeriod == nil {
		invalidParams.Add(aws.NewErrParamRequired("TimePeriod"))
	}
	if s.TimePeriod != nil {
		if err := s.TimePeriod.Validate(); err != nil {
			invalidParams.AddNested("TimePeriod", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ce-2017-10-25/GetReservationUtilizationResponse
type GetReservationUtilizationOutput struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response

	// The token for the next set of retrievable results. AWS provides the token
	// when the response from a previous call has more results than the maximum
	// page size.
	NextPageToken *string `type:"string"`

	// The total amount of time that you utilized your RIs.
	Total *ReservationAggregates `type:"structure"`

	// The amount of time that you utilized your RIs.
	//
	// UtilizationsByTime is a required field
	UtilizationsByTime []UtilizationByTime `type:"list" required:"true"`
}

// String returns the string representation
func (s GetReservationUtilizationOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetReservationUtilizationOutput) GoString() string {
	return s.String()
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s GetReservationUtilizationOutput) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ce-2017-10-25/GetTagsRequest
type GetTagsInput struct {
	_ struct{} `type:"structure"`

	// The token to retrieve the next set of results. AWS provides the token when
	// the response from a previous call has more results than the maximum page
	// size.
	NextPageToken *string `type:"string"`

	// The value that you want to search for.
	SearchString *string `type:"string"`

	// The key of the tag that you want to return values for.
	TagKey *string `type:"string"`

	// The start and end dates for retrieving the dimension values. The start date
	// is inclusive, but the end date is exclusive. For example, if start is 2017-01-01
	// and end is 2017-05-01, then the cost and usage data is retrieved from 2017-01-01
	// up to and including 2017-04-30 but not including 2017-05-01.
	//
	// TimePeriod is a required field
	TimePeriod *DateInterval `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetTagsInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetTagsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetTagsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetTagsInput"}

	if s.TimePeriod == nil {
		invalidParams.Add(aws.NewErrParamRequired("TimePeriod"))
	}
	if s.TimePeriod != nil {
		if err := s.TimePeriod.Validate(); err != nil {
			invalidParams.AddNested("TimePeriod", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ce-2017-10-25/GetTagsResponse
type GetTagsOutput struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response

	// The token for the next set of retrievable results. AWS provides the token
	// when the response from a previous call has more results than the maximum
	// page size.
	NextPageToken *string `type:"string"`

	// The number of query results that AWS returns at a time.
	//
	// ReturnSize is a required field
	ReturnSize *int64 `type:"integer" required:"true"`

	// The tags that match your request.
	//
	// Tags is a required field
	Tags []string `type:"list" required:"true"`

	// The total number of query results.
	//
	// TotalSize is a required field
	TotalSize *int64 `type:"integer" required:"true"`
}

// String returns the string representation
func (s GetTagsOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetTagsOutput) GoString() string {
	return s.String()
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s GetTagsOutput) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// One level of grouped data within the results.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ce-2017-10-25/Group
type Group struct {
	_ struct{} `type:"structure"`

	// The keys included in this group.
	Keys []string `type:"list"`

	// The metrics included in this group.
	Metrics map[string]MetricValue `type:"map"`
}

// String returns the string representation
func (s Group) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Group) GoString() string {
	return s.String()
}

// Represents a group when you specify a group by criteria, or in the response
// to a query with a specific grouping.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ce-2017-10-25/GroupDefinition
type GroupDefinition struct {
	_ struct{} `type:"structure"`

	// The string that represents a key for a specified group.
	Key *string `type:"string"`

	// The string that represents the type of group.
	Type GroupDefinitionType `type:"string" enum:"true"`
}

// String returns the string representation
func (s GroupDefinition) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GroupDefinition) GoString() string {
	return s.String()
}

// The aggregated value for a metric.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ce-2017-10-25/MetricValue
type MetricValue struct {
	_ struct{} `type:"structure"`

	// The actual number that represents the metric.
	Amount *string `type:"string"`

	// The unit that the metric is given in.
	Unit *string `type:"string"`
}

// String returns the string representation
func (s MetricValue) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s MetricValue) GoString() string {
	return s.String()
}

// The aggregated numbers for your RI usage.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ce-2017-10-25/ReservationAggregates
type ReservationAggregates struct {
	_ struct{} `type:"structure"`

	// How many RI hours you purchased.
	PurchasedHours *string `type:"string"`

	// The total number of RI hours that you used.
	TotalActualHours *string `type:"string"`

	// The number of RI hours that you didn't use.
	UnusedHours *string `type:"string"`

	// The percentage of RI time that you used.
	UtilizationPercentage *string `type:"string"`
}

// String returns the string representation
func (s ReservationAggregates) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ReservationAggregates) GoString() string {
	return s.String()
}

// A group of RIs that share a set of attributes.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ce-2017-10-25/ReservationUtilizationGroup
type ReservationUtilizationGroup struct {
	_ struct{} `type:"structure"`

	// The attributes for this group of RIs.
	Attributes map[string]string `type:"map"`

	// The key for a specific RI attribute.
	Key *string `type:"string"`

	// How much you used this group of RIs.
	Utilization *ReservationAggregates `type:"structure"`

	// The value of a specific RI attribute.
	Value *string `type:"string"`
}

// String returns the string representation
func (s ReservationUtilizationGroup) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ReservationUtilizationGroup) GoString() string {
	return s.String()
}

// The result that is associated with a time period.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ce-2017-10-25/ResultByTime
type ResultByTime struct {
	_ struct{} `type:"structure"`

	// Whether or not this result is estimated.
	Estimated *bool `type:"boolean"`

	// The groups that are included in this time period.
	Groups []Group `type:"list"`

	// The time period covered by a result.
	TimePeriod *DateInterval `type:"structure"`

	// The total amount of cost or usage accrued during the time period.
	Total map[string]MetricValue `type:"map"`
}

// String returns the string representation
func (s ResultByTime) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ResultByTime) GoString() string {
	return s.String()
}

// The values that are available for a tag.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ce-2017-10-25/TagValues
type TagValues struct {
	_ struct{} `type:"structure"`

	// The key for a tag.
	Key *string `type:"string"`

	// The specific value of a tag.
	Values []string `type:"list"`
}

// String returns the string representation
func (s TagValues) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s TagValues) GoString() string {
	return s.String()
}

// The amount of utilization, in hours.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ce-2017-10-25/UtilizationByTime
type UtilizationByTime struct {
	_ struct{} `type:"structure"`

	// The groups that are included in this utilization result.
	Groups []ReservationUtilizationGroup `type:"list"`

	// The period of time over which this utilization was used.
	TimePeriod *DateInterval `type:"structure"`

	// The total number of RI hours that were used.
	Total *ReservationAggregates `type:"structure"`
}

// String returns the string representation
func (s UtilizationByTime) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s UtilizationByTime) GoString() string {
	return s.String()
}

type Context string

// Enum values for Context
const (
	ContextCostAndUsage Context = "COST_AND_USAGE"
	ContextReservations Context = "RESERVATIONS"
)

type Dimension string

// Enum values for Dimension
const (
	DimensionAz              Dimension = "AZ"
	DimensionInstanceType    Dimension = "INSTANCE_TYPE"
	DimensionLinkedAccount   Dimension = "LINKED_ACCOUNT"
	DimensionOperation       Dimension = "OPERATION"
	DimensionPurchaseType    Dimension = "PURCHASE_TYPE"
	DimensionRegion          Dimension = "REGION"
	DimensionService         Dimension = "SERVICE"
	DimensionUsageType       Dimension = "USAGE_TYPE"
	DimensionUsageTypeGroup  Dimension = "USAGE_TYPE_GROUP"
	DimensionRecordType      Dimension = "RECORD_TYPE"
	DimensionOperatingSystem Dimension = "OPERATING_SYSTEM"
	DimensionTenancy         Dimension = "TENANCY"
	DimensionScope           Dimension = "SCOPE"
	DimensionPlatform        Dimension = "PLATFORM"
	DimensionSubscriptionId  Dimension = "SUBSCRIPTION_ID"
)

type Granularity string

// Enum values for Granularity
const (
	GranularityDaily   Granularity = "DAILY"
	GranularityMonthly Granularity = "MONTHLY"
)

type GroupDefinitionType string

// Enum values for GroupDefinitionType
const (
	GroupDefinitionTypeDimension GroupDefinitionType = "DIMENSION"
	GroupDefinitionTypeTag       GroupDefinitionType = "TAG"
)