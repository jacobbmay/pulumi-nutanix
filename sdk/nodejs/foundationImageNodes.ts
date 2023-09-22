// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class FoundationImageNodes extends pulumi.CustomResource {
    /**
     * Get an existing FoundationImageNodes resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FoundationImageNodesState, opts?: pulumi.CustomResourceOptions): FoundationImageNodes {
        return new FoundationImageNodes(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'nutanix:index/foundationImageNodes:FoundationImageNodes';

    /**
     * Returns true if the given object is an instance of FoundationImageNodes.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FoundationImageNodes {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FoundationImageNodes.__pulumiType;
    }

    public readonly blocks!: pulumi.Output<outputs.FoundationImageNodesBlock[]>;
    /**
     * - list containing cluster name and cluster urls for created clusters in current session
     * * `cluster_urls.#.cluster_name` :- clusterName
     * * `cluster_urls.#.cluster_url` :- url to access the cluster login
     */
    public /*out*/ readonly clusterUrls!: pulumi.Output<outputs.FoundationImageNodesClusterUrl[]>;
    public readonly clusters!: pulumi.Output<outputs.FoundationImageNodesCluster[] | undefined>;
    /**
     * - (Required) CVM gateway.
     */
    public readonly cvmGateway!: pulumi.Output<string>;
    /**
     * - (Required) CVM netmask.
     */
    public readonly cvmNetmask!: pulumi.Output<string>;
    /**
     * - Contains user data from Eos portal.
     */
    public readonly eosMetadata!: pulumi.Output<outputs.FoundationImageNodesEosMetadata | undefined>;
    /**
     * - Foundation Central specific settings.
     */
    public readonly fcSettings!: pulumi.Output<outputs.FoundationImageNodesFcSettings | undefined>;
    /**
     * - Hyperv External virtual network adapter name.
     */
    public readonly hypervExternalVnic!: pulumi.Output<string | undefined>;
    /**
     * - Hyperv External vswitch name.
     */
    public readonly hypervExternalVswitch!: pulumi.Output<string | undefined>;
    /**
     * - Hyperv product key.
     */
    public readonly hypervProductKey!: pulumi.Output<string | undefined>;
    /**
     * - Hyperv SKU.
     */
    public readonly hypervSku!: pulumi.Output<boolean | undefined>;
    /**
     * - (Required) Hypervisor gateway.
     */
    public readonly hypervisorGateway!: pulumi.Output<string>;
    /**
     * - Hypervisor ISO.
     */
    public readonly hypervisorIso!: pulumi.Output<outputs.FoundationImageNodesHypervisorIso | undefined>;
    public readonly hypervisorNameserver!: pulumi.Output<string | undefined>;
    /**
     * - (Required) Hypervisor netmask.
     */
    public readonly hypervisorNetmask!: pulumi.Output<string>;
    /**
     * - Hypervisor password.
     */
    public readonly hypervisorPassword!: pulumi.Output<string | undefined>;
    /**
     * - install script.
     */
    public readonly installScript!: pulumi.Output<string | undefined>;
    /**
     * - (Required incase using IPMI based imaging either here or inside node spec) default IPMI gateway
     */
    public readonly ipmiGateway!: pulumi.Output<string | undefined>;
    /**
     * - (Required incase using IPMI based imaging either here or inside node spec) default IPMI netmask
     */
    public readonly ipmiNetmask!: pulumi.Output<string | undefined>;
    /**
     * - (Required incase using IPMI based imaging either here or inside node spec) IPMI password.
     */
    public readonly ipmiPassword!: pulumi.Output<string | undefined>;
    /**
     * - (Required incase using IPMI based imaging either here or inside node spec) IPMI username.
     */
    public readonly ipmiUser!: pulumi.Output<string | undefined>;
    /**
     * - Id of the custom layout which needs to be passed to imaging request.
     */
    public readonly layoutEggUuid!: pulumi.Output<string | undefined>;
    /**
     * - (Required) NOS package.
     */
    public readonly nosPackage!: pulumi.Output<string>;
    /**
     * - sessionId of the imaging session
     */
    public /*out*/ readonly sessionId!: pulumi.Output<string>;
    /**
     * - If hypervisor installation should be skipped.
     */
    public readonly skipHypervisor!: pulumi.Output<boolean | undefined>;
    /**
     * - Arguments to be passed to svmRescue for AOS installation. Ensure that the arguments provided are supported by the AOS version used for imaging.
     */
    public readonly svmRescueArgs!: pulumi.Output<string[] | undefined>;
    /**
     * - Types of tests to be performed.
     */
    public readonly tests!: pulumi.Output<outputs.FoundationImageNodesTests | undefined>;
    /**
     * - UCSM IP address.
     */
    public readonly ucsmIp!: pulumi.Output<string | undefined>;
    /**
     * - UCSM password.
     */
    public readonly ucsmPassword!: pulumi.Output<string | undefined>;
    /**
     * - UCSM username.
     */
    public readonly ucsmUser!: pulumi.Output<string | undefined>;
    /**
     * - UNC password.
     */
    public readonly uncPassword!: pulumi.Output<string | undefined>;
    /**
     * - UNC Path.
     */
    public readonly uncPath!: pulumi.Output<string | undefined>;
    /**
     * - UNC username.
     */
    public readonly uncUsername!: pulumi.Output<string | undefined>;
    /**
     * - xen config types.
     */
    public readonly xenConfigType!: pulumi.Output<string | undefined>;
    /**
     * - xen server master IP address.
     */
    public readonly xsMasterIp!: pulumi.Output<string | undefined>;
    /**
     * - xen server master label.
     */
    public readonly xsMasterLabel!: pulumi.Output<string | undefined>;
    /**
     * - xen server master password.
     */
    public readonly xsMasterPassword!: pulumi.Output<string | undefined>;
    /**
     * - xen server master username.
     */
    public readonly xsMasterUsername!: pulumi.Output<string | undefined>;

    /**
     * Create a FoundationImageNodes resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FoundationImageNodesArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FoundationImageNodesArgs | FoundationImageNodesState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FoundationImageNodesState | undefined;
            resourceInputs["blocks"] = state ? state.blocks : undefined;
            resourceInputs["clusterUrls"] = state ? state.clusterUrls : undefined;
            resourceInputs["clusters"] = state ? state.clusters : undefined;
            resourceInputs["cvmGateway"] = state ? state.cvmGateway : undefined;
            resourceInputs["cvmNetmask"] = state ? state.cvmNetmask : undefined;
            resourceInputs["eosMetadata"] = state ? state.eosMetadata : undefined;
            resourceInputs["fcSettings"] = state ? state.fcSettings : undefined;
            resourceInputs["hypervExternalVnic"] = state ? state.hypervExternalVnic : undefined;
            resourceInputs["hypervExternalVswitch"] = state ? state.hypervExternalVswitch : undefined;
            resourceInputs["hypervProductKey"] = state ? state.hypervProductKey : undefined;
            resourceInputs["hypervSku"] = state ? state.hypervSku : undefined;
            resourceInputs["hypervisorGateway"] = state ? state.hypervisorGateway : undefined;
            resourceInputs["hypervisorIso"] = state ? state.hypervisorIso : undefined;
            resourceInputs["hypervisorNameserver"] = state ? state.hypervisorNameserver : undefined;
            resourceInputs["hypervisorNetmask"] = state ? state.hypervisorNetmask : undefined;
            resourceInputs["hypervisorPassword"] = state ? state.hypervisorPassword : undefined;
            resourceInputs["installScript"] = state ? state.installScript : undefined;
            resourceInputs["ipmiGateway"] = state ? state.ipmiGateway : undefined;
            resourceInputs["ipmiNetmask"] = state ? state.ipmiNetmask : undefined;
            resourceInputs["ipmiPassword"] = state ? state.ipmiPassword : undefined;
            resourceInputs["ipmiUser"] = state ? state.ipmiUser : undefined;
            resourceInputs["layoutEggUuid"] = state ? state.layoutEggUuid : undefined;
            resourceInputs["nosPackage"] = state ? state.nosPackage : undefined;
            resourceInputs["sessionId"] = state ? state.sessionId : undefined;
            resourceInputs["skipHypervisor"] = state ? state.skipHypervisor : undefined;
            resourceInputs["svmRescueArgs"] = state ? state.svmRescueArgs : undefined;
            resourceInputs["tests"] = state ? state.tests : undefined;
            resourceInputs["ucsmIp"] = state ? state.ucsmIp : undefined;
            resourceInputs["ucsmPassword"] = state ? state.ucsmPassword : undefined;
            resourceInputs["ucsmUser"] = state ? state.ucsmUser : undefined;
            resourceInputs["uncPassword"] = state ? state.uncPassword : undefined;
            resourceInputs["uncPath"] = state ? state.uncPath : undefined;
            resourceInputs["uncUsername"] = state ? state.uncUsername : undefined;
            resourceInputs["xenConfigType"] = state ? state.xenConfigType : undefined;
            resourceInputs["xsMasterIp"] = state ? state.xsMasterIp : undefined;
            resourceInputs["xsMasterLabel"] = state ? state.xsMasterLabel : undefined;
            resourceInputs["xsMasterPassword"] = state ? state.xsMasterPassword : undefined;
            resourceInputs["xsMasterUsername"] = state ? state.xsMasterUsername : undefined;
        } else {
            const args = argsOrState as FoundationImageNodesArgs | undefined;
            if ((!args || args.blocks === undefined) && !opts.urn) {
                throw new Error("Missing required property 'blocks'");
            }
            if ((!args || args.cvmGateway === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cvmGateway'");
            }
            if ((!args || args.cvmNetmask === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cvmNetmask'");
            }
            if ((!args || args.hypervisorGateway === undefined) && !opts.urn) {
                throw new Error("Missing required property 'hypervisorGateway'");
            }
            if ((!args || args.hypervisorNetmask === undefined) && !opts.urn) {
                throw new Error("Missing required property 'hypervisorNetmask'");
            }
            if ((!args || args.nosPackage === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nosPackage'");
            }
            resourceInputs["blocks"] = args ? args.blocks : undefined;
            resourceInputs["clusters"] = args ? args.clusters : undefined;
            resourceInputs["cvmGateway"] = args ? args.cvmGateway : undefined;
            resourceInputs["cvmNetmask"] = args ? args.cvmNetmask : undefined;
            resourceInputs["eosMetadata"] = args ? args.eosMetadata : undefined;
            resourceInputs["fcSettings"] = args ? args.fcSettings : undefined;
            resourceInputs["hypervExternalVnic"] = args ? args.hypervExternalVnic : undefined;
            resourceInputs["hypervExternalVswitch"] = args ? args.hypervExternalVswitch : undefined;
            resourceInputs["hypervProductKey"] = args ? args.hypervProductKey : undefined;
            resourceInputs["hypervSku"] = args ? args.hypervSku : undefined;
            resourceInputs["hypervisorGateway"] = args ? args.hypervisorGateway : undefined;
            resourceInputs["hypervisorIso"] = args ? args.hypervisorIso : undefined;
            resourceInputs["hypervisorNameserver"] = args ? args.hypervisorNameserver : undefined;
            resourceInputs["hypervisorNetmask"] = args ? args.hypervisorNetmask : undefined;
            resourceInputs["hypervisorPassword"] = args ? args.hypervisorPassword : undefined;
            resourceInputs["installScript"] = args ? args.installScript : undefined;
            resourceInputs["ipmiGateway"] = args ? args.ipmiGateway : undefined;
            resourceInputs["ipmiNetmask"] = args ? args.ipmiNetmask : undefined;
            resourceInputs["ipmiPassword"] = args ? args.ipmiPassword : undefined;
            resourceInputs["ipmiUser"] = args ? args.ipmiUser : undefined;
            resourceInputs["layoutEggUuid"] = args ? args.layoutEggUuid : undefined;
            resourceInputs["nosPackage"] = args ? args.nosPackage : undefined;
            resourceInputs["skipHypervisor"] = args ? args.skipHypervisor : undefined;
            resourceInputs["svmRescueArgs"] = args ? args.svmRescueArgs : undefined;
            resourceInputs["tests"] = args ? args.tests : undefined;
            resourceInputs["ucsmIp"] = args ? args.ucsmIp : undefined;
            resourceInputs["ucsmPassword"] = args ? args.ucsmPassword : undefined;
            resourceInputs["ucsmUser"] = args ? args.ucsmUser : undefined;
            resourceInputs["uncPassword"] = args ? args.uncPassword : undefined;
            resourceInputs["uncPath"] = args ? args.uncPath : undefined;
            resourceInputs["uncUsername"] = args ? args.uncUsername : undefined;
            resourceInputs["xenConfigType"] = args ? args.xenConfigType : undefined;
            resourceInputs["xsMasterIp"] = args ? args.xsMasterIp : undefined;
            resourceInputs["xsMasterLabel"] = args ? args.xsMasterLabel : undefined;
            resourceInputs["xsMasterPassword"] = args ? args.xsMasterPassword : undefined;
            resourceInputs["xsMasterUsername"] = args ? args.xsMasterUsername : undefined;
            resourceInputs["clusterUrls"] = undefined /*out*/;
            resourceInputs["sessionId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FoundationImageNodes.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FoundationImageNodes resources.
 */
export interface FoundationImageNodesState {
    blocks?: pulumi.Input<pulumi.Input<inputs.FoundationImageNodesBlock>[]>;
    /**
     * - list containing cluster name and cluster urls for created clusters in current session
     * * `cluster_urls.#.cluster_name` :- clusterName
     * * `cluster_urls.#.cluster_url` :- url to access the cluster login
     */
    clusterUrls?: pulumi.Input<pulumi.Input<inputs.FoundationImageNodesClusterUrl>[]>;
    clusters?: pulumi.Input<pulumi.Input<inputs.FoundationImageNodesCluster>[]>;
    /**
     * - (Required) CVM gateway.
     */
    cvmGateway?: pulumi.Input<string>;
    /**
     * - (Required) CVM netmask.
     */
    cvmNetmask?: pulumi.Input<string>;
    /**
     * - Contains user data from Eos portal.
     */
    eosMetadata?: pulumi.Input<inputs.FoundationImageNodesEosMetadata>;
    /**
     * - Foundation Central specific settings.
     */
    fcSettings?: pulumi.Input<inputs.FoundationImageNodesFcSettings>;
    /**
     * - Hyperv External virtual network adapter name.
     */
    hypervExternalVnic?: pulumi.Input<string>;
    /**
     * - Hyperv External vswitch name.
     */
    hypervExternalVswitch?: pulumi.Input<string>;
    /**
     * - Hyperv product key.
     */
    hypervProductKey?: pulumi.Input<string>;
    /**
     * - Hyperv SKU.
     */
    hypervSku?: pulumi.Input<boolean>;
    /**
     * - (Required) Hypervisor gateway.
     */
    hypervisorGateway?: pulumi.Input<string>;
    /**
     * - Hypervisor ISO.
     */
    hypervisorIso?: pulumi.Input<inputs.FoundationImageNodesHypervisorIso>;
    hypervisorNameserver?: pulumi.Input<string>;
    /**
     * - (Required) Hypervisor netmask.
     */
    hypervisorNetmask?: pulumi.Input<string>;
    /**
     * - Hypervisor password.
     */
    hypervisorPassword?: pulumi.Input<string>;
    /**
     * - install script.
     */
    installScript?: pulumi.Input<string>;
    /**
     * - (Required incase using IPMI based imaging either here or inside node spec) default IPMI gateway
     */
    ipmiGateway?: pulumi.Input<string>;
    /**
     * - (Required incase using IPMI based imaging either here or inside node spec) default IPMI netmask
     */
    ipmiNetmask?: pulumi.Input<string>;
    /**
     * - (Required incase using IPMI based imaging either here or inside node spec) IPMI password.
     */
    ipmiPassword?: pulumi.Input<string>;
    /**
     * - (Required incase using IPMI based imaging either here or inside node spec) IPMI username.
     */
    ipmiUser?: pulumi.Input<string>;
    /**
     * - Id of the custom layout which needs to be passed to imaging request.
     */
    layoutEggUuid?: pulumi.Input<string>;
    /**
     * - (Required) NOS package.
     */
    nosPackage?: pulumi.Input<string>;
    /**
     * - sessionId of the imaging session
     */
    sessionId?: pulumi.Input<string>;
    /**
     * - If hypervisor installation should be skipped.
     */
    skipHypervisor?: pulumi.Input<boolean>;
    /**
     * - Arguments to be passed to svmRescue for AOS installation. Ensure that the arguments provided are supported by the AOS version used for imaging.
     */
    svmRescueArgs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * - Types of tests to be performed.
     */
    tests?: pulumi.Input<inputs.FoundationImageNodesTests>;
    /**
     * - UCSM IP address.
     */
    ucsmIp?: pulumi.Input<string>;
    /**
     * - UCSM password.
     */
    ucsmPassword?: pulumi.Input<string>;
    /**
     * - UCSM username.
     */
    ucsmUser?: pulumi.Input<string>;
    /**
     * - UNC password.
     */
    uncPassword?: pulumi.Input<string>;
    /**
     * - UNC Path.
     */
    uncPath?: pulumi.Input<string>;
    /**
     * - UNC username.
     */
    uncUsername?: pulumi.Input<string>;
    /**
     * - xen config types.
     */
    xenConfigType?: pulumi.Input<string>;
    /**
     * - xen server master IP address.
     */
    xsMasterIp?: pulumi.Input<string>;
    /**
     * - xen server master label.
     */
    xsMasterLabel?: pulumi.Input<string>;
    /**
     * - xen server master password.
     */
    xsMasterPassword?: pulumi.Input<string>;
    /**
     * - xen server master username.
     */
    xsMasterUsername?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FoundationImageNodes resource.
 */
export interface FoundationImageNodesArgs {
    blocks: pulumi.Input<pulumi.Input<inputs.FoundationImageNodesBlock>[]>;
    clusters?: pulumi.Input<pulumi.Input<inputs.FoundationImageNodesCluster>[]>;
    /**
     * - (Required) CVM gateway.
     */
    cvmGateway: pulumi.Input<string>;
    /**
     * - (Required) CVM netmask.
     */
    cvmNetmask: pulumi.Input<string>;
    /**
     * - Contains user data from Eos portal.
     */
    eosMetadata?: pulumi.Input<inputs.FoundationImageNodesEosMetadata>;
    /**
     * - Foundation Central specific settings.
     */
    fcSettings?: pulumi.Input<inputs.FoundationImageNodesFcSettings>;
    /**
     * - Hyperv External virtual network adapter name.
     */
    hypervExternalVnic?: pulumi.Input<string>;
    /**
     * - Hyperv External vswitch name.
     */
    hypervExternalVswitch?: pulumi.Input<string>;
    /**
     * - Hyperv product key.
     */
    hypervProductKey?: pulumi.Input<string>;
    /**
     * - Hyperv SKU.
     */
    hypervSku?: pulumi.Input<boolean>;
    /**
     * - (Required) Hypervisor gateway.
     */
    hypervisorGateway: pulumi.Input<string>;
    /**
     * - Hypervisor ISO.
     */
    hypervisorIso?: pulumi.Input<inputs.FoundationImageNodesHypervisorIso>;
    hypervisorNameserver?: pulumi.Input<string>;
    /**
     * - (Required) Hypervisor netmask.
     */
    hypervisorNetmask: pulumi.Input<string>;
    /**
     * - Hypervisor password.
     */
    hypervisorPassword?: pulumi.Input<string>;
    /**
     * - install script.
     */
    installScript?: pulumi.Input<string>;
    /**
     * - (Required incase using IPMI based imaging either here or inside node spec) default IPMI gateway
     */
    ipmiGateway?: pulumi.Input<string>;
    /**
     * - (Required incase using IPMI based imaging either here or inside node spec) default IPMI netmask
     */
    ipmiNetmask?: pulumi.Input<string>;
    /**
     * - (Required incase using IPMI based imaging either here or inside node spec) IPMI password.
     */
    ipmiPassword?: pulumi.Input<string>;
    /**
     * - (Required incase using IPMI based imaging either here or inside node spec) IPMI username.
     */
    ipmiUser?: pulumi.Input<string>;
    /**
     * - Id of the custom layout which needs to be passed to imaging request.
     */
    layoutEggUuid?: pulumi.Input<string>;
    /**
     * - (Required) NOS package.
     */
    nosPackage: pulumi.Input<string>;
    /**
     * - If hypervisor installation should be skipped.
     */
    skipHypervisor?: pulumi.Input<boolean>;
    /**
     * - Arguments to be passed to svmRescue for AOS installation. Ensure that the arguments provided are supported by the AOS version used for imaging.
     */
    svmRescueArgs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * - Types of tests to be performed.
     */
    tests?: pulumi.Input<inputs.FoundationImageNodesTests>;
    /**
     * - UCSM IP address.
     */
    ucsmIp?: pulumi.Input<string>;
    /**
     * - UCSM password.
     */
    ucsmPassword?: pulumi.Input<string>;
    /**
     * - UCSM username.
     */
    ucsmUser?: pulumi.Input<string>;
    /**
     * - UNC password.
     */
    uncPassword?: pulumi.Input<string>;
    /**
     * - UNC Path.
     */
    uncPath?: pulumi.Input<string>;
    /**
     * - UNC username.
     */
    uncUsername?: pulumi.Input<string>;
    /**
     * - xen config types.
     */
    xenConfigType?: pulumi.Input<string>;
    /**
     * - xen server master IP address.
     */
    xsMasterIp?: pulumi.Input<string>;
    /**
     * - xen server master label.
     */
    xsMasterLabel?: pulumi.Input<string>;
    /**
     * - xen server master password.
     */
    xsMasterPassword?: pulumi.Input<string>;
    /**
     * - xen server master username.
     */
    xsMasterUsername?: pulumi.Input<string>;
}