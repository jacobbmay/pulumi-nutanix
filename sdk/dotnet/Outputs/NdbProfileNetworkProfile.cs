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
    public sealed class NdbProfileNetworkProfile
    {
        public readonly ImmutableArray<Outputs.NdbProfileNetworkProfilePostgresDatabase> PostgresDatabases;
        public readonly string Topology;
        public readonly ImmutableArray<Outputs.NdbProfileNetworkProfileVersionClusterAssociation> VersionClusterAssociations;

        [OutputConstructor]
        private NdbProfileNetworkProfile(
            ImmutableArray<Outputs.NdbProfileNetworkProfilePostgresDatabase> postgresDatabases,

            string topology,

            ImmutableArray<Outputs.NdbProfileNetworkProfileVersionClusterAssociation> versionClusterAssociations)
        {
            PostgresDatabases = postgresDatabases;
            Topology = topology;
            VersionClusterAssociations = versionClusterAssociations;
        }
    }
}