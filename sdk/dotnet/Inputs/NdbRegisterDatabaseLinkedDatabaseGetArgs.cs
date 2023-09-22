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

    public sealed class NdbRegisterDatabaseLinkedDatabaseGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// name of database
        /// </summary>
        [Input("databaseName")]
        public Input<string>? DatabaseName { get; set; }

        /// <summary>
        /// status of database
        /// </summary>
        [Input("databaseStatus")]
        public Input<string>? DatabaseStatus { get; set; }

        /// <summary>
        /// date created for db instance
        /// </summary>
        [Input("dateCreated")]
        public Input<string>? DateCreated { get; set; }

        /// <summary>
        /// date modified for instance
        /// </summary>
        [Input("dateModified")]
        public Input<string>? DateModified { get; set; }

        /// <summary>
        /// description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("infos")]
        private InputList<Inputs.NdbRegisterDatabaseLinkedDatabaseInfoGetArgs>? _infos;

        /// <summary>
        /// info of instance
        /// </summary>
        public InputList<Inputs.NdbRegisterDatabaseLinkedDatabaseInfoGetArgs> Infos
        {
            get => _infos ?? (_infos = new InputList<Inputs.NdbRegisterDatabaseLinkedDatabaseInfoGetArgs>());
            set => _infos = value;
        }

        [Input("metric")]
        private InputMap<string>? _metric;

        /// <summary>
        /// Stores storage info regarding size, allocatedSize, usedSize and unit of calculation that seems to have been fetched from PRISM.
        /// </summary>
        public InputMap<string> Metric
        {
            get => _metric ?? (_metric = new InputMap<string>());
            set => _metric = value;
        }

        /// <summary>
        /// - (Required) name of time machine
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// parent database id
        /// </summary>
        [Input("parentDatabaseId")]
        public Input<string>? ParentDatabaseId { get; set; }

        [Input("parentLinkedDatabaseId")]
        public Input<string>? ParentLinkedDatabaseId { get; set; }

        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

        /// <summary>
        /// status of instance
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("timezone")]
        public Input<string>? Timezone { get; set; }

        public NdbRegisterDatabaseLinkedDatabaseGetArgs()
        {
        }
        public static new NdbRegisterDatabaseLinkedDatabaseGetArgs Empty => new NdbRegisterDatabaseLinkedDatabaseGetArgs();
    }
}