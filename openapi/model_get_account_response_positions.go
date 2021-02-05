/*
 * Webull API
 *
 * Webull API Documentation
 *
 * API version: 3.0.1
 * Contact: austin.millan@gmail.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// GetAccountResponsePositions struct for GetAccountResponsePositions
type GetAccountResponsePositions struct {
	AssetType *string `json:"assetType,omitempty"`
	BrokerId *int32 `json:"brokerId,omitempty"`
	Cost *string `json:"cost,omitempty"`
	CostPrice *string `json:"costPrice,omitempty"`
	Currency *string `json:"currency,omitempty"`
	ExchangeRate *string `json:"exchangeRate,omitempty"`
	Id *int32 `json:"id,omitempty"`
	LastOpenTime *string `json:"lastOpenTime,omitempty"`
	LastPrice *string `json:"lastPrice,omitempty"`
	Lock *bool `json:"lock,omitempty"`
	MarketValue *string `json:"marketValue,omitempty"`
	Position *string `json:"position,omitempty"`
	PositionProportion *string `json:"positionProportion,omitempty"`
	Ticker *GetAcccountRequestTicker `json:"ticker,omitempty"`
	UnrealizedProfitLoss *string `json:"unrealizedProfitLoss,omitempty"`
	UnrealizedProfitLossRate *string `json:"unrealizedProfitLossRate,omitempty"`
	UpdatePositionTimeStamp *int64 `json:"updatePositionTimeStamp,omitempty"`
}

// NewGetAccountResponsePositions instantiates a new GetAccountResponsePositions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccountResponsePositions() *GetAccountResponsePositions {
	this := GetAccountResponsePositions{}
	return &this
}

// NewGetAccountResponsePositionsWithDefaults instantiates a new GetAccountResponsePositions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccountResponsePositionsWithDefaults() *GetAccountResponsePositions {
	this := GetAccountResponsePositions{}
	return &this
}

// GetAssetType returns the AssetType field value if set, zero value otherwise.
func (o *GetAccountResponsePositions) GetAssetType() string {
	if o == nil || o.AssetType == nil {
		var ret string
		return ret
	}
	return *o.AssetType
}

// GetAssetTypeOk returns a tuple with the AssetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponsePositions) GetAssetTypeOk() (*string, bool) {
	if o == nil || o.AssetType == nil {
		return nil, false
	}
	return o.AssetType, true
}

// HasAssetType returns a boolean if a field has been set.
func (o *GetAccountResponsePositions) HasAssetType() bool {
	if o != nil && o.AssetType != nil {
		return true
	}

	return false
}

// SetAssetType gets a reference to the given string and assigns it to the AssetType field.
func (o *GetAccountResponsePositions) SetAssetType(v string) {
	o.AssetType = &v
}

// GetBrokerId returns the BrokerId field value if set, zero value otherwise.
func (o *GetAccountResponsePositions) GetBrokerId() int32 {
	if o == nil || o.BrokerId == nil {
		var ret int32
		return ret
	}
	return *o.BrokerId
}

// GetBrokerIdOk returns a tuple with the BrokerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponsePositions) GetBrokerIdOk() (*int32, bool) {
	if o == nil || o.BrokerId == nil {
		return nil, false
	}
	return o.BrokerId, true
}

// HasBrokerId returns a boolean if a field has been set.
func (o *GetAccountResponsePositions) HasBrokerId() bool {
	if o != nil && o.BrokerId != nil {
		return true
	}

	return false
}

// SetBrokerId gets a reference to the given int32 and assigns it to the BrokerId field.
func (o *GetAccountResponsePositions) SetBrokerId(v int32) {
	o.BrokerId = &v
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *GetAccountResponsePositions) GetCost() string {
	if o == nil || o.Cost == nil {
		var ret string
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponsePositions) GetCostOk() (*string, bool) {
	if o == nil || o.Cost == nil {
		return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *GetAccountResponsePositions) HasCost() bool {
	if o != nil && o.Cost != nil {
		return true
	}

	return false
}

// SetCost gets a reference to the given string and assigns it to the Cost field.
func (o *GetAccountResponsePositions) SetCost(v string) {
	o.Cost = &v
}

// GetCostPrice returns the CostPrice field value if set, zero value otherwise.
func (o *GetAccountResponsePositions) GetCostPrice() string {
	if o == nil || o.CostPrice == nil {
		var ret string
		return ret
	}
	return *o.CostPrice
}

// GetCostPriceOk returns a tuple with the CostPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponsePositions) GetCostPriceOk() (*string, bool) {
	if o == nil || o.CostPrice == nil {
		return nil, false
	}
	return o.CostPrice, true
}

// HasCostPrice returns a boolean if a field has been set.
func (o *GetAccountResponsePositions) HasCostPrice() bool {
	if o != nil && o.CostPrice != nil {
		return true
	}

	return false
}

// SetCostPrice gets a reference to the given string and assigns it to the CostPrice field.
func (o *GetAccountResponsePositions) SetCostPrice(v string) {
	o.CostPrice = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *GetAccountResponsePositions) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponsePositions) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *GetAccountResponsePositions) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *GetAccountResponsePositions) SetCurrency(v string) {
	o.Currency = &v
}

// GetExchangeRate returns the ExchangeRate field value if set, zero value otherwise.
func (o *GetAccountResponsePositions) GetExchangeRate() string {
	if o == nil || o.ExchangeRate == nil {
		var ret string
		return ret
	}
	return *o.ExchangeRate
}

// GetExchangeRateOk returns a tuple with the ExchangeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponsePositions) GetExchangeRateOk() (*string, bool) {
	if o == nil || o.ExchangeRate == nil {
		return nil, false
	}
	return o.ExchangeRate, true
}

// HasExchangeRate returns a boolean if a field has been set.
func (o *GetAccountResponsePositions) HasExchangeRate() bool {
	if o != nil && o.ExchangeRate != nil {
		return true
	}

	return false
}

// SetExchangeRate gets a reference to the given string and assigns it to the ExchangeRate field.
func (o *GetAccountResponsePositions) SetExchangeRate(v string) {
	o.ExchangeRate = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetAccountResponsePositions) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponsePositions) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetAccountResponsePositions) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GetAccountResponsePositions) SetId(v int32) {
	o.Id = &v
}

// GetLastOpenTime returns the LastOpenTime field value if set, zero value otherwise.
func (o *GetAccountResponsePositions) GetLastOpenTime() string {
	if o == nil || o.LastOpenTime == nil {
		var ret string
		return ret
	}
	return *o.LastOpenTime
}

// GetLastOpenTimeOk returns a tuple with the LastOpenTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponsePositions) GetLastOpenTimeOk() (*string, bool) {
	if o == nil || o.LastOpenTime == nil {
		return nil, false
	}
	return o.LastOpenTime, true
}

// HasLastOpenTime returns a boolean if a field has been set.
func (o *GetAccountResponsePositions) HasLastOpenTime() bool {
	if o != nil && o.LastOpenTime != nil {
		return true
	}

	return false
}

// SetLastOpenTime gets a reference to the given string and assigns it to the LastOpenTime field.
func (o *GetAccountResponsePositions) SetLastOpenTime(v string) {
	o.LastOpenTime = &v
}

// GetLastPrice returns the LastPrice field value if set, zero value otherwise.
func (o *GetAccountResponsePositions) GetLastPrice() string {
	if o == nil || o.LastPrice == nil {
		var ret string
		return ret
	}
	return *o.LastPrice
}

// GetLastPriceOk returns a tuple with the LastPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponsePositions) GetLastPriceOk() (*string, bool) {
	if o == nil || o.LastPrice == nil {
		return nil, false
	}
	return o.LastPrice, true
}

// HasLastPrice returns a boolean if a field has been set.
func (o *GetAccountResponsePositions) HasLastPrice() bool {
	if o != nil && o.LastPrice != nil {
		return true
	}

	return false
}

// SetLastPrice gets a reference to the given string and assigns it to the LastPrice field.
func (o *GetAccountResponsePositions) SetLastPrice(v string) {
	o.LastPrice = &v
}

// GetLock returns the Lock field value if set, zero value otherwise.
func (o *GetAccountResponsePositions) GetLock() bool {
	if o == nil || o.Lock == nil {
		var ret bool
		return ret
	}
	return *o.Lock
}

// GetLockOk returns a tuple with the Lock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponsePositions) GetLockOk() (*bool, bool) {
	if o == nil || o.Lock == nil {
		return nil, false
	}
	return o.Lock, true
}

// HasLock returns a boolean if a field has been set.
func (o *GetAccountResponsePositions) HasLock() bool {
	if o != nil && o.Lock != nil {
		return true
	}

	return false
}

// SetLock gets a reference to the given bool and assigns it to the Lock field.
func (o *GetAccountResponsePositions) SetLock(v bool) {
	o.Lock = &v
}

// GetMarketValue returns the MarketValue field value if set, zero value otherwise.
func (o *GetAccountResponsePositions) GetMarketValue() string {
	if o == nil || o.MarketValue == nil {
		var ret string
		return ret
	}
	return *o.MarketValue
}

// GetMarketValueOk returns a tuple with the MarketValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponsePositions) GetMarketValueOk() (*string, bool) {
	if o == nil || o.MarketValue == nil {
		return nil, false
	}
	return o.MarketValue, true
}

// HasMarketValue returns a boolean if a field has been set.
func (o *GetAccountResponsePositions) HasMarketValue() bool {
	if o != nil && o.MarketValue != nil {
		return true
	}

	return false
}

// SetMarketValue gets a reference to the given string and assigns it to the MarketValue field.
func (o *GetAccountResponsePositions) SetMarketValue(v string) {
	o.MarketValue = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *GetAccountResponsePositions) GetPosition() string {
	if o == nil || o.Position == nil {
		var ret string
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponsePositions) GetPositionOk() (*string, bool) {
	if o == nil || o.Position == nil {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *GetAccountResponsePositions) HasPosition() bool {
	if o != nil && o.Position != nil {
		return true
	}

	return false
}

// SetPosition gets a reference to the given string and assigns it to the Position field.
func (o *GetAccountResponsePositions) SetPosition(v string) {
	o.Position = &v
}

// GetPositionProportion returns the PositionProportion field value if set, zero value otherwise.
func (o *GetAccountResponsePositions) GetPositionProportion() string {
	if o == nil || o.PositionProportion == nil {
		var ret string
		return ret
	}
	return *o.PositionProportion
}

// GetPositionProportionOk returns a tuple with the PositionProportion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponsePositions) GetPositionProportionOk() (*string, bool) {
	if o == nil || o.PositionProportion == nil {
		return nil, false
	}
	return o.PositionProportion, true
}

// HasPositionProportion returns a boolean if a field has been set.
func (o *GetAccountResponsePositions) HasPositionProportion() bool {
	if o != nil && o.PositionProportion != nil {
		return true
	}

	return false
}

// SetPositionProportion gets a reference to the given string and assigns it to the PositionProportion field.
func (o *GetAccountResponsePositions) SetPositionProportion(v string) {
	o.PositionProportion = &v
}

// GetTicker returns the Ticker field value if set, zero value otherwise.
func (o *GetAccountResponsePositions) GetTicker() GetAcccountRequestTicker {
	if o == nil || o.Ticker == nil {
		var ret GetAcccountRequestTicker
		return ret
	}
	return *o.Ticker
}

// GetTickerOk returns a tuple with the Ticker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponsePositions) GetTickerOk() (*GetAcccountRequestTicker, bool) {
	if o == nil || o.Ticker == nil {
		return nil, false
	}
	return o.Ticker, true
}

// HasTicker returns a boolean if a field has been set.
func (o *GetAccountResponsePositions) HasTicker() bool {
	if o != nil && o.Ticker != nil {
		return true
	}

	return false
}

// SetTicker gets a reference to the given GetAcccountRequestTicker and assigns it to the Ticker field.
func (o *GetAccountResponsePositions) SetTicker(v GetAcccountRequestTicker) {
	o.Ticker = &v
}

// GetUnrealizedProfitLoss returns the UnrealizedProfitLoss field value if set, zero value otherwise.
func (o *GetAccountResponsePositions) GetUnrealizedProfitLoss() string {
	if o == nil || o.UnrealizedProfitLoss == nil {
		var ret string
		return ret
	}
	return *o.UnrealizedProfitLoss
}

// GetUnrealizedProfitLossOk returns a tuple with the UnrealizedProfitLoss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponsePositions) GetUnrealizedProfitLossOk() (*string, bool) {
	if o == nil || o.UnrealizedProfitLoss == nil {
		return nil, false
	}
	return o.UnrealizedProfitLoss, true
}

// HasUnrealizedProfitLoss returns a boolean if a field has been set.
func (o *GetAccountResponsePositions) HasUnrealizedProfitLoss() bool {
	if o != nil && o.UnrealizedProfitLoss != nil {
		return true
	}

	return false
}

// SetUnrealizedProfitLoss gets a reference to the given string and assigns it to the UnrealizedProfitLoss field.
func (o *GetAccountResponsePositions) SetUnrealizedProfitLoss(v string) {
	o.UnrealizedProfitLoss = &v
}

// GetUnrealizedProfitLossRate returns the UnrealizedProfitLossRate field value if set, zero value otherwise.
func (o *GetAccountResponsePositions) GetUnrealizedProfitLossRate() string {
	if o == nil || o.UnrealizedProfitLossRate == nil {
		var ret string
		return ret
	}
	return *o.UnrealizedProfitLossRate
}

// GetUnrealizedProfitLossRateOk returns a tuple with the UnrealizedProfitLossRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponsePositions) GetUnrealizedProfitLossRateOk() (*string, bool) {
	if o == nil || o.UnrealizedProfitLossRate == nil {
		return nil, false
	}
	return o.UnrealizedProfitLossRate, true
}

// HasUnrealizedProfitLossRate returns a boolean if a field has been set.
func (o *GetAccountResponsePositions) HasUnrealizedProfitLossRate() bool {
	if o != nil && o.UnrealizedProfitLossRate != nil {
		return true
	}

	return false
}

// SetUnrealizedProfitLossRate gets a reference to the given string and assigns it to the UnrealizedProfitLossRate field.
func (o *GetAccountResponsePositions) SetUnrealizedProfitLossRate(v string) {
	o.UnrealizedProfitLossRate = &v
}

// GetUpdatePositionTimeStamp returns the UpdatePositionTimeStamp field value if set, zero value otherwise.
func (o *GetAccountResponsePositions) GetUpdatePositionTimeStamp() int64 {
	if o == nil || o.UpdatePositionTimeStamp == nil {
		var ret int64
		return ret
	}
	return *o.UpdatePositionTimeStamp
}

// GetUpdatePositionTimeStampOk returns a tuple with the UpdatePositionTimeStamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponsePositions) GetUpdatePositionTimeStampOk() (*int64, bool) {
	if o == nil || o.UpdatePositionTimeStamp == nil {
		return nil, false
	}
	return o.UpdatePositionTimeStamp, true
}

// HasUpdatePositionTimeStamp returns a boolean if a field has been set.
func (o *GetAccountResponsePositions) HasUpdatePositionTimeStamp() bool {
	if o != nil && o.UpdatePositionTimeStamp != nil {
		return true
	}

	return false
}

// SetUpdatePositionTimeStamp gets a reference to the given int64 and assigns it to the UpdatePositionTimeStamp field.
func (o *GetAccountResponsePositions) SetUpdatePositionTimeStamp(v int64) {
	o.UpdatePositionTimeStamp = &v
}

func (o GetAccountResponsePositions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AssetType != nil {
		toSerialize["assetType"] = o.AssetType
	}
	if o.BrokerId != nil {
		toSerialize["brokerId"] = o.BrokerId
	}
	if o.Cost != nil {
		toSerialize["cost"] = o.Cost
	}
	if o.CostPrice != nil {
		toSerialize["costPrice"] = o.CostPrice
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.ExchangeRate != nil {
		toSerialize["exchangeRate"] = o.ExchangeRate
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastOpenTime != nil {
		toSerialize["lastOpenTime"] = o.LastOpenTime
	}
	if o.LastPrice != nil {
		toSerialize["lastPrice"] = o.LastPrice
	}
	if o.Lock != nil {
		toSerialize["lock"] = o.Lock
	}
	if o.MarketValue != nil {
		toSerialize["marketValue"] = o.MarketValue
	}
	if o.Position != nil {
		toSerialize["position"] = o.Position
	}
	if o.PositionProportion != nil {
		toSerialize["positionProportion"] = o.PositionProportion
	}
	if o.Ticker != nil {
		toSerialize["ticker"] = o.Ticker
	}
	if o.UnrealizedProfitLoss != nil {
		toSerialize["unrealizedProfitLoss"] = o.UnrealizedProfitLoss
	}
	if o.UnrealizedProfitLossRate != nil {
		toSerialize["unrealizedProfitLossRate"] = o.UnrealizedProfitLossRate
	}
	if o.UpdatePositionTimeStamp != nil {
		toSerialize["updatePositionTimeStamp"] = o.UpdatePositionTimeStamp
	}
	return json.Marshal(toSerialize)
}

type NullableGetAccountResponsePositions struct {
	value *GetAccountResponsePositions
	isSet bool
}

func (v NullableGetAccountResponsePositions) Get() *GetAccountResponsePositions {
	return v.value
}

func (v *NullableGetAccountResponsePositions) Set(val *GetAccountResponsePositions) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccountResponsePositions) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccountResponsePositions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccountResponsePositions(val *GetAccountResponsePositions) *NullableGetAccountResponsePositions {
	return &NullableGetAccountResponsePositions{value: val, isSet: true}
}

func (v NullableGetAccountResponsePositions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccountResponsePositions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


