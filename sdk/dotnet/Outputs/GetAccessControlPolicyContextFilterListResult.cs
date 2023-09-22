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
    public sealed class GetAccessControlPolicyContextFilterListResult
    {
        /// <summary>
        /// A list of Entity filter expressions.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAccessControlPolicyContextFilterListEntityFilterExpressionListResult> EntityFilterExpressionLists;
        /// <summary>
        /// - The device ID which is used to uniquely identify this particular disk.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAccessControlPolicyContextFilterListScopeFilterExpressionListResult> ScopeFilterExpressionLists;

        [OutputConstructor]
        private GetAccessControlPolicyContextFilterListResult(
            ImmutableArray<Outputs.GetAccessControlPolicyContextFilterListEntityFilterExpressionListResult> entityFilterExpressionLists,

            ImmutableArray<Outputs.GetAccessControlPolicyContextFilterListScopeFilterExpressionListResult> scopeFilterExpressionLists)
        {
            EntityFilterExpressionLists = entityFilterExpressionLists;
            ScopeFilterExpressionLists = scopeFilterExpressionLists;
        }
    }
}