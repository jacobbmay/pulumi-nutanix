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

    public sealed class NdbMaintenanceWindowPropertyGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name for the maintenance window.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("value")]
        public Input<string>? Value { get; set; }

        public NdbMaintenanceWindowPropertyGetArgs()
        {
        }
        public static new NdbMaintenanceWindowPropertyGetArgs Empty => new NdbMaintenanceWindowPropertyGetArgs();
    }
}