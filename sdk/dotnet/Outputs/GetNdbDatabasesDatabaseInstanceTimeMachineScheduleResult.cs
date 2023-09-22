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
    public sealed class GetNdbDatabasesDatabaseInstanceTimeMachineScheduleResult
    {
        public readonly ImmutableArray<Outputs.GetNdbDatabasesDatabaseInstanceTimeMachineScheduleContinuousScheduleResult> ContinuousSchedules;
        public readonly ImmutableArray<Outputs.GetNdbDatabasesDatabaseInstanceTimeMachineScheduleDailyScheduleResult> DailySchedules;
        public readonly string DateCreated;
        public readonly string DateModified;
        public readonly string Description;
        public readonly bool GlobalPolicy;
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetNdbDatabasesDatabaseInstanceTimeMachineScheduleMonthlyScheduleResult> MonthlySchedules;
        public readonly string Name;
        public readonly string OwnerId;
        public readonly ImmutableArray<Outputs.GetNdbDatabasesDatabaseInstanceTimeMachineScheduleQuartelyScheduleResult> QuartelySchedules;
        public readonly int ReferenceCount;
        public readonly ImmutableArray<Outputs.GetNdbDatabasesDatabaseInstanceTimeMachineScheduleSnapshotTimeOfDayResult> SnapshotTimeOfDays;
        public readonly string StartTime;
        public readonly bool SystemPolicy;
        public readonly string TimeZone;
        public readonly string UniqueName;
        public readonly ImmutableArray<Outputs.GetNdbDatabasesDatabaseInstanceTimeMachineScheduleWeeklyScheduleResult> WeeklySchedules;
        public readonly ImmutableArray<Outputs.GetNdbDatabasesDatabaseInstanceTimeMachineScheduleYearlyScheduleResult> YearlySchedules;

        [OutputConstructor]
        private GetNdbDatabasesDatabaseInstanceTimeMachineScheduleResult(
            ImmutableArray<Outputs.GetNdbDatabasesDatabaseInstanceTimeMachineScheduleContinuousScheduleResult> continuousSchedules,

            ImmutableArray<Outputs.GetNdbDatabasesDatabaseInstanceTimeMachineScheduleDailyScheduleResult> dailySchedules,

            string dateCreated,

            string dateModified,

            string description,

            bool globalPolicy,

            string id,

            ImmutableArray<Outputs.GetNdbDatabasesDatabaseInstanceTimeMachineScheduleMonthlyScheduleResult> monthlySchedules,

            string name,

            string ownerId,

            ImmutableArray<Outputs.GetNdbDatabasesDatabaseInstanceTimeMachineScheduleQuartelyScheduleResult> quartelySchedules,

            int referenceCount,

            ImmutableArray<Outputs.GetNdbDatabasesDatabaseInstanceTimeMachineScheduleSnapshotTimeOfDayResult> snapshotTimeOfDays,

            string startTime,

            bool systemPolicy,

            string timeZone,

            string uniqueName,

            ImmutableArray<Outputs.GetNdbDatabasesDatabaseInstanceTimeMachineScheduleWeeklyScheduleResult> weeklySchedules,

            ImmutableArray<Outputs.GetNdbDatabasesDatabaseInstanceTimeMachineScheduleYearlyScheduleResult> yearlySchedules)
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