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

type DiffsApiService service

/*
DiffsApiService
Find the closest common commit (the merge base) between base and target commits, and get the diff between either the base and target commits or common and target commits.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param repositoryId The name or ID of the repository.
 * @param project Project ID or project name
 * @param apiVersion Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api.
 * @param optional nil or *GetDiffsOpts - Optional Parameters:
     * @param "DiffCommonCommit" (optional.Bool) -  If true, diff between common and target commits. If false, diff between base and target commits.
     * @param "Top" (optional.Int32) -  Maximum number of changes to return. Defaults to 100.
     * @param "Skip" (optional.Int32) -  Number of changes to skip
     * @param "BaseVersionOptions" (optional.String) -  Version options - Specify additional modifiers to version (e.g Previous)
     * @param "BaseVersion" (optional.String) -  Version string identifier (name of tag/branch, SHA1 of commit)
     * @param "BaseVersionType" (optional.String) -  Version type (branch, tag, or commit). Determines how Id is interpreted
     * @param "TargetVersionOptions" (optional.String) -  Version options - Specify additional modifiers to version (e.g Previous)
     * @param "TargetVersion" (optional.String) -  Version string identifier (name of tag/branch, SHA1 of commit)
     * @param "TargetVersionType" (optional.String) -  Version type (branch, tag, or commit). Determines how Id is interpreted

@return GitCommitDiffs
*/

type GetDiffsOpts struct {
	DiffCommonCommit     optional.Bool
	Top                  optional.Int32
	Skip                 optional.Int32
	BaseVersionOptions   optional.String
	BaseVersion          optional.String
	BaseVersionType      optional.String
	TargetVersionOptions optional.String
	TargetVersion        optional.String
	TargetVersionType    optional.String
}

func (a *DiffsApiService) Get(ctx context.Context, repositoryId string, project string, apiVersion string, localVarOptionals *GetDiffsOpts) (GitCommitDiffs, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue GitCommitDiffs
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/{project}/_apis/git/repositories/{repositoryId}/diffs/commits"
	localVarPath = strings.Replace(localVarPath, "{"+"repositoryId"+"}", fmt.Sprintf("%v", repositoryId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", fmt.Sprintf("%v", project), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.DiffCommonCommit.IsSet() {
		localVarQueryParams.Add("diffCommonCommit", parameterToString(localVarOptionals.DiffCommonCommit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Top.IsSet() {
		localVarQueryParams.Add("$top", parameterToString(localVarOptionals.Top.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Skip.IsSet() {
		localVarQueryParams.Add("$skip", parameterToString(localVarOptionals.Skip.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BaseVersionOptions.IsSet() {
		localVarQueryParams.Add("baseVersionOptions", parameterToString(localVarOptionals.BaseVersionOptions.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BaseVersion.IsSet() {
		localVarQueryParams.Add("baseVersion", parameterToString(localVarOptionals.BaseVersion.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BaseVersionType.IsSet() {
		localVarQueryParams.Add("baseVersionType", parameterToString(localVarOptionals.BaseVersionType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TargetVersionOptions.IsSet() {
		localVarQueryParams.Add("targetVersionOptions", parameterToString(localVarOptionals.TargetVersionOptions.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TargetVersion.IsSet() {
		localVarQueryParams.Add("targetVersion", parameterToString(localVarOptionals.TargetVersion.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TargetVersionType.IsSet() {
		localVarQueryParams.Add("targetVersionType", parameterToString(localVarOptionals.TargetVersionType.Value(), ""))
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

			var v GitCommitDiffs
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
