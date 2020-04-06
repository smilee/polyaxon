// Copyright 2018-2020 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * OpenAPI spec version: 1.0.72
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.4.10
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient', 'model/RuntimeError', 'model/V1ListOrganizationMembersResponse', 'model/V1ListOrganizationsResponse', 'model/V1Organization', 'model/V1OrganizationMember'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('../model/RuntimeError'), require('../model/V1ListOrganizationMembersResponse'), require('../model/V1ListOrganizationsResponse'), require('../model/V1Organization'), require('../model/V1OrganizationMember'));
  } else {
    // Browser globals (root is window)
    if (!root.PolyaxonSdk) {
      root.PolyaxonSdk = {};
    }
    root.PolyaxonSdk.OrganizationsV1Api = factory(root.PolyaxonSdk.ApiClient, root.PolyaxonSdk.RuntimeError, root.PolyaxonSdk.V1ListOrganizationMembersResponse, root.PolyaxonSdk.V1ListOrganizationsResponse, root.PolyaxonSdk.V1Organization, root.PolyaxonSdk.V1OrganizationMember);
  }
}(this, function(ApiClient, RuntimeError, V1ListOrganizationMembersResponse, V1ListOrganizationsResponse, V1Organization, V1OrganizationMember) {
  'use strict';

  /**
   * OrganizationsV1 service.
   * @module api/OrganizationsV1Api
   * @version 1.0.72
   */

  /**
   * Constructs a new OrganizationsV1Api. 
   * @alias module:api/OrganizationsV1Api
   * @class
   * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
   * default to {@link module:ApiClient#instance} if unspecified.
   */
  var exports = function(apiClient) {
    this.apiClient = apiClient || ApiClient.instance;


    /**
     * Callback function to receive the result of the organizationsV1CreateOrganization operation.
     * @callback module:api/OrganizationsV1Api~organizationsV1CreateOrganizationCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1Organization} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * @param {module:model/V1Organization} body 
     * @param {module:api/OrganizationsV1Api~organizationsV1CreateOrganizationCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1Organization}
     */
    this.organizationsV1CreateOrganization = function(body, callback) {
      var postBody = body;

      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling organizationsV1CreateOrganization");
      }


      var pathParams = {
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['ApiKey'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = V1Organization;

      return this.apiClient.callApi(
        '/api/v1/orgs/create', 'POST',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the organizationsV1CreateOrganizationMember operation.
     * @callback module:api/OrganizationsV1Api~organizationsV1CreateOrganizationMemberCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1OrganizationMember} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * @param {String} owner Owner of the namespace
     * @param {module:model/V1OrganizationMember} body Organization body
     * @param {module:api/OrganizationsV1Api~organizationsV1CreateOrganizationMemberCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1OrganizationMember}
     */
    this.organizationsV1CreateOrganizationMember = function(owner, body, callback) {
      var postBody = body;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling organizationsV1CreateOrganizationMember");
      }

      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling organizationsV1CreateOrganizationMember");
      }


      var pathParams = {
        'owner': owner
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['ApiKey'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = V1OrganizationMember;

      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/members', 'POST',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the organizationsV1DeleteOrganization operation.
     * @callback module:api/OrganizationsV1Api~organizationsV1DeleteOrganizationCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * @param {String} owner Owner of the namespace
     * @param {module:api/OrganizationsV1Api~organizationsV1DeleteOrganizationCallback} callback The callback function, accepting three arguments: error, data, response
     */
    this.organizationsV1DeleteOrganization = function(owner, callback) {
      var postBody = null;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling organizationsV1DeleteOrganization");
      }


      var pathParams = {
        'owner': owner
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['ApiKey'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = null;

      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}', 'DELETE',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the organizationsV1DeleteOrganizationMember operation.
     * @callback module:api/OrganizationsV1Api~organizationsV1DeleteOrganizationMemberCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * @param {String} owner Owner of the namespace
     * @param {String} user Memeber under namesapce
     * @param {module:api/OrganizationsV1Api~organizationsV1DeleteOrganizationMemberCallback} callback The callback function, accepting three arguments: error, data, response
     */
    this.organizationsV1DeleteOrganizationMember = function(owner, user, callback) {
      var postBody = null;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling organizationsV1DeleteOrganizationMember");
      }

      // verify the required parameter 'user' is set
      if (user === undefined || user === null) {
        throw new Error("Missing the required parameter 'user' when calling organizationsV1DeleteOrganizationMember");
      }


      var pathParams = {
        'owner': owner,
        'user': user
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['ApiKey'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = null;

      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/members/{user}', 'DELETE',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the organizationsV1GetOrganization operation.
     * @callback module:api/OrganizationsV1Api~organizationsV1GetOrganizationCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1Organization} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * @param {String} owner Owner of the namespace
     * @param {module:api/OrganizationsV1Api~organizationsV1GetOrganizationCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1Organization}
     */
    this.organizationsV1GetOrganization = function(owner, callback) {
      var postBody = null;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling organizationsV1GetOrganization");
      }


      var pathParams = {
        'owner': owner
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['ApiKey'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = V1Organization;

      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}', 'GET',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the organizationsV1GetOrganizationMember operation.
     * @callback module:api/OrganizationsV1Api~organizationsV1GetOrganizationMemberCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1OrganizationMember} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * @param {String} owner Owner of the namespace
     * @param {String} user Memeber under namesapce
     * @param {module:api/OrganizationsV1Api~organizationsV1GetOrganizationMemberCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1OrganizationMember}
     */
    this.organizationsV1GetOrganizationMember = function(owner, user, callback) {
      var postBody = null;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling organizationsV1GetOrganizationMember");
      }

      // verify the required parameter 'user' is set
      if (user === undefined || user === null) {
        throw new Error("Missing the required parameter 'user' when calling organizationsV1GetOrganizationMember");
      }


      var pathParams = {
        'owner': owner,
        'user': user
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['ApiKey'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = V1OrganizationMember;

      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/members/{user}', 'GET',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the organizationsV1ListOrganizationMembers operation.
     * @callback module:api/OrganizationsV1Api~organizationsV1ListOrganizationMembersCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1ListOrganizationMembersResponse} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * @param {String} owner Owner of the namespace
     * @param {Object} opts Optional parameters
     * @param {Number} opts.offset Pagination offset.
     * @param {Number} opts.limit Limit size.
     * @param {String} opts.sort Sort to order the search.
     * @param {String} opts.query Query filter the search search.
     * @param {module:api/OrganizationsV1Api~organizationsV1ListOrganizationMembersCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1ListOrganizationMembersResponse}
     */
    this.organizationsV1ListOrganizationMembers = function(owner, opts, callback) {
      opts = opts || {};
      var postBody = null;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling organizationsV1ListOrganizationMembers");
      }


      var pathParams = {
        'owner': owner
      };
      var queryParams = {
        'offset': opts['offset'],
        'limit': opts['limit'],
        'sort': opts['sort'],
        'query': opts['query'],
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['ApiKey'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = V1ListOrganizationMembersResponse;

      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/members', 'GET',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the organizationsV1ListOrganizationNames operation.
     * @callback module:api/OrganizationsV1Api~organizationsV1ListOrganizationNamesCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1ListOrganizationsResponse} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get versions
     * @param {module:api/OrganizationsV1Api~organizationsV1ListOrganizationNamesCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1ListOrganizationsResponse}
     */
    this.organizationsV1ListOrganizationNames = function(callback) {
      var postBody = null;


      var pathParams = {
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['ApiKey'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = V1ListOrganizationsResponse;

      return this.apiClient.callApi(
        '/api/v1/orgs/names', 'GET',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the organizationsV1ListOrganizations operation.
     * @callback module:api/OrganizationsV1Api~organizationsV1ListOrganizationsCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1ListOrganizationsResponse} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get log handler
     * @param {module:api/OrganizationsV1Api~organizationsV1ListOrganizationsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1ListOrganizationsResponse}
     */
    this.organizationsV1ListOrganizations = function(callback) {
      var postBody = null;


      var pathParams = {
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['ApiKey'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = V1ListOrganizationsResponse;

      return this.apiClient.callApi(
        '/api/v1/orgs/list', 'GET',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the organizationsV1PatchOrganization operation.
     * @callback module:api/OrganizationsV1Api~organizationsV1PatchOrganizationCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1Organization} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * @param {String} owner Owner of the namespace
     * @param {module:model/V1Organization} body Organization body
     * @param {module:api/OrganizationsV1Api~organizationsV1PatchOrganizationCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1Organization}
     */
    this.organizationsV1PatchOrganization = function(owner, body, callback) {
      var postBody = body;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling organizationsV1PatchOrganization");
      }

      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling organizationsV1PatchOrganization");
      }


      var pathParams = {
        'owner': owner
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['ApiKey'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = V1Organization;

      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}', 'PATCH',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the organizationsV1PatchOrganizationMember operation.
     * @callback module:api/OrganizationsV1Api~organizationsV1PatchOrganizationMemberCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1OrganizationMember} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * @param {String} owner Owner of the namespace
     * @param {String} member_user User
     * @param {module:model/V1OrganizationMember} body Organization body
     * @param {module:api/OrganizationsV1Api~organizationsV1PatchOrganizationMemberCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1OrganizationMember}
     */
    this.organizationsV1PatchOrganizationMember = function(owner, member_user, body, callback) {
      var postBody = body;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling organizationsV1PatchOrganizationMember");
      }

      // verify the required parameter 'member_user' is set
      if (member_user === undefined || member_user === null) {
        throw new Error("Missing the required parameter 'member_user' when calling organizationsV1PatchOrganizationMember");
      }

      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling organizationsV1PatchOrganizationMember");
      }


      var pathParams = {
        'owner': owner,
        'member.user': member_user
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['ApiKey'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = V1OrganizationMember;

      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/members/{member.user}', 'PATCH',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the organizationsV1UpdateOrganization operation.
     * @callback module:api/OrganizationsV1Api~organizationsV1UpdateOrganizationCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1Organization} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * @param {String} owner Owner of the namespace
     * @param {module:model/V1Organization} body Organization body
     * @param {module:api/OrganizationsV1Api~organizationsV1UpdateOrganizationCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1Organization}
     */
    this.organizationsV1UpdateOrganization = function(owner, body, callback) {
      var postBody = body;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling organizationsV1UpdateOrganization");
      }

      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling organizationsV1UpdateOrganization");
      }


      var pathParams = {
        'owner': owner
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['ApiKey'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = V1Organization;

      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}', 'PUT',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the organizationsV1UpdateOrganizationMember operation.
     * @callback module:api/OrganizationsV1Api~organizationsV1UpdateOrganizationMemberCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1OrganizationMember} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * @param {String} owner Owner of the namespace
     * @param {String} member_user User
     * @param {module:model/V1OrganizationMember} body Organization body
     * @param {module:api/OrganizationsV1Api~organizationsV1UpdateOrganizationMemberCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1OrganizationMember}
     */
    this.organizationsV1UpdateOrganizationMember = function(owner, member_user, body, callback) {
      var postBody = body;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling organizationsV1UpdateOrganizationMember");
      }

      // verify the required parameter 'member_user' is set
      if (member_user === undefined || member_user === null) {
        throw new Error("Missing the required parameter 'member_user' when calling organizationsV1UpdateOrganizationMember");
      }

      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling organizationsV1UpdateOrganizationMember");
      }


      var pathParams = {
        'owner': owner,
        'member.user': member_user
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['ApiKey'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = V1OrganizationMember;

      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/members/{member.user}', 'PUT',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }
  };

  return exports;
}));