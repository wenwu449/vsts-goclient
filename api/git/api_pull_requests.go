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

type PullRequestsApiService service

/*
PullRequestsApiService
Create a pull request.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body The pull request to create.
 * @param repositoryId The repository ID of the pull request&#39;s target branch.
 * @param project Project ID or project name
 * @param apiVersion Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api.
 * @param optional nil or *CreatePullRequestOpts - Optional Parameters:
     * @param "SupportsIterations" (optional.Bool) -  If true, subsequent pushes to the pull request will be individually reviewable. Set this to false for large pull requests for performance reasons if this functionality is not needed.

@return GitPullRequest
*/

type CreatePullRequestOpts struct {
	SupportsIterations optional.Bool
}

func (a *PullRequestsApiService) Create(ctx context.Context, body GitPullRequest, repositoryId string, project string, apiVersion string, localVarOptionals *CreatePullRequestOpts) (GitPullRequest, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue GitPullRequest
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/{project}/_apis/git/repositories/{repositoryId}/pullrequests"
	localVarPath = strings.Replace(localVarPath, "{"+"repositoryId"+"}", fmt.Sprintf("%v", repositoryId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", fmt.Sprintf("%v", project), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.SupportsIterations.IsSet() {
		localVarQueryParams.Add("supportsIterations", parameterToString(localVarOptionals.SupportsIterations.Value(), ""))
	}
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

			var v GitPullRequest
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
PullRequestsApiService
Retrieve a pull request.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param repositoryId The repository ID of the pull request&#39;s target branch.
 * @param pullRequestId The ID of the pull request to retrieve.
 * @param project Project ID or project name
 * @param apiVersion Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api.
 * @param optional nil or *GetPullRequestOpts - Optional Parameters:
     * @param "MaxCommentLength" (optional.Int32) -  Not used.
     * @param "Skip" (optional.Int32) -  Not used.
     * @param "Top" (optional.Int32) -  Not used.
     * @param "IncludeCommits" (optional.Bool) -  If true, the pull request will be returned with the associated commits.
     * @param "IncludeWorkItemRefs" (optional.Bool) -  If true, the pull request will be returned with the associated work item references.

@return GitPullRequest
*/

type GetPullRequestOpts struct {
	MaxCommentLength    optional.Int32
	Skip                optional.Int32
	Top                 optional.Int32
	IncludeCommits      optional.Bool
	IncludeWorkItemRefs optional.Bool
}

func (a *PullRequestsApiService) GetPullRequest(ctx context.Context, repositoryId string, pullRequestId int32, project string, apiVersion string, localVarOptionals *GetPullRequestOpts) (GitPullRequest, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue GitPullRequest
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/{project}/_apis/git/repositories/{repositoryId}/pullrequests/{pullRequestId}"
	localVarPath = strings.Replace(localVarPath, "{"+"repositoryId"+"}", fmt.Sprintf("%v", repositoryId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pullRequestId"+"}", fmt.Sprintf("%v", pullRequestId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", fmt.Sprintf("%v", project), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.MaxCommentLength.IsSet() {
		localVarQueryParams.Add("maxCommentLength", parameterToString(localVarOptionals.MaxCommentLength.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Skip.IsSet() {
		localVarQueryParams.Add("$skip", parameterToString(localVarOptionals.Skip.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Top.IsSet() {
		localVarQueryParams.Add("$top", parameterToString(localVarOptionals.Top.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IncludeCommits.IsSet() {
		localVarQueryParams.Add("includeCommits", parameterToString(localVarOptionals.IncludeCommits.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IncludeWorkItemRefs.IsSet() {
		localVarQueryParams.Add("includeWorkItemRefs", parameterToString(localVarOptionals.IncludeWorkItemRefs.Value(), ""))
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

			var v GitPullRequest
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
PullRequestsApiService
Retrieve a pull request.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pullRequestId The ID of the pull request to retrieve.
 * @param project Project ID or project name
 * @param apiVersion Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api.

@return GitPullRequest
*/
func (a *PullRequestsApiService) GetPullRequestById(ctx context.Context, pullRequestId int32, project string, apiVersion string) (GitPullRequest, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue GitPullRequest
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/{project}/_apis/git/pullrequests/{pullRequestId}"
	localVarPath = strings.Replace(localVarPath, "{"+"pullRequestId"+"}", fmt.Sprintf("%v", pullRequestId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", fmt.Sprintf("%v", project), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

			var v GitPullRequest
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
PullRequestsApiService
Retrieve all pull requests matching a specified criteria.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param repositoryId The repository ID of the pull request&#39;s target branch.
 * @param project Project ID or project name
 * @param apiVersion Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api.
 * @param optional nil or *GetPullRequestsOpts - Optional Parameters:
     * @param "SearchCriteriaIncludeLinks" (optional.Bool) -  Whether to include the _links field on the shallow references
     * @param "SearchCriteriaSourceRefName" (optional.String) -  If set, search for pull requests from this branch.
     * @param "SearchCriteriaSourceRepositoryId" (optional.Interface of string) -  If set, search for pull requests whose source branch is in this repository.
     * @param "SearchCriteriaTargetRefName" (optional.String) -  If set, search for pull requests into this branch.
     * @param "SearchCriteriaStatus" (optional.String) -  If set, search for pull requests that are in this state. Defaults to Active if unset.
     * @param "SearchCriteriaReviewerId" (optional.Interface of string) -  If set, search for pull requests that have this identity as a reviewer.
     * @param "SearchCriteriaCreatorId" (optional.Interface of string) -  If set, search for pull requests that were created by this identity.
     * @param "SearchCriteriaRepositoryId" (optional.Interface of string) -  If set, search for pull requests whose target branch is in this repository.
     * @param "MaxCommentLength" (optional.Int32) -  Not used.
     * @param "Skip" (optional.Int32) -  The number of pull requests to ignore. For example, to retrieve results 101-150, set top to 50 and skip to 100.
     * @param "Top" (optional.Int32) -  The number of pull requests to retrieve.

@return []GitPullRequest
*/

type GetPullRequestsOpts struct {
	SearchCriteriaIncludeLinks       optional.Bool
	SearchCriteriaSourceRefName      optional.String
	SearchCriteriaSourceRepositoryId optional.Interface
	SearchCriteriaTargetRefName      optional.String
	SearchCriteriaStatus             optional.String
	SearchCriteriaReviewerId         optional.Interface
	SearchCriteriaCreatorId          optional.Interface
	SearchCriteriaRepositoryId       optional.Interface
	MaxCommentLength                 optional.Int32
	Skip                             optional.Int32
	Top                              optional.Int32
}

func (a *PullRequestsApiService) GetPullRequests(ctx context.Context, repositoryId string, project string, apiVersion string, localVarOptionals *GetPullRequestsOpts) ([]GitPullRequest, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue []GitPullRequest
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/{project}/_apis/git/repositories/{repositoryId}/pullrequests"
	localVarPath = strings.Replace(localVarPath, "{"+"repositoryId"+"}", fmt.Sprintf("%v", repositoryId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", fmt.Sprintf("%v", project), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.SearchCriteriaIncludeLinks.IsSet() {
		localVarQueryParams.Add("searchCriteria.includeLinks", parameterToString(localVarOptionals.SearchCriteriaIncludeLinks.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SearchCriteriaSourceRefName.IsSet() {
		localVarQueryParams.Add("searchCriteria.sourceRefName", parameterToString(localVarOptionals.SearchCriteriaSourceRefName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SearchCriteriaSourceRepositoryId.IsSet() {
		localVarQueryParams.Add("searchCriteria.sourceRepositoryId", parameterToString(localVarOptionals.SearchCriteriaSourceRepositoryId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SearchCriteriaTargetRefName.IsSet() {
		localVarQueryParams.Add("searchCriteria.targetRefName", parameterToString(localVarOptionals.SearchCriteriaTargetRefName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SearchCriteriaStatus.IsSet() {
		localVarQueryParams.Add("searchCriteria.status", parameterToString(localVarOptionals.SearchCriteriaStatus.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SearchCriteriaReviewerId.IsSet() {
		localVarQueryParams.Add("searchCriteria.reviewerId", parameterToString(localVarOptionals.SearchCriteriaReviewerId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SearchCriteriaCreatorId.IsSet() {
		localVarQueryParams.Add("searchCriteria.creatorId", parameterToString(localVarOptionals.SearchCriteriaCreatorId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SearchCriteriaRepositoryId.IsSet() {
		localVarQueryParams.Add("searchCriteria.repositoryId", parameterToString(localVarOptionals.SearchCriteriaRepositoryId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.MaxCommentLength.IsSet() {
		localVarQueryParams.Add("maxCommentLength", parameterToString(localVarOptionals.MaxCommentLength.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Skip.IsSet() {
		localVarQueryParams.Add("$skip", parameterToString(localVarOptionals.Skip.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Top.IsSet() {
		localVarQueryParams.Add("$top", parameterToString(localVarOptionals.Top.Value(), ""))
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

			var v []GitPullRequest
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
PullRequestsApiService
Retrieve all pull requests matching a specified criteria.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param project Project ID or project name
 * @param apiVersion Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api.
 * @param optional nil or *GetPullRequestsByProjectOpts - Optional Parameters:
     * @param "SearchCriteriaIncludeLinks" (optional.Bool) -  Whether to include the _links field on the shallow references
     * @param "SearchCriteriaSourceRefName" (optional.String) -  If set, search for pull requests from this branch.
     * @param "SearchCriteriaSourceRepositoryId" (optional.Interface of string) -  If set, search for pull requests whose source branch is in this repository.
     * @param "SearchCriteriaTargetRefName" (optional.String) -  If set, search for pull requests into this branch.
     * @param "SearchCriteriaStatus" (optional.String) -  If set, search for pull requests that are in this state. Defaults to Active if unset.
     * @param "SearchCriteriaReviewerId" (optional.Interface of string) -  If set, search for pull requests that have this identity as a reviewer.
     * @param "SearchCriteriaCreatorId" (optional.Interface of string) -  If set, search for pull requests that were created by this identity.
     * @param "SearchCriteriaRepositoryId" (optional.Interface of string) -  If set, search for pull requests whose target branch is in this repository.
     * @param "MaxCommentLength" (optional.Int32) -  Not used.
     * @param "Skip" (optional.Int32) -  The number of pull requests to ignore. For example, to retrieve results 101-150, set top to 50 and skip to 100.
     * @param "Top" (optional.Int32) -  The number of pull requests to retrieve.

@return []GitPullRequest
*/

type GetPullRequestsByProjectOpts struct {
	SearchCriteriaIncludeLinks       optional.Bool
	SearchCriteriaSourceRefName      optional.String
	SearchCriteriaSourceRepositoryId optional.Interface
	SearchCriteriaTargetRefName      optional.String
	SearchCriteriaStatus             optional.String
	SearchCriteriaReviewerId         optional.Interface
	SearchCriteriaCreatorId          optional.Interface
	SearchCriteriaRepositoryId       optional.Interface
	MaxCommentLength                 optional.Int32
	Skip                             optional.Int32
	Top                              optional.Int32
}

func (a *PullRequestsApiService) GetPullRequestsByProject(ctx context.Context, project string, apiVersion string, localVarOptionals *GetPullRequestsByProjectOpts) ([]GitPullRequest, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue []GitPullRequest
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/{project}/_apis/git/pullrequests"
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", fmt.Sprintf("%v", project), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.SearchCriteriaIncludeLinks.IsSet() {
		localVarQueryParams.Add("searchCriteria.includeLinks", parameterToString(localVarOptionals.SearchCriteriaIncludeLinks.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SearchCriteriaSourceRefName.IsSet() {
		localVarQueryParams.Add("searchCriteria.sourceRefName", parameterToString(localVarOptionals.SearchCriteriaSourceRefName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SearchCriteriaSourceRepositoryId.IsSet() {
		localVarQueryParams.Add("searchCriteria.sourceRepositoryId", parameterToString(localVarOptionals.SearchCriteriaSourceRepositoryId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SearchCriteriaTargetRefName.IsSet() {
		localVarQueryParams.Add("searchCriteria.targetRefName", parameterToString(localVarOptionals.SearchCriteriaTargetRefName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SearchCriteriaStatus.IsSet() {
		localVarQueryParams.Add("searchCriteria.status", parameterToString(localVarOptionals.SearchCriteriaStatus.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SearchCriteriaReviewerId.IsSet() {
		localVarQueryParams.Add("searchCriteria.reviewerId", parameterToString(localVarOptionals.SearchCriteriaReviewerId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SearchCriteriaCreatorId.IsSet() {
		localVarQueryParams.Add("searchCriteria.creatorId", parameterToString(localVarOptionals.SearchCriteriaCreatorId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SearchCriteriaRepositoryId.IsSet() {
		localVarQueryParams.Add("searchCriteria.repositoryId", parameterToString(localVarOptionals.SearchCriteriaRepositoryId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.MaxCommentLength.IsSet() {
		localVarQueryParams.Add("maxCommentLength", parameterToString(localVarOptionals.MaxCommentLength.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Skip.IsSet() {
		localVarQueryParams.Add("$skip", parameterToString(localVarOptionals.Skip.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Top.IsSet() {
		localVarQueryParams.Add("$top", parameterToString(localVarOptionals.Top.Value(), ""))
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

			var v []GitPullRequest
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
PullRequestsApiService
Update a pull request.  These are the properties that can be updated with the API:  - Status  - Title  - Description  - CompletionOptions  - MergeOptions  - AutoCompleteSetBy.Id  - TargetRefName (when the PR retargeting feature is enabled)  Attempting to update other properties outside of this list will either cause the server to throw an &#x60;InvalidArgumentValueException&#x60;,  or to silently ignore the update.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body The pull request content to update.
 * @param repositoryId The repository ID of the pull request&#39;s target branch.
 * @param pullRequestId The ID of the pull request to retrieve.
 * @param project Project ID or project name
 * @param apiVersion Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api.

@return GitPullRequest
*/
func (a *PullRequestsApiService) Update(ctx context.Context, body GitPullRequest, repositoryId string, pullRequestId int32, project string, apiVersion string) (GitPullRequest, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Patch")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue GitPullRequest
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/{project}/_apis/git/repositories/{repositoryId}/pullrequests/{pullRequestId}"
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

			var v GitPullRequest
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
