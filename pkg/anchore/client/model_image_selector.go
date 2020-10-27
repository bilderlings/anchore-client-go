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
// ImageSelector A set of selection criteria to match an image by a tagged pullstring based on its components, with regex support in each field
type ImageSelector struct {
	// The registry section of a pull string. e.g. with \"docker.io/anchore/anchore-engine:latest\", this is \"docker.io\"
	Registry string `json:"registry,omitempty"`
	// The repository section of a pull string. e.g. with \"docker.io/anchore/anchore-engine:latest\", this is \"anchore/anchore-engine\"
	Repository string `json:"repository,omitempty"`
	// The tag-only section of a pull string. e.g. with \"docker.io/anchore/anchore-engine:latest\", this is \"latest\"
	Tag string `json:"tag,omitempty"`
}
