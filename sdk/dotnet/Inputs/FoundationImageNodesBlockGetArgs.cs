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

    public sealed class FoundationImageNodesBlockGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// - Block ID.
        /// </summary>
        [Input("blockId")]
        public Input<string>? BlockId { get; set; }

        [Input("nodes", required: true)]
        private InputList<Inputs.FoundationImageNodesBlockNodeGetArgs>? _nodes;
        public InputList<Inputs.FoundationImageNodesBlockNodeGetArgs> Nodes
        {
            get => _nodes ?? (_nodes = new InputList<Inputs.FoundationImageNodesBlockNodeGetArgs>());
            set => _nodes = value;
        }

        public FoundationImageNodesBlockGetArgs()
        {
        }
        public static new FoundationImageNodesBlockGetArgs Empty => new FoundationImageNodesBlockGetArgs();
    }
}