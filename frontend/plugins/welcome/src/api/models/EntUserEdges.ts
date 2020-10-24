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
    EntRoom,
    EntRoomFromJSON,
    EntRoomFromJSONTyped,
    EntRoomToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntUserEdges
 */
export interface EntUserEdges {
    /**
     * User2problem holds the value of the user2problem edge.
     * @type {Array<EntProblem>}
     * @memberof EntUserEdges
     */
    user2problem?: Array<EntProblem>;
    /**
     * User2room holds the value of the user2room edge.
     * @type {Array<EntRoom>}
     * @memberof EntUserEdges
     */
    user2room?: Array<EntRoom>;
}

export function EntUserEdgesFromJSON(json: any): EntUserEdges {
    return EntUserEdgesFromJSONTyped(json, false);
}

export function EntUserEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntUserEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'user2problem': !exists(json, 'user2problem') ? undefined : ((json['user2problem'] as Array<any>).map(EntProblemFromJSON)),
        'user2room': !exists(json, 'user2room') ? undefined : ((json['user2room'] as Array<any>).map(EntRoomFromJSON)),
    };
}

export function EntUserEdgesToJSON(value?: EntUserEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'user2problem': value.user2problem === undefined ? undefined : ((value.user2problem as Array<any>).map(EntProblemToJSON)),
        'user2room': value.user2room === undefined ? undefined : ((value.user2room as Array<any>).map(EntRoomToJSON)),
    };
}


