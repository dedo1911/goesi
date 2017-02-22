# \MailApi

All URIs are relative to *https://esi.tech.ccp.is/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCharactersCharacterIdMailLabelsLabelId**](MailApi.md#DeleteCharactersCharacterIdMailLabelsLabelId) | **Delete** /characters/{character_id}/mail/labels/{label_id}/ | Delete a mail label
[**DeleteCharactersCharacterIdMailMailId**](MailApi.md#DeleteCharactersCharacterIdMailMailId) | **Delete** /characters/{character_id}/mail/{mail_id}/ | Delete a mail
[**GetCharactersCharacterIdMail**](MailApi.md#GetCharactersCharacterIdMail) | **Get** /characters/{character_id}/mail/ | Return mail headers
[**GetCharactersCharacterIdMailLists**](MailApi.md#GetCharactersCharacterIdMailLists) | **Get** /characters/{character_id}/mail/lists/ | Return mailing list subscriptions
[**GetCharactersCharacterIdMailMailId**](MailApi.md#GetCharactersCharacterIdMailMailId) | **Get** /characters/{character_id}/mail/{mail_id}/ | Return a mail
[**GetCharactersCharacterIdMailUnread**](MailApi.md#GetCharactersCharacterIdMailUnread) | **Get** /characters/{character_id}/mail/unread/ | Return the number of unread mails
[**PostCharactersCharacterIdMail**](MailApi.md#PostCharactersCharacterIdMail) | **Post** /characters/{character_id}/mail/ | Send a new mail
[**PutCharactersCharacterIdMailMailId**](MailApi.md#PutCharactersCharacterIdMailMailId) | **Put** /characters/{character_id}/mail/{mail_id}/ | Update metadata about a mail


# **DeleteCharactersCharacterIdMailLabelsLabelId**
> DeleteCharactersCharacterIdMailLabelsLabelId(ctx, characterId, labelId, optional)
Delete a mail label

Delete a mail label  ---  Alternate route: `/legacy/characters/{character_id}/mail/labels/{label_id}/`  Alternate route: `/latest/characters/{character_id}/mail/labels/{label_id}/`  Alternate route: `/dev/characters/{character_id}/mail/labels/{label_id}/` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **characterId** | **int32**| An EVE character ID | 
  **labelId** | **int32**| An EVE label id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int32**| An EVE character ID | 
 **labelId** | **int32**| An EVE label id | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **token** | **string**| Access token to use, if preferred over a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

 (empty response body)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCharactersCharacterIdMailMailId**
> DeleteCharactersCharacterIdMailMailId(ctx, characterId, mailId, optional)
Delete a mail

Delete a mail  ---  Alternate route: `/legacy/characters/{character_id}/mail/{mail_id}/`  Alternate route: `/latest/characters/{character_id}/mail/{mail_id}/`  Alternate route: `/dev/characters/{character_id}/mail/{mail_id}/` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **characterId** | **int32**| An EVE character ID | 
  **mailId** | **int32**| An EVE mail ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int32**| An EVE character ID | 
 **mailId** | **int32**| An EVE mail ID | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **token** | **string**| Access token to use, if preferred over a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

 (empty response body)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersCharacterIdMail**
> []GetCharactersCharacterIdMail200Ok GetCharactersCharacterIdMail(ctx, characterId, optional)
Return mail headers

Return the 50 most recent mail headers belonging to the character that match the query criteria. Queries can be filtered by label, and last_mail_id can be used to paginate backwards.  ---  Alternate route: `/legacy/characters/{character_id}/mail/`  Alternate route: `/latest/characters/{character_id}/mail/`  Alternate route: `/dev/characters/{character_id}/mail/`   ---  This route is cached for up to 30 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **characterId** | **int32**| An EVE character ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int32**| An EVE character ID | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **labels** | [**[]int64**](int64.md)| Fetch only mails that match one or more of the given labels | 
 **lastMailId** | **int32**| List only mail with an ID lower than the given ID, if present | 
 **token** | **string**| Access token to use, if preferred over a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**[]GetCharactersCharacterIdMail200Ok**](get_characters_character_id_mail_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersCharacterIdMailLists**
> []GetCharactersCharacterIdMailLists200Ok GetCharactersCharacterIdMailLists(ctx, characterId, optional)
Return mailing list subscriptions

Return all mailing lists that the character is subscribed to   ---  Alternate route: `/legacy/characters/{character_id}/mail/lists/`  Alternate route: `/latest/characters/{character_id}/mail/lists/`  Alternate route: `/dev/characters/{character_id}/mail/lists/`   ---  This route is cached for up to 120 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **characterId** | **int32**| An EVE character ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int32**| An EVE character ID | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **token** | **string**| Access token to use, if preferred over a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**[]GetCharactersCharacterIdMailLists200Ok**](get_characters_character_id_mail_lists_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersCharacterIdMailMailId**
> GetCharactersCharacterIdMailMailIdOk GetCharactersCharacterIdMailMailId(ctx, characterId, mailId, optional)
Return a mail

Return the contents of an EVE mail  ---  Alternate route: `/legacy/characters/{character_id}/mail/{mail_id}/`  Alternate route: `/latest/characters/{character_id}/mail/{mail_id}/`  Alternate route: `/dev/characters/{character_id}/mail/{mail_id}/`   ---  This route is cached for up to 30 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **characterId** | **int32**| An EVE character ID | 
  **mailId** | **int32**| An EVE mail ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int32**| An EVE character ID | 
 **mailId** | **int32**| An EVE mail ID | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **token** | **string**| Access token to use, if preferred over a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**GetCharactersCharacterIdMailMailIdOk**](get_characters_character_id_mail_mail_id_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersCharacterIdMailUnread**
> int32 GetCharactersCharacterIdMailUnread(ctx, characterId, optional)
Return the number of unread mails

Return the number of unread mails for the character   ---  Alternate route: `/legacy/characters/{character_id}/mail/unread/`   ---  This route is cached for up to 5 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **characterId** | **int32**| An EVE character ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int32**| An EVE character ID | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **token** | **string**| Access token to use, if preferred over a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

**int32**

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCharactersCharacterIdMail**
> int32 PostCharactersCharacterIdMail(ctx, characterId, mail, optional)
Send a new mail

Create and send a new mail  ---  Alternate route: `/legacy/characters/{character_id}/mail/`  Alternate route: `/latest/characters/{character_id}/mail/`  Alternate route: `/dev/characters/{character_id}/mail/` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **characterId** | **int32**| The sender&#39;s character ID | 
  **mail** | [**PostCharactersCharacterIdMailMail**](PostCharactersCharacterIdMailMail.md)| The mail to send | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int32**| The sender&#39;s character ID | 
 **mail** | [**PostCharactersCharacterIdMailMail**](PostCharactersCharacterIdMailMail.md)| The mail to send | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **token** | **string**| Access token to use, if preferred over a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

**int32**

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutCharactersCharacterIdMailMailId**
> PutCharactersCharacterIdMailMailId(ctx, characterId, contents, mailId, optional)
Update metadata about a mail

Update metadata about a mail  ---  Alternate route: `/legacy/characters/{character_id}/mail/{mail_id}/`  Alternate route: `/latest/characters/{character_id}/mail/{mail_id}/`  Alternate route: `/dev/characters/{character_id}/mail/{mail_id}/` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **characterId** | **int32**| An EVE character ID | 
  **contents** | [**PutCharactersCharacterIdMailMailIdContents**](PutCharactersCharacterIdMailMailIdContents.md)| Data used to update the mail | 
  **mailId** | **int32**| An EVE mail ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int32**| An EVE character ID | 
 **contents** | [**PutCharactersCharacterIdMailMailIdContents**](PutCharactersCharacterIdMailMailIdContents.md)| Data used to update the mail | 
 **mailId** | **int32**| An EVE mail ID | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **token** | **string**| Access token to use, if preferred over a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

 (empty response body)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
