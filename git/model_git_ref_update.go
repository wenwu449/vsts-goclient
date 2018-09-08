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
type GitRefUpdate struct {
	IsLocked bool `json:"isLocked,omitempty"`
	Name string `json:"name,omitempty"`
	NewObjectId string `json:"newObjectId,omitempty"`
	OldObjectId string `json:"oldObjectId,omitempty"`
	RepositoryId string `json:"repositoryId,omitempty"`
}
