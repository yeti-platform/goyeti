# FileInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] [default to "file"]
**Created** | Pointer to **time.Time** |  | [optional] 
**Context** | Pointer to **[]map[string]interface{}** |  | [optional] [default to []]
**LastAnalysis** | Pointer to [**map[string]time.Time**](time.Time.md) |  | [optional] [default to {}]
**Name** | Pointer to **NullableString** |  | [optional] 
**Size** | Pointer to **NullableInt32** |  | [optional] 
**Sha256** | Pointer to **NullableString** |  | [optional] 
**Md5** | Pointer to **NullableString** |  | [optional] 
**Sha1** | Pointer to **NullableString** |  | [optional] 
**MimeType** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFileInput

`func NewFileInput(value string, ) *FileInput`

NewFileInput instantiates a new FileInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileInputWithDefaults

`func NewFileInputWithDefaults() *FileInput`

NewFileInputWithDefaults instantiates a new FileInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *FileInput) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FileInput) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FileInput) SetValue(v string)`

SetValue sets Value field to given value.


### GetType

`func (o *FileInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FileInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FileInput) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FileInput) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreated

`func (o *FileInput) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *FileInput) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *FileInput) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *FileInput) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetContext

`func (o *FileInput) GetContext() []map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *FileInput) GetContextOk() (*[]map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *FileInput) SetContext(v []map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *FileInput) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetLastAnalysis

`func (o *FileInput) GetLastAnalysis() map[string]time.Time`

GetLastAnalysis returns the LastAnalysis field if non-nil, zero value otherwise.

### GetLastAnalysisOk

`func (o *FileInput) GetLastAnalysisOk() (*map[string]time.Time, bool)`

GetLastAnalysisOk returns a tuple with the LastAnalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAnalysis

`func (o *FileInput) SetLastAnalysis(v map[string]time.Time)`

SetLastAnalysis sets LastAnalysis field to given value.

### HasLastAnalysis

`func (o *FileInput) HasLastAnalysis() bool`

HasLastAnalysis returns a boolean if a field has been set.

### GetName

`func (o *FileInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FileInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FileInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FileInput) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *FileInput) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FileInput) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSize

`func (o *FileInput) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FileInput) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FileInput) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *FileInput) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *FileInput) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *FileInput) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetSha256

`func (o *FileInput) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *FileInput) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *FileInput) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *FileInput) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### SetSha256Nil

`func (o *FileInput) SetSha256Nil(b bool)`

 SetSha256Nil sets the value for Sha256 to be an explicit nil

### UnsetSha256
`func (o *FileInput) UnsetSha256()`

UnsetSha256 ensures that no value is present for Sha256, not even an explicit nil
### GetMd5

`func (o *FileInput) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *FileInput) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *FileInput) SetMd5(v string)`

SetMd5 sets Md5 field to given value.

### HasMd5

`func (o *FileInput) HasMd5() bool`

HasMd5 returns a boolean if a field has been set.

### SetMd5Nil

`func (o *FileInput) SetMd5Nil(b bool)`

 SetMd5Nil sets the value for Md5 to be an explicit nil

### UnsetMd5
`func (o *FileInput) UnsetMd5()`

UnsetMd5 ensures that no value is present for Md5, not even an explicit nil
### GetSha1

`func (o *FileInput) GetSha1() string`

GetSha1 returns the Sha1 field if non-nil, zero value otherwise.

### GetSha1Ok

`func (o *FileInput) GetSha1Ok() (*string, bool)`

GetSha1Ok returns a tuple with the Sha1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha1

`func (o *FileInput) SetSha1(v string)`

SetSha1 sets Sha1 field to given value.

### HasSha1

`func (o *FileInput) HasSha1() bool`

HasSha1 returns a boolean if a field has been set.

### SetSha1Nil

`func (o *FileInput) SetSha1Nil(b bool)`

 SetSha1Nil sets the value for Sha1 to be an explicit nil

### UnsetSha1
`func (o *FileInput) UnsetSha1()`

UnsetSha1 ensures that no value is present for Sha1, not even an explicit nil
### GetMimeType

`func (o *FileInput) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *FileInput) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *FileInput) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *FileInput) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### SetMimeTypeNil

`func (o *FileInput) SetMimeTypeNil(b bool)`

 SetMimeTypeNil sets the value for MimeType to be an explicit nil

### UnsetMimeType
`func (o *FileInput) UnsetMimeType()`

UnsetMimeType ensures that no value is present for MimeType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


