# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetPolicyTypeResult',
    'AwaitableGetPolicyTypeResult',
    'get_policy_type',
    'get_policy_type_output',
]

@pulumi.output_type
class GetPolicyTypeResult:
    """
    A collection of values returned by getPolicyType.
    """
    def __init__(__self__, creation_time=None, description=None, enabled=None, id=None, modified_by=None, modified_time=None, name=None, policy_type=None, rules=None, sorted=None):
        if creation_time and not isinstance(creation_time, str):
            raise TypeError("Expected argument 'creation_time' to be a str")
        pulumi.set(__self__, "creation_time", creation_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if modified_by and not isinstance(modified_by, str):
            raise TypeError("Expected argument 'modified_by' to be a str")
        pulumi.set(__self__, "modified_by", modified_by)
        if modified_time and not isinstance(modified_time, str):
            raise TypeError("Expected argument 'modified_time' to be a str")
        pulumi.set(__self__, "modified_time", modified_time)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if policy_type and not isinstance(policy_type, str):
            raise TypeError("Expected argument 'policy_type' to be a str")
        pulumi.set(__self__, "policy_type", policy_type)
        if rules and not isinstance(rules, list):
            raise TypeError("Expected argument 'rules' to be a list")
        pulumi.set(__self__, "rules", rules)
        if sorted and not isinstance(sorted, bool):
            raise TypeError("Expected argument 'sorted' to be a bool")
        pulumi.set(__self__, "sorted", sorted)

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> str:
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="modifiedBy")
    def modified_by(self) -> str:
        return pulumi.get(self, "modified_by")

    @property
    @pulumi.getter(name="modifiedTime")
    def modified_time(self) -> str:
        return pulumi.get(self, "modified_time")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="policyType")
    def policy_type(self) -> str:
        return pulumi.get(self, "policy_type")

    @property
    @pulumi.getter
    def rules(self) -> Sequence['outputs.GetPolicyTypeRuleResult']:
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter
    def sorted(self) -> bool:
        return pulumi.get(self, "sorted")


