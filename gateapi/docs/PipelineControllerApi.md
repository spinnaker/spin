# \PipelineControllerApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelPipelineUsingPUT1**](PipelineControllerApi.md#CancelPipelineUsingPUT1) | **Put** /pipelines/{id}/cancel | Cancel a pipeline execution
[**DeletePipelineUsingDELETE**](PipelineControllerApi.md#DeletePipelineUsingDELETE) | **Delete** /pipelines/{application}/{pipelineName} | Delete a pipeline definition
[**DeletePipelineUsingDELETE1**](PipelineControllerApi.md#DeletePipelineUsingDELETE1) | **Delete** /pipelines/{id} | Delete a pipeline execution
[**EvaluateExpressionForExecutionAtStageUsingDELETE**](PipelineControllerApi.md#EvaluateExpressionForExecutionAtStageUsingDELETE) | **Delete** /pipelines/{id}/{stageId}/evaluateExpression | Evaluate a pipeline expression at a specific stage using the provided execution as context
[**EvaluateExpressionForExecutionAtStageUsingGET**](PipelineControllerApi.md#EvaluateExpressionForExecutionAtStageUsingGET) | **Get** /pipelines/{id}/{stageId}/evaluateExpression | Evaluate a pipeline expression at a specific stage using the provided execution as context
[**EvaluateExpressionForExecutionAtStageUsingHEAD**](PipelineControllerApi.md#EvaluateExpressionForExecutionAtStageUsingHEAD) | **Head** /pipelines/{id}/{stageId}/evaluateExpression | Evaluate a pipeline expression at a specific stage using the provided execution as context
[**EvaluateExpressionForExecutionAtStageUsingOPTIONS**](PipelineControllerApi.md#EvaluateExpressionForExecutionAtStageUsingOPTIONS) | **Options** /pipelines/{id}/{stageId}/evaluateExpression | Evaluate a pipeline expression at a specific stage using the provided execution as context
[**EvaluateExpressionForExecutionAtStageUsingPATCH**](PipelineControllerApi.md#EvaluateExpressionForExecutionAtStageUsingPATCH) | **Patch** /pipelines/{id}/{stageId}/evaluateExpression | Evaluate a pipeline expression at a specific stage using the provided execution as context
[**EvaluateExpressionForExecutionAtStageUsingPOST**](PipelineControllerApi.md#EvaluateExpressionForExecutionAtStageUsingPOST) | **Post** /pipelines/{id}/{stageId}/evaluateExpression | Evaluate a pipeline expression at a specific stage using the provided execution as context
[**EvaluateExpressionForExecutionAtStageUsingPUT**](PipelineControllerApi.md#EvaluateExpressionForExecutionAtStageUsingPUT) | **Put** /pipelines/{id}/{stageId}/evaluateExpression | Evaluate a pipeline expression at a specific stage using the provided execution as context
[**EvaluateExpressionForExecutionUsingDELETE**](PipelineControllerApi.md#EvaluateExpressionForExecutionUsingDELETE) | **Delete** /pipelines/{id}/evaluateExpression | Evaluate a pipeline expression using the provided execution as context
[**EvaluateExpressionForExecutionUsingGET**](PipelineControllerApi.md#EvaluateExpressionForExecutionUsingGET) | **Get** /pipelines/{id}/evaluateExpression | Evaluate a pipeline expression using the provided execution as context
[**EvaluateExpressionForExecutionUsingHEAD**](PipelineControllerApi.md#EvaluateExpressionForExecutionUsingHEAD) | **Head** /pipelines/{id}/evaluateExpression | Evaluate a pipeline expression using the provided execution as context
[**EvaluateExpressionForExecutionUsingOPTIONS**](PipelineControllerApi.md#EvaluateExpressionForExecutionUsingOPTIONS) | **Options** /pipelines/{id}/evaluateExpression | Evaluate a pipeline expression using the provided execution as context
[**EvaluateExpressionForExecutionUsingPATCH**](PipelineControllerApi.md#EvaluateExpressionForExecutionUsingPATCH) | **Patch** /pipelines/{id}/evaluateExpression | Evaluate a pipeline expression using the provided execution as context
[**EvaluateExpressionForExecutionUsingPUT**](PipelineControllerApi.md#EvaluateExpressionForExecutionUsingPUT) | **Put** /pipelines/{id}/evaluateExpression | Evaluate a pipeline expression using the provided execution as context
[**EvaluateExpressionForExecutionViaPOSTUsingPOST**](PipelineControllerApi.md#EvaluateExpressionForExecutionViaPOSTUsingPOST) | **Post** /pipelines/{id}/evaluateExpression | Evaluate a pipeline expression using the provided execution as context
[**GetPipelineUsingGET**](PipelineControllerApi.md#GetPipelineUsingGET) | **Get** /pipelines/{id} | Retrieve a pipeline execution
[**InvokePipelineConfigUsingPOST1**](PipelineControllerApi.md#InvokePipelineConfigUsingPOST1) | **Post** /pipelines/{application}/{pipelineNameOrId} | Trigger a pipeline execution
[**InvokePipelineConfigViaEchoUsingPOST**](PipelineControllerApi.md#InvokePipelineConfigViaEchoUsingPOST) | **Post** /pipelines/v2/{application}/{pipelineNameOrId} | Trigger a pipeline execution
[**PausePipelineUsingPUT**](PipelineControllerApi.md#PausePipelineUsingPUT) | **Put** /pipelines/{id}/pause | Pause a pipeline execution
[**RenamePipelineUsingPOST**](PipelineControllerApi.md#RenamePipelineUsingPOST) | **Post** /pipelines/move | Rename a pipeline definition
[**RestartStageUsingPUT**](PipelineControllerApi.md#RestartStageUsingPUT) | **Put** /pipelines/{id}/stages/{stageId}/restart | Restart a stage execution
[**ResumePipelineUsingPUT**](PipelineControllerApi.md#ResumePipelineUsingPUT) | **Put** /pipelines/{id}/resume | Resume a pipeline execution
[**SavePipelineUsingPOST**](PipelineControllerApi.md#SavePipelineUsingPOST) | **Post** /pipelines | Save a pipeline definition
[**StartUsingPOST**](PipelineControllerApi.md#StartUsingPOST) | **Post** /pipelines/start | Initiate a pipeline execution
[**UpdatePipelineUsingPUT**](PipelineControllerApi.md#UpdatePipelineUsingPUT) | **Put** /pipelines/{id} | Update a pipeline definition
[**UpdateStageUsingPATCH**](PipelineControllerApi.md#UpdateStageUsingPATCH) | **Patch** /pipelines/{id}/stages/{stageId} | Update a stage execution


# **CancelPipelineUsingPUT1**
> CancelPipelineUsingPUT1(ctx, id, optional)
Cancel a pipeline execution

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**| id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| id | 
 **force** | **bool**| force | [default to false]
 **reason** | **string**| reason | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePipelineUsingDELETE**
> DeletePipelineUsingDELETE(ctx, application, pipelineName)
Delete a pipeline definition

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **application** | **string**| application | 
  **pipelineName** | **string**| pipelineName | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePipelineUsingDELETE1**
> map[string]interface{} DeletePipelineUsingDELETE1(ctx, id)
Delete a pipeline execution

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**| id | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EvaluateExpressionForExecutionAtStageUsingDELETE**
> map[string]interface{} EvaluateExpressionForExecutionAtStageUsingDELETE(ctx, expression, id, stageId)
Evaluate a pipeline expression at a specific stage using the provided execution as context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **expression** | **string**| expression | 
  **id** | **string**| id | 
  **stageId** | **string**| stageId | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EvaluateExpressionForExecutionAtStageUsingGET**
> map[string]interface{} EvaluateExpressionForExecutionAtStageUsingGET(ctx, expression, id, stageId)
Evaluate a pipeline expression at a specific stage using the provided execution as context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **expression** | **string**| expression | 
  **id** | **string**| id | 
  **stageId** | **string**| stageId | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EvaluateExpressionForExecutionAtStageUsingHEAD**
> map[string]interface{} EvaluateExpressionForExecutionAtStageUsingHEAD(ctx, expression, id, stageId)
Evaluate a pipeline expression at a specific stage using the provided execution as context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **expression** | **string**| expression | 
  **id** | **string**| id | 
  **stageId** | **string**| stageId | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EvaluateExpressionForExecutionAtStageUsingOPTIONS**
> map[string]interface{} EvaluateExpressionForExecutionAtStageUsingOPTIONS(ctx, expression, id, stageId)
Evaluate a pipeline expression at a specific stage using the provided execution as context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **expression** | **string**| expression | 
  **id** | **string**| id | 
  **stageId** | **string**| stageId | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EvaluateExpressionForExecutionAtStageUsingPATCH**
> map[string]interface{} EvaluateExpressionForExecutionAtStageUsingPATCH(ctx, expression, id, stageId)
Evaluate a pipeline expression at a specific stage using the provided execution as context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **expression** | **string**| expression | 
  **id** | **string**| id | 
  **stageId** | **string**| stageId | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EvaluateExpressionForExecutionAtStageUsingPOST**
> map[string]interface{} EvaluateExpressionForExecutionAtStageUsingPOST(ctx, expression, id, stageId)
Evaluate a pipeline expression at a specific stage using the provided execution as context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **expression** | **string**| expression | 
  **id** | **string**| id | 
  **stageId** | **string**| stageId | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EvaluateExpressionForExecutionAtStageUsingPUT**
> map[string]interface{} EvaluateExpressionForExecutionAtStageUsingPUT(ctx, expression, id, stageId)
Evaluate a pipeline expression at a specific stage using the provided execution as context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **expression** | **string**| expression | 
  **id** | **string**| id | 
  **stageId** | **string**| stageId | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EvaluateExpressionForExecutionUsingDELETE**
> map[string]interface{} EvaluateExpressionForExecutionUsingDELETE(ctx, expression, id)
Evaluate a pipeline expression using the provided execution as context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **expression** | **string**| expression | 
  **id** | **string**| id | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EvaluateExpressionForExecutionUsingGET**
> map[string]interface{} EvaluateExpressionForExecutionUsingGET(ctx, expression, id)
Evaluate a pipeline expression using the provided execution as context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **expression** | **string**| expression | 
  **id** | **string**| id | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EvaluateExpressionForExecutionUsingHEAD**
> map[string]interface{} EvaluateExpressionForExecutionUsingHEAD(ctx, expression, id)
Evaluate a pipeline expression using the provided execution as context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **expression** | **string**| expression | 
  **id** | **string**| id | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EvaluateExpressionForExecutionUsingOPTIONS**
> map[string]interface{} EvaluateExpressionForExecutionUsingOPTIONS(ctx, expression, id)
Evaluate a pipeline expression using the provided execution as context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **expression** | **string**| expression | 
  **id** | **string**| id | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EvaluateExpressionForExecutionUsingPATCH**
> map[string]interface{} EvaluateExpressionForExecutionUsingPATCH(ctx, expression, id)
Evaluate a pipeline expression using the provided execution as context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **expression** | **string**| expression | 
  **id** | **string**| id | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EvaluateExpressionForExecutionUsingPUT**
> map[string]interface{} EvaluateExpressionForExecutionUsingPUT(ctx, expression, id)
Evaluate a pipeline expression using the provided execution as context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **expression** | **string**| expression | 
  **id** | **string**| id | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EvaluateExpressionForExecutionViaPOSTUsingPOST**
> map[string]interface{} EvaluateExpressionForExecutionViaPOSTUsingPOST(ctx, id, pipelineExpression)
Evaluate a pipeline expression using the provided execution as context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**| id | 
  **pipelineExpression** | [**interface{}**](interface{}.md)| pipelineExpression | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPipelineUsingGET**
> interface{} GetPipelineUsingGET(ctx, id)
Retrieve a pipeline execution

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**| id | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InvokePipelineConfigUsingPOST1**
> HttpEntity InvokePipelineConfigUsingPOST1(ctx, application, pipelineNameOrId, optional)
Trigger a pipeline execution

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **application** | **string**| application | 
  **pipelineNameOrId** | **string**| pipelineNameOrId | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **application** | **string**| application | 
 **pipelineNameOrId** | **string**| pipelineNameOrId | 
 **trigger** | [**interface{}**](interface{}.md)| trigger | 

### Return type

[**HttpEntity**](HttpEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InvokePipelineConfigViaEchoUsingPOST**
> HttpEntity InvokePipelineConfigViaEchoUsingPOST(ctx, application, pipelineNameOrId, optional)
Trigger a pipeline execution

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **application** | **string**| application | 
  **pipelineNameOrId** | **string**| pipelineNameOrId | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **application** | **string**| application | 
 **pipelineNameOrId** | **string**| pipelineNameOrId | 
 **trigger** | [**interface{}**](interface{}.md)| trigger | 

### Return type

[**HttpEntity**](HttpEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PausePipelineUsingPUT**
> PausePipelineUsingPUT(ctx, id)
Pause a pipeline execution

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**| id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RenamePipelineUsingPOST**
> RenamePipelineUsingPOST(ctx, renameCommand)
Rename a pipeline definition

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **renameCommand** | [**interface{}**](interface{}.md)| renameCommand | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestartStageUsingPUT**
> map[string]interface{} RestartStageUsingPUT(ctx, context, id, stageId)
Restart a stage execution

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **context** | [**interface{}**](interface{}.md)| context | 
  **id** | **string**| id | 
  **stageId** | **string**| stageId | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResumePipelineUsingPUT**
> map[string]interface{} ResumePipelineUsingPUT(ctx, id)
Resume a pipeline execution

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**| id | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SavePipelineUsingPOST**
> SavePipelineUsingPOST(ctx, pipeline)
Save a pipeline definition

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **pipeline** | [**interface{}**](interface{}.md)| pipeline | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartUsingPOST**
> ResponseEntity StartUsingPOST(ctx, map_)
Initiate a pipeline execution

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **map_** | [**interface{}**](interface{}.md)| map | 

### Return type

[**ResponseEntity**](ResponseEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePipelineUsingPUT**
> map[string]interface{} UpdatePipelineUsingPUT(ctx, id, pipeline)
Update a pipeline definition

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**| id | 
  **pipeline** | [**interface{}**](interface{}.md)| pipeline | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateStageUsingPATCH**
> map[string]interface{} UpdateStageUsingPATCH(ctx, context, id, stageId)
Update a stage execution

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **context** | [**interface{}**](interface{}.md)| context | 
  **id** | **string**| id | 
  **stageId** | **string**| stageId | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

