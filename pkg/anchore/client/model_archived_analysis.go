/*
 * Anchore Engine API Server
 *
 * This is the Anchore Engine API. Provides the primary external API for users of the service.
 *
 * API version: 0.1.15
 * Contact: nurmi@anchore.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client
import (
	"time"
)
// ArchivedAnalysis struct for ArchivedAnalysis
type ArchivedAnalysis struct {
	AnalyzedAt time.Time `json:"analyzed_at,omitempty"`
	// User provided annotations as key-value pairs
	Annotations map[string]interface{} `json:"annotations,omitempty"`
	// The size, in bytes, of the analysis archive file
	ArchiveSizeBytes int32 `json:"archive_size_bytes,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	// The image digest (digest of the manifest describing the image, per docker spec)
	ImageDigest string `json:"imageDigest,omitempty"`
	// List of tags associated with the image digest
	ImageDetail []TagEntry `json:"image_detail,omitempty"`
	LastUpdated time.Time `json:"last_updated,omitempty"`
	// The digest of a parent manifest (for manifest-list images)
	ParentDigest string `json:"parentDigest,omitempty"`
	// The archival status
	Status string `json:"status,omitempty"`
}
