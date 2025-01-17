/*
 * HCS API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package hcsschema

// memory usage as viewed from within the container
type ContainerMemoryInformation struct {
	TotalPhysicalBytes int32 `json:"TotalPhysicalBytes,omitempty"`

	TotalUsage int32 `json:"TotalUsage,omitempty"`

	CommittedBytes int32 `json:"CommittedBytes,omitempty"`

	SharedCommittedBytes int32 `json:"SharedCommittedBytes,omitempty"`

	CommitLimitBytes int32 `json:"CommitLimitBytes,omitempty"`

	PeakCommitmentBytes int32 `json:"PeakCommitmentBytes,omitempty"`
}
