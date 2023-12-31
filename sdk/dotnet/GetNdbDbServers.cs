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
    public static class GetNdbDbServers
    {
        /// <summary>
        /// List of all Database Server VM in Nutanix Database Service
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
        ///     var dbservers = Nutanix.GetNdbDbServers.Invoke();
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetNdbDbServersResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNdbDbServersResult>("nutanix:index/getNdbDbServers:getNdbDbServers", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// List of all Database Server VM in Nutanix Database Service
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
        ///     var dbservers = Nutanix.GetNdbDbServers.Invoke();
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetNdbDbServersResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNdbDbServersResult>("nutanix:index/getNdbDbServers:getNdbDbServers", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetNdbDbServersResult
    {
        public readonly ImmutableArray<Outputs.GetNdbDbServersDbserverResult> Dbservers;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetNdbDbServersResult(
            ImmutableArray<Outputs.GetNdbDbServersDbserverResult> dbservers,

            string id)
        {
            Dbservers = dbservers;
            Id = id;
        }
    }
}
