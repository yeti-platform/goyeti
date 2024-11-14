# ResponseDetailsApiV2ObservablesObservableIdGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] [default to "url"]
**Created** | Pointer to **time.Time** |  | [optional] 
**Context** | Pointer to **[]map[string]interface{}** |  | [optional] [default to []]
**LastAnalysis** | Pointer to [**map[string]time.Time**](time.Time.md) |  | [optional] [default to {}]
**Id** | **string** |  | [readonly] 
**Tags** | [**map[string]TagRelationshipOutput**](TagRelationshipOutput.md) |  | [readonly] 
**RootType** | **string** |  | [readonly] 
**LastSeen** | Pointer to **time.Time** |  | [optional] 
**FirstSeen** | Pointer to **time.Time** |  | [optional] 
**Issuer** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**After** | Pointer to **time.Time** |  | [optional] 
**Before** | Pointer to **time.Time** |  | [optional] 
**Fingerprint** | Pointer to **string** |  | [optional] 
**Key** | **string** |  | 
**Data** | ***os.File** |  | 
**Hive** | [**RegistryHive**](RegistryHive.md) |  | 
**PathFile** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**RegitryType** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Sha256** | Pointer to **string** |  | [optional] 
**Md5** | Pointer to **string** |  | [optional] 
**Sha1** | Pointer to **string** |  | [optional] 
**MimeType** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**Credential** | Pointer to **string** |  | [optional] 
**AccountLogin** | Pointer to **string** |  | [optional] 
**AccountType** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**IsServiceAccount** | Pointer to **bool** |  | [optional] 
**IsPrivileged** | Pointer to **bool** |  | [optional] 
**CanEscalatePrivs** | Pointer to **bool** |  | [optional] 
**IsDisabled** | Pointer to **bool** |  | [optional] 
**AccountCreated** | Pointer to **time.Time** |  | [optional] 
**AccountExpires** | Pointer to **time.Time** |  | [optional] 
**CredentialLastChanged** | Pointer to **time.Time** |  | [optional] 
**AccountFirstLogin** | Pointer to **time.Time** |  | [optional] 
**AccountLastLogin** | Pointer to **time.Time** |  | [optional] 
**Coin** | Pointer to **string** |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 

## Methods

### NewResponseDetailsApiV2ObservablesObservableIdGet

`func NewResponseDetailsApiV2ObservablesObservableIdGet(value string, id string, tags map[string]TagRelationshipOutput, rootType string, key string, data *os.File, hive RegistryHive, ) *ResponseDetailsApiV2ObservablesObservableIdGet`

NewResponseDetailsApiV2ObservablesObservableIdGet instantiates a new ResponseDetailsApiV2ObservablesObservableIdGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseDetailsApiV2ObservablesObservableIdGetWithDefaults

`func NewResponseDetailsApiV2ObservablesObservableIdGetWithDefaults() *ResponseDetailsApiV2ObservablesObservableIdGet`

NewResponseDetailsApiV2ObservablesObservableIdGetWithDefaults instantiates a new ResponseDetailsApiV2ObservablesObservableIdGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) SetValue(v string)`

SetValue sets Value field to given value.


### GetType

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreated

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetContext

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetContext() []map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetContextOk() (*[]map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) SetContext(v []map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetLastAnalysis

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetLastAnalysis() map[string]time.Time`

GetLastAnalysis returns the LastAnalysis field if non-nil, zero value otherwise.

### GetLastAnalysisOk

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetLastAnalysisOk() (*map[string]time.Time, bool)`

GetLastAnalysisOk returns a tuple with the LastAnalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAnalysis

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) SetLastAnalysis(v map[string]time.Time)`

SetLastAnalysis sets LastAnalysis field to given value.

### HasLastAnalysis

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) HasLastAnalysis() bool`

HasLastAnalysis returns a boolean if a field has been set.

### GetId

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) SetId(v string)`

SetId sets Id field to given value.


### GetTags

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetTags() map[string]TagRelationshipOutput`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetTagsOk() (*map[string]TagRelationshipOutput, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) SetTags(v map[string]TagRelationshipOutput)`

SetTags sets Tags field to given value.


### GetRootType

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetRootType() string`

GetRootType returns the RootType field if non-nil, zero value otherwise.

### GetRootTypeOk

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetRootTypeOk() (*string, bool)`

