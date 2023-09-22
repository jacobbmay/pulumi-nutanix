// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace JacobBMay.Nutanix.Inputs
{

    public sealed class UserDirectoryServiceUserArgs : global::Pulumi.ResourceArgs
    {
        [Input("defaultUserPrincipalName")]
        public Input<string>? DefaultUserPrincipalName { get; set; }

        /// <summary>
        /// - (Optional) The reference to a directory service. See #reference for to look the supported attributes.
        /// </summary>
        [Input("directoryServiceReference", required: true)]
        public Input<Inputs.UserDirectoryServiceUserDirectoryServiceReferenceArgs> DirectoryServiceReference { get; set; } = null!;

        /// <summary>
        /// - (Optional) The UserPrincipalName of the user from the directory service.
        /// </summary>
        [Input("userPrincipalName")]
        public Input<string>? UserPrincipalName { get; set; }

        public UserDirectoryServiceUserArgs()
        {
        }
        public static new UserDirectoryServiceUserArgs Empty => new UserDirectoryServiceUserArgs();
    }
}