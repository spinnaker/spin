/*
 * Spinnaker API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type OverrideVerificationRequest struct {
	ArtifactVersion string `json:"artifactVersion,omitempty"`
	VerificationId string `json:"verificationId,omitempty"`
	ArtifactReference string `json:"artifactReference,omitempty"`
	Comment string `json:"comment,omitempty"`
	Status string `json:"status,omitempty"`
}