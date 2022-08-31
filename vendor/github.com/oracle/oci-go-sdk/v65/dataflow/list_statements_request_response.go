// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataflow

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListStatementsRequest wrapper for the ListStatements operation
type ListStatementsRequest struct {

	// The unique ID for the run
	RunId *string `mandatory:"true" contributesTo:"path" name:"runId"`

	// The LifecycleState of the statement.
	LifecycleState ListStatementsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The field used to sort the results. Multiple fields are not supported.
	SortBy ListStatementsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The ordering of results in ascending or descending order.
	SortOrder ListStatementsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of results to return in a paginated `List` call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` or `opc-prev-page` response header from the last `List` call
	// to sent back to server for getting the next page of results.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique identifier for the request. If provided, the returned request ID will include this value.
	// Otherwise, a random request ID will be generated by the service.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListStatementsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListStatementsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListStatementsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListStatementsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListStatementsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListStatementsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListStatementsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListStatementsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListStatementsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListStatementsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListStatementsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListStatementsResponse wrapper for the ListStatements operation
type ListStatementsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of StatementCollection instances
	StatementCollection `presentIn:"body"`

	// Retrieves the previous page of results.
	// When this header appears in the response, previous pages of results exist.
	// See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Retrieves the next page of results. When this header appears in the response,
	// additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListStatementsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListStatementsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListStatementsLifecycleStateEnum Enum with underlying type: string
type ListStatementsLifecycleStateEnum string

// Set of constants representing the allowable values for ListStatementsLifecycleStateEnum
const (
	ListStatementsLifecycleStateAccepted   ListStatementsLifecycleStateEnum = "ACCEPTED"
	ListStatementsLifecycleStateCancelling ListStatementsLifecycleStateEnum = "CANCELLING"
	ListStatementsLifecycleStateCancelled  ListStatementsLifecycleStateEnum = "CANCELLED"
	ListStatementsLifecycleStateFailed     ListStatementsLifecycleStateEnum = "FAILED"
	ListStatementsLifecycleStateInProgress ListStatementsLifecycleStateEnum = "IN_PROGRESS"
	ListStatementsLifecycleStateSucceeded  ListStatementsLifecycleStateEnum = "SUCCEEDED"
)

var mappingListStatementsLifecycleStateEnum = map[string]ListStatementsLifecycleStateEnum{
	"ACCEPTED":    ListStatementsLifecycleStateAccepted,
	"CANCELLING":  ListStatementsLifecycleStateCancelling,
	"CANCELLED":   ListStatementsLifecycleStateCancelled,
	"FAILED":      ListStatementsLifecycleStateFailed,
	"IN_PROGRESS": ListStatementsLifecycleStateInProgress,
	"SUCCEEDED":   ListStatementsLifecycleStateSucceeded,
}

var mappingListStatementsLifecycleStateEnumLowerCase = map[string]ListStatementsLifecycleStateEnum{
	"accepted":    ListStatementsLifecycleStateAccepted,
	"cancelling":  ListStatementsLifecycleStateCancelling,
	"cancelled":   ListStatementsLifecycleStateCancelled,
	"failed":      ListStatementsLifecycleStateFailed,
	"in_progress": ListStatementsLifecycleStateInProgress,
	"succeeded":   ListStatementsLifecycleStateSucceeded,
}

// GetListStatementsLifecycleStateEnumValues Enumerates the set of values for ListStatementsLifecycleStateEnum
func GetListStatementsLifecycleStateEnumValues() []ListStatementsLifecycleStateEnum {
	values := make([]ListStatementsLifecycleStateEnum, 0)
	for _, v := range mappingListStatementsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListStatementsLifecycleStateEnumStringValues Enumerates the set of values in String for ListStatementsLifecycleStateEnum
func GetListStatementsLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"CANCELLING",
		"CANCELLED",
		"FAILED",
		"IN_PROGRESS",
		"SUCCEEDED",
	}
}

// GetMappingListStatementsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListStatementsLifecycleStateEnum(val string) (ListStatementsLifecycleStateEnum, bool) {
	enum, ok := mappingListStatementsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListStatementsSortByEnum Enum with underlying type: string
type ListStatementsSortByEnum string

// Set of constants representing the allowable values for ListStatementsSortByEnum
const (
	ListStatementsSortByTimecreated ListStatementsSortByEnum = "timeCreated"
)

var mappingListStatementsSortByEnum = map[string]ListStatementsSortByEnum{
	"timeCreated": ListStatementsSortByTimecreated,
}

var mappingListStatementsSortByEnumLowerCase = map[string]ListStatementsSortByEnum{
	"timecreated": ListStatementsSortByTimecreated,
}

// GetListStatementsSortByEnumValues Enumerates the set of values for ListStatementsSortByEnum
func GetListStatementsSortByEnumValues() []ListStatementsSortByEnum {
	values := make([]ListStatementsSortByEnum, 0)
	for _, v := range mappingListStatementsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListStatementsSortByEnumStringValues Enumerates the set of values in String for ListStatementsSortByEnum
func GetListStatementsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
	}
}

// GetMappingListStatementsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListStatementsSortByEnum(val string) (ListStatementsSortByEnum, bool) {
	enum, ok := mappingListStatementsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListStatementsSortOrderEnum Enum with underlying type: string
type ListStatementsSortOrderEnum string

// Set of constants representing the allowable values for ListStatementsSortOrderEnum
const (
	ListStatementsSortOrderAsc  ListStatementsSortOrderEnum = "ASC"
	ListStatementsSortOrderDesc ListStatementsSortOrderEnum = "DESC"
)

var mappingListStatementsSortOrderEnum = map[string]ListStatementsSortOrderEnum{
	"ASC":  ListStatementsSortOrderAsc,
	"DESC": ListStatementsSortOrderDesc,
}

var mappingListStatementsSortOrderEnumLowerCase = map[string]ListStatementsSortOrderEnum{
	"asc":  ListStatementsSortOrderAsc,
	"desc": ListStatementsSortOrderDesc,
}

// GetListStatementsSortOrderEnumValues Enumerates the set of values for ListStatementsSortOrderEnum
func GetListStatementsSortOrderEnumValues() []ListStatementsSortOrderEnum {
	values := make([]ListStatementsSortOrderEnum, 0)
	for _, v := range mappingListStatementsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListStatementsSortOrderEnumStringValues Enumerates the set of values in String for ListStatementsSortOrderEnum
func GetListStatementsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListStatementsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListStatementsSortOrderEnum(val string) (ListStatementsSortOrderEnum, bool) {
	enum, ok := mappingListStatementsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
