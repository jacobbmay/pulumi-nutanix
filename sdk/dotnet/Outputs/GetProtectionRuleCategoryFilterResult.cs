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
    public sealed class GetProtectionRuleCategoryFilterResult
    {
        public readonly ImmutableArray<string> KindLists;
        public readonly ImmutableArray<Outputs.GetProtectionRuleCategoryFilterParamResult> Params;
        public readonly string Type;

        [OutputConstructor]
        private GetProtectionRuleCategoryFilterResult(
            ImmutableArray<string> kindLists,

            ImmutableArray<Outputs.GetProtectionRuleCategoryFilterParamResult> @params,

            string type)
        {
            KindLists = kindLists;
            Params = @params;
            Type = type;
        }
    }
}
