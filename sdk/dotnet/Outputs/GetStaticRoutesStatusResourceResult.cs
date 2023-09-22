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
    public sealed class GetStaticRoutesStatusResourceResult
    {
        /// <summary>
        /// default route. (present in status resource only )
        /// </summary>
        public readonly ImmutableArray<Outputs.GetStaticRoutesStatusResourceDefaultRouteResult> DefaultRoutes;
        /// <summary>
        /// list of dynamic routes (present in status resource only)
        /// </summary>
        public readonly ImmutableArray<Outputs.GetStaticRoutesStatusResourceDynamicRoutesListResult> DynamicRoutesLists;
        /// <summary>
        /// list of local routes (present in status resource only )
        /// </summary>
        public readonly ImmutableArray<Outputs.GetStaticRoutesStatusResourceLocalRoutesListResult> LocalRoutesLists;
        /// <summary>
        /// list of static routes
        /// </summary>
        public readonly ImmutableArray<Outputs.GetStaticRoutesStatusResourceStaticRoutesListResult> StaticRoutesLists;

        [OutputConstructor]
        private GetStaticRoutesStatusResourceResult(
            ImmutableArray<Outputs.GetStaticRoutesStatusResourceDefaultRouteResult> defaultRoutes,

            ImmutableArray<Outputs.GetStaticRoutesStatusResourceDynamicRoutesListResult> dynamicRoutesLists,

            ImmutableArray<Outputs.GetStaticRoutesStatusResourceLocalRoutesListResult> localRoutesLists,

            ImmutableArray<Outputs.GetStaticRoutesStatusResourceStaticRoutesListResult> staticRoutesLists)
        {
            DefaultRoutes = defaultRoutes;
            DynamicRoutesLists = dynamicRoutesLists;
            LocalRoutesLists = localRoutesLists;
            StaticRoutesLists = staticRoutesLists;
        }
    }
}