# GetOrdersItemOrders

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**OrderSide**](OrderSide.md) |  | [optional] 
**AssetType** | Pointer to **string** |  | [optional] 
**AvgFilledPrice** | Pointer to **string** |  | [optional] 
**BrokerId** | Pointer to **int32** |  | [optional] 
**CanCancel** | Pointer to **bool** |  | [optional] 
**CanModify** | Pointer to **bool** |  | [optional] 
**CreateTime** | Pointer to **string** |  | [optional] 
**CreateTime0** | Pointer to **int64** |  | [optional] 
**FilledQuantity** | Pointer to **string** |  | [optional] 
**FilledTime** | Pointer to **string** |  | [optional] 
**FilledTime0** | Pointer to **float32** |  | [optional] 
**FilledValue** | Pointer to **string** |  | [optional] 
**OptionContractMultiplier** | Pointer to **string** |  | [optional] 
**OptionExercisePrice** | Pointer to **string** |  | [optional] 
**OptionCycle** | Pointer to **int64** |  | [optional] 
**OrderId** | Pointer to **int32** |  | [optional] 
**OrderType** | Pointer to [**OrderType**](OrderType.md) |  | [optional] 
**OptionType** | Pointer to [**OptionType**](OptionType.md) |  | [optional] 
**OptionExpirationDate** | Pointer to **string** |  | [optional] 
**Relation** | Pointer to **string** |  | [optional] 
**OptionCategory** | Pointer to **string** |  | [optional] 
**RemainQuantity** | Pointer to **string** |  | [optional] 
**StatusCode** | Pointer to **string** |  | [optional] 
**StatusStr** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Ticker** | Pointer to [**GetOrdersItemTicker**](GetOrdersItem_ticker.md) |  | [optional] 
**TickerId** | Pointer to **int32** |  | [optional] 
**TickerPriceDefineList** | Pointer to [**[]GetOrdersItemTickerPriceDefineList**](GetOrdersItemTickerPriceDefineList.md) |  | [optional] 
**TimeInForce** | Pointer to [**TimeInForce**](TimeInForce.md) |  | [optional] 
**TotalQuantity** | Pointer to **string** |  | [optional] 
**UpdateTime** | Pointer to **string** |  | [optional] 
**UpdateTime0** | Pointer to **float32** |  | [optional] 

## Methods

### NewGetOrdersItemOrders

`func NewGetOrdersItemOrders() *GetOrdersItemOrders`

NewGetOrdersItemOrders instantiates a new GetOrdersItemOrders object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrdersItemOrdersWithDefaults

`func NewGetOrdersItemOrdersWithDefaults() *GetOrdersItemOrders`

NewGetOrdersItemOrdersWithDefaults instantiates a new GetOrdersItemOrders object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *GetOrdersItemOrders) GetAction() OrderSide`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *GetOrdersItemOrders) GetActionOk() (*OrderSide, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *GetOrdersItemOrders) SetAction(v OrderSide)`

SetAction sets Action field to given value.

### HasAction

`func (o *GetOrdersItemOrders) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetAssetType

`func (o *GetOrdersItemOrders) GetAssetType() string`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *GetOrdersItemOrders) GetAssetTypeOk() (*string, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *GetOrdersItemOrders) SetAssetType(v string)`

SetAssetType sets AssetType field to given value.

### HasAssetType

`func (o *GetOrdersItemOrders) HasAssetType() bool`

HasAssetType returns a boolean if a field has been set.

### GetAvgFilledPrice

`func (o *GetOrdersItemOrders) GetAvgFilledPrice() string`

GetAvgFilledPrice returns the AvgFilledPrice field if non-nil, zero value otherwise.

### GetAvgFilledPriceOk

`func (o *GetOrdersItemOrders) GetAvgFilledPriceOk() (*string, bool)`

GetAvgFilledPriceOk returns a tuple with the AvgFilledPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgFilledPrice

`func (o *GetOrdersItemOrders) SetAvgFilledPrice(v string)`

SetAvgFilledPrice sets AvgFilledPrice field to given value.

### HasAvgFilledPrice

`func (o *GetOrdersItemOrders) HasAvgFilledPrice() bool`

HasAvgFilledPrice returns a boolean if a field has been set.

