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
    public sealed class NdbSoftwareVersionProfilePostgresDatabase
    {
        public readonly string? DbSoftwareNotes;
        public readonly string? OsNotes;
        public readonly string? SourceDbserverId;

        [OutputConstructor]
        private NdbSoftwareVersionProfilePostgresDatabase(
            string? dbSoftwareNotes,

            string? osNotes,

            string? sourceDbserverId)
        {
            DbSoftwareNotes = dbSoftwareNotes;
            OsNotes = osNotes;
            SourceDbserverId = sourceDbserverId;
        }
    }
}
