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

    public sealed class NdbDbserverVMMaintenanceTasksGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("maintenanceWindowId")]
        public Input<string>? MaintenanceWindowId { get; set; }

        [Input("tasks")]
        private InputList<Inputs.NdbDbserverVMMaintenanceTasksTaskGetArgs>? _tasks;
        public InputList<Inputs.NdbDbserverVMMaintenanceTasksTaskGetArgs> Tasks
        {
            get => _tasks ?? (_tasks = new InputList<Inputs.NdbDbserverVMMaintenanceTasksTaskGetArgs>());
            set => _tasks = value;
        }

        public NdbDbserverVMMaintenanceTasksGetArgs()
        {
        }
        public static new NdbDbserverVMMaintenanceTasksGetArgs Empty => new NdbDbserverVMMaintenanceTasksGetArgs();
    }
}
