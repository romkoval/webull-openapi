# GetOrdersItemTicker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyCode** | Pointer to **string** |  | [optional] 
**CurrencyId** | Pointer to **int32** |  | [optional] 
**DisExchangeCode** | Pointer to **string** |  | [optional] 
**DisSymbol** | Pointer to **string** |  | [optional] 
**ExchangeCode** | Pointer to **string** |  | [optional] 
**ExchangeId** | Pointer to **int32** |  | [optional] 
**ExchangeTrade** | Pointer to **bool** |  | [optional] 
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

### NewGetOrdersItemTicker

`func NewGetOrdersItemTicker() *GetOrdersItemTicker`

NewGetOrdersItemTicker instantiates a new GetOrdersItemTicker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrdersItemTickerWithDefaults

`func NewGetOrdersItemTickerWithDefaults() *GetOrdersItemTicker`

NewGetOrdersItemTickerWithDefaults instantiates a new GetOrdersItemTicker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencyCode

`func (o *GetOrdersItemTicker) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GetOrdersItemTicker) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GetOrdersItemTicker) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *GetOrdersItemTicker) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetCurrencyId

`func (o *GetOrdersItemTicker) GetCurrencyId() int32`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *GetOrdersItemTicker) GetCurrencyIdOk() (*int32, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *GetOrdersItemTicker) SetCurrencyId(v int32)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *GetOrdersItemTicker) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetDisExchangeCode

`func (o *GetOrdersItemTicker) GetDisExchangeCode() string`

GetDisExchangeCode returns the DisExchangeCode field if non-nil, zero value otherwise.

### GetDisExchangeCodeOk

`func (o *GetOrdersItemTicker) GetDisExchangeCodeOk() (*string, bool)`

GetDisExchangeCodeOk returns a tuple with the DisExchangeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisExchangeCode

`func (o *GetOrdersItemTicker) SetDisExchangeCode(v string)`

SetDisExchangeCode sets DisExchangeCode field to given value.

### HasDisExchangeCode

`func (o *GetOrdersItemTicker) HasDisExchangeCode() bool`

HasDisExchangeCode returns a boolean if a field has been set.

### GetDisSymbol

`func (o *GetOrdersItemTicker) GetDisSymbol() string`

GetDisSymbol returns the DisSymbol field if non-nil, zero value otherwise.

### GetDisSymbolOk

`func (o *GetOrdersItemTicker) GetDisSymbolOk() (*string, bool)`

GetDisSymbolOk returns a tuple with the DisSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisSymbol

`func (o *GetOrdersItemTicker) SetDisSymbol(v string)`

SetDisSymbol sets DisSymbol field to given value.

### HasDisSymbol

`func (o *GetOrdersItemTicker) HasDisSymbol() bool`

HasDisSymbol returns a boolean if a field has been set.

### GetExchangeCode

`func (o *GetOrdersItemTicker) GetExchangeCode() string`

GetExchangeCode returns the ExchangeCode field if non-nil, zero value otherwise.

### GetExchangeCodeOk

`func (o *GetOrdersItemTicker) GetExchangeCodeOk() (*string, bool)`

GetExchangeCodeOk returns a tuple with the ExchangeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeCode

`func (o *GetOrdersItemTicker) SetExchangeCode(v string)`

SetExchangeCode sets ExchangeCode field to given value.

### HasExchangeCode

`func (o *GetOrdersItemTicker) HasExchangeCode() bool`

HasExchangeCode returns a boolean if a field has been set.

### GetExchangeId

`func (o *GetOrdersItemTicker) GetExchangeId() int32`

GetExchangeId returns the ExchangeId field if non-nil, zero value otherwise.

### GetExchangeIdOk

`func (o *GetOrdersItemTicker) GetExchangeIdOk() (*int32, bool)`

GetExchangeIdOk returns a tuple with the ExchangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeId

`func (o *GetOrdersItemTicker) SetExchangeId(v int32)`

SetExchangeId sets ExchangeId field to given value.

### HasExchangeId

`func (o *GetOrdersItemTicker) HasExchangeId() bool`

HasExchangeId returns a boolean if a field has been set.

### GetExchangeTrade

`func (o *GetOrdersItemTicker) GetExchangeTrade() bool`

GetExchangeTrade returns the ExchangeTrade field if non-nil, zero value otherwise.

### GetExchangeTradeOk

`func (o *GetOrdersItemTicker) GetExchangeTradeOk() (*bool, bool)`

GetExchangeTradeOk returns a tuple with the ExchangeTrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeTrade

`func (o *GetOrdersItemTicker) SetExchangeTrade(v bool)`

SetExchangeTrade sets ExchangeTrade field to given value.

### HasExchangeTrade

`func (o *GetOrdersItemTicker) HasExchangeTrade() bool`

HasExchangeTrade returns a boolean if a field has been set.

### GetExtType

`func (o *GetOrdersItemTicker) GetExtType() []map[string]interface{}`

GetExtType returns the ExtType field if non-nil, zero value otherwise.

### GetExtTypeOk

`func (o *GetOrdersItemTicker) GetExtTypeOk() (*[]map[string]interface{}, bool)`

GetExtTypeOk returns a tuple with the ExtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtType

`func (o *GetOrdersItemTicker) SetExtType(v []map[string]interface{})`

SetExtType sets ExtType field to given value.

### HasExtType

`func (o *GetOrdersItemTicker) HasExtType() bool`

HasExtType returns a boolean if a field has been set.

### GetListStatus

`func (o *GetOrdersItemTicker) GetListStatus() int32`

GetListStatus returns the ListStatus field if non-nil, zero value otherwise.

### GetListStatusOk

`func (o *GetOrdersItemTicker) GetListStatusOk() (*int32, bool)`

GetListStatusOk returns a tuple with the ListStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListStatus

`func (o *GetOrdersItemTicker) SetListStatus(v int32)`

SetListStatus sets ListStatus field to given value.

### HasListStatus

`func (o *GetOrdersItemTicker) HasListStatus() bool`

HasListStatus returns a boolean if a field has been set.

### GetName

`func (o *GetOrdersItemTicker) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetOrdersItemTicker) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetOrdersItemTicker) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetOrdersItemTicker) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegionId

`func (o *GetOrdersItemTicker) GetRegionId() int32`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *GetOrdersItemTicker) GetRegionIdOk() (*int32, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *GetOrdersItemTicker) SetRegionId(v int32)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *GetOrdersItemTicker) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### GetRegionIsoCode

`func (o *GetOrdersItemTicker) GetRegionIsoCode() string`

GetRegionIsoCode returns the RegionIsoCode field if non-nil, zero value otherwise.

### GetRegionIsoCodeOk

`func (o *GetOrdersItemTicker) GetRegionIsoCodeOk() (*string, bool)`

GetRegionIsoCodeOk returns a tuple with the RegionIsoCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionIsoCode

`func (o *GetOrdersItemTicker) SetRegionIsoCode(v string)`

SetRegionIsoCode sets RegionIsoCode field to given value.

### HasRegionIsoCode

`func (o *GetOrdersItemTicker) HasRegionIsoCode() bool`

HasRegionIsoCode returns a boolean if a field has been set.

### GetRegionName

`func (o *GetOrdersItemTicker) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *GetOrdersItemTicker) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *GetOrdersItemTicker) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *GetOrdersItemTicker) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.

### GetSecType

`func (o *GetOrdersItemTicker) GetSecType() []int32`

GetSecType returns the SecType field if non-nil, zero value otherwise.

### GetSecTypeOk

`func (o *GetOrdersItemTicker) GetSecTypeOk() (*[]int32, bool)`

GetSecTypeOk returns a tuple with the SecType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecType

`func (o *GetOrdersItemTicker) SetSecType(v []int32)`

SetSecType sets SecType field to given value.

### HasSecType

`func (o *GetOrdersItemTicker) HasSecType() bool`

HasSecType returns a boolean if a field has been set.

### GetSymbol

`func (o *GetOrdersItemTicker) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetOrdersItemTicker) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetOrdersItemTicker) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *GetOrdersItemTicker) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTickerId

`func (o *GetOrdersItemTicker) GetTickerId() int32`

GetTickerId returns the TickerId field if non-nil, zero value otherwise.

### GetTickerIdOk

`func (o *GetOrdersItemTicker) GetTickerIdOk() (*int32, bool)`

GetTickerIdOk returns a tuple with the TickerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickerId

`func (o *GetOrdersItemTicker) SetTickerId(v int32)`

SetTickerId sets TickerId field to given value.

### HasTickerId

`func (o *GetOrdersItemTicker) HasTickerId() bool`

HasTickerId returns a boolean if a field has been set.

### GetTinyName

`func (o *GetOrdersItemTicker) GetTinyName() string`

GetTinyName returns the TinyName field if non-nil, zero value otherwise.

### GetTinyNameOk

`func (o *GetOrdersItemTicker) GetTinyNameOk() (*string, bool)`

GetTinyNameOk returns a tuple with the TinyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTinyName

`func (o *GetOrdersItemTicker) SetTinyName(v string)`

SetTinyName sets TinyName field to given value.

### HasTinyName

`func (o *GetOrdersItemTicker) HasTinyName() bool`

HasTinyName returns a boolean if a field has been set.

### GetType

`func (o *GetOrdersItemTicker) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetOrdersItemTicker) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetOrdersItemTicker) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *GetOrdersItemTicker) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


