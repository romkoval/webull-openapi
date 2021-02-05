# GetAccountsResponseV5Positions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PositionId** | Pointer to **float32** |  | [optional] 
**BrokerPosId** | Pointer to **string** |  | [optional] 
**BrokerId** | Pointer to **float32** |  | [optional] 
**AssetType** | Pointer to **string** |  | [optional] 
**TickerType** | Pointer to **string** |  | [optional] 
**Ticker** | Pointer to [**GetAccountsResponseV5Ticker**](GetAccountsResponseV5_ticker.md) |  | [optional] 
**Action** | Pointer to **string** |  | [optional] 
**Position** | Pointer to **string** |  | [optional] 
**TickerId** | Pointer to **float32** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**CostPrice** | Pointer to **string** |  | [optional] 
**Cost** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**MarketValue** | Pointer to **string** |  | [optional] 
**LastPrice** | Pointer to **string** |  | [optional] 
**UnrealizedProfitLoss** | Pointer to **string** |  | [optional] 
**UnrealizedProfitLossRate** | Pointer to **string** |  | [optional] 
**PositionProportion** | Pointer to **string** |  | [optional] 
**UpdatePositionTime** | Pointer to **float32** |  | [optional] 
**OptionCanExercise** | Pointer to **float32** |  | [optional] 
**RecentlyExpireFlag** | Pointer to **bool** |  | [optional] 
**RemainDay** | Pointer to **float32** |  | [optional] 
**IsLending** | Pointer to **string** |  | [optional] 

## Methods

### NewGetAccountsResponseV5Positions

`func NewGetAccountsResponseV5Positions() *GetAccountsResponseV5Positions`

NewGetAccountsResponseV5Positions instantiates a new GetAccountsResponseV5Positions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountsResponseV5PositionsWithDefaults

`func NewGetAccountsResponseV5PositionsWithDefaults() *GetAccountsResponseV5Positions`

NewGetAccountsResponseV5PositionsWithDefaults instantiates a new GetAccountsResponseV5Positions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPositionId

`func (o *GetAccountsResponseV5Positions) GetPositionId() float32`

GetPositionId returns the PositionId field if non-nil, zero value otherwise.

### GetPositionIdOk

`func (o *GetAccountsResponseV5Positions) GetPositionIdOk() (*float32, bool)`

GetPositionIdOk returns a tuple with the PositionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionId

`func (o *GetAccountsResponseV5Positions) SetPositionId(v float32)`

SetPositionId sets PositionId field to given value.

### HasPositionId

`func (o *GetAccountsResponseV5Positions) HasPositionId() bool`

HasPositionId returns a boolean if a field has been set.

### GetBrokerPosId

`func (o *GetAccountsResponseV5Positions) GetBrokerPosId() string`

GetBrokerPosId returns the BrokerPosId field if non-nil, zero value otherwise.

### GetBrokerPosIdOk

`func (o *GetAccountsResponseV5Positions) GetBrokerPosIdOk() (*string, bool)`

GetBrokerPosIdOk returns a tuple with the BrokerPosId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerPosId

`func (o *GetAccountsResponseV5Positions) SetBrokerPosId(v string)`

SetBrokerPosId sets BrokerPosId field to given value.

### HasBrokerPosId

`func (o *GetAccountsResponseV5Positions) HasBrokerPosId() bool`

HasBrokerPosId returns a boolean if a field has been set.

### GetBrokerId

`func (o *GetAccountsResponseV5Positions) GetBrokerId() float32`

GetBrokerId returns the BrokerId field if non-nil, zero value otherwise.

### GetBrokerIdOk

`func (o *GetAccountsResponseV5Positions) GetBrokerIdOk() (*float32, bool)`

GetBrokerIdOk returns a tuple with the BrokerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerId

`func (o *GetAccountsResponseV5Positions) SetBrokerId(v float32)`

SetBrokerId sets BrokerId field to given value.

### HasBrokerId

`func (o *GetAccountsResponseV5Positions) HasBrokerId() bool`

