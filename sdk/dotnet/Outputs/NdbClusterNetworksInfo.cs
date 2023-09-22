// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace JacobBMay.Nutanix.Outputs
{

    [OutputType]
    public sealed class NdbClusterNetworksInfo
    {
        /// <summary>
        /// VLAN access types for which you want to configure network segmentation. Supports [PRISM, DSIP, DBSERVER ]. 
        /// Prism Element: Select this VLAN access type to configure a VLAN that the NDB agent VM can use to communicate with Prism.
        /// Prism iSCSI Data Service. Select this VLAN access type to configure a VLAN that the agent VM can use to make connection requests to the iSCSI data services IP.
        /// DBServer Access from NDB server. Select this VLAN access type to configure a VLAN that is used for communications between the NDB agent VM and the database server VM on the newly registered NDB server cluster.
        /// </summary>
        public readonly ImmutableArray<string> AccessTypes;
        /// <summary>
        /// network segmentation to segment the network traffic
        /// </summary>
        public readonly ImmutableArray<Outputs.NdbClusterNetworksInfoNetworkInfo> NetworkInfos;
        /// <summary>
        /// type of vlan. Supported [DHCP, Static, IPAM]
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private NdbClusterNetworksInfo(
            ImmutableArray<string> accessTypes,

            ImmutableArray<Outputs.NdbClusterNetworksInfoNetworkInfo> networkInfos,

            string? type)
        {
            AccessTypes = accessTypes;
            NetworkInfos = networkInfos;
            Type = type;
        }
    }
}