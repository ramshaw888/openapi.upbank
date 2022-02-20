# GetWebhookResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**WebhookResource**](WebhookResource.md) | The webhook returned in this response.  | 

## Methods

### NewGetWebhookResponse

`func NewGetWebhookResponse(data WebhookResource, ) *GetWebhookResponse`

NewGetWebhookResponse instantiates a new GetWebhookResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWebhookResponseWithDefaults

`func NewGetWebhookResponseWithDefaults() *GetWebhookResponse`

NewGetWebhookResponseWithDefaults instantiates a new GetWebhookResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetWebhookResponse) GetData() WebhookResource`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetWebhookResponse) GetDataOk() (*WebhookResource, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetWebhookResponse) SetData(v WebhookResource)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


