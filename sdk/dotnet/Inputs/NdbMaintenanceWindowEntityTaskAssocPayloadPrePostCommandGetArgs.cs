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

    public sealed class NdbMaintenanceWindowEntityTaskAssocPayloadPrePostCommandGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("postCommand")]
        public Input<string>? PostCommand { get; set; }

        [Input("preCommand")]
        public Input<string>? PreCommand { get; set; }

        public NdbMaintenanceWindowEntityTaskAssocPayloadPrePostCommandGetArgs()
        {
        }
        public static new NdbMaintenanceWindowEntityTaskAssocPayloadPrePostCommandGetArgs Empty => new NdbMaintenanceWindowEntityTaskAssocPayloadPrePostCommandGetArgs();
    }
}