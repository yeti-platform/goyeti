# UserAccountOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] [default to "user_account"]
**Created** | Pointer to **time.Time** |  | [optional] 
**Context** | Pointer to **[]map[string]interface{}** |  | [optional] [default to []]
**LastAnalysis** | Pointer to [**map[string]time.Time**](time.Time.md) |  | [optional] [default to {}]
**UserId** | Pointer to **NullableString** |  | [optional] 
**Credential** | Pointer to **NullableString** |  | [optional] 
**AccountLogin** | Pointer to **NullableString** |  | [optional] 
**AccountType** | Pointer to **NullableString** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**IsServiceAccount** | Pointer to **NullableBool** |  | [optional] 
**IsPrivileged** | Pointer to **NullableBool** |  | [optional] 
**CanEscalatePrivs** | Pointer to **NullableBool** |  | [optional] 
**IsDisabled** | Pointer to **NullableBool** |  | [optional] 
**AccountCreated** | Pointer to **NullableTime** |  | [optional] 
**AccountExpires** | Pointer to **NullableTime** |  | [optional] 
**CredentialLastChanged** | Pointer to **NullableTime** |  | [optional] 
**AccountFirstLogin** | Pointer to **NullableTime** |  | [optional] 
**AccountLastLogin** | Pointer to **NullableTime** |  | [optional] 
**Id** | **string** |  | [readonly] 
**Tags** | [**map[string]TagRelationshipOutput**](TagRelationshipOutput.md) |  | [readonly] 
**RootType** | **string** |  | [readonly] 

## Methods

### NewUserAccountOutput

`func NewUserAccountOutput(value string, id string, tags map[string]TagRelationshipOutput, rootType string, ) *UserAccountOutput`

NewUserAccountOutput instantiates a new UserAccountOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAccountOutputWithDefaults

`func NewUserAccountOutputWithDefaults() *UserAccountOutput`

NewUserAccountOutputWithDefaults instantiates a new UserAccountOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *UserAccountOutput) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UserAccountOutput) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UserAccountOutput) SetValue(v string)`

SetValue sets Value field to given value.


### GetType

`func (o *UserAccountOutput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserAccountOutput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserAccountOutput) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserAccountOutput) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreated

`func (o *UserAccountOutput) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *UserAccountOutput) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *UserAccountOutput) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *UserAccountOutput) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetContext

`func (o *UserAccountOutput) GetContext() []map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *UserAccountOutput) GetContextOk() (*[]map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *UserAccountOutput) SetContext(v []map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *UserAccountOutput) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetLastAnalysis

`func (o *UserAccountOutput) GetLastAnalysis() map[string]time.Time`

GetLastAnalysis returns the LastAnalysis field if non-nil, zero value otherwise.

### GetLastAnalysisOk

`func (o *UserAccountOutput) GetLastAnalysisOk() (*map[string]time.Time, bool)`

GetLastAnalysisOk returns a tuple with the LastAnalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAnalysis

`func (o *UserAccountOutput) SetLastAnalysis(v map[string]time.Time)`

SetLastAnalysis sets LastAnalysis field to given value.

### HasLastAnalysis

`func (o *UserAccountOutput) HasLastAnalysis() bool`

HasLastAnalysis returns a boolean if a field has been set.

### GetUserId

