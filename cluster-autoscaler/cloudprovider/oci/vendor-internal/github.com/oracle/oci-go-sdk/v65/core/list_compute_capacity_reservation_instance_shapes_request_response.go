// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package core

import (
	"fmt"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/oci/vendor-internal/github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListComputeCapacityReservationInstanceShapesRequest wrapper for the ListComputeCapacityReservationInstanceShapes operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/core/ListComputeCapacityReservationInstanceShapes.go.html to see an example of how to use ListComputeCapacityReservationInstanceShapesRequest.
type ListComputeCapacityReservationInstanceShapesRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The name of the availability domain.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"false" contributesTo:"query" name:"availabilityDomain"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List"
	// call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by availability domain if the scope of the resource type is within a
	// single availability domain. If you call one of these "List" operations without specifying
	// an availability domain, the resources are grouped by availability domain, then sorted.
	SortBy ListComputeCapacityReservationInstanceShapesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	// is case sensitive.
	SortOrder ListComputeCapacityReservationInstanceShapesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListComputeCapacityReservationInstanceShapesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListComputeCapacityReservationInstanceShapesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListComputeCapacityReservationInstanceShapesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListComputeCapacityReservationInstanceShapesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListComputeCapacityReservationInstanceShapesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListComputeCapacityReservationInstanceShapesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListComputeCapacityReservationInstanceShapesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListComputeCapacityReservationInstanceShapesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListComputeCapacityReservationInstanceShapesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListComputeCapacityReservationInstanceShapesResponse wrapper for the ListComputeCapacityReservationInstanceShapes operation
type ListComputeCapacityReservationInstanceShapesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ComputeCapacityReservationInstanceShapeSummary instances
	Items []ComputeCapacityReservationInstanceShapeSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListComputeCapacityReservationInstanceShapesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListComputeCapacityReservationInstanceShapesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListComputeCapacityReservationInstanceShapesSortByEnum Enum with underlying type: string
type ListComputeCapacityReservationInstanceShapesSortByEnum string

// Set of constants representing the allowable values for ListComputeCapacityReservationInstanceShapesSortByEnum
const (
	ListComputeCapacityReservationInstanceShapesSortByTimecreated ListComputeCapacityReservationInstanceShapesSortByEnum = "TIMECREATED"
	ListComputeCapacityReservationInstanceShapesSortByDisplayname ListComputeCapacityReservationInstanceShapesSortByEnum = "DISPLAYNAME"
)

var mappingListComputeCapacityReservationInstanceShapesSortByEnum = map[string]ListComputeCapacityReservationInstanceShapesSortByEnum{
	"TIMECREATED": ListComputeCapacityReservationInstanceShapesSortByTimecreated,
	"DISPLAYNAME": ListComputeCapacityReservationInstanceShapesSortByDisplayname,
}

var mappingListComputeCapacityReservationInstanceShapesSortByEnumLowerCase = map[string]ListComputeCapacityReservationInstanceShapesSortByEnum{
	"timecreated": ListComputeCapacityReservationInstanceShapesSortByTimecreated,
	"displayname": ListComputeCapacityReservationInstanceShapesSortByDisplayname,
}

// GetListComputeCapacityReservationInstanceShapesSortByEnumValues Enumerates the set of values for ListComputeCapacityReservationInstanceShapesSortByEnum
func GetListComputeCapacityReservationInstanceShapesSortByEnumValues() []ListComputeCapacityReservationInstanceShapesSortByEnum {
	values := make([]ListComputeCapacityReservationInstanceShapesSortByEnum, 0)
	for _, v := range mappingListComputeCapacityReservationInstanceShapesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListComputeCapacityReservationInstanceShapesSortByEnumStringValues Enumerates the set of values in String for ListComputeCapacityReservationInstanceShapesSortByEnum
func GetListComputeCapacityReservationInstanceShapesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListComputeCapacityReservationInstanceShapesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListComputeCapacityReservationInstanceShapesSortByEnum(val string) (ListComputeCapacityReservationInstanceShapesSortByEnum, bool) {
	enum, ok := mappingListComputeCapacityReservationInstanceShapesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListComputeCapacityReservationInstanceShapesSortOrderEnum Enum with underlying type: string
type ListComputeCapacityReservationInstanceShapesSortOrderEnum string

// Set of constants representing the allowable values for ListComputeCapacityReservationInstanceShapesSortOrderEnum
const (
	ListComputeCapacityReservationInstanceShapesSortOrderAsc  ListComputeCapacityReservationInstanceShapesSortOrderEnum = "ASC"
	ListComputeCapacityReservationInstanceShapesSortOrderDesc ListComputeCapacityReservationInstanceShapesSortOrderEnum = "DESC"
)

var mappingListComputeCapacityReservationInstanceShapesSortOrderEnum = map[string]ListComputeCapacityReservationInstanceShapesSortOrderEnum{
	"ASC":  ListComputeCapacityReservationInstanceShapesSortOrderAsc,
	"DESC": ListComputeCapacityReservationInstanceShapesSortOrderDesc,
}

var mappingListComputeCapacityReservationInstanceShapesSortOrderEnumLowerCase = map[string]ListComputeCapacityReservationInstanceShapesSortOrderEnum{
	"asc":  ListComputeCapacityReservationInstanceShapesSortOrderAsc,
	"desc": ListComputeCapacityReservationInstanceShapesSortOrderDesc,
}

// GetListComputeCapacityReservationInstanceShapesSortOrderEnumValues Enumerates the set of values for ListComputeCapacityReservationInstanceShapesSortOrderEnum
func GetListComputeCapacityReservationInstanceShapesSortOrderEnumValues() []ListComputeCapacityReservationInstanceShapesSortOrderEnum {
	values := make([]ListComputeCapacityReservationInstanceShapesSortOrderEnum, 0)
	for _, v := range mappingListComputeCapacityReservationInstanceShapesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListComputeCapacityReservationInstanceShapesSortOrderEnumStringValues Enumerates the set of values in String for ListComputeCapacityReservationInstanceShapesSortOrderEnum
func GetListComputeCapacityReservationInstanceShapesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListComputeCapacityReservationInstanceShapesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListComputeCapacityReservationInstanceShapesSortOrderEnum(val string) (ListComputeCapacityReservationInstanceShapesSortOrderEnum, bool) {
	enum, ok := mappingListComputeCapacityReservationInstanceShapesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
