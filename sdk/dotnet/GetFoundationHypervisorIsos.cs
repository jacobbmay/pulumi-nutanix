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
    public static class GetFoundationHypervisorIsos
    {
        /// <summary>
        /// Describes a list of hypervisor isos image file details present in foundation vm
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
        ///     var hypervisorIsos = Nutanix.GetFoundationHypervisorIsos.Invoke();
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Note
        /// 
        /// * This data source only lists .iso files details.
        /// 
        /// See detailed information in [Nutanix Foundation Hypervisor Isos](https://www.nutanix.dev/api_references/foundation/#/b3A6MjIyMjM0MDE-list-hypervisor-images-available-in-foundation).
        /// </summary>
        public static Task<GetFoundationHypervisorIsosResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFoundationHypervisorIsosResult>("nutanix:index/getFoundationHypervisorIsos:getFoundationHypervisorIsos", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// Describes a list of hypervisor isos image file details present in foundation vm
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
        ///     var hypervisorIsos = Nutanix.GetFoundationHypervisorIsos.Invoke();
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Note
        /// 
        /// * This data source only lists .iso files details.
        /// 
        /// See detailed information in [Nutanix Foundation Hypervisor Isos](https://www.nutanix.dev/api_references/foundation/#/b3A6MjIyMjM0MDE-list-hypervisor-images-available-in-foundation).
        /// </summary>
        public static Output<GetFoundationHypervisorIsosResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFoundationHypervisorIsosResult>("nutanix:index/getFoundationHypervisorIsos:getFoundationHypervisorIsos", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetFoundationHypervisorIsosResult
    {
        /// <summary>
        /// List of esx isos and theirdetails present in foundation vm
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFoundationHypervisorIsosEsxResult> Esxes;
        /// <summary>
        /// List of hyperv isos and their details present in foundation vm
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFoundationHypervisorIsosHypervResult> Hypervs;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// List of kvm isos and their details present in foundation vm
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFoundationHypervisorIsosKvmResult> Kvms;
        /// <summary>
        /// List of linux isos and their details present in foundation vm
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFoundationHypervisorIsosLinuxResult> Linuxes;
        /// <summary>
        /// List of esx isos and theirdetails present in foundation vm
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFoundationHypervisorIsosXenResult> Xens;

        [OutputConstructor]
        private GetFoundationHypervisorIsosResult(
            ImmutableArray<Outputs.GetFoundationHypervisorIsosEsxResult> esxes,

            ImmutableArray<Outputs.GetFoundationHypervisorIsosHypervResult> hypervs,

            string id,

            ImmutableArray<Outputs.GetFoundationHypervisorIsosKvmResult> kvms,

            ImmutableArray<Outputs.GetFoundationHypervisorIsosLinuxResult> linuxes,

            ImmutableArray<Outputs.GetFoundationHypervisorIsosXenResult> xens)
        {
            Esxes = esxes;
            Hypervs = hypervs;
            Id = id;
            Kvms = kvms;
            Linuxes = linuxes;
            Xens = xens;
        }
    }
}