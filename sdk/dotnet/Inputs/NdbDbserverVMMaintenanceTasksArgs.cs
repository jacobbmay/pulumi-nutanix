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

    public sealed class NdbDbserverVMMaintenanceTasksArgs : global::Pulumi.ResourceArgs
    {
        [Input("maintenanceWindowId")]
        public Input<string>? MaintenanceWindowId { get; set; }

        [Input("tasks")]
        private InputList<Inputs.NdbDbserverVMMaintenanceTasksTaskArgs>? _tasks;
        public InputList<Inputs.NdbDbserverVMMaintenanceTasksTaskArgs> Tasks
        {
            get => _tasks ?? (_tasks = new InputList<Inputs.NdbDbserverVMMaintenanceTasksTaskArgs>());
            set => _tasks = value;
        }

        public NdbDbserverVMMaintenanceTasksArgs()
        {
        }
        public static new NdbDbserverVMMaintenanceTasksArgs Empty => new NdbDbserverVMMaintenanceTasksArgs();
    }
}
