# UserAccountInput

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

## Methods

### NewUserAccountInput

`func NewUserAccountInput(value string, ) *UserAccountInput`

NewUserAccountInput instantiates a new UserAccountInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAccountInputWithDefaults

`func NewUserAccountInputWithDefaults() *UserAccountInput`

NewUserAccountInputWithDefaults instantiates a new UserAccountInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *UserAccountInput) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UserAccountInput) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UserAccountInput) SetValue(v string)`

SetValue sets Value field to given value.


### GetType

`func (o *UserAccountInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserAccountInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserAccountInput) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserAccountInput) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreated

`func (o *UserAccountInput) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *UserAccountInput) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *UserAccountInput) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *UserAccountInput) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetContext

`func (o *UserAccountInput) GetContext() []map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *UserAccountInput) GetContextOk() (*[]map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *UserAccountInput) SetContext(v []map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *UserAccountInput) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetLastAnalysis

`func (o *UserAccountInput) GetLastAnalysis() map[string]time.Time`

GetLastAnalysis returns the LastAnalysis field if non-nil, zero value otherwise.

### GetLastAnalysisOk

`func (o *UserAccountInput) GetLastAnalysisOk() (*map[string]time.Time, bool)`

GetLastAnalysisOk returns a tuple with the LastAnalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAnalysis

`func (o *UserAccountInput) SetLastAnalysis(v map[string]time.Time)`

SetLastAnalysis sets LastAnalysis field to given value.

### HasLastAnalysis

`func (o *UserAccountInput) HasLastAnalysis() bool`

HasLastAnalysis returns a boolean if a field has been set.

### GetUserId

`func (o *UserAccountInput) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserAccountInput) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserAccountInput) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UserAccountInput) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *UserAccountInput) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *UserAccountInput) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetCredential

`func (o *UserAccountInput) GetCredential() string`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *UserAccountInput) GetCredentialOk() (*string, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *UserAccountInput) SetCredential(v string)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *UserAccountInput) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### SetCredentialNil

`func (o *UserAccountInput) SetCredentialNil(b bool)`

 SetCredentialNil sets the value for Credential to be an explicit nil

### UnsetCredential
`func (o *UserAccountInput) UnsetCredential()`

UnsetCredential ensures that no value is present for Credential, not even an explicit nil
### GetAccountLogin

`func (o *UserAccountInput) GetAccountLogin() string`

GetAccountLogin returns the AccountLogin field if non-nil, zero value otherwise.

### GetAccountLoginOk

`func (o *UserAccountInput) GetAccountLoginOk() (*string, bool)`

GetAccountLoginOk returns a tuple with the AccountLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLogin

`func (o *UserAccountInput) SetAccountLogin(v string)`

SetAccountLogin sets AccountLogin field to given value.

### HasAccountLogin

`func (o *UserAccountInput) HasAccountLogin() bool`

HasAccountLogin returns a boolean if a field has been set.

### SetAccountLoginNil

`func (o *UserAccountInput) SetAccountLoginNil(b bool)`

 SetAccountLoginNil sets the value for AccountLogin to be an explicit nil

### UnsetAccountLogin
`func (o *UserAccountInput) UnsetAccountLogin()`

UnsetAccountLogin ensures that no value is present for AccountLogin, not even an explicit nil
### GetAccountType

`func (o *UserAccountInput) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *UserAccountInput) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *UserAccountInput) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *UserAccountInput) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### SetAccountTypeNil

`func (o *UserAccountInput) SetAccountTypeNil(b bool)`

 SetAccountTypeNil sets the value for AccountType to be an explicit nil

### UnsetAccountType
`func (o *UserAccountInput) UnsetAccountType()`

UnsetAccountType ensures that no value is present for AccountType, not even an explicit nil
### GetDisplayName

