# GetSecurityAccountsResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountTypes** | Pointer to **[]string** |  | [optional] 
**AllowDeposit** | Pointer to **bool** |  | [optional] 
**BrokerAccountId** | Pointer to **string** |  | [optional] 
**BrokerId** | Pointer to **int32** |  | [optional] 
**BrokerName** | Pointer to **string** |  | [optional] 
**ComboTypes** | Pointer to **[]string** |  | [optional] 
**CustomerType** | Pointer to **string** |  | [optional] 
**Deposit** | Pointer to **bool** |  | [optional] 
**DepositStatus** | Pointer to **string** |  | [optional] 
**GiftStockStatus** | Pointer to **int32** |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 
**IsDefaultChecked** | Pointer to **bool** |  | [optional] 
**OpenAccountUrl** | Pointer to **string** |  | [optional] 
**OptionOpenStatus** | Pointer to **string** |  | [optional] 
**RegisterRegionId** | Pointer to **int32** |  | [optional] 
**RegisterTradeUrl** | Pointer to **string** |  | [optional] 
**SecAccountId** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SupportClickIPO** | Pointer to **bool** |  | [optional] 
**SupportOpenOption** | Pointer to **bool** |  | [optional] 
**SupportOutsideRth** | Pointer to **bool** |  | [optional] 
**TickerTypes** | Pointer to [**[]GetSecurityAccountsResponseTickerTypes**](GetSecurityAccountsResponseTickerTypes.md) |  | [optional] 
**TimeInForces** | Pointer to [**[]GetSecurityAccountsResponseTimeInForces**](GetSecurityAccountsResponseTimeInForces.md) |  | [optional] 
**UserTradePermissionVOs** | Pointer to [**[]GetSecurityAccountsResponseUserTradePermissionVOs**](GetSecurityAccountsResponseUserTradePermissionVOs.md) |  | [optional] 

## Methods

### NewGetSecurityAccountsResponseData

`func NewGetSecurityAccountsResponseData() *GetSecurityAccountsResponseData`

NewGetSecurityAccountsResponseData instantiates a new GetSecurityAccountsResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSecurityAccountsResponseDataWithDefaults

`func NewGetSecurityAccountsResponseDataWithDefaults() *GetSecurityAccountsResponseData`

NewGetSecurityAccountsResponseDataWithDefaults instantiates a new GetSecurityAccountsResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountTypes

`func (o *GetSecurityAccountsResponseData) GetAccountTypes() []string`

GetAccountTypes returns the AccountTypes field if non-nil, zero value otherwise.

### GetAccountTypesOk

`func (o *GetSecurityAccountsResponseData) GetAccountTypesOk() (*[]string, bool)`

GetAccountTypesOk returns a tuple with the AccountTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTypes

`func (o *GetSecurityAccountsResponseData) SetAccountTypes(v []string)`

SetAccountTypes sets AccountTypes field to given value.

### HasAccountTypes

`func (o *GetSecurityAccountsResponseData) HasAccountTypes() bool`

HasAccountTypes returns a boolean if a field has been set.

### GetAllowDeposit

`func (o *GetSecurityAccountsResponseData) GetAllowDeposit() bool`

GetAllowDeposit returns the AllowDeposit field if non-nil, zero value otherwise.

### GetAllowDepositOk

`func (o *GetSecurityAccountsResponseData) GetAllowDepositOk() (*bool, bool)`

GetAllowDepositOk returns a tuple with the AllowDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDeposit

`func (o *GetSecurityAccountsResponseData) SetAllowDeposit(v bool)`

SetAllowDeposit sets AllowDeposit field to given value.

### HasAllowDeposit

`func (o *GetSecurityAccountsResponseData) HasAllowDeposit() bool`

HasAllowDeposit returns a boolean if a field has been set.

### GetBrokerAccountId

`func (o *GetSecurityAccountsResponseData) GetBrokerAccountId() string`

