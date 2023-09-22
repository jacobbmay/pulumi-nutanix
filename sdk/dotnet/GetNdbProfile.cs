// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace JacobBMay.Nutanix
{
    public static class GetNdbProfile
    {
        /// <summary>
        /// Describes a profile in Nutanix Database Service
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Nutanix = Pulumi.Nutanix;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var profile1 = Nutanix.GetNdbProfile.Invoke(new()
        ///     {
        ///         ProfileType = "Network",
        ///         ProfileName = "TEST_NETWORK_PROFILE",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["profile"] = profile1,
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetNdbProfileResult> InvokeAsync(GetNdbProfileArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNdbProfileResult>("nutanix:index/getNdbProfile:getNdbProfile", args ?? new GetNdbProfileArgs(), options.WithDefaults());

        /// <summary>
        /// Describes a profile in Nutanix Database Service
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Nutanix = Pulumi.Nutanix;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var profile1 = Nutanix.GetNdbProfile.Invoke(new()
        ///     {
        ///         ProfileType = "Network",
        ///         ProfileName = "TEST_NETWORK_PROFILE",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["profile"] = profile1,
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetNdbProfileResult> Invoke(GetNdbProfileInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNdbProfileResult>("nutanix:index/getNdbProfile:getNdbProfile", args ?? new GetNdbProfileInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNdbProfileArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Database engine. For eg. postgres_database
        /// </summary>
        [Input("engine")]
        public string? Engine { get; set; }

        /// <summary>
        /// Profile ID for query
        /// </summary>
        [Input("profileId")]
        public string? ProfileId { get; set; }

        /// <summary>
        /// Profile Name for query
        /// </summary>
        [Input("profileName")]
        public string? ProfileName { get; set; }

        /// <summary>
        /// Profile type. Types: Software, Compute, Network and Database_Parameter
        /// </summary>
        [Input("profileType")]
        public string? ProfileType { get; set; }

        public GetNdbProfileArgs()
        {
        }
        public static new GetNdbProfileArgs Empty => new GetNdbProfileArgs();
    }

    public sealed class GetNdbProfileInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Database engine. For eg. postgres_database
        /// </summary>
        [Input("engine")]
        public Input<string>? Engine { get; set; }

        /// <summary>
        /// Profile ID for query
        /// </summary>
        [Input("profileId")]
        public Input<string>? ProfileId { get; set; }

        /// <summary>
        /// Profile Name for query
        /// </summary>
        [Input("profileName")]
        public Input<string>? ProfileName { get; set; }

        /// <summary>
        /// Profile type. Types: Software, Compute, Network and Database_Parameter
        /// </summary>
        [Input("profileType")]
        public Input<string>? ProfileType { get; set; }

        public GetNdbProfileInvokeArgs()
        {
        }
        public static new GetNdbProfileInvokeArgs Empty => new GetNdbProfileInvokeArgs();
    }


    [OutputType]
    public sealed class GetNdbProfileResult
    {
        /// <summary>
        /// - associated databases
        /// </summary>
        public readonly ImmutableArray<string> AssocDatabases;
        /// <summary>
        /// - associated DB servers
        /// </summary>
        public readonly ImmutableArray<string> AssocDbServers;
        /// <summary>
        /// - list of clusters availability
        /// </summary>
        public readonly ImmutableArray<Outputs.GetNdbProfileClusterAvailabilityResult> ClusterAvailabilities;
        /// <summary>
        /// - database version
        /// </summary>
        public readonly string DbVersion;
        /// <summary>
        /// - description of profile
        /// </summary>
        public readonly string Description;
        public readonly string? Engine;
        /// <summary>
        /// - database engine type
        /// </summary>
        public readonly string EngineType;
        /// <summary>
        /// - id of profile
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// - latest version for engine software
        /// </summary>
        public readonly string LatestVersion;
        /// <summary>
        /// - ID of latest version for engine software
        /// </summary>
        public readonly string LatestVersionId;
        /// <summary>
        /// - profile name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// - era cluster ID
        /// </summary>
        public readonly string NxClusterId;
        /// <summary>
        /// - owner name
        /// </summary>
        public readonly string Owner;
        public readonly string? ProfileId;
        public readonly string? ProfileName;
        public readonly string? ProfileType;
        /// <summary>
        /// - status of profile
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// - if system profile or not
        /// </summary>
        public readonly bool SystemProfile;
        /// <summary>
        /// - topology
        /// </summary>
        public readonly string Topology;
        public readonly string Type;
        /// <summary>
        /// - profile's different version config
        /// </summary>
        public readonly ImmutableArray<Outputs.GetNdbProfileVersionResult> Versions;

        [OutputConstructor]
        private GetNdbProfileResult(
            ImmutableArray<string> assocDatabases,

            ImmutableArray<string> assocDbServers,

            ImmutableArray<Outputs.GetNdbProfileClusterAvailabilityResult> clusterAvailabilities,

            string dbVersion,

            string description,

            string? engine,

            string engineType,

            string id,

            string latestVersion,

            string latestVersionId,

            string name,

            string nxClusterId,

            string owner,

            string? profileId,

            string? profileName,

            string? profileType,

            string status,

            bool systemProfile,

            string topology,

            string type,

            ImmutableArray<Outputs.GetNdbProfileVersionResult> versions)
        {
            AssocDatabases = assocDatabases;
            AssocDbServers = assocDbServers;
            ClusterAvailabilities = clusterAvailabilities;
            DbVersion = dbVersion;
            Description = description;
            Engine = engine;
            EngineType = engineType;
            Id = id;
            LatestVersion = latestVersion;
            LatestVersionId = latestVersionId;
            Name = name;
            NxClusterId = nxClusterId;
            Owner = owner;
            ProfileId = profileId;
            ProfileName = profileName;
            ProfileType = profileType;
            Status = status;
            SystemProfile = systemProfile;
            Topology = topology;
            Type = type;
            Versions = versions;
        }
    }
}