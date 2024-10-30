# ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] [default to "ipv6"]
**Created** | Pointer to **time.Time** |  | [optional] 
**Context** | Pointer to **[]map[string]interface{}** |  | [optional] [default to []]
**LastAnalysis** | Pointer to [**map[string]time.Time**](time.Time.md) |  | [optional] [default to {}]
**Id** | **string** |  | [readonly] 
**Tags** | [**map[string]TagRelationshipOutput**](TagRelationshipOutput.md) |  | [readonly] 
**RootType** | **string** |  | [readonly] 
**Key** | **string** |  | 
**Data** | ***os.File** |  | 
**Hive** | [**RegistryHive**](RegistryHive.md) |  | 
**PathFile** | Pointer to **string** |  | [optional] 
**LastSeen** | Pointer to **time.Time** |  | [optional] 
**FirstSeen** | Pointer to **time.Time** |  | [optional] 
**Issuer** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**After** | Pointer to **time.Time** |  | [optional] 
**Before** | Pointer to **time.Time** |  | [optional] 
**Fingerprint** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**RegitryType** | Pointer to **string** |  | [optional] 
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
**Country** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Sha256** | Pointer to **string** |  | [optional] 
**Md5** | Pointer to **string** |  | [optional] 
**Sha1** | Pointer to **string** |  | [optional] 
**MimeType** | Pointer to **string** |  | [optional] 

## Methods

### NewResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost

`func NewResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost(value string, id string, tags map[string]TagRelationshipOutput, rootType string, key string, data *os.File, hive RegistryHive, ) *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost`

NewResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost instantiates a new ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseDeleteContextApiV2ObservablesObservableIdContextDeletePostWithDefaults

`func NewResponseDeleteContextApiV2ObservablesObservableIdContextDeletePostWithDefaults() *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost`

NewResponseDeleteContextApiV2ObservablesObservableIdContextDeletePostWithDefaults instantiates a new ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) SetValue(v string)`

SetValue sets Value field to given value.


### GetType

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreated

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetContext

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetContext() []map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetContextOk() (*[]map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) SetContext(v []map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetLastAnalysis

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetLastAnalysis() map[string]time.Time`

GetLastAnalysis returns the LastAnalysis field if non-nil, zero value otherwise.

### GetLastAnalysisOk

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetLastAnalysisOk() (*map[string]time.Time, bool)`

GetLastAnalysisOk returns a tuple with the LastAnalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAnalysis

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) SetLastAnalysis(v map[string]time.Time)`

SetLastAnalysis sets LastAnalysis field to given value.

### HasLastAnalysis

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) HasLastAnalysis() bool`

HasLastAnalysis returns a boolean if a field has been set.

### GetId

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) SetId(v string)`

SetId sets Id field to given value.


### GetTags

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetTags() map[string]TagRelationshipOutput`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetTagsOk() (*map[string]TagRelationshipOutput, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) SetTags(v map[string]TagRelationshipOutput)`

SetTags sets Tags field to given value.


### GetRootType

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetRootType() string`

GetRootType returns the RootType field if non-nil, zero value otherwise.

### GetRootTypeOk

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetRootTypeOk() (*string, bool)`

GetRootTypeOk returns a tuple with the RootType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootType

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) SetRootType(v string)`

SetRootType sets RootType field to given value.


### GetKey

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) SetKey(v string)`

SetKey sets Key field to given value.


### GetData

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetData() *os.File`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetDataOk() (**os.File, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) SetData(v *os.File)`

SetData sets Data field to given value.


### GetHive

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetHive() RegistryHive`

GetHive returns the Hive field if non-nil, zero value otherwise.

### GetHiveOk

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetHiveOk() (*RegistryHive, bool)`

GetHiveOk returns a tuple with the Hive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHive

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) SetHive(v RegistryHive)`

SetHive sets Hive field to given value.


### GetPathFile

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetPathFile() string`

GetPathFile returns the PathFile field if non-nil, zero value otherwise.

### GetPathFileOk

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetPathFileOk() (*string, bool)`

GetPathFileOk returns a tuple with the PathFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathFile

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) SetPathFile(v string)`

