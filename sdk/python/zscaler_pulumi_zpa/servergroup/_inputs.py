# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'ServerGroupAppConnectorGroupArgs',
    'ServerGroupApplicationArgs',
    'ServerGroupServerArgs',
]

@pulumi.input_type
class ServerGroupAppConnectorGroupArgs:
    def __init__(__self__, *,
                 ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        if ids is not None:
            pulumi.set(__self__, "ids", ids)

    @property
    @pulumi.getter
    def ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "ids")

    @ids.setter
    def ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ids", value)


@pulumi.input_type
class ServerGroupApplicationArgs:
    def __init__(__self__, *,
                 ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        if ids is not None:
            pulumi.set(__self__, "ids", ids)

    @property
    @pulumi.getter
    def ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "ids")

    @ids.setter
    def ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ids", value)


@pulumi.input_type
class ServerGroupServerArgs:
    def __init__(__self__, *,
                 ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        if ids is not None:
            pulumi.set(__self__, "ids", ids)

    @property
    @pulumi.getter
    def ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "ids")

    @ids.setter
    def ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ids", value)

