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

__all__ = [
    'GetAddressGroupResult',
    'AwaitableGetAddressGroupResult',
    'get_address_group',
    'get_address_group_output',
]

@pulumi.output_type
class GetAddressGroupResult:
    """
    A collection of values returned by getAddressGroup.
    """
    def __init__(__self__, address_group_string=None, description=None, id=None, ip_address_block_lists=None, name=None, uuid=None):
        if address_group_string and not isinstance(address_group_string, str):
            raise TypeError("Expected argument 'address_group_string' to be a str")
        pulumi.set(__self__, "address_group_string", address_group_string)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip_address_block_lists and not isinstance(ip_address_block_lists, list):
            raise TypeError("Expected argument 'ip_address_block_lists' to be a list")
        pulumi.set(__self__, "ip_address_block_lists", ip_address_block_lists)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if uuid and not isinstance(uuid, str):
            raise TypeError("Expected argument 'uuid' to be a str")
        pulumi.set(__self__, "uuid", uuid)

    @property
    @pulumi.getter(name="addressGroupString")
    def address_group_string(self) -> str:
        """
        - (ReadOnly) Address Group string
        """
        return pulumi.get(self, "address_group_string")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        - (ReadOnly) Description of the address group
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ipAddressBlockLists")
    def ip_address_block_lists(self) -> Sequence['outputs.GetAddressGroupIpAddressBlockListResult']:
        """
        - (ReadOnly) list of IP address blocks with their prefix length
        """
        return pulumi.get(self, "ip_address_block_lists")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        - (ReadOnly) Name of the address group
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def uuid(self) -> str:
        """
        - (Required) UUID of the address group
        """
        return pulumi.get(self, "uuid")


class AwaitableGetAddressGroupResult(GetAddressGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAddressGroupResult(
            address_group_string=self.address_group_string,
            description=self.description,
            id=self.id,
            ip_address_block_lists=self.ip_address_block_lists,
            name=self.name,
            uuid=self.uuid)


def get_address_group(uuid: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAddressGroupResult:
    """
    Provides a datasource to retrieve a address group.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_nutanix as nutanix

    test_address = nutanix.AddressGroup("testAddress",
        description="test address groups resource",
        ip_address_block_lists=[nutanix.AddressGroupIpAddressBlockListArgs(
            ip="10.0.0.0",
            prefix_length=24,
        )])
    addr_group = nutanix.get_address_group_output(uuid=test_address.id)
    ```


    :param str uuid: - (Required) UUID of the address group
    """
    __args__ = dict()
    __args__['uuid'] = uuid
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('nutanix:index/getAddressGroup:getAddressGroup', __args__, opts=opts, typ=GetAddressGroupResult).value

    return AwaitableGetAddressGroupResult(
        address_group_string=pulumi.get(__ret__, 'address_group_string'),
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        ip_address_block_lists=pulumi.get(__ret__, 'ip_address_block_lists'),
        name=pulumi.get(__ret__, 'name'),
        uuid=pulumi.get(__ret__, 'uuid'))


@_utilities.lift_output_func(get_address_group)
def get_address_group_output(uuid: Optional[pulumi.Input[str]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAddressGroupResult]:
    """
    Provides a datasource to retrieve a address group.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_nutanix as nutanix

    test_address = nutanix.AddressGroup("testAddress",
        description="test address groups resource",
        ip_address_block_lists=[nutanix.AddressGroupIpAddressBlockListArgs(
            ip="10.0.0.0",
            prefix_length=24,
        )])
    addr_group = nutanix.get_address_group_output(uuid=test_address.id)
    ```


    :param str uuid: - (Required) UUID of the address group
    """
    ...