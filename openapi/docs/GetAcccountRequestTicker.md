# GetAcccountRequestTicker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyCode** | Pointer to **string** |  | [optional] 
**CurrencyId** | Pointer to **int32** |  | [optional] 
**DisExchangeCode** | Pointer to **string** |  | [optional] 
**DisSymbol** | Pointer to **string** |  | [optional] 
**ExchangeCode** | Pointer to **string** |  | [optional] 
**ExchangeId** | Pointer to **int32** |  | [optional] 
**ExtType** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ListStatus** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**RegionId** | Pointer to **int32** |  | [optional] 
**RegionIsoCode** | Pointer to **string** |  | [optional] 
**RegionName** | Pointer to **string** |  | [optional] 
**SecType** | Pointer to **[]int32** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**TickerId** | Pointer to **int32** |  | [optional] 
**TinyName** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetAcccountRequestTicker

`func NewGetAcccountRequestTicker() *GetAcccountRequestTicker`

NewGetAcccountRequestTicker instantiates a new GetAcccountRequestTicker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAcccountRequestTickerWithDefaults

`func NewGetAcccountRequestTickerWithDefaults() *GetAcccountRequestTicker`

NewGetAcccountRequestTickerWithDefaults instantiates a new GetAcccountRequestTicker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencyCode

`func (o *GetAcccountRequestTicker) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GetAcccountRequestTicker) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GetAcccountRequestTicker) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *GetAcccountRequestTicker) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetCurrencyId

`func (o *GetAcccountRequestTicker) GetCurrencyId() int32`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *GetAcccountRequestTicker) GetCurrencyIdOk() (*int32, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *GetAcccountRequestTicker) SetCurrencyId(v int32)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *GetAcccountRequestTicker) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetDisExchangeCode

`func (o *GetAcccountRequestTicker) GetDisExchangeCode() string`

GetDisExchangeCode returns the DisExchangeCode field if non-nil, zero value otherwise.

### GetDisExchangeCodeOk

`func (o *GetAcccountRequestTicker) GetDisExchangeCodeOk() (*string, bool)`

GetDisExchangeCodeOk returns a tuple with the DisExchangeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisExchangeCode

`func (o *GetAcccountRequestTicker) SetDisExchangeCode(v string)`

SetDisExchangeCode sets DisExchangeCode field to given value.

### HasDisExchangeCode

`func (o *GetAcccountRequestTicker) HasDisExchangeCode() bool`

HasDisExchangeCode returns a boolean if a field has been set.

### GetDisSymbol

`func (o *GetAcccountRequestTicker) GetDisSymbol() string`

GetDisSymbol returns the DisSymbol field if non-nil, zero value otherwise.

### GetDisSymbolOk

`func (o *GetAcccountRequestTicker) GetDisSymbolOk() (*string, bool)`

GetDisSymbolOk returns a tuple with the DisSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisSymbol

`func (o *GetAcccountRequestTicker) SetDisSymbol(v string)`

SetDisSymbol sets DisSymbol field to given value.

### HasDisSymbol

`func (o *GetAcccountRequestTicker) HasDisSymbol() bool`

HasDisSymbol returns a boolean if a field has been set.

### GetExchangeCode

`func (o *GetAcccountRequestTicker) GetExchangeCode() string`

GetExchangeCode returns the ExchangeCode field if non-nil, zero value otherwise.

### GetExchangeCodeOk

`func (o *GetAcccountRequestTicker) GetExchangeCodeOk() (*string, bool)`

GetExchangeCodeOk returns a tuple with the ExchangeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeCode

`func (o *GetAcccountRequestTicker) SetExchangeCode(v string)`

SetExchangeCode sets ExchangeCode field to given value.

### HasExchangeCode

`func (o *GetAcccountRequestTicker) HasExchangeCode() bool`

HasExchangeCode returns a boolean if a field has been set.

### GetExchangeId

`func (o *GetAcccountRequestTicker) GetExchangeId() int32`

GetExchangeId returns the ExchangeId field if non-nil, zero value otherwise.

### GetExchangeIdOk

`func (o *GetAcccountRequestTicker) GetExchangeIdOk() (*int32, bool)`

GetExchangeIdOk returns a tuple with the ExchangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeId

`func (o *GetAcccountRequestTicker) SetExchangeId(v int32)`

SetExchangeId sets ExchangeId field to given value.

### HasExchangeId

`func (o *GetAcccountRequestTicker) HasExchangeId() bool`

HasExchangeId returns a boolean if a field has been set.

### GetExtType

`func (o *GetAcccountRequestTicker) GetExtType() []map[string]interface{}`

GetExtType returns the ExtType field if non-nil, zero value otherwise.

### GetExtTypeOk

`func (o *GetAcccountRequestTicker) GetExtTypeOk() (*[]map[string]interface{}, bool)`

GetExtTypeOk returns a tuple with the ExtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtType

`func (o *GetAcccountRequestTicker) SetExtType(v []map[string]interface{})`

SetExtType sets ExtType field to given value.

### HasExtType

`func (o *GetAcccountRequestTicker) HasExtType() bool`

HasExtType returns a boolean if a field has been set.

### GetListStatus

`func (o *GetAcccountRequestTicker) GetListStatus() int32`

GetListStatus returns the ListStatus field if non-nil, zero value otherwise.

### GetListStatusOk

`func (o *GetAcccountRequestTicker) GetListStatusOk() (*int32, bool)`

GetListStatusOk returns a tuple with the ListStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListStatus

`func (o *GetAcccountRequestTicker) SetListStatus(v int32)`

SetListStatus sets ListStatus field to given value.

### HasListStatus

`func (o *GetAcccountRequestTicker) HasListStatus() bool`

HasListStatus returns a boolean if a field has been set.

### GetName

`func (o *GetAcccountRequestTicker) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetAcccountRequestTicker) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetAcccountRequestTicker) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetAcccountRequestTicker) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegionId

