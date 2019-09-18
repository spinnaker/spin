# \ClusterControllerApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetClusterLoadBalancersUsingGET**](ClusterControllerApi.md#GetClusterLoadBalancersUsingGET) | **Get** /applications/{application}/clusters/{account}/{clusterName}/{type}/loadBalancers | Retrieve a cluster&#39;s loadbalancers
[**GetClustersUsingGET**](ClusterControllerApi.md#GetClustersUsingGET) | **Get** /applications/{application}/clusters/{account}/{clusterName} | Retrieve a cluster&#39;s details
[**GetClustersUsingGET1**](ClusterControllerApi.md#GetClustersUsingGET1) | **Get** /applications/{application}/clusters/{account} | Retrieve a list of clusters for an account
[**GetClustersUsingGET2**](ClusterControllerApi.md#GetClustersUsingGET2) | **Get** /applications/{application}/clusters | Retrieve a list of cluster names for an application, grouped by account
[**GetScalingActivitiesUsingGET**](ClusterControllerApi.md#GetScalingActivitiesUsingGET) | **Get** /applications/{application}/clusters/{account}/{clusterName}/serverGroups/{serverGroupName}/scalingActivities | Retrieve a list of scaling activities for a server group
[**GetServerGroupsUsingGET**](ClusterControllerApi.md#GetServerGroupsUsingGET) | **Get** /applications/{application}/clusters/{account}/{clusterName}/serverGroups/{serverGroupName} | Retrieve a server group&#39;s details
[**GetServerGroupsUsingGET1**](ClusterControllerApi.md#GetServerGroupsUsingGET1) | **Get** /applications/{application}/clusters/{account}/{clusterName}/serverGroups | Retrieve a list of server groups for a cluster
[**GetTargetServerGroupUsingGET**](ClusterControllerApi.md#GetTargetServerGroupUsingGET) | **Get** /applications/{application}/clusters/{account}/{clusterName}/{cloudProvider}/{scope}/serverGroups/target/{target} | Retrieve a server group that matches a target coordinate (e.g., newest, ancestor) relative to a cluster


# **GetClusterLoadBalancersUsingGET**
> []interface{} GetClusterLoadBalancersUsingGET(ctx, account, applicationName, clusterName, type_, optional)
Retrieve a cluster's loadbalancers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **account** | **string**| account | 
  **applicationName** | **string**| applicationName | 
  **clusterName** | **string**| clusterName | 
  **type_** | **string**| type | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string**| account | 
 **applicationName** | **string**| applicationName | 
 **clusterName** | **string**| clusterName | 
 **type_** | **string**| type | 
 **xRateLimitApp** | **string**| X-RateLimit-App | 

### Return type

[**[]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClustersUsingGET**
> map[string]interface{} GetClustersUsingGET(ctx, account, application, clusterName, optional)
Retrieve a cluster's details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **account** | **string**| account | 
  **application** | **string**| application | 
  **clusterName** | **string**| clusterName | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string**| account | 
 **application** | **string**| application | 
 **clusterName** | **string**| clusterName | 
 **xRateLimitApp** | **string**| X-RateLimit-App | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClustersUsingGET1**
> []interface{} GetClustersUsingGET1(ctx, account, application, optional)
Retrieve a list of clusters for an account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **account** | **string**| account | 
  **application** | **string**| application | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string**| account | 
 **application** | **string**| application | 
 **xRateLimitApp** | **string**| X-RateLimit-App | 

### Return type

[**[]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClustersUsingGET2**
> map[string]interface{} GetClustersUsingGET2(ctx, application, optional)
Retrieve a list of cluster names for an application, grouped by account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **application** | **string**| application | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **application** | **string**| application | 
 **xRateLimitApp** | **string**| X-RateLimit-App | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetScalingActivitiesUsingGET**
> []interface{} GetScalingActivitiesUsingGET(ctx, account, application, clusterName, serverGroupName, optional)
Retrieve a list of scaling activities for a server group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **account** | **string**| account | 
  **application** | **string**| application | 
  **clusterName** | **string**| clusterName | 
  **serverGroupName** | **string**| serverGroupName | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string**| account | 
 **application** | **string**| application | 
 **clusterName** | **string**| clusterName | 
 **serverGroupName** | **string**| serverGroupName | 
 **xRateLimitApp** | **string**| X-RateLimit-App | 
 **provider** | **string**| provider | [default to aws]
 **region** | **string**| region | 

### Return type

[**[]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServerGroupsUsingGET**
> []interface{} GetServerGroupsUsingGET(ctx, account, application, clusterName, serverGroupName, optional)
Retrieve a server group's details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **account** | **string**| account | 
  **application** | **string**| application | 
  **clusterName** | **string**| clusterName | 
  **serverGroupName** | **string**| serverGroupName | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string**| account | 
 **application** | **string**| application | 
 **clusterName** | **string**| clusterName | 
 **serverGroupName** | **string**| serverGroupName | 
 **xRateLimitApp** | **string**| X-RateLimit-App | 

### Return type

[**[]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServerGroupsUsingGET1**
> []interface{} GetServerGroupsUsingGET1(ctx, account, application, clusterName, optional)
Retrieve a list of server groups for a cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **account** | **string**| account | 
  **application** | **string**| application | 
  **clusterName** | **string**| clusterName | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string**| account | 
 **application** | **string**| application | 
 **clusterName** | **string**| clusterName | 
 **xRateLimitApp** | **string**| X-RateLimit-App | 

### Return type

[**[]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTargetServerGroupUsingGET**
> map[string]interface{} GetTargetServerGroupUsingGET(ctx, account, application, cloudProvider, clusterName, scope, target, optional)
Retrieve a server group that matches a target coordinate (e.g., newest, ancestor) relative to a cluster

`scope` is either a zone or a region

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **account** | **string**| account | 
  **application** | **string**| application | 
  **cloudProvider** | **string**| cloudProvider | 
  **clusterName** | **string**| clusterName | 
  **scope** | **string**| scope | 
  **target** | **string**| target | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string**| account | 
 **application** | **string**| application | 
 **cloudProvider** | **string**| cloudProvider | 
 **clusterName** | **string**| clusterName | 
 **scope** | **string**| scope | 
 **target** | **string**| target | 
 **xRateLimitApp** | **string**| X-RateLimit-App | 
 **onlyEnabled** | **bool**| onlyEnabled | 
 **validateOldest** | **bool**| validateOldest | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

