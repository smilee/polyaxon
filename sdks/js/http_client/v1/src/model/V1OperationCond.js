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
 *
 */

import ApiClient from '../ApiClient';
import V1IoCond from './V1IoCond';
import V1StatusCond from './V1StatusCond';

/**
 * The V1OperationCond model module.
 * @module model/V1OperationCond
 * @version 1.1.8-rc4
 */
class V1OperationCond {
    /**
     * Constructs a new <code>V1OperationCond</code>.
     * @alias module:model/V1OperationCond
     */
    constructor() { 
        
        V1OperationCond.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1OperationCond</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1OperationCond} obj Optional instance to populate.
     * @return {module:model/V1OperationCond} The populated <code>V1OperationCond</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1OperationCond();

            if (data.hasOwnProperty('io_conidtion')) {
                obj['io_conidtion'] = V1IoCond.constructFromObject(data['io_conidtion']);
            }
            if (data.hasOwnProperty('status_condition')) {
                obj['status_condition'] = V1StatusCond.constructFromObject(data['status_condition']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/V1IoCond} io_conidtion
 */
V1OperationCond.prototype['io_conidtion'] = undefined;

/**
 * @member {module:model/V1StatusCond} status_condition
 */
V1OperationCond.prototype['status_condition'] = undefined;






export default V1OperationCond;

