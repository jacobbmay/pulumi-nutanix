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

    public sealed class NdbClusterEntityCountEngineCountOracleDatabaseGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("profiles")]
        private InputList<Inputs.NdbClusterEntityCountEngineCountOracleDatabaseProfileGetArgs>? _profiles;
        public InputList<Inputs.NdbClusterEntityCountEngineCountOracleDatabaseProfileGetArgs> Profiles
        {
            get => _profiles ?? (_profiles = new InputList<Inputs.NdbClusterEntityCountEngineCountOracleDatabaseProfileGetArgs>());
            set => _profiles = value;
        }

        [Input("timeMachines")]
        public Input<int>? TimeMachines { get; set; }

        public NdbClusterEntityCountEngineCountOracleDatabaseGetArgs()
        {
        }
        public static new NdbClusterEntityCountEngineCountOracleDatabaseGetArgs Empty => new NdbClusterEntityCountEngineCountOracleDatabaseGetArgs();
    }
}