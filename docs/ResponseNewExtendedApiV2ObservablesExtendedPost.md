# ResponseNewExtendedApiV2ObservablesExtendedPost

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

### NewResponseNewExtendedApiV2ObservablesExtendedPost

`func NewResponseNewExtendedApiV2ObservablesExtendedPost(value string, id string, tags map[string]TagRelationshipOutput, rootType string, key string, data *os.File, hive RegistryHive, ) *ResponseNewExtendedApiV2ObservablesExtendedPost`

NewResponseNewExtendedApiV2ObservablesExtendedPost instantiates a new ResponseNewExtendedApiV2ObservablesExtendedPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseNewExtendedApiV2ObservablesExtendedPostWithDefaults

`func NewResponseNewExtendedApiV2ObservablesExtendedPostWithDefaults() *ResponseNewExtendedApiV2ObservablesExtendedPost`

NewResponseNewExtendedApiV2ObservablesExtendedPostWithDefaults instantiates a new ResponseNewExtendedApiV2ObservablesExtendedPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) SetValue(v string)`

SetValue sets Value field to given value.


### GetType

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreated

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetContext

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetContext() []map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetContextOk() (*[]map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) SetContext(v []map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetLastAnalysis

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetLastAnalysis() map[string]time.Time`

GetLastAnalysis returns the LastAnalysis field if non-nil, zero value otherwise.

### GetLastAnalysisOk

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetLastAnalysisOk() (*map[string]time.Time, bool)`

GetLastAnalysisOk returns a tuple with the LastAnalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAnalysis

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) SetLastAnalysis(v map[string]time.Time)`

SetLastAnalysis sets LastAnalysis field to given value.

### HasLastAnalysis

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) HasLastAnalysis() bool`

HasLastAnalysis returns a boolean if a field has been set.

### GetId

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) SetId(v string)`

SetId sets Id field to given value.


### GetTags

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetTags() map[string]TagRelationshipOutput`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetTagsOk() (*map[string]TagRelationshipOutput, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) SetTags(v map[string]TagRelationshipOutput)`

SetTags sets Tags field to given value.


### GetRootType

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetRootType() string`

GetRootType returns the RootType field if non-nil, zero value otherwise.

### GetRootTypeOk

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetRootTypeOk() (*string, bool)`

GetRootTypeOk returns a tuple with the RootType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootType

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) SetRootType(v string)`

SetRootType sets RootType field to given value.


### GetVersion

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetRegitryType

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetRegitryType() string`

GetRegitryType returns the RegitryType field if non-nil, zero value otherwise.

### GetRegitryTypeOk

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetRegitryTypeOk() (*string, bool)`

GetRegitryTypeOk returns a tuple with the RegitryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegitryType

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) SetRegitryType(v string)`

SetRegitryType sets RegitryType field to given value.

### HasRegitryType

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) HasRegitryType() bool`

HasRegitryType returns a boolean if a field has been set.

### GetUserId

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetCredential

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetCredential() string`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetCredentialOk() (*string, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) SetCredential(v string)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetAccountLogin

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetAccountLogin() string`

GetAccountLogin returns the AccountLogin field if non-nil, zero value otherwise.

### GetAccountLoginOk

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetAccountLoginOk() (*string, bool)`

GetAccountLoginOk returns a tuple with the AccountLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLogin

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) SetAccountLogin(v string)`

SetAccountLogin sets AccountLogin field to given value.

### HasAccountLogin

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) HasAccountLogin() bool`

HasAccountLogin returns a boolean if a field has been set.

### GetAccountType

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetDisplayName

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetIsServiceAccount

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetIsServiceAccount() bool`

GetIsServiceAccount returns the IsServiceAccount field if non-nil, zero value otherwise.

### GetIsServiceAccountOk

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetIsServiceAccountOk() (*bool, bool)`

GetIsServiceAccountOk returns a tuple with the IsServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsServiceAccount

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) SetIsServiceAccount(v bool)`

SetIsServiceAccount sets IsServiceAccount field to given value.

### HasIsServiceAccount

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) HasIsServiceAccount() bool`

HasIsServiceAccount returns a boolean if a field has been set.

### GetIsPrivileged

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetIsPrivileged() bool`

GetIsPrivileged returns the IsPrivileged field if non-nil, zero value otherwise.

### GetIsPrivilegedOk

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetIsPrivilegedOk() (*bool, bool)`

GetIsPrivilegedOk returns a tuple with the IsPrivileged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivileged

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) SetIsPrivileged(v bool)`

SetIsPrivileged sets IsPrivileged field to given value.

### HasIsPrivileged

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) HasIsPrivileged() bool`

HasIsPrivileged returns a boolean if a field has been set.

### GetCanEscalatePrivs

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetCanEscalatePrivs() bool`

GetCanEscalatePrivs returns the CanEscalatePrivs field if non-nil, zero value otherwise.

### GetCanEscalatePrivsOk

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetCanEscalatePrivsOk() (*bool, bool)`

GetCanEscalatePrivsOk returns a tuple with the CanEscalatePrivs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEscalatePrivs

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) SetCanEscalatePrivs(v bool)`

SetCanEscalatePrivs sets CanEscalatePrivs field to given value.

### HasCanEscalatePrivs

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) HasCanEscalatePrivs() bool`

HasCanEscalatePrivs returns a boolean if a field has been set.

### GetIsDisabled

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetAccountCreated

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetAccountCreated() time.Time`

GetAccountCreated returns the AccountCreated field if non-nil, zero value otherwise.

### GetAccountCreatedOk

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetAccountCreatedOk() (*time.Time, bool)`