GetRootTypeOk returns a tuple with the RootType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootType

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) SetRootType(v string)`

SetRootType sets RootType field to given value.


### GetLastSeen

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetLastSeen() time.Time`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetLastSeenOk() (*time.Time, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) SetLastSeen(v time.Time)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetFirstSeen

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetFirstSeen() time.Time`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetFirstSeenOk() (*time.Time, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) SetFirstSeen(v time.Time)`

SetFirstSeen sets FirstSeen field to given value.

### HasFirstSeen

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### GetIssuer

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetSubject

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetSerialNumber

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetAfter

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetAfter() time.Time`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetAfterOk() (*time.Time, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) SetAfter(v time.Time)`

SetAfter sets After field to given value.

### HasAfter

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) HasAfter() bool`

HasAfter returns a boolean if a field has been set.

### GetBefore

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetBefore() time.Time`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetBeforeOk() (*time.Time, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) SetBefore(v time.Time)`

SetBefore sets Before field to given value.

### HasBefore

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) HasBefore() bool`

HasBefore returns a boolean if a field has been set.

### GetFingerprint

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetKey

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) SetKey(v string)`

SetKey sets Key field to given value.


### GetData

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetData() *os.File`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetDataOk() (**os.File, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) SetData(v *os.File)`

SetData sets Data field to given value.


### GetHive

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetHive() RegistryHive`

GetHive returns the Hive field if non-nil, zero value otherwise.

### GetHiveOk

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetHiveOk() (*RegistryHive, bool)`

GetHiveOk returns a tuple with the Hive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHive

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) SetHive(v RegistryHive)`

SetHive sets Hive field to given value.


### GetPathFile

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetPathFile() string`

GetPathFile returns the PathFile field if non-nil, zero value otherwise.

### GetPathFileOk

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetPathFileOk() (*string, bool)`

GetPathFileOk returns a tuple with the PathFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathFile

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) SetPathFile(v string)`

SetPathFile sets PathFile field to given value.

### HasPathFile

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) HasPathFile() bool`

HasPathFile returns a boolean if a field has been set.

### GetCountry

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetDescription

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVersion

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetRegitryType

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetRegitryType() string`

GetRegitryType returns the RegitryType field if non-nil, zero value otherwise.

### GetRegitryTypeOk

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetRegitryTypeOk() (*string, bool)`

GetRegitryTypeOk returns a tuple with the RegitryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegitryType

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) SetRegitryType(v string)`

SetRegitryType sets RegitryType field to given value.

### HasRegitryType

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) HasRegitryType() bool`

HasRegitryType returns a boolean if a field has been set.

### GetName

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSize

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSha256

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### GetMd5

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) SetMd5(v string)`

SetMd5 sets Md5 field to given value.

### HasMd5

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) HasMd5() bool`

HasMd5 returns a boolean if a field has been set.

### GetSha1

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetSha1() string`

GetSha1 returns the Sha1 field if non-nil, zero value otherwise.

### GetSha1Ok

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetSha1Ok() (*string, bool)`

GetSha1Ok returns a tuple with the Sha1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha1

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) SetSha1(v string)`

SetSha1 sets Sha1 field to given value.

### HasSha1

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) HasSha1() bool`

HasSha1 returns a boolean if a field has been set.

### GetMimeType

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetUserId

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetCredential

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetCredential() string`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetCredentialOk() (*string, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) SetCredential(v string)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetAccountLogin

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetAccountLogin() string`

GetAccountLogin returns the AccountLogin field if non-nil, zero value otherwise.

### GetAccountLoginOk

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetAccountLoginOk() (*string, bool)`

GetAccountLoginOk returns a tuple with the AccountLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLogin

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) SetAccountLogin(v string)`

SetAccountLogin sets AccountLogin field to given value.

### HasAccountLogin

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) HasAccountLogin() bool`

HasAccountLogin returns a boolean if a field has been set.

### GetAccountType

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetDisplayName

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetIsServiceAccount

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetIsServiceAccount() bool`

GetIsServiceAccount returns the IsServiceAccount field if non-nil, zero value otherwise.

### GetIsServiceAccountOk

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetIsServiceAccountOk() (*bool, bool)`

GetIsServiceAccountOk returns a tuple with the IsServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsServiceAccount

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) SetIsServiceAccount(v bool)`

SetIsServiceAccount sets IsServiceAccount field to given value.

### HasIsServiceAccount

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) HasIsServiceAccount() bool`

HasIsServiceAccount returns a boolean if a field has been set.

### GetIsPrivileged

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetIsPrivileged() bool`

GetIsPrivileged returns the IsPrivileged field if non-nil, zero value otherwise.

