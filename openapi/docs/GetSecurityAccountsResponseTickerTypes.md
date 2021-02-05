# GetSecurityAccountsResponseTickerTypes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderTypes** | Pointer to [**[]OrderType**](OrderType.md) |  | [optional] 
**RegionId** | Pointer to **int32** |  | [optional] 
**TickerType** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetSecurityAccountsResponseTickerTypes

`func NewGetSecurityAccountsResponseTickerTypes() *GetSecurityAccountsResponseTickerTypes`

NewGetSecurityAccountsResponseTickerTypes instantiates a new GetSecurityAccountsResponseTickerTypes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSecurityAccountsResponseTickerTypesWithDefaults

`func NewGetSecurityAccountsResponseTickerTypesWithDefaults() *GetSecurityAccountsResponseTickerTypes`

NewGetSecurityAccountsResponseTickerTypesWithDefaults instantiates a new GetSecurityAccountsResponseTickerTypes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderTypes

`func (o *GetSecurityAccountsResponseTickerTypes) GetOrderTypes() []OrderType`

GetOrderTypes returns the OrderTypes field if non-nil, zero value otherwise.

### GetOrderTypesOk

`func (o *GetSecurityAccountsResponseTickerTypes) GetOrderTypesOk() (*[]OrderType, bool)`

GetOrderTypesOk returns a tuple with the OrderTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderTypes

`func (o *GetSecurityAccountsResponseTickerTypes) SetOrderTypes(v []OrderType)`

SetOrderTypes sets OrderTypes field to given value.

### HasOrderTypes

`func (o *GetSecurityAccountsResponseTickerTypes) HasOrderTypes() bool`

HasOrderTypes returns a boolean if a field has been set.

### GetRegionId

`func (o *GetSecurityAccountsResponseTickerTypes) GetRegionId() int32`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *GetSecurityAccountsResponseTickerTypes) GetRegionIdOk() (*int32, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *GetSecurityAccountsResponseTickerTypes) SetRegionId(v int32)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *GetSecurityAccountsResponseTickerTypes) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### GetTickerType

`func (o *GetSecurityAccountsResponseTickerTypes) GetTickerType() int32`

GetTickerType returns the TickerType field if non-nil, zero value otherwise.

### GetTickerTypeOk

`func (o *GetSecurityAccountsResponseTickerTypes) GetTickerTypeOk() (*int32, bool)`

GetTickerTypeOk returns a tuple with the TickerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickerType

`func (o *GetSecurityAccountsResponseTickerTypes) SetTickerType(v int32)`

SetTickerType sets TickerType field to given value.

### HasTickerType

`func (o *GetSecurityAccountsResponseTickerTypes) HasTickerType() bool`

HasTickerType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


