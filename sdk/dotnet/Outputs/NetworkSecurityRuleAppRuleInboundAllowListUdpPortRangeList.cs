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
    public sealed class NetworkSecurityRuleAppRuleInboundAllowListUdpPortRangeList
    {
        public readonly int? EndPort;
        public readonly int? StartPort;

        [OutputConstructor]
        private NetworkSecurityRuleAppRuleInboundAllowListUdpPortRangeList(
            int? endPort,

            int? startPort)
        {
            EndPort = endPort;
            StartPort = startPort;
        }
    }
}