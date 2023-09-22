# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = [
    'GetVpcsResult',
    'AwaitableGetVpcsResult',
    'get_vpcs',
    'get_vpcs_output',
]

@pulumi.output_type
class GetVpcsResult:
    """
    A collection of values returned by getVpcs.
    """
    def __init__(__self__, api_version=None, entities=None, id=None, metadatas=None):
        if api_version and not isinstance(api_version, str):
            raise TypeError("Expected argument 'api_version' to be a str")
        pulumi.set(__self__, "api_version", api_version)
        if entities and not isinstance(entities, list):
            raise TypeError("Expected argument 'entities' to be a list")
        pulumi.set(__self__, "entities", entities)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if metadatas and not isinstance(metadatas, list):
            raise TypeError("Expected argument 'metadatas' to be a list")
        pulumi.set(__self__, "metadatas", metadatas)

    @property
    @pulumi.getter(name="apiVersion")
    def api_version(self) -> str:
        """
        version of the API
        """
        return pulumi.get(self, "api_version")

    @property
    @pulumi.getter
    def entities(self) -> Sequence['outputs.GetVpcsEntityResult']:
        """
        List of VPCs
        """
        return pulumi.get(self, "entities")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def metadatas(self) -> Sequence['outputs.GetVpcsMetadataResult']:
        """
        - The vpc kind metadata.
        """
        return pulumi.get(self, "metadatas")


class AwaitableGetVpcsResult(GetVpcsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVpcsResult(
            api_version=self.api_version,
            entities=self.entities,
            id=self.id,
            metadatas=self.metadatas)


def get_vpcs(metadatas: Optional[Sequence[pulumi.InputType['GetVpcsMetadataArgs']]] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVpcsResult:
    """
    Provides a datasource to retrieve all the vpcs.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_nutanix as nutanix

    test = nutanix.get_vpcs()
    ```


    :param Sequence[pulumi.InputType['GetVpcsMetadataArgs']] metadatas: - The vpc kind metadata.
    """
    __args__ = dict()
    __args__['metadatas'] = metadatas
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('nutanix:index/getVpcs:getVpcs', __args__, opts=opts, typ=GetVpcsResult).value

    return AwaitableGetVpcsResult(
        api_version=pulumi.get(__ret__, 'api_version'),
        entities=pulumi.get(__ret__, 'entities'),
        id=pulumi.get(__ret__, 'id'),
        metadatas=pulumi.get(__ret__, 'metadatas'))


@_utilities.lift_output_func(get_vpcs)
def get_vpcs_output(metadatas: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetVpcsMetadataArgs']]]]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVpcsResult]:
    """
    Provides a datasource to retrieve all the vpcs.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_nutanix as nutanix

    test = nutanix.get_vpcs()
    ```


    :param Sequence[pulumi.InputType['GetVpcsMetadataArgs']] metadatas: - The vpc kind metadata.
    """
    ...