# GetAcccountRequestPositions

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
**UpdatePositionTimeStamp** | Pointer to **float32** |  | [optional] 

## Methods

### NewGetAcccountRequestPositions

`func NewGetAcccountRequestPositions() *GetAcccountRequestPositions`

NewGetAcccountRequestPositions instantiates a new GetAcccountRequestPositions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAcccountRequestPositionsWithDefaults

`func NewGetAcccountRequestPositionsWithDefaults() *GetAcccountRequestPositions`

NewGetAcccountRequestPositionsWithDefaults instantiates a new GetAcccountRequestPositions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetType

`func (o *GetAcccountRequestPositions) GetAssetType() string`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *GetAcccountRequestPositions) GetAssetTypeOk() (*string, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *GetAcccountRequestPositions) SetAssetType(v string)`

SetAssetType sets AssetType field to given value.

### HasAssetType

`func (o *GetAcccountRequestPositions) HasAssetType() bool`

HasAssetType returns a boolean if a field has been set.

### GetBrokerId

`func (o *GetAcccountRequestPositions) GetBrokerId() int32`

GetBrokerId returns the BrokerId field if non-nil, zero value otherwise.

### GetBrokerIdOk

`func (o *GetAcccountRequestPositions) GetBrokerIdOk() (*int32, bool)`

GetBrokerIdOk returns a tuple with the BrokerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerId

`func (o *GetAcccountRequestPositions) SetBrokerId(v int32)`

SetBrokerId sets BrokerId field to given value.

### HasBrokerId

`func (o *GetAcccountRequestPositions) HasBrokerId() bool`

HasBrokerId returns a boolean if a field has been set.

### GetCost

`func (o *GetAcccountRequestPositions) GetCost() string`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *GetAcccountRequestPositions) GetCostOk() (*string, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *GetAcccountRequestPositions) SetCost(v string)`

SetCost sets Cost field to given value.

### HasCost

`func (o *GetAcccountRequestPositions) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetCostPrice

`func (o *GetAcccountRequestPositions) GetCostPrice() string`

GetCostPrice returns the CostPrice field if non-nil, zero value otherwise.

### GetCostPriceOk

`func (o *GetAcccountRequestPositions) GetCostPriceOk() (*string, bool)`

GetCostPriceOk returns a tuple with the CostPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPrice

`func (o *GetAcccountRequestPositions) SetCostPrice(v string)`

SetCostPrice sets CostPrice field to given value.

### HasCostPrice

`func (o *GetAcccountRequestPositions) HasCostPrice() bool`

HasCostPrice returns a boolean if a field has been set.

### GetCurrency

`func (o *GetAcccountRequestPositions) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetAcccountRequestPositions) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetAcccountRequestPositions) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GetAcccountRequestPositions) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetExchangeRate

`func (o *GetAcccountRequestPositions) GetExchangeRate() string`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *GetAcccountRequestPositions) GetExchangeRateOk() (*string, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *GetAcccountRequestPositions) SetExchangeRate(v string)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *GetAcccountRequestPositions) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetId

`func (o *GetAcccountRequestPositions) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetAcccountRequestPositions) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetAcccountRequestPositions) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetAcccountRequestPositions) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastOpenTime

`func (o *GetAcccountRequestPositions) GetLastOpenTime() string`

GetLastOpenTime returns the LastOpenTime field if non-nil, zero value otherwise.

### GetLastOpenTimeOk

`func (o *GetAcccountRequestPositions) GetLastOpenTimeOk() (*string, bool)`

GetLastOpenTimeOk returns a tuple with the LastOpenTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOpenTime

`func (o *GetAcccountRequestPositions) SetLastOpenTime(v string)`

SetLastOpenTime sets LastOpenTime field to given value.

### HasLastOpenTime

`func (o *GetAcccountRequestPositions) HasLastOpenTime() bool`

HasLastOpenTime returns a boolean if a field has been set.

### GetLastPrice

`func (o *GetAcccountRequestPositions) GetLastPrice() string`

GetLastPrice returns the LastPrice field if non-nil, zero value otherwise.

### GetLastPriceOk

`func (o *GetAcccountRequestPositions) GetLastPriceOk() (*string, bool)`

GetLastPriceOk returns a tuple with the LastPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPrice

`func (o *GetAcccountRequestPositions) SetLastPrice(v string)`

SetLastPrice sets LastPrice field to given value.

### HasLastPrice

`func (o *GetAcccountRequestPositions) HasLastPrice() bool`

HasLastPrice returns a boolean if a field has been set.

### GetLock

`func (o *GetAcccountRequestPositions) GetLock() bool`

GetLock returns the Lock field if non-nil, zero value otherwise.

### GetLockOk

`func (o *GetAcccountRequestPositions) GetLockOk() (*bool, bool)`

GetLockOk returns a tuple with the Lock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLock

`func (o *GetAcccountRequestPositions) SetLock(v bool)`

SetLock sets Lock field to given value.

### HasLock

`func (o *GetAcccountRequestPositions) HasLock() bool`

HasLock returns a boolean if a field has been set.

### GetMarketValue

`func (o *GetAcccountRequestPositions) GetMarketValue() string`

GetMarketValue returns the MarketValue field if non-nil, zero value otherwise.

### GetMarketValueOk

`func (o *GetAcccountRequestPositions) GetMarketValueOk() (*string, bool)`

GetMarketValueOk returns a tuple with the MarketValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketValue

`func (o *GetAcccountRequestPositions) SetMarketValue(v string)`

SetMarketValue sets MarketValue field to given value.

### HasMarketValue

`func (o *GetAcccountRequestPositions) HasMarketValue() bool`

HasMarketValue returns a boolean if a field has been set.

### GetPosition

`func (o *GetAcccountRequestPositions) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *GetAcccountRequestPositions) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *GetAcccountRequestPositions) SetPosition(v string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *GetAcccountRequestPositions) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetPositionProportion

