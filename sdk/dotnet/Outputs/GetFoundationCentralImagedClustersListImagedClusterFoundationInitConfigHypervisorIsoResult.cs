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
    public sealed class GetFoundationCentralImagedClustersListImagedClusterFoundationInitConfigHypervisorIsoResult
    {
        public readonly string HypervisorType;
        public readonly string Sha256sum;
        public readonly string Url;

        [OutputConstructor]
        private GetFoundationCentralImagedClustersListImagedClusterFoundationInitConfigHypervisorIsoResult(
            string hypervisorType,

            string sha256sum,

            string url)
        {
            HypervisorType = hypervisorType;
            Sha256sum = sha256sum;
            Url = url;
        }
    }
}
