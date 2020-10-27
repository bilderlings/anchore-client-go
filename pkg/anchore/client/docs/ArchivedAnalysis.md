# ArchivedAnalysis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalyzedAt** | [**time.Time**](time.Time.md) |  | [optional] 
**Annotations** | [**map[string]interface{}**](.md) | User provided annotations as key-value pairs | [optional] 
**ArchiveSizeBytes** | **int32** | The size, in bytes, of the analysis archive file | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] 
**ImageDigest** | **string** | The image digest (digest of the manifest describing the image, per docker spec) | [optional] 
**ImageDetail** | [**[]TagEntry**](TagEntry.md) | List of tags associated with the image digest | [optional] 
**LastUpdated** | [**time.Time**](time.Time.md) |  | [optional] 
**ParentDigest** | **string** | The digest of a parent manifest (for manifest-list images) | [optional] 
**Status** | **string** | The archival status | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


