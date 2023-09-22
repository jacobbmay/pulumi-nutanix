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

    public sealed class ProtectionRuleAvailabilityZoneConnectivityListSnapshotScheduleListGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("autoSuspendTimeoutSecs")]
        public Input<int>? AutoSuspendTimeoutSecs { get; set; }

        [Input("localSnapshotRetentionPolicy")]
        public Input<Inputs.ProtectionRuleAvailabilityZoneConnectivityListSnapshotScheduleListLocalSnapshotRetentionPolicyGetArgs>? LocalSnapshotRetentionPolicy { get; set; }

        [Input("recoveryPointObjectiveSecs", required: true)]
        public Input<int> RecoveryPointObjectiveSecs { get; set; } = null!;

        [Input("remoteSnapshotRetentionPolicy")]
        public Input<Inputs.ProtectionRuleAvailabilityZoneConnectivityListSnapshotScheduleListRemoteSnapshotRetentionPolicyGetArgs>? RemoteSnapshotRetentionPolicy { get; set; }

        [Input("snapshotType")]
        public Input<string>? SnapshotType { get; set; }

        public ProtectionRuleAvailabilityZoneConnectivityListSnapshotScheduleListGetArgs()
        {
        }
        public static new ProtectionRuleAvailabilityZoneConnectivityListSnapshotScheduleListGetArgs Empty => new ProtectionRuleAvailabilityZoneConnectivityListSnapshotScheduleListGetArgs();
    }
}