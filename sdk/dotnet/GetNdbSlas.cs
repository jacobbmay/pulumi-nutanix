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
    public static class GetNdbSlas
    {
        /// <summary>
        /// Lists all SLAs in Nutanix Database Service
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
        ///     var slas = Nutanix.GetNdbSlas.Invoke();
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["sla"] = slas,
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetNdbSlasResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNdbSlasResult>("nutanix:index/getNdbSlas:getNdbSlas", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// Lists all SLAs in Nutanix Database Service
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
        ///     var slas = Nutanix.GetNdbSlas.Invoke();
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["sla"] = slas,
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetNdbSlasResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNdbSlasResult>("nutanix:index/getNdbSlas:getNdbSlas", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetNdbSlasResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// - list of slas
        /// </summary>
        public readonly ImmutableArray<Outputs.GetNdbSlasSlaResult> Slas;

        [OutputConstructor]
        private GetNdbSlasResult(
            string id,

            ImmutableArray<Outputs.GetNdbSlasSlaResult> slas)
        {
            Id = id;
            Slas = slas;
        }
    }
}