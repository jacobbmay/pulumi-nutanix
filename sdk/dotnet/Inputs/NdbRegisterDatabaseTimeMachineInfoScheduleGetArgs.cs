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

    public sealed class NdbRegisterDatabaseTimeMachineInfoScheduleGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// - (Optional) snapshot freq and log config
        /// </summary>
        [Input("continuousschedule")]
        public Input<Inputs.NdbRegisterDatabaseTimeMachineInfoScheduleContinuousscheduleGetArgs>? Continuousschedule { get; set; }

        /// <summary>
        /// - (Optional) monthly snapshot config
        /// </summary>
        [Input("monthlyschedule")]
        public Input<Inputs.NdbRegisterDatabaseTimeMachineInfoScheduleMonthlyscheduleGetArgs>? Monthlyschedule { get; set; }

        /// <summary>
        /// - (Optional) quaterly snapshot config
        /// </summary>
        [Input("quartelyschedule")]
        public Input<Inputs.NdbRegisterDatabaseTimeMachineInfoScheduleQuartelyscheduleGetArgs>? Quartelyschedule { get; set; }

        /// <summary>
        /// - (Optional) daily snapshot config
        /// </summary>
        [Input("snapshottimeofday")]
        public Input<Inputs.NdbRegisterDatabaseTimeMachineInfoScheduleSnapshottimeofdayGetArgs>? Snapshottimeofday { get; set; }

        /// <summary>
        /// - (Optional) weekly snapshot config
        /// </summary>
        [Input("weeklyschedule")]
        public Input<Inputs.NdbRegisterDatabaseTimeMachineInfoScheduleWeeklyscheduleGetArgs>? Weeklyschedule { get; set; }

        /// <summary>
        /// - (Optional) yearly snapshot config
        /// </summary>
        [Input("yearlyschedule")]
        public Input<Inputs.NdbRegisterDatabaseTimeMachineInfoScheduleYearlyscheduleGetArgs>? Yearlyschedule { get; set; }

        public NdbRegisterDatabaseTimeMachineInfoScheduleGetArgs()
        {
        }
        public static new NdbRegisterDatabaseTimeMachineInfoScheduleGetArgs Empty => new NdbRegisterDatabaseTimeMachineInfoScheduleGetArgs();
    }
}