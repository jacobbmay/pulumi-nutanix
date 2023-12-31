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

    public sealed class NdbSoftwareVersionProfileVersionClusterAssociationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("dateCreated")]
        public Input<string>? DateCreated { get; set; }

        [Input("dateModified")]
        public Input<string>? DateModified { get; set; }

        [Input("nxClusterId")]
        public Input<string>? NxClusterId { get; set; }

        [Input("optimizedForProvisioning")]
        public Input<bool>? OptimizedForProvisioning { get; set; }

        [Input("ownerId")]
        public Input<string>? OwnerId { get; set; }

        [Input("profileVersionId")]
        public Input<string>? ProfileVersionId { get; set; }

        [Input("properties")]
        private InputList<Inputs.NdbSoftwareVersionProfileVersionClusterAssociationPropertyGetArgs>? _properties;
        public InputList<Inputs.NdbSoftwareVersionProfileVersionClusterAssociationPropertyGetArgs> Properties
        {
            get => _properties ?? (_properties = new InputList<Inputs.NdbSoftwareVersionProfileVersionClusterAssociationPropertyGetArgs>());
            set => _properties = value;
        }

        [Input("status")]
        public Input<string>? Status { get; set; }

        public NdbSoftwareVersionProfileVersionClusterAssociationGetArgs()
        {
        }
        public static new NdbSoftwareVersionProfileVersionClusterAssociationGetArgs Empty => new NdbSoftwareVersionProfileVersionClusterAssociationGetArgs();
    }
}
