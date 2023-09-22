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
    public sealed class GetNdbClonesCloneLinkedDatabaseResult
    {
        /// <summary>
        /// database name
        /// </summary>
        public readonly string DatabaseName;
        /// <summary>
        /// database status
        /// </summary>
        public readonly string DatabaseStatus;
        /// <summary>
        /// date created for clone
        /// </summary>
        public readonly string DateCreated;
        /// <summary>
        /// last modified date for clone
        /// </summary>
        public readonly string DateModified;
        /// <summary>
        /// cloned description
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// cloned id
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// cloned info
        /// </summary>
        public readonly ImmutableArray<Outputs.GetNdbClonesCloneLinkedDatabaseInfoResult> Infos;
        /// <summary>
        /// Metric of clone
        /// </summary>
        public readonly ImmutableDictionary<string, string> Metric;
        /// <summary>
        /// cloned name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// parent database id
        /// </summary>
        public readonly string ParentDatabaseId;
        public readonly string ParentLinkedDatabaseId;
        public readonly string SnapshotId;
        /// <summary>
        /// status of clone
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Default is UTC
        /// </summary>
        public readonly string Timezone;

        [OutputConstructor]
        private GetNdbClonesCloneLinkedDatabaseResult(
            string databaseName,

            string databaseStatus,

            string dateCreated,

            string dateModified,

            string description,

            string id,

            ImmutableArray<Outputs.GetNdbClonesCloneLinkedDatabaseInfoResult> infos,

            ImmutableDictionary<string, string> metric,

            string name,

            string parentDatabaseId,

            string parentLinkedDatabaseId,

            string snapshotId,

            string status,

            string timezone)
        {
            DatabaseName = databaseName;
            DatabaseStatus = databaseStatus;
            DateCreated = dateCreated;
            DateModified = dateModified;
            Description = description;
            Id = id;
            Infos = infos;
            Metric = metric;
            Name = name;
            ParentDatabaseId = parentDatabaseId;
            ParentLinkedDatabaseId = parentLinkedDatabaseId;
            SnapshotId = snapshotId;
            Status = status;
            Timezone = timezone;
        }
    }
}