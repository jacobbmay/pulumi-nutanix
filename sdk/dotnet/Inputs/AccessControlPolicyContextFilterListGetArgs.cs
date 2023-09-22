// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace JacobBMay.Nutanix.Inputs
{

    public sealed class AccessControlPolicyContextFilterListGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("entityFilterExpressionLists", required: true)]
        private InputList<Inputs.AccessControlPolicyContextFilterListEntityFilterExpressionListGetArgs>? _entityFilterExpressionLists;

        /// <summary>
        /// A list of Entity filter expressions.
        /// </summary>
        public InputList<Inputs.AccessControlPolicyContextFilterListEntityFilterExpressionListGetArgs> EntityFilterExpressionLists
        {
            get => _entityFilterExpressionLists ?? (_entityFilterExpressionLists = new InputList<Inputs.AccessControlPolicyContextFilterListEntityFilterExpressionListGetArgs>());
            set => _entityFilterExpressionLists = value;
        }

        [Input("scopeFilterExpressionLists")]
        private InputList<Inputs.AccessControlPolicyContextFilterListScopeFilterExpressionListGetArgs>? _scopeFilterExpressionLists;

        /// <summary>
        /// - (Optional) Filter the scope of an Access Control Policy.
        /// </summary>
        public InputList<Inputs.AccessControlPolicyContextFilterListScopeFilterExpressionListGetArgs> ScopeFilterExpressionLists
        {
            get => _scopeFilterExpressionLists ?? (_scopeFilterExpressionLists = new InputList<Inputs.AccessControlPolicyContextFilterListScopeFilterExpressionListGetArgs>());
            set => _scopeFilterExpressionLists = value;
        }

        public AccessControlPolicyContextFilterListGetArgs()
        {
        }
        public static new AccessControlPolicyContextFilterListGetArgs Empty => new AccessControlPolicyContextFilterListGetArgs();
    }
}