`func (o *UserAccountOutput) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserAccountOutput) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserAccountOutput) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UserAccountOutput) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *UserAccountOutput) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *UserAccountOutput) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetCredential

`func (o *UserAccountOutput) GetCredential() string`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *UserAccountOutput) GetCredentialOk() (*string, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *UserAccountOutput) SetCredential(v string)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *UserAccountOutput) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### SetCredentialNil

`func (o *UserAccountOutput) SetCredentialNil(b bool)`

 SetCredentialNil sets the value for Credential to be an explicit nil

### UnsetCredential
`func (o *UserAccountOutput) UnsetCredential()`

UnsetCredential ensures that no value is present for Credential, not even an explicit nil
### GetAccountLogin

`func (o *UserAccountOutput) GetAccountLogin() string`

GetAccountLogin returns the AccountLogin field if non-nil, zero value otherwise.

### GetAccountLoginOk

`func (o *UserAccountOutput) GetAccountLoginOk() (*string, bool)`

GetAccountLoginOk returns a tuple with the AccountLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLogin

`func (o *UserAccountOutput) SetAccountLogin(v string)`

SetAccountLogin sets AccountLogin field to given value.

### HasAccountLogin

`func (o *UserAccountOutput) HasAccountLogin() bool`

HasAccountLogin returns a boolean if a field has been set.

### SetAccountLoginNil

`func (o *UserAccountOutput) SetAccountLoginNil(b bool)`

 SetAccountLoginNil sets the value for AccountLogin to be an explicit nil

### UnsetAccountLogin
`func (o *UserAccountOutput) UnsetAccountLogin()`

UnsetAccountLogin ensures that no value is present for AccountLogin, not even an explicit nil
### GetAccountType

`func (o *UserAccountOutput) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *UserAccountOutput) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *UserAccountOutput) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *UserAccountOutput) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### SetAccountTypeNil

`func (o *UserAccountOutput) SetAccountTypeNil(b bool)`

 SetAccountTypeNil sets the value for AccountType to be an explicit nil

### UnsetAccountType
`func (o *UserAccountOutput) UnsetAccountType()`

UnsetAccountType ensures that no value is present for AccountType, not even an explicit nil
### GetDisplayName

`func (o *UserAccountOutput) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UserAccountOutput) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UserAccountOutput) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UserAccountOutput) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *UserAccountOutput) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *UserAccountOutput) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetIsServiceAccount

`func (o *UserAccountOutput) GetIsServiceAccount() bool`

GetIsServiceAccount returns the IsServiceAccount field if non-nil, zero value otherwise.

### GetIsServiceAccountOk

`func (o *UserAccountOutput) GetIsServiceAccountOk() (*bool, bool)`

GetIsServiceAccountOk returns a tuple with the IsServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsServiceAccount

`func (o *UserAccountOutput) SetIsServiceAccount(v bool)`

SetIsServiceAccount sets IsServiceAccount field to given value.

### HasIsServiceAccount

`func (o *UserAccountOutput) HasIsServiceAccount() bool`

HasIsServiceAccount returns a boolean if a field has been set.

### SetIsServiceAccountNil

`func (o *UserAccountOutput) SetIsServiceAccountNil(b bool)`

 SetIsServiceAccountNil sets the value for IsServiceAccount to be an explicit nil

### UnsetIsServiceAccount
`func (o *UserAccountOutput) UnsetIsServiceAccount()`

UnsetIsServiceAccount ensures that no value is present for IsServiceAccount, not even an explicit nil
### GetIsPrivileged

`func (o *UserAccountOutput) GetIsPrivileged() bool`

GetIsPrivileged returns the IsPrivileged field if non-nil, zero value otherwise.

### GetIsPrivilegedOk

`func (o *UserAccountOutput) GetIsPrivilegedOk() (*bool, bool)`

GetIsPrivilegedOk returns a tuple with the IsPrivileged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivileged

`func (o *UserAccountOutput) SetIsPrivileged(v bool)`

SetIsPrivileged sets IsPrivileged field to given value.

### HasIsPrivileged

`func (o *UserAccountOutput) HasIsPrivileged() bool`

HasIsPrivileged returns a boolean if a field has been set.

### SetIsPrivilegedNil

`func (o *UserAccountOutput) SetIsPrivilegedNil(b bool)`

 SetIsPrivilegedNil sets the value for IsPrivileged to be an explicit nil

### UnsetIsPrivileged
`func (o *UserAccountOutput) UnsetIsPrivileged()`

UnsetIsPrivileged ensures that no value is present for IsPrivileged, not even an explicit nil
### GetCanEscalatePrivs

`func (o *UserAccountOutput) GetCanEscalatePrivs() bool`

GetCanEscalatePrivs returns the CanEscalatePrivs field if non-nil, zero value otherwise.

### GetCanEscalatePrivsOk

`func (o *UserAccountOutput) GetCanEscalatePrivsOk() (*bool, bool)`

GetCanEscalatePrivsOk returns a tuple with the CanEscalatePrivs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEscalatePrivs

`func (o *UserAccountOutput) SetCanEscalatePrivs(v bool)`

SetCanEscalatePrivs sets CanEscalatePrivs field to given value.

### HasCanEscalatePrivs

`func (o *UserAccountOutput) HasCanEscalatePrivs() bool`

HasCanEscalatePrivs returns a boolean if a field has been set.

### SetCanEscalatePrivsNil

`func (o *UserAccountOutput) SetCanEscalatePrivsNil(b bool)`

 SetCanEscalatePrivsNil sets the value for CanEscalatePrivs to be an explicit nil

### UnsetCanEscalatePrivs
`func (o *UserAccountOutput) UnsetCanEscalatePrivs()`

UnsetCanEscalatePrivs ensures that no value is present for CanEscalatePrivs, not even an explicit nil
### GetIsDisabled

`func (o *UserAccountOutput) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *UserAccountOutput) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *UserAccountOutput) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *UserAccountOutput) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### SetIsDisabledNil

