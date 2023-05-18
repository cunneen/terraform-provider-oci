// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// Use Object Storage and Archive Storage APIs to manage buckets, objects, and related resources.
// For more information, see Overview of Object Storage (https://docs.cloud.oracle.com/Content/Object/Concepts/objectstorageoverview.htm) and
// Overview of Archive Storage (https://docs.cloud.oracle.com/Content/Archive/Concepts/archivestorageoverview.htm).
//

package objectstorage

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"

	"regexp"
)

//ObjectStorageClient a client for ObjectStorage
type ObjectStorageClient struct {
	common.BaseClient
	config                   *common.ConfigurationProvider
	requiredParamsInEndpoint map[string][]common.TemplateParamForPerRealmEndpoint
}

// NewObjectStorageClientWithConfigurationProvider Creates a new default ObjectStorage client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewObjectStorageClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ObjectStorageClient, err error) {
	if enabled := common.CheckForEnabledServices("objectstorage"); !enabled {
		return client, fmt.Errorf("the Alloy configuration disabled this service, this behavior is controlled by OciSdkEnabledServicesMap variables. Please check if your local alloy_config file configured the service you're targeting or contact the cloud provider on the availability of this service")
	}
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newObjectStorageClientFromBaseClient(baseClient, provider)
}

// NewObjectStorageClientWithOboToken Creates a new default ObjectStorage client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewObjectStorageClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ObjectStorageClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newObjectStorageClientFromBaseClient(baseClient, configProvider)
}

func newObjectStorageClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ObjectStorageClient, err error) {
	// ObjectStorage service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("ObjectStorage"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = ObjectStorageClient{BaseClient: baseClient}
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ObjectStorageClient) SetRegion(region string) {
	client.Host, _ = common.StringToRegion(region).EndpointForTemplateDottedRegion("objectstorage", client.getEndpointTemplatePerRealm(region), "objectstorage")
	client.parseEndpointTemplatePerRealm()
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ObjectStorageClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.SetRegion(region)
	if client.Host == "" {
		return fmt.Errorf("Invalid region or Host. Endpoint cannot be constructed without endpointServiceName or serviceEndpointTemplate for a dotted region")
	}
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *ObjectStorageClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// getEndpointTemplatePerRealm returns the endpoint template for the given region, if not found, returns the default endpoint template
func (client *ObjectStorageClient) getEndpointTemplatePerRealm(region string) string {
	if client.IsOciRealmSpecificServiceEndpointTemplateEnabled() {
		realm, _ := common.StringToRegion(region).RealmID()
		templatePerRealmDict := map[string]string{
			"oc1": "https://{namespaceName+Dot}objectstorage.{region}.oci.customer-oci.com",
		}
		if template, ok := templatePerRealmDict[realm]; ok {
			return template
		}
	}
	return "https://objectstorage.{region}.{secondLevelDomain}"
}

// parseEndpointTemplatePerRealm parses the endpoint template per realm from the service endpoint template
// This function will build a map of template params to their values, this map is used when building the API endpoint
func (client *ObjectStorageClient) parseEndpointTemplatePerRealm() {
	client.requiredParamsInEndpoint = make(map[string][]common.TemplateParamForPerRealmEndpoint)
	templateRegex := regexp.MustCompile(`{.*?}`)
	templateSubRegex := regexp.MustCompile(`{(.+)\+Dot}`)
	templates := templateRegex.FindAllString(client.Host, -1)
	for _, template := range templates {
		templateParam := templateSubRegex.FindStringSubmatch(template)
		if len(templateParam) > 1 {
			client.requiredParamsInEndpoint[templateParam[1]] = append(client.requiredParamsInEndpoint[templateParam[1]], common.TemplateParamForPerRealmEndpoint{
				Template:    templateParam[0],
				EndsWithDot: true,
			})
		} else {
			templateParam := template[1 : len(template)-1]
			client.requiredParamsInEndpoint[templateParam] = append(client.requiredParamsInEndpoint[templateParam], common.TemplateParamForPerRealmEndpoint{
				Template:    template,
				EndsWithDot: false,
			})
		}
	}
}

// SetCustomClientConfiguration sets client with retry and other custom configurations
func (client *ObjectStorageClient) SetCustomClientConfiguration(config common.CustomClientConfiguration) {
	client.Configuration = config
	client.refreshRegion()
}

// refreshRegion will refresh the region of this client, this function will be called after setting the CustomClientConfiguration
func (client *ObjectStorageClient) refreshRegion() {
	configProvider := *client.config
	region, _ := configProvider.Region()
	client.SetRegion(region)
}

// AbortMultipartUpload Aborts an in-progress multipart upload and deletes all parts that have been uploaded.
// A default retry strategy applies to this operation AbortMultipartUpload()
func (client ObjectStorageClient) AbortMultipartUpload(ctx context.Context, request AbortMultipartUploadRequest) (response AbortMultipartUploadResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.abortMultipartUpload, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AbortMultipartUploadResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AbortMultipartUploadResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AbortMultipartUploadResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AbortMultipartUploadResponse")
	}
	return
}

