# polyaxon_sdk.OrganizationsV1Api

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**organizations_v1_create_organization**](OrganizationsV1Api.md#organizations_v1_create_organization) | **POST** /api/v1/orgs/create | 
[**organizations_v1_create_organization_member**](OrganizationsV1Api.md#organizations_v1_create_organization_member) | **POST** /api/v1/orgs/{owner}/members | 
[**organizations_v1_delete_organization**](OrganizationsV1Api.md#organizations_v1_delete_organization) | **DELETE** /api/v1/orgs/{owner} | 
[**organizations_v1_delete_organization_member**](OrganizationsV1Api.md#organizations_v1_delete_organization_member) | **DELETE** /api/v1/orgs/{owner}/members/{user} | 
[**organizations_v1_get_organization**](OrganizationsV1Api.md#organizations_v1_get_organization) | **GET** /api/v1/orgs/{owner} | 
[**organizations_v1_get_organization_member**](OrganizationsV1Api.md#organizations_v1_get_organization_member) | **GET** /api/v1/orgs/{owner}/members/{user} | 
[**organizations_v1_list_organization_members**](OrganizationsV1Api.md#organizations_v1_list_organization_members) | **GET** /api/v1/orgs/{owner}/members | 
[**organizations_v1_list_organization_names**](OrganizationsV1Api.md#organizations_v1_list_organization_names) | **GET** /api/v1/orgs/names | Get versions
[**organizations_v1_list_organizations**](OrganizationsV1Api.md#organizations_v1_list_organizations) | **GET** /api/v1/orgs/list | Get log handler
[**organizations_v1_patch_organization**](OrganizationsV1Api.md#organizations_v1_patch_organization) | **PATCH** /api/v1/orgs/{owner} | 
[**organizations_v1_patch_organization_member**](OrganizationsV1Api.md#organizations_v1_patch_organization_member) | **PATCH** /api/v1/orgs/{owner}/members/{member.user} | 
[**organizations_v1_update_organization**](OrganizationsV1Api.md#organizations_v1_update_organization) | **PUT** /api/v1/orgs/{owner} | 
[**organizations_v1_update_organization_member**](OrganizationsV1Api.md#organizations_v1_update_organization_member) | **PUT** /api/v1/orgs/{owner}/members/{member.user} | 


# **organizations_v1_create_organization**
> V1Organization organizations_v1_create_organization(body)



### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.OrganizationsV1Api(polyaxon_sdk.ApiClient(configuration))
body = polyaxon_sdk.V1Organization() # V1Organization | 

try:
    api_response = api_instance.organizations_v1_create_organization(body)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling OrganizationsV1Api->organizations_v1_create_organization: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1Organization**](V1Organization.md)|  | 

### Return type

[**V1Organization**](V1Organization.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **organizations_v1_create_organization_member**
> V1OrganizationMember organizations_v1_create_organization_member(owner, body)



### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.OrganizationsV1Api(polyaxon_sdk.ApiClient(configuration))
owner = 'owner_example' # str | Owner of the namespace
body = polyaxon_sdk.V1OrganizationMember() # V1OrganizationMember | Organization body

try:
    api_response = api_instance.organizations_v1_create_organization_member(owner, body)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling OrganizationsV1Api->organizations_v1_create_organization_member: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **str**| Owner of the namespace | 
 **body** | [**V1OrganizationMember**](V1OrganizationMember.md)| Organization body | 

### Return type

[**V1OrganizationMember**](V1OrganizationMember.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **organizations_v1_delete_organization**
> organizations_v1_delete_organization(owner)



### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.OrganizationsV1Api(polyaxon_sdk.ApiClient(configuration))
owner = 'owner_example' # str | Owner of the namespace

try:
    api_instance.organizations_v1_delete_organization(owner)
except ApiException as e:
    print("Exception when calling OrganizationsV1Api->organizations_v1_delete_organization: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **str**| Owner of the namespace | 

### Return type

void (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **organizations_v1_delete_organization_member**
> organizations_v1_delete_organization_member(owner, user)



### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.OrganizationsV1Api(polyaxon_sdk.ApiClient(configuration))
owner = 'owner_example' # str | Owner of the namespace
user = 'user_example' # str | Memeber under namesapce

try:
    api_instance.organizations_v1_delete_organization_member(owner, user)
except ApiException as e:
    print("Exception when calling OrganizationsV1Api->organizations_v1_delete_organization_member: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **str**| Owner of the namespace | 
 **user** | **str**| Memeber under namesapce | 

### Return type

void (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **organizations_v1_get_organization**
> V1Organization organizations_v1_get_organization(owner)



### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.OrganizationsV1Api(polyaxon_sdk.ApiClient(configuration))
owner = 'owner_example' # str | Owner of the namespace

try:
    api_response = api_instance.organizations_v1_get_organization(owner)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling OrganizationsV1Api->organizations_v1_get_organization: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **str**| Owner of the namespace | 

### Return type

[**V1Organization**](V1Organization.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **organizations_v1_get_organization_member**
> V1OrganizationMember organizations_v1_get_organization_member(owner, user)



### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.OrganizationsV1Api(polyaxon_sdk.ApiClient(configuration))
owner = 'owner_example' # str | Owner of the namespace
user = 'user_example' # str | Memeber under namesapce

try:
    api_response = api_instance.organizations_v1_get_organization_member(owner, user)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling OrganizationsV1Api->organizations_v1_get_organization_member: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **str**| Owner of the namespace | 
 **user** | **str**| Memeber under namesapce | 

### Return type

[**V1OrganizationMember**](V1OrganizationMember.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **organizations_v1_list_organization_members**
> V1ListOrganizationMembersResponse organizations_v1_list_organization_members(owner, offset=offset, limit=limit, sort=sort, query=query)



### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.OrganizationsV1Api(polyaxon_sdk.ApiClient(configuration))
owner = 'owner_example' # str | Owner of the namespace
offset = 56 # int | Pagination offset. (optional)
limit = 56 # int | Limit size. (optional)
sort = 'sort_example' # str | Sort to order the search. (optional)
query = 'query_example' # str | Query filter the search search. (optional)

try:
    api_response = api_instance.organizations_v1_list_organization_members(owner, offset=offset, limit=limit, sort=sort, query=query)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling OrganizationsV1Api->organizations_v1_list_organization_members: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **str**| Owner of the namespace | 
 **offset** | **int**| Pagination offset. | [optional] 
 **limit** | **int**| Limit size. | [optional] 
 **sort** | **str**| Sort to order the search. | [optional] 
 **query** | **str**| Query filter the search search. | [optional] 

### Return type

[**V1ListOrganizationMembersResponse**](V1ListOrganizationMembersResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **organizations_v1_list_organization_names**
> V1ListOrganizationsResponse organizations_v1_list_organization_names()

Get versions

### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.OrganizationsV1Api(polyaxon_sdk.ApiClient(configuration))

try:
    # Get versions
    api_response = api_instance.organizations_v1_list_organization_names()
    pprint(api_response)
except ApiException as e:
    print("Exception when calling OrganizationsV1Api->organizations_v1_list_organization_names: %s\n" % e)
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**V1ListOrganizationsResponse**](V1ListOrganizationsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **organizations_v1_list_organizations**
> V1ListOrganizationsResponse organizations_v1_list_organizations()

Get log handler

### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.OrganizationsV1Api(polyaxon_sdk.ApiClient(configuration))

try:
    # Get log handler
    api_response = api_instance.organizations_v1_list_organizations()
    pprint(api_response)
except ApiException as e:
    print("Exception when calling OrganizationsV1Api->organizations_v1_list_organizations: %s\n" % e)
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**V1ListOrganizationsResponse**](V1ListOrganizationsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **organizations_v1_patch_organization**
> V1Organization organizations_v1_patch_organization(owner, body)



### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.OrganizationsV1Api(polyaxon_sdk.ApiClient(configuration))
owner = 'owner_example' # str | Owner of the namespace
body = polyaxon_sdk.V1Organization() # V1Organization | Organization body

try:
    api_response = api_instance.organizations_v1_patch_organization(owner, body)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling OrganizationsV1Api->organizations_v1_patch_organization: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **str**| Owner of the namespace | 
 **body** | [**V1Organization**](V1Organization.md)| Organization body | 

### Return type

[**V1Organization**](V1Organization.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **organizations_v1_patch_organization_member**
> V1OrganizationMember organizations_v1_patch_organization_member(owner, member_user, body)



### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.OrganizationsV1Api(polyaxon_sdk.ApiClient(configuration))
owner = 'owner_example' # str | Owner of the namespace
member_user = 'member_user_example' # str | User
body = polyaxon_sdk.V1OrganizationMember() # V1OrganizationMember | Organization body

try:
    api_response = api_instance.organizations_v1_patch_organization_member(owner, member_user, body)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling OrganizationsV1Api->organizations_v1_patch_organization_member: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **str**| Owner of the namespace | 
 **member_user** | **str**| User | 
 **body** | [**V1OrganizationMember**](V1OrganizationMember.md)| Organization body | 

### Return type

[**V1OrganizationMember**](V1OrganizationMember.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **organizations_v1_update_organization**
> V1Organization organizations_v1_update_organization(owner, body)



### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.OrganizationsV1Api(polyaxon_sdk.ApiClient(configuration))
owner = 'owner_example' # str | Owner of the namespace
body = polyaxon_sdk.V1Organization() # V1Organization | Organization body

try:
    api_response = api_instance.organizations_v1_update_organization(owner, body)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling OrganizationsV1Api->organizations_v1_update_organization: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **str**| Owner of the namespace | 
 **body** | [**V1Organization**](V1Organization.md)| Organization body | 

### Return type

[**V1Organization**](V1Organization.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **organizations_v1_update_organization_member**
> V1OrganizationMember organizations_v1_update_organization_member(owner, member_user, body)



### Example
```python
from __future__ import print_function
import time
import polyaxon_sdk
from polyaxon_sdk.rest import ApiException
from pprint import pprint

# Configure API key authorization: ApiKey
configuration = polyaxon_sdk.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = polyaxon_sdk.OrganizationsV1Api(polyaxon_sdk.ApiClient(configuration))
owner = 'owner_example' # str | Owner of the namespace
member_user = 'member_user_example' # str | User
body = polyaxon_sdk.V1OrganizationMember() # V1OrganizationMember | Organization body

try:
    api_response = api_instance.organizations_v1_update_organization_member(owner, member_user, body)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling OrganizationsV1Api->organizations_v1_update_organization_member: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **str**| Owner of the namespace | 
 **member_user** | **str**| User | 
 **body** | [**V1OrganizationMember**](V1OrganizationMember.md)| Organization body | 

### Return type

[**V1OrganizationMember**](V1OrganizationMember.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