`func (o *UserAccountOutput) SetIsDisabledNil(b bool)`

 SetIsDisabledNil sets the value for IsDisabled to be an explicit nil

### UnsetIsDisabled
`func (o *UserAccountOutput) UnsetIsDisabled()`

UnsetIsDisabled ensures that no value is present for IsDisabled, not even an explicit nil
### GetAccountCreated

`func (o *UserAccountOutput) GetAccountCreated() time.Time`

GetAccountCreated returns the AccountCreated field if non-nil, zero value otherwise.

### GetAccountCreatedOk

`func (o *UserAccountOutput) GetAccountCreatedOk() (*time.Time, bool)`

GetAccountCreatedOk returns a tuple with the AccountCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCreated

`func (o *UserAccountOutput) SetAccountCreated(v time.Time)`

SetAccountCreated sets AccountCreated field to given value.

### HasAccountCreated

`func (o *UserAccountOutput) HasAccountCreated() bool`

HasAccountCreated returns a boolean if a field has been set.

### SetAccountCreatedNil

`func (o *UserAccountOutput) SetAccountCreatedNil(b bool)`

 SetAccountCreatedNil sets the value for AccountCreated to be an explicit nil

### UnsetAccountCreated
`func (o *UserAccountOutput) UnsetAccountCreated()`

UnsetAccountCreated ensures that no value is present for AccountCreated, not even an explicit nil
### GetAccountExpires

`func (o *UserAccountOutput) GetAccountExpires() time.Time`

GetAccountExpires returns the AccountExpires field if non-nil, zero value otherwise.

### GetAccountExpiresOk

`func (o *UserAccountOutput) GetAccountExpiresOk() (*time.Time, bool)`

GetAccountExpiresOk returns a tuple with the AccountExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountExpires

`func (o *UserAccountOutput) SetAccountExpires(v time.Time)`

SetAccountExpires sets AccountExpires field to given value.

### HasAccountExpires

`func (o *UserAccountOutput) HasAccountExpires() bool`

HasAccountExpires returns a boolean if a field has been set.

### SetAccountExpiresNil

`func (o *UserAccountOutput) SetAccountExpiresNil(b bool)`

 SetAccountExpiresNil sets the value for AccountExpires to be an explicit nil

### UnsetAccountExpires
`func (o *UserAccountOutput) UnsetAccountExpires()`

UnsetAccountExpires ensures that no value is present for AccountExpires, not even an explicit nil
### GetCredentialLastChanged

`func (o *UserAccountOutput) GetCredentialLastChanged() time.Time`

GetCredentialLastChanged returns the CredentialLastChanged field if non-nil, zero value otherwise.

### GetCredentialLastChangedOk

`func (o *UserAccountOutput) GetCredentialLastChangedOk() (*time.Time, bool)`

