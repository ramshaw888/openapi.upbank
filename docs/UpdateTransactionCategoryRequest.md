# UpdateTransactionCategoryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**NullableCategoryInputResourceIdentifier**](CategoryInputResourceIdentifier.md) | The category to set on the transaction. Set this entire key to &#x60;null&#x60; de-categorize a transaction.  | 

## Methods

### NewUpdateTransactionCategoryRequest

`func NewUpdateTransactionCategoryRequest(data NullableCategoryInputResourceIdentifier, ) *UpdateTransactionCategoryRequest`

NewUpdateTransactionCategoryRequest instantiates a new UpdateTransactionCategoryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTransactionCategoryRequestWithDefaults

`func NewUpdateTransactionCategoryRequestWithDefaults() *UpdateTransactionCategoryRequest`

NewUpdateTransactionCategoryRequestWithDefaults instantiates a new UpdateTransactionCategoryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *UpdateTransactionCategoryRequest) GetData() CategoryInputResourceIdentifier`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UpdateTransactionCategoryRequest) GetDataOk() (*CategoryInputResourceIdentifier, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UpdateTransactionCategoryRequest) SetData(v CategoryInputResourceIdentifier)`

SetData sets Data field to given value.


### SetDataNil

`func (o *UpdateTransactionCategoryRequest) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *UpdateTransactionCategoryRequest) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


