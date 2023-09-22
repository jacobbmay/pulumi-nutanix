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

    public sealed class NdbDatabaseScaleLcmConfigGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("expiryDetails")]
        private InputList<Inputs.NdbDatabaseScaleLcmConfigExpiryDetailGetArgs>? _expiryDetails;
        public InputList<Inputs.NdbDatabaseScaleLcmConfigExpiryDetailGetArgs> ExpiryDetails
        {
            get => _expiryDetails ?? (_expiryDetails = new InputList<Inputs.NdbDatabaseScaleLcmConfigExpiryDetailGetArgs>());
            set => _expiryDetails = value;
        }

        [Input("postDeleteCommands")]
        private InputList<Inputs.NdbDatabaseScaleLcmConfigPostDeleteCommandGetArgs>? _postDeleteCommands;
        public InputList<Inputs.NdbDatabaseScaleLcmConfigPostDeleteCommandGetArgs> PostDeleteCommands
        {
            get => _postDeleteCommands ?? (_postDeleteCommands = new InputList<Inputs.NdbDatabaseScaleLcmConfigPostDeleteCommandGetArgs>());
            set => _postDeleteCommands = value;
        }

        [Input("preDeleteCommands")]
        private InputList<Inputs.NdbDatabaseScaleLcmConfigPreDeleteCommandGetArgs>? _preDeleteCommands;
        public InputList<Inputs.NdbDatabaseScaleLcmConfigPreDeleteCommandGetArgs> PreDeleteCommands
        {
            get => _preDeleteCommands ?? (_preDeleteCommands = new InputList<Inputs.NdbDatabaseScaleLcmConfigPreDeleteCommandGetArgs>());
            set => _preDeleteCommands = value;
        }

        [Input("refreshDetails")]
        private InputList<Inputs.NdbDatabaseScaleLcmConfigRefreshDetailGetArgs>? _refreshDetails;
        public InputList<Inputs.NdbDatabaseScaleLcmConfigRefreshDetailGetArgs> RefreshDetails
        {
            get => _refreshDetails ?? (_refreshDetails = new InputList<Inputs.NdbDatabaseScaleLcmConfigRefreshDetailGetArgs>());
            set => _refreshDetails = value;
        }

        public NdbDatabaseScaleLcmConfigGetArgs()
        {
        }
        public static new NdbDatabaseScaleLcmConfigGetArgs Empty => new NdbDatabaseScaleLcmConfigGetArgs();
    }
}