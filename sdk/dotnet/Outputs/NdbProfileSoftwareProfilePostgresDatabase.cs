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
    public sealed class NdbProfileSoftwareProfilePostgresDatabase
    {
        public readonly string? BaseProfileVersionDescription;
        public readonly string? BaseProfileVersionName;
        public readonly string? DbSoftwareNotes;
        public readonly string? OsNotes;
        public readonly string? SourceDbserverId;

        [OutputConstructor]
        private NdbProfileSoftwareProfilePostgresDatabase(
            string? baseProfileVersionDescription,

            string? baseProfileVersionName,

            string? dbSoftwareNotes,

            string? osNotes,

            string? sourceDbserverId)
        {
            BaseProfileVersionDescription = baseProfileVersionDescription;
            BaseProfileVersionName = baseProfileVersionName;
            DbSoftwareNotes = dbSoftwareNotes;
            OsNotes = osNotes;
            SourceDbserverId = sourceDbserverId;
        }
    }
}