GetAccountCreatedOk returns a tuple with the AccountCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCreated

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) SetAccountCreated(v time.Time)`

SetAccountCreated sets AccountCreated field to given value.

### HasAccountCreated

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) HasAccountCreated() bool`

HasAccountCreated returns a boolean if a field has been set.

### GetAccountExpires

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetAccountExpires() time.Time`

GetAccountExpires returns the AccountExpires field if non-nil, zero value otherwise.

### GetAccountExpiresOk

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetAccountExpiresOk() (*time.Time, bool)`

GetAccountExpiresOk returns a tuple with the AccountExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountExpires

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) SetAccountExpires(v time.Time)`

SetAccountExpires sets AccountExpires field to given value.

### HasAccountExpires

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) HasAccountExpires() bool`

HasAccountExpires returns a boolean if a field has been set.

### GetCredentialLastChanged

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetCredentialLastChanged() time.Time`

GetCredentialLastChanged returns the CredentialLastChanged field if non-nil, zero value otherwise.

### GetCredentialLastChangedOk

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetCredentialLastChangedOk() (*time.Time, bool)`

GetCredentialLastChangedOk returns a tuple with the CredentialLastChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialLastChanged

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) SetCredentialLastChanged(v time.Time)`

SetCredentialLastChanged sets CredentialLastChanged field to given value.

### HasCredentialLastChanged

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) HasCredentialLastChanged() bool`

HasCredentialLastChanged returns a boolean if a field has been set.

### GetAccountFirstLogin

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetAccountFirstLogin() time.Time`

GetAccountFirstLogin returns the AccountFirstLogin field if non-nil, zero value otherwise.

### GetAccountFirstLoginOk

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetAccountFirstLoginOk() (*time.Time, bool)`

GetAccountFirstLoginOk returns a tuple with the AccountFirstLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountFirstLogin

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) SetAccountFirstLogin(v time.Time)`

SetAccountFirstLogin sets AccountFirstLogin field to given value.

### HasAccountFirstLogin

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) HasAccountFirstLogin() bool`

HasAccountFirstLogin returns a boolean if a field has been set.

### GetAccountLastLogin

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetAccountLastLogin() time.Time`

GetAccountLastLogin returns the AccountLastLogin field if non-nil, zero value otherwise.

### GetAccountLastLoginOk

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetAccountLastLoginOk() (*time.Time, bool)`

GetAccountLastLoginOk returns a tuple with the AccountLastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLastLogin

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) SetAccountLastLogin(v time.Time)`

SetAccountLastLogin sets AccountLastLogin field to given value.

### HasAccountLastLogin

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) HasAccountLastLogin() bool`

HasAccountLastLogin returns a boolean if a field has been set.

### GetCountry

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetDescription

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLastSeen

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetLastSeen() time.Time`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetLastSeenOk() (*time.Time, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) SetLastSeen(v time.Time)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetFirstSeen

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetFirstSeen() time.Time`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetFirstSeenOk() (*time.Time, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) SetFirstSeen(v time.Time)`

SetFirstSeen sets FirstSeen field to given value.

### HasFirstSeen

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### GetIssuer

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetSubject

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetSerialNumber

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetAfter

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetAfter() time.Time`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetAfterOk() (*time.Time, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) SetAfter(v time.Time)`

SetAfter sets After field to given value.

### HasAfter

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) HasAfter() bool`

HasAfter returns a boolean if a field has been set.

### GetBefore

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetBefore() time.Time`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetBeforeOk() (*time.Time, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) SetBefore(v time.Time)`

SetBefore sets Before field to given value.

### HasBefore

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) HasBefore() bool`

HasBefore returns a boolean if a field has been set.

### GetFingerprint

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetKey

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) SetKey(v string)`

SetKey sets Key field to given value.


### GetData

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetData() *os.File`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetDataOk() (**os.File, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) SetData(v *os.File)`

SetData sets Data field to given value.


### GetHive

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetHive() RegistryHive`

GetHive returns the Hive field if non-nil, zero value otherwise.

### GetHiveOk

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetHiveOk() (*RegistryHive, bool)`

GetHiveOk returns a tuple with the Hive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHive

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) SetHive(v RegistryHive)`

SetHive sets Hive field to given value.


### GetPathFile

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetPathFile() string`

GetPathFile returns the PathFile field if non-nil, zero value otherwise.

### GetPathFileOk

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetPathFileOk() (*string, bool)`

GetPathFileOk returns a tuple with the PathFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathFile

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) SetPathFile(v string)`

SetPathFile sets PathFile field to given value.

### HasPathFile

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) HasPathFile() bool`

HasPathFile returns a boolean if a field has been set.

### GetName

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSize

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSha256

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### GetMd5

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) SetMd5(v string)`

SetMd5 sets Md5 field to given value.

### HasMd5

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) HasMd5() bool`

HasMd5 returns a boolean if a field has been set.

### GetSha1

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetSha1() string`

GetSha1 returns the Sha1 field if non-nil, zero value otherwise.

### GetSha1Ok

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetSha1Ok() (*string, bool)`

GetSha1Ok returns a tuple with the Sha1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha1

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) SetSha1(v string)`

SetSha1 sets Sha1 field to given value.

### HasSha1

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) HasSha1() bool`

HasSha1 returns a boolean if a field has been set.

### GetMimeType

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetCoin

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetCoin() string`

GetCoin returns the Coin field if non-nil, zero value otherwise.

### GetCoinOk

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetCoinOk() (*string, bool)`

GetCoinOk returns a tuple with the Coin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoin

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) SetCoin(v string)`

SetCoin sets Coin field to given value.

### HasCoin

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) HasCoin() bool`

HasCoin returns a boolean if a field has been set.

### GetAddress

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ResponseNewExtendedApiV2ObservablesExtendedPost) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


