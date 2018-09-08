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

// Real time event (SignalR) for a status update on a pull request
type StatusUpdatedEvent struct {
	// The id of this event. Can be used to track send/receive state between client and server.
	EventId string `json:"eventId,omitempty"`
	// The id of the pull request this event was generated for.
	PullRequestId int32 `json:"pullRequestId,omitempty"`
}
