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
    public sealed class GetRecoveryPlansEntityParameterNetworkMappingListAvailabilityZoneNetworkMappingListRecoveryIpAssignmentListResult
    {
        public readonly ImmutableArray<Outputs.GetRecoveryPlansEntityParameterNetworkMappingListAvailabilityZoneNetworkMappingListRecoveryIpAssignmentListIpConfigListResult> IpConfigLists;
        public readonly Outputs.GetRecoveryPlansEntityParameterNetworkMappingListAvailabilityZoneNetworkMappingListRecoveryIpAssignmentListVmReferenceResult VmReference;

        [OutputConstructor]
        private GetRecoveryPlansEntityParameterNetworkMappingListAvailabilityZoneNetworkMappingListRecoveryIpAssignmentListResult(
            ImmutableArray<Outputs.GetRecoveryPlansEntityParameterNetworkMappingListAvailabilityZoneNetworkMappingListRecoveryIpAssignmentListIpConfigListResult> ipConfigLists,

            Outputs.GetRecoveryPlansEntityParameterNetworkMappingListAvailabilityZoneNetworkMappingListRecoveryIpAssignmentListVmReferenceResult vmReference)
        {
            IpConfigLists = ipConfigLists;
            VmReference = vmReference;
        }
    }
}
