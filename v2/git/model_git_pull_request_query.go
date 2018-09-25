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

// A set of pull request queries and their results.
type GitPullRequestQuery struct {
	// The queries to perform.
	Queries []GitPullRequestQueryInput `json:"queries,omitempty"`
	// The results of the queries. This matches the QueryInputs list so Results[n] are the results of QueryInputs[n]. Each entry in the list is a dictionary of commit->pull requests.
	Results []map[string][]GitPullRequest `json:"results,omitempty"`
}