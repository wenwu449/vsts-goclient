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
	"time"
)

// 
type GitLastChangeTreeItems struct {
	// The list of commits referenced by Items, if they were requested.
	Commits []GitCommitRef `json:"commits,omitempty"`
	// The last change of items.
	Items []GitLastChangeItem `json:"items,omitempty"`
	// The last explored time, in case the result is not comprehensive. Null otherwise.
	LastExploredTime time.Time `json:"lastExploredTime,omitempty"`
}
