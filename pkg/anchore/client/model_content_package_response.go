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
// ContentPackageResponse Package content listings from images
type ContentPackageResponse struct {
	Content []ContentPackageResponseContent `json:"content,omitempty"`
	ContentType string `json:"content_type,omitempty"`
	ImageDigest string `json:"imageDigest,omitempty"`
}
