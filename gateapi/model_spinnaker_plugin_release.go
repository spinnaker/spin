/*
 * Spinnaker API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

type SpinnakerPluginRelease struct {
	Url              string            `json:"url,omitempty"`
	Date             time.Time         `json:"date,omitempty"`
	Sha512sum        string            `json:"sha512sum,omitempty"`
	Requires         string            `json:"requires,omitempty"`
	RemoteExtensions []RemoteExtension `json:"remoteExtensions,omitempty"`
	Preferred        bool              `json:"preferred,omitempty"`
	Version          string            `json:"version,omitempty"`
}
