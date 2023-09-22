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
    public sealed class GetNdbDatabasesDatabaseInstanceInfoResult
    {
        public readonly ImmutableArray<Outputs.GetNdbDatabasesDatabaseInstanceInfoBpgConfigResult> BpgConfigs;
        public readonly ImmutableDictionary<string, string> SecureInfo;

        [OutputConstructor]
        private GetNdbDatabasesDatabaseInstanceInfoResult(
            ImmutableArray<Outputs.GetNdbDatabasesDatabaseInstanceInfoBpgConfigResult> bpgConfigs,

            ImmutableDictionary<string, string> secureInfo)
        {
            BpgConfigs = bpgConfigs;
            SecureInfo = secureInfo;
        }
    }
}