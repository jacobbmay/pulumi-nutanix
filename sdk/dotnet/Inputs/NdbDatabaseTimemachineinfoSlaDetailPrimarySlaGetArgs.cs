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

    public sealed class NdbDatabaseTimemachineinfoSlaDetailPrimarySlaGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("nxClusterIds")]
        private InputList<string>? _nxClusterIds;
        public InputList<string> NxClusterIds
        {
            get => _nxClusterIds ?? (_nxClusterIds = new InputList<string>());
            set => _nxClusterIds = value;
        }

        [Input("slaId", required: true)]
        public Input<string> SlaId { get; set; } = null!;

        public NdbDatabaseTimemachineinfoSlaDetailPrimarySlaGetArgs()
        {
        }
        public static new NdbDatabaseTimemachineinfoSlaDetailPrimarySlaGetArgs Empty => new NdbDatabaseTimemachineinfoSlaDetailPrimarySlaGetArgs();
    }
}