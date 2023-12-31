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
    public sealed class GetPbrSpecResourceProtocolParameterIcmpResult
    {
        public readonly int IcmpCode;
        public readonly int IcmpType;

        [OutputConstructor]
        private GetPbrSpecResourceProtocolParameterIcmpResult(
            int icmpCode,

            int icmpType)
        {
            IcmpCode = icmpCode;
            IcmpType = icmpType;
        }
    }
}
