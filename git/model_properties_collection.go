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

// The class represents a property bag as a collection of key-value pairs. Values of all primitive types (any type with a `TypeCode != TypeCode.Object`) except for `DBNull` are accepted. Values of type Byte[], Int32, Double, DateType and String preserve their type, other primitives are retuned as a String. Byte[] expected as base64 encoded string.
type PropertiesCollection struct {
	// The count of properties in the collection.
	Count int32 `json:"count,omitempty"`
	Item *interface{} `json:"item,omitempty"`
	// The set of keys in the collection.
	Keys []string `json:"keys,omitempty"`
	// The set of values in the collection.
	Values []string `json:"values,omitempty"`
}