HasBrokerId returns a boolean if a field has been set.

### GetAssetType

`func (o *GetAccountsResponseV5Positions) GetAssetType() string`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *GetAccountsResponseV5Positions) GetAssetTypeOk() (*string, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *GetAccountsResponseV5Positions) SetAssetType(v string)`

SetAssetType sets AssetType field to given value.

### HasAssetType

`func (o *GetAccountsResponseV5Positions) HasAssetType() bool`

HasAssetType returns a boolean if a field has been set.

### GetTickerType

`func (o *GetAccountsResponseV5Positions) GetTickerType() string`

GetTickerType returns the TickerType field if non-nil, zero value otherwise.

### GetTickerTypeOk

`func (o *GetAccountsResponseV5Positions) GetTickerTypeOk() (*string, bool)`

GetTickerTypeOk returns a tuple with the TickerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickerType

`func (o *GetAccountsResponseV5Positions) SetTickerType(v string)`

SetTickerType sets TickerType field to given value.

### HasTickerType

`func (o *GetAccountsResponseV5Positions) HasTickerType() bool`

HasTickerType returns a boolean if a field has been set.

### GetTicker

`func (o *GetAccountsResponseV5Positions) GetTicker() GetAccountsResponseV5Ticker`

GetTicker returns the Ticker field if non-nil, zero value otherwise.

### GetTickerOk

`func (o *GetAccountsResponseV5Positions) GetTickerOk() (*GetAccountsResponseV5Ticker, bool)`

GetTickerOk returns a tuple with the Ticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicker

`func (o *GetAccountsResponseV5Positions) SetTicker(v GetAccountsResponseV5Ticker)`

SetTicker sets Ticker field to given value.

### HasTicker

`func (o *GetAccountsResponseV5Positions) HasTicker() bool`

HasTicker returns a boolean if a field has been set.

### GetAction

`func (o *GetAccountsResponseV5Positions) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *GetAccountsResponseV5Positions) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *GetAccountsResponseV5Positions) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *GetAccountsResponseV5Positions) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetPosition

`func (o *GetAccountsResponseV5Positions) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *GetAccountsResponseV5Positions) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *GetAccountsResponseV5Positions) SetPosition(v string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *GetAccountsResponseV5Positions) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetTickerId

`func (o *GetAccountsResponseV5Positions) GetTickerId() float32`

GetTickerId returns the TickerId field if non-nil, zero value otherwise.

### GetTickerIdOk

`func (o *GetAccountsResponseV5Positions) GetTickerIdOk() (*float32, bool)`

GetTickerIdOk returns a tuple with the TickerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickerId

`func (o *GetAccountsResponseV5Positions) SetTickerId(v float32)`

SetTickerId sets TickerId field to given value.

### HasTickerId

`func (o *GetAccountsResponseV5Positions) HasTickerId() bool`

HasTickerId returns a boolean if a field has been set.

### GetSymbol

`func (o *GetAccountsResponseV5Positions) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetAccountsResponseV5Positions) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetAccountsResponseV5Positions) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *GetAccountsResponseV5Positions) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetCostPrice

`func (o *GetAccountsResponseV5Positions) GetCostPrice() string`

GetCostPrice returns the CostPrice field if non-nil, zero value otherwise.

### GetCostPriceOk

`func (o *GetAccountsResponseV5Positions) GetCostPriceOk() (*string, bool)`

GetCostPriceOk returns a tuple with the CostPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPrice

`func (o *GetAccountsResponseV5Positions) SetCostPrice(v string)`

SetCostPrice sets CostPrice field to given value.

### HasCostPrice

`func (o *GetAccountsResponseV5Positions) HasCostPrice() bool`

HasCostPrice returns a boolean if a field has been set.

### GetCost

`func (o *GetAccountsResponseV5Positions) GetCost() string`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *GetAccountsResponseV5Positions) GetCostOk() (*string, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *GetAccountsResponseV5Positions) SetCost(v string)`

SetCost sets Cost field to given value.

### HasCost

`func (o *GetAccountsResponseV5Positions) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetCurrency

`func (o *GetAccountsResponseV5Positions) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetAccountsResponseV5Positions) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetAccountsResponseV5Positions) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GetAccountsResponseV5Positions) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetMarketValue

`func (o *GetAccountsResponseV5Positions) GetMarketValue() string`

GetMarketValue returns the MarketValue field if non-nil, zero value otherwise.

### GetMarketValueOk

`func (o *GetAccountsResponseV5Positions) GetMarketValueOk() (*string, bool)`

GetMarketValueOk returns a tuple with the MarketValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketValue

`func (o *GetAccountsResponseV5Positions) SetMarketValue(v string)`

SetMarketValue sets MarketValue field to given value.

### HasMarketValue

`func (o *GetAccountsResponseV5Positions) HasMarketValue() bool`

HasMarketValue returns a boolean if a field has been set.

### GetLastPrice

`func (o *GetAccountsResponseV5Positions) GetLastPrice() string`

GetLastPrice returns the LastPrice field if non-nil, zero value otherwise.

### GetLastPriceOk

`func (o *GetAccountsResponseV5Positions) GetLastPriceOk() (*string, bool)`

GetLastPriceOk returns a tuple with the LastPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPrice

`func (o *GetAccountsResponseV5Positions) SetLastPrice(v string)`

SetLastPrice sets LastPrice field to given value.

### HasLastPrice

`func (o *GetAccountsResponseV5Positions) HasLastPrice() bool`

HasLastPrice returns a boolean if a field has been set.

### GetUnrealizedProfitLoss

`func (o *GetAccountsResponseV5Positions) GetUnrealizedProfitLoss() string`

GetUnrealizedProfitLoss returns the UnrealizedProfitLoss field if non-nil, zero value otherwise.

### GetUnrealizedProfitLossOk

`func (o *GetAccountsResponseV5Positions) GetUnrealizedProfitLossOk() (*string, bool)`

GetUnrealizedProfitLossOk returns a tuple with the UnrealizedProfitLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedProfitLoss

`func (o *GetAccountsResponseV5Positions) SetUnrealizedProfitLoss(v string)`

SetUnrealizedProfitLoss sets UnrealizedProfitLoss field to given value.

### HasUnrealizedProfitLoss

`func (o *GetAccountsResponseV5Positions) HasUnrealizedProfitLoss() bool`

HasUnrealizedProfitLoss returns a boolean if a field has been set.

### GetUnrealizedProfitLossRate

`func (o *GetAccountsResponseV5Positions) GetUnrealizedProfitLossRate() string`

GetUnrealizedProfitLossRate returns the UnrealizedProfitLossRate field if non-nil, zero value otherwise.

### GetUnrealizedProfitLossRateOk

`func (o *GetAccountsResponseV5Positions) GetUnrealizedProfitLossRateOk() (*string, bool)`

GetUnrealizedProfitLossRateOk returns a tuple with the UnrealizedProfitLossRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedProfitLossRate

`func (o *GetAccountsResponseV5Positions) SetUnrealizedProfitLossRate(v string)`

SetUnrealizedProfitLossRate sets UnrealizedProfitLossRate field to given value.

### HasUnrealizedProfitLossRate

`func (o *GetAccountsResponseV5Positions) HasUnrealizedProfitLossRate() bool`

HasUnrealizedProfitLossRate returns a boolean if a field has been set.

### GetPositionProportion

`func (o *GetAccountsResponseV5Positions) GetPositionProportion() string`

GetPositionProportion returns the PositionProportion field if non-nil, zero value otherwise.

### GetPositionProportionOk

`func (o *GetAccountsResponseV5Positions) GetPositionProportionOk() (*string, bool)`

GetPositionProportionOk returns a tuple with the PositionProportion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionProportion

`func (o *GetAccountsResponseV5Positions) SetPositionProportion(v string)`

SetPositionProportion sets PositionProportion field to given value.

### HasPositionProportion

`func (o *GetAccountsResponseV5Positions) HasPositionProportion() bool`

HasPositionProportion returns a boolean if a field has been set.

### GetUpdatePositionTime

`func (o *GetAccountsResponseV5Positions) GetUpdatePositionTime() float32`

GetUpdatePositionTime returns the UpdatePositionTime field if non-nil, zero value otherwise.

### GetUpdatePositionTimeOk

`func (o *GetAccountsResponseV5Positions) GetUpdatePositionTimeOk() (*float32, bool)`

GetUpdatePositionTimeOk returns a tuple with the UpdatePositionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatePositionTime

`func (o *GetAccountsResponseV5Positions) SetUpdatePositionTime(v float32)`

SetUpdatePositionTime sets UpdatePositionTime field to given value.

### HasUpdatePositionTime

`func (o *GetAccountsResponseV5Positions) HasUpdatePositionTime() bool`

HasUpdatePositionTime returns a boolean if a field has been set.

### GetOptionCanExercise

`func (o *GetAccountsResponseV5Positions) GetOptionCanExercise() float32`

GetOptionCanExercise returns the OptionCanExercise field if non-nil, zero value otherwise.

### GetOptionCanExerciseOk

`func (o *GetAccountsResponseV5Positions) GetOptionCanExerciseOk() (*float32, bool)`

GetOptionCanExerciseOk returns a tuple with the OptionCanExercise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionCanExercise

`func (o *GetAccountsResponseV5Positions) SetOptionCanExercise(v float32)`

SetOptionCanExercise sets OptionCanExercise field to given value.

### HasOptionCanExercise

`func (o *GetAccountsResponseV5Positions) HasOptionCanExercise() bool`

HasOptionCanExercise returns a boolean if a field has been set.

### GetRecentlyExpireFlag

`func (o *GetAccountsResponseV5Positions) GetRecentlyExpireFlag() bool`

GetRecentlyExpireFlag returns the RecentlyExpireFlag field if non-nil, zero value otherwise.

### GetRecentlyExpireFlagOk

`func (o *GetAccountsResponseV5Positions) GetRecentlyExpireFlagOk() (*bool, bool)`

GetRecentlyExpireFlagOk returns a tuple with the RecentlyExpireFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentlyExpireFlag

`func (o *GetAccountsResponseV5Positions) SetRecentlyExpireFlag(v bool)`

SetRecentlyExpireFlag sets RecentlyExpireFlag field to given value.

### HasRecentlyExpireFlag

`func (o *GetAccountsResponseV5Positions) HasRecentlyExpireFlag() bool`

HasRecentlyExpireFlag returns a boolean if a field has been set.

### GetRemainDay

`func (o *GetAccountsResponseV5Positions) GetRemainDay() float32`

GetRemainDay returns the RemainDay field if non-nil, zero value otherwise.

### GetRemainDayOk

`func (o *GetAccountsResponseV5Positions) GetRemainDayOk() (*float32, bool)`

GetRemainDayOk returns a tuple with the RemainDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainDay

`func (o *GetAccountsResponseV5Positions) SetRemainDay(v float32)`

SetRemainDay sets RemainDay field to given value.

### HasRemainDay

`func (o *GetAccountsResponseV5Positions) HasRemainDay() bool`

HasRemainDay returns a boolean if a field has been set.

### GetIsLending

`func (o *GetAccountsResponseV5Positions) GetIsLending() string`

GetIsLending returns the IsLending field if non-nil, zero value otherwise.

### GetIsLendingOk

`func (o *GetAccountsResponseV5Positions) GetIsLendingOk() (*string, bool)`

GetIsLendingOk returns a tuple with the IsLending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLending

`func (o *GetAccountsResponseV5Positions) SetIsLending(v string)`

SetIsLending sets IsLending field to given value.

### HasIsLending

`func (o *GetAccountsResponseV5Positions) HasIsLending() bool`

HasIsLending returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


