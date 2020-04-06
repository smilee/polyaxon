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
    define(['ApiClient', 'model/RuntimeError', 'model/V1ListSearchesResponse', 'model/V1Search'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('../model/RuntimeError'), require('../model/V1ListSearchesResponse'), require('../model/V1Search'));
  } else {
    // Browser globals (root is window)
    if (!root.PolyaxonSdk) {
      root.PolyaxonSdk = {};
    }
    root.PolyaxonSdk.SearchesV1Api = factory(root.PolyaxonSdk.ApiClient, root.PolyaxonSdk.RuntimeError, root.PolyaxonSdk.V1ListSearchesResponse, root.PolyaxonSdk.V1Search);
  }
}(this, function(ApiClient, RuntimeError, V1ListSearchesResponse, V1Search) {
  'use strict';

  /**
   * SearchesV1 service.
   * @module api/SearchesV1Api
   * @version 1.0.72
   */

  /**
   * Constructs a new SearchesV1Api. 
   * @alias module:api/SearchesV1Api
   * @class
   * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
   * default to {@link module:ApiClient#instance} if unspecified.
   */
  var exports = function(apiClient) {
    this.apiClient = apiClient || ApiClient.instance;


    /**
     * Callback function to receive the result of the searchesV1CreateSearch operation.
     * @callback module:api/SearchesV1Api~searchesV1CreateSearchCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1Search} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Create search
     * @param {String} owner Owner of the namespace
     * @param {module:model/V1Search} body Search body
     * @param {module:api/SearchesV1Api~searchesV1CreateSearchCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1Search}
     */
    this.searchesV1CreateSearch = function(owner, body, callback) {
      var postBody = body;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling searchesV1CreateSearch");
      }

      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling searchesV1CreateSearch");
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
      var returnType = V1Search;

      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/searches', 'POST',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the searchesV1DeleteSearch operation.
     * @callback module:api/SearchesV1Api~searchesV1DeleteSearchCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Delete search
     * @param {String} owner Owner of the namespace
     * @param {String} uuid Uuid identifier of the entity
     * @param {module:api/SearchesV1Api~searchesV1DeleteSearchCallback} callback The callback function, accepting three arguments: error, data, response
     */
    this.searchesV1DeleteSearch = function(owner, uuid, callback) {
      var postBody = null;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling searchesV1DeleteSearch");
      }

      // verify the required parameter 'uuid' is set
      if (uuid === undefined || uuid === null) {
        throw new Error("Missing the required parameter 'uuid' when calling searchesV1DeleteSearch");
      }


      var pathParams = {
        'owner': owner,
        'uuid': uuid
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
        '/api/v1/orgs/{owner}/searches/{uuid}', 'DELETE',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the searchesV1GetSearch operation.
     * @callback module:api/SearchesV1Api~searchesV1GetSearchCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1Search} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get search
     * @param {String} owner Owner of the namespace
     * @param {String} uuid Uuid identifier of the entity
     * @param {module:api/SearchesV1Api~searchesV1GetSearchCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1Search}
     */
    this.searchesV1GetSearch = function(owner, uuid, callback) {
      var postBody = null;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling searchesV1GetSearch");
      }

      // verify the required parameter 'uuid' is set
      if (uuid === undefined || uuid === null) {
        throw new Error("Missing the required parameter 'uuid' when calling searchesV1GetSearch");
      }


      var pathParams = {
        'owner': owner,
        'uuid': uuid
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
      var returnType = V1Search;

      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/searches/{uuid}', 'GET',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the searchesV1ListSearchNames operation.
     * @callback module:api/SearchesV1Api~searchesV1ListSearchNamesCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1ListSearchesResponse} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * List search names
     * @param {String} owner Owner of the namespace
     * @param {Object} opts Optional parameters
     * @param {Number} opts.offset Pagination offset.
     * @param {Number} opts.limit Limit size.
     * @param {String} opts.sort Sort to order the search.
     * @param {String} opts.query Query filter the search search.
     * @param {module:api/SearchesV1Api~searchesV1ListSearchNamesCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1ListSearchesResponse}
     */
    this.searchesV1ListSearchNames = function(owner, opts, callback) {
      opts = opts || {};
      var postBody = null;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling searchesV1ListSearchNames");
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
      var returnType = V1ListSearchesResponse;

      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/searches/names', 'GET',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the searchesV1ListSearches operation.
     * @callback module:api/SearchesV1Api~searchesV1ListSearchesCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1ListSearchesResponse} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * List searches
     * @param {String} owner Owner of the namespace
     * @param {Object} opts Optional parameters
     * @param {Number} opts.offset Pagination offset.
     * @param {Number} opts.limit Limit size.
     * @param {String} opts.sort Sort to order the search.
     * @param {String} opts.query Query filter the search search.
     * @param {module:api/SearchesV1Api~searchesV1ListSearchesCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1ListSearchesResponse}
     */
    this.searchesV1ListSearches = function(owner, opts, callback) {
      opts = opts || {};
      var postBody = null;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling searchesV1ListSearches");
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
      var returnType = V1ListSearchesResponse;

      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/searches', 'GET',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the searchesV1PatchSearch operation.
     * @callback module:api/SearchesV1Api~searchesV1PatchSearchCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1Search} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Patch search
     * @param {String} owner Owner of the namespace
     * @param {String} search_uuid UUID
     * @param {module:model/V1Search} body Search body
     * @param {module:api/SearchesV1Api~searchesV1PatchSearchCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1Search}
     */
    this.searchesV1PatchSearch = function(owner, search_uuid, body, callback) {
      var postBody = body;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling searchesV1PatchSearch");
      }

      // verify the required parameter 'search_uuid' is set
      if (search_uuid === undefined || search_uuid === null) {
        throw new Error("Missing the required parameter 'search_uuid' when calling searchesV1PatchSearch");
      }

      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling searchesV1PatchSearch");
      }


      var pathParams = {
        'owner': owner,
        'search.uuid': search_uuid
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
      var returnType = V1Search;

      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/searches/{search.uuid}', 'PATCH',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the searchesV1UpdateSearch operation.
     * @callback module:api/SearchesV1Api~searchesV1UpdateSearchCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1Search} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Update search
     * @param {String} owner Owner of the namespace
     * @param {String} search_uuid UUID
     * @param {module:model/V1Search} body Search body
     * @param {module:api/SearchesV1Api~searchesV1UpdateSearchCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1Search}
     */
    this.searchesV1UpdateSearch = function(owner, search_uuid, body, callback) {
      var postBody = body;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling searchesV1UpdateSearch");
      }

      // verify the required parameter 'search_uuid' is set
      if (search_uuid === undefined || search_uuid === null) {
        throw new Error("Missing the required parameter 'search_uuid' when calling searchesV1UpdateSearch");
      }

      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling searchesV1UpdateSearch");
      }


      var pathParams = {
        'owner': owner,
        'search.uuid': search_uuid
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
      var returnType = V1Search;

      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/searches/{search.uuid}', 'PUT',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }
  };

  return exports;
}));