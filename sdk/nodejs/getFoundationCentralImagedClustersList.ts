// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getFoundationCentralImagedClustersList(args?: GetFoundationCentralImagedClustersListArgs, opts?: pulumi.InvokeOptions): Promise<GetFoundationCentralImagedClustersListResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nutanix:index/getFoundationCentralImagedClustersList:getFoundationCentralImagedClustersList", {
        "filters": args.filters,
        "length": args.length,
        "offset": args.offset,
    }, opts);
}

/**
 * A collection of arguments for invoking getFoundationCentralImagedClustersList.
 */
export interface GetFoundationCentralImagedClustersListArgs {
    filters?: inputs.GetFoundationCentralImagedClustersListFilters;
    length?: number;
    offset?: number;
}

/**
 * A collection of values returned by getFoundationCentralImagedClustersList.
 */
export interface GetFoundationCentralImagedClustersListResult {
    readonly filters?: outputs.GetFoundationCentralImagedClustersListFilters;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly imagedClusters: outputs.GetFoundationCentralImagedClustersListImagedCluster[];
    readonly length?: number;
    readonly metadatas: outputs.GetFoundationCentralImagedClustersListMetadata[];
    readonly offset?: number;
}
export function getFoundationCentralImagedClustersListOutput(args?: GetFoundationCentralImagedClustersListOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFoundationCentralImagedClustersListResult> {
    return pulumi.output(args).apply((a: any) => getFoundationCentralImagedClustersList(a, opts))
}

/**
 * A collection of arguments for invoking getFoundationCentralImagedClustersList.
 */
export interface GetFoundationCentralImagedClustersListOutputArgs {
    filters?: pulumi.Input<inputs.GetFoundationCentralImagedClustersListFiltersArgs>;
    length?: pulumi.Input<number>;
    offset?: pulumi.Input<number>;
}
