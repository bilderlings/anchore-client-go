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
// MappingRule struct for MappingRule
type MappingRule struct {
	Id string `json:"id,omitempty"`
	Image ImageRef `json:"image"`
	Name string `json:"name"`
	// Optional single policy to evalute, if set will override any value in policy_ids, for backwards compatibility. Generally, policy_ids should be used even with a array of length 1.
	PolicyId string `json:"policy_id,omitempty"`
	// List of policyIds to evaluate in order, to completion
	PolicyIds []string `json:"policy_ids,omitempty"`
	Registry string `json:"registry"`
	Repository string `json:"repository"`
	WhitelistIds []string `json:"whitelist_ids,omitempty"`
}