GetBrokerAccountId returns the BrokerAccountId field if non-nil, zero value otherwise.

### GetBrokerAccountIdOk

`func (o *GetSecurityAccountsResponseData) GetBrokerAccountIdOk() (*string, bool)`

GetBrokerAccountIdOk returns a tuple with the BrokerAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerAccountId

`func (o *GetSecurityAccountsResponseData) SetBrokerAccountId(v string)`

SetBrokerAccountId sets BrokerAccountId field to given value.

### HasBrokerAccountId

`func (o *GetSecurityAccountsResponseData) HasBrokerAccountId() bool`

HasBrokerAccountId returns a boolean if a field has been set.

### GetBrokerId

`func (o *GetSecurityAccountsResponseData) GetBrokerId() int32`

GetBrokerId returns the BrokerId field if non-nil, zero value otherwise.

### GetBrokerIdOk

`func (o *GetSecurityAccountsResponseData) GetBrokerIdOk() (*int32, bool)`

GetBrokerIdOk returns a tuple with the BrokerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerId

`func (o *GetSecurityAccountsResponseData) SetBrokerId(v int32)`

SetBrokerId sets BrokerId field to given value.

### HasBrokerId

`func (o *GetSecurityAccountsResponseData) HasBrokerId() bool`

HasBrokerId returns a boolean if a field has been set.

### GetBrokerName

`func (o *GetSecurityAccountsResponseData) GetBrokerName() string`

GetBrokerName returns the BrokerName field if non-nil, zero value otherwise.

### GetBrokerNameOk

`func (o *GetSecurityAccountsResponseData) GetBrokerNameOk() (*string, bool)`

GetBrokerNameOk returns a tuple with the BrokerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerName

`func (o *GetSecurityAccountsResponseData) SetBrokerName(v string)`

SetBrokerName sets BrokerName field to given value.

### HasBrokerName

`func (o *GetSecurityAccountsResponseData) HasBrokerName() bool`

HasBrokerName returns a boolean if a field has been set.

### GetComboTypes

`func (o *GetSecurityAccountsResponseData) GetComboTypes() []string`

GetComboTypes returns the ComboTypes field if non-nil, zero value otherwise.

### GetComboTypesOk

`func (o *GetSecurityAccountsResponseData) GetComboTypesOk() (*[]string, bool)`

GetComboTypesOk returns a tuple with the ComboTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComboTypes

`func (o *GetSecurityAccountsResponseData) SetComboTypes(v []string)`

SetComboTypes sets ComboTypes field to given value.

### HasComboTypes

`func (o *GetSecurityAccountsResponseData) HasComboTypes() bool`

HasComboTypes returns a boolean if a field has been set.

### GetCustomerType

`func (o *GetSecurityAccountsResponseData) GetCustomerType() string`

GetCustomerType returns the CustomerType field if non-nil, zero value otherwise.

### GetCustomerTypeOk

`func (o *GetSecurityAccountsResponseData) GetCustomerTypeOk() (*string, bool)`

GetCustomerTypeOk returns a tuple with the CustomerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerType

`func (o *GetSecurityAccountsResponseData) SetCustomerType(v string)`

SetCustomerType sets CustomerType field to given value.

### HasCustomerType

`func (o *GetSecurityAccountsResponseData) HasCustomerType() bool`

HasCustomerType returns a boolean if a field has been set.

### GetDeposit

`func (o *GetSecurityAccountsResponseData) GetDeposit() bool`

GetDeposit returns the Deposit field if non-nil, zero value otherwise.

### GetDepositOk

`func (o *GetSecurityAccountsResponseData) GetDepositOk() (*bool, bool)`

GetDepositOk returns a tuple with the Deposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeposit

`func (o *GetSecurityAccountsResponseData) SetDeposit(v bool)`

SetDeposit sets Deposit field to given value.

### HasDeposit

`func (o *GetSecurityAccountsResponseData) HasDeposit() bool`

HasDeposit returns a boolean if a field has been set.

