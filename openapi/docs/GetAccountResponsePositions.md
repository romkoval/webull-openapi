# GetAccountResponsePositions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetType** | Pointer to **string** |  | [optional] 
**BrokerId** | Pointer to **int32** |  | [optional] 
**Cost** | Pointer to **string** |  | [optional] 
**CostPrice** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**ExchangeRate** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**LastOpenTime** | Pointer to **string** |  | [optional] 
**LastPrice** | Pointer to **string** |  | [optional] 
**Lock** | Pointer to **bool** |  | [optional] 
**MarketValue** | Pointer to **string** |  | [optional] 
**Position** | Pointer to **string** |  | [optional] 
**PositionProportion** | Pointer to **string** |  | [optional] 
**Ticker** | Pointer to [**GetAcccountRequestTicker**](GetAcccountRequest_ticker.md) |  | [optional] 
**UnrealizedProfitLoss** | Pointer to **string** |  | [optional] 
**UnrealizedProfitLossRate** | Pointer to **string** |  | [optional] 
**UpdatePositionTimeStamp** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetAccountResponsePositions

`func NewGetAccountResponsePositions() *GetAccountResponsePositions`

NewGetAccountResponsePositions instantiates a new GetAccountResponsePositions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountResponsePositionsWithDefaults

`func NewGetAccountResponsePositionsWithDefaults() *GetAccountResponsePositions`

NewGetAccountResponsePositionsWithDefaults instantiates a new GetAccountResponsePositions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetType

`func (o *GetAccountResponsePositions) GetAssetType() string`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *GetAccountResponsePositions) GetAssetTypeOk() (*string, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *GetAccountResponsePositions) SetAssetType(v string)`

SetAssetType sets AssetType field to given value.

### HasAssetType

`func (o *GetAccountResponsePositions) HasAssetType() bool`

HasAssetType returns a boolean if a field has been set.

### GetBrokerId

`func (o *GetAccountResponsePositions) GetBrokerId() int32`

GetBrokerId returns the BrokerId field if non-nil, zero value otherwise.

### GetBrokerIdOk

`func (o *GetAccountResponsePositions) GetBrokerIdOk() (*int32, bool)`

GetBrokerIdOk returns a tuple with the BrokerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerId

`func (o *GetAccountResponsePositions) SetBrokerId(v int32)`

SetBrokerId sets BrokerId field to given value.

### HasBrokerId

`func (o *GetAccountResponsePositions) HasBrokerId() bool`

HasBrokerId returns a boolean if a field has been set.

### GetCost

`func (o *GetAccountResponsePositions) GetCost() string`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *GetAccountResponsePositions) GetCostOk() (*string, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *GetAccountResponsePositions) SetCost(v string)`

SetCost sets Cost field to given value.

### HasCost

`func (o *GetAccountResponsePositions) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetCostPrice

`func (o *GetAccountResponsePositions) GetCostPrice() string`

GetCostPrice returns the CostPrice field if non-nil, zero value otherwise.

### GetCostPriceOk

`func (o *GetAccountResponsePositions) GetCostPriceOk() (*string, bool)`

GetCostPriceOk returns a tuple with the CostPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPrice

`func (o *GetAccountResponsePositions) SetCostPrice(v string)`

SetCostPrice sets CostPrice field to given value.

### HasCostPrice

`func (o *GetAccountResponsePositions) HasCostPrice() bool`

HasCostPrice returns a boolean if a field has been set.

### GetCurrency

`func (o *GetAccountResponsePositions) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetAccountResponsePositions) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetAccountResponsePositions) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GetAccountResponsePositions) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetExchangeRate

`func (o *GetAccountResponsePositions) GetExchangeRate() string`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *GetAccountResponsePositions) GetExchangeRateOk() (*string, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *GetAccountResponsePositions) SetExchangeRate(v string)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *GetAccountResponsePositions) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetId

`func (o *GetAccountResponsePositions) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetAccountResponsePositions) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetAccountResponsePositions) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetAccountResponsePositions) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastOpenTime

`func (o *GetAccountResponsePositions) GetLastOpenTime() string`

GetLastOpenTime returns the LastOpenTime field if non-nil, zero value otherwise.

### GetLastOpenTimeOk

`func (o *GetAccountResponsePositions) GetLastOpenTimeOk() (*string, bool)`

GetLastOpenTimeOk returns a tuple with the LastOpenTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOpenTime

`func (o *GetAccountResponsePositions) SetLastOpenTime(v string)`

SetLastOpenTime sets LastOpenTime field to given value.

### HasLastOpenTime

`func (o *GetAccountResponsePositions) HasLastOpenTime() bool`

HasLastOpenTime returns a boolean if a field has been set.

### GetLastPrice

`func (o *GetAccountResponsePositions) GetLastPrice() string`

GetLastPrice returns the LastPrice field if non-nil, zero value otherwise.

### GetLastPriceOk

`func (o *GetAccountResponsePositions) GetLastPriceOk() (*string, bool)`

GetLastPriceOk returns a tuple with the LastPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPrice

`func (o *GetAccountResponsePositions) SetLastPrice(v string)`

SetLastPrice sets LastPrice field to given value.

### HasLastPrice

`func (o *GetAccountResponsePositions) HasLastPrice() bool`

HasLastPrice returns a boolean if a field has been set.

### GetLock

`func (o *GetAccountResponsePositions) GetLock() bool`

GetLock returns the Lock field if non-nil, zero value otherwise.

### GetLockOk

`func (o *GetAccountResponsePositions) GetLockOk() (*bool, bool)`

GetLockOk returns a tuple with the Lock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLock

`func (o *GetAccountResponsePositions) SetLock(v bool)`

SetLock sets Lock field to given value.

### HasLock

`func (o *GetAccountResponsePositions) HasLock() bool`

HasLock returns a boolean if a field has been set.

### GetMarketValue

`func (o *GetAccountResponsePositions) GetMarketValue() string`

GetMarketValue returns the MarketValue field if non-nil, zero value otherwise.

### GetMarketValueOk

`func (o *GetAccountResponsePositions) GetMarketValueOk() (*string, bool)`

GetMarketValueOk returns a tuple with the MarketValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketValue

`func (o *GetAccountResponsePositions) SetMarketValue(v string)`

SetMarketValue sets MarketValue field to given value.

### HasMarketValue

`func (o *GetAccountResponsePositions) HasMarketValue() bool`

HasMarketValue returns a boolean if a field has been set.

### GetPosition

`func (o *GetAccountResponsePositions) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *GetAccountResponsePositions) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *GetAccountResponsePositions) SetPosition(v string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *GetAccountResponsePositions) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetPositionProportion

