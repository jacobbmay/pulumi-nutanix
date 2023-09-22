// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Provides a resource to create maintenance window based on the input parameters.
 *
 * ## Example Usage
 * ### resource to create weekly maintenance window
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nutanix from "@jacobbmay/nutanix";
 *
 * const acctest_managed = new nutanix.NdbMaintenanceWindow("acctest-managed", {
 *     dayOfWeek: "TUESDAY",
 *     description: "desc",
 *     duration: 3,
 *     recurrence: "WEEKLY",
 *     startTime: "17:04:47",
 * });
 * ```
 * ### resource to create monthly maintenance window
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nutanix from "@jacobbmay/nutanix";
 *
 * const acctest_managed = new nutanix.NdbMaintenanceWindow("acctest-managed", {
 *     dayOfWeek: "TUESDAY",
 *     description: "description",
 *     duration: 2,
 *     recurrence: "MONTHLY",
 *     startTime: "17:04:47",
 *     weekOfMonth: 4,
 * });
 * ```
 */
export class NdbMaintenanceWindow extends pulumi.CustomResource {
    /**
     * Get an existing NdbMaintenanceWindow resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NdbMaintenanceWindowState, opts?: pulumi.CustomResourceOptions): NdbMaintenanceWindow {
        return new NdbMaintenanceWindow(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'nutanix:index/ndbMaintenanceWindow:NdbMaintenanceWindow';

    /**
     * Returns true if the given object is an instance of NdbMaintenanceWindow.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NdbMaintenanceWindow {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NdbMaintenanceWindow.__pulumiType;
    }

    /**
     * access level
     */
    public /*out*/ readonly accessLevel!: pulumi.Output<string>;
    /**
     * created date of maintenance window
     */
    public /*out*/ readonly dateCreated!: pulumi.Output<string>;
    /**
     * modified date of maintenance window
     */
    public /*out*/ readonly dateModified!: pulumi.Output<string>;
    /**
     * Day of the week to trigger maintenance window. Supports [ MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY, SUNDAY ]
     */
    public readonly dayOfWeek!: pulumi.Output<string | undefined>;
    /**
     * Description for maintenance window
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * duration in hours. Default is 2
     */
    public readonly duration!: pulumi.Output<number | undefined>;
    /**
     * entity task association for maintenance window
     */
    public /*out*/ readonly entityTaskAssocs!: pulumi.Output<outputs.NdbMaintenanceWindowEntityTaskAssoc[]>;
    /**
     * Name for the maintenance window.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * next run time for maintenance window to trigger
     */
    public /*out*/ readonly nextRunTime!: pulumi.Output<string>;
    /**
     * owner id of maintenance window
     */
    public /*out*/ readonly ownerId!: pulumi.Output<string>;
    /**
     * properties of maintenance window
     */
    public /*out*/ readonly properties!: pulumi.Output<outputs.NdbMaintenanceWindowProperty[]>;
    /**
     * Supported values [ MONTHLY, WEEKLY ]
     */
    public readonly recurrence!: pulumi.Output<string>;
    /**
     * schedule of maintenance window
     */
    public /*out*/ readonly schedules!: pulumi.Output<outputs.NdbMaintenanceWindowSchedule[]>;
    /**
     * start time for maintenance window to trigger
     */
    public readonly startTime!: pulumi.Output<string>;
    /**
     * status of maintennace window
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * tags of maintenance window
     */
    public readonly tags!: pulumi.Output<outputs.NdbMaintenanceWindowTag[]>;
    /**
     * timezone . Default is Asia/Calcutta .
     */
    public readonly timezone!: pulumi.Output<string | undefined>;
    /**
     * week of the month. Supports [1, 2, 3, 4] .
     */
    public readonly weekOfMonth!: pulumi.Output<number | undefined>;