SetPathFile sets PathFile field to given value.

### HasPathFile

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) HasPathFile() bool`

HasPathFile returns a boolean if a field has been set.

### GetLastSeen

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetLastSeen() time.Time`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetLastSeenOk() (*time.Time, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) SetLastSeen(v time.Time)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetFirstSeen

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetFirstSeen() time.Time`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetFirstSeenOk() (*time.Time, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) SetFirstSeen(v time.Time)`

SetFirstSeen sets FirstSeen field to given value.

### HasFirstSeen

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### GetIssuer

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetSubject

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetSerialNumber

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetAfter

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetAfter() time.Time`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetAfterOk() (*time.Time, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) SetAfter(v time.Time)`

SetAfter sets After field to given value.

### HasAfter

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) HasAfter() bool`

HasAfter returns a boolean if a field has been set.

### GetBefore

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetBefore() time.Time`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetBeforeOk() (*time.Time, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) SetBefore(v time.Time)`

SetBefore sets Before field to given value.

### HasBefore

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) HasBefore() bool`

HasBefore returns a boolean if a field has been set.

### GetFingerprint

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetVersion

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetRegitryType

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetRegitryType() string`

GetRegitryType returns the RegitryType field if non-nil, zero value otherwise.

### GetRegitryTypeOk

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetRegitryTypeOk() (*string, bool)`

GetRegitryTypeOk returns a tuple with the RegitryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegitryType

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) SetRegitryType(v string)`

SetRegitryType sets RegitryType field to given value.

### HasRegitryType

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) HasRegitryType() bool`

HasRegitryType returns a boolean if a field has been set.

### GetUserId

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetCredential

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetCredential() string`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetCredentialOk() (*string, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) SetCredential(v string)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetAccountLogin

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetAccountLogin() string`

GetAccountLogin returns the AccountLogin field if non-nil, zero value otherwise.

### GetAccountLoginOk

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetAccountLoginOk() (*string, bool)`

GetAccountLoginOk returns a tuple with the AccountLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLogin

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) SetAccountLogin(v string)`

SetAccountLogin sets AccountLogin field to given value.

### HasAccountLogin

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) HasAccountLogin() bool`

HasAccountLogin returns a boolean if a field has been set.

### GetAccountType

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetDisplayName

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetIsServiceAccount

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetIsServiceAccount() bool`

GetIsServiceAccount returns the IsServiceAccount field if non-nil, zero value otherwise.

### GetIsServiceAccountOk

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetIsServiceAccountOk() (*bool, bool)`

GetIsServiceAccountOk returns a tuple with the IsServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsServiceAccount

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) SetIsServiceAccount(v bool)`

SetIsServiceAccount sets IsServiceAccount field to given value.

### HasIsServiceAccount

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) HasIsServiceAccount() bool`

HasIsServiceAccount returns a boolean if a field has been set.

### GetIsPrivileged

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetIsPrivileged() bool`

GetIsPrivileged returns the IsPrivileged field if non-nil, zero value otherwise.

### GetIsPrivilegedOk

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetIsPrivilegedOk() (*bool, bool)`

GetIsPrivilegedOk returns a tuple with the IsPrivileged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivileged

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) SetIsPrivileged(v bool)`

SetIsPrivileged sets IsPrivileged field to given value.

### HasIsPrivileged

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) HasIsPrivileged() bool`

HasIsPrivileged returns a boolean if a field has been set.

### GetCanEscalatePrivs

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetCanEscalatePrivs() bool`

GetCanEscalatePrivs returns the CanEscalatePrivs field if non-nil, zero value otherwise.

### GetCanEscalatePrivsOk

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetCanEscalatePrivsOk() (*bool, bool)`

GetCanEscalatePrivsOk returns a tuple with the CanEscalatePrivs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEscalatePrivs

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) SetCanEscalatePrivs(v bool)`

SetCanEscalatePrivs sets CanEscalatePrivs field to given value.

### HasCanEscalatePrivs

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) HasCanEscalatePrivs() bool`

HasCanEscalatePrivs returns a boolean if a field has been set.

### GetIsDisabled

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetAccountCreated

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetAccountCreated() time.Time`

