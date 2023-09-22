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
    public sealed class GetNdbProfileVersionVersionClusterAssociationResult
    {
        public readonly string DateCreated;
        public readonly string DateModified;
        /// <summary>
        /// - era cluster ID
        /// </summary>
        public readonly string NxClusterId;
        public readonly bool OptimizedForProvisioning;
        public readonly string OwnerId;
        public readonly string ProfileVersionId;
        public readonly ImmutableArray<Outputs.GetNdbProfileVersionVersionClusterAssociationPropertyResult> Properties;
        /// <summary>
        /// - status of profile
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetNdbProfileVersionVersionClusterAssociationResult(
            string dateCreated,

            string dateModified,

            string nxClusterId,

            bool optimizedForProvisioning,

            string ownerId,

            string profileVersionId,

            ImmutableArray<Outputs.GetNdbProfileVersionVersionClusterAssociationPropertyResult> properties,

            string status)
        {
            DateCreated = dateCreated;
            DateModified = dateModified;
            NxClusterId = nxClusterId;
            OptimizedForProvisioning = optimizedForProvisioning;
            OwnerId = ownerId;
            ProfileVersionId = profileVersionId;
            Properties = properties;
            Status = status;
        }
    }
}