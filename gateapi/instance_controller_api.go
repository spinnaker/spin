/*
 * Spinnaker API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/context"
)

// Linger please
var (
	_ context.Context
)

type InstanceControllerApiService service

/* InstanceControllerApiService Retrieve an instance&#39;s console output
* @param ctx context.Context for authentication, logging, tracing, etc.
@param account account
@param instanceId instanceId
@param region region
@param optional (nil or map[string]interface{}) with one or more of:
    @param "xRateLimitApp" (string) X-RateLimit-App
    @param "provider" (string) provider
@return interface{}*/
func (a *InstanceControllerApiService) GetConsoleOutputUsingGET(ctx context.Context, account string, instanceId string, region string, localVarOptionals map[string]interface{}) (interface{}, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/instances/{account}/{region}/{instanceId}/console"
	localVarPath = strings.Replace(localVarPath, "{"+"account"+"}", fmt.Sprintf("%v", account), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"instanceId"+"}", fmt.Sprintf("%v", instanceId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"region"+"}", fmt.Sprintf("%v", region), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["xRateLimitApp"], "string", "xRateLimitApp"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["provider"], "string", "provider"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["provider"].(string); localVarOk {
		localVarQueryParams.Add("provider", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"*/*",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["xRateLimitApp"].(string); localVarOk {
		localVarHeaderParams["X-RateLimit-App"] = parameterToString(localVarTempParam, "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* InstanceControllerApiService Retrieve an instance&#39;s details
* @param ctx context.Context for authentication, logging, tracing, etc.
@param account account
@param instanceId instanceId
@param region region
@param optional (nil or map[string]interface{}) with one or more of:
    @param "xRateLimitApp" (string) X-RateLimit-App
@return interface{}*/
func (a *InstanceControllerApiService) GetInstanceDetailsUsingGET(ctx context.Context, account string, instanceId string, region string, localVarOptionals map[string]interface{}) (interface{}, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/instances/{account}/{region}/{instanceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"account"+"}", fmt.Sprintf("%v", account), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"instanceId"+"}", fmt.Sprintf("%v", instanceId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"region"+"}", fmt.Sprintf("%v", region), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["xRateLimitApp"], "string", "xRateLimitApp"); err != nil {
		return successPayload, nil, err
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"*/*",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["xRateLimitApp"].(string); localVarOk {
		localVarHeaderParams["X-RateLimit-App"] = parameterToString(localVarTempParam, "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}
