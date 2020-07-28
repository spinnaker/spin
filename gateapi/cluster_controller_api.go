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
	"io/ioutil"
	"net/url"
	"net/http"
	"strings"
	"golang.org/x/net/context"
	"encoding/json"
	"fmt"
)

// Linger please
var (
	_ context.Context
)

type ClusterControllerApiService service


/* ClusterControllerApiService Retrieve a cluster&#39;s loadbalancers
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param account account
 @param applicationName applicationName
 @param clusterName clusterName
 @param type_ type
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "xRateLimitApp" (string) X-RateLimit-App
 @return []interface{}*/
func (a *ClusterControllerApiService) GetClusterLoadBalancersUsingGET(ctx context.Context, account string, applicationName string, clusterName string, type_ string, localVarOptionals map[string]interface{}) ([]interface{},  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/applications/{application}/clusters/{account}/{clusterName}/{type}/loadBalancers"
	localVarPath = strings.Replace(localVarPath, "{"+"account"+"}", fmt.Sprintf("%v", account), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"applicationName"+"}", fmt.Sprintf("%v", applicationName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", fmt.Sprintf("%v", clusterName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"type"+"}", fmt.Sprintf("%v", type_), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["xRateLimitApp"], "string", "xRateLimitApp"); err != nil {
		return successPayload, nil, err
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

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

/* ClusterControllerApiService Retrieve a cluster&#39;s details
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param account account
 @param application application
 @param clusterName clusterName
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "xRateLimitApp" (string) X-RateLimit-App
 @return map[string]interface{}*/
func (a *ClusterControllerApiService) GetClustersUsingGET(ctx context.Context, account string, application string, clusterName string, localVarOptionals map[string]interface{}) (map[string]interface{},  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  map[string]interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/applications/{application}/clusters/{account}/{clusterName}"
	localVarPath = strings.Replace(localVarPath, "{"+"account"+"}", fmt.Sprintf("%v", account), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"application"+"}", fmt.Sprintf("%v", application), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", fmt.Sprintf("%v", clusterName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["xRateLimitApp"], "string", "xRateLimitApp"); err != nil {
		return successPayload, nil, err
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

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

/* ClusterControllerApiService Retrieve a list of clusters for an account
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param account account
 @param application application
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "xRateLimitApp" (string) X-RateLimit-App
 @return []interface{}*/
func (a *ClusterControllerApiService) GetClustersUsingGET1(ctx context.Context, account string, application string, localVarOptionals map[string]interface{}) ([]interface{},  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/applications/{application}/clusters/{account}"
	localVarPath = strings.Replace(localVarPath, "{"+"account"+"}", fmt.Sprintf("%v", account), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"application"+"}", fmt.Sprintf("%v", application), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["xRateLimitApp"], "string", "xRateLimitApp"); err != nil {
		return successPayload, nil, err
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

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

/* ClusterControllerApiService Retrieve a list of cluster names for an application, grouped by account
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param application application
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "xRateLimitApp" (string) X-RateLimit-App
 @return map[string]interface{}*/
func (a *ClusterControllerApiService) GetClustersUsingGET2(ctx context.Context, application string, localVarOptionals map[string]interface{}) (map[string]interface{},  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  map[string]interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/applications/{application}/clusters"
	localVarPath = strings.Replace(localVarPath, "{"+"application"+"}", fmt.Sprintf("%v", application), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["xRateLimitApp"], "string", "xRateLimitApp"); err != nil {
		return successPayload, nil, err
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

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

/* ClusterControllerApiService Retrieve a list of scaling activities for a server group
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param account account
 @param application application
 @param clusterName clusterName
 @param serverGroupName serverGroupName
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "xRateLimitApp" (string) X-RateLimit-App
     @param "provider" (string) provider
     @param "region" (string) region
 @return []interface{}*/
func (a *ClusterControllerApiService) GetScalingActivitiesUsingGET(ctx context.Context, account string, application string, clusterName string, serverGroupName string, localVarOptionals map[string]interface{}) ([]interface{},  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/applications/{application}/clusters/{account}/{clusterName}/serverGroups/{serverGroupName}/scalingActivities"
	localVarPath = strings.Replace(localVarPath, "{"+"account"+"}", fmt.Sprintf("%v", account), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"application"+"}", fmt.Sprintf("%v", application), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", fmt.Sprintf("%v", clusterName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"serverGroupName"+"}", fmt.Sprintf("%v", serverGroupName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["xRateLimitApp"], "string", "xRateLimitApp"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["provider"], "string", "provider"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["region"], "string", "region"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["provider"].(string); localVarOk {
		localVarQueryParams.Add("provider", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["region"].(string); localVarOk {
		localVarQueryParams.Add("region", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

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

/* ClusterControllerApiService Retrieve a server group&#39;s details
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param account account
 @param application application
 @param clusterName clusterName
 @param serverGroupName serverGroupName
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "xRateLimitApp" (string) X-RateLimit-App
 @return []interface{}*/
func (a *ClusterControllerApiService) GetServerGroupsUsingGET(ctx context.Context, account string, application string, clusterName string, serverGroupName string, localVarOptionals map[string]interface{}) ([]interface{},  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/applications/{application}/clusters/{account}/{clusterName}/serverGroups/{serverGroupName}"
	localVarPath = strings.Replace(localVarPath, "{"+"account"+"}", fmt.Sprintf("%v", account), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"application"+"}", fmt.Sprintf("%v", application), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", fmt.Sprintf("%v", clusterName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"serverGroupName"+"}", fmt.Sprintf("%v", serverGroupName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["xRateLimitApp"], "string", "xRateLimitApp"); err != nil {
		return successPayload, nil, err
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

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

/* ClusterControllerApiService Retrieve a list of server groups for a cluster
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param account account
 @param application application
 @param clusterName clusterName
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "xRateLimitApp" (string) X-RateLimit-App
 @return []interface{}*/
func (a *ClusterControllerApiService) GetServerGroupsUsingGET1(ctx context.Context, account string, application string, clusterName string, localVarOptionals map[string]interface{}) ([]interface{},  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/applications/{application}/clusters/{account}/{clusterName}/serverGroups"
	localVarPath = strings.Replace(localVarPath, "{"+"account"+"}", fmt.Sprintf("%v", account), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"application"+"}", fmt.Sprintf("%v", application), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", fmt.Sprintf("%v", clusterName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["xRateLimitApp"], "string", "xRateLimitApp"); err != nil {
		return successPayload, nil, err
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

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

/* ClusterControllerApiService Retrieve a server group that matches a target coordinate (e.g., newest, ancestor) relative to a cluster
 &#x60;scope&#x60; is either a zone or a region
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param account account
 @param application application
 @param cloudProvider cloudProvider
 @param clusterName clusterName
 @param scope scope
 @param target target
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "xRateLimitApp" (string) X-RateLimit-App
     @param "onlyEnabled" (bool) onlyEnabled
     @param "validateOldest" (bool) validateOldest
 @return map[string]interface{}*/
func (a *ClusterControllerApiService) GetTargetServerGroupUsingGET(ctx context.Context, account string, application string, cloudProvider string, clusterName string, scope string, target string, localVarOptionals map[string]interface{}) (map[string]interface{},  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  map[string]interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/applications/{application}/clusters/{account}/{clusterName}/{cloudProvider}/{scope}/serverGroups/target/{target}"
	localVarPath = strings.Replace(localVarPath, "{"+"account"+"}", fmt.Sprintf("%v", account), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"application"+"}", fmt.Sprintf("%v", application), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cloudProvider"+"}", fmt.Sprintf("%v", cloudProvider), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", fmt.Sprintf("%v", clusterName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"scope"+"}", fmt.Sprintf("%v", scope), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"target"+"}", fmt.Sprintf("%v", target), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["xRateLimitApp"], "string", "xRateLimitApp"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["onlyEnabled"], "bool", "onlyEnabled"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["validateOldest"], "bool", "validateOldest"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["onlyEnabled"].(bool); localVarOk {
		localVarQueryParams.Add("onlyEnabled", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["validateOldest"].(bool); localVarOk {
		localVarQueryParams.Add("validateOldest", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

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

