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
type GitTreeRef struct {
	Links *ReferenceLinks `json:"_links,omitempty"`
	// SHA1 hash of git object
	ObjectId string `json:"objectId,omitempty"`
	// Sum of sizes of all children
	Size int64 `json:"size,omitempty"`
	// Blobs and trees under this tree
	TreeEntries []GitTreeEntryRef `json:"treeEntries,omitempty"`
	// Url to tree
	Url string `json:"url,omitempty"`
}