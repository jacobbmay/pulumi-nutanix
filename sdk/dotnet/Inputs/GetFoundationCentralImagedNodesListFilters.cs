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

    public sealed class GetFoundationCentralImagedNodesListFiltersArgs : global::Pulumi.InvokeArgs
    {
        [Input("nodeState")]
        public string? NodeState { get; set; }

        public GetFoundationCentralImagedNodesListFiltersArgs()
        {
        }
        public static new GetFoundationCentralImagedNodesListFiltersArgs Empty => new GetFoundationCentralImagedNodesListFiltersArgs();
    }
}