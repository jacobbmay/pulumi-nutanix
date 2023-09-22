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
    'GetNdbNetworkResult',
    'AwaitableGetNdbNetworkResult',
    'get_ndb_network',
    'get_ndb_network_output',
]

@pulumi.output_type
class GetNdbNetworkResult:
    """
    A collection of values returned by getNdbNetwork.
    """
    def __init__(__self__, cluster_id=None, id=None, ip_addresses=None, ip_pools=None, managed=None, name=None, properties=None, properties_maps=None, stretched_vlan_id=None, type=None):
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip_addresses and not isinstance(ip_addresses, list):
            raise TypeError("Expected argument 'ip_addresses' to be a list")
        pulumi.set(__self__, "ip_addresses", ip_addresses)
        if ip_pools and not isinstance(ip_pools, list):
            raise TypeError("Expected argument 'ip_pools' to be a list")
        pulumi.set(__self__, "ip_pools", ip_pools)
        if managed and not isinstance(managed, bool):
            raise TypeError("Expected argument 'managed' to be a bool")
        pulumi.set(__self__, "managed", managed)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if properties and not isinstance(properties, list):
            raise TypeError("Expected argument 'properties' to be a list")
        pulumi.set(__self__, "properties", properties)
        if properties_maps and not isinstance(properties_maps, list):
            raise TypeError("Expected argument 'properties_maps' to be a list")
        pulumi.set(__self__, "properties_maps", properties_maps)
        if stretched_vlan_id and not isinstance(stretched_vlan_id, str):
            raise TypeError("Expected argument 'stretched_vlan_id' to be a str")
        pulumi.set(__self__, "stretched_vlan_id", stretched_vlan_id)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> str:
        """
        cluster id where network is present
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        network id
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ipAddresses")
    def ip_addresses(self) -> Sequence['outputs.GetNdbNetworkIpAddressResult']:
        """
        IP addresses of network
        """
        return pulumi.get(self, "ip_addresses")

    @property
    @pulumi.getter(name="ipPools")
    def ip_pools(self) -> Sequence['outputs.GetNdbNetworkIpPoolResult']:
        """
        IP Pools of network
        """
        return pulumi.get(self, "ip_pools")

    @property
    @pulumi.getter
    def managed(self) -> bool:
        """
        network managed by NDB or not
        """
        return pulumi.get(self, "managed")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        network name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def properties(self) -> Sequence['outputs.GetNdbNetworkPropertyResult']:
        """
        properties of network
        """
        return pulumi.get(self, "properties")

    @property
    @pulumi.getter(name="propertiesMaps")
    def properties_maps(self) -> Sequence['outputs.GetNdbNetworkPropertiesMapResult']:
        """
        properties map of network
        """
        return pulumi.get(self, "properties_maps")

    @property
    @pulumi.getter(name="stretchedVlanId")
    def stretched_vlan_id(self) -> str:
        """
        stretched vlan id
        """
        return pulumi.get(self, "stretched_vlan_id")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        type of network
        """
        return pulumi.get(self, "type")


class AwaitableGetNdbNetworkResult(GetNdbNetworkResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNdbNetworkResult(
            cluster_id=self.cluster_id,
            id=self.id,
            ip_addresses=self.ip_addresses,
            ip_pools=self.ip_pools,
            managed=self.managed,
            name=self.name,
            properties=self.properties,
            properties_maps=self.properties_maps,
            stretched_vlan_id=self.stretched_vlan_id,
            type=self.type)


def get_ndb_network(id: Optional[str] = None,
                    name: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNdbNetworkResult:
    """
    Describes a network in Nutanix Database Service

    ## Example Usage

    ```python
    import pulumi
    import pulumi_nutanix as nutanix

    nw = nutanix.get_ndb_network(id="{{ id of network }}")
    ```


    :param str id: id of network
    :param str name: name of network
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('nutanix:index/getNdbNetwork:getNdbNetwork', __args__, opts=opts, typ=GetNdbNetworkResult).value

    return AwaitableGetNdbNetworkResult(
        cluster_id=pulumi.get(__ret__, 'cluster_id'),
        id=pulumi.get(__ret__, 'id'),
        ip_addresses=pulumi.get(__ret__, 'ip_addresses'),
        ip_pools=pulumi.get(__ret__, 'ip_pools'),
        managed=pulumi.get(__ret__, 'managed'),
        name=pulumi.get(__ret__, 'name'),
        properties=pulumi.get(__ret__, 'properties'),
        properties_maps=pulumi.get(__ret__, 'properties_maps'),
        stretched_vlan_id=pulumi.get(__ret__, 'stretched_vlan_id'),
        type=pulumi.get(__ret__, 'type'))


@_utilities.lift_output_func(get_ndb_network)
def get_ndb_network_output(id: Optional[pulumi.Input[Optional[str]]] = None,
                           name: Optional[pulumi.Input[Optional[str]]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetNdbNetworkResult]:
    """
    Describes a network in Nutanix Database Service

    ## Example Usage

    ```python
    import pulumi
    import pulumi_nutanix as nutanix

    nw = nutanix.get_ndb_network(id="{{ id of network }}")
    ```


    :param str id: id of network
    :param str name: name of network
    """
    ...