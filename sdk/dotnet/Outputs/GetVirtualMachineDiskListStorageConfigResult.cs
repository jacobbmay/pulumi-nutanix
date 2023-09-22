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
    public sealed class GetVirtualMachineDiskListStorageConfigResult
    {
        /// <summary>
        /// - State of the storage policy to pin virtual disks to the hot tier. When specified as a VM attribute, the storage policy applies to all virtual disks of the VM unless overridden by the same attribute specified for a virtual disk.
        /// </summary>
        public readonly string FlashMode;
        /// <summary>
        /// - Reference to a kind. Either one of (kind, uuid) or url needs to be specified.
        /// * `storage_container_reference.#.url`: - GET query on the URL will provide information on the source.
        /// * `storage_container_reference.#.kind`: - kind of the container reference
        /// * `storage_container_reference.#.name`: - name of the container reference
        /// * `storage_container_reference.#.uuid`: - uiid of the container reference
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVirtualMachineDiskListStorageConfigStorageContainerReferenceResult> StorageContainerReferences;

        [OutputConstructor]
        private GetVirtualMachineDiskListStorageConfigResult(
            string flashMode,

            ImmutableArray<Outputs.GetVirtualMachineDiskListStorageConfigStorageContainerReferenceResult> storageContainerReferences)
        {
            FlashMode = flashMode;
            StorageContainerReferences = storageContainerReferences;
        }
    }
}