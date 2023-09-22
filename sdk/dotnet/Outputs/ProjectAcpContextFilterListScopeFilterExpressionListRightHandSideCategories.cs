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
    public sealed class ProjectAcpContextFilterListScopeFilterExpressionListRightHandSideCategories
    {
        /// <summary>
        /// The name for the project.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// value of the key.
        /// </summary>
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private ProjectAcpContextFilterListScopeFilterExpressionListRightHandSideCategories(
            string? name,

            ImmutableArray<string> values)
        {
            Name = name;
            Values = values;
        }
    }
}