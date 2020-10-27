# MappingRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] 
**Image** | [**ImageRef**](ImageRef.md) |  | 
**Name** | **string** |  | 
**PolicyId** | **string** | Optional single policy to evalute, if set will override any value in policy_ids, for backwards compatibility. Generally, policy_ids should be used even with a array of length 1. | [optional] 
**PolicyIds** | **[]string** | List of policyIds to evaluate in order, to completion | [optional] 
**Registry** | **string** |  | 
**Repository** | **string** |  | 
**WhitelistIds** | **[]string** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


