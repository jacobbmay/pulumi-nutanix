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

    public sealed class KarbonWorkerNodepoolNodeArgs : global::Pulumi.ResourceArgs
    {
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        [Input("ipv4Address")]
        public Input<string>? Ipv4Address { get; set; }

        public KarbonWorkerNodepoolNodeArgs()
        {
        }
        public static new KarbonWorkerNodepoolNodeArgs Empty => new KarbonWorkerNodepoolNodeArgs();
    }
}
