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
    'PolicyAccessForwardingRuleCondition',
    'PolicyAccessForwardingRuleConditionOperand',
]

@pulumi.output_type
class PolicyAccessForwardingRuleCondition(dict):
    def __init__(__self__, *,
                 operator: str,
                 id: Optional[str] = None,
                 negated: Optional[bool] = None,
                 operands: Optional[Sequence['outputs.PolicyAccessForwardingRuleConditionOperand']] = None):
        pulumi.set(__self__, "operator", operator)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if negated is not None:
            pulumi.set(__self__, "negated", negated)
        if operands is not None:
            pulumi.set(__self__, "operands", operands)

    @property
    @pulumi.getter
    def operator(self) -> str:
        return pulumi.get(self, "operator")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def negated(self) -> Optional[bool]:
        return pulumi.get(self, "negated")

    @property
    @pulumi.getter
    def operands(self) -> Optional[Sequence['outputs.PolicyAccessForwardingRuleConditionOperand']]:
        return pulumi.get(self, "operands")


@pulumi.output_type
class PolicyAccessForwardingRuleConditionOperand(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "objectType":
            suggest = "object_type"
        elif key == "idpId":
            suggest = "idp_id"
        elif key == "rhsLists":
            suggest = "rhs_lists"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PolicyAccessForwardingRuleConditionOperand. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PolicyAccessForwardingRuleConditionOperand.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PolicyAccessForwardingRuleConditionOperand.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 lhs: str,
                 object_type: str,
                 id: Optional[str] = None,
                 idp_id: Optional[str] = None,
                 name: Optional[str] = None,
                 rhs: Optional[str] = None,
                 rhs_lists: Optional[Sequence[str]] = None):
        pulumi.set(__self__, "lhs", lhs)
        pulumi.set(__self__, "object_type", object_type)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if idp_id is not None:
            pulumi.set(__self__, "idp_id", idp_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if rhs is not None:
            pulumi.set(__self__, "rhs", rhs)
        if rhs_lists is not None:
            pulumi.set(__self__, "rhs_lists", rhs_lists)

    @property
    @pulumi.getter
    def lhs(self) -> str:
        return pulumi.get(self, "lhs")

    @property
    @pulumi.getter(name="objectType")
    def object_type(self) -> str:
        return pulumi.get(self, "object_type")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="idpId")
    def idp_id(self) -> Optional[str]:
        return pulumi.get(self, "idp_id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def rhs(self) -> Optional[str]:
        return pulumi.get(self, "rhs")

    @property
    @pulumi.getter(name="rhsLists")
    def rhs_lists(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "rhs_lists")


