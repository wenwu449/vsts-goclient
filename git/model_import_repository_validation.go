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

// 
type ImportRepositoryValidation struct {
	GitSource *GitImportGitSource `json:"gitSource,omitempty"`
	Password string `json:"password,omitempty"`
	TfvcSource *GitImportTfvcSource `json:"tfvcSource,omitempty"`
	Username string `json:"username,omitempty"`
}
