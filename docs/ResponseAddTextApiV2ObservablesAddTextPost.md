# ResponseAddTextApiV2ObservablesAddTextPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] [default to "imphash"]
**Created** | Pointer to **time.Time** |  | [optional] 
**Context** | Pointer to **[]map[string]interface{}** |  | [optional] [default to []]
**LastAnalysis** | Pointer to [**map[string]time.Time**](time.Time.md) |  | [optional] [default to {}]
**Id** | **string** |  | [readonly] 
**Tags** | [**map[string]TagRelationshipOutput**](TagRelationshipOutput.md) |  | [readonly] 
**RootType** | **string** |  | [readonly] 
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
**Country** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
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
**Name** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Sha256** | Pointer to **string** |  | [optional] 
**Md5** | Pointer to **string** |  | [optional] 
**Sha1** | Pointer to **string** |  | [optional] 
**MimeType** | Pointer to **string** |  | [optional] 
**Coin** | Pointer to **string** |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 

## Methods

### NewResponseAddTextApiV2ObservablesAddTextPost

`func NewResponseAddTextApiV2ObservablesAddTextPost(value string, id string, tags map[string]TagRelationshipOutput, rootType string, key string, data *os.File, hive RegistryHive, ) *ResponseAddTextApiV2ObservablesAddTextPost`

NewResponseAddTextApiV2ObservablesAddTextPost instantiates a new ResponseAddTextApiV2ObservablesAddTextPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseAddTextApiV2ObservablesAddTextPostWithDefaults

`func NewResponseAddTextApiV2ObservablesAddTextPostWithDefaults() *ResponseAddTextApiV2ObservablesAddTextPost`

NewResponseAddTextApiV2ObservablesAddTextPostWithDefaults instantiates a new ResponseAddTextApiV2ObservablesAddTextPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) SetValue(v string)`

SetValue sets Value field to given value.


### GetType

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreated

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetContext

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetContext() []map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetContextOk() (*[]map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) SetContext(v []map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetLastAnalysis

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetLastAnalysis() map[string]time.Time`

GetLastAnalysis returns the LastAnalysis field if non-nil, zero value otherwise.

### GetLastAnalysisOk

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetLastAnalysisOk() (*map[string]time.Time, bool)`

GetLastAnalysisOk returns a tuple with the LastAnalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAnalysis

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) SetLastAnalysis(v map[string]time.Time)`

SetLastAnalysis sets LastAnalysis field to given value.

### HasLastAnalysis

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) HasLastAnalysis() bool`

HasLastAnalysis returns a boolean if a field has been set.

### GetId

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) SetId(v string)`

SetId sets Id field to given value.


### GetTags

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetTags() map[string]TagRelationshipOutput`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetTagsOk() (*map[string]TagRelationshipOutput, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) SetTags(v map[string]TagRelationshipOutput)`

SetTags sets Tags field to given value.


### GetRootType

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetRootType() string`

GetRootType returns the RootType field if non-nil, zero value otherwise.

### GetRootTypeOk

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetRootTypeOk() (*string, bool)`

GetRootTypeOk returns a tuple with the RootType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootType

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) SetRootType(v string)`

SetRootType sets RootType field to given value.


### GetVersion

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetRegitryType

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetRegitryType() string`

GetRegitryType returns the RegitryType field if non-nil, zero value otherwise.

### GetRegitryTypeOk

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetRegitryTypeOk() (*string, bool)`

GetRegitryTypeOk returns a tuple with the RegitryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegitryType

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) SetRegitryType(v string)`

SetRegitryType sets RegitryType field to given value.

### HasRegitryType

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) HasRegitryType() bool`

HasRegitryType returns a boolean if a field has been set.

### GetUserId

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetCredential

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetCredential() string`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetCredentialOk() (*string, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) SetCredential(v string)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetAccountLogin

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetAccountLogin() string`

GetAccountLogin returns the AccountLogin field if non-nil, zero value otherwise.

### GetAccountLoginOk

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetAccountLoginOk() (*string, bool)`

