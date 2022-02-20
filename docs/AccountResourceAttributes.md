# AccountResourceAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** | The name associated with the account in the Up application.  | 
**AccountType** | [**AccountTypeEnum**](AccountTypeEnum.md) | The bank account type of this account.  | 
**OwnershipType** | [**OwnershipTypeEnum**](OwnershipTypeEnum.md) | The ownership structure for this account.  | 
**Balance** | [**MoneyObject**](MoneyObject.md) | The available balance of the account, taking into account any amounts that are currently on hold.  | 
**CreatedAt** | **time.Time** | The date-time at which this account was first opened.  | 

## Methods

### NewAccountResourceAttributes

`func NewAccountResourceAttributes(displayName string, accountType AccountTypeEnum, ownershipType OwnershipTypeEnum, balance MoneyObject, createdAt time.Time, ) *AccountResourceAttributes`

NewAccountResourceAttributes instantiates a new AccountResourceAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountResourceAttributesWithDefaults

`func NewAccountResourceAttributesWithDefaults() *AccountResourceAttributes`

NewAccountResourceAttributesWithDefaults instantiates a new AccountResourceAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *AccountResourceAttributes) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AccountResourceAttributes) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AccountResourceAttributes) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetAccountType

`func (o *AccountResourceAttributes) GetAccountType() AccountTypeEnum`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *AccountResourceAttributes) GetAccountTypeOk() (*AccountTypeEnum, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *AccountResourceAttributes) SetAccountType(v AccountTypeEnum)`

SetAccountType sets AccountType field to given value.


### SetAccountTypeNil

`func (o *AccountResourceAttributes) SetAccountTypeNil(b bool)`

 SetAccountTypeNil sets the value for AccountType to be an explicit nil

### UnsetAccountType
`func (o *AccountResourceAttributes) UnsetAccountType()`

UnsetAccountType ensures that no value is present for AccountType, not even an explicit nil
### GetOwnershipType

`func (o *AccountResourceAttributes) GetOwnershipType() OwnershipTypeEnum`

GetOwnershipType returns the OwnershipType field if non-nil, zero value otherwise.

### GetOwnershipTypeOk

`func (o *AccountResourceAttributes) GetOwnershipTypeOk() (*OwnershipTypeEnum, bool)`

GetOwnershipTypeOk returns a tuple with the OwnershipType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnershipType

`func (o *AccountResourceAttributes) SetOwnershipType(v OwnershipTypeEnum)`

SetOwnershipType sets OwnershipType field to given value.


### SetOwnershipTypeNil

`func (o *AccountResourceAttributes) SetOwnershipTypeNil(b bool)`

 SetOwnershipTypeNil sets the value for OwnershipType to be an explicit nil

### UnsetOwnershipType
`func (o *AccountResourceAttributes) UnsetOwnershipType()`

UnsetOwnershipType ensures that no value is present for OwnershipType, not even an explicit nil
### GetBalance

`func (o *AccountResourceAttributes) GetBalance() MoneyObject`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *AccountResourceAttributes) GetBalanceOk() (*MoneyObject, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *AccountResourceAttributes) SetBalance(v MoneyObject)`

SetBalance sets Balance field to given value.


### SetBalanceNil

`func (o *AccountResourceAttributes) SetBalanceNil(b bool)`

 SetBalanceNil sets the value for Balance to be an explicit nil

### UnsetBalance
`func (o *AccountResourceAttributes) UnsetBalance()`

UnsetBalance ensures that no value is present for Balance, not even an explicit nil
### GetCreatedAt

`func (o *AccountResourceAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccountResourceAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccountResourceAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


