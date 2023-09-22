// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Describes Clusters
 */
export function getCluster(args?: GetClusterArgs, opts?: pulumi.InvokeOptions): Promise<GetClusterResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nutanix:index/getCluster:getCluster", {
        "categories": args.categories,
        "clusterId": args.clusterId,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getCluster.
 */
export interface GetClusterArgs {
    /**
     * - Categories for the image.
     */
    categories?: inputs.GetClusterCategory[];
    /**
     * Represents clusters uuid
     */
    clusterId?: string;
    /**
     * Represents the name of cluster
     */
    name?: string;
}

/**
 * A collection of values returned by getCluster.
 */
export interface GetClusterResult {
    /**
     * - Map of cluster efficiency which includes numbers of inefficient vms. The value is populated by analytics on PC. (Readonly)
     */
    readonly analysisVmEfficiencyMap: {[key: string]: string};
    /**
     * The API version.
     */
    readonly apiVersion: string;
    readonly authorizedPublicKeyLists: outputs.GetClusterAuthorizedPublicKeyList[];
    /**
     * - Cluster build details.
     */
    readonly build: {[key: string]: string};
    /**
     * - Zone name used in value of TZ environment variable.
     */
    readonly caCertificateLists: outputs.GetClusterCaCertificateList[];
    /**
     * - Categories for the image.
     */
    readonly categories: outputs.GetClusterCategory[];
    /**
     * - Customer information used in Certificate Signing Request for creating digital certificates.
     */
    readonly certificationSigningInfo: {[key: string]: string};
    /**
     * - Client authentication config.
     */
    readonly clientAuth: {[key: string]: string};
    /**
     * - Cluster architecture. (Readonly, Options: Options : X86_64 , PPC64LE)
     */
    readonly clusterArch: string;
    readonly clusterId: string;
    /**
     * - Cluster domain credentials.
     */
    readonly domainServerCredentials: {[key: string]: string};
    /**
     * - Joined domain name. In 'put' request, empty name will unjoin the cluster from current domain.
     */
    readonly domainServerName: string;
    /**
     * -  The IP of the nameserver that can resolve the domain name. Must set when joining the domain.
     */
    readonly domainServerNameserver: string;
    /**
     * - Array of enabled features.
     */
    readonly enabledFeatureLists: string[];
    /**
     * - Cluster encryption status.
     */
    readonly encryptionStatus: string;
    /**
     * - The cluster IP address that provides external entities access to various cluster data services.
     */
    readonly externalDataServicesIp: string;
    /**
     * - The local IP of cluster visible externally.
     */
    readonly externalIp: string;
    /**
     * - External subnet for cross server communication. The format is IP/netmask. (default 172.16.0.0/255.240.0.0)
     */
    readonly externalSubnet: string;
    /**
     * - GPU driver version.
     */
    readonly gpuDriverVersion: string;
    /**
     * - List of proxies to connect to the service centers.
     */
    readonly httpProxyLists: outputs.GetClusterHttpProxyList[];
    /**
     * - HTTP proxy whitelist.
     */
    readonly httpProxyWhitelists: outputs.GetClusterHttpProxyWhitelist[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * - The internal subnet is local to every server - its not visible outside.iSCSI requests generated internally within the appliance (by user VMs or VMFS) are sent to the internal subnet. The format is IP/netmask.
     */
    readonly internalSubnet: string;
    /**
     * - Indicates if cluster is available to contact. (Readonly)
     */
    readonly isAvailable: boolean;
    /**
     * - List of cluster management servers. (Readonly)
     */
    readonly managementServerLists: outputs.GetClusterManagementServerList[];
    /**
     * - The cluster NAT'd or proxy IP which maps to the cluster local IP.
     */
    readonly masqueradingIp: string;
    /**
     * - Port used together with masqueradingIp to connect to the cluster.
     */
    readonly masqueradingPort: number;
    /**
     * - The image kind metadata.
     */
    readonly metadata: {[key: string]: string};
    /**
     * - the name.
     */
    readonly name: string;
    /**
     * - The list of IP addresses of the name servers.
     */
    readonly nameServerIpLists: string[];
    /**
     * - Comma separated list of subnets (of the form 'a.b.c.d/l.m.n.o') that are allowed to send NFS requests to this container. If not specified, the global NFS whitelist will be looked up for access permission. The internal subnet is always automatically considered part of the whitelist, even if the field below does not explicitly specify it. Similarly, all the hypervisor IPs are considered part of the whitelist. Finally, to permit debugging, all of the SVMs local IPs are considered to be implicitly part of the whitelist.
     */
    readonly nfsSubnetWhitelists: string[];
    readonly nodes: outputs.GetClusterNode[];
    /**
     * - The list of IP addresses or FQDNs of the NTP servers.
     */
    readonly ntpServerIpLists: string[];
    /**
     * - Cluster operation mode. - 'NORMAL': Cluster is operating normally. - 'READ_ONLY': Cluster is operating in read only mode. - 'STAND_ALONE': Only one node is operational in the cluster. This is valid only for single node or two node clusters. - 'SWITCH_TO_TWO_NODE': Cluster is moving from single node to two node cluster. - 'OVERRIDE': Valid only for single node cluster. If the user wants to run vms on a single node cluster in read only mode, he can set the cluster peration mode to override. Writes will be allowed in override mode.
     */
    readonly operationMode: string;
    /**
     * - The reference to a user.
     */
    readonly ownerReference: {[key: string]: string};
    /**
     * - The reference to a project.
     */
    readonly projectReference: {[key: string]: string};
    /**
     * - Array of enabled cluster services. For example, a cluster can function as both AOS and cloud data gateway. - 'AOS': Regular Prism Element - 'PRISM_CENTRAL': Prism Central - 'CLOUD_DATA_GATEWAY': Cloud backup and DR gateway - 'AFS': Cluster for file server - 'WITNESS' : Witness cluster - 'XI_PORTAL': Xi cluster.
     */
    readonly serviceLists: string[];
    /**
     * - SMTP Server Address.
     */
    readonly smtpServerAddress: {[key: string]: string};
    /**
     * - SMTP Server Credentials.
     */
    readonly smtpServerCredentials: {[key: string]: string};
    /**
     * - SMTP Server Email Address.
     */
    readonly smtpServerEmailAddress: string;
    /**
     * - SMTP Server Proxy Type List
     */
    readonly smtpServerProxyTypeLists: string[];
    /**
     * - SMTP Server type.
     */
    readonly smtpServerType: string;
    /**
     * - Map of software on the cluster with software type as the key.
     */
    readonly softwareMapNcc: {[key: string]: any};
    /**
     * - Map of software on the cluster with software type as the key.
     */
    readonly softwareMapNos: {[key: string]: any};
    /**
     * - UTC date and time in RFC-3339 format when the key expires
     */
    readonly sslKeyExpireDatetime: string;
    readonly sslKeyName: string;
    /**
     * - Customer information used in Certificate Signing Request for creating digital certificates.
     */
    readonly sslKeySigningInfo: {[key: string]: string};
    /**
     * - SSL key type. Key types with RSA_2048, ECDSA_256 and ECDSA_384 are supported for key generation and importing.
     */
    readonly sslKeyType: string;
    /**
     * - The state of the cluster entity.
     */
    readonly state: string;
    /**
     * - Verbosity level settings for populating support information. - 'Nothing': Send nothing - 'Basic': Send basic information - skip core dump and hypervisor stats information - 'BasicPlusCoreDump': Send basic and core dump information - 'All': Send all information (Default value: BASIC_PLUS_CORE_DUMP)
     */
    readonly supportedInformationVerbosity: string;
    /**
     * - Zone name used in value of TZ environment variable.
     */
    readonly timezone: string;
}
/**
 * Describes Clusters
 */
export function getClusterOutput(args?: GetClusterOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetClusterResult> {
    return pulumi.output(args).apply((a: any) => getCluster(a, opts))
}

/**
 * A collection of arguments for invoking getCluster.
 */
export interface GetClusterOutputArgs {
    /**
     * - Categories for the image.
     */
    categories?: pulumi.Input<pulumi.Input<inputs.GetClusterCategoryArgs>[]>;
    /**
     * Represents clusters uuid
     */
    clusterId?: pulumi.Input<string>;
    /**
     * Represents the name of cluster
     */
    name?: pulumi.Input<string>;
}