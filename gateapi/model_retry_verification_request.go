/*
 * Spinnaker API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type RetryVerificationRequest struct {
	VerificationId string `json:"verificationId,omitempty"`
	ArtifactReference string `json:"artifactReference,omitempty"`
	ArtifactVersion string `json:"artifactVersion,omitempty"`
}