### GetBrokerId

`func (o *GetOrdersItemOrders) GetBrokerId() int32`

GetBrokerId returns the BrokerId field if non-nil, zero value otherwise.

### GetBrokerIdOk

`func (o *GetOrdersItemOrders) GetBrokerIdOk() (*int32, bool)`

GetBrokerIdOk returns a tuple with the BrokerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerId

`func (o *GetOrdersItemOrders) SetBrokerId(v int32)`

SetBrokerId sets BrokerId field to given value.

### HasBrokerId

`func (o *GetOrdersItemOrders) HasBrokerId() bool`

HasBrokerId returns a boolean if a field has been set.

### GetCanCancel

`func (o *GetOrdersItemOrders) GetCanCancel() bool`

GetCanCancel returns the CanCancel field if non-nil, zero value otherwise.

### GetCanCancelOk

`func (o *GetOrdersItemOrders) GetCanCancelOk() (*bool, bool)`

GetCanCancelOk returns a tuple with the CanCancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCancel

`func (o *GetOrdersItemOrders) SetCanCancel(v bool)`

SetCanCancel sets CanCancel field to given value.

### HasCanCancel

`func (o *GetOrdersItemOrders) HasCanCancel() bool`

HasCanCancel returns a boolean if a field has been set.

### GetCanModify

`func (o *GetOrdersItemOrders) GetCanModify() bool`

GetCanModify returns the CanModify field if non-nil, zero value otherwise.

### GetCanModifyOk

`func (o *GetOrdersItemOrders) GetCanModifyOk() (*bool, bool)`

GetCanModifyOk returns a tuple with the CanModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanModify

`func (o *GetOrdersItemOrders) SetCanModify(v bool)`

SetCanModify sets CanModify field to given value.

### HasCanModify

`func (o *GetOrdersItemOrders) HasCanModify() bool`

HasCanModify returns a boolean if a field has been set.

### GetCreateTime

`func (o *GetOrdersItemOrders) GetCreateTime() string`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *GetOrdersItemOrders) GetCreateTimeOk() (*string, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *GetOrdersItemOrders) SetCreateTime(v string)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *GetOrdersItemOrders) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCreateTime0

`func (o *GetOrdersItemOrders) GetCreateTime0() int64`

GetCreateTime0 returns the CreateTime0 field if non-nil, zero value otherwise.

### GetCreateTime0Ok

`func (o *GetOrdersItemOrders) GetCreateTime0Ok() (*int64, bool)`

GetCreateTime0Ok returns a tuple with the CreateTime0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime0

`func (o *GetOrdersItemOrders) SetCreateTime0(v int64)`

SetCreateTime0 sets CreateTime0 field to given value.

### HasCreateTime0

`func (o *GetOrdersItemOrders) HasCreateTime0() bool`

HasCreateTime0 returns a boolean if a field has been set.

### GetFilledQuantity

`func (o *GetOrdersItemOrders) GetFilledQuantity() string`

GetFilledQuantity returns the FilledQuantity field if non-nil, zero value otherwise.

### GetFilledQuantityOk

`func (o *GetOrdersItemOrders) GetFilledQuantityOk() (*string, bool)`

GetFilledQuantityOk returns a tuple with the FilledQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilledQuantity

`func (o *GetOrdersItemOrders) SetFilledQuantity(v string)`

SetFilledQuantity sets FilledQuantity field to given value.

### HasFilledQuantity

`func (o *GetOrdersItemOrders) HasFilledQuantity() bool`

HasFilledQuantity returns a boolean if a field has been set.

### GetFilledTime

`func (o *GetOrdersItemOrders) GetFilledTime() string`

GetFilledTime returns the FilledTime field if non-nil, zero value otherwise.

### GetFilledTimeOk

`func (o *GetOrdersItemOrders) GetFilledTimeOk() (*string, bool)`

GetFilledTimeOk returns a tuple with the FilledTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilledTime

`func (o *GetOrdersItemOrders) SetFilledTime(v string)`

SetFilledTime sets FilledTime field to given value.

### HasFilledTime

`func (o *GetOrdersItemOrders) HasFilledTime() bool`

HasFilledTime returns a boolean if a field has been set.

### GetFilledTime0

