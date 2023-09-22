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
    public sealed class GetRecoveryPlanParameterResult
    {
        public readonly ImmutableArray<Outputs.GetRecoveryPlanParameterFloatingIpAssignmentListResult> FloatingIpAssignmentLists;
        public readonly ImmutableArray<Outputs.GetRecoveryPlanParameterNetworkMappingListResult> NetworkMappingLists;

        [OutputConstructor]
        private GetRecoveryPlanParameterResult(
            ImmutableArray<Outputs.GetRecoveryPlanParameterFloatingIpAssignmentListResult> floatingIpAssignmentLists,

            ImmutableArray<Outputs.GetRecoveryPlanParameterNetworkMappingListResult> networkMappingLists)
        {
            FloatingIpAssignmentLists = floatingIpAssignmentLists;
            NetworkMappingLists = networkMappingLists;
        }
    }
}