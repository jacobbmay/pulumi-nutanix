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
    public sealed class GetClusterNodeResult
    {
        public readonly string Ip;
        public readonly string Type;
        public readonly string Version;

        [OutputConstructor]
        private GetClusterNodeResult(
            string ip,

            string type,

            string version)
        {
            Ip = ip;
            Type = type;
            Version = version;
        }
    }
}