`func (o *UserAccountInput) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UserAccountInput) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UserAccountInput) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UserAccountInput) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *UserAccountInput) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *UserAccountInput) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetIsServiceAccount

`func (o *UserAccountInput) GetIsServiceAccount() bool`

GetIsServiceAccount returns the IsServiceAccount field if non-nil, zero value otherwise.

### GetIsServiceAccountOk

`func (o *UserAccountInput) GetIsServiceAccountOk() (*bool, bool)`

GetIsServiceAccountOk returns a tuple with the IsServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsServiceAccount

`func (o *UserAccountInput) SetIsServiceAccount(v bool)`

SetIsServiceAccount sets IsServiceAccount field to given value.

### HasIsServiceAccount

`func (o *UserAccountInput) HasIsServiceAccount() bool`

HasIsServiceAccount returns a boolean if a field has been set.

### SetIsServiceAccountNil

`func (o *UserAccountInput) SetIsServiceAccountNil(b bool)`

 SetIsServiceAccountNil sets the value for IsServiceAccount to be an explicit nil

### UnsetIsServiceAccount
`func (o *UserAccountInput) UnsetIsServiceAccount()`

UnsetIsServiceAccount ensures that no value is present for IsServiceAccount, not even an explicit nil
### GetIsPrivileged

`func (o *UserAccountInput) GetIsPrivileged() bool`

GetIsPrivileged returns the IsPrivileged field if non-nil, zero value otherwise.

### GetIsPrivilegedOk

`func (o *UserAccountInput) GetIsPrivilegedOk() (*bool, bool)`

GetIsPrivilegedOk returns a tuple with the IsPrivileged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivileged

`func (o *UserAccountInput) SetIsPrivileged(v bool)`

SetIsPrivileged sets IsPrivileged field to given value.

### HasIsPrivileged

`func (o *UserAccountInput) HasIsPrivileged() bool`

HasIsPrivileged returns a boolean if a field has been set.

### SetIsPrivilegedNil

`func (o *UserAccountInput) SetIsPrivilegedNil(b bool)`

 SetIsPrivilegedNil sets the value for IsPrivileged to be an explicit nil

### UnsetIsPrivileged
`func (o *UserAccountInput) UnsetIsPrivileged()`

UnsetIsPrivileged ensures that no value is present for IsPrivileged, not even an explicit nil
### GetCanEscalatePrivs

`func (o *UserAccountInput) GetCanEscalatePrivs() bool`

GetCanEscalatePrivs returns the CanEscalatePrivs field if non-nil, zero value otherwise.

### GetCanEscalatePrivsOk

`func (o *UserAccountInput) GetCanEscalatePrivsOk() (*bool, bool)`

GetCanEscalatePrivsOk returns a tuple with the CanEscalatePrivs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEscalatePrivs

`func (o *UserAccountInput) SetCanEscalatePrivs(v bool)`

SetCanEscalatePrivs sets CanEscalatePrivs field to given value.

### HasCanEscalatePrivs

`func (o *UserAccountInput) HasCanEscalatePrivs() bool`

HasCanEscalatePrivs returns a boolean if a field has been set.

### SetCanEscalatePrivsNil

`func (o *UserAccountInput) SetCanEscalatePrivsNil(b bool)`

 SetCanEscalatePrivsNil sets the value for CanEscalatePrivs to be an explicit nil

### UnsetCanEscalatePrivs
`func (o *UserAccountInput) UnsetCanEscalatePrivs()`

UnsetCanEscalatePrivs ensures that no value is present for CanEscalatePrivs, not even an explicit nil
### GetIsDisabled

`func (o *UserAccountInput) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *UserAccountInput) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *UserAccountInput) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *UserAccountInput) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### SetIsDisabledNil

`func (o *UserAccountInput) SetIsDisabledNil(b bool)`

 SetIsDisabledNil sets the value for IsDisabled to be an explicit nil

### UnsetIsDisabled
`func (o *UserAccountInput) UnsetIsDisabled()`

UnsetIsDisabled ensures that no value is present for IsDisabled, not even an explicit nil
### GetAccountCreated

`func (o *UserAccountInput) GetAccountCreated() time.Time`

GetAccountCreated returns the AccountCreated field if non-nil, zero value otherwise.

### GetAccountCreatedOk

`func (o *UserAccountInput) GetAccountCreatedOk() (*time.Time, bool)`

GetAccountCreatedOk returns a tuple with the AccountCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCreated

`func (o *UserAccountInput) SetAccountCreated(v time.Time)`

SetAccountCreated sets AccountCreated field to given value.

### HasAccountCreated

`func (o *UserAccountInput) HasAccountCreated() bool`

HasAccountCreated returns a boolean if a field has been set.

### SetAccountCreatedNil

`func (o *UserAccountInput) SetAccountCreatedNil(b bool)`

 SetAccountCreatedNil sets the value for AccountCreated to be an explicit nil

### UnsetAccountCreated
`func (o *UserAccountInput) UnsetAccountCreated()`

UnsetAccountCreated ensures that no value is present for AccountCreated, not even an explicit nil
### GetAccountExpires