`func (o *GetAccountResponsePositions) GetPositionProportion() string`

GetPositionProportion returns the PositionProportion field if non-nil, zero value otherwise.

### GetPositionProportionOk

`func (o *GetAccountResponsePositions) GetPositionProportionOk() (*string, bool)`

GetPositionProportionOk returns a tuple with the PositionProportion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionProportion

`func (o *GetAccountResponsePositions) SetPositionProportion(v string)`

SetPositionProportion sets PositionProportion field to given value.

### HasPositionProportion

`func (o *GetAccountResponsePositions) HasPositionProportion() bool`

HasPositionProportion returns a boolean if a field has been set.

### GetTicker

`func (o *GetAccountResponsePositions) GetTicker() GetAcccountRequestTicker`

GetTicker returns the Ticker field if non-nil, zero value otherwise.

### GetTickerOk

`func (o *GetAccountResponsePositions) GetTickerOk() (*GetAcccountRequestTicker, bool)`

GetTickerOk returns a tuple with the Ticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicker

`func (o *GetAccountResponsePositions) SetTicker(v GetAcccountRequestTicker)`

SetTicker sets Ticker field to given value.

### HasTicker

`func (o *GetAccountResponsePositions) HasTicker() bool`

HasTicker returns a boolean if a field has been set.

### GetUnrealizedProfitLoss

`func (o *GetAccountResponsePositions) GetUnrealizedProfitLoss() string`

GetUnrealizedProfitLoss returns the UnrealizedProfitLoss field if non-nil, zero value otherwise.

### GetUnrealizedProfitLossOk

`func (o *GetAccountResponsePositions) GetUnrealizedProfitLossOk() (*string, bool)`

GetUnrealizedProfitLossOk returns a tuple with the UnrealizedProfitLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedProfitLoss

`func (o *GetAccountResponsePositions) SetUnrealizedProfitLoss(v string)`

SetUnrealizedProfitLoss sets UnrealizedProfitLoss field to given value.

### HasUnrealizedProfitLoss

`func (o *GetAccountResponsePositions) HasUnrealizedProfitLoss() bool`

HasUnrealizedProfitLoss returns a boolean if a field has been set.

### GetUnrealizedProfitLossRate

`func (o *GetAccountResponsePositions) GetUnrealizedProfitLossRate() string`

GetUnrealizedProfitLossRate returns the UnrealizedProfitLossRate field if non-nil, zero value otherwise.

### GetUnrealizedProfitLossRateOk

`func (o *GetAccountResponsePositions) GetUnrealizedProfitLossRateOk() (*string, bool)`

GetUnrealizedProfitLossRateOk returns a tuple with the UnrealizedProfitLossRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedProfitLossRate

`func (o *GetAccountResponsePositions) SetUnrealizedProfitLossRate(v string)`

SetUnrealizedProfitLossRate sets UnrealizedProfitLossRate field to given value.

### HasUnrealizedProfitLossRate

`func (o *GetAccountResponsePositions) HasUnrealizedProfitLossRate() bool`

HasUnrealizedProfitLossRate returns a boolean if a field has been set.

### GetUpdatePositionTimeStamp

`func (o *GetAccountResponsePositions) GetUpdatePositionTimeStamp() int64`

GetUpdatePositionTimeStamp returns the UpdatePositionTimeStamp field if non-nil, zero value otherwise.

### GetUpdatePositionTimeStampOk

`func (o *GetAccountResponsePositions) GetUpdatePositionTimeStampOk() (*int64, bool)`

GetUpdatePositionTimeStampOk returns a tuple with the UpdatePositionTimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatePositionTimeStamp

`func (o *GetAccountResponsePositions) SetUpdatePositionTimeStamp(v int64)`

SetUpdatePositionTimeStamp sets UpdatePositionTimeStamp field to given value.

### HasUpdatePositionTimeStamp

`func (o *GetAccountResponsePositions) HasUpdatePositionTimeStamp() bool`

HasUpdatePositionTimeStamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


