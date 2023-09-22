// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace JacobBMay.Nutanix.Outputs
{

    [OutputType]
    public sealed class NdbDatabaseTimeMachineSchedule
    {
        public readonly ImmutableArray<Outputs.NdbDatabaseTimeMachineScheduleContinuousSchedule> ContinuousSchedules;
        public readonly ImmutableArray<Outputs.NdbDatabaseTimeMachineScheduleDailySchedule> DailySchedules;
        public readonly string? DateCreated;
        public readonly string? DateModified;
        /// <summary>
        /// - (Optional) The description
        /// </summary>
        public readonly string? Description;
        public readonly bool? GlobalPolicy;
        public readonly string? Id;
        public readonly ImmutableArray<Outputs.NdbDatabaseTimeMachineScheduleMonthlySchedule> MonthlySchedules;
        /// <summary>
        /// - (Required) Name of the instance.
        /// </summary>
        public readonly string? Name;
        public readonly string? OwnerId;
        public readonly ImmutableArray<Outputs.NdbDatabaseTimeMachineScheduleQuartelySchedule> QuartelySchedules;
        public readonly int? ReferenceCount;
        public readonly ImmutableArray<Outputs.NdbDatabaseTimeMachineScheduleSnapshotTimeOfDay> SnapshotTimeOfDays;
        public readonly string? StartTime;
        public readonly bool? SystemPolicy;
        public readonly string? TimeZone;
        public readonly string? UniqueName;
        public readonly ImmutableArray<Outputs.NdbDatabaseTimeMachineScheduleWeeklySchedule> WeeklySchedules;
        public readonly ImmutableArray<Outputs.NdbDatabaseTimeMachineScheduleYearlySchedule> YearlySchedules;

        [OutputConstructor]
        private NdbDatabaseTimeMachineSchedule(
            ImmutableArray<Outputs.NdbDatabaseTimeMachineScheduleContinuousSchedule> continuousSchedules,

            ImmutableArray<Outputs.NdbDatabaseTimeMachineScheduleDailySchedule> dailySchedules,

            string? dateCreated,

            string? dateModified,

            string? description,

            bool? globalPolicy,

            string? id,

            ImmutableArray<Outputs.NdbDatabaseTimeMachineScheduleMonthlySchedule> monthlySchedules,

            string? name,

            string? ownerId,

            ImmutableArray<Outputs.NdbDatabaseTimeMachineScheduleQuartelySchedule> quartelySchedules,

            int? referenceCount,

            ImmutableArray<Outputs.NdbDatabaseTimeMachineScheduleSnapshotTimeOfDay> snapshotTimeOfDays,

            string? startTime,

            bool? systemPolicy,

            string? timeZone,

            string? uniqueName,

            ImmutableArray<Outputs.NdbDatabaseTimeMachineScheduleWeeklySchedule> weeklySchedules,

            ImmutableArray<Outputs.NdbDatabaseTimeMachineScheduleYearlySchedule> yearlySchedules)
        {
            ContinuousSchedules = continuousSchedules;
            DailySchedules = dailySchedules;
            DateCreated = dateCreated;
            DateModified = dateModified;
            Description = description;
            GlobalPolicy = globalPolicy;
            Id = id;
            MonthlySchedules = monthlySchedules;
            Name = name;
            OwnerId = ownerId;
            QuartelySchedules = quartelySchedules;
            ReferenceCount = referenceCount;
            SnapshotTimeOfDays = snapshotTimeOfDays;
            StartTime = startTime;
            SystemPolicy = systemPolicy;
            TimeZone = timeZone;
            UniqueName = uniqueName;
            WeeklySchedules = weeklySchedules;
            YearlySchedules = yearlySchedules;
        }
    }
}