    /**
     * Create a NdbMaintenanceWindow resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NdbMaintenanceWindowArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NdbMaintenanceWindowArgs | NdbMaintenanceWindowState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NdbMaintenanceWindowState | undefined;
            resourceInputs["accessLevel"] = state ? state.accessLevel : undefined;
            resourceInputs["dateCreated"] = state ? state.dateCreated : undefined;
            resourceInputs["dateModified"] = state ? state.dateModified : undefined;
            resourceInputs["dayOfWeek"] = state ? state.dayOfWeek : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["duration"] = state ? state.duration : undefined;
            resourceInputs["entityTaskAssocs"] = state ? state.entityTaskAssocs : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nextRunTime"] = state ? state.nextRunTime : undefined;
            resourceInputs["ownerId"] = state ? state.ownerId : undefined;
            resourceInputs["properties"] = state ? state.properties : undefined;
            resourceInputs["recurrence"] = state ? state.recurrence : undefined;
            resourceInputs["schedules"] = state ? state.schedules : undefined;
            resourceInputs["startTime"] = state ? state.startTime : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["timezone"] = state ? state.timezone : undefined;
            resourceInputs["weekOfMonth"] = state ? state.weekOfMonth : undefined;
        } else {
            const args = argsOrState as NdbMaintenanceWindowArgs | undefined;
            if ((!args || args.recurrence === undefined) && !opts.urn) {
                throw new Error("Missing required property 'recurrence'");
            }
            if ((!args || args.startTime === undefined) && !opts.urn) {
                throw new Error("Missing required property 'startTime'");
            }
            resourceInputs["dayOfWeek"] = args ? args.dayOfWeek : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["duration"] = args ? args.duration : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["recurrence"] = args ? args.recurrence : undefined;
            resourceInputs["startTime"] = args ? args.startTime : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["timezone"] = args ? args.timezone : undefined;
            resourceInputs["weekOfMonth"] = args ? args.weekOfMonth : undefined;
            resourceInputs["accessLevel"] = undefined /*out*/;
            resourceInputs["dateCreated"] = undefined /*out*/;
            resourceInputs["dateModified"] = undefined /*out*/;
            resourceInputs["entityTaskAssocs"] = undefined /*out*/;
            resourceInputs["nextRunTime"] = undefined /*out*/;
            resourceInputs["ownerId"] = undefined /*out*/;
            resourceInputs["properties"] = undefined /*out*/;
            resourceInputs["schedules"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NdbMaintenanceWindow.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NdbMaintenanceWindow resources.
 */
export interface NdbMaintenanceWindowState {
    /**
     * access level
     */
    accessLevel?: pulumi.Input<string>;
    /**
     * created date of maintenance window
     */
    dateCreated?: pulumi.Input<string>;
    /**
     * modified date of maintenance window
     */
    dateModified?: pulumi.Input<string>;
    /**
     * Day of the week to trigger maintenance window. Supports [ MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY, SUNDAY ]
     */
    dayOfWeek?: pulumi.Input<string>;
    /**
     * Description for maintenance window
     */
    description?: pulumi.Input<string>;
    /**
     * duration in hours. Default is 2
     */
    duration?: pulumi.Input<number>;
    /**
     * entity task association for maintenance window
     */
    entityTaskAssocs?: pulumi.Input<pulumi.Input<inputs.NdbMaintenanceWindowEntityTaskAssoc>[]>;
    /**
     * Name for the maintenance window.
     */
    name?: pulumi.Input<string>;
    /**
     * next run time for maintenance window to trigger
     */
    nextRunTime?: pulumi.Input<string>;
    /**
     * owner id of maintenance window
     */
    ownerId?: pulumi.Input<string>;
    /**
     * properties of maintenance window
     */
    properties?: pulumi.Input<pulumi.Input<inputs.NdbMaintenanceWindowProperty>[]>;
    /**
     * Supported values [ MONTHLY, WEEKLY ]
     */
    recurrence?: pulumi.Input<string>;
    /**
     * schedule of maintenance window
     */
    schedules?: pulumi.Input<pulumi.Input<inputs.NdbMaintenanceWindowSchedule>[]>;
    /**
     * start time for maintenance window to trigger
     */
    startTime?: pulumi.Input<string>;
    /**
     * status of maintennace window
     */
    status?: pulumi.Input<string>;
    /**
     * tags of maintenance window
     */
    tags?: pulumi.Input<pulumi.Input<inputs.NdbMaintenanceWindowTag>[]>;
    /**
     * timezone . Default is Asia/Calcutta .
     */
    timezone?: pulumi.Input<string>;
    /**
     * week of the month. Supports [1, 2, 3, 4] .
     */
    weekOfMonth?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a NdbMaintenanceWindow resource.
 */
export interface NdbMaintenanceWindowArgs {
    /**
     * Day of the week to trigger maintenance window. Supports [ MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY, SUNDAY ]
     */
    dayOfWeek?: pulumi.Input<string>;
    /**
     * Description for maintenance window
     */
    description?: pulumi.Input<string>;
    /**
     * duration in hours. Default is 2
     */
    duration?: pulumi.Input<number>;
    /**
     * Name for the maintenance window.
     */
    name?: pulumi.Input<string>;
    /**
     * Supported values [ MONTHLY, WEEKLY ]
     */
    recurrence: pulumi.Input<string>;
    /**
     * start time for maintenance window to trigger
     */
    startTime: pulumi.Input<string>;
    /**
     * tags of maintenance window
     */
    tags?: pulumi.Input<pulumi.Input<inputs.NdbMaintenanceWindowTag>[]>;
    /**
     * timezone . Default is Asia/Calcutta .
     */
    timezone?: pulumi.Input<string>;
    /**
     * week of the month. Supports [1, 2, 3, 4] .
     */
    weekOfMonth?: pulumi.Input<number>;
}