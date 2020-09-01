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

/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * The version of the OpenAPI document: 1.1.8-rc4
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface V1OperationBody
 */
export interface V1OperationBody {
    /**
     * 
     * @type {string}
     * @memberof V1OperationBody
     */
    content?: string;
    /**
     * 
     * @type {boolean}
     * @memberof V1OperationBody
     */
    is_managed?: boolean;
    /**
     * 
     * @type {string}
     * @memberof V1OperationBody
     */
    name?: string;
    /**
     * 
     * @type {string}
     * @memberof V1OperationBody
     */
    description?: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof V1OperationBody
     */
    tags?: Array<string>;
}

export function V1OperationBodyFromJSON(json: any): V1OperationBody {
    return V1OperationBodyFromJSONTyped(json, false);
}

export function V1OperationBodyFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1OperationBody {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'content': !exists(json, 'content') ? undefined : json['content'],
        'is_managed': !exists(json, 'is_managed') ? undefined : json['is_managed'],
        'name': !exists(json, 'name') ? undefined : json['name'],
        'description': !exists(json, 'description') ? undefined : json['description'],
        'tags': !exists(json, 'tags') ? undefined : json['tags'],
    };
}

export function V1OperationBodyToJSON(value?: V1OperationBody | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'content': value.content,
        'is_managed': value.is_managed,
        'name': value.name,
        'description': value.description,
        'tags': value.tags,
    };
}


