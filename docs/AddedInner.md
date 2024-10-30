# AddedInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] [default to "mutex"]
**Created** | Pointer to **time.Time** |  | [optional] 
**Context** | Pointer to **[]map[string]interface{}** |  | [optional] [default to []]
**LastAnalysis** | Pointer to [**map[string]time.Time**](time.Time.md) |  | [optional] [default to {}]
**Id** | **string** |  | [readonly] 
**Tags** | [**map[string]TagRelationshipOutput**](TagRelationshipOutput.md) |  | [readonly] 
**RootType** | **string** |  | [readonly] 
**Version** | Pointer to **string** |  | [optional] 
**RegitryType** | Pointer to **string** |  | [optional] 
**Key** | **string** |  | 
**Data** | ***os.File** |  | 
**Hive** | [**RegistryHive**](RegistryHive.md) |  | 
**PathFile** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
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
**LastSeen** | Pointer to **time.Time** |  | [optional] 
**FirstSeen** | Pointer to **time.Time** |  | [optional] 
**Issuer** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**After** | Pointer to **time.Time** |  | [optional] 
**Before** | Pointer to **time.Time** |  | [optional] 
**Fingerprint** | Pointer to **string** |  | [optional] 

## Methods

### NewAddedInner

`func NewAddedInner(value string, id string, tags map[string]TagRelationshipOutput, rootType string, key string, data *os.File, hive RegistryHive, ) *AddedInner`

NewAddedInner instantiates a new AddedInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddedInnerWithDefaults

`func NewAddedInnerWithDefaults() *AddedInner`

NewAddedInnerWithDefaults instantiates a new AddedInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *AddedInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AddedInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AddedInner) SetValue(v string)`

SetValue sets Value field to given value.


### GetType

`func (o *AddedInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddedInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddedInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AddedInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreated

`func (o *AddedInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AddedInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AddedInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AddedInner) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetContext

`func (o *AddedInner) GetContext() []map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *AddedInner) GetContextOk() (*[]map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *AddedInner) SetContext(v []map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *AddedInner) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetLastAnalysis

`func (o *AddedInner) GetLastAnalysis() map[string]time.Time`

GetLastAnalysis returns the LastAnalysis field if non-nil, zero value otherwise.

### GetLastAnalysisOk

`func (o *AddedInner) GetLastAnalysisOk() (*map[string]time.Time, bool)`

GetLastAnalysisOk returns a tuple with the LastAnalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAnalysis

`func (o *AddedInner) SetLastAnalysis(v map[string]time.Time)`

SetLastAnalysis sets LastAnalysis field to given value.

### HasLastAnalysis

`func (o *AddedInner) HasLastAnalysis() bool`

HasLastAnalysis returns a boolean if a field has been set.

### GetId

`func (o *AddedInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddedInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddedInner) SetId(v string)`

SetId sets Id field to given value.


### GetTags

`func (o *AddedInner) GetTags() map[string]TagRelationshipOutput`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AddedInner) GetTagsOk() (*map[string]TagRelationshipOutput, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AddedInner) SetTags(v map[string]TagRelationshipOutput)`

SetTags sets Tags field to given value.


### GetRootType

`func (o *AddedInner) GetRootType() string`

GetRootType returns the RootType field if non-nil, zero value otherwise.

### GetRootTypeOk

`func (o *AddedInner) GetRootTypeOk() (*string, bool)`

GetRootTypeOk returns a tuple with the RootType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootType

`func (o *AddedInner) SetRootType(v string)`

SetRootType sets RootType field to given value.


### GetVersion

`func (o *AddedInner) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AddedInner) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AddedInner) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AddedInner) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetRegitryType

`func (o *AddedInner) GetRegitryType() string`

GetRegitryType returns the RegitryType field if non-nil, zero value otherwise.

### GetRegitryTypeOk

`func (o *AddedInner) GetRegitryTypeOk() (*string, bool)`

