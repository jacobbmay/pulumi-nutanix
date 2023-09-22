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
    'GetAccessControlPolicyResult',
    'AwaitableGetAccessControlPolicyResult',
    'get_access_control_policy',
    'get_access_control_policy_output',
]

@pulumi.output_type
class GetAccessControlPolicyResult:
    """
    A collection of values returned by getAccessControlPolicy.
    """
    def __init__(__self__, access_control_policy_id=None, access_control_policy_name=None, api_version=None, categories=None, context_filter_lists=None, description=None, id=None, metadata=None, name=None, owner_reference=None, project_reference=None, role_references=None, state=None, user_group_reference_lists=None, user_reference_lists=None):
        if access_control_policy_id and not isinstance(access_control_policy_id, str):
            raise TypeError("Expected argument 'access_control_policy_id' to be a str")
        pulumi.set(__self__, "access_control_policy_id", access_control_policy_id)
        if access_control_policy_name and not isinstance(access_control_policy_name, str):
            raise TypeError("Expected argument 'access_control_policy_name' to be a str")
        pulumi.set(__self__, "access_control_policy_name", access_control_policy_name)
        if api_version and not isinstance(api_version, str):
            raise TypeError("Expected argument 'api_version' to be a str")
        pulumi.set(__self__, "api_version", api_version)
        if categories and not isinstance(categories, list):
            raise TypeError("Expected argument 'categories' to be a list")
        pulumi.set(__self__, "categories", categories)
        if context_filter_lists and not isinstance(context_filter_lists, list):
            raise TypeError("Expected argument 'context_filter_lists' to be a list")
        pulumi.set(__self__, "context_filter_lists", context_filter_lists)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if metadata and not isinstance(metadata, dict):
            raise TypeError("Expected argument 'metadata' to be a dict")
        pulumi.set(__self__, "metadata", metadata)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if owner_reference and not isinstance(owner_reference, dict):
            raise TypeError("Expected argument 'owner_reference' to be a dict")
        pulumi.set(__self__, "owner_reference", owner_reference)
        if project_reference and not isinstance(project_reference, dict):
            raise TypeError("Expected argument 'project_reference' to be a dict")
        pulumi.set(__self__, "project_reference", project_reference)
        if role_references and not isinstance(role_references, list):
            raise TypeError("Expected argument 'role_references' to be a list")
        pulumi.set(__self__, "role_references", role_references)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if user_group_reference_lists and not isinstance(user_group_reference_lists, list):
            raise TypeError("Expected argument 'user_group_reference_lists' to be a list")
        pulumi.set(__self__, "user_group_reference_lists", user_group_reference_lists)
        if user_reference_lists and not isinstance(user_reference_lists, list):
            raise TypeError("Expected argument 'user_reference_lists' to be a list")
        pulumi.set(__self__, "user_reference_lists", user_reference_lists)

    @property
    @pulumi.getter(name="accessControlPolicyId")
    def access_control_policy_id(self) -> Optional[str]:
        return pulumi.get(self, "access_control_policy_id")

    @property
    @pulumi.getter(name="accessControlPolicyName")
    def access_control_policy_name(self) -> Optional[str]:
        return pulumi.get(self, "access_control_policy_name")

    @property
    @pulumi.getter(name="apiVersion")
    def api_version(self) -> str:
        """
        The version of the API.
        """
        return pulumi.get(self, "api_version")

    @property
    @pulumi.getter
    def categories(self) -> Sequence['outputs.GetAccessControlPolicyCategoryResult']:
        """
        - The category values represented as a dictionary of key > list of values.
        """
        return pulumi.get(self, "categories")

    @property
    @pulumi.getter(name="contextFilterLists")
    def context_filter_lists(self) -> Sequence['outputs.GetAccessControlPolicyContextFilterListResult']:
        return pulumi.get(self, "context_filter_lists")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        - The description of the Access Control Policy.
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
    @pulumi.getter
    def metadata(self) -> Mapping[str, str]:
        """
        - The Access Control Policy kind metadata.
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        - the name(Optional).
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="ownerReference")
    def owner_reference(self) -> Mapping[str, str]:
        """
        - The reference to a user.
        """
        return pulumi.get(self, "owner_reference")

    @property
    @pulumi.getter(name="projectReference")
    def project_reference(self) -> Mapping[str, str]:
        """
        - The reference to a project.
        """
        return pulumi.get(self, "project_reference")

    @property
    @pulumi.getter(name="roleReferences")
    def role_references(self) -> Sequence['outputs.GetAccessControlPolicyRoleReferenceResult']:
        """
        - The reference to a role.
        """
        return pulumi.get(self, "role_references")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        - The state of the Access Control Policy.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="userGroupReferenceLists")
    def user_group_reference_lists(self) -> Sequence['outputs.GetAccessControlPolicyUserGroupReferenceListResult']:
        """
        - The User group(s) being assigned a given role.
        """
        return pulumi.get(self, "user_group_reference_lists")

    @property
    @pulumi.getter(name="userReferenceLists")
    def user_reference_lists(self) -> Sequence['outputs.GetAccessControlPolicyUserReferenceListResult']:
        """
        - The User(s) being assigned a given role.
        """
        return pulumi.get(self, "user_reference_lists")


class AwaitableGetAccessControlPolicyResult(GetAccessControlPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAccessControlPolicyResult(
            access_control_policy_id=self.access_control_policy_id,
            access_control_policy_name=self.access_control_policy_name,
            api_version=self.api_version,
            categories=self.categories,
            context_filter_lists=self.context_filter_lists,
            description=self.description,
            id=self.id,
            metadata=self.metadata,
            name=self.name,
            owner_reference=self.owner_reference,
            project_reference=self.project_reference,
            role_references=self.role_references,
            state=self.state,
            user_group_reference_lists=self.user_group_reference_lists,
            user_reference_lists=self.user_reference_lists)


def get_access_control_policy(access_control_policy_id: Optional[str] = None,
                              access_control_policy_name: Optional[str] = None,
                              categories: Optional[Sequence[pulumi.InputType['GetAccessControlPolicyCategoryArgs']]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAccessControlPolicyResult:
    """
    Describes an Access Control Policy.


    :param str access_control_policy_id: - (Required) The UUID of an access control policy.
    :param Sequence[pulumi.InputType['GetAccessControlPolicyCategoryArgs']] categories: - The category values represented as a dictionary of key > list of values.
    """
    __args__ = dict()
    __args__['accessControlPolicyId'] = access_control_policy_id
    __args__['accessControlPolicyName'] = access_control_policy_name
    __args__['categories'] = categories
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('nutanix:index/getAccessControlPolicy:getAccessControlPolicy', __args__, opts=opts, typ=GetAccessControlPolicyResult).value

    return AwaitableGetAccessControlPolicyResult(
        access_control_policy_id=pulumi.get(__ret__, 'access_control_policy_id'),
        access_control_policy_name=pulumi.get(__ret__, 'access_control_policy_name'),
        api_version=pulumi.get(__ret__, 'api_version'),
        categories=pulumi.get(__ret__, 'categories'),
        context_filter_lists=pulumi.get(__ret__, 'context_filter_lists'),
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        metadata=pulumi.get(__ret__, 'metadata'),
        name=pulumi.get(__ret__, 'name'),
        owner_reference=pulumi.get(__ret__, 'owner_reference'),
        project_reference=pulumi.get(__ret__, 'project_reference'),
        role_references=pulumi.get(__ret__, 'role_references'),
        state=pulumi.get(__ret__, 'state'),
        user_group_reference_lists=pulumi.get(__ret__, 'user_group_reference_lists'),
        user_reference_lists=pulumi.get(__ret__, 'user_reference_lists'))


@_utilities.lift_output_func(get_access_control_policy)
def get_access_control_policy_output(access_control_policy_id: Optional[pulumi.Input[Optional[str]]] = None,
                                     access_control_policy_name: Optional[pulumi.Input[Optional[str]]] = None,
                                     categories: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetAccessControlPolicyCategoryArgs']]]]] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAccessControlPolicyResult]:
    """
    Describes an Access Control Policy.


    :param str access_control_policy_id: - (Required) The UUID of an access control policy.
    :param Sequence[pulumi.InputType['GetAccessControlPolicyCategoryArgs']] categories: - The category values represented as a dictionary of key > list of values.
    """
    ...