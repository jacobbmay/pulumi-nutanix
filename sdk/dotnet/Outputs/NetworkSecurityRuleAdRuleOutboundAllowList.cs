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
    public sealed class NetworkSecurityRuleAdRuleOutboundAllowList
    {
        public readonly ImmutableArray<Outputs.NetworkSecurityRuleAdRuleOutboundAllowListAddressGroupInclusionList> AddressGroupInclusionLists;
        public readonly string? ExpirationTime;
        public readonly ImmutableArray<string> FilterKindLists;
        public readonly ImmutableArray<Outputs.NetworkSecurityRuleAdRuleOutboundAllowListFilterParam> FilterParams;
        public readonly string? FilterType;
        public readonly ImmutableArray<Outputs.NetworkSecurityRuleAdRuleOutboundAllowListIcmpTypeCodeList> IcmpTypeCodeLists;
        public readonly string? IpSubnet;
        public readonly string? IpSubnetPrefixLength;
        public readonly ImmutableDictionary<string, string>? NetworkFunctionChainReference;
        public readonly string? PeerSpecificationType;
        public readonly string? Protocol;
        public readonly ImmutableArray<Outputs.NetworkSecurityRuleAdRuleOutboundAllowListServiceGroupList> ServiceGroupLists;
        public readonly ImmutableArray<Outputs.NetworkSecurityRuleAdRuleOutboundAllowListTcpPortRangeList> TcpPortRangeLists;
        public readonly ImmutableArray<Outputs.NetworkSecurityRuleAdRuleOutboundAllowListUdpPortRangeList> UdpPortRangeLists;

        [OutputConstructor]
        private NetworkSecurityRuleAdRuleOutboundAllowList(
            ImmutableArray<Outputs.NetworkSecurityRuleAdRuleOutboundAllowListAddressGroupInclusionList> addressGroupInclusionLists,

            string? expirationTime,

            ImmutableArray<string> filterKindLists,

            ImmutableArray<Outputs.NetworkSecurityRuleAdRuleOutboundAllowListFilterParam> filterParams,

            string? filterType,

            ImmutableArray<Outputs.NetworkSecurityRuleAdRuleOutboundAllowListIcmpTypeCodeList> icmpTypeCodeLists,

            string? ipSubnet,

            string? ipSubnetPrefixLength,

            ImmutableDictionary<string, string>? networkFunctionChainReference,

            string? peerSpecificationType,

            string? protocol,

            ImmutableArray<Outputs.NetworkSecurityRuleAdRuleOutboundAllowListServiceGroupList> serviceGroupLists,

            ImmutableArray<Outputs.NetworkSecurityRuleAdRuleOutboundAllowListTcpPortRangeList> tcpPortRangeLists,

            ImmutableArray<Outputs.NetworkSecurityRuleAdRuleOutboundAllowListUdpPortRangeList> udpPortRangeLists)
        {
            AddressGroupInclusionLists = addressGroupInclusionLists;
            ExpirationTime = expirationTime;
            FilterKindLists = filterKindLists;
            FilterParams = filterParams;
            FilterType = filterType;
            IcmpTypeCodeLists = icmpTypeCodeLists;
            IpSubnet = ipSubnet;
            IpSubnetPrefixLength = ipSubnetPrefixLength;
            NetworkFunctionChainReference = networkFunctionChainReference;
            PeerSpecificationType = peerSpecificationType;
            Protocol = protocol;
            ServiceGroupLists = serviceGroupLists;
            TcpPortRangeLists = tcpPortRangeLists;
            UdpPortRangeLists = udpPortRangeLists;
        }
    }
}