GetRegitryTypeOk returns a tuple with the RegitryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegitryType

`func (o *AddedInner) SetRegitryType(v string)`

SetRegitryType sets RegitryType field to given value.

### HasRegitryType

`func (o *AddedInner) HasRegitryType() bool`

HasRegitryType returns a boolean if a field has been set.

### GetKey

`func (o *AddedInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AddedInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AddedInner) SetKey(v string)`

SetKey sets Key field to given value.


### GetData

`func (o *AddedInner) GetData() *os.File`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AddedInner) GetDataOk() (**os.File, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AddedInner) SetData(v *os.File)`

SetData sets Data field to given value.


### GetHive

`func (o *AddedInner) GetHive() RegistryHive`

GetHive returns the Hive field if non-nil, zero value otherwise.

### GetHiveOk

`func (o *AddedInner) GetHiveOk() (*RegistryHive, bool)`

GetHiveOk returns a tuple with the Hive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHive

`func (o *AddedInner) SetHive(v RegistryHive)`

SetHive sets Hive field to given value.


### GetPathFile

`func (o *AddedInner) GetPathFile() string`

GetPathFile returns the PathFile field if non-nil, zero value otherwise.

### GetPathFileOk

`func (o *AddedInner) GetPathFileOk() (*string, bool)`

GetPathFileOk returns a tuple with the PathFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathFile

`func (o *AddedInner) SetPathFile(v string)`

SetPathFile sets PathFile field to given value.

### HasPathFile

`func (o *AddedInner) HasPathFile() bool`

HasPathFile returns a boolean if a field has been set.

### GetCountry

`func (o *AddedInner) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *AddedInner) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *AddedInner) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *AddedInner) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetDescription

`func (o *AddedInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddedInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddedInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddedInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *AddedInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddedInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddedInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddedInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSize

`func (o *AddedInner) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *AddedInner) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *AddedInner) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *AddedInner) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSha256

`func (o *AddedInner) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *AddedInner) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *AddedInner) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *AddedInner) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### GetMd5

`func (o *AddedInner) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *AddedInner) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *AddedInner) SetMd5(v string)`

SetMd5 sets Md5 field to given value.

### HasMd5

`func (o *AddedInner) HasMd5() bool`

HasMd5 returns a boolean if a field has been set.

### GetSha1

`func (o *AddedInner) GetSha1() string`

GetSha1 returns the Sha1 field if non-nil, zero value otherwise.

### GetSha1Ok

`func (o *AddedInner) GetSha1Ok() (*string, bool)`

GetSha1Ok returns a tuple with the Sha1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha1

`func (o *AddedInner) SetSha1(v string)`

SetSha1 sets Sha1 field to given value.

### HasSha1

`func (o *AddedInner) HasSha1() bool`

HasSha1 returns a boolean if a field has been set.

### GetMimeType

`func (o *AddedInner) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *AddedInner) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *AddedInner) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *AddedInner) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetUserId

`func (o *AddedInner) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AddedInner) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AddedInner) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AddedInner) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetCredential

`func (o *AddedInner) GetCredential() string`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *AddedInner) GetCredentialOk() (*string, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *AddedInner) SetCredential(v string)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *AddedInner) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetAccountLogin

`func (o *AddedInner) GetAccountLogin() string`

GetAccountLogin returns the AccountLogin field if non-nil, zero value otherwise.

### GetAccountLoginOk

`func (o *AddedInner) GetAccountLoginOk() (*string, bool)`

GetAccountLoginOk returns a tuple with the AccountLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLogin

`func (o *AddedInner) SetAccountLogin(v string)`

SetAccountLogin sets AccountLogin field to given value.

### HasAccountLogin

`func (o *AddedInner) HasAccountLogin() bool`

HasAccountLogin returns a boolean if a field has been set.

### GetAccountType

`func (o *AddedInner) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *AddedInner) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *AddedInner) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *AddedInner) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetDisplayName

`func (o *AddedInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AddedInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AddedInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AddedInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetIsServiceAccount

`func (o *AddedInner) GetIsServiceAccount() bool`

GetIsServiceAccount returns the IsServiceAccount field if non-nil, zero value otherwise.

### GetIsServiceAccountOk

`func (o *AddedInner) GetIsServiceAccountOk() (*bool, bool)`

GetIsServiceAccountOk returns a tuple with the IsServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsServiceAccount

`func (o *AddedInner) SetIsServiceAccount(v bool)`

SetIsServiceAccount sets IsServiceAccount field to given value.

### HasIsServiceAccount

`func (o *AddedInner) HasIsServiceAccount() bool`

HasIsServiceAccount returns a boolean if a field has been set.

### GetIsPrivileged

`func (o *AddedInner) GetIsPrivileged() bool`

GetIsPrivileged returns the IsPrivileged field if non-nil, zero value otherwise.

### GetIsPrivilegedOk

`func (o *AddedInner) GetIsPrivilegedOk() (*bool, bool)`

GetIsPrivilegedOk returns a tuple with the IsPrivileged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivileged

`func (o *AddedInner) SetIsPrivileged(v bool)`

SetIsPrivileged sets IsPrivileged field to given value.

### HasIsPrivileged

`func (o *AddedInner) HasIsPrivileged() bool`

HasIsPrivileged returns a boolean if a field has been set.

### GetCanEscalatePrivs

`func (o *AddedInner) GetCanEscalatePrivs() bool`

GetCanEscalatePrivs returns the CanEscalatePrivs field if non-nil, zero value otherwise.

### GetCanEscalatePrivsOk

`func (o *AddedInner) GetCanEscalatePrivsOk() (*bool, bool)`

GetCanEscalatePrivsOk returns a tuple with the CanEscalatePrivs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEscalatePrivs

`func (o *AddedInner) SetCanEscalatePrivs(v bool)`

SetCanEscalatePrivs sets CanEscalatePrivs field to given value.

### HasCanEscalatePrivs

`func (o *AddedInner) HasCanEscalatePrivs() bool`

HasCanEscalatePrivs returns a boolean if a field has been set.

### GetIsDisabled

`func (o *AddedInner) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *AddedInner) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *AddedInner) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *AddedInner) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetAccountCreated

`func (o *AddedInner) GetAccountCreated() time.Time`

GetAccountCreated returns the AccountCreated field if non-nil, zero value otherwise.

### GetAccountCreatedOk

`func (o *AddedInner) GetAccountCreatedOk() (*time.Time, bool)`

GetAccountCreatedOk returns a tuple with the AccountCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCreated

`func (o *AddedInner) SetAccountCreated(v time.Time)`

SetAccountCreated sets AccountCreated field to given value.

### HasAccountCreated

`func (o *AddedInner) HasAccountCreated() bool`

HasAccountCreated returns a boolean if a field has been set.

### GetAccountExpires

`func (o *AddedInner) GetAccountExpires() time.Time`

GetAccountExpires returns the AccountExpires field if non-nil, zero value otherwise.

### GetAccountExpiresOk

`func (o *AddedInner) GetAccountExpiresOk() (*time.Time, bool)`

GetAccountExpiresOk returns a tuple with the AccountExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountExpires

`func (o *AddedInner) SetAccountExpires(v time.Time)`

SetAccountExpires sets AccountExpires field to given value.

### HasAccountExpires

`func (o *AddedInner) HasAccountExpires() bool`

HasAccountExpires returns a boolean if a field has been set.

### GetCredentialLastChanged

`func (o *AddedInner) GetCredentialLastChanged() time.Time`

GetCredentialLastChanged returns the CredentialLastChanged field if non-nil, zero value otherwise.

### GetCredentialLastChangedOk

`func (o *AddedInner) GetCredentialLastChangedOk() (*time.Time, bool)`

GetCredentialLastChangedOk returns a tuple with the CredentialLastChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialLastChanged

`func (o *AddedInner) SetCredentialLastChanged(v time.Time)`

SetCredentialLastChanged sets CredentialLastChanged field to given value.

### HasCredentialLastChanged

`func (o *AddedInner) HasCredentialLastChanged() bool`

HasCredentialLastChanged returns a boolean if a field has been set.

### GetAccountFirstLogin

`func (o *AddedInner) GetAccountFirstLogin() time.Time`

GetAccountFirstLogin returns the AccountFirstLogin field if non-nil, zero value otherwise.

### GetAccountFirstLoginOk

`func (o *AddedInner) GetAccountFirstLoginOk() (*time.Time, bool)`

GetAccountFirstLoginOk returns a tuple with the AccountFirstLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountFirstLogin

`func (o *AddedInner) SetAccountFirstLogin(v time.Time)`

SetAccountFirstLogin sets AccountFirstLogin field to given value.

### HasAccountFirstLogin

`func (o *AddedInner) HasAccountFirstLogin() bool`

HasAccountFirstLogin returns a boolean if a field has been set.

### GetAccountLastLogin

`func (o *AddedInner) GetAccountLastLogin() time.Time`

GetAccountLastLogin returns the AccountLastLogin field if non-nil, zero value otherwise.

### GetAccountLastLoginOk

`func (o *AddedInner) GetAccountLastLoginOk() (*time.Time, bool)`

GetAccountLastLoginOk returns a tuple with the AccountLastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLastLogin

`func (o *AddedInner) SetAccountLastLogin(v time.Time)`

SetAccountLastLogin sets AccountLastLogin field to given value.

### HasAccountLastLogin

`func (o *AddedInner) HasAccountLastLogin() bool`

HasAccountLastLogin returns a boolean if a field has been set.

### GetCoin

`func (o *AddedInner) GetCoin() string`

GetCoin returns the Coin field if non-nil, zero value otherwise.

### GetCoinOk

`func (o *AddedInner) GetCoinOk() (*string, bool)`

GetCoinOk returns a tuple with the Coin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoin

`func (o *AddedInner) SetCoin(v string)`

SetCoin sets Coin field to given value.

### HasCoin

`func (o *AddedInner) HasCoin() bool`

HasCoin returns a boolean if a field has been set.

### GetAddress

`func (o *AddedInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AddedInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AddedInner) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *AddedInner) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetLastSeen

`func (o *AddedInner) GetLastSeen() time.Time`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *AddedInner) GetLastSeenOk() (*time.Time, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *AddedInner) SetLastSeen(v time.Time)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *AddedInner) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetFirstSeen

`func (o *AddedInner) GetFirstSeen() time.Time`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *AddedInner) GetFirstSeenOk() (*time.Time, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *AddedInner) SetFirstSeen(v time.Time)`

SetFirstSeen sets FirstSeen field to given value.

### HasFirstSeen

`func (o *AddedInner) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### GetIssuer

`func (o *AddedInner) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *AddedInner) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *AddedInner) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *AddedInner) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetSubject

`func (o *AddedInner) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *AddedInner) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *AddedInner) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *AddedInner) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetSerialNumber

`func (o *AddedInner) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *AddedInner) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *AddedInner) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *AddedInner) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetAfter

`func (o *AddedInner) GetAfter() time.Time`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *AddedInner) GetAfterOk() (*time.Time, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *AddedInner) SetAfter(v time.Time)`

SetAfter sets After field to given value.

### HasAfter

`func (o *AddedInner) HasAfter() bool`

HasAfter returns a boolean if a field has been set.

### GetBefore

`func (o *AddedInner) GetBefore() time.Time`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *AddedInner) GetBeforeOk() (*time.Time, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *AddedInner) SetBefore(v time.Time)`

SetBefore sets Before field to given value.

### HasBefore

`func (o *AddedInner) HasBefore() bool`

HasBefore returns a boolean if a field has been set.

### GetFingerprint

`func (o *AddedInner) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *AddedInner) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *AddedInner) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *AddedInner) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


