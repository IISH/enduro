/* tslint:disable */
/* eslint-disable */
/**
 * Enduro API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * StoredPackage describes a collection retrieved by the service. (default view)
 * @export
 * @interface EnduroStoredCollectionResponseBody
 */
export interface EnduroStoredCollectionResponseBody {
    /**
     * Identifier of Archivematica AIP
     * @type {string}
     * @memberof EnduroStoredCollectionResponseBody
     */
    aipId?: string;
    /**
     * Completion datetime
     * @type {Date}
     * @memberof EnduroStoredCollectionResponseBody
     */
    completedAt?: Date;
    /**
     * Creation datetime
     * @type {Date}
     * @memberof EnduroStoredCollectionResponseBody
     */
    createdAt: Date;
    /**
     * Identifier of collection
     * @type {number}
     * @memberof EnduroStoredCollectionResponseBody
     */
    id: number;
    /**
     * Name of the collection
     * @type {string}
     * @memberof EnduroStoredCollectionResponseBody
     */
    name?: string;
    /**
     * Identifier provided by the client
     * @type {string}
     * @memberof EnduroStoredCollectionResponseBody
     */
    originalId?: string;
    /**
     * Identifier of Archivematica pipeline
     * @type {string}
     * @memberof EnduroStoredCollectionResponseBody
     */
    pipelineId?: string;
    /**
     * Identifier of latest processing workflow run
     * @type {string}
     * @memberof EnduroStoredCollectionResponseBody
     */
    runId?: string;
    /**
     * Start datetime
     * @type {Date}
     * @memberof EnduroStoredCollectionResponseBody
     */
    startedAt?: Date;
    /**
     * Status of the collection
     * @type {string}
     * @memberof EnduroStoredCollectionResponseBody
     */
    status: EnduroStoredCollectionResponseBodyStatusEnum;
    /**
     * Identifier of Archivematica transfer
     * @type {string}
     * @memberof EnduroStoredCollectionResponseBody
     */
    transferId?: string;
    /**
     * Identifier of processing workflow
     * @type {string}
     * @memberof EnduroStoredCollectionResponseBody
     */
    workflowId?: string;
}

export function EnduroStoredCollectionResponseBodyFromJSON(json: any): EnduroStoredCollectionResponseBody {
    return EnduroStoredCollectionResponseBodyFromJSONTyped(json, false);
}

export function EnduroStoredCollectionResponseBodyFromJSONTyped(json: any, ignoreDiscriminator: boolean): EnduroStoredCollectionResponseBody {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'aipId': !exists(json, 'aip_id') ? undefined : json['aip_id'],
        'completedAt': !exists(json, 'completed_at') ? undefined : (new Date(json['completed_at'])),
        'createdAt': (new Date(json['created_at'])),
        'id': json['id'],
        'name': !exists(json, 'name') ? undefined : json['name'],
        'originalId': !exists(json, 'original_id') ? undefined : json['original_id'],
        'pipelineId': !exists(json, 'pipeline_id') ? undefined : json['pipeline_id'],
        'runId': !exists(json, 'run_id') ? undefined : json['run_id'],
        'startedAt': !exists(json, 'started_at') ? undefined : (new Date(json['started_at'])),
        'status': json['status'],
        'transferId': !exists(json, 'transfer_id') ? undefined : json['transfer_id'],
        'workflowId': !exists(json, 'workflow_id') ? undefined : json['workflow_id'],
    };
}

export function EnduroStoredCollectionResponseBodyToJSON(value?: EnduroStoredCollectionResponseBody | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'aip_id': value.aipId,
        'completed_at': value.completedAt === undefined ? undefined : (value.completedAt.toISOString()),
        'created_at': (value.createdAt.toISOString()),
        'id': value.id,
        'name': value.name,
        'original_id': value.originalId,
        'pipeline_id': value.pipelineId,
        'run_id': value.runId,
        'started_at': value.startedAt === undefined ? undefined : (value.startedAt.toISOString()),
        'status': value.status,
        'transfer_id': value.transferId,
        'workflow_id': value.workflowId,
    };
}

/**
* @export
* @enum {string}
*/
export enum EnduroStoredCollectionResponseBodyStatusEnum {
    New = 'new',
    InProgress = 'in progress',
    Done = 'done',
    Error = 'error',
    Unknown = 'unknown',
    Queued = 'queued'
}


