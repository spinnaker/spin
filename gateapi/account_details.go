/*
 * Spinnaker API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type AccountDetails struct {

	Environment string `json:"environment,omitempty"`

	Name string `json:"name,omitempty"`

	PrimaryAccount bool `json:"primaryAccount,omitempty"`

	RequiredGroupMembership []string `json:"requiredGroupMembership,omitempty"`

	Permissions map[string][]string `json:"permissions,omitempty"`

	CloudProvider string `json:"cloudProvider,omitempty"`

	AccountId string `json:"accountId,omitempty"`

	ChallengeDestructiveActions bool `json:"challengeDestructiveActions,omitempty"`

	Type_ string `json:"type,omitempty"`

	Skin string `json:"skin,omitempty"`

	AccountType string `json:"accountType,omitempty"`

	ProviderVersion string `json:"providerVersion,omitempty"`
}
