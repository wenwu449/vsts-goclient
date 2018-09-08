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

// Parameters for creating an import request
type GitImportRequestParameters struct {
	// Option to delete service endpoint when import is done
	DeleteServiceEndpointAfterImportIsDone bool `json:"deleteServiceEndpointAfterImportIsDone,omitempty"`
	// Source for importing git repository
	GitSource *GitImportGitSource `json:"gitSource,omitempty"`
	// Service Endpoint for connection to external endpoint
	ServiceEndpointId string `json:"serviceEndpointId,omitempty"`
	// Source for importing tfvc repository
	TfvcSource *GitImportTfvcSource `json:"tfvcSource,omitempty"`
}