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

// Policy configuration reference.
type PolicyConfigurationRef struct {
	// The policy configuration ID.
	Id int32 `json:"id,omitempty"`
	// The policy configuration type.
	Type_ *PolicyTypeRef `json:"type,omitempty"`
	// The URL where the policy configuration can be retrieved.
	Url string `json:"url,omitempty"`
}
