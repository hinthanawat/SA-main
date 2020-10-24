/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Playlist Vidoe
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntProblem,
    EntProblemFromJSON,
    EntProblemFromJSONTyped,
    EntProblemToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntProblemStatusEdges
 */
export interface EntProblemStatusEdges {
    /**
     * Problemstatus2problem holds the value of the problemstatus2problem edge.
     * @type {Array<EntProblem>}
     * @memberof EntProblemStatusEdges
     */
    problemstatus2problem?: Array<EntProblem>;
}

export function EntProblemStatusEdgesFromJSON(json: any): EntProblemStatusEdges {
    return EntProblemStatusEdgesFromJSONTyped(json, false);
}

export function EntProblemStatusEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntProblemStatusEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'problemstatus2problem': !exists(json, 'problemstatus2problem') ? undefined : ((json['problemstatus2problem'] as Array<any>).map(EntProblemFromJSON)),
    };
}

export function EntProblemStatusEdgesToJSON(value?: EntProblemStatusEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'problemstatus2problem': value.problemstatus2problem === undefined ? undefined : ((value.problemstatus2problem as Array<any>).map(EntProblemToJSON)),
    };
}


