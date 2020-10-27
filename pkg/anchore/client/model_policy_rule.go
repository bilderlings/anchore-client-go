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
// PolicyRule A rule that defines and decision value if the match is found true for a given image.
type PolicyRule struct {
	Action string `json:"action"`
	Gate string `json:"gate"`
	Id string `json:"id,omitempty"`
	Params []PolicyRuleParams `json:"params,omitempty"`
	Trigger string `json:"trigger"`
}
