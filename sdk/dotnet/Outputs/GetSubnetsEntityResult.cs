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
    public sealed class GetSubnetsEntityResult
    {
        /// <summary>
        /// version of the API
        /// </summary>
        public readonly string ApiVersion;
        /// <summary>
        /// The reference to a availability_zone.
        /// </summary>
        public readonly ImmutableDictionary<string, string> AvailabilityZoneReference;
        /// <summary>
        /// The API Version.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSubnetsEntityCategoryResult> Categories;
        /// <summary>
        /// The name of a cluster.
        /// </summary>
        public readonly string ClusterName;
        /// <summary>
        /// The reference to a cluster.
        /// </summary>
        public readonly ImmutableDictionary<string, string> ClusterReference;
        public readonly string ClusterUuid;
        /// <summary>
        /// Default gateway IP address.
        /// </summary>
        public readonly string DefaultGatewayIp;
        /// <summary>
        /// A description for subnet.
        /// </summary>
        public readonly string Description;
        public readonly ImmutableArray<string> DhcpDomainNameServerLists;
        /// <summary>
        /// DHCP domain search list for a subnet.
        /// </summary>
        public readonly ImmutableArray<string> DhcpDomainSearchLists;
        /// <summary>
        /// Spec for defining DHCP options.
        /// </summary>
        public readonly ImmutableDictionary<string, string> DhcpOptions;
        /// <summary>
        /// Host address.
        /// </summary>
        public readonly ImmutableDictionary<string, string> DhcpServerAddress;
        /// <summary>
        /// Port Number.
        /// </summary>
        public readonly int DhcpServerAddressPort;
        public readonly bool EnableNat;
        public readonly ImmutableArray<string> IpConfigPoolListRanges;
        public readonly bool IsExternal;
        public readonly ImmutableArray<Outputs.GetSubnetsEntityMessageListResult> MessageLists;
        /// <summary>
        /// The subnet kind metadata.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Metadata;
        /// <summary>
        /// the name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The reference to a network_function_chain.
        /// </summary>
        public readonly ImmutableDictionary<string, string> NetworkFunctionChainReference;
        /// <summary>
        /// The reference to a user.
        /// </summary>
        public readonly ImmutableDictionary<string, string> OwnerReference;
        /// <summary>
        /// -. IP prefix length of the Subnet.
        /// </summary>
        public readonly int PrefixLength;
        /// <summary>
        /// The reference to a project.
        /// </summary>
        public readonly ImmutableDictionary<string, string> ProjectReference;
        /// <summary>
        /// The state of the subnet.
        /// </summary>
        public readonly string State;
        public readonly string SubnetId;
        /// <summary>
        /// Subnet IP address.
        /// </summary>
        public readonly string SubnetIp;
        public readonly string SubnetName;
        /// <summary>
        /// The type of the subnet.
        /// </summary>
        public readonly string SubnetType;
        /// <summary>
        /// VLAN assigned to the subnet.
        /// </summary>
        public readonly int VlanId;
        public readonly ImmutableDictionary<string, string> VpcReference;
        /// <summary>
        /// The name of the vswitch.
        /// </summary>
        public readonly string VswitchName;

        [OutputConstructor]
        private GetSubnetsEntityResult(
            string apiVersion,

            ImmutableDictionary<string, string> availabilityZoneReference,

            ImmutableArray<Outputs.GetSubnetsEntityCategoryResult> categories,

            string clusterName,

            ImmutableDictionary<string, string> clusterReference,

            string clusterUuid,

            string defaultGatewayIp,

            string description,

            ImmutableArray<string> dhcpDomainNameServerLists,

            ImmutableArray<string> dhcpDomainSearchLists,

            ImmutableDictionary<string, string> dhcpOptions,

            ImmutableDictionary<string, string> dhcpServerAddress,

            int dhcpServerAddressPort,

            bool enableNat,

            ImmutableArray<string> ipConfigPoolListRanges,

            bool isExternal,

            ImmutableArray<Outputs.GetSubnetsEntityMessageListResult> messageLists,

            ImmutableDictionary<string, string> metadata,

            string name,

            ImmutableDictionary<string, string> networkFunctionChainReference,

            ImmutableDictionary<string, string> ownerReference,

            int prefixLength,

            ImmutableDictionary<string, string> projectReference,

            string state,

            string subnetId,

            string subnetIp,

            string subnetName,

            string subnetType,

            int vlanId,

            ImmutableDictionary<string, string> vpcReference,

            string vswitchName)
        {
            ApiVersion = apiVersion;
            AvailabilityZoneReference = availabilityZoneReference;
            Categories = categories;
            ClusterName = clusterName;
            ClusterReference = clusterReference;
            ClusterUuid = clusterUuid;
            DefaultGatewayIp = defaultGatewayIp;
            Description = description;
            DhcpDomainNameServerLists = dhcpDomainNameServerLists;
            DhcpDomainSearchLists = dhcpDomainSearchLists;
            DhcpOptions = dhcpOptions;
            DhcpServerAddress = dhcpServerAddress;
            DhcpServerAddressPort = dhcpServerAddressPort;
            EnableNat = enableNat;
            IpConfigPoolListRanges = ipConfigPoolListRanges;
            IsExternal = isExternal;
            MessageLists = messageLists;
            Metadata = metadata;
            Name = name;
            NetworkFunctionChainReference = networkFunctionChainReference;
            OwnerReference = ownerReference;
            PrefixLength = prefixLength;
            ProjectReference = projectReference;
            State = state;
            SubnetId = subnetId;
            SubnetIp = subnetIp;
            SubnetName = subnetName;
            SubnetType = subnetType;
            VlanId = vlanId;
            VpcReference = vpcReference;
            VswitchName = vswitchName;
        }
    }
}