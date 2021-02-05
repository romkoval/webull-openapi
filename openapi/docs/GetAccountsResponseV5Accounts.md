# GetAccountsResponseV5Accounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecAccountId** | Pointer to **float32** |  | [optional] 
**BrokerId** | Pointer to **float32** |  | [optional] 
**AccountType** | Pointer to **string** |  | [optional] 
**BrokerAccountId** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**CurrencyId** | Pointer to **float32** |  | [optional] 
**Pdt** | Pointer to **bool** |  | [optional] 
**Professional** | Pointer to **bool** |  | [optional] 
**ShowUpgrade** | Pointer to **bool** |  | [optional] 
**TotalCost** | Pointer to **string** |  | [optional] 
**NetLiquidation** | Pointer to **string** |  | [optional] 
**UnrealizedProfitLoss** | Pointer to **string** |  | [optional] 
**UnrealizedProfitLossRate** | Pointer to **string** |  | [optional] 
**UnrealizedProfitLossBase** | Pointer to **string** |  | [optional] 
**Warning** | Pointer to **bool** |  | [optional] 
**RemindModifyPwd** | Pointer to **bool** |  | [optional] 
**AccountMembers** | Pointer to [**[]GetAccountsResponseV5AccountMembers**](GetAccountsResponseV5AccountMembers.md) |  | [optional] 
**OpenOrderSize** | Pointer to **float32** |  | [optional] 

## Methods

### NewGetAccountsResponseV5Accounts

`func NewGetAccountsResponseV5Accounts() *GetAccountsResponseV5Accounts`

NewGetAccountsResponseV5Accounts instantiates a new GetAccountsResponseV5Accounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountsResponseV5AccountsWithDefaults

`func NewGetAccountsResponseV5AccountsWithDefaults() *GetAccountsResponseV5Accounts`

NewGetAccountsResponseV5AccountsWithDefaults instantiates a new GetAccountsResponseV5Accounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecAccountId

`func (o *GetAccountsResponseV5Accounts) GetSecAccountId() float32`

GetSecAccountId returns the SecAccountId field if non-nil, zero value otherwise.

### GetSecAccountIdOk

`func (o *GetAccountsResponseV5Accounts) GetSecAccountIdOk() (*float32, bool)`

GetSecAccountIdOk returns a tuple with the SecAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecAccountId

`func (o *GetAccountsResponseV5Accounts) SetSecAccountId(v float32)`

SetSecAccountId sets SecAccountId field to given value.

### HasSecAccountId

`func (o *GetAccountsResponseV5Accounts) HasSecAccountId() bool`

HasSecAccountId returns a boolean if a field has been set.

### GetBrokerId

`func (o *GetAccountsResponseV5Accounts) GetBrokerId() float32`

GetBrokerId returns the BrokerId field if non-nil, zero value otherwise.

### GetBrokerIdOk

`func (o *GetAccountsResponseV5Accounts) GetBrokerIdOk() (*float32, bool)`

GetBrokerIdOk returns a tuple with the BrokerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerId

`func (o *GetAccountsResponseV5Accounts) SetBrokerId(v float32)`

SetBrokerId sets BrokerId field to given value.

### HasBrokerId

`func (o *GetAccountsResponseV5Accounts) HasBrokerId() bool`

HasBrokerId returns a boolean if a field has been set.

### GetAccountType

`func (o *GetAccountsResponseV5Accounts) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *GetAccountsResponseV5Accounts) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *GetAccountsResponseV5Accounts) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *GetAccountsResponseV5Accounts) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetBrokerAccountId

`func (o *GetAccountsResponseV5Accounts) GetBrokerAccountId() string`

GetBrokerAccountId returns the BrokerAccountId field if non-nil, zero value otherwise.

### GetBrokerAccountIdOk

`func (o *GetAccountsResponseV5Accounts) GetBrokerAccountIdOk() (*string, bool)`

GetBrokerAccountIdOk returns a tuple with the BrokerAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerAccountId

`func (o *GetAccountsResponseV5Accounts) SetBrokerAccountId(v string)`

SetBrokerAccountId sets BrokerAccountId field to given value.

### HasBrokerAccountId

`func (o *GetAccountsResponseV5Accounts) HasBrokerAccountId() bool`

HasBrokerAccountId returns a boolean if a field has been set.

### GetCurrency

`func (o *GetAccountsResponseV5Accounts) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetAccountsResponseV5Accounts) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetAccountsResponseV5Accounts) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GetAccountsResponseV5Accounts) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCurrencyId

