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

    public sealed class NdbStretchedVlanVlansListPropertiesMapGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("vlanGateway")]
        public Input<string>? VlanGateway { get; set; }

        [Input("vlanPrimaryDns")]
        public Input<string>? VlanPrimaryDns { get; set; }

        [Input("vlanSecondaryDns")]
        public Input<string>? VlanSecondaryDns { get; set; }

        [Input("vlanSubnetMask")]
        public Input<string>? VlanSubnetMask { get; set; }

        public NdbStretchedVlanVlansListPropertiesMapGetArgs()
        {
        }
        public static new NdbStretchedVlanVlansListPropertiesMapGetArgs Empty => new NdbStretchedVlanVlansListPropertiesMapGetArgs();
    }
}