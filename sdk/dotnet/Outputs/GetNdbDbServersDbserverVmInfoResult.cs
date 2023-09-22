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
    public sealed class GetNdbDbServersDbserverVmInfoResult
    {
        public readonly ImmutableArray<Outputs.GetNdbDbServersDbserverVmInfoDeregisterInfoResult> DeregisterInfos;
        public readonly ImmutableDictionary<string, string> Distribution;
        public readonly ImmutableArray<Outputs.GetNdbDbServersDbserverVmInfoInfoResult> Infos;
        public readonly ImmutableArray<Outputs.GetNdbDbServersDbserverVmInfoNetworkInfoResult> NetworkInfos;
        public readonly string OsType;
        public readonly string OsVersion;
        public readonly ImmutableDictionary<string, string> SecureInfo;

        [OutputConstructor]
        private GetNdbDbServersDbserverVmInfoResult(
            ImmutableArray<Outputs.GetNdbDbServersDbserverVmInfoDeregisterInfoResult> deregisterInfos,

            ImmutableDictionary<string, string> distribution,

            ImmutableArray<Outputs.GetNdbDbServersDbserverVmInfoInfoResult> infos,

            ImmutableArray<Outputs.GetNdbDbServersDbserverVmInfoNetworkInfoResult> networkInfos,

            string osType,

            string osVersion,

            ImmutableDictionary<string, string> secureInfo)
        {
            DeregisterInfos = deregisterInfos;
            Distribution = distribution;
            Infos = infos;
            NetworkInfos = networkInfos;
            OsType = osType;
            OsVersion = osVersion;
            SecureInfo = secureInfo;
        }
    }
}