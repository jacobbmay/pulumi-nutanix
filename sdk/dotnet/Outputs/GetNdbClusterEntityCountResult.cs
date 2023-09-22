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
    public sealed class GetNdbClusterEntityCountResult
    {
        public readonly int DbServers;
        public readonly ImmutableArray<Outputs.GetNdbClusterEntityCountEngineCountResult> EngineCounts;

        [OutputConstructor]
        private GetNdbClusterEntityCountResult(
            int dbServers,

            ImmutableArray<Outputs.GetNdbClusterEntityCountEngineCountResult> engineCounts)
        {
            DbServers = dbServers;
            EngineCounts = engineCounts;
        }
    }
}