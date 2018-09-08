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

// Context used while sharing a pull request.
type ShareNotificationContext struct {
	// Optional user note or message.
	Message string `json:"message,omitempty"`
	// Identities of users who will receive a share notification.
	Receivers []IdentityRef `json:"receivers,omitempty"`
}