GetAccountLoginOk returns a tuple with the AccountLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLogin

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) SetAccountLogin(v string)`

SetAccountLogin sets AccountLogin field to given value.

### HasAccountLogin

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) HasAccountLogin() bool`

HasAccountLogin returns a boolean if a field has been set.

### GetAccountType

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetDisplayName

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetIsServiceAccount

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetIsServiceAccount() bool`

GetIsServiceAccount returns the IsServiceAccount field if non-nil, zero value otherwise.

### GetIsServiceAccountOk

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetIsServiceAccountOk() (*bool, bool)`

GetIsServiceAccountOk returns a tuple with the IsServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsServiceAccount

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) SetIsServiceAccount(v bool)`

SetIsServiceAccount sets IsServiceAccount field to given value.

### HasIsServiceAccount

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) HasIsServiceAccount() bool`

HasIsServiceAccount returns a boolean if a field has been set.

### GetIsPrivileged

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetIsPrivileged() bool`

GetIsPrivileged returns the IsPrivileged field if non-nil, zero value otherwise.

### GetIsPrivilegedOk

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetIsPrivilegedOk() (*bool, bool)`

GetIsPrivilegedOk returns a tuple with the IsPrivileged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivileged

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) SetIsPrivileged(v bool)`

SetIsPrivileged sets IsPrivileged field to given value.

### HasIsPrivileged

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) HasIsPrivileged() bool`

HasIsPrivileged returns a boolean if a field has been set.

### GetCanEscalatePrivs

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetCanEscalatePrivs() bool`

GetCanEscalatePrivs returns the CanEscalatePrivs field if non-nil, zero value otherwise.

### GetCanEscalatePrivsOk

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetCanEscalatePrivsOk() (*bool, bool)`

GetCanEscalatePrivsOk returns a tuple with the CanEscalatePrivs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEscalatePrivs

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) SetCanEscalatePrivs(v bool)`

SetCanEscalatePrivs sets CanEscalatePrivs field to given value.

### HasCanEscalatePrivs

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) HasCanEscalatePrivs() bool`

HasCanEscalatePrivs returns a boolean if a field has been set.

### GetIsDisabled

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetAccountCreated

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetAccountCreated() time.Time`

GetAccountCreated returns the AccountCreated field if non-nil, zero value otherwise.

### GetAccountCreatedOk

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetAccountCreatedOk() (*time.Time, bool)`

GetAccountCreatedOk returns a tuple with the AccountCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCreated

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) SetAccountCreated(v time.Time)`

SetAccountCreated sets AccountCreated field to given value.

### HasAccountCreated

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) HasAccountCreated() bool`

HasAccountCreated returns a boolean if a field has been set.

### GetAccountExpires

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetAccountExpires() time.Time`

GetAccountExpires returns the AccountExpires field if non-nil, zero value otherwise.

### GetAccountExpiresOk

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetAccountExpiresOk() (*time.Time, bool)`

GetAccountExpiresOk returns a tuple with the AccountExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountExpires

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) SetAccountExpires(v time.Time)`

SetAccountExpires sets AccountExpires field to given value.

### HasAccountExpires

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) HasAccountExpires() bool`

HasAccountExpires returns a boolean if a field has been set.

### GetCredentialLastChanged

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetCredentialLastChanged() time.Time`

GetCredentialLastChanged returns the CredentialLastChanged field if non-nil, zero value otherwise.

### GetCredentialLastChangedOk

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetCredentialLastChangedOk() (*time.Time, bool)`

GetCredentialLastChangedOk returns a tuple with the CredentialLastChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialLastChanged

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) SetCredentialLastChanged(v time.Time)`

SetCredentialLastChanged sets CredentialLastChanged field to given value.

### HasCredentialLastChanged

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) HasCredentialLastChanged() bool`

HasCredentialLastChanged returns a boolean if a field has been set.

### GetAccountFirstLogin

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetAccountFirstLogin() time.Time`

GetAccountFirstLogin returns the AccountFirstLogin field if non-nil, zero value otherwise.

### GetAccountFirstLoginOk

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetAccountFirstLoginOk() (*time.Time, bool)`

