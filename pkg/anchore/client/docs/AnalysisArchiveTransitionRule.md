# AnalysisArchiveTransitionRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalysisAgeDays** | **int32** | Matches if the analysis is strictly older than this number of days | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] 
**LastUpdated** | [**time.Time**](time.Time.md) |  | [optional] 
**Selector** | [**ImageSelector**](ImageSelector.md) |  | [optional] 
**SystemGlobal** | **bool** | True if the rule applies to all accounts in the system. This is only available to admin users to update/modify, but all users with permission to list rules can see them | [optional] 
**TagVersionsNewer** | **int32** | Number of images mapped to the tag that are newer | [optional] 
**Transition** | **string** | The type of transition to make. If \&quot;archive\&quot;, then archive an image from the working set and remove it from the working set. If \&quot;delete\&quot;, then match against archived images and delete from the archive if match. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


