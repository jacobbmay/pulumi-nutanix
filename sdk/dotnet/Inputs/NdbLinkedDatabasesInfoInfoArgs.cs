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

    public sealed class NdbLinkedDatabasesInfoInfoArgs : global::Pulumi.ResourceArgs
    {
        [Input("createdBy")]
        public Input<string>? CreatedBy { get; set; }

        public NdbLinkedDatabasesInfoInfoArgs()
        {
        }
        public static new NdbLinkedDatabasesInfoInfoArgs Empty => new NdbLinkedDatabasesInfoInfoArgs();
    }
}