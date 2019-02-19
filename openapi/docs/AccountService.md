# AccountService

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataContext** | **string** | The OData description of a payload. | [optional] 
**OdataEtag** | **string** | The current ETag of the resource. | [optional] 
**OdataId** | **string** | The unique identifier for a resource. | 
**OdataType** | **string** | The type of a resource. | 
**AccountLockoutCounterResetAfter** | **int32** | The interval of time in seconds between the last failed login attempt and reset of the lockout threshold counter. This value must be less than or equal to AccountLockoutDuration. Reset sets the counter to zero. | [optional] 
**AccountLockoutDuration** | **int32** | The time in seconds an account is locked out. The value must be greater than or equal to the value of the AccountLockoutCounterResetAfter property. If set to 0, no lockout occurs. | [optional] 
**AccountLockoutThreshold** | **int32** | The number of failed login attempts allowed before a user account is locked for a specified duration. A value of 0 means it is never locked. | [optional] 
**Accounts** | [**IdRef**](idRef.md) |  | [optional] 
**Actions** | [**Actions**](Actions.md) |  | [optional] 
**ActiveDirectory** | [**ExternalAccountProvider**](ExternalAccountProvider.md) |  | [optional] 
**AdditionalExternalAccountProviders** | [**IdRef**](idRef.md) |  | [optional] 
**AuthFailureLoggingThreshold** | **int32** | The number of authorization failures allowed before the failure attempt is logged to the manager log. | [optional] 
**Description** | **string** | Provides a description of this resource and is used for commonality  in the schema definitions. | [optional] 
**Id** | **string** | Uniquely identifies the resource within the collection of like resources. | 
**LDAP** | [**ExternalAccountProvider**](ExternalAccountProvider.md) |  | [optional] 
**LocalAccountAuth** | [**LocalAccountAuth**](LocalAccountAuth.md) |  | [optional] 
**MaxPasswordLength** | **int32** | The maximum password length for this service. | [optional] 
**MinPasswordLength** | **int32** | The minimum password length for this service. | [optional] 
**Name** | **string** | The name of the resource or array element. | 
**Oem** | [**Oem**](Oem.md) |  | [optional] 
**PrivilegeMap** | [**IdRef**](idRef.md) |  | [optional] 
**Roles** | [**IdRef**](idRef.md) |  | [optional] 
**ServiceEnabled** | **bool** | Indicates whether this service is enabled.  If set to false, the AccountService is disabled.  This means no users can be created, deleted or modified.  Any service attempting to access the AccountService resource (for example, the Session Service) will fail.  New sessions cannot be started when the service is disabled. However, established sessions may still continue operating. This does not affect Basic AUTH connections. | [optional] 
**Status** | [**Status**](Status.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


