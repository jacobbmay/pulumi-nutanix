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
    public sealed class GetNdbTmsCapabilityCapabilityContinuousRegionDbLogMetadataResult
    {
        public readonly bool CreatedDirectly;
        public readonly int CurationRetryCount;
        public readonly ImmutableArray<Outputs.GetNdbTmsCapabilityCapabilityContinuousRegionDbLogMetadataDeregisterInfoResult> DeregisterInfos;
        public readonly ImmutableDictionary<string, string> Info;
        public readonly ImmutableDictionary<string, string> SecureInfo;
        public readonly bool UpdatedDirectly;

        [OutputConstructor]
        private GetNdbTmsCapabilityCapabilityContinuousRegionDbLogMetadataResult(
            bool createdDirectly,

            int curationRetryCount,

            ImmutableArray<Outputs.GetNdbTmsCapabilityCapabilityContinuousRegionDbLogMetadataDeregisterInfoResult> deregisterInfos,

            ImmutableDictionary<string, string> info,

            ImmutableDictionary<string, string> secureInfo,

            bool updatedDirectly)
        {
            CreatedDirectly = createdDirectly;
            CurationRetryCount = curationRetryCount;
            DeregisterInfos = deregisterInfos;
            Info = info;
            SecureInfo = secureInfo;
            UpdatedDirectly = updatedDirectly;
        }
    }
}