`func (o *GetAcccountRequestTicker) GetRegionId() int32`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *GetAcccountRequestTicker) GetRegionIdOk() (*int32, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *GetAcccountRequestTicker) SetRegionId(v int32)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *GetAcccountRequestTicker) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### GetRegionIsoCode

`func (o *GetAcccountRequestTicker) GetRegionIsoCode() string`

GetRegionIsoCode returns the RegionIsoCode field if non-nil, zero value otherwise.

### GetRegionIsoCodeOk

`func (o *GetAcccountRequestTicker) GetRegionIsoCodeOk() (*string, bool)`

GetRegionIsoCodeOk returns a tuple with the RegionIsoCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionIsoCode

`func (o *GetAcccountRequestTicker) SetRegionIsoCode(v string)`

SetRegionIsoCode sets RegionIsoCode field to given value.

### HasRegionIsoCode

`func (o *GetAcccountRequestTicker) HasRegionIsoCode() bool`

HasRegionIsoCode returns a boolean if a field has been set.

### GetRegionName

`func (o *GetAcccountRequestTicker) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *GetAcccountRequestTicker) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *GetAcccountRequestTicker) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *GetAcccountRequestTicker) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.

### GetSecType

`func (o *GetAcccountRequestTicker) GetSecType() []int32`

GetSecType returns the SecType field if non-nil, zero value otherwise.

### GetSecTypeOk

`func (o *GetAcccountRequestTicker) GetSecTypeOk() (*[]int32, bool)`

GetSecTypeOk returns a tuple with the SecType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecType

`func (o *GetAcccountRequestTicker) SetSecType(v []int32)`

SetSecType sets SecType field to given value.

### HasSecType

`func (o *GetAcccountRequestTicker) HasSecType() bool`

HasSecType returns a boolean if a field has been set.

### GetSymbol

`func (o *GetAcccountRequestTicker) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetAcccountRequestTicker) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetAcccountRequestTicker) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *GetAcccountRequestTicker) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTickerId

`func (o *GetAcccountRequestTicker) GetTickerId() int32`

GetTickerId returns the TickerId field if non-nil, zero value otherwise.

### GetTickerIdOk

`func (o *GetAcccountRequestTicker) GetTickerIdOk() (*int32, bool)`

GetTickerIdOk returns a tuple with the TickerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickerId

`func (o *GetAcccountRequestTicker) SetTickerId(v int32)`

SetTickerId sets TickerId field to given value.

### HasTickerId

`func (o *GetAcccountRequestTicker) HasTickerId() bool`

HasTickerId returns a boolean if a field has been set.

### GetTinyName

`func (o *GetAcccountRequestTicker) GetTinyName() string`

GetTinyName returns the TinyName field if non-nil, zero value otherwise.

### GetTinyNameOk

`func (o *GetAcccountRequestTicker) GetTinyNameOk() (*string, bool)`

GetTinyNameOk returns a tuple with the TinyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTinyName

`func (o *GetAcccountRequestTicker) SetTinyName(v string)`

SetTinyName sets TinyName field to given value.

### HasTinyName

`func (o *GetAcccountRequestTicker) HasTinyName() bool`

HasTinyName returns a boolean if a field has been set.

### GetType

`func (o *GetAcccountRequestTicker) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetAcccountRequestTicker) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetAcccountRequestTicker) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *GetAcccountRequestTicker) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