`func (o *GetOrdersItemOrders) GetFilledTime0() float32`

GetFilledTime0 returns the FilledTime0 field if non-nil, zero value otherwise.

### GetFilledTime0Ok

`func (o *GetOrdersItemOrders) GetFilledTime0Ok() (*float32, bool)`

GetFilledTime0Ok returns a tuple with the FilledTime0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilledTime0

`func (o *GetOrdersItemOrders) SetFilledTime0(v float32)`

SetFilledTime0 sets FilledTime0 field to given value.

### HasFilledTime0

`func (o *GetOrdersItemOrders) HasFilledTime0() bool`

HasFilledTime0 returns a boolean if a field has been set.

### GetFilledValue

`func (o *GetOrdersItemOrders) GetFilledValue() string`

GetFilledValue returns the FilledValue field if non-nil, zero value otherwise.

### GetFilledValueOk

`func (o *GetOrdersItemOrders) GetFilledValueOk() (*string, bool)`

GetFilledValueOk returns a tuple with the FilledValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilledValue

`func (o *GetOrdersItemOrders) SetFilledValue(v string)`

SetFilledValue sets FilledValue field to given value.

### HasFilledValue

`func (o *GetOrdersItemOrders) HasFilledValue() bool`

HasFilledValue returns a boolean if a field has been set.

### GetOptionContractMultiplier

`func (o *GetOrdersItemOrders) GetOptionContractMultiplier() string`

GetOptionContractMultiplier returns the OptionContractMultiplier field if non-nil, zero value otherwise.

### GetOptionContractMultiplierOk

`func (o *GetOrdersItemOrders) GetOptionContractMultiplierOk() (*string, bool)`

GetOptionContractMultiplierOk returns a tuple with the OptionContractMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionContractMultiplier

`func (o *GetOrdersItemOrders) SetOptionContractMultiplier(v string)`

SetOptionContractMultiplier sets OptionContractMultiplier field to given value.

### HasOptionContractMultiplier

`func (o *GetOrdersItemOrders) HasOptionContractMultiplier() bool`

HasOptionContractMultiplier returns a boolean if a field has been set.

### GetOptionExercisePrice

`func (o *GetOrdersItemOrders) GetOptionExercisePrice() string`

GetOptionExercisePrice returns the OptionExercisePrice field if non-nil, zero value otherwise.

### GetOptionExercisePriceOk

`func (o *GetOrdersItemOrders) GetOptionExercisePriceOk() (*string, bool)`

GetOptionExercisePriceOk returns a tuple with the OptionExercisePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionExercisePrice

`func (o *GetOrdersItemOrders) SetOptionExercisePrice(v string)`

SetOptionExercisePrice sets OptionExercisePrice field to given value.

### HasOptionExercisePrice

`func (o *GetOrdersItemOrders) HasOptionExercisePrice() bool`

HasOptionExercisePrice returns a boolean if a field has been set.

### GetOptionCycle

`func (o *GetOrdersItemOrders) GetOptionCycle() int64`

GetOptionCycle returns the OptionCycle field if non-nil, zero value otherwise.

### GetOptionCycleOk

`func (o *GetOrdersItemOrders) GetOptionCycleOk() (*int64, bool)`

GetOptionCycleOk returns a tuple with the OptionCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionCycle

`func (o *GetOrdersItemOrders) SetOptionCycle(v int64)`

SetOptionCycle sets OptionCycle field to given value.

### HasOptionCycle

`func (o *GetOrdersItemOrders) HasOptionCycle() bool`

HasOptionCycle returns a boolean if a field has been set.

### GetOrderId