GetAccountCreated returns the AccountCreated field if non-nil, zero value otherwise.

### GetAccountCreatedOk

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetAccountCreatedOk() (*time.Time, bool)`

GetAccountCreatedOk returns a tuple with the AccountCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCreated

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) SetAccountCreated(v time.Time)`

SetAccountCreated sets AccountCreated field to given value.

### HasAccountCreated

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) HasAccountCreated() bool`

HasAccountCreated returns a boolean if a field has been set.

### GetAccountExpires

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetAccountExpires() time.Time`

GetAccountExpires returns the AccountExpires field if non-nil, zero value otherwise.

### GetAccountExpiresOk

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetAccountExpiresOk() (*time.Time, bool)`

GetAccountExpiresOk returns a tuple with the AccountExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountExpires

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) SetAccountExpires(v time.Time)`

SetAccountExpires sets AccountExpires field to given value.

### HasAccountExpires

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) HasAccountExpires() bool`

HasAccountExpires returns a boolean if a field has been set.

### GetCredentialLastChanged

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetCredentialLastChanged() time.Time`

GetCredentialLastChanged returns the CredentialLastChanged field if non-nil, zero value otherwise.

### GetCredentialLastChangedOk

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetCredentialLastChangedOk() (*time.Time, bool)`

GetCredentialLastChangedOk returns a tuple with the CredentialLastChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialLastChanged

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) SetCredentialLastChanged(v time.Time)`

SetCredentialLastChanged sets CredentialLastChanged field to given value.

### HasCredentialLastChanged

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) HasCredentialLastChanged() bool`

HasCredentialLastChanged returns a boolean if a field has been set.

### GetAccountFirstLogin

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetAccountFirstLogin() time.Time`

GetAccountFirstLogin returns the AccountFirstLogin field if non-nil, zero value otherwise.

### GetAccountFirstLoginOk

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetAccountFirstLoginOk() (*time.Time, bool)`

GetAccountFirstLoginOk returns a tuple with the AccountFirstLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountFirstLogin

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) SetAccountFirstLogin(v time.Time)`

SetAccountFirstLogin sets AccountFirstLogin field to given value.

### HasAccountFirstLogin

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) HasAccountFirstLogin() bool`

HasAccountFirstLogin returns a boolean if a field has been set.

### GetAccountLastLogin

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetAccountLastLogin() time.Time`

GetAccountLastLogin returns the AccountLastLogin field if non-nil, zero value otherwise.

### GetAccountLastLoginOk

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetAccountLastLoginOk() (*time.Time, bool)`

GetAccountLastLoginOk returns a tuple with the AccountLastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLastLogin

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) SetAccountLastLogin(v time.Time)`

SetAccountLastLogin sets AccountLastLogin field to given value.

### HasAccountLastLogin

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) HasAccountLastLogin() bool`

HasAccountLastLogin returns a boolean if a field has been set.

### GetCoin

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetCoin() string`

GetCoin returns the Coin field if non-nil, zero value otherwise.

### GetCoinOk

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetCoinOk() (*string, bool)`

GetCoinOk returns a tuple with the Coin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoin

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) SetCoin(v string)`

SetCoin sets Coin field to given value.

### HasCoin

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) HasCoin() bool`

HasCoin returns a boolean if a field has been set.

### GetAddress

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCountry

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetDescription

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSize

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSha256

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### GetMd5

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) SetMd5(v string)`

SetMd5 sets Md5 field to given value.

### HasMd5

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) HasMd5() bool`

HasMd5 returns a boolean if a field has been set.

### GetSha1

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetSha1() string`

GetSha1 returns the Sha1 field if non-nil, zero value otherwise.

### GetSha1Ok

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetSha1Ok() (*string, bool)`

GetSha1Ok returns a tuple with the Sha1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha1

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) SetSha1(v string)`

SetSha1 sets Sha1 field to given value.

### HasSha1

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) HasSha1() bool`

HasSha1 returns a boolean if a field has been set.

### GetMimeType

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


