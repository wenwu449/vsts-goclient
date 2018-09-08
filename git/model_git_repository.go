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

// 
type GitRepository struct {
	Links *ReferenceLinks `json:"_links,omitempty"`
	DefaultBranch string `json:"defaultBranch,omitempty"`
	Id string `json:"id,omitempty"`
	// True if the repository was created as a fork
	IsFork bool `json:"isFork,omitempty"`
	Name string `json:"name,omitempty"`
	ParentRepository *GitRepositoryRef `json:"parentRepository,omitempty"`
	Project *TeamProjectReference `json:"project,omitempty"`
	RemoteUrl string `json:"remoteUrl,omitempty"`
	// Compressed size (bytes) of the repository.
	Size int64 `json:"size,omitempty"`
	SshUrl string `json:"sshUrl,omitempty"`
	Url string `json:"url,omitempty"`
	ValidRemoteUrls []string `json:"validRemoteUrls,omitempty"`
}