### GetIsPrivilegedOk

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetIsPrivilegedOk() (*bool, bool)`

GetIsPrivilegedOk returns a tuple with the IsPrivileged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivileged

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) SetIsPrivileged(v bool)`

SetIsPrivileged sets IsPrivileged field to given value.

### HasIsPrivileged

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) HasIsPrivileged() bool`

HasIsPrivileged returns a boolean if a field has been set.

### GetCanEscalatePrivs

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetCanEscalatePrivs() bool`

GetCanEscalatePrivs returns the CanEscalatePrivs field if non-nil, zero value otherwise.

### GetCanEscalatePrivsOk

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetCanEscalatePrivsOk() (*bool, bool)`

GetCanEscalatePrivsOk returns a tuple with the CanEscalatePrivs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEscalatePrivs

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) SetCanEscalatePrivs(v bool)`

SetCanEscalatePrivs sets CanEscalatePrivs field to given value.

### HasCanEscalatePrivs

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) HasCanEscalatePrivs() bool`

HasCanEscalatePrivs returns a boolean if a field has been set.

### GetIsDisabled

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetAccountCreated

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetAccountCreated() time.Time`

GetAccountCreated returns the AccountCreated field if non-nil, zero value otherwise.

### GetAccountCreatedOk

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetAccountCreatedOk() (*time.Time, bool)`

GetAccountCreatedOk returns a tuple with the AccountCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCreated

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) SetAccountCreated(v time.Time)`

SetAccountCreated sets AccountCreated field to given value.

### HasAccountCreated

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) HasAccountCreated() bool`

HasAccountCreated returns a boolean if a field has been set.

### GetAccountExpires

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetAccountExpires() time.Time`

GetAccountExpires returns the AccountExpires field if non-nil, zero value otherwise.

### GetAccountExpiresOk

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetAccountExpiresOk() (*time.Time, bool)`

GetAccountExpiresOk returns a tuple with the AccountExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountExpires

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) SetAccountExpires(v time.Time)`

SetAccountExpires sets AccountExpires field to given value.

### HasAccountExpires

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) HasAccountExpires() bool`

HasAccountExpires returns a boolean if a field has been set.

### GetCredentialLastChanged

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetCredentialLastChanged() time.Time`

GetCredentialLastChanged returns the CredentialLastChanged field if non-nil, zero value otherwise.

### GetCredentialLastChangedOk

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetCredentialLastChangedOk() (*time.Time, bool)`

GetCredentialLastChangedOk returns a tuple with the CredentialLastChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialLastChanged

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) SetCredentialLastChanged(v time.Time)`

SetCredentialLastChanged sets CredentialLastChanged field to given value.

### HasCredentialLastChanged

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) HasCredentialLastChanged() bool`

HasCredentialLastChanged returns a boolean if a field has been set.

### GetAccountFirstLogin

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetAccountFirstLogin() time.Time`

GetAccountFirstLogin returns the AccountFirstLogin field if non-nil, zero value otherwise.

### GetAccountFirstLoginOk

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetAccountFirstLoginOk() (*time.Time, bool)`

GetAccountFirstLoginOk returns a tuple with the AccountFirstLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountFirstLogin

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) SetAccountFirstLogin(v time.Time)`

SetAccountFirstLogin sets AccountFirstLogin field to given value.

### HasAccountFirstLogin

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) HasAccountFirstLogin() bool`

HasAccountFirstLogin returns a boolean if a field has been set.

### GetAccountLastLogin

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetAccountLastLogin() time.Time`

GetAccountLastLogin returns the AccountLastLogin field if non-nil, zero value otherwise.

### GetAccountLastLoginOk

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetAccountLastLoginOk() (*time.Time, bool)`

GetAccountLastLoginOk returns a tuple with the AccountLastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLastLogin

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) SetAccountLastLogin(v time.Time)`

SetAccountLastLogin sets AccountLastLogin field to given value.

### HasAccountLastLogin

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) HasAccountLastLogin() bool`

HasAccountLastLogin returns a boolean if a field has been set.

### GetCoin

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetCoin() string`

GetCoin returns the Coin field if non-nil, zero value otherwise.

### GetCoinOk

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetCoinOk() (*string, bool)`

GetCoinOk returns a tuple with the Coin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoin

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) SetCoin(v string)`

SetCoin sets Coin field to given value.

### HasCoin

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) HasCoin() bool`

HasCoin returns a boolean if a field has been set.

### GetAddress

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ResponseDetailsApiV2ObservablesObservableIdGet) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