// abortMultipartUpload implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) abortMultipartUpload(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/n/{namespaceName}/b/{bucketName}/u/{objectName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(AbortMultipartUploadRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response AbortMultipartUploadResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/MultipartUpload/AbortMultipartUpload"
		err = common.PostProcessServiceError(err, "ObjectStorage", "AbortMultipartUpload", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// BulkCopyObjects Creates a request to copy a batch of objects within a region or to another region.
// See Object Names (https://docs.cloud.oracle.com/Content/Object/Tasks/managingobjects.htm#namerequirements)
// for object naming requirements.
// A default retry strategy applies to this operation BulkCopyObjects()
func (client ObjectStorageClient) BulkCopyObjects(ctx context.Context, request BulkCopyObjectsRequest) (response BulkCopyObjectsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.bulkCopyObjects, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = BulkCopyObjectsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = BulkCopyObjectsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(BulkCopyObjectsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into BulkCopyObjectsResponse")
	}
	return
}

// bulkCopyObjects implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) bulkCopyObjects(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/actions/bulkCopyObjects", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(BulkCopyObjectsRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response BulkCopyObjectsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/Object/BulkCopyObjects"
		err = common.PostProcessServiceError(err, "ObjectStorage", "BulkCopyObjects", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CancelWorkRequest Cancels a work request.
// A default retry strategy applies to this operation CancelWorkRequest()
func (client ObjectStorageClient) CancelWorkRequest(ctx context.Context, request CancelWorkRequestRequest) (response CancelWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.cancelWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CancelWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CancelWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CancelWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelWorkRequestResponse")
	}
	return
}

// cancelWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) cancelWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/workRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(CancelWorkRequestRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response CancelWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/WorkRequest/CancelWorkRequest"
		err = common.PostProcessServiceError(err, "ObjectStorage", "CancelWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CheckObject Retrieve an object's stored and calculated digests, specifically:
//   - The content MD5;
//   - The SHA-256 of each chunk on each storage server.
// If the MD5 digest calculated from the decrypted object data
// does not match the MD5 digest stored in the object's metadata,
// or any chunk's' SHA-256 does not match the one stored for it,
// this API will signal an error.  The stored and calculated
// digests are always included in the response.
// This internal API allows Object Storage tooling to verify the
// correctness of stored data.  Any tenancy's objects can be
// verified with this API, so it must only protected by BOAT AAA.
// A default retry strategy applies to this operation CheckObject()
func (client ObjectStorageClient) CheckObject(ctx context.Context, request CheckObjectRequest) (response CheckObjectResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.checkObject, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CheckObjectResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CheckObjectResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CheckObjectResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CheckObjectResponse")
	}
	return
}

// checkObject implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) checkObject(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/actions/checkObject/{objectName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(CheckObjectRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response CheckObjectResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/Object/CheckObject"
		err = common.PostProcessServiceError(err, "ObjectStorage", "CheckObject", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CommitMultipartUpload Commits a multipart upload, which involves checking part numbers and entity tags (ETags) of the parts, to create an aggregate object.
// A default retry strategy applies to this operation CommitMultipartUpload()
func (client ObjectStorageClient) CommitMultipartUpload(ctx context.Context, request CommitMultipartUploadRequest) (response CommitMultipartUploadResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.commitMultipartUpload, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CommitMultipartUploadResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CommitMultipartUploadResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CommitMultipartUploadResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CommitMultipartUploadResponse")
	}
	return
}

// commitMultipartUpload implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) commitMultipartUpload(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/u/{objectName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(CommitMultipartUploadRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response CommitMultipartUploadResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/MultipartUpload/CommitMultipartUpload"
		err = common.PostProcessServiceError(err, "ObjectStorage", "CommitMultipartUpload", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CopyObject Creates a request to copy an object within a region or to another region.
// See Object Names (https://docs.cloud.oracle.com/Content/Object/Tasks/managingobjects.htm#namerequirements)
// for object naming requirements.
// A default retry strategy applies to this operation CopyObject()
func (client ObjectStorageClient) CopyObject(ctx context.Context, request CopyObjectRequest) (response CopyObjectResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.copyObject, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CopyObjectResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CopyObjectResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CopyObjectResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CopyObjectResponse")
	}
	return
}

// copyObject implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) copyObject(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/actions/copyObject", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(CopyObjectRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response CopyObjectResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/Object/CopyObject"
		err = common.PostProcessServiceError(err, "ObjectStorage", "CopyObject", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CopyPart Copy part of an existing object to a destination object
// A default retry strategy applies to this operation CopyPart()
func (client ObjectStorageClient) CopyPart(ctx context.Context, request CopyPartRequest) (response CopyPartResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.copyPart, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CopyPartResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CopyPartResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CopyPartResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CopyPartResponse")
	}
	return
}

// copyPart implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) copyPart(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/n/{namespaceName}/b/{bucketName}/cp/{objectName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(CopyPartRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response CopyPartResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/MultipartUpload/CopyPart"
		err = common.PostProcessServiceError(err, "ObjectStorage", "CopyPart", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateAcl Creates an Access-Control List. The ACL can be included in an ACL Group to restrict network access to certain
// buckets and objects.
// A default retry strategy applies to this operation CreateAcl()
func (client ObjectStorageClient) CreateAcl(ctx context.Context, request CreateAclRequest) (response CreateAclResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createAcl, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateAclResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateAclResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAclResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAclResponse")
	}
	return
}

// createAcl implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) createAcl(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/n/{namespaceName}/security/acl", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(CreateAclRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response CreateAclResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/Acl/CreateAcl"
		err = common.PostProcessServiceError(err, "ObjectStorage", "CreateAcl", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateAclGroup Creates an ACL Group. These are used to restrict network access to certain buckets and objects.
// A default retry strategy applies to this operation CreateAclGroup()
func (client ObjectStorageClient) CreateAclGroup(ctx context.Context, request CreateAclGroupRequest) (response CreateAclGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createAclGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateAclGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateAclGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAclGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAclGroupResponse")
	}
	return
}

// createAclGroup implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) createAclGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/n/{namespaceName}/security/aclgroup", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(CreateAclGroupRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response CreateAclGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/AclGroup/CreateAclGroup"
		err = common.PostProcessServiceError(err, "ObjectStorage", "CreateAclGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateBucket Creates a bucket in the given namespace with a bucket name and optional user-defined metadata. Avoid entering
// confidential information in bucket names.
// A default retry strategy applies to this operation CreateBucket()
func (client ObjectStorageClient) CreateBucket(ctx context.Context, request CreateBucketRequest) (response CreateBucketResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.createBucket, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateBucketResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateBucketResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateBucketResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateBucketResponse")
	}
	return
}

// createBucket implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) createBucket(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/n/{namespaceName}/b", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(CreateBucketRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response CreateBucketResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/Bucket/CreateBucket"
		err = common.PostProcessServiceError(err, "ObjectStorage", "CreateBucket", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateMultipartUpload Starts a new multipart upload to a specific object in the given bucket in the given namespace.
// See Object Names (https://docs.cloud.oracle.com/Content/Object/Tasks/managingobjects.htm#namerequirements)
// for object naming requirements.
// A default retry strategy applies to this operation CreateMultipartUpload()
func (client ObjectStorageClient) CreateMultipartUpload(ctx context.Context, request CreateMultipartUploadRequest) (response CreateMultipartUploadResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.createMultipartUpload, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateMultipartUploadResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateMultipartUploadResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateMultipartUploadResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateMultipartUploadResponse")
	}
	return
}

// createMultipartUpload implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) createMultipartUpload(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/u", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(CreateMultipartUploadRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response CreateMultipartUploadResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/MultipartUpload/CreateMultipartUpload"
		err = common.PostProcessServiceError(err, "ObjectStorage", "CreateMultipartUpload", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreatePreauthenticatedRequest Creates a pre-authenticated request specific to the bucket.
// A default retry strategy applies to this operation CreatePreauthenticatedRequest()
func (client ObjectStorageClient) CreatePreauthenticatedRequest(ctx context.Context, request CreatePreauthenticatedRequestRequest) (response CreatePreauthenticatedRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.createPreauthenticatedRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreatePreauthenticatedRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreatePreauthenticatedRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreatePreauthenticatedRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreatePreauthenticatedRequestResponse")
	}
	return
}

// createPreauthenticatedRequest implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) createPreauthenticatedRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/p", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(CreatePreauthenticatedRequestRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response CreatePreauthenticatedRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/PreauthenticatedRequest/CreatePreauthenticatedRequest"
		err = common.PostProcessServiceError(err, "ObjectStorage", "CreatePreauthenticatedRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreatePrivateEndpoint Create a PrivateEndpoint.
// A default retry strategy applies to this operation CreatePrivateEndpoint()
func (client ObjectStorageClient) CreatePrivateEndpoint(ctx context.Context, request CreatePrivateEndpointRequest) (response CreatePrivateEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createPrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreatePrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreatePrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreatePrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreatePrivateEndpointResponse")
	}
	return
}

// createPrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) createPrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/n/{namespaceName}/pe", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(CreatePrivateEndpointRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response CreatePrivateEndpointResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/PrivateEndpoint/CreatePrivateEndpoint"
		err = common.PostProcessServiceError(err, "ObjectStorage", "CreatePrivateEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateReplicationPolicy Creates a replication policy for the specified bucket.
// A default retry strategy applies to this operation CreateReplicationPolicy()
func (client ObjectStorageClient) CreateReplicationPolicy(ctx context.Context, request CreateReplicationPolicyRequest) (response CreateReplicationPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.createReplicationPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateReplicationPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateReplicationPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateReplicationPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateReplicationPolicyResponse")
	}
	return
}

// createReplicationPolicy implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) createReplicationPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/replicationPolicies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(CreateReplicationPolicyRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response CreateReplicationPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/Replication/CreateReplicationPolicy"
		err = common.PostProcessServiceError(err, "ObjectStorage", "CreateReplicationPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateRetentionRule Creates a new retention rule in the specified bucket. The new rule will take effect typically within 30 seconds.
// Note that a maximum of 100 rules are supported on a bucket.
// A default retry strategy applies to this operation CreateRetentionRule()
func (client ObjectStorageClient) CreateRetentionRule(ctx context.Context, request CreateRetentionRuleRequest) (response CreateRetentionRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.createRetentionRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateRetentionRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateRetentionRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateRetentionRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateRetentionRuleResponse")
	}
	return
}

// createRetentionRule implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) createRetentionRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/retentionRules", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(CreateRetentionRuleRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response CreateRetentionRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/RetentionRule/CreateRetentionRule"
		err = common.PostProcessServiceError(err, "ObjectStorage", "CreateRetentionRule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteAcl Deletes the specified ACL. The ACL must not be in any existing ACL Group.
// A default retry strategy applies to this operation DeleteAcl()
func (client ObjectStorageClient) DeleteAcl(ctx context.Context, request DeleteAclRequest) (response DeleteAclResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteAcl, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteAclResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteAclResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteAclResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteAclResponse")
	}
	return
}

// deleteAcl implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) deleteAcl(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/n/{namespaceName}/security/acl/{aclId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(DeleteAclRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response DeleteAclResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/Acl/DeleteAcl"
		err = common.PostProcessServiceError(err, "ObjectStorage", "DeleteAcl", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteAclGroup Deletes the specified ACL Group.
// A default retry strategy applies to this operation DeleteAclGroup()
func (client ObjectStorageClient) DeleteAclGroup(ctx context.Context, request DeleteAclGroupRequest) (response DeleteAclGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteAclGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteAclGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteAclGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteAclGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteAclGroupResponse")
	}
	return
}

// deleteAclGroup implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) deleteAclGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/n/{namespaceName}/security/aclgroup/{aclGroupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(DeleteAclGroupRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response DeleteAclGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/AclGroup/DeleteAclGroup"
		err = common.PostProcessServiceError(err, "ObjectStorage", "DeleteAclGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteBucket Deletes a bucket if the bucket is already empty. If the bucket is not empty, use
// DeleteObject first. In addition,
// you cannot delete a bucket that has a multipart upload in progress or a pre-authenticated
// request associated with that bucket.
// A default retry strategy applies to this operation DeleteBucket()
func (client ObjectStorageClient) DeleteBucket(ctx context.Context, request DeleteBucketRequest) (response DeleteBucketResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteBucket, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteBucketResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteBucketResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteBucketResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteBucketResponse")
	}
	return
}

// deleteBucket implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) deleteBucket(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/n/{namespaceName}/b/{bucketName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(DeleteBucketRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response DeleteBucketResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/Bucket/DeleteBucket"
		err = common.PostProcessServiceError(err, "ObjectStorage", "DeleteBucket", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteObject Deletes an object.
// A default retry strategy applies to this operation DeleteObject()
func (client ObjectStorageClient) DeleteObject(ctx context.Context, request DeleteObjectRequest) (response DeleteObjectResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteObject, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteObjectResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteObjectResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteObjectResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteObjectResponse")
	}
	return
}

// deleteObject implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) deleteObject(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/n/{namespaceName}/b/{bucketName}/o/{objectName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(DeleteObjectRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response DeleteObjectResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/Object/DeleteObject"
		err = common.PostProcessServiceError(err, "ObjectStorage", "DeleteObject", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteObjectLifecyclePolicy Deletes the object lifecycle policy for the bucket.
// A default retry strategy applies to this operation DeleteObjectLifecyclePolicy()
func (client ObjectStorageClient) DeleteObjectLifecyclePolicy(ctx context.Context, request DeleteObjectLifecyclePolicyRequest) (response DeleteObjectLifecyclePolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteObjectLifecyclePolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteObjectLifecyclePolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteObjectLifecyclePolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteObjectLifecyclePolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteObjectLifecyclePolicyResponse")
	}
	return
}

// deleteObjectLifecyclePolicy implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) deleteObjectLifecyclePolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/n/{namespaceName}/b/{bucketName}/l", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(DeleteObjectLifecyclePolicyRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response DeleteObjectLifecyclePolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/ObjectLifecyclePolicy/DeleteObjectLifecyclePolicy"
		err = common.PostProcessServiceError(err, "ObjectStorage", "DeleteObjectLifecyclePolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeletePreauthenticatedRequest Deletes the pre-authenticated request for the bucket.
// A default retry strategy applies to this operation DeletePreauthenticatedRequest()
func (client ObjectStorageClient) DeletePreauthenticatedRequest(ctx context.Context, request DeletePreauthenticatedRequestRequest) (response DeletePreauthenticatedRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deletePreauthenticatedRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeletePreauthenticatedRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeletePreauthenticatedRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeletePreauthenticatedRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeletePreauthenticatedRequestResponse")
	}
	return
}

// deletePreauthenticatedRequest implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) deletePreauthenticatedRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/n/{namespaceName}/b/{bucketName}/p/{parId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(DeletePreauthenticatedRequestRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response DeletePreauthenticatedRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/PreauthenticatedRequest/DeletePreauthenticatedRequest"
		err = common.PostProcessServiceError(err, "ObjectStorage", "DeletePreauthenticatedRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeletePrivateEndpoint Deletes a Private Endpoint if it exists in the given namespace.
// A default retry strategy applies to this operation DeletePrivateEndpoint()
func (client ObjectStorageClient) DeletePrivateEndpoint(ctx context.Context, request DeletePrivateEndpointRequest) (response DeletePrivateEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deletePrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeletePrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeletePrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeletePrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeletePrivateEndpointResponse")
	}
	return
}

// deletePrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) deletePrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/n/{namespaceName}/pe/{peName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(DeletePrivateEndpointRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response DeletePrivateEndpointResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/PrivateEndpoint/DeletePrivateEndpoint"
		err = common.PostProcessServiceError(err, "ObjectStorage", "DeletePrivateEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteReplicationPolicy Deletes the replication policy associated with the source bucket.
// A default retry strategy applies to this operation DeleteReplicationPolicy()
func (client ObjectStorageClient) DeleteReplicationPolicy(ctx context.Context, request DeleteReplicationPolicyRequest) (response DeleteReplicationPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteReplicationPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteReplicationPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteReplicationPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteReplicationPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteReplicationPolicyResponse")
	}
	return
}

// deleteReplicationPolicy implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) deleteReplicationPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/n/{namespaceName}/b/{bucketName}/replicationPolicies/{replicationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(DeleteReplicationPolicyRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response DeleteReplicationPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/Replication/DeleteReplicationPolicy"
		err = common.PostProcessServiceError(err, "ObjectStorage", "DeleteReplicationPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteRetentionRule Deletes the specified rule. The deletion takes effect typically within 30 seconds.
// A default retry strategy applies to this operation DeleteRetentionRule()
func (client ObjectStorageClient) DeleteRetentionRule(ctx context.Context, request DeleteRetentionRuleRequest) (response DeleteRetentionRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteRetentionRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteRetentionRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteRetentionRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteRetentionRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteRetentionRuleResponse")
	}
	return
}

// deleteRetentionRule implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) deleteRetentionRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/n/{namespaceName}/b/{bucketName}/retentionRules/{retentionRuleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(DeleteRetentionRuleRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response DeleteRetentionRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/RetentionRule/DeleteRetentionRule"
		err = common.PostProcessServiceError(err, "ObjectStorage", "DeleteRetentionRule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAcl Get the specified ACL.
// A default retry strategy applies to this operation GetAcl()
func (client ObjectStorageClient) GetAcl(ctx context.Context, request GetAclRequest) (response GetAclResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAcl, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAclResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAclResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAclResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAclResponse")
	}
	return
}

// getAcl implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) getAcl(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/n/{namespaceName}/security/acl/{aclId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(GetAclRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response GetAclResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/Acl/GetAcl"
		err = common.PostProcessServiceError(err, "ObjectStorage", "GetAcl", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAclGroup Get the specified ACL Group.
// A default retry strategy applies to this operation GetAclGroup()
func (client ObjectStorageClient) GetAclGroup(ctx context.Context, request GetAclGroupRequest) (response GetAclGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAclGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAclGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAclGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAclGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAclGroupResponse")
	}
	return
}

// getAclGroup implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) getAclGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/n/{namespaceName}/security/aclgroup/{aclGroupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(GetAclGroupRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response GetAclGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/AclGroup/GetAclGroup"
		err = common.PostProcessServiceError(err, "ObjectStorage", "GetAclGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetBucket Gets the current representation of the given bucket in the given Object Storage namespace.
// A default retry strategy applies to this operation GetBucket()
func (client ObjectStorageClient) GetBucket(ctx context.Context, request GetBucketRequest) (response GetBucketResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getBucket, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetBucketResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetBucketResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetBucketResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetBucketResponse")
	}
	return
}

// getBucket implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) getBucket(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/n/{namespaceName}/b/{bucketName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(GetBucketRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response GetBucketResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/Bucket/GetBucket"
		err = common.PostProcessServiceError(err, "ObjectStorage", "GetBucket", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetBucketOptions Lists the various options associated with the bucket. The options are returned as a JSON-formatted object.
// This API is for internal-use only.
// A default retry strategy applies to this operation GetBucketOptions()
func (client ObjectStorageClient) GetBucketOptions(ctx context.Context, request GetBucketOptionsRequest) (response GetBucketOptionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getBucketOptions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetBucketOptionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetBucketOptionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetBucketOptionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetBucketOptionsResponse")
	}
	return
}

// getBucketOptions implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) getBucketOptions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/actions/options", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(GetBucketOptionsRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response GetBucketOptionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/Bucket/GetBucketOptions"
		err = common.PostProcessServiceError(err, "ObjectStorage", "GetBucketOptions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetNamespace Each Oracle Cloud Infrastructure tenant is assigned one unique and uneditable Object Storage namespace. The namespace
// is a system-generated string assigned during account creation. For some older tenancies, the namespace string may be
// the tenancy name in all lower-case letters. You cannot edit a namespace.
// GetNamespace returns the name of the Object Storage namespace for the user making the request.
// If an optional compartmentId query parameter is provided, GetNamespace returns the namespace name of the corresponding
// tenancy, provided the user has access to it.
// A default retry strategy applies to this operation GetNamespace()
func (client ObjectStorageClient) GetNamespace(ctx context.Context, request GetNamespaceRequest) (response GetNamespaceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getNamespace, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetNamespaceResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetNamespaceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetNamespaceResponse")
	}
	return
}

// getNamespace implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) getNamespace(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/n", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(GetNamespaceRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response GetNamespaceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/Namespace/GetNamespace"
		err = common.PostProcessServiceError(err, "ObjectStorage", "GetNamespace", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetNamespaceMetadata Gets the metadata for the Object Storage namespace, which contains defaultS3CompartmentId and
// defaultSwiftCompartmentId.
// Any user with the OBJECTSTORAGE_NAMESPACE_READ permission will be able to see the current metadata. If you are
// not authorized, talk to an administrator. If you are an administrator who needs to write policies
// to give users access, see
// Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
// A default retry strategy applies to this operation GetNamespaceMetadata()
func (client ObjectStorageClient) GetNamespaceMetadata(ctx context.Context, request GetNamespaceMetadataRequest) (response GetNamespaceMetadataResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getNamespaceMetadata, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetNamespaceMetadataResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetNamespaceMetadataResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetNamespaceMetadataResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetNamespaceMetadataResponse")
	}
	return
}

// getNamespaceMetadata implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) getNamespaceMetadata(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/n/{namespaceName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(GetNamespaceMetadataRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response GetNamespaceMetadataResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/Namespace/GetNamespaceMetadata"
		err = common.PostProcessServiceError(err, "ObjectStorage", "GetNamespaceMetadata", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetObject Gets the metadata and body of an object.
// A default retry strategy applies to this operation GetObject()
func (client ObjectStorageClient) GetObject(ctx context.Context, request GetObjectRequest) (response GetObjectResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getObject, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetObjectResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetObjectResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetObjectResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetObjectResponse")
	}
	return
}

// getObject implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) getObject(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/o/{objectName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(GetObjectRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response GetObjectResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/Object/GetObject"
		err = common.PostProcessServiceError(err, "ObjectStorage", "GetObject", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetObjectLifecyclePolicy Gets the object lifecycle policy for the bucket.
// A default retry strategy applies to this operation GetObjectLifecyclePolicy()
func (client ObjectStorageClient) GetObjectLifecyclePolicy(ctx context.Context, request GetObjectLifecyclePolicyRequest) (response GetObjectLifecyclePolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getObjectLifecyclePolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetObjectLifecyclePolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetObjectLifecyclePolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetObjectLifecyclePolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetObjectLifecyclePolicyResponse")
	}
	return
}

// getObjectLifecyclePolicy implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) getObjectLifecyclePolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/l", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(GetObjectLifecyclePolicyRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response GetObjectLifecyclePolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/ObjectLifecyclePolicy/GetObjectLifecyclePolicy"
		err = common.PostProcessServiceError(err, "ObjectStorage", "GetObjectLifecyclePolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetPreauthenticatedRequest Gets the pre-authenticated request for the bucket.
// A default retry strategy applies to this operation GetPreauthenticatedRequest()
func (client ObjectStorageClient) GetPreauthenticatedRequest(ctx context.Context, request GetPreauthenticatedRequestRequest) (response GetPreauthenticatedRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPreauthenticatedRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPreauthenticatedRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPreauthenticatedRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPreauthenticatedRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPreauthenticatedRequestResponse")
	}
	return
}

// getPreauthenticatedRequest implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) getPreauthenticatedRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/p/{parId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(GetPreauthenticatedRequestRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response GetPreauthenticatedRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/PreauthenticatedRequest/GetPreauthenticatedRequest"
		err = common.PostProcessServiceError(err, "ObjectStorage", "GetPreauthenticatedRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetPrefixRenameStatus Get the status of prefix rename for the given work request ID.
// A default retry strategy applies to this operation GetPrefixRenameStatus()
func (client ObjectStorageClient) GetPrefixRenameStatus(ctx context.Context, request GetPrefixRenameStatusRequest) (response GetPrefixRenameStatusResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPrefixRenameStatus, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPrefixRenameStatusResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPrefixRenameStatusResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPrefixRenameStatusResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPrefixRenameStatusResponse")
	}
	return
}

// getPrefixRenameStatus implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) getPrefixRenameStatus(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/actions/prefixRename/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(GetPrefixRenameStatusRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response GetPrefixRenameStatusResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/Object/GetPrefixRenameStatus"
		err = common.PostProcessServiceError(err, "ObjectStorage", "GetPrefixRenameStatus", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetPrivateEndpoint Gets the current representation of the given Private Endpoint in the given Object Storage namespace.
// A default retry strategy applies to this operation GetPrivateEndpoint()
func (client ObjectStorageClient) GetPrivateEndpoint(ctx context.Context, request GetPrivateEndpointRequest) (response GetPrivateEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPrivateEndpointResponse")
	}
	return
}

// getPrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) getPrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/n/{namespaceName}/pe/{peName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(GetPrivateEndpointRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response GetPrivateEndpointResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/PrivateEndpoint/GetPrivateEndpoint"
		err = common.PostProcessServiceError(err, "ObjectStorage", "GetPrivateEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetReplicationPolicy Get the replication policy.
// A default retry strategy applies to this operation GetReplicationPolicy()
func (client ObjectStorageClient) GetReplicationPolicy(ctx context.Context, request GetReplicationPolicyRequest) (response GetReplicationPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getReplicationPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetReplicationPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetReplicationPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetReplicationPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetReplicationPolicyResponse")
	}
	return
}

// getReplicationPolicy implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) getReplicationPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/replicationPolicies/{replicationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(GetReplicationPolicyRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response GetReplicationPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/Replication/GetReplicationPolicy"
		err = common.PostProcessServiceError(err, "ObjectStorage", "GetReplicationPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetRetentionRule Get the specified retention rule.
// A default retry strategy applies to this operation GetRetentionRule()
func (client ObjectStorageClient) GetRetentionRule(ctx context.Context, request GetRetentionRuleRequest) (response GetRetentionRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getRetentionRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetRetentionRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetRetentionRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetRetentionRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetRetentionRuleResponse")
	}
	return
}

// getRetentionRule implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) getRetentionRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/retentionRules/{retentionRuleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(GetRetentionRuleRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response GetRetentionRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/RetentionRule/GetRetentionRule"
		err = common.PostProcessServiceError(err, "ObjectStorage", "GetRetentionRule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Gets the status of the work request for the given ID.
// A default retry strategy applies to this operation GetWorkRequest()
func (client ObjectStorageClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetWorkRequestResponse")
	}
	return
}

// getWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(GetWorkRequestRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response GetWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/WorkRequest/GetWorkRequest"
		err = common.PostProcessServiceError(err, "ObjectStorage", "GetWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// HeadBucket Efficiently checks to see if a bucket exists and gets the current entity tag (ETag) for the bucket.
// A default retry strategy applies to this operation HeadBucket()
func (client ObjectStorageClient) HeadBucket(ctx context.Context, request HeadBucketRequest) (response HeadBucketResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.headBucket, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = HeadBucketResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = HeadBucketResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(HeadBucketResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into HeadBucketResponse")
	}
	return
}

// headBucket implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) headBucket(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodHead, "/n/{namespaceName}/b/{bucketName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(HeadBucketRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response HeadBucketResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/Bucket/HeadBucket"
		err = common.PostProcessServiceError(err, "ObjectStorage", "HeadBucket", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// HeadObject Gets the user-defined metadata and entity tag (ETag) for an object.
// A default retry strategy applies to this operation HeadObject()
func (client ObjectStorageClient) HeadObject(ctx context.Context, request HeadObjectRequest) (response HeadObjectResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.headObject, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = HeadObjectResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = HeadObjectResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(HeadObjectResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into HeadObjectResponse")
	}
	return
}

// headObject implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) headObject(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodHead, "/n/{namespaceName}/b/{bucketName}/o/{objectName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(HeadObjectRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response HeadObjectResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/Object/HeadObject"
		err = common.PostProcessServiceError(err, "ObjectStorage", "HeadObject", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAclGroups List the ACL Groups in the namespace residing in the given compartment.
// A default retry strategy applies to this operation ListAclGroups()
func (client ObjectStorageClient) ListAclGroups(ctx context.Context, request ListAclGroupsRequest) (response ListAclGroupsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAclGroups, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAclGroupsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAclGroupsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAclGroupsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAclGroupsResponse")
	}
	return
}

// listAclGroups implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) listAclGroups(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/n/{namespaceName}/security/aclgroup", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(ListAclGroupsRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response ListAclGroupsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/AclGroup/ListAclGroups"
		err = common.PostProcessServiceError(err, "ObjectStorage", "ListAclGroups", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAcls List the ACLs in the namespace.
// A default retry strategy applies to this operation ListAcls()
func (client ObjectStorageClient) ListAcls(ctx context.Context, request ListAclsRequest) (response ListAclsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAcls, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAclsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAclsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAclsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAclsResponse")
	}
	return
}

// listAcls implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) listAcls(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/n/{namespaceName}/security/acl", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(ListAclsRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response ListAclsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/Acl/ListAcls"
		err = common.PostProcessServiceError(err, "ObjectStorage", "ListAcls", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListBuckets Gets a list of all BucketSummary items in a compartment. A BucketSummary contains only summary fields for the bucket
// and does not contain fields like the user-defined metadata.
// ListBuckets returns a BucketSummary containing at most 1000 buckets. To paginate through more buckets, use the returned
// `opc-next-page` value with the `page` request parameter.
// To use this and other API operations, you must be authorized in an IAM policy. If you are not authorized,
// talk to an administrator. If you are an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
// A default retry strategy applies to this operation ListBuckets()
func (client ObjectStorageClient) ListBuckets(ctx context.Context, request ListBucketsRequest) (response ListBucketsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listBuckets, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListBucketsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListBucketsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListBucketsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListBucketsResponse")
	}
	return
}

// listBuckets implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) listBuckets(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/n/{namespaceName}/b", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(ListBucketsRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response ListBucketsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/Bucket/ListBuckets"
		err = common.PostProcessServiceError(err, "ObjectStorage", "ListBuckets", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMultipartUploadParts Lists the parts of an in-progress multipart upload.
// A default retry strategy applies to this operation ListMultipartUploadParts()
func (client ObjectStorageClient) ListMultipartUploadParts(ctx context.Context, request ListMultipartUploadPartsRequest) (response ListMultipartUploadPartsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMultipartUploadParts, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMultipartUploadPartsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMultipartUploadPartsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMultipartUploadPartsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMultipartUploadPartsResponse")
	}
	return
}

// listMultipartUploadParts implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) listMultipartUploadParts(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/u/{objectName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(ListMultipartUploadPartsRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response ListMultipartUploadPartsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/MultipartUpload/ListMultipartUploadParts"
		err = common.PostProcessServiceError(err, "ObjectStorage", "ListMultipartUploadParts", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMultipartUploads Lists all of the in-progress multipart uploads for the given bucket in the given Object Storage namespace.
// A default retry strategy applies to this operation ListMultipartUploads()
func (client ObjectStorageClient) ListMultipartUploads(ctx context.Context, request ListMultipartUploadsRequest) (response ListMultipartUploadsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMultipartUploads, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMultipartUploadsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMultipartUploadsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMultipartUploadsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMultipartUploadsResponse")
	}
	return
}

// listMultipartUploads implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) listMultipartUploads(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/u", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(ListMultipartUploadsRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response ListMultipartUploadsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/MultipartUpload/ListMultipartUploads"
		err = common.PostProcessServiceError(err, "ObjectStorage", "ListMultipartUploads", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListObjectVersions Lists the object versions in a bucket.
// ListObjectVersions returns an ObjectVersionCollection containing at most 1000 object versions. To paginate through
// more object versions, use the returned `opc-next-page` value with the `page` request parameter.
// To use this and other API operations, you must be authorized in an IAM policy. If you are not authorized,
// talk to an administrator. If you are an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
// A default retry strategy applies to this operation ListObjectVersions()
func (client ObjectStorageClient) ListObjectVersions(ctx context.Context, request ListObjectVersionsRequest) (response ListObjectVersionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listObjectVersions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListObjectVersionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListObjectVersionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListObjectVersionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListObjectVersionsResponse")
	}
	return
}

// listObjectVersions implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) listObjectVersions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/objectversions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(ListObjectVersionsRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response ListObjectVersionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/Object/ListObjectVersions"
		err = common.PostProcessServiceError(err, "ObjectStorage", "ListObjectVersions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListObjects Lists the objects in a bucket. By default, ListObjects returns object names only. See the `fields`
// parameter for other fields that you can optionally include in ListObjects response.
// ListObjects returns at most 1000 objects. To paginate through more objects, use the returned 'nextStartWith'
// value with the 'start' parameter. To filter which objects ListObjects returns, use the 'start' and 'end'
// parameters.
// To use this and other API operations, you must be authorized in an IAM policy. If you are not authorized,
// talk to an administrator. If you are an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
// A default retry strategy applies to this operation ListObjects()
func (client ObjectStorageClient) ListObjects(ctx context.Context, request ListObjectsRequest) (response ListObjectsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listObjects, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListObjectsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListObjectsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListObjectsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListObjectsResponse")
	}
	return
}

// listObjects implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) listObjects(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/o", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(ListObjectsRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response ListObjectsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/Object/ListObjects"
		err = common.PostProcessServiceError(err, "ObjectStorage", "ListObjects", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPreauthenticatedRequests Lists pre-authenticated requests for the bucket.
// A default retry strategy applies to this operation ListPreauthenticatedRequests()
func (client ObjectStorageClient) ListPreauthenticatedRequests(ctx context.Context, request ListPreauthenticatedRequestsRequest) (response ListPreauthenticatedRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPreauthenticatedRequests, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPreauthenticatedRequestsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPreauthenticatedRequestsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPreauthenticatedRequestsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPreauthenticatedRequestsResponse")
	}
	return
}

// listPreauthenticatedRequests implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) listPreauthenticatedRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/p", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(ListPreauthenticatedRequestsRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response ListPreauthenticatedRequestsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/PreauthenticatedRequest/ListPreauthenticatedRequests"
		err = common.PostProcessServiceError(err, "ObjectStorage", "ListPreauthenticatedRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPrivateEndpoints Gets a list of all PrivateEndpointSummary in a compartment associated with a namespace.
// To use this and other API operations, you must be authorized in an IAM policy. If you are not authorized,
// talk to an administrator. If you are an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
// A default retry strategy applies to this operation ListPrivateEndpoints()
func (client ObjectStorageClient) ListPrivateEndpoints(ctx context.Context, request ListPrivateEndpointsRequest) (response ListPrivateEndpointsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPrivateEndpoints, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPrivateEndpointsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPrivateEndpointsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPrivateEndpointsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPrivateEndpointsResponse")
	}
	return
}

// listPrivateEndpoints implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) listPrivateEndpoints(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/n/{namespaceName}/pe", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(ListPrivateEndpointsRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response ListPrivateEndpointsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/PrivateEndpointSummary/ListPrivateEndpoints"
		err = common.PostProcessServiceError(err, "ObjectStorage", "ListPrivateEndpoints", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListReplicationPolicies List the replication policies associated with a bucket.
// A default retry strategy applies to this operation ListReplicationPolicies()
func (client ObjectStorageClient) ListReplicationPolicies(ctx context.Context, request ListReplicationPoliciesRequest) (response ListReplicationPoliciesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listReplicationPolicies, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListReplicationPoliciesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListReplicationPoliciesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListReplicationPoliciesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListReplicationPoliciesResponse")
	}
	return
}

// listReplicationPolicies implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) listReplicationPolicies(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/replicationPolicies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(ListReplicationPoliciesRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response ListReplicationPoliciesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/Replication/ListReplicationPolicies"
		err = common.PostProcessServiceError(err, "ObjectStorage", "ListReplicationPolicies", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListReplicationSources List the replication sources of a destination bucket.
// A default retry strategy applies to this operation ListReplicationSources()
func (client ObjectStorageClient) ListReplicationSources(ctx context.Context, request ListReplicationSourcesRequest) (response ListReplicationSourcesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listReplicationSources, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListReplicationSourcesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListReplicationSourcesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListReplicationSourcesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListReplicationSourcesResponse")
	}
	return
}

// listReplicationSources implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) listReplicationSources(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/replicationSources", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(ListReplicationSourcesRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response ListReplicationSourcesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/Replication/ListReplicationSources"
		err = common.PostProcessServiceError(err, "ObjectStorage", "ListReplicationSources", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListRetentionRules List the retention rules for a bucket. The retention rules are sorted based on creation time,
// with the most recently created retention rule returned first.
// A default retry strategy applies to this operation ListRetentionRules()
func (client ObjectStorageClient) ListRetentionRules(ctx context.Context, request ListRetentionRulesRequest) (response ListRetentionRulesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listRetentionRules, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListRetentionRulesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListRetentionRulesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListRetentionRulesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListRetentionRulesResponse")
	}
	return
}

// listRetentionRules implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) listRetentionRules(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/retentionRules", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(ListRetentionRulesRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response ListRetentionRulesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/RetentionRule/ListRetentionRules"
		err = common.PostProcessServiceError(err, "ObjectStorage", "ListRetentionRules", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Lists the errors of the work request with the given ID.
// A default retry strategy applies to this operation ListWorkRequestErrors()
func (client ObjectStorageClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequestErrors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestErrorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestErrorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestErrorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestErrorsResponse")
	}
	return
}

// listWorkRequestErrors implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/errors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(ListWorkRequestErrorsRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response ListWorkRequestErrorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/WorkRequestError/ListWorkRequestErrors"
		err = common.PostProcessServiceError(err, "ObjectStorage", "ListWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Lists the logs of the work request with the given ID.
// A default retry strategy applies to this operation ListWorkRequestLogs()
func (client ObjectStorageClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequestLogs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestLogsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestLogsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestLogsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestLogsResponse")
	}
	return
}

// listWorkRequestLogs implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/logs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(ListWorkRequestLogsRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response ListWorkRequestLogsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/WorkRequestLogEntry/ListWorkRequestLogs"
		err = common.PostProcessServiceError(err, "ObjectStorage", "ListWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests Lists the work requests in a compartment.
// A default retry strategy applies to this operation ListWorkRequests()
func (client ObjectStorageClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequests, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestsResponse")
	}
	return
}

// listWorkRequests implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(ListWorkRequestsRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response ListWorkRequestsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/WorkRequest/ListWorkRequests"
		err = common.PostProcessServiceError(err, "ObjectStorage", "ListWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// MakeBucketWritable Stops replication to the destination bucket and removes the replication policy. When the replication
// policy was created, this destination bucket became read-only except for new and changed objects replicated
// automatically from the source bucket. MakeBucketWritable removes the replication policy. This bucket is no
// longer the target for replication and is now writable, allowing users to make changes to bucket contents.
// A default retry strategy applies to this operation MakeBucketWritable()
func (client ObjectStorageClient) MakeBucketWritable(ctx context.Context, request MakeBucketWritableRequest) (response MakeBucketWritableResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.makeBucketWritable, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = MakeBucketWritableResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = MakeBucketWritableResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(MakeBucketWritableResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into MakeBucketWritableResponse")
	}
	return
}

// makeBucketWritable implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) makeBucketWritable(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/actions/makeBucketWritable", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(MakeBucketWritableRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response MakeBucketWritableResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/Replication/MakeBucketWritable"
		err = common.PostProcessServiceError(err, "ObjectStorage", "MakeBucketWritable", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// MergeObjectMetadata Merges an Object's user-defined metadata with existing metadata.
// A default retry strategy applies to this operation MergeObjectMetadata()
func (client ObjectStorageClient) MergeObjectMetadata(ctx context.Context, request MergeObjectMetadataRequest) (response MergeObjectMetadataResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.mergeObjectMetadata, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = MergeObjectMetadataResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = MergeObjectMetadataResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(MergeObjectMetadataResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into MergeObjectMetadataResponse")
	}
	return
}

// mergeObjectMetadata implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) mergeObjectMetadata(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/actions/mergeObjectMetadata/{objectName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(MergeObjectMetadataRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response MergeObjectMetadataResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/Object/MergeObjectMetadata"
		err = common.PostProcessServiceError(err, "ObjectStorage", "MergeObjectMetadata", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PutObject Creates a new object or overwrites an existing object with the same name. The maximum object size allowed by
// PutObject is 50 GiB.
// See Object Names (https://docs.cloud.oracle.com/Content/Object/Tasks/managingobjects.htm#namerequirements)
// for object naming requirements.
// See Special Instructions for Object Storage PUT (https://docs.cloud.oracle.com/Content/API/Concepts/signingrequests.htm#ObjectStoragePut)
// for request signature requirements.
// A default retry strategy applies to this operation PutObject()
func (client ObjectStorageClient) PutObject(ctx context.Context, request PutObjectRequest) (response PutObjectResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.putObject, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PutObjectResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PutObjectResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PutObjectResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PutObjectResponse")
	}
	return
}

// putObject implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) putObject(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/n/{namespaceName}/b/{bucketName}/o/{objectName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(PutObjectRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response PutObjectResponse
	var httpResponse *http.Response
	var customSigner common.HTTPRequestSigner
	excludeBodySigningPredicate := func(r *http.Request) bool { return false }
	customSigner, err = common.NewSignerFromOCIRequestSigner(client.Signer, excludeBodySigningPredicate)

	//if there was an error overriding the signer, then use the signer from the client itself
	if err != nil {
		customSigner = client.Signer
	}

	//Execute the request with a custom signer
	httpResponse, err = client.CallWithDetails(ctx, &httpRequest, common.ClientCallDetails{Signer: customSigner})
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/Object/PutObject"
		err = common.PostProcessServiceError(err, "ObjectStorage", "PutObject", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PutObjectLifecyclePolicy Creates or replaces the object lifecycle policy for the bucket.
// A default retry strategy applies to this operation PutObjectLifecyclePolicy()
func (client ObjectStorageClient) PutObjectLifecyclePolicy(ctx context.Context, request PutObjectLifecyclePolicyRequest) (response PutObjectLifecyclePolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.putObjectLifecyclePolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PutObjectLifecyclePolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PutObjectLifecyclePolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PutObjectLifecyclePolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PutObjectLifecyclePolicyResponse")
	}
	return
}

// putObjectLifecyclePolicy implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) putObjectLifecyclePolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/n/{namespaceName}/b/{bucketName}/l", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(PutObjectLifecyclePolicyRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response PutObjectLifecyclePolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/ObjectLifecyclePolicy/PutObjectLifecyclePolicy"
		err = common.PostProcessServiceError(err, "ObjectStorage", "PutObjectLifecyclePolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// QueryObject Returns query result as a byte string in the response body.The default format of results is CSV.
// A default retry strategy applies to this operation QueryObject()
func (client ObjectStorageClient) QueryObject(ctx context.Context, request QueryObjectRequest) (response QueryObjectResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.queryObject, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = QueryObjectResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = QueryObjectResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(QueryObjectResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into QueryObjectResponse")
	}
	return
}

// queryObject implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) queryObject(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/actions/query", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(QueryObjectRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response QueryObjectResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/Object/QueryObject"
		err = common.PostProcessServiceError(err, "ObjectStorage", "QueryObject", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ReadObjectSchema Returns object schema(metadata that describes columns if the object represents file in supported columnar file format) result as a json in the response body.
// A default retry strategy applies to this operation ReadObjectSchema()
func (client ObjectStorageClient) ReadObjectSchema(ctx context.Context, request ReadObjectSchemaRequest) (response ReadObjectSchemaResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.readObjectSchema, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ReadObjectSchemaResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ReadObjectSchemaResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ReadObjectSchemaResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ReadObjectSchemaResponse")
	}
	return
}

// readObjectSchema implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) readObjectSchema(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/actions/readObjectSchema", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(ReadObjectSchemaRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response ReadObjectSchemaResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/Object/ReadObjectSchema"
		err = common.PostProcessServiceError(err, "ObjectStorage", "ReadObjectSchema", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ReencryptBucket Re-encrypts the unique data encryption key that encrypts each object written to the bucket by using the most recent
// version of the master encryption key assigned to the bucket. (All data encryption keys are encrypted by a master
// encryption key. Master encryption keys are assigned to buckets and managed by Oracle by default, but you can assign
// a key that you created and control through the Oracle Cloud Infrastructure Key Management service.) The kmsKeyId property
// of the bucket determines which master encryption key is assigned to the bucket. If you assigned a different Key Management
// master encryption key to the bucket, you can call this API to re-encrypt all data encryption keys with the newly
// assigned key. Similarly, you might want to re-encrypt all data encryption keys if the assigned key has been rotated to
// a new key version since objects were last added to the bucket. If you call this API and there is no kmsKeyId associated
// with the bucket, the call will fail.
// Calling this API starts a work request task to re-encrypt the data encryption key of all objects in the bucket. Only
// objects created before the time of the API call will be re-encrypted. The call can take a long time, depending on how many
// objects are in the bucket and how big they are. This API returns a work request ID that you can use to retrieve the status
// of the work request task.
// All the versions of objects will be re-encrypted whether versioning is enabled or suspended at the bucket.
// A default retry strategy applies to this operation ReencryptBucket()
func (client ObjectStorageClient) ReencryptBucket(ctx context.Context, request ReencryptBucketRequest) (response ReencryptBucketResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.reencryptBucket, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ReencryptBucketResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ReencryptBucketResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ReencryptBucketResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ReencryptBucketResponse")
	}
	return
}

// reencryptBucket implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) reencryptBucket(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/actions/reencrypt", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(ReencryptBucketRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response ReencryptBucketResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/Bucket/ReencryptBucket"
		err = common.PostProcessServiceError(err, "ObjectStorage", "ReencryptBucket", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ReencryptObject Re-encrypts the data encryption keys that encrypt the object and its chunks. By default, when you create a bucket, the Object Storage
// service manages the master encryption key used to encrypt each object's data encryption keys. The encryption mechanism that you specify for
// the bucket applies to the objects it contains.
// You can alternatively employ one of these encryption strategies for an object:
// - You can assign a key that you created and control through the Oracle Cloud Infrastructure Vault service.
// - You can encrypt an object using your own encryption key. The key you supply is known as a customer-provided encryption key (SSE-C).
// A default retry strategy applies to this operation ReencryptObject()
func (client ObjectStorageClient) ReencryptObject(ctx context.Context, request ReencryptObjectRequest) (response ReencryptObjectResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.reencryptObject, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ReencryptObjectResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ReencryptObjectResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ReencryptObjectResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ReencryptObjectResponse")
	}
	return
}

// reencryptObject implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) reencryptObject(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/actions/reencrypt/{objectName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(ReencryptObjectRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response ReencryptObjectResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/Object/ReencryptObject"
		err = common.PostProcessServiceError(err, "ObjectStorage", "ReencryptObject", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RenameObject Rename an object in the given Object Storage namespace.
// See Object Names (https://docs.cloud.oracle.com/Content/Object/Tasks/managingobjects.htm#namerequirements)
// for object naming requirements.
// A default retry strategy applies to this operation RenameObject()
func (client ObjectStorageClient) RenameObject(ctx context.Context, request RenameObjectRequest) (response RenameObjectResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.renameObject, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RenameObjectResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RenameObjectResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RenameObjectResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RenameObjectResponse")
	}
	return
}

// renameObject implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) renameObject(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/actions/renameObject", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(RenameObjectRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response RenameObjectResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/Object/RenameObject"
		err = common.PostProcessServiceError(err, "ObjectStorage", "RenameObject", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ReplaceObjectMetadata Overwrites an object's user-defined metadata.
// A default retry strategy applies to this operation ReplaceObjectMetadata()
func (client ObjectStorageClient) ReplaceObjectMetadata(ctx context.Context, request ReplaceObjectMetadataRequest) (response ReplaceObjectMetadataResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.replaceObjectMetadata, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ReplaceObjectMetadataResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ReplaceObjectMetadataResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ReplaceObjectMetadataResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ReplaceObjectMetadataResponse")
	}
	return
}

// replaceObjectMetadata implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) replaceObjectMetadata(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/actions/replaceObjectMetadata/{objectName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(ReplaceObjectMetadataRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response ReplaceObjectMetadataResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/Object/ReplaceObjectMetadata"
		err = common.PostProcessServiceError(err, "ObjectStorage", "ReplaceObjectMetadata", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RestoreObjects Restores the object specified by the objectName parameter.
// By default object will be restored for 24 hours. Duration can be configured using the hours parameter.
// A default retry strategy applies to this operation RestoreObjects()
func (client ObjectStorageClient) RestoreObjects(ctx context.Context, request RestoreObjectsRequest) (response RestoreObjectsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.restoreObjects, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RestoreObjectsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RestoreObjectsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RestoreObjectsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RestoreObjectsResponse")
	}
	return
}

// restoreObjects implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) restoreObjects(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/actions/restoreObjects", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(RestoreObjectsRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response RestoreObjectsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/Object/RestoreObjects"
		err = common.PostProcessServiceError(err, "ObjectStorage", "RestoreObjects", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StartPrefixRename Rename prefix of all the objects in the given source prefix with destination prefix.
// A default retry strategy applies to this operation StartPrefixRename()
func (client ObjectStorageClient) StartPrefixRename(ctx context.Context, request StartPrefixRenameRequest) (response StartPrefixRenameResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.startPrefixRename, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StartPrefixRenameResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StartPrefixRenameResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StartPrefixRenameResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StartPrefixRenameResponse")
	}
	return
}

// startPrefixRename implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) startPrefixRename(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/actions/prefixRename", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(StartPrefixRenameRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response StartPrefixRenameResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/Object/StartPrefixRename"
		err = common.PostProcessServiceError(err, "ObjectStorage", "StartPrefixRename", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateAcl Updates the specified ACL.
// A default retry strategy applies to this operation UpdateAcl()
func (client ObjectStorageClient) UpdateAcl(ctx context.Context, request UpdateAclRequest) (response UpdateAclResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateAcl, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAclResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAclResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAclResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAclResponse")
	}
	return
}

// updateAcl implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) updateAcl(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/n/{namespaceName}/security/acl/{aclId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(UpdateAclRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response UpdateAclResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/Acl/UpdateAcl"
		err = common.PostProcessServiceError(err, "ObjectStorage", "UpdateAcl", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateAclGroup Updates the specified ACL Group.
// A default retry strategy applies to this operation UpdateAclGroup()
func (client ObjectStorageClient) UpdateAclGroup(ctx context.Context, request UpdateAclGroupRequest) (response UpdateAclGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateAclGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAclGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAclGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAclGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAclGroupResponse")
	}
	return
}

// updateAclGroup implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) updateAclGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/n/{namespaceName}/security/aclgroup/{aclGroupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(UpdateAclGroupRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response UpdateAclGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/AclGroup/UpdateAclGroup"
		err = common.PostProcessServiceError(err, "ObjectStorage", "UpdateAclGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateBucket Performs a partial or full update of a bucket's user-defined metadata.
// Use UpdateBucket to move a bucket from one compartment to another within the same tenancy. Supply the compartmentID
// of the compartment that you want to move the bucket to. For more information about moving resources between compartments,
// see Moving Resources to a Different Compartment (https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
// A default retry strategy applies to this operation UpdateBucket()
func (client ObjectStorageClient) UpdateBucket(ctx context.Context, request UpdateBucketRequest) (response UpdateBucketResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateBucket, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateBucketResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateBucketResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateBucketResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateBucketResponse")
	}
	return
}

// updateBucket implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) updateBucket(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/n/{namespaceName}/b/{bucketName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(UpdateBucketRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response UpdateBucketResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/Bucket/UpdateBucket"
		err = common.PostProcessServiceError(err, "ObjectStorage", "UpdateBucket", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateBucketOptions Updates internal options associated with a bucket. The options to be updated/added/removed are specified in a
// JSON object in the body of the request. This API only affects the bucket options that are specified in the JSON
// body. Other options that have been set (by prior use of this API) are left unchanged.
// Options that have a value set to null are removed.
// All the existing options are removed, if the value associated with "freeformOptions" is null.
// This API is for internal-use only.
// A default retry strategy applies to this operation UpdateBucketOptions()
func (client ObjectStorageClient) UpdateBucketOptions(ctx context.Context, request UpdateBucketOptionsRequest) (response UpdateBucketOptionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateBucketOptions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateBucketOptionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateBucketOptionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateBucketOptionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateBucketOptionsResponse")
	}
	return
}

// updateBucketOptions implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) updateBucketOptions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/actions/options", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(UpdateBucketOptionsRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response UpdateBucketOptionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/Bucket/UpdateBucketOptions"
		err = common.PostProcessServiceError(err, "ObjectStorage", "UpdateBucketOptions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateNamespaceMetadata By default, buckets created using the Amazon S3 Compatibility API or the Swift API are created in the root
// compartment of the Oracle Cloud Infrastructure tenancy.
// You can change the default Swift/Amazon S3 compartmentId designation to a different compartmentId. All
// subsequent bucket creations will use the new default compartment, but no previously created
// buckets will be modified. A user must have OBJECTSTORAGE_NAMESPACE_UPDATE permission to make changes to the default
// compartments for Amazon S3 and Swift.
// A default retry strategy applies to this operation UpdateNamespaceMetadata()
func (client ObjectStorageClient) UpdateNamespaceMetadata(ctx context.Context, request UpdateNamespaceMetadataRequest) (response UpdateNamespaceMetadataResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateNamespaceMetadata, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateNamespaceMetadataResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateNamespaceMetadataResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateNamespaceMetadataResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateNamespaceMetadataResponse")
	}
	return
}

// updateNamespaceMetadata implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) updateNamespaceMetadata(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/n/{namespaceName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(UpdateNamespaceMetadataRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response UpdateNamespaceMetadataResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/Namespace/UpdateNamespaceMetadata"
		err = common.PostProcessServiceError(err, "ObjectStorage", "UpdateNamespaceMetadata", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateObjectStorageTier Changes the storage tier of the object specified by the objectName parameter.
// A default retry strategy applies to this operation UpdateObjectStorageTier()
func (client ObjectStorageClient) UpdateObjectStorageTier(ctx context.Context, request UpdateObjectStorageTierRequest) (response UpdateObjectStorageTierResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateObjectStorageTier, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateObjectStorageTierResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateObjectStorageTierResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateObjectStorageTierResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateObjectStorageTierResponse")
	}
	return
}

// updateObjectStorageTier implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) updateObjectStorageTier(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/actions/updateObjectStorageTier", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(UpdateObjectStorageTierRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response UpdateObjectStorageTierResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/Object/UpdateObjectStorageTier"
		err = common.PostProcessServiceError(err, "ObjectStorage", "UpdateObjectStorageTier", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdatePrivateEndpoint Performs a partial or full update of a user-defined data associated with the Private Endpoint.
// Use UpdatePrivateEndpoint to move a Private Endpoint from one compartment to another within the same tenancy. Supply the compartmentID
// of the compartment that you want to move the Private Endpoint to. Or use it to update the name, subnetId, endpointFqdn or privateEndpointIp of the Private Endpoint.
// For more information about moving resources between compartments, see Moving Resources to a Different Compartment (https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
// This API follows replace semantics (rather than merge semantics). That means if the body provides values for
// parameters and the resource has exisiting ones, this operation will replace those existing values.
// A default retry strategy applies to this operation UpdatePrivateEndpoint()
func (client ObjectStorageClient) UpdatePrivateEndpoint(ctx context.Context, request UpdatePrivateEndpointRequest) (response UpdatePrivateEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updatePrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdatePrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdatePrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdatePrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdatePrivateEndpointResponse")
	}
	return
}

// updatePrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) updatePrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/n/{namespaceName}/pe/{peName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(UpdatePrivateEndpointRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response UpdatePrivateEndpointResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/PrivateEndpoint/UpdatePrivateEndpoint"
		err = common.PostProcessServiceError(err, "ObjectStorage", "UpdatePrivateEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateRetentionRule Updates the specified retention rule. Rule changes take effect typically within 30 seconds.
// A default retry strategy applies to this operation UpdateRetentionRule()
func (client ObjectStorageClient) UpdateRetentionRule(ctx context.Context, request UpdateRetentionRuleRequest) (response UpdateRetentionRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateRetentionRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateRetentionRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateRetentionRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateRetentionRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateRetentionRuleResponse")
	}
	return
}

// updateRetentionRule implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) updateRetentionRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/n/{namespaceName}/b/{bucketName}/retentionRules/{retentionRuleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(UpdateRetentionRuleRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response UpdateRetentionRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/RetentionRule/UpdateRetentionRule"
		err = common.PostProcessServiceError(err, "ObjectStorage", "UpdateRetentionRule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UploadPart Uploads a single part of a multipart upload.
// A default retry strategy applies to this operation UploadPart()
func (client ObjectStorageClient) UploadPart(ctx context.Context, request UploadPartRequest) (response UploadPartResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.uploadPart, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UploadPartResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UploadPartResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UploadPartResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UploadPartResponse")
	}
	return
}

// uploadPart implements the OCIOperation interface (enables retrying operations)
func (client ObjectStorageClient) uploadPart(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {
	if !common.IsEnvVarFalse(common.UsingExpectHeaderEnvVar) {
		extraHeaders["Expect"] = "100-continue"
	}
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/n/{namespaceName}/b/{bucketName}/u/{objectName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(UploadPartRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response UploadPartResponse
	var httpResponse *http.Response
	var customSigner common.HTTPRequestSigner
	excludeBodySigningPredicate := func(r *http.Request) bool { return false }
	customSigner, err = common.NewSignerFromOCIRequestSigner(client.Signer, excludeBodySigningPredicate)

	//if there was an error overriding the signer, then use the signer from the client itself
	if err != nil {
		customSigner = client.Signer
	}

	//Execute the request with a custom signer
	httpResponse, err = client.CallWithDetails(ctx, &httpRequest, common.ClientCallDetails{Signer: customSigner})
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/objectstorage/20160918/MultipartUpload/UploadPart"
		err = common.PostProcessServiceError(err, "ObjectStorage", "UploadPart", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
