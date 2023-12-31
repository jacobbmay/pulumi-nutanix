// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class NdbTmsCluster extends pulumi.CustomResource {
    /**
     * Get an existing NdbTmsCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NdbTmsClusterState, opts?: pulumi.CustomResourceOptions): NdbTmsCluster {
        return new NdbTmsCluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'nutanix:index/ndbTmsCluster:NdbTmsCluster';

    /**
     * Returns true if the given object is an instance of NdbTmsCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NdbTmsCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NdbTmsCluster.__pulumiType;
    }

    public /*out*/ readonly dateCreated!: pulumi.Output<string>;
    public /*out*/ readonly dateModified!: pulumi.Output<string>;
    public /*out*/ readonly description!: pulumi.Output<string>;
    public /*out*/ readonly logDriveId!: pulumi.Output<string>;
    public /*out*/ readonly logDriveStatus!: pulumi.Output<string>;
    public readonly nxClusterId!: pulumi.Output<string>;
    public /*out*/ readonly ownerId!: pulumi.Output<string>;
    public /*out*/ readonly scheduleId!: pulumi.Output<string>;
    public readonly slaId!: pulumi.Output<string>;
    public /*out*/ readonly source!: pulumi.Output<boolean>;
    public /*out*/ readonly sourceClusters!: pulumi.Output<string[]>;
    public /*out*/ readonly status!: pulumi.Output<string>;
    public readonly timeMachineId!: pulumi.Output<string>;
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a NdbTmsCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NdbTmsClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NdbTmsClusterArgs | NdbTmsClusterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NdbTmsClusterState | undefined;
            resourceInputs["dateCreated"] = state ? state.dateCreated : undefined;
            resourceInputs["dateModified"] = state ? state.dateModified : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["logDriveId"] = state ? state.logDriveId : undefined;
            resourceInputs["logDriveStatus"] = state ? state.logDriveStatus : undefined;
            resourceInputs["nxClusterId"] = state ? state.nxClusterId : undefined;
            resourceInputs["ownerId"] = state ? state.ownerId : undefined;
            resourceInputs["scheduleId"] = state ? state.scheduleId : undefined;
            resourceInputs["slaId"] = state ? state.slaId : undefined;
            resourceInputs["source"] = state ? state.source : undefined;
            resourceInputs["sourceClusters"] = state ? state.sourceClusters : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["timeMachineId"] = state ? state.timeMachineId : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as NdbTmsClusterArgs | undefined;
            if ((!args || args.nxClusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nxClusterId'");
            }
            if ((!args || args.slaId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'slaId'");
            }
            if ((!args || args.timeMachineId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'timeMachineId'");
            }
            resourceInputs["nxClusterId"] = args ? args.nxClusterId : undefined;
            resourceInputs["slaId"] = args ? args.slaId : undefined;
            resourceInputs["timeMachineId"] = args ? args.timeMachineId : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["dateCreated"] = undefined /*out*/;
            resourceInputs["dateModified"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["logDriveId"] = undefined /*out*/;
            resourceInputs["logDriveStatus"] = undefined /*out*/;
            resourceInputs["ownerId"] = undefined /*out*/;
            resourceInputs["scheduleId"] = undefined /*out*/;
            resourceInputs["source"] = undefined /*out*/;
            resourceInputs["sourceClusters"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NdbTmsCluster.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NdbTmsCluster resources.
 */
export interface NdbTmsClusterState {
    dateCreated?: pulumi.Input<string>;
    dateModified?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    logDriveId?: pulumi.Input<string>;
    logDriveStatus?: pulumi.Input<string>;
    nxClusterId?: pulumi.Input<string>;
    ownerId?: pulumi.Input<string>;
    scheduleId?: pulumi.Input<string>;
    slaId?: pulumi.Input<string>;
    source?: pulumi.Input<boolean>;
    sourceClusters?: pulumi.Input<pulumi.Input<string>[]>;
    status?: pulumi.Input<string>;
    timeMachineId?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NdbTmsCluster resource.
 */
export interface NdbTmsClusterArgs {
    nxClusterId: pulumi.Input<string>;
    slaId: pulumi.Input<string>;
    timeMachineId: pulumi.Input<string>;
    type?: pulumi.Input<string>;
}
