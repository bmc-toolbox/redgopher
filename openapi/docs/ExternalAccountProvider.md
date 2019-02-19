# ExternalAccountProvider

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountProviderType** | [**AccountProviderTypes**](AccountProviderTypes.md) |  | [optional] 
**Authentication** | [**Authentication**](Authentication.md) |  | [optional] 
**Certificates** | [**IdRef**](idRef.md) |  | [optional] 
**LDAPService** | [**LdapService**](LDAPService.md) |  | [optional] 
**RemoteRoleMapping** | [**[]RoleMapping**](RoleMapping.md) | This property contains a collection of the mapping rules to convert the external account providers account information to the local Redfish Role. | [optional] 
**ServiceAddresses** | **[]string** | This property contains the addresses of the user account providers this resource references. The format of this field depends on the Type. | [optional] 
**ServiceEnabled** | **bool** | This indicates whether this service is enabled. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


