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

    public sealed class NdbScaleDatabaseTimeMachineScheduleSnapshotTimeOfDayArgs : global::Pulumi.ResourceArgs
    {
        [Input("extra")]
        public Input<bool>? Extra { get; set; }

        [Input("hours")]
        public Input<int>? Hours { get; set; }

        [Input("minutes")]
        public Input<int>? Minutes { get; set; }

        [Input("seconds")]
        public Input<int>? Seconds { get; set; }

        public NdbScaleDatabaseTimeMachineScheduleSnapshotTimeOfDayArgs()
        {
        }
        public static new NdbScaleDatabaseTimeMachineScheduleSnapshotTimeOfDayArgs Empty => new NdbScaleDatabaseTimeMachineScheduleSnapshotTimeOfDayArgs();
    }
}