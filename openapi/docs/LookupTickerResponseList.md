# LookupTickerResponseList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyId** | Pointer to **int32** |  | [optional] 
**DerivativeSupport** | Pointer to **int32** |  | [optional] 
**DisExchangeCode** | Pointer to **string** |  | [optional] 
**DisSymbol** | Pointer to **string** |  | [optional] 
**ExchangeCode** | Pointer to **string** |  | [optional] 
**ExchangeId** | Pointer to **int32** |  | [optional] 
**ListStatus** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**RegionCode** | Pointer to **string** |  | [optional] 
**RegionId** | Pointer to **int32** |  | [optional] 
**SecType** | Pointer to **[]int32** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Template** | Pointer to **string** |  | [optional] 
**TickerId** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **int32** |  | [optional] 

## Methods

### NewLookupTickerResponseList

`func NewLookupTickerResponseList() *LookupTickerResponseList`

NewLookupTickerResponseList instantiates a new LookupTickerResponseList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLookupTickerResponseListWithDefaults

`func NewLookupTickerResponseListWithDefaults() *LookupTickerResponseList`

NewLookupTickerResponseListWithDefaults instantiates a new LookupTickerResponseList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencyId

`func (o *LookupTickerResponseList) GetCurrencyId() int32`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *LookupTickerResponseList) GetCurrencyIdOk() (*int32, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *LookupTickerResponseList) SetCurrencyId(v int32)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *LookupTickerResponseList) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetDerivativeSupport

`func (o *LookupTickerResponseList) GetDerivativeSupport() int32`

GetDerivativeSupport returns the DerivativeSupport field if non-nil, zero value otherwise.

### GetDerivativeSupportOk

`func (o *LookupTickerResponseList) GetDerivativeSupportOk() (*int32, bool)`

GetDerivativeSupportOk returns a tuple with the DerivativeSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivativeSupport

`func (o *LookupTickerResponseList) SetDerivativeSupport(v int32)`

SetDerivativeSupport sets DerivativeSupport field to given value.

### HasDerivativeSupport

`func (o *LookupTickerResponseList) HasDerivativeSupport() bool`

HasDerivativeSupport returns a boolean if a field has been set.

### GetDisExchangeCode

`func (o *LookupTickerResponseList) GetDisExchangeCode() string`

GetDisExchangeCode returns the DisExchangeCode field if non-nil, zero value otherwise.

### GetDisExchangeCodeOk

`func (o *LookupTickerResponseList) GetDisExchangeCodeOk() (*string, bool)`

GetDisExchangeCodeOk returns a tuple with the DisExchangeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisExchangeCode

`func (o *LookupTickerResponseList) SetDisExchangeCode(v string)`

SetDisExchangeCode sets DisExchangeCode field to given value.

### HasDisExchangeCode

`func (o *LookupTickerResponseList) HasDisExchangeCode() bool`

HasDisExchangeCode returns a boolean if a field has been set.

### GetDisSymbol

`func (o *LookupTickerResponseList) GetDisSymbol() string`

GetDisSymbol returns the DisSymbol field if non-nil, zero value otherwise.

### GetDisSymbolOk

`func (o *LookupTickerResponseList) GetDisSymbolOk() (*string, bool)`

GetDisSymbolOk returns a tuple with the DisSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisSymbol

`func (o *LookupTickerResponseList) SetDisSymbol(v string)`

SetDisSymbol sets DisSymbol field to given value.

### HasDisSymbol

`func (o *LookupTickerResponseList) HasDisSymbol() bool`

HasDisSymbol returns a boolean if a field has been set.

### GetExchangeCode

`func (o *LookupTickerResponseList) GetExchangeCode() string`

GetExchangeCode returns the ExchangeCode field if non-nil, zero value otherwise.

### GetExchangeCodeOk

`func (o *LookupTickerResponseList) GetExchangeCodeOk() (*string, bool)`

GetExchangeCodeOk returns a tuple with the ExchangeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeCode

`func (o *LookupTickerResponseList) SetExchangeCode(v string)`

SetExchangeCode sets ExchangeCode field to given value.

### HasExchangeCode

`func (o *LookupTickerResponseList) HasExchangeCode() bool`

HasExchangeCode returns a boolean if a field has been set.

### GetExchangeId

`func (o *LookupTickerResponseList) GetExchangeId() int32`

GetExchangeId returns the ExchangeId field if non-nil, zero value otherwise.

### GetExchangeIdOk

`func (o *LookupTickerResponseList) GetExchangeIdOk() (*int32, bool)`

GetExchangeIdOk returns a tuple with the ExchangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeId

`func (o *LookupTickerResponseList) SetExchangeId(v int32)`

SetExchangeId sets ExchangeId field to given value.

### HasExchangeId

`func (o *LookupTickerResponseList) HasExchangeId() bool`

HasExchangeId returns a boolean if a field has been set.

### GetListStatus

`func (o *LookupTickerResponseList) GetListStatus() int32`

GetListStatus returns the ListStatus field if non-nil, zero value otherwise.

### GetListStatusOk

`func (o *LookupTickerResponseList) GetListStatusOk() (*int32, bool)`

GetListStatusOk returns a tuple with the ListStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListStatus

`func (o *LookupTickerResponseList) SetListStatus(v int32)`

SetListStatus sets ListStatus field to given value.

### HasListStatus

`func (o *LookupTickerResponseList) HasListStatus() bool`

HasListStatus returns a boolean if a field has been set.

### GetName

`func (o *LookupTickerResponseList) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LookupTickerResponseList) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LookupTickerResponseList) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LookupTickerResponseList) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegionCode

`func (o *LookupTickerResponseList) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *LookupTickerResponseList) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *LookupTickerResponseList) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *LookupTickerResponseList) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### GetRegionId

`func (o *LookupTickerResponseList) GetRegionId() int32`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *LookupTickerResponseList) GetRegionIdOk() (*int32, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *LookupTickerResponseList) SetRegionId(v int32)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *LookupTickerResponseList) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### GetSecType

`func (o *LookupTickerResponseList) GetSecType() []int32`

GetSecType returns the SecType field if non-nil, zero value otherwise.

### GetSecTypeOk

`func (o *LookupTickerResponseList) GetSecTypeOk() (*[]int32, bool)`

GetSecTypeOk returns a tuple with the SecType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecType

`func (o *LookupTickerResponseList) SetSecType(v []int32)`

SetSecType sets SecType field to given value.

### HasSecType

`func (o *LookupTickerResponseList) HasSecType() bool`

HasSecType returns a boolean if a field has been set.

### GetSymbol

`func (o *LookupTickerResponseList) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *LookupTickerResponseList) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *LookupTickerResponseList) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *LookupTickerResponseList) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTemplate

`func (o *LookupTickerResponseList) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *LookupTickerResponseList) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *LookupTickerResponseList) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *LookupTickerResponseList) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetTickerId

`func (o *LookupTickerResponseList) GetTickerId() int32`

GetTickerId returns the TickerId field if non-nil, zero value otherwise.

### GetTickerIdOk

`func (o *LookupTickerResponseList) GetTickerIdOk() (*int32, bool)`

GetTickerIdOk returns a tuple with the TickerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickerId

`func (o *LookupTickerResponseList) SetTickerId(v int32)`

SetTickerId sets TickerId field to given value.

### HasTickerId

`func (o *LookupTickerResponseList) HasTickerId() bool`

HasTickerId returns a boolean if a field has been set.

### GetType

`func (o *LookupTickerResponseList) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LookupTickerResponseList) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LookupTickerResponseList) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *LookupTickerResponseList) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


