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

// Comment thread context contains details about what diffs were being viewed at the time of thread creation and whether or not the thread has been tracked from that original diff.
type GitPullRequestCommentThreadContext struct {
	// Used to track a comment across iterations. This value can be found by looking at the iteration's changes list. Must be set for pull requests with iteration support. Otherwise, it's not required for 'legacy' pull requests.
	ChangeTrackingId int32 `json:"changeTrackingId,omitempty"`
	// The iteration context being viewed when the thread was created.
	IterationContext *CommentIterationContext `json:"iterationContext,omitempty"`
	// The criteria used to track this thread. If this property is filled out when the thread is returned, then the thread has been tracked from its original location using the given criteria.
	TrackingCriteria *CommentTrackingCriteria `json:"trackingCriteria,omitempty"`
}
