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
    public static class GetKarbonPrivateRegistry
    {
        /// <summary>
        /// Describes Karbon private registry entry
        /// </summary>
        public static Task<GetKarbonPrivateRegistryResult> InvokeAsync(GetKarbonPrivateRegistryArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetKarbonPrivateRegistryResult>("nutanix:index/getKarbonPrivateRegistry:getKarbonPrivateRegistry", args ?? new GetKarbonPrivateRegistryArgs(), options.WithDefaults());

        /// <summary>
        /// Describes Karbon private registry entry
        /// </summary>
        public static Output<GetKarbonPrivateRegistryResult> Invoke(GetKarbonPrivateRegistryInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetKarbonPrivateRegistryResult>("nutanix:index/getKarbonPrivateRegistry:getKarbonPrivateRegistry", args ?? new GetKarbonPrivateRegistryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetKarbonPrivateRegistryArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Represents karbon private registry uuid
        /// </summary>
        [Input("privateRegistryId")]
        public string? PrivateRegistryId { get; set; }

        /// <summary>
        /// Represents the name of karbon private registry
        /// </summary>
        [Input("privateRegistryName")]
        public string? PrivateRegistryName { get; set; }

        public GetKarbonPrivateRegistryArgs()
        {
        }
        public static new GetKarbonPrivateRegistryArgs Empty => new GetKarbonPrivateRegistryArgs();
    }

    public sealed class GetKarbonPrivateRegistryInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Represents karbon private registry uuid
        /// </summary>
        [Input("privateRegistryId")]
        public Input<string>? PrivateRegistryId { get; set; }

        /// <summary>
        /// Represents the name of karbon private registry
        /// </summary>
        [Input("privateRegistryName")]
        public Input<string>? PrivateRegistryName { get; set; }

        public GetKarbonPrivateRegistryInvokeArgs()
        {
        }
        public static new GetKarbonPrivateRegistryInvokeArgs Empty => new GetKarbonPrivateRegistryInvokeArgs();
    }


    [OutputType]
    public sealed class GetKarbonPrivateRegistryResult
    {
        /// <summary>
        /// - Endpoint of the private in format `url:port`.
        /// </summary>
        public readonly string Endpoint;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// - Name of the private registry.
        /// </summary>
        public readonly string Name;
        public readonly string? PrivateRegistryId;
        public readonly string? PrivateRegistryName;
        /// <summary>
        /// - UUID of the private registry.
        /// </summary>
        public readonly string Uuid;

        [OutputConstructor]
        private GetKarbonPrivateRegistryResult(
            string endpoint,

            string id,

            string name,

            string? privateRegistryId,

            string? privateRegistryName,

            string uuid)
        {
            Endpoint = endpoint;
            Id = id;
            Name = name;
            PrivateRegistryId = privateRegistryId;
            PrivateRegistryName = privateRegistryName;
            Uuid = uuid;
        }
    }
}