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

// Parameters that are provided in the request body when requesting to cherry pick or revert.
type GitAsyncRefOperationParameters struct {
	// Proposed target branch name for the cherry pick or revert operation.
	GeneratedRefName string `json:"generatedRefName,omitempty"`
	// The target branch for the cherry pick or revert operation.
	OntoRefName string `json:"ontoRefName,omitempty"`
	// The git repository for the cherry pick or revert operation.
	Repository *GitRepository `json:"repository,omitempty"`
	// Details about the source of the cherry pick or revert operation (e.g. A pull request or a specific commit).
	Source *GitAsyncRefOperationSource `json:"source,omitempty"`
}
