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
// AnalysisArchiveRulesSummary Summary of the transition rule set
type AnalysisArchiveRulesSummary struct {
	// The number of rules for this account
	Count int32 `json:"count,omitempty"`
	// The newest last_updated timestamp from the set of rules
	LastUpdated time.Time `json:"last_updated,omitempty"`
}
