// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace JacobBMay.Nutanix.Inputs
{

    public sealed class NetworkSecurityRuleAdRuleOutboundAllowListUdpPortRangeListGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("endPort")]
        public Input<int>? EndPort { get; set; }

        [Input("startPort")]
        public Input<int>? StartPort { get; set; }

        public NetworkSecurityRuleAdRuleOutboundAllowListUdpPortRangeListGetArgs()
        {
        }
        public static new NetworkSecurityRuleAdRuleOutboundAllowListUdpPortRangeListGetArgs Empty => new NetworkSecurityRuleAdRuleOutboundAllowListUdpPortRangeListGetArgs();
    }
}