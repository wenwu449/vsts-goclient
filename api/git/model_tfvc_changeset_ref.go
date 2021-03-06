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
type TfvcChangesetRef struct {
	// A collection of REST reference links.
	Links *ReferenceLinks `json:"_links,omitempty"`
	// Alias or display name of user
	Author *IdentityRef `json:"author,omitempty"`
	// Id of the changeset.
	ChangesetId int32 `json:"changesetId,omitempty"`
	// Alias or display name of user
	CheckedInBy *IdentityRef `json:"checkedInBy,omitempty"`
	// Comment for the changeset.
	Comment string `json:"comment,omitempty"`
	// Was the Comment result truncated?
	CommentTruncated bool `json:"commentTruncated,omitempty"`
	// Creation date of the changeset.
	CreatedDate time.Time `json:"createdDate,omitempty"`
	// URL to retrieve the item.
	Url string `json:"url,omitempty"`
}
