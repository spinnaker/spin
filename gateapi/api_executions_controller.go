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
	"context"
	"fmt"
	"github.com/antihax/optional"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type ExecutionsControllerApiService service

/*
ExecutionsControllerApiService Retrieves an ad-hoc collection of executions based on a number of user-supplied parameters. Either executionIds or pipelineConfigIds must be supplied in order to return any results. If both are supplied, an exception will be thrown.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *ExecutionsControllerApiGetLatestExecutionsByConfigIdsUsingGETOpts - Optional Parameters:
     * @param "ExecutionIds" (optional.String) -  A comma-separated list of executions to retrieve. Either this OR pipelineConfigIds must be supplied, but not both.
     * @param "Expand" (optional.Bool) -  Expands each execution object in the resulting list. If this value is missing, it is defaulted to true.
     * @param "Limit" (optional.Int32) -  The number of executions to return per pipeline configuration. Ignored if executionIds parameter is supplied. If this value is missing, it is defaulted to 1.
     * @param "PipelineConfigIds" (optional.String) -  A comma-separated list of pipeline configuration IDs to retrieve recent executions for. Either this OR pipelineConfigIds must be supplied, but not both.
     * @param "Statuses" (optional.String) -  A comma-separated list of execution statuses to filter by. Ignored if executionIds parameter is supplied. If this value is missing, it is defaulted to all statuses.

@return []interface{}
*/

type ExecutionsControllerApiGetLatestExecutionsByConfigIdsUsingGETOpts struct {
	ExecutionIds      optional.String
	Expand            optional.Bool
	Limit             optional.Int32
	PipelineConfigIds optional.String
	Statuses          optional.String
}