`func (o *GetOrdersItemOrders) GetOrderId() int32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *GetOrdersItemOrders) GetOrderIdOk() (*int32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *GetOrdersItemOrders) SetOrderId(v int32)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *GetOrdersItemOrders) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrderType

`func (o *GetOrdersItemOrders) GetOrderType() OrderType`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *GetOrdersItemOrders) GetOrderTypeOk() (*OrderType, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *GetOrdersItemOrders) SetOrderType(v OrderType)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *GetOrdersItemOrders) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetOptionType

`func (o *GetOrdersItemOrders) GetOptionType() OptionType`

GetOptionType returns the OptionType field if non-nil, zero value otherwise.

### GetOptionTypeOk

`func (o *GetOrdersItemOrders) GetOptionTypeOk() (*OptionType, bool)`

GetOptionTypeOk returns a tuple with the OptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionType

`func (o *GetOrdersItemOrders) SetOptionType(v OptionType)`

SetOptionType sets OptionType field to given value.

### HasOptionType

`func (o *GetOrdersItemOrders) HasOptionType() bool`

HasOptionType returns a boolean if a field has been set.

### GetOptionExpirationDate

`func (o *GetOrdersItemOrders) GetOptionExpirationDate() string`

GetOptionExpirationDate returns the OptionExpirationDate field if non-nil, zero value otherwise.

### GetOptionExpirationDateOk

`func (o *GetOrdersItemOrders) GetOptionExpirationDateOk() (*string, bool)`

GetOptionExpirationDateOk returns a tuple with the OptionExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionExpirationDate

`func (o *GetOrdersItemOrders) SetOptionExpirationDate(v string)`

SetOptionExpirationDate sets OptionExpirationDate field to given value.

### HasOptionExpirationDate

`func (o *GetOrdersItemOrders) HasOptionExpirationDate() bool`

HasOptionExpirationDate returns a boolean if a field has been set.

### GetRelation

`func (o *GetOrdersItemOrders) GetRelation() string`

GetRelation returns the Relation field if non-nil, zero value otherwise.

### GetRelationOk

`func (o *GetOrdersItemOrders) GetRelationOk() (*string, bool)`

GetRelationOk returns a tuple with the Relation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelation

`func (o *GetOrdersItemOrders) SetRelation(v string)`

SetRelation sets Relation field to given value.

### HasRelation

`func (o *GetOrdersItemOrders) HasRelation() bool`

HasRelation returns a boolean if a field has been set.

### GetOptionCategory

`func (o *GetOrdersItemOrders) GetOptionCategory() string`

GetOptionCategory returns the OptionCategory field if non-nil, zero value otherwise.

### GetOptionCategoryOk

`func (o *GetOrdersItemOrders) GetOptionCategoryOk() (*string, bool)`

GetOptionCategoryOk returns a tuple with the OptionCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionCategory

`func (o *GetOrdersItemOrders) SetOptionCategory(v string)`

SetOptionCategory sets OptionCategory field to given value.

### HasOptionCategory

`func (o *GetOrdersItemOrders) HasOptionCategory() bool`

HasOptionCategory returns a boolean if a field has been set.

### GetRemainQuantity

`func (o *GetOrdersItemOrders) GetRemainQuantity() string`

GetRemainQuantity returns the RemainQuantity field if non-nil, zero value otherwise.

### GetRemainQuantityOk

`func (o *GetOrdersItemOrders) GetRemainQuantityOk() (*string, bool)`

GetRemainQuantityOk returns a tuple with the RemainQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainQuantity

`func (o *GetOrdersItemOrders) SetRemainQuantity(v string)`

SetRemainQuantity sets RemainQuantity field to given value.

### HasRemainQuantity

`func (o *GetOrdersItemOrders) HasRemainQuantity() bool`

HasRemainQuantity returns a boolean if a field has been set.

### GetStatusCode

`func (o *GetOrdersItemOrders) GetStatusCode() string`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *GetOrdersItemOrders) GetStatusCodeOk() (*string, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *GetOrdersItemOrders) SetStatusCode(v string)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *GetOrdersItemOrders) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetStatusStr

`func (o *GetOrdersItemOrders) GetStatusStr() string`

GetStatusStr returns the StatusStr field if non-nil, zero value otherwise.

### GetStatusStrOk

`func (o *GetOrdersItemOrders) GetStatusStrOk() (*string, bool)`

GetStatusStrOk returns a tuple with the StatusStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusStr

`func (o *GetOrdersItemOrders) SetStatusStr(v string)`

SetStatusStr sets StatusStr field to given value.

### HasStatusStr

`func (o *GetOrdersItemOrders) HasStatusStr() bool`

HasStatusStr returns a boolean if a field has been set.

### GetSymbol

`func (o *GetOrdersItemOrders) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetOrdersItemOrders) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetOrdersItemOrders) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *GetOrdersItemOrders) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTicker

`func (o *GetOrdersItemOrders) GetTicker() GetOrdersItemTicker`

GetTicker returns the Ticker field if non-nil, zero value otherwise.

### GetTickerOk

`func (o *GetOrdersItemOrders) GetTickerOk() (*GetOrdersItemTicker, bool)`

GetTickerOk returns a tuple with the Ticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicker

`func (o *GetOrdersItemOrders) SetTicker(v GetOrdersItemTicker)`

SetTicker sets Ticker field to given value.

### HasTicker

`func (o *GetOrdersItemOrders) HasTicker() bool`

HasTicker returns a boolean if a field has been set.

### GetTickerId

`func (o *GetOrdersItemOrders) GetTickerId() int32`

GetTickerId returns the TickerId field if non-nil, zero value otherwise.

### GetTickerIdOk

`func (o *GetOrdersItemOrders) GetTickerIdOk() (*int32, bool)`

GetTickerIdOk returns a tuple with the TickerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickerId

`func (o *GetOrdersItemOrders) SetTickerId(v int32)`

SetTickerId sets TickerId field to given value.

### HasTickerId

`func (o *GetOrdersItemOrders) HasTickerId() bool`

HasTickerId returns a boolean if a field has been set.

### GetTickerPriceDefineList

`func (o *GetOrdersItemOrders) GetTickerPriceDefineList() []GetOrdersItemTickerPriceDefineList`

GetTickerPriceDefineList returns the TickerPriceDefineList field if non-nil, zero value otherwise.

### GetTickerPriceDefineListOk

`func (o *GetOrdersItemOrders) GetTickerPriceDefineListOk() (*[]GetOrdersItemTickerPriceDefineList, bool)`

GetTickerPriceDefineListOk returns a tuple with the TickerPriceDefineList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickerPriceDefineList

`func (o *GetOrdersItemOrders) SetTickerPriceDefineList(v []GetOrdersItemTickerPriceDefineList)`

SetTickerPriceDefineList sets TickerPriceDefineList field to given value.

### HasTickerPriceDefineList

`func (o *GetOrdersItemOrders) HasTickerPriceDefineList() bool`

HasTickerPriceDefineList returns a boolean if a field has been set.

### GetTimeInForce

`func (o *GetOrdersItemOrders) GetTimeInForce() TimeInForce`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *GetOrdersItemOrders) GetTimeInForceOk() (*TimeInForce, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *GetOrdersItemOrders) SetTimeInForce(v TimeInForce)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *GetOrdersItemOrders) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetTotalQuantity

`func (o *GetOrdersItemOrders) GetTotalQuantity() string`

GetTotalQuantity returns the TotalQuantity field if non-nil, zero value otherwise.

### GetTotalQuantityOk

`func (o *GetOrdersItemOrders) GetTotalQuantityOk() (*string, bool)`

GetTotalQuantityOk returns a tuple with the TotalQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalQuantity

`func (o *GetOrdersItemOrders) SetTotalQuantity(v string)`

SetTotalQuantity sets TotalQuantity field to given value.

### HasTotalQuantity

`func (o *GetOrdersItemOrders) HasTotalQuantity() bool`

HasTotalQuantity returns a boolean if a field has been set.

### GetUpdateTime

`func (o *GetOrdersItemOrders) GetUpdateTime() string`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *GetOrdersItemOrders) GetUpdateTimeOk() (*string, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *GetOrdersItemOrders) SetUpdateTime(v string)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *GetOrdersItemOrders) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetUpdateTime0

`func (o *GetOrdersItemOrders) GetUpdateTime0() float32`

GetUpdateTime0 returns the UpdateTime0 field if non-nil, zero value otherwise.

### GetUpdateTime0Ok

`func (o *GetOrdersItemOrders) GetUpdateTime0Ok() (*float32, bool)`

GetUpdateTime0Ok returns a tuple with the UpdateTime0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime0

`func (o *GetOrdersItemOrders) SetUpdateTime0(v float32)`

SetUpdateTime0 sets UpdateTime0 field to given value.

### HasUpdateTime0

`func (o *GetOrdersItemOrders) HasUpdateTime0() bool`

HasUpdateTime0 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


