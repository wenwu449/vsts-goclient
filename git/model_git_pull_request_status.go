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

// This class contains the metadata of a service/extension posting pull request status. Status can be associated with a pull request or an iteration.
type GitPullRequestStatus struct {
	// Reference links.
	Links *ReferenceLinks `json:"_links,omitempty"`
	// Context of the status.
	Context *GitStatusContext `json:"context,omitempty"`
	// Identity that created the status.
	CreatedBy *IdentityRef `json:"createdBy,omitempty"`
	// Creation date and time of the status.
	CreationDate time.Time `json:"creationDate,omitempty"`
	// Status description. Typically describes current state of the status.
	Description string `json:"description,omitempty"`
	// Status identifier.
	Id int32 `json:"id,omitempty"`
	// State of the status.
	State string `json:"state,omitempty"`
	// URL with status details.
	TargetUrl string `json:"targetUrl,omitempty"`
	// Last update date and time of the status.
	UpdatedDate time.Time `json:"updatedDate,omitempty"`
}
