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
    public sealed class GetNdbMaintenanceWindowsMaintenanceWindowEntityTaskAssocResult
    {
        /// <summary>
        /// access level
        /// </summary>
        public readonly string AccessLevel;
        /// <summary>
        /// created date of maintenance window
        /// </summary>
        public readonly string DateCreated;
        /// <summary>
        /// modified date of maintenance window
        /// </summary>
        public readonly string DateModified;
        /// <summary>
        /// description of maintenance window
        /// </summary>
        public readonly string Description;
        public readonly string Entity;
        public readonly string EntityId;
        public readonly string EntityType;
        public readonly string Id;
        public readonly string MaintenanceWindowId;
        public readonly string MaintenanceWindowOwnerId;
        /// <summary>
        /// name of maintenance window
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// owner id of maintenance window
        /// </summary>
        public readonly string OwnerId;
        public readonly ImmutableArray<Outputs.GetNdbMaintenanceWindowsMaintenanceWindowEntityTaskAssocPayloadResult> Payloads;
        /// <summary>
        /// properties of maintenance window
        /// </summary>
        public readonly ImmutableArray<Outputs.GetNdbMaintenanceWindowsMaintenanceWindowEntityTaskAssocPropertyResult> Properties;
        /// <summary>
        /// status of maintennace window
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// tags of maintenance window
        /// </summary>
        public readonly ImmutableArray<Outputs.GetNdbMaintenanceWindowsMaintenanceWindowEntityTaskAssocTagResult> Tags;
        public readonly string TaskType;

        [OutputConstructor]
        private GetNdbMaintenanceWindowsMaintenanceWindowEntityTaskAssocResult(
            string accessLevel,

            string dateCreated,

            string dateModified,

            string description,

            string entity,

            string entityId,

            string entityType,

            string id,

            string maintenanceWindowId,

            string maintenanceWindowOwnerId,

            string name,

            string ownerId,

            ImmutableArray<Outputs.GetNdbMaintenanceWindowsMaintenanceWindowEntityTaskAssocPayloadResult> payloads,

            ImmutableArray<Outputs.GetNdbMaintenanceWindowsMaintenanceWindowEntityTaskAssocPropertyResult> properties,

            string status,

            ImmutableArray<Outputs.GetNdbMaintenanceWindowsMaintenanceWindowEntityTaskAssocTagResult> tags,

            string taskType)
        {
            AccessLevel = accessLevel;
            DateCreated = dateCreated;
            DateModified = dateModified;
            Description = description;
            Entity = entity;
            EntityId = entityId;
            EntityType = entityType;
            Id = id;
            MaintenanceWindowId = maintenanceWindowId;
            MaintenanceWindowOwnerId = maintenanceWindowOwnerId;
            Name = name;
            OwnerId = ownerId;
            Payloads = payloads;
            Properties = properties;
            Status = status;
            Tags = tags;
            TaskType = taskType;
        }
    }
}