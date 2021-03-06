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
type GitPullRequestReviewFileContentInfo struct {
	Links *ReferenceLinks `json:"_links,omitempty"`
	// The file change path.
	Path string `json:"path,omitempty"`
	// Content hash of on-disk representation of file content. Its calculated by the client by using SHA1 hash function. Ensure that uploaded file has same encoding as in source control.
	SHA1Hash string `json:"sHA1Hash,omitempty"`
}
