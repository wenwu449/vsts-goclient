/*
 * Git
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.0-preview
 * Contact: nugetvss@microsoft.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type PullRequestThreadsApiService service

/*
PullRequestThreadsApiService
Create a thread in a pull request.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body The thread to create. Thread must contain at least one comment.
 * @param repositoryId Repository ID of the pull request&#39;s target branch.
 * @param pullRequestId ID of the pull request.
 * @param project Project ID or project name
 * @param apiVersion Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api.

@return GitPullRequestCommentThread
*/
func (a *PullRequestThreadsApiService) Create(ctx context.Context, body GitPullRequestCommentThread, repositoryId string, pullRequestId int32, project string, apiVersion string) (GitPullRequestCommentThread, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue GitPullRequestCommentThread
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/threads"
	localVarPath = strings.Replace(localVarPath, "{"+"repositoryId"+"}", fmt.Sprintf("%v", repositoryId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pullRequestId"+"}", fmt.Sprintf("%v", pullRequestId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", fmt.Sprintf("%v", project), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("api-version", parameterToString(apiVersion, ""))
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
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
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarReturnValue, localVarHttpResponse, err
			}

			var v GitPullRequestCommentThread
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
PullRequestThreadsApiService
Retrieve a thread in a pull request.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param repositoryId The repository ID of the pull request&#39;s target branch.
 * @param pullRequestId ID of the pull request.
 * @param threadId ID of the thread.
 * @param project Project ID or project name
 * @param apiVersion Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api.
 * @param optional nil or *GetOpts - Optional Parameters:
     * @param "Iteration" (optional.Int32) -  If specified, thread position will be tracked using this iteration as the right side of the diff.
     * @param "BaseIteration" (optional.Int32) -  If specified, thread position will be tracked using this iteration as the left side of the diff.

@return GitPullRequestCommentThread
*/

type GetOpts struct {
	Iteration     optional.Int32
	BaseIteration optional.Int32
}

func (a *PullRequestThreadsApiService) Get(ctx context.Context, repositoryId string, pullRequestId int32, threadId int32, project string, apiVersion string, localVarOptionals *GetOpts) (GitPullRequestCommentThread, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue GitPullRequestCommentThread
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/threads/{threadId}"
	localVarPath = strings.Replace(localVarPath, "{"+"repositoryId"+"}", fmt.Sprintf("%v", repositoryId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pullRequestId"+"}", fmt.Sprintf("%v", pullRequestId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"threadId"+"}", fmt.Sprintf("%v", threadId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", fmt.Sprintf("%v", project), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Iteration.IsSet() {
		localVarQueryParams.Add("$iteration", parameterToString(localVarOptionals.Iteration.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BaseIteration.IsSet() {
		localVarQueryParams.Add("$baseIteration", parameterToString(localVarOptionals.BaseIteration.Value(), ""))
	}
	localVarQueryParams.Add("api-version", parameterToString(apiVersion, ""))
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

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
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarReturnValue, localVarHttpResponse, err
			}

			var v GitPullRequestCommentThread
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
PullRequestThreadsApiService
Retrieve all threads in a pull request.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param repositoryId The repository ID of the pull request&#39;s target branch.
 * @param pullRequestId ID of the pull request.
 * @param project Project ID or project name
 * @param apiVersion Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api.
 * @param optional nil or *ListPRThreadsOpts - Optional Parameters:
     * @param "Iteration" (optional.Int32) -  If specified, thread positions will be tracked using this iteration as the right side of the diff.
     * @param "BaseIteration" (optional.Int32) -  If specified, thread positions will be tracked using this iteration as the left side of the diff.

@return []GitPullRequestCommentThread
*/

type ListPRThreadsOpts struct {
	Iteration     optional.Int32
	BaseIteration optional.Int32
}

func (a *PullRequestThreadsApiService) List(ctx context.Context, repositoryId string, pullRequestId int32, project string, apiVersion string, localVarOptionals *ListPRThreadsOpts) ([]GitPullRequestCommentThread, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue []GitPullRequestCommentThread
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/threads"
	localVarPath = strings.Replace(localVarPath, "{"+"repositoryId"+"}", fmt.Sprintf("%v", repositoryId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pullRequestId"+"}", fmt.Sprintf("%v", pullRequestId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", fmt.Sprintf("%v", project), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Iteration.IsSet() {
		localVarQueryParams.Add("$iteration", parameterToString(localVarOptionals.Iteration.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BaseIteration.IsSet() {
		localVarQueryParams.Add("$baseIteration", parameterToString(localVarOptionals.BaseIteration.Value(), ""))
	}
	localVarQueryParams.Add("api-version", parameterToString(apiVersion, ""))
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

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
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarReturnValue, localVarHttpResponse, err
			}

			var v []GitPullRequestCommentThread
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
PullRequestThreadsApiService
Update a thread in a pull request.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body The thread content that should be updated.
 * @param repositoryId The repository ID of the pull request&#39;s target branch.
 * @param pullRequestId ID of the pull request.
 * @param threadId ID of the thread to update.
 * @param project Project ID or project name
 * @param apiVersion Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api.

@return GitPullRequestCommentThread
*/
func (a *PullRequestThreadsApiService) Update(ctx context.Context, body GitPullRequestCommentThread, repositoryId string, pullRequestId int32, threadId int32, project string, apiVersion string) (GitPullRequestCommentThread, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Patch")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue GitPullRequestCommentThread
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/threads/{threadId}"
	localVarPath = strings.Replace(localVarPath, "{"+"repositoryId"+"}", fmt.Sprintf("%v", repositoryId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pullRequestId"+"}", fmt.Sprintf("%v", pullRequestId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"threadId"+"}", fmt.Sprintf("%v", threadId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", fmt.Sprintf("%v", project), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("api-version", parameterToString(apiVersion, ""))
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
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
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarReturnValue, localVarHttpResponse, err
			}

			var v GitPullRequestCommentThread
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