GetCredentialLastChangedOk returns a tuple with the CredentialLastChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialLastChanged

`func (o *UserAccountOutput) SetCredentialLastChanged(v time.Time)`

SetCredentialLastChanged sets CredentialLastChanged field to given value.

### HasCredentialLastChanged

`func (o *UserAccountOutput) HasCredentialLastChanged() bool`

HasCredentialLastChanged returns a boolean if a field has been set.

### SetCredentialLastChangedNil

`func (o *UserAccountOutput) SetCredentialLastChangedNil(b bool)`

 SetCredentialLastChangedNil sets the value for CredentialLastChanged to be an explicit nil

### UnsetCredentialLastChanged
`func (o *UserAccountOutput) UnsetCredentialLastChanged()`

UnsetCredentialLastChanged ensures that no value is present for CredentialLastChanged, not even an explicit nil
### GetAccountFirstLogin

`func (o *UserAccountOutput) GetAccountFirstLogin() time.Time`

GetAccountFirstLogin returns the AccountFirstLogin field if non-nil, zero value otherwise.

### GetAccountFirstLoginOk

`func (o *UserAccountOutput) GetAccountFirstLoginOk() (*time.Time, bool)`

GetAccountFirstLoginOk returns a tuple with the AccountFirstLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountFirstLogin

`func (o *UserAccountOutput) SetAccountFirstLogin(v time.Time)`

SetAccountFirstLogin sets AccountFirstLogin field to given value.

### HasAccountFirstLogin

`func (o *UserAccountOutput) HasAccountFirstLogin() bool`

HasAccountFirstLogin returns a boolean if a field has been set.

### SetAccountFirstLoginNil

`func (o *UserAccountOutput) SetAccountFirstLoginNil(b bool)`

 SetAccountFirstLoginNil sets the value for AccountFirstLogin to be an explicit nil

### UnsetAccountFirstLogin
`func (o *UserAccountOutput) UnsetAccountFirstLogin()`

UnsetAccountFirstLogin ensures that no value is present for AccountFirstLogin, not even an explicit nil
### GetAccountLastLogin

`func (o *UserAccountOutput) GetAccountLastLogin() time.Time`

GetAccountLastLogin returns the AccountLastLogin field if non-nil, zero value otherwise.

### GetAccountLastLoginOk

`func (o *UserAccountOutput) GetAccountLastLoginOk() (*time.Time, bool)`

GetAccountLastLoginOk returns a tuple with the AccountLastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLastLogin

`func (o *UserAccountOutput) SetAccountLastLogin(v time.Time)`

SetAccountLastLogin sets AccountLastLogin field to given value.

### HasAccountLastLogin

`func (o *UserAccountOutput) HasAccountLastLogin() bool`

HasAccountLastLogin returns a boolean if a field has been set.

### SetAccountLastLoginNil

`func (o *UserAccountOutput) SetAccountLastLoginNil(b bool)`

 SetAccountLastLoginNil sets the value for AccountLastLogin to be an explicit nil

### UnsetAccountLastLogin
`func (o *UserAccountOutput) UnsetAccountLastLogin()`

UnsetAccountLastLogin ensures that no value is present for AccountLastLogin, not even an explicit nil
### GetId

`func (o *UserAccountOutput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserAccountOutput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserAccountOutput) SetId(v string)`

SetId sets Id field to given value.


### GetTags

`func (o *UserAccountOutput) GetTags() map[string]TagRelationshipOutput`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UserAccountOutput) GetTagsOk() (*map[string]TagRelationshipOutput, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UserAccountOutput) SetTags(v map[string]TagRelationshipOutput)`

SetTags sets Tags field to given value.


### GetRootType

`func (o *UserAccountOutput) GetRootType() string`

GetRootType returns the RootType field if non-nil, zero value otherwise.

### GetRootTypeOk

`func (o *UserAccountOutput) GetRootTypeOk() (*string, bool)`

GetRootTypeOk returns a tuple with the RootType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootType

`func (o *UserAccountOutput) SetRootType(v string)`

SetRootType sets RootType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


