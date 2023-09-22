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
    public sealed class NdbScaleDatabaseDatabaseNodeProtectionDomain
    {
        public readonly ImmutableArray<string> AssocEntities;
        public readonly string? CloudId;
        public readonly string? DateCreated;
        public readonly string? DateModified;
        public readonly string? Description;
        public readonly bool? EraCreated;
        public readonly string? Id;
        public readonly string? Name;
        public readonly string? OwnerId;
        public readonly string? PrimaryHost;
        public readonly ImmutableArray<Outputs.NdbScaleDatabaseDatabaseNodeProtectionDomainProperty> Properties;
        public readonly string? Status;
        public readonly string? Type;

        [OutputConstructor]
        private NdbScaleDatabaseDatabaseNodeProtectionDomain(
            ImmutableArray<string> assocEntities,

            string? cloudId,

            string? dateCreated,

            string? dateModified,

            string? description,

            bool? eraCreated,

            string? id,

            string? name,

            string? ownerId,

            string? primaryHost,

            ImmutableArray<Outputs.NdbScaleDatabaseDatabaseNodeProtectionDomainProperty> properties,

            string? status,

            string? type)
        {
            AssocEntities = assocEntities;
            CloudId = cloudId;
            DateCreated = dateCreated;
            DateModified = dateModified;
            Description = description;
            EraCreated = eraCreated;
            Id = id;
            Name = name;
            OwnerId = ownerId;
            PrimaryHost = primaryHost;
            Properties = properties;
            Status = status;
            Type = type;
        }
    }
}