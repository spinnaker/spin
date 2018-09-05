/*
 * Spinnaker API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Account struct {

	RequiredGroupMembership []string `json:"requiredGroupMembership,omitempty"`

	Skin string `json:"skin,omitempty"`

	Permissions map[string][]string `json:"permissions,omitempty"`

	AccountId string `json:"accountId,omitempty"`

	Name string `json:"name,omitempty"`

	ProviderVersion string `json:"providerVersion,omitempty"`

	Type_ string `json:"type,omitempty"`
}