`func (o *GetAccountsResponseV5Accounts) GetCurrencyId() float32`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *GetAccountsResponseV5Accounts) GetCurrencyIdOk() (*float32, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *GetAccountsResponseV5Accounts) SetCurrencyId(v float32)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *GetAccountsResponseV5Accounts) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetPdt

`func (o *GetAccountsResponseV5Accounts) GetPdt() bool`

GetPdt returns the Pdt field if non-nil, zero value otherwise.

### GetPdtOk

`func (o *GetAccountsResponseV5Accounts) GetPdtOk() (*bool, bool)`

GetPdtOk returns a tuple with the Pdt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdt

`func (o *GetAccountsResponseV5Accounts) SetPdt(v bool)`

SetPdt sets Pdt field to given value.

### HasPdt

`func (o *GetAccountsResponseV5Accounts) HasPdt() bool`

HasPdt returns a boolean if a field has been set.

### GetProfessional

`func (o *GetAccountsResponseV5Accounts) GetProfessional() bool`

GetProfessional returns the Professional field if non-nil, zero value otherwise.

### GetProfessionalOk

`func (o *GetAccountsResponseV5Accounts) GetProfessionalOk() (*bool, bool)`

GetProfessionalOk returns a tuple with the Professional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfessional

`func (o *GetAccountsResponseV5Accounts) SetProfessional(v bool)`

SetProfessional sets Professional field to given value.

### HasProfessional

`func (o *GetAccountsResponseV5Accounts) HasProfessional() bool`

HasProfessional returns a boolean if a field has been set.

### GetShowUpgrade

`func (o *GetAccountsResponseV5Accounts) GetShowUpgrade() bool`

GetShowUpgrade returns the ShowUpgrade field if non-nil, zero value otherwise.

### GetShowUpgradeOk

`func (o *GetAccountsResponseV5Accounts) GetShowUpgradeOk() (*bool, bool)`

GetShowUpgradeOk returns a tuple with the ShowUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowUpgrade

`func (o *GetAccountsResponseV5Accounts) SetShowUpgrade(v bool)`

SetShowUpgrade sets ShowUpgrade field to given value.

### HasShowUpgrade

`func (o *GetAccountsResponseV5Accounts) HasShowUpgrade() bool`

HasShowUpgrade returns a boolean if a field has been set.

### GetTotalCost

`func (o *GetAccountsResponseV5Accounts) GetTotalCost() string`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *GetAccountsResponseV5Accounts) GetTotalCostOk() (*string, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *GetAccountsResponseV5Accounts) SetTotalCost(v string)`

SetTotalCost sets TotalCost field to given value.

### HasTotalCost

`func (o *GetAccountsResponseV5Accounts) HasTotalCost() bool`

HasTotalCost returns a boolean if a field has been set.

### GetNetLiquidation

`func (o *GetAccountsResponseV5Accounts) GetNetLiquidation() string`

GetNetLiquidation returns the NetLiquidation field if non-nil, zero value otherwise.

### GetNetLiquidationOk

`func (o *GetAccountsResponseV5Accounts) GetNetLiquidationOk() (*string, bool)`

GetNetLiquidationOk returns a tuple with the NetLiquidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetLiquidation

`func (o *GetAccountsResponseV5Accounts) SetNetLiquidation(v string)`

SetNetLiquidation sets NetLiquidation field to given value.

### HasNetLiquidation

`func (o *GetAccountsResponseV5Accounts) HasNetLiquidation() bool`

HasNetLiquidation returns a boolean if a field has been set.

### GetUnrealizedProfitLoss

`func (o *GetAccountsResponseV5Accounts) GetUnrealizedProfitLoss() string`

GetUnrealizedProfitLoss returns the UnrealizedProfitLoss field if non-nil, zero value otherwise.

### GetUnrealizedProfitLossOk

`func (o *GetAccountsResponseV5Accounts) GetUnrealizedProfitLossOk() (*string, bool)`

GetUnrealizedProfitLossOk returns a tuple with the UnrealizedProfitLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedProfitLoss

`func (o *GetAccountsResponseV5Accounts) SetUnrealizedProfitLoss(v string)`

SetUnrealizedProfitLoss sets UnrealizedProfitLoss field to given value.

### HasUnrealizedProfitLoss

`func (o *GetAccountsResponseV5Accounts) HasUnrealizedProfitLoss() bool`

HasUnrealizedProfitLoss returns a boolean if a field has been set.

### GetUnrealizedProfitLossRate

`func (o *GetAccountsResponseV5Accounts) GetUnrealizedProfitLossRate() string`

GetUnrealizedProfitLossRate returns the UnrealizedProfitLossRate field if non-nil, zero value otherwise.

### GetUnrealizedProfitLossRateOk

`func (o *GetAccountsResponseV5Accounts) GetUnrealizedProfitLossRateOk() (*string, bool)`

GetUnrealizedProfitLossRateOk returns a tuple with the UnrealizedProfitLossRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedProfitLossRate

`func (o *GetAccountsResponseV5Accounts) SetUnrealizedProfitLossRate(v string)`

SetUnrealizedProfitLossRate sets UnrealizedProfitLossRate field to given value.

### HasUnrealizedProfitLossRate

`func (o *GetAccountsResponseV5Accounts) HasUnrealizedProfitLossRate() bool`

HasUnrealizedProfitLossRate returns a boolean if a field has been set.

### GetUnrealizedProfitLossBase

`func (o *GetAccountsResponseV5Accounts) GetUnrealizedProfitLossBase() string`

GetUnrealizedProfitLossBase returns the UnrealizedProfitLossBase field if non-nil, zero value otherwise.

### GetUnrealizedProfitLossBaseOk

`func (o *GetAccountsResponseV5Accounts) GetUnrealizedProfitLossBaseOk() (*string, bool)`

GetUnrealizedProfitLossBaseOk returns a tuple with the UnrealizedProfitLossBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedProfitLossBase

`func (o *GetAccountsResponseV5Accounts) SetUnrealizedProfitLossBase(v string)`

SetUnrealizedProfitLossBase sets UnrealizedProfitLossBase field to given value.

### HasUnrealizedProfitLossBase

`func (o *GetAccountsResponseV5Accounts) HasUnrealizedProfitLossBase() bool`

HasUnrealizedProfitLossBase returns a boolean if a field has been set.

### GetWarning

`func (o *GetAccountsResponseV5Accounts) GetWarning() bool`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *GetAccountsResponseV5Accounts) GetWarningOk() (*bool, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *GetAccountsResponseV5Accounts) SetWarning(v bool)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *GetAccountsResponseV5Accounts) HasWarning() bool`