GetAccountFirstLoginOk returns a tuple with the AccountFirstLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountFirstLogin

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) SetAccountFirstLogin(v time.Time)`

SetAccountFirstLogin sets AccountFirstLogin field to given value.

### HasAccountFirstLogin

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) HasAccountFirstLogin() bool`

HasAccountFirstLogin returns a boolean if a field has been set.

### GetAccountLastLogin

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetAccountLastLogin() time.Time`

GetAccountLastLogin returns the AccountLastLogin field if non-nil, zero value otherwise.

### GetAccountLastLoginOk

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetAccountLastLoginOk() (*time.Time, bool)`

GetAccountLastLoginOk returns a tuple with the AccountLastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLastLogin

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) SetAccountLastLogin(v time.Time)`

SetAccountLastLogin sets AccountLastLogin field to given value.

### HasAccountLastLogin

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) HasAccountLastLogin() bool`

HasAccountLastLogin returns a boolean if a field has been set.

### GetCountry

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetDescription

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLastSeen

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetLastSeen() time.Time`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetLastSeenOk() (*time.Time, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) SetLastSeen(v time.Time)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetFirstSeen

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetFirstSeen() time.Time`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetFirstSeenOk() (*time.Time, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) SetFirstSeen(v time.Time)`

SetFirstSeen sets FirstSeen field to given value.

### HasFirstSeen

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### GetIssuer

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetSubject

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetSerialNumber

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetAfter

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetAfter() time.Time`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetAfterOk() (*time.Time, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) SetAfter(v time.Time)`

SetAfter sets After field to given value.

### HasAfter

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) HasAfter() bool`

HasAfter returns a boolean if a field has been set.

### GetBefore

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetBefore() time.Time`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetBeforeOk() (*time.Time, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) SetBefore(v time.Time)`

SetBefore sets Before field to given value.

### HasBefore

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) HasBefore() bool`

HasBefore returns a boolean if a field has been set.

### GetFingerprint

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetKey

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) SetKey(v string)`

SetKey sets Key field to given value.


### GetData

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetData() *os.File`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetDataOk() (**os.File, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) SetData(v *os.File)`

SetData sets Data field to given value.


### GetHive

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetHive() RegistryHive`

GetHive returns the Hive field if non-nil, zero value otherwise.

### GetHiveOk

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetHiveOk() (*RegistryHive, bool)`

GetHiveOk returns a tuple with the Hive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHive

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) SetHive(v RegistryHive)`

SetHive sets Hive field to given value.


### GetPathFile

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetPathFile() string`

GetPathFile returns the PathFile field if non-nil, zero value otherwise.

### GetPathFileOk

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetPathFileOk() (*string, bool)`

GetPathFileOk returns a tuple with the PathFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathFile

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) SetPathFile(v string)`

SetPathFile sets PathFile field to given value.

### HasPathFile

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) HasPathFile() bool`

HasPathFile returns a boolean if a field has been set.

### GetName

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSize

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSha256

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### GetMd5

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) SetMd5(v string)`

SetMd5 sets Md5 field to given value.

### HasMd5

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) HasMd5() bool`

HasMd5 returns a boolean if a field has been set.

### GetSha1

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetSha1() string`

GetSha1 returns the Sha1 field if non-nil, zero value otherwise.

### GetSha1Ok

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetSha1Ok() (*string, bool)`

GetSha1Ok returns a tuple with the Sha1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha1

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) SetSha1(v string)`

SetSha1 sets Sha1 field to given value.

### HasSha1

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) HasSha1() bool`

HasSha1 returns a boolean if a field has been set.

### GetMimeType

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetCoin

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetCoin() string`

GetCoin returns the Coin field if non-nil, zero value otherwise.

### GetCoinOk

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetCoinOk() (*string, bool)`

GetCoinOk returns a tuple with the Coin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoin

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) SetCoin(v string)`

SetCoin sets Coin field to given value.

### HasCoin

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) HasCoin() bool`

HasCoin returns a boolean if a field has been set.

### GetAddress

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ResponseAddTextApiV2ObservablesAddTextPost) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


