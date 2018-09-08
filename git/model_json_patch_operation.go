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

// The JSON model for a JSON Patch operation
type JsonPatchOperation struct {
	// The path to copy from for the Move/Copy operation.
	From string `json:"from,omitempty"`
	// The patch operation
	Op string `json:"op,omitempty"`
	// The path for the operation
	Path string `json:"path,omitempty"`
	// The value for the operation. This is either a primitive or a JToken.
	Value *interface{} `json:"value,omitempty"`
}
