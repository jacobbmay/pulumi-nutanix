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

    public sealed class NdbProfileDatabaseParameterProfileGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("postgresDatabases")]
        private InputList<Inputs.NdbProfileDatabaseParameterProfilePostgresDatabaseGetArgs>? _postgresDatabases;
        public InputList<Inputs.NdbProfileDatabaseParameterProfilePostgresDatabaseGetArgs> PostgresDatabases
        {
            get => _postgresDatabases ?? (_postgresDatabases = new InputList<Inputs.NdbProfileDatabaseParameterProfilePostgresDatabaseGetArgs>());
            set => _postgresDatabases = value;
        }

        public NdbProfileDatabaseParameterProfileGetArgs()
        {
        }
        public static new NdbProfileDatabaseParameterProfileGetArgs Empty => new NdbProfileDatabaseParameterProfileGetArgs();
    }
}