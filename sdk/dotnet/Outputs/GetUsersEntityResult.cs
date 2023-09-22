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
    public sealed class GetUsersEntityResult
    {
        /// <summary>
        /// - List of ACP references. See #reference for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetUsersEntityAccessControlPolicyReferenceListResult> AccessControlPolicyReferenceLists;
        /// <summary>
        /// The version of the API.
        /// </summary>
        public readonly string ApiVersion;
        /// <summary>
        /// - (Optional) Categories for the user.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetUsersEntityCategoryResult> Categories;
        /// <summary>
        /// - (Optional) The directory service user configuration. See below for more information.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetUsersEntityDirectoryServiceUserResult> DirectoryServiceUsers;
        /// <summary>
        /// - The display name of the user (common name) provided by the directory service.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// - (Optional) (Optional) The identity provider user configuration. See below for more information.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetUsersEntityIdentityProviderUserResult> IdentityProviderUsers;
        /// <summary>
        /// - The user kind metadata.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Metadata;
        /// <summary>
        /// - the name(Optional).
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// - (Optional) The reference to a user.
        /// </summary>
        public readonly ImmutableDictionary<string, string> OwnerReference;
        /// <summary>
        /// - (Optional) The reference to a project.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? ProjectReference;
        /// <summary>
        /// - A list of projects the user is part of. See #reference for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetUsersEntityProjectReferenceListResult> ProjectReferenceLists;
        /// <summary>
        /// - The state of the entity.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// - The name of the user.
        /// </summary>
        public readonly string UserType;

        [OutputConstructor]
        private GetUsersEntityResult(
            ImmutableArray<Outputs.GetUsersEntityAccessControlPolicyReferenceListResult> accessControlPolicyReferenceLists,

            string apiVersion,

            ImmutableArray<Outputs.GetUsersEntityCategoryResult> categories,

            ImmutableArray<Outputs.GetUsersEntityDirectoryServiceUserResult> directoryServiceUsers,

            string displayName,

            ImmutableArray<Outputs.GetUsersEntityIdentityProviderUserResult> identityProviderUsers,

            ImmutableDictionary<string, string> metadata,

            string name,

            ImmutableDictionary<string, string> ownerReference,

            ImmutableDictionary<string, string>? projectReference,

            ImmutableArray<Outputs.GetUsersEntityProjectReferenceListResult> projectReferenceLists,

            string state,

            string userType)
        {
            AccessControlPolicyReferenceLists = accessControlPolicyReferenceLists;
            ApiVersion = apiVersion;
            Categories = categories;
            DirectoryServiceUsers = directoryServiceUsers;
            DisplayName = displayName;
            IdentityProviderUsers = identityProviderUsers;
            Metadata = metadata;
            Name = name;
            OwnerReference = ownerReference;
            ProjectReference = projectReference;
            ProjectReferenceLists = projectReferenceLists;
            State = state;
            UserType = userType;
        }
    }
}