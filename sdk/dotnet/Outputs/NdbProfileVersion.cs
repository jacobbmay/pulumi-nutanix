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
    public sealed class NdbProfileVersion
    {
        public readonly string? DbVersion;
        public readonly bool? Deprecated;
        public readonly string? Description;
        public readonly string? EngineType;
        public readonly string? Id;
        public readonly string? Name;
        public readonly string? Owner;
        public readonly string? ProfileId;
        public readonly ImmutableArray<Outputs.NdbProfileVersionProperty> Properties;
        public readonly ImmutableDictionary<string, string>? PropertiesMap;
        public readonly bool? Published;
        public readonly string? Status;
        public readonly bool? SystemProfile;
        public readonly string? Topology;
        public readonly string? Type;
        public readonly string? Version;
        public readonly ImmutableArray<Outputs.NdbProfileVersionVersionClusterAssociation> VersionClusterAssociations;

        [OutputConstructor]
        private NdbProfileVersion(
            string? dbVersion,

            bool? deprecated,

            string? description,

            string? engineType,

            string? id,

            string? name,

            string? owner,

            string? profileId,

            ImmutableArray<Outputs.NdbProfileVersionProperty> properties,

            ImmutableDictionary<string, string>? propertiesMap,

            bool? published,

            string? status,

            bool? systemProfile,

            string? topology,

            string? type,

            string? version,

            ImmutableArray<Outputs.NdbProfileVersionVersionClusterAssociation> versionClusterAssociations)
        {
            DbVersion = dbVersion;
            Deprecated = deprecated;
            Description = description;
            EngineType = engineType;
            Id = id;
            Name = name;
            Owner = owner;
            ProfileId = profileId;
            Properties = properties;
            PropertiesMap = propertiesMap;
            Published = published;
            Status = status;
            SystemProfile = systemProfile;
            Topology = topology;
            Type = type;
            Version = version;
            VersionClusterAssociations = versionClusterAssociations;
        }
    }
}