### GetDepositStatus

`func (o *GetSecurityAccountsResponseData) GetDepositStatus() string`

GetDepositStatus returns the DepositStatus field if non-nil, zero value otherwise.

### GetDepositStatusOk

`func (o *GetSecurityAccountsResponseData) GetDepositStatusOk() (*string, bool)`

GetDepositStatusOk returns a tuple with the DepositStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositStatus

`func (o *GetSecurityAccountsResponseData) SetDepositStatus(v string)`

SetDepositStatus sets DepositStatus field to given value.

### HasDepositStatus

`func (o *GetSecurityAccountsResponseData) HasDepositStatus() bool`

HasDepositStatus returns a boolean if a field has been set.

### GetGiftStockStatus

`func (o *GetSecurityAccountsResponseData) GetGiftStockStatus() int32`

GetGiftStockStatus returns the GiftStockStatus field if non-nil, zero value otherwise.

### GetGiftStockStatusOk

`func (o *GetSecurityAccountsResponseData) GetGiftStockStatusOk() (*int32, bool)`

GetGiftStockStatusOk returns a tuple with the GiftStockStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftStockStatus

`func (o *GetSecurityAccountsResponseData) SetGiftStockStatus(v int32)`

SetGiftStockStatus sets GiftStockStatus field to given value.

### HasGiftStockStatus

`func (o *GetSecurityAccountsResponseData) HasGiftStockStatus() bool`

HasGiftStockStatus returns a boolean if a field has been set.

### GetIsDefault