HasWarning returns a boolean if a field has been set.

### GetRemindModifyPwd

`func (o *GetAccountsResponseV5Accounts) GetRemindModifyPwd() bool`

GetRemindModifyPwd returns the RemindModifyPwd field if non-nil, zero value otherwise.

### GetRemindModifyPwdOk

`func (o *GetAccountsResponseV5Accounts) GetRemindModifyPwdOk() (*bool, bool)`

GetRemindModifyPwdOk returns a tuple with the RemindModifyPwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemindModifyPwd

`func (o *GetAccountsResponseV5Accounts) SetRemindModifyPwd(v bool)`

SetRemindModifyPwd sets RemindModifyPwd field to given value.

### HasRemindModifyPwd

`func (o *GetAccountsResponseV5Accounts) HasRemindModifyPwd() bool`

HasRemindModifyPwd returns a boolean if a field has been set.

### GetAccountMembers

`func (o *GetAccountsResponseV5Accounts) GetAccountMembers() []GetAccountsResponseV5AccountMembers`

GetAccountMembers returns the AccountMembers field if non-nil, zero value otherwise.

### GetAccountMembersOk

`func (o *GetAccountsResponseV5Accounts) GetAccountMembersOk() (*[]GetAccountsResponseV5AccountMembers, bool)`

GetAccountMembersOk returns a tuple with the AccountMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMembers

`func (o *GetAccountsResponseV5Accounts) SetAccountMembers(v []GetAccountsResponseV5AccountMembers)`

SetAccountMembers sets AccountMembers field to given value.

### HasAccountMembers

`func (o *GetAccountsResponseV5Accounts) HasAccountMembers() bool`

HasAccountMembers returns a boolean if a field has been set.

### GetOpenOrderSize

`func (o *GetAccountsResponseV5Accounts) GetOpenOrderSize() float32`

GetOpenOrderSize returns the OpenOrderSize field if non-nil, zero value otherwise.

### GetOpenOrderSizeOk

`func (o *GetAccountsResponseV5Accounts) GetOpenOrderSizeOk() (*float32, bool)`

GetOpenOrderSizeOk returns a tuple with the OpenOrderSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenOrderSize

`func (o *GetAccountsResponseV5Accounts) SetOpenOrderSize(v float32)`

SetOpenOrderSize sets OpenOrderSize field to given value.

### HasOpenOrderSize

`func (o *GetAccountsResponseV5Accounts) HasOpenOrderSize() bool`

HasOpenOrderSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


