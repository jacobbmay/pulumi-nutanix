// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace JacobBMay.Nutanix
{
    public static class GetFoundationCentralImagedClustersList
    {
        public static Task<GetFoundationCentralImagedClustersListResult> InvokeAsync(GetFoundationCentralImagedClustersListArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFoundationCentralImagedClustersListResult>("nutanix:index/getFoundationCentralImagedClustersList:getFoundationCentralImagedClustersList", args ?? new GetFoundationCentralImagedClustersListArgs(), options.WithDefaults());

        public static Output<GetFoundationCentralImagedClustersListResult> Invoke(GetFoundationCentralImagedClustersListInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFoundationCentralImagedClustersListResult>("nutanix:index/getFoundationCentralImagedClustersList:getFoundationCentralImagedClustersList", args ?? new GetFoundationCentralImagedClustersListInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFoundationCentralImagedClustersListArgs : global::Pulumi.InvokeArgs
    {
        [Input("filters")]
        public Inputs.GetFoundationCentralImagedClustersListFiltersArgs? Filters { get; set; }

        [Input("length")]
        public int? Length { get; set; }

        [Input("offset")]
        public int? Offset { get; set; }

        public GetFoundationCentralImagedClustersListArgs()
        {
        }
        public static new GetFoundationCentralImagedClustersListArgs Empty => new GetFoundationCentralImagedClustersListArgs();
    }

    public sealed class GetFoundationCentralImagedClustersListInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("filters")]
        public Input<Inputs.GetFoundationCentralImagedClustersListFiltersInputArgs>? Filters { get; set; }

        [Input("length")]
        public Input<int>? Length { get; set; }

        [Input("offset")]
        public Input<int>? Offset { get; set; }

        public GetFoundationCentralImagedClustersListInvokeArgs()
        {
        }
        public static new GetFoundationCentralImagedClustersListInvokeArgs Empty => new GetFoundationCentralImagedClustersListInvokeArgs();
    }


    [OutputType]
    public sealed class GetFoundationCentralImagedClustersListResult
    {
        public readonly Outputs.GetFoundationCentralImagedClustersListFiltersResult? Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetFoundationCentralImagedClustersListImagedClusterResult> ImagedClusters;
        public readonly int? Length;
        public readonly ImmutableArray<Outputs.GetFoundationCentralImagedClustersListMetadataResult> Metadatas;
        public readonly int? Offset;

        [OutputConstructor]
        private GetFoundationCentralImagedClustersListResult(
            Outputs.GetFoundationCentralImagedClustersListFiltersResult? filters,

            string id,

            ImmutableArray<Outputs.GetFoundationCentralImagedClustersListImagedClusterResult> imagedClusters,

            int? length,

            ImmutableArray<Outputs.GetFoundationCentralImagedClustersListMetadataResult> metadatas,

            int? offset)
        {
            Filters = filters;
            Id = id;
            ImagedClusters = imagedClusters;
            Length = length;
            Metadatas = metadatas;
            Offset = offset;
        }
    }
}