class AwaitableGetPolicyTypeResult(GetPolicyTypeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPolicyTypeResult(
            creation_time=self.creation_time,
            description=self.description,
            enabled=self.enabled,
            id=self.id,
            modified_by=self.modified_by,
            modified_time=self.modified_time,
            name=self.name,
            policy_type=self.policy_type,
            rules=self.rules,
            sorted=self.sorted)


def get_policy_type(id: Optional[str] = None,
                    policy_type: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPolicyTypeResult:
    """
    Use the **zpa_policy_type** data source to get information about an a ``policy_set_id`` and ``policy_type``. This data source is required when creating:

    1. Access policy Rules
    2. Access policy timeout rules
    3. Access policy forwarding rules
    4. Access policy inspection rules

    > **NOTE** The parameters ``policy_set_id`` is required in all circumstances and is exported when checking for the policy_type parameter. The policy_type value is used for differentiating the policy types, in the request endpoint. The supported values are:

    * ``ACCESS_POLICY/GLOBAL_POLICY``
    * ``TIMEOUT_POLICY/REAUTH_POLICY``
    * ``BYPASS_POLICY/CLIENT_FORWARDING_POLICY``
    * ``INSPECTION_POLICY``

    ## Example Usage

    ```python
    import pulumi
    import pulumi_zpa as zpa

    access_policy = zpa.PolicyType.get_policy_type(policy_type="ACCESS_POLICY")
    pulumi.export("zpaPolicyTypeAccessPolicy", access_policy.id)
    ```

    ```python
    import pulumi
    import pulumi_zpa as zpa

    global_policy = zpa.PolicyType.get_policy_type(policy_type="GLOBAL_POLICY")
    pulumi.export("zpaPolicyTypeAccessPolicy", global_policy.id)
    ```

    ```python
    import pulumi
    import pulumi_zpa as zpa

    timeout_policy = zpa.PolicyType.get_policy_type(policy_type="TIMEOUT_POLICY")
    pulumi.export("zpaPolicyTypeTimeoutPolicy", timeout_policy.id)
    ```

    ```python
    import pulumi
    import pulumi_zpa as zpa

    reauth_policy = zpa.PolicyType.get_policy_type(policy_type="REAUTH_POLICY")
    pulumi.export("zpaPolicyTypeReauthPolicy", reauth_policy.id)
    ```

    ```python
    import pulumi
    import pulumi_zpa as zpa

    client_forwarding_policy = zpa.PolicyType.get_policy_type(policy_type="CLIENT_FORWARDING_POLICY")
    pulumi.export("zpaPolicyTypeClientForwardingPolicy", client_forwarding_policy.id)
    ```

    ```python
    import pulumi
    import pulumi_zpa as zpa

    inspection_policy = zpa.PolicyType.get_policy_type(policy_type="INSPECTION_POLICY")
    pulumi.export("zpaPolicyTypeInspectionPolicy", inspection_policy.id)
    ```


    :param str policy_type: The value for differentiating the policy types.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['policyType'] = policy_type
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('zpa:PolicyType/getPolicyType:getPolicyType', __args__, opts=opts, typ=GetPolicyTypeResult).value

    return AwaitableGetPolicyTypeResult(
        creation_time=__ret__.creation_time,
        description=__ret__.description,
        enabled=__ret__.enabled,
        id=__ret__.id,
        modified_by=__ret__.modified_by,
        modified_time=__ret__.modified_time,
        name=__ret__.name,
        policy_type=__ret__.policy_type,
        rules=__ret__.rules,
        sorted=__ret__.sorted)


@_utilities.lift_output_func(get_policy_type)
def get_policy_type_output(id: Optional[pulumi.Input[Optional[str]]] = None,
                           policy_type: Optional[pulumi.Input[Optional[str]]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPolicyTypeResult]:
    """
    Use the **zpa_policy_type** data source to get information about an a ``policy_set_id`` and ``policy_type``. This data source is required when creating:

    1. Access policy Rules
    2. Access policy timeout rules
    3. Access policy forwarding rules
    4. Access policy inspection rules

    > **NOTE** The parameters ``policy_set_id`` is required in all circumstances and is exported when checking for the policy_type parameter. The policy_type value is used for differentiating the policy types, in the request endpoint. The supported values are:

    * ``ACCESS_POLICY/GLOBAL_POLICY``
    * ``TIMEOUT_POLICY/REAUTH_POLICY``
    * ``BYPASS_POLICY/CLIENT_FORWARDING_POLICY``
    * ``INSPECTION_POLICY``

    ## Example Usage

    ```python
    import pulumi
    import pulumi_zpa as zpa

    access_policy = zpa.PolicyType.get_policy_type(policy_type="ACCESS_POLICY")
    pulumi.export("zpaPolicyTypeAccessPolicy", access_policy.id)
    ```

    ```python
    import pulumi
    import pulumi_zpa as zpa

    global_policy = zpa.PolicyType.get_policy_type(policy_type="GLOBAL_POLICY")
    pulumi.export("zpaPolicyTypeAccessPolicy", global_policy.id)
    ```

    ```python
    import pulumi
    import pulumi_zpa as zpa

    timeout_policy = zpa.PolicyType.get_policy_type(policy_type="TIMEOUT_POLICY")
    pulumi.export("zpaPolicyTypeTimeoutPolicy", timeout_policy.id)
    ```

    ```python
    import pulumi
    import pulumi_zpa as zpa

    reauth_policy = zpa.PolicyType.get_policy_type(policy_type="REAUTH_POLICY")
    pulumi.export("zpaPolicyTypeReauthPolicy", reauth_policy.id)
    ```

    ```python
    import pulumi
    import pulumi_zpa as zpa

    client_forwarding_policy = zpa.PolicyType.get_policy_type(policy_type="CLIENT_FORWARDING_POLICY")
    pulumi.export("zpaPolicyTypeClientForwardingPolicy", client_forwarding_policy.id)
    ```

    ```python
    import pulumi
    import pulumi_zpa as zpa

    inspection_policy = zpa.PolicyType.get_policy_type(policy_type="INSPECTION_POLICY")
    pulumi.export("zpaPolicyTypeInspectionPolicy", inspection_policy.id)
    ```


    :param str policy_type: The value for differentiating the policy types.
    """
    ...