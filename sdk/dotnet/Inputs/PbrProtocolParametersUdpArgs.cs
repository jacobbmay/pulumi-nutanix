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

    public sealed class PbrProtocolParametersUdpArgs : global::Pulumi.ResourceArgs
    {
        [Input("destinationPortRangeLists")]
        private InputList<Inputs.PbrProtocolParametersUdpDestinationPortRangeListArgs>? _destinationPortRangeLists;
        public InputList<Inputs.PbrProtocolParametersUdpDestinationPortRangeListArgs> DestinationPortRangeLists
        {
            get => _destinationPortRangeLists ?? (_destinationPortRangeLists = new InputList<Inputs.PbrProtocolParametersUdpDestinationPortRangeListArgs>());
            set => _destinationPortRangeLists = value;
        }

        [Input("sourcePortRangeLists")]
        private InputList<Inputs.PbrProtocolParametersUdpSourcePortRangeListArgs>? _sourcePortRangeLists;
        public InputList<Inputs.PbrProtocolParametersUdpSourcePortRangeListArgs> SourcePortRangeLists
        {
            get => _sourcePortRangeLists ?? (_sourcePortRangeLists = new InputList<Inputs.PbrProtocolParametersUdpSourcePortRangeListArgs>());
            set => _sourcePortRangeLists = value;
        }

        public PbrProtocolParametersUdpArgs()
        {
        }
        public static new PbrProtocolParametersUdpArgs Empty => new PbrProtocolParametersUdpArgs();
    }
}