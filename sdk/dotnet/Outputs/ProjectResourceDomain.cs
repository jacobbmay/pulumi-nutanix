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
    public sealed class ProjectResourceDomain
    {
        public readonly ImmutableArray<Outputs.ProjectResourceDomainResource> Resources;

        [OutputConstructor]
        private ProjectResourceDomain(ImmutableArray<Outputs.ProjectResourceDomainResource> resources)
        {
            Resources = resources;
        }
    }
}