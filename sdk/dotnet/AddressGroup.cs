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
    /// <summary>
    /// Provides a resource to create a address group based on the input parameters.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Nutanix = JacobBMay.Nutanix;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var testAddress = new Nutanix.AddressGroup("testAddress", new()
    ///     {
    ///         Description = "test address groups resource",
    ///         IpAddressBlockLists = new[]
    ///         {
    ///             new Nutanix.Inputs.AddressGroupIpAddressBlockListArgs
    ///             {
    ///                 Ip = "10.0.0.0",
    ///                 PrefixLength = 24,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [NutanixResourceType("nutanix:index/addressGroup:AddressGroup")]
    public partial class AddressGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// - (ReadOnly) Address Group string
        /// </summary>
        [Output("addressGroupString")]
        public Output<string> AddressGroupString { get; private set; } = null!;

        /// <summary>
        /// - (Optional) Description of the service group
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// - (Required) list of IP address blocks with their prefix length
        /// </summary>
        [Output("ipAddressBlockLists")]
        public Output<ImmutableArray<Outputs.AddressGroupIpAddressBlockList>> IpAddressBlockLists { get; private set; } = null!;

        /// <summary>
        /// - (Required) Name of the service group
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a AddressGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AddressGroup(string name, AddressGroupArgs args, CustomResourceOptions? options = null)
            : base("nutanix:index/addressGroup:AddressGroup", name, args ?? new AddressGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AddressGroup(string name, Input<string> id, AddressGroupState? state = null, CustomResourceOptions? options = null)
            : base("nutanix:index/addressGroup:AddressGroup", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AddressGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AddressGroup Get(string name, Input<string> id, AddressGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new AddressGroup(name, id, state, options);
        }
    }

    public sealed class AddressGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// - (Optional) Description of the service group
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("ipAddressBlockLists", required: true)]
        private InputList<Inputs.AddressGroupIpAddressBlockListArgs>? _ipAddressBlockLists;

        /// <summary>
        /// - (Required) list of IP address blocks with their prefix length
        /// </summary>
        public InputList<Inputs.AddressGroupIpAddressBlockListArgs> IpAddressBlockLists
        {
            get => _ipAddressBlockLists ?? (_ipAddressBlockLists = new InputList<Inputs.AddressGroupIpAddressBlockListArgs>());
            set => _ipAddressBlockLists = value;
        }

        /// <summary>
        /// - (Required) Name of the service group
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public AddressGroupArgs()
        {
        }
        public static new AddressGroupArgs Empty => new AddressGroupArgs();
    }

    public sealed class AddressGroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// - (ReadOnly) Address Group string
        /// </summary>
        [Input("addressGroupString")]
        public Input<string>? AddressGroupString { get; set; }

        /// <summary>
        /// - (Optional) Description of the service group
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("ipAddressBlockLists")]
        private InputList<Inputs.AddressGroupIpAddressBlockListGetArgs>? _ipAddressBlockLists;

        /// <summary>
        /// - (Required) list of IP address blocks with their prefix length
        /// </summary>
        public InputList<Inputs.AddressGroupIpAddressBlockListGetArgs> IpAddressBlockLists
        {
            get => _ipAddressBlockLists ?? (_ipAddressBlockLists = new InputList<Inputs.AddressGroupIpAddressBlockListGetArgs>());
            set => _ipAddressBlockLists = value;
        }

        /// <summary>
        /// - (Required) Name of the service group
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public AddressGroupState()
        {
        }
        public static new AddressGroupState Empty => new AddressGroupState();
    }
}