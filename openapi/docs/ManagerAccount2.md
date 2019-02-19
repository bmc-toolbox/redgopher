# ManagerAccount2

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataContext** | **string** | The OData description of a payload. | [optional] 
**OdataId** | **string** | The unique identifier for a resource. | [optional] 
**OdataType** | **string** | The type of a resource. | [optional] 
**Description** | **string** | Provides a description of this resource and is used for commonality  in the schema definitions. | [optional] 
**Enabled** | **bool** | This property is used by a User Administrator to disable an account w/o having to delet the user information.  When set to true, the user can login.  When set to false, the account is administratively disabled and the user cannot login. | [optional] 
**Id** | **string** | Uniquely identifies the resource within the collection of like resources. | 
**Links** | [**ManagerAccount2Links**](ManagerAccount_2_Links.md) |  | [optional] 
**Locked** | **bool** | This property indicates that the account has been auto-locked by the account service because the lockout threshold has been exceeded.  When set to true, the account is locked. A user admin can write the property to false to manually unlock, or the account service will unlock it once the lockout duration period has passed. | [optional] 
**Name** | **string** | The name of the resource or array element. | 
**Oem** | [**Oem**](Oem.md) |  | [optional] 
**Password** | **string** | This property is used with a PATCH or PUT to write the password for the account.  This property is null on a GET. | [optional] 
**RoleId** | **string** | This property contains the Role for this account. | [optional] 
**UserName** | **string** | This property contains the user name for the account. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


