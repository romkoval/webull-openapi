# PostLoginParametersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** | username for your account | [optional] 
**AccountType** | **float32** | 2 is email, 1 is phone | [optional] [default to 2]
**DeviceId** | **string** |  | [optional] 
**DeviceName** | **string** | device name | [optional] [default to test]
**ExtInfo** | [**PostLoginParametersRequestExtInfo**](PostLoginParametersRequest_extInfo.md) |  | [optional] 
**Grade** | **float32** |  | [optional] [default to 1]
**Pwd** | **string** | pwd &#x3D; md5(passwordHash + password) | [optional] 
**RegionId** | **float32** |  | [optional] [default to 1]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


