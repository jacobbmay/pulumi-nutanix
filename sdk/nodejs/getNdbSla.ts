// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Describes a SLA in Nutanix Database Service
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nutanix from "@pulumi/nutanix";
 *
 * const sla1 = nutanix.getNdbSla({
 *     slaName: "test-sla",
 * });
 * export const sla = sla1;
 * ```
 */
export function getNdbSla(args?: GetNdbSlaArgs, opts?: pulumi.InvokeOptions): Promise<GetNdbSlaResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nutanix:index/getNdbSla:getNdbSla", {
        "slaId": args.slaId,
        "slaName": args.slaName,
    }, opts);
}

/**
 * A collection of arguments for invoking getNdbSla.
 */
export interface GetNdbSlaArgs {
    /**
     * and `slaName` are mutually exclusive.
     */
    slaId?: string;
    /**
     * SLA Name for query
     */
    slaName?: string;
}

/**
 * A collection of values returned by getNdbSla.
 */
export interface GetNdbSlaResult {
    /**
     * - continuous retention of logs limit
     */
    readonly continuousRetention: number;
    /**
     * - Current active frequency
     */
    readonly currentActiveFrequency: string;
    /**
     * - Daily snapshots retention limit
     */
    readonly dailyRetention: number;
    /**
     * - creation date
     */
    readonly dateCreated: string;
    /**
     * - last modified
     */
    readonly dateModified: string;
    /**
     * - description of sla
     */
    readonly description: string;
    /**
     * - id of sla
     */
    readonly id: string;
    /**
     * - Monthly snapshots retention limit
     */
    readonly monthlyRetention: number;
    /**
     * - sla name
     */
    readonly name: string;
    /**
     * - owner ID
     */
    readonly ownerId: string;
    /**
     * - If point in time recovery enabled
     */
    readonly pitrEnabled: boolean;
    /**
     * - Daily snapshots retention limit
     */
    readonly quartelyRetention: number;
    /**
     * - Reference count
     */
    readonly referenceCount: number;
    readonly slaId?: string;
    readonly slaName?: string;
    /**
     * - if system sla
     */
    readonly systemSla: boolean;
    /**
     * - unique name
     */
    readonly uniqueName: string;
    /**
     * - weeky snapshots retention limit
     */
    readonly weeklyRetention: number;
    /**
     * - Yearly snapshots retention limit
     */
    readonly yearlyRetention: number;
}
/**
 * Describes a SLA in Nutanix Database Service
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nutanix from "@pulumi/nutanix";
 *
 * const sla1 = nutanix.getNdbSla({
 *     slaName: "test-sla",
 * });
 * export const sla = sla1;
 * ```
 */
export function getNdbSlaOutput(args?: GetNdbSlaOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNdbSlaResult> {
    return pulumi.output(args).apply((a: any) => getNdbSla(a, opts))
}

/**
 * A collection of arguments for invoking getNdbSla.
 */
export interface GetNdbSlaOutputArgs {
    /**
     * and `slaName` are mutually exclusive.
     */
    slaId?: pulumi.Input<string>;
    /**
     * SLA Name for query
     */
    slaName?: pulumi.Input<string>;
}