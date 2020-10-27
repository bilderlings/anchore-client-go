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
// EventSubcategory A collection of events related to each other
type EventSubcategory struct {
	Description string `json:"description,omitempty"`
	Events []EventDescription `json:"events,omitempty"`
	Name string `json:"name,omitempty"`
}