`func (o *GetSecurityAccountsResponseData) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *GetSecurityAccountsResponseData) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *GetSecurityAccountsResponseData) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *GetSecurityAccountsResponseData) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetIsDefaultChecked

`func (o *GetSecurityAccountsResponseData) GetIsDefaultChecked() bool`

GetIsDefaultChecked returns the IsDefaultChecked field if non-nil, zero value otherwise.

### GetIsDefaultCheckedOk

`func (o *GetSecurityAccountsResponseData) GetIsDefaultCheckedOk() (*bool, bool)`

GetIsDefaultCheckedOk returns a tuple with the IsDefaultChecked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultChecked

`func (o *GetSecurityAccountsResponseData) SetIsDefaultChecked(v bool)`

SetIsDefaultChecked sets IsDefaultChecked field to given value.

### HasIsDefaultChecked

`func (o *GetSecurityAccountsResponseData) HasIsDefaultChecked() bool`

HasIsDefaultChecked returns a boolean if a field has been set.

### GetOpenAccountUrl

`func (o *GetSecurityAccountsResponseData) GetOpenAccountUrl() string`

GetOpenAccountUrl returns the OpenAccountUrl field if non-nil, zero value otherwise.

### GetOpenAccountUrlOk

`func (o *GetSecurityAccountsResponseData) GetOpenAccountUrlOk() (*string, bool)`

GetOpenAccountUrlOk returns a tuple with the OpenAccountUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenAccountUrl

`func (o *GetSecurityAccountsResponseData) SetOpenAccountUrl(v string)`

SetOpenAccountUrl sets OpenAccountUrl field to given value.

### HasOpenAccountUrl

`func (o *GetSecurityAccountsResponseData) HasOpenAccountUrl() bool`

HasOpenAccountUrl returns a boolean if a field has been set.

### GetOptionOpenStatus

`func (o *GetSecurityAccountsResponseData) GetOptionOpenStatus() string`

GetOptionOpenStatus returns the OptionOpenStatus field if non-nil, zero value otherwise.

### GetOptionOpenStatusOk

`func (o *GetSecurityAccountsResponseData) GetOptionOpenStatusOk() (*string, bool)`

GetOptionOpenStatusOk returns a tuple with the OptionOpenStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionOpenStatus

`func (o *GetSecurityAccountsResponseData) SetOptionOpenStatus(v string)`

SetOptionOpenStatus sets OptionOpenStatus field to given value.

### HasOptionOpenStatus

`func (o *GetSecurityAccountsResponseData) HasOptionOpenStatus() bool`

HasOptionOpenStatus returns a boolean if a field has been set.

### GetRegisterRegionId

`func (o *GetSecurityAccountsResponseData) GetRegisterRegionId() int32`

GetRegisterRegionId returns the RegisterRegionId field if non-nil, zero value otherwise.

### GetRegisterRegionIdOk

`func (o *GetSecurityAccountsResponseData) GetRegisterRegionIdOk() (*int32, bool)`

GetRegisterRegionIdOk returns a tuple with the RegisterRegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterRegionId

`func (o *GetSecurityAccountsResponseData) SetRegisterRegionId(v int32)`

SetRegisterRegionId sets RegisterRegionId field to given value.

### HasRegisterRegionId

`func (o *GetSecurityAccountsResponseData) HasRegisterRegionId() bool`

HasRegisterRegionId returns a boolean if a field has been set.

### GetRegisterTradeUrl

`func (o *GetSecurityAccountsResponseData) GetRegisterTradeUrl() string`

GetRegisterTradeUrl returns the RegisterTradeUrl field if non-nil, zero value otherwise.

### GetRegisterTradeUrlOk

`func (o *GetSecurityAccountsResponseData) GetRegisterTradeUrlOk() (*string, bool)`

GetRegisterTradeUrlOk returns a tuple with the RegisterTradeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterTradeUrl

`func (o *GetSecurityAccountsResponseData) SetRegisterTradeUrl(v string)`

SetRegisterTradeUrl sets RegisterTradeUrl field to given value.

### HasRegisterTradeUrl

`func (o *GetSecurityAccountsResponseData) HasRegisterTradeUrl() bool`

HasRegisterTradeUrl returns a boolean if a field has been set.

### GetSecAccountId

`func (o *GetSecurityAccountsResponseData) GetSecAccountId() int32`

GetSecAccountId returns the SecAccountId field if non-nil, zero value otherwise.

### GetSecAccountIdOk

`func (o *GetSecurityAccountsResponseData) GetSecAccountIdOk() (*int32, bool)`

GetSecAccountIdOk returns a tuple with the SecAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecAccountId

`func (o *GetSecurityAccountsResponseData) SetSecAccountId(v int32)`

SetSecAccountId sets SecAccountId field to given value.

### HasSecAccountId

`func (o *GetSecurityAccountsResponseData) HasSecAccountId() bool`

HasSecAccountId returns a boolean if a field has been set.

### GetStatus

`func (o *GetSecurityAccountsResponseData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetSecurityAccountsResponseData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetSecurityAccountsResponseData) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetSecurityAccountsResponseData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSupportClickIPO

`func (o *GetSecurityAccountsResponseData) GetSupportClickIPO() bool`

GetSupportClickIPO returns the SupportClickIPO field if non-nil, zero value otherwise.

### GetSupportClickIPOOk

`func (o *GetSecurityAccountsResponseData) GetSupportClickIPOOk() (*bool, bool)`

GetSupportClickIPOOk returns a tuple with the SupportClickIPO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportClickIPO

`func (o *GetSecurityAccountsResponseData) SetSupportClickIPO(v bool)`

SetSupportClickIPO sets SupportClickIPO field to given value.

### HasSupportClickIPO

`func (o *GetSecurityAccountsResponseData) HasSupportClickIPO() bool`

HasSupportClickIPO returns a boolean if a field has been set.

### GetSupportOpenOption

`func (o *GetSecurityAccountsResponseData) GetSupportOpenOption() bool`

GetSupportOpenOption returns the SupportOpenOption field if non-nil, zero value otherwise.

### GetSupportOpenOptionOk

`func (o *GetSecurityAccountsResponseData) GetSupportOpenOptionOk() (*bool, bool)`

GetSupportOpenOptionOk returns a tuple with the SupportOpenOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportOpenOption

`func (o *GetSecurityAccountsResponseData) SetSupportOpenOption(v bool)`

SetSupportOpenOption sets SupportOpenOption field to given value.

### HasSupportOpenOption

`func (o *GetSecurityAccountsResponseData) HasSupportOpenOption() bool`

HasSupportOpenOption returns a boolean if a field has been set.

### GetSupportOutsideRth

`func (o *GetSecurityAccountsResponseData) GetSupportOutsideRth() bool`

GetSupportOutsideRth returns the SupportOutsideRth field if non-nil, zero value otherwise.

### GetSupportOutsideRthOk

`func (o *GetSecurityAccountsResponseData) GetSupportOutsideRthOk() (*bool, bool)`

GetSupportOutsideRthOk returns a tuple with the SupportOutsideRth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportOutsideRth

`func (o *GetSecurityAccountsResponseData) SetSupportOutsideRth(v bool)`

SetSupportOutsideRth sets SupportOutsideRth field to given value.

### HasSupportOutsideRth

`func (o *GetSecurityAccountsResponseData) HasSupportOutsideRth() bool`

HasSupportOutsideRth returns a boolean if a field has been set.

### GetTickerTypes

`func (o *GetSecurityAccountsResponseData) GetTickerTypes() []GetSecurityAccountsResponseTickerTypes`

GetTickerTypes returns the TickerTypes field if non-nil, zero value otherwise.

### GetTickerTypesOk

`func (o *GetSecurityAccountsResponseData) GetTickerTypesOk() (*[]GetSecurityAccountsResponseTickerTypes, bool)`

GetTickerTypesOk returns a tuple with the TickerTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickerTypes

`func (o *GetSecurityAccountsResponseData) SetTickerTypes(v []GetSecurityAccountsResponseTickerTypes)`

SetTickerTypes sets TickerTypes field to given value.

### HasTickerTypes

`func (o *GetSecurityAccountsResponseData) HasTickerTypes() bool`

HasTickerTypes returns a boolean if a field has been set.

### GetTimeInForces

`func (o *GetSecurityAccountsResponseData) GetTimeInForces() []GetSecurityAccountsResponseTimeInForces`

GetTimeInForces returns the TimeInForces field if non-nil, zero value otherwise.

### GetTimeInForcesOk

`func (o *GetSecurityAccountsResponseData) GetTimeInForcesOk() (*[]GetSecurityAccountsResponseTimeInForces, bool)`

GetTimeInForcesOk returns a tuple with the TimeInForces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForces

`func (o *GetSecurityAccountsResponseData) SetTimeInForces(v []GetSecurityAccountsResponseTimeInForces)`

SetTimeInForces sets TimeInForces field to given value.

### HasTimeInForces

`func (o *GetSecurityAccountsResponseData) HasTimeInForces() bool`

HasTimeInForces returns a boolean if a field has been set.

### GetUserTradePermissionVOs

`func (o *GetSecurityAccountsResponseData) GetUserTradePermissionVOs() []GetSecurityAccountsResponseUserTradePermissionVOs`

GetUserTradePermissionVOs returns the UserTradePermissionVOs field if non-nil, zero value otherwise.

### GetUserTradePermissionVOsOk

`func (o *GetSecurityAccountsResponseData) GetUserTradePermissionVOsOk() (*[]GetSecurityAccountsResponseUserTradePermissionVOs, bool)`

GetUserTradePermissionVOsOk returns a tuple with the UserTradePermissionVOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTradePermissionVOs

`func (o *GetSecurityAccountsResponseData) SetUserTradePermissionVOs(v []GetSecurityAccountsResponseUserTradePermissionVOs)`

SetUserTradePermissionVOs sets UserTradePermissionVOs field to given value.

### HasUserTradePermissionVOs

`func (o *GetSecurityAccountsResponseData) HasUserTradePermissionVOs() bool`

HasUserTradePermissionVOs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