`func (o *UserAccountInput) GetAccountExpires() time.Time`

GetAccountExpires returns the AccountExpires field if non-nil, zero value otherwise.

### GetAccountExpiresOk

`func (o *UserAccountInput) GetAccountExpiresOk() (*time.Time, bool)`

GetAccountExpiresOk returns a tuple with the AccountExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountExpires

`func (o *UserAccountInput) SetAccountExpires(v time.Time)`

SetAccountExpires sets AccountExpires field to given value.

### HasAccountExpires

`func (o *UserAccountInput) HasAccountExpires() bool`

HasAccountExpires returns a boolean if a field has been set.

### SetAccountExpiresNil

`func (o *UserAccountInput) SetAccountExpiresNil(b bool)`

 SetAccountExpiresNil sets the value for AccountExpires to be an explicit nil

### UnsetAccountExpires
`func (o *UserAccountInput) UnsetAccountExpires()`

UnsetAccountExpires ensures that no value is present for AccountExpires, not even an explicit nil
### GetCredentialLastChanged

`func (o *UserAccountInput) GetCredentialLastChanged() time.Time`

GetCredentialLastChanged returns the CredentialLastChanged field if non-nil, zero value otherwise.

### GetCredentialLastChangedOk

`func (o *UserAccountInput) GetCredentialLastChangedOk() (*time.Time, bool)`

GetCredentialLastChangedOk returns a tuple with the CredentialLastChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialLastChanged

`func (o *UserAccountInput) SetCredentialLastChanged(v time.Time)`

SetCredentialLastChanged sets CredentialLastChanged field to given value.

### HasCredentialLastChanged

`func (o *UserAccountInput) HasCredentialLastChanged() bool`

HasCredentialLastChanged returns a boolean if a field has been set.

### SetCredentialLastChangedNil

`func (o *UserAccountInput) SetCredentialLastChangedNil(b bool)`

 SetCredentialLastChangedNil sets the value for CredentialLastChanged to be an explicit nil

### UnsetCredentialLastChanged
`func (o *UserAccountInput) UnsetCredentialLastChanged()`

UnsetCredentialLastChanged ensures that no value is present for CredentialLastChanged, not even an explicit nil
### GetAccountFirstLogin

`func (o *UserAccountInput) GetAccountFirstLogin() time.Time`

GetAccountFirstLogin returns the AccountFirstLogin field if non-nil, zero value otherwise.

### GetAccountFirstLoginOk

`func (o *UserAccountInput) GetAccountFirstLoginOk() (*time.Time, bool)`

GetAccountFirstLoginOk returns a tuple with the AccountFirstLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountFirstLogin

`func (o *UserAccountInput) SetAccountFirstLogin(v time.Time)`

SetAccountFirstLogin sets AccountFirstLogin field to given value.

### HasAccountFirstLogin

`func (o *UserAccountInput) HasAccountFirstLogin() bool`

HasAccountFirstLogin returns a boolean if a field has been set.

### SetAccountFirstLoginNil

`func (o *UserAccountInput) SetAccountFirstLoginNil(b bool)`

 SetAccountFirstLoginNil sets the value for AccountFirstLogin to be an explicit nil

### UnsetAccountFirstLogin
`func (o *UserAccountInput) UnsetAccountFirstLogin()`

UnsetAccountFirstLogin ensures that no value is present for AccountFirstLogin, not even an explicit nil
### GetAccountLastLogin

`func (o *UserAccountInput) GetAccountLastLogin() time.Time`

GetAccountLastLogin returns the AccountLastLogin field if non-nil, zero value otherwise.

### GetAccountLastLoginOk

`func (o *UserAccountInput) GetAccountLastLoginOk() (*time.Time, bool)`

GetAccountLastLoginOk returns a tuple with the AccountLastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLastLogin

`func (o *UserAccountInput) SetAccountLastLogin(v time.Time)`

SetAccountLastLogin sets AccountLastLogin field to given value.

### HasAccountLastLogin

`func (o *UserAccountInput) HasAccountLastLogin() bool`

HasAccountLastLogin returns a boolean if a field has been set.

### SetAccountLastLoginNil

`func (o *UserAccountInput) SetAccountLastLoginNil(b bool)`

 SetAccountLastLoginNil sets the value for AccountLastLogin to be an explicit nil

### UnsetAccountLastLogin
`func (o *UserAccountInput) UnsetAccountLastLogin()`

UnsetAccountLastLogin ensures that no value is present for AccountLastLogin, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


