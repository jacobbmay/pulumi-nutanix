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
    public sealed class GetFoundationCentralImagedClustersListMetadataResult
    {
        public readonly int Length;
        public readonly int Offset;
        public readonly int TotalMatches;

        [OutputConstructor]
        private GetFoundationCentralImagedClustersListMetadataResult(
            int length,

            int offset,

            int totalMatches)
        {
            Length = length;
            Offset = offset;
            TotalMatches = totalMatches;
        }
    }
}
