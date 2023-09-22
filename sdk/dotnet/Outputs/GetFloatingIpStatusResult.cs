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
    public sealed class GetFloatingIpStatusResult
    {
        /// <summary>
        /// Execution Context of Floating IP.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFloatingIpStatusExecutionContextResult> ExecutionContexts;
        /// <summary>
        /// - the name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Floating IP allocation status.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFloatingIpStatusResourceResult> Resources;
        /// <summary>
        /// The state of the floating_ip.
        /// </summary>
        public readonly string State;

        [OutputConstructor]
        private GetFloatingIpStatusResult(
            ImmutableArray<Outputs.GetFloatingIpStatusExecutionContextResult> executionContexts,

            string name,

            ImmutableArray<Outputs.GetFloatingIpStatusResourceResult> resources,

            string state)
        {
            ExecutionContexts = executionContexts;
            Name = name;
            Resources = resources;
            State = state;
        }
    }
}