func (a *ExecutionsControllerApiService) GetLatestExecutionsByConfigIdsUsingGET(ctx context.Context, localVarOptionals *ExecutionsControllerApiGetLatestExecutionsByConfigIdsUsingGETOpts) ([]interface{}, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue []interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/executions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.ExecutionIds.IsSet() {
		localVarQueryParams.Add("executionIds", parameterToString(localVarOptionals.ExecutionIds.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Expand.IsSet() {
		localVarQueryParams.Add("expand", parameterToString(localVarOptionals.Expand.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PipelineConfigIds.IsSet() {
		localVarQueryParams.Add("pipelineConfigIds", parameterToString(localVarOptionals.PipelineConfigIds.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Statuses.IsSet() {
		localVarQueryParams.Add("statuses", parameterToString(localVarOptionals.Statuses.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v []interface{}
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
ExecutionsControllerApiService Search for pipeline executions using a combination of criteria. The returned list is sorted by buildTime (trigger time) in reverse order so that newer executions are first in the list.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param application Only includes executions that are part of this application. If this value is \&quot;*\&quot;, results will include executions of all applications.
 * @param optional nil or *ExecutionsControllerApiSearchForPipelineExecutionsByTriggerUsingGETOpts - Optional Parameters:
     * @param "EventId" (optional.String) -  Only includes executions that were triggered by a trigger with this eventId.
     * @param "Expand" (optional.Bool) -  Expands each execution object in the resulting list. If this value is missing, it is defaulted to false.
     * @param "PipelineName" (optional.String) -  Only includes executions that with this pipeline name.
     * @param "Reverse" (optional.Bool) -  Reverses the resulting list before it is paginated. If this value is missing, it is defaulted to false.
     * @param "Size" (optional.Int32) -  Sets the size of the resulting list for pagination. This value must be &gt; 0. If this value is missing, it is defaulted to 10.
     * @param "StartIndex" (optional.Int32) -  Sets the first item of the resulting list for pagination. The list is 0-indexed. This value must be &gt;&#x3D; 0. If this value is missing, it is defaulted to 0.
     * @param "Statuses" (optional.String) -  Only includes executions with a status that is equal to a status provided in this field. The list of statuses should be given as a comma-delimited string. If this value is missing, includes executions of all statuses. Allowed statuses are: NOT_STARTED, RUNNING, PAUSED, SUSPENDED, SUCCEEDED, FAILED_CONTINUE, TERMINAL, CANCELED, REDIRECT, STOPPED, SKIPPED, BUFFERED.
     * @param "Trigger" (optional.String) -  Only includes executions that were triggered by a trigger that matches the subset of fields provided by this value. This value should be a base64-encoded string of a JSON representation of a trigger object. The comparison succeeds if the execution trigger contains all the fields of the input trigger, the fields are of the same type, and each value of the field \&quot;matches\&quot;. The term \&quot;matches\&quot; is specific for each field&#39;s type: - For Strings: A String value in the execution&#39;s trigger matches the input trigger&#39;s String value if the former equals the latter (case-insensitive) OR if the former matches the latter as a regular expression. - For Maps: A Map value in the execution&#39;s trigger matches the input trigger&#39;s Map value if the former contains all keys of the latter and their values match. - For Collections: A Collection value in the execution&#39;s trigger matches the input trigger&#39;s Collection value if the former has a unique element that matches each element of the latter. - Every other value is compared using the Java \&quot;equals\&quot; method (Groovy \&quot;&#x3D;&#x3D;\&quot; operator)
     * @param "TriggerTimeEndBoundary" (optional.Int64) -  Only includes executions that were built at or before the given time, represented as a Unix timestamp in ms (UTC). This value must be &lt;&#x3D; 9223372036854775807 (Long.MAX_VALUE) and &gt;&#x3D; the value of [triggerTimeStartBoundary], if provided. If this value is missing, it is defaulted to 9223372036854775807.
     * @param "TriggerTimeStartBoundary" (optional.Int64) -  Only includes executions that were built at or after the given time, represented as a Unix timestamp in ms (UTC). This value must be &gt;&#x3D; 0 and &lt;&#x3D; the value of [triggerTimeEndBoundary], if provided. If this value is missing, it is defaulted to 0.
     * @param "TriggerTypes" (optional.String) -  Only includes executions that were triggered by a trigger with a type that is equal to a type provided in this field. The list of trigger types should be a comma-delimited string. If this value is missing, results will includes executions of all trigger types.

@return []interface{}
*/

type ExecutionsControllerApiSearchForPipelineExecutionsByTriggerUsingGETOpts struct {
	EventId                  optional.String
	Expand                   optional.Bool
	PipelineName             optional.String
	Reverse                  optional.Bool
	Size                     optional.Int32
	StartIndex               optional.Int32
	Statuses                 optional.String
	Trigger                  optional.String
	TriggerTimeEndBoundary   optional.Int64
	TriggerTimeStartBoundary optional.Int64
	TriggerTypes             optional.String
}

func (a *ExecutionsControllerApiService) SearchForPipelineExecutionsByTriggerUsingGET(ctx context.Context, application string, localVarOptionals *ExecutionsControllerApiSearchForPipelineExecutionsByTriggerUsingGETOpts) ([]interface{}, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue []interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/applications/{application}/executions/search"
	localVarPath = strings.Replace(localVarPath, "{"+"application"+"}", fmt.Sprintf("%v", application), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.EventId.IsSet() {
		localVarQueryParams.Add("eventId", parameterToString(localVarOptionals.EventId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Expand.IsSet() {
		localVarQueryParams.Add("expand", parameterToString(localVarOptionals.Expand.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PipelineName.IsSet() {
		localVarQueryParams.Add("pipelineName", parameterToString(localVarOptionals.PipelineName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Reverse.IsSet() {
		localVarQueryParams.Add("reverse", parameterToString(localVarOptionals.Reverse.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Size.IsSet() {
		localVarQueryParams.Add("size", parameterToString(localVarOptionals.Size.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.StartIndex.IsSet() {
		localVarQueryParams.Add("startIndex", parameterToString(localVarOptionals.StartIndex.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Statuses.IsSet() {
		localVarQueryParams.Add("statuses", parameterToString(localVarOptionals.Statuses.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Trigger.IsSet() {
		localVarQueryParams.Add("trigger", parameterToString(localVarOptionals.Trigger.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TriggerTimeEndBoundary.IsSet() {
		localVarQueryParams.Add("triggerTimeEndBoundary", parameterToString(localVarOptionals.TriggerTimeEndBoundary.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TriggerTimeStartBoundary.IsSet() {
		localVarQueryParams.Add("triggerTimeStartBoundary", parameterToString(localVarOptionals.TriggerTimeStartBoundary.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TriggerTypes.IsSet() {
		localVarQueryParams.Add("triggerTypes", parameterToString(localVarOptionals.TriggerTypes.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v []interface{}
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
