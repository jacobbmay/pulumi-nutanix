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
    public sealed class GetFloatingIpsMetadataResult
    {
        public readonly string Filter;
        /// <summary>
        /// - The kind name (Default value: project).
        /// </summary>
        public readonly string Kind;
        public readonly int Length;
        public readonly int Offset;
        public readonly string SortAttribute;
        public readonly string SortOrder;
        public readonly int TotalMatches;

        [OutputConstructor]
        private GetFloatingIpsMetadataResult(
            string filter,

            string kind,

            int length,

            int offset,

            string sortAttribute,

            string sortOrder,

            int totalMatches)
        {
            Filter = filter;
            Kind = kind;
            Length = length;
            Offset = offset;
            SortAttribute = sortAttribute;
            SortOrder = sortOrder;
            TotalMatches = totalMatches;
        }
    }
}