`func (o *GetAcccountRequestPositions) GetPositionProportion() string`

GetPositionProportion returns the PositionProportion field if non-nil, zero value otherwise.

### GetPositionProportionOk

`func (o *GetAcccountRequestPositions) GetPositionProportionOk() (*string, bool)`

GetPositionProportionOk returns a tuple with the PositionProportion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionProportion

`func (o *GetAcccountRequestPositions) SetPositionProportion(v string)`

SetPositionProportion sets PositionProportion field to given value.

### HasPositionProportion

`func (o *GetAcccountRequestPositions) HasPositionProportion() bool`

HasPositionProportion returns a boolean if a field has been set.

### GetTicker

`func (o *GetAcccountRequestPositions) GetTicker() GetAcccountRequestTicker`

GetTicker returns the Ticker field if non-nil, zero value otherwise.

### GetTickerOk

`func (o *GetAcccountRequestPositions) GetTickerOk() (*GetAcccountRequestTicker, bool)`

GetTickerOk returns a tuple with the Ticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicker

`func (o *GetAcccountRequestPositions) SetTicker(v GetAcccountRequestTicker)`

SetTicker sets Ticker field to given value.

### HasTicker

`func (o *GetAcccountRequestPositions) HasTicker() bool`

HasTicker returns a boolean if a field has been set.

### GetUnrealizedProfitLoss

`func (o *GetAcccountRequestPositions) GetUnrealizedProfitLoss() string`

GetUnrealizedProfitLoss returns the UnrealizedProfitLoss field if non-nil, zero value otherwise.

### GetUnrealizedProfitLossOk

`func (o *GetAcccountRequestPositions) GetUnrealizedProfitLossOk() (*string, bool)`

GetUnrealizedProfitLossOk returns a tuple with the UnrealizedProfitLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedProfitLoss

`func (o *GetAcccountRequestPositions) SetUnrealizedProfitLoss(v string)`

SetUnrealizedProfitLoss sets UnrealizedProfitLoss field to given value.

### HasUnrealizedProfitLoss

`func (o *GetAcccountRequestPositions) HasUnrealizedProfitLoss() bool`

HasUnrealizedProfitLoss returns a boolean if a field has been set.

### GetUnrealizedProfitLossRate

`func (o *GetAcccountRequestPositions) GetUnrealizedProfitLossRate() string`

GetUnrealizedProfitLossRate returns the UnrealizedProfitLossRate field if non-nil, zero value otherwise.

### GetUnrealizedProfitLossRateOk

`func (o *GetAcccountRequestPositions) GetUnrealizedProfitLossRateOk() (*string, bool)`

GetUnrealizedProfitLossRateOk returns a tuple with the UnrealizedProfitLossRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedProfitLossRate

`func (o *GetAcccountRequestPositions) SetUnrealizedProfitLossRate(v string)`

SetUnrealizedProfitLossRate sets UnrealizedProfitLossRate field to given value.

### HasUnrealizedProfitLossRate

`func (o *GetAcccountRequestPositions) HasUnrealizedProfitLossRate() bool`

HasUnrealizedProfitLossRate returns a boolean if a field has been set.

### GetUpdatePositionTimeStamp

`func (o *GetAcccountRequestPositions) GetUpdatePositionTimeStamp() float32`

GetUpdatePositionTimeStamp returns the UpdatePositionTimeStamp field if non-nil, zero value otherwise.

### GetUpdatePositionTimeStampOk

`func (o *GetAcccountRequestPositions) GetUpdatePositionTimeStampOk() (*float32, bool)`

GetUpdatePositionTimeStampOk returns a tuple with the UpdatePositionTimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatePositionTimeStamp

`func (o *GetAcccountRequestPositions) SetUpdatePositionTimeStamp(v float32)`

SetUpdatePositionTimeStamp sets UpdatePositionTimeStamp field to given value.

### HasUpdatePositionTimeStamp

`func (o *GetAcccountRequestPositions) HasUpdatePositionTimeStamp() bool`

HasUpdatePositionTimeStamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


