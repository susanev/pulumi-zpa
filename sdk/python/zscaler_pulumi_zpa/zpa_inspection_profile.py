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

__all__ = ['ZPAInspectionProfileArgs', 'ZPAInspectionProfile']

@pulumi.input_type
class ZPAInspectionProfileArgs:
    def __init__(__self__, *,
                 associate_all_controls: Optional[pulumi.Input[bool]] = None,
                 common_global_override_actions_config: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 controls_infos: Optional[pulumi.Input[Sequence[pulumi.Input['ZPAInspectionProfileControlsInfoArgs']]]] = None,
                 custom_controls: Optional[pulumi.Input[Sequence[pulumi.Input['ZPAInspectionProfileCustomControlArgs']]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 global_control_actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 incarnation_number: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 paranoia_level: Optional[pulumi.Input[str]] = None,
                 predefined_controls: Optional[pulumi.Input[Sequence[pulumi.Input['ZPAInspectionProfilePredefinedControlArgs']]]] = None,
                 predefined_controls_version: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ZPAInspectionProfile resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] common_global_override_actions_config: (Optional)
        :param pulumi.Input[Sequence[pulumi.Input['ZPAInspectionProfileControlsInfoArgs']]] controls_infos: (Optional) Types for custom controls
        :param pulumi.Input[Sequence[pulumi.Input['ZPAInspectionProfileCustomControlArgs']]] custom_controls: (Optional) Types for custom controls
        :param pulumi.Input[str] description: Description of the inspection profile.
        :param pulumi.Input[str] name: The name of the inspection profile.
        :param pulumi.Input[str] paranoia_level: OWASP Predefined Paranoia Level. Range: [1-4], inclusive
        :param pulumi.Input[Sequence[pulumi.Input['ZPAInspectionProfilePredefinedControlArgs']]] predefined_controls: The predefined controls. The default predefined control `Preprocessors` are mandatory and injected in the request by default. Individual `predefined_controls` can be set by using the data source `data_source_zpa_predefined_controls` or by group using the data source `get_zpa_inspection_all_predefined_controls`.
        """
        if associate_all_controls is not None:
            pulumi.set(__self__, "associate_all_controls", associate_all_controls)
        if common_global_override_actions_config is not None:
            pulumi.set(__self__, "common_global_override_actions_config", common_global_override_actions_config)
        if controls_infos is not None:
            pulumi.set(__self__, "controls_infos", controls_infos)
        if custom_controls is not None:
            pulumi.set(__self__, "custom_controls", custom_controls)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if global_control_actions is not None:
            pulumi.set(__self__, "global_control_actions", global_control_actions)
        if incarnation_number is not None:
            pulumi.set(__self__, "incarnation_number", incarnation_number)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if paranoia_level is not None:
            pulumi.set(__self__, "paranoia_level", paranoia_level)
        if predefined_controls is not None:
            pulumi.set(__self__, "predefined_controls", predefined_controls)
        if predefined_controls_version is not None:
            pulumi.set(__self__, "predefined_controls_version", predefined_controls_version)

    @property
    @pulumi.getter(name="associateAllControls")
    def associate_all_controls(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "associate_all_controls")

    @associate_all_controls.setter
    def associate_all_controls(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "associate_all_controls", value)

    @property
    @pulumi.getter(name="commonGlobalOverrideActionsConfig")
    def common_global_override_actions_config(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        (Optional)
        """
        return pulumi.get(self, "common_global_override_actions_config")

    @common_global_override_actions_config.setter
    def common_global_override_actions_config(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "common_global_override_actions_config", value)

    @property
    @pulumi.getter(name="controlsInfos")
    def controls_infos(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ZPAInspectionProfileControlsInfoArgs']]]]:
        """
        (Optional) Types for custom controls
        """
        return pulumi.get(self, "controls_infos")

    @controls_infos.setter
    def controls_infos(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ZPAInspectionProfileControlsInfoArgs']]]]):
        pulumi.set(self, "controls_infos", value)

    @property
    @pulumi.getter(name="customControls")
    def custom_controls(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ZPAInspectionProfileCustomControlArgs']]]]:
        """
        (Optional) Types for custom controls
        """
        return pulumi.get(self, "custom_controls")

    @custom_controls.setter
    def custom_controls(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ZPAInspectionProfileCustomControlArgs']]]]):
        pulumi.set(self, "custom_controls", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the inspection profile.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="globalControlActions")
    def global_control_actions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "global_control_actions")

    @global_control_actions.setter
    def global_control_actions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "global_control_actions", value)

    @property
    @pulumi.getter(name="incarnationNumber")
    def incarnation_number(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "incarnation_number")

    @incarnation_number.setter
    def incarnation_number(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "incarnation_number", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the inspection profile.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="paranoiaLevel")
    def paranoia_level(self) -> Optional[pulumi.Input[str]]:
        """
        OWASP Predefined Paranoia Level. Range: [1-4], inclusive
        """
        return pulumi.get(self, "paranoia_level")

    @paranoia_level.setter
    def paranoia_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "paranoia_level", value)

    @property
    @pulumi.getter(name="predefinedControls")
    def predefined_controls(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ZPAInspectionProfilePredefinedControlArgs']]]]:
        """
        The predefined controls. The default predefined control `Preprocessors` are mandatory and injected in the request by default. Individual `predefined_controls` can be set by using the data source `data_source_zpa_predefined_controls` or by group using the data source `get_zpa_inspection_all_predefined_controls`.
        """
        return pulumi.get(self, "predefined_controls")

    @predefined_controls.setter
    def predefined_controls(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ZPAInspectionProfilePredefinedControlArgs']]]]):
        pulumi.set(self, "predefined_controls", value)

    @property
    @pulumi.getter(name="predefinedControlsVersion")
    def predefined_controls_version(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "predefined_controls_version")

    @predefined_controls_version.setter
    def predefined_controls_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "predefined_controls_version", value)


@pulumi.input_type
class _ZPAInspectionProfileState:
    def __init__(__self__, *,
                 associate_all_controls: Optional[pulumi.Input[bool]] = None,
                 common_global_override_actions_config: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 controls_infos: Optional[pulumi.Input[Sequence[pulumi.Input['ZPAInspectionProfileControlsInfoArgs']]]] = None,
                 custom_controls: Optional[pulumi.Input[Sequence[pulumi.Input['ZPAInspectionProfileCustomControlArgs']]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 global_control_actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 incarnation_number: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 paranoia_level: Optional[pulumi.Input[str]] = None,
                 predefined_controls: Optional[pulumi.Input[Sequence[pulumi.Input['ZPAInspectionProfilePredefinedControlArgs']]]] = None,
                 predefined_controls_version: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ZPAInspectionProfile resources.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] common_global_override_actions_config: (Optional)
        :param pulumi.Input[Sequence[pulumi.Input['ZPAInspectionProfileControlsInfoArgs']]] controls_infos: (Optional) Types for custom controls
        :param pulumi.Input[Sequence[pulumi.Input['ZPAInspectionProfileCustomControlArgs']]] custom_controls: (Optional) Types for custom controls
        :param pulumi.Input[str] description: Description of the inspection profile.
        :param pulumi.Input[str] name: The name of the inspection profile.
        :param pulumi.Input[str] paranoia_level: OWASP Predefined Paranoia Level. Range: [1-4], inclusive
        :param pulumi.Input[Sequence[pulumi.Input['ZPAInspectionProfilePredefinedControlArgs']]] predefined_controls: The predefined controls. The default predefined control `Preprocessors` are mandatory and injected in the request by default. Individual `predefined_controls` can be set by using the data source `data_source_zpa_predefined_controls` or by group using the data source `get_zpa_inspection_all_predefined_controls`.
        """
        if associate_all_controls is not None:
            pulumi.set(__self__, "associate_all_controls", associate_all_controls)
        if common_global_override_actions_config is not None:
            pulumi.set(__self__, "common_global_override_actions_config", common_global_override_actions_config)
        if controls_infos is not None:
            pulumi.set(__self__, "controls_infos", controls_infos)
        if custom_controls is not None:
            pulumi.set(__self__, "custom_controls", custom_controls)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if global_control_actions is not None:
            pulumi.set(__self__, "global_control_actions", global_control_actions)
        if incarnation_number is not None:
            pulumi.set(__self__, "incarnation_number", incarnation_number)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if paranoia_level is not None:
            pulumi.set(__self__, "paranoia_level", paranoia_level)
        if predefined_controls is not None:
            pulumi.set(__self__, "predefined_controls", predefined_controls)
        if predefined_controls_version is not None:
            pulumi.set(__self__, "predefined_controls_version", predefined_controls_version)

    @property
    @pulumi.getter(name="associateAllControls")
    def associate_all_controls(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "associate_all_controls")

    @associate_all_controls.setter
    def associate_all_controls(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "associate_all_controls", value)

    @property
    @pulumi.getter(name="commonGlobalOverrideActionsConfig")
    def common_global_override_actions_config(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        (Optional)
        """
        return pulumi.get(self, "common_global_override_actions_config")

    @common_global_override_actions_config.setter
    def common_global_override_actions_config(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "common_global_override_actions_config", value)

    @property
    @pulumi.getter(name="controlsInfos")
    def controls_infos(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ZPAInspectionProfileControlsInfoArgs']]]]:
        """
        (Optional) Types for custom controls
        """
        return pulumi.get(self, "controls_infos")

    @controls_infos.setter
    def controls_infos(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ZPAInspectionProfileControlsInfoArgs']]]]):
        pulumi.set(self, "controls_infos", value)

    @property
    @pulumi.getter(name="customControls")
    def custom_controls(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ZPAInspectionProfileCustomControlArgs']]]]:
        """
        (Optional) Types for custom controls
        """
        return pulumi.get(self, "custom_controls")

    @custom_controls.setter
    def custom_controls(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ZPAInspectionProfileCustomControlArgs']]]]):
        pulumi.set(self, "custom_controls", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the inspection profile.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="globalControlActions")
    def global_control_actions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "global_control_actions")

    @global_control_actions.setter
    def global_control_actions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "global_control_actions", value)

    @property
    @pulumi.getter(name="incarnationNumber")
    def incarnation_number(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "incarnation_number")

    @incarnation_number.setter
    def incarnation_number(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "incarnation_number", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the inspection profile.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="paranoiaLevel")
    def paranoia_level(self) -> Optional[pulumi.Input[str]]:
        """
        OWASP Predefined Paranoia Level. Range: [1-4], inclusive
        """
        return pulumi.get(self, "paranoia_level")

    @paranoia_level.setter
    def paranoia_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "paranoia_level", value)

    @property
    @pulumi.getter(name="predefinedControls")
    def predefined_controls(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ZPAInspectionProfilePredefinedControlArgs']]]]:
        """
        The predefined controls. The default predefined control `Preprocessors` are mandatory and injected in the request by default. Individual `predefined_controls` can be set by using the data source `data_source_zpa_predefined_controls` or by group using the data source `get_zpa_inspection_all_predefined_controls`.
        """
        return pulumi.get(self, "predefined_controls")

    @predefined_controls.setter
    def predefined_controls(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ZPAInspectionProfilePredefinedControlArgs']]]]):
        pulumi.set(self, "predefined_controls", value)

    @property
    @pulumi.getter(name="predefinedControlsVersion")
    def predefined_controls_version(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "predefined_controls_version")

    @predefined_controls_version.setter
    def predefined_controls_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "predefined_controls_version", value)


class ZPAInspectionProfile(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 associate_all_controls: Optional[pulumi.Input[bool]] = None,
                 common_global_override_actions_config: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 controls_infos: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ZPAInspectionProfileControlsInfoArgs']]]]] = None,
                 custom_controls: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ZPAInspectionProfileCustomControlArgs']]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 global_control_actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 incarnation_number: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 paranoia_level: Optional[pulumi.Input[str]] = None,
                 predefined_controls: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ZPAInspectionProfilePredefinedControlArgs']]]]] = None,
                 predefined_controls_version: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The  **zpa_inspection_profile** resource creates an inspection profile in the Zscaler Private Access cloud. This resource can then be referenced in an inspection custom control resource.

        ## Import

        Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language. [Visit](https://github.com/zscaler/zscaler-terraformer)

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] common_global_override_actions_config: (Optional)
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ZPAInspectionProfileControlsInfoArgs']]]] controls_infos: (Optional) Types for custom controls
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ZPAInspectionProfileCustomControlArgs']]]] custom_controls: (Optional) Types for custom controls
        :param pulumi.Input[str] description: Description of the inspection profile.
        :param pulumi.Input[str] name: The name of the inspection profile.
        :param pulumi.Input[str] paranoia_level: OWASP Predefined Paranoia Level. Range: [1-4], inclusive
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ZPAInspectionProfilePredefinedControlArgs']]]] predefined_controls: The predefined controls. The default predefined control `Preprocessors` are mandatory and injected in the request by default. Individual `predefined_controls` can be set by using the data source `data_source_zpa_predefined_controls` or by group using the data source `get_zpa_inspection_all_predefined_controls`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ZPAInspectionProfileArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The  **zpa_inspection_profile** resource creates an inspection profile in the Zscaler Private Access cloud. This resource can then be referenced in an inspection custom control resource.

        ## Import

        Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language. [Visit](https://github.com/zscaler/zscaler-terraformer)

        :param str resource_name: The name of the resource.
        :param ZPAInspectionProfileArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ZPAInspectionProfileArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 associate_all_controls: Optional[pulumi.Input[bool]] = None,
                 common_global_override_actions_config: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 controls_infos: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ZPAInspectionProfileControlsInfoArgs']]]]] = None,
                 custom_controls: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ZPAInspectionProfileCustomControlArgs']]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 global_control_actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 incarnation_number: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 paranoia_level: Optional[pulumi.Input[str]] = None,
                 predefined_controls: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ZPAInspectionProfilePredefinedControlArgs']]]]] = None,
                 predefined_controls_version: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ZPAInspectionProfileArgs.__new__(ZPAInspectionProfileArgs)

            __props__.__dict__["associate_all_controls"] = associate_all_controls
            __props__.__dict__["common_global_override_actions_config"] = common_global_override_actions_config
            __props__.__dict__["controls_infos"] = controls_infos
            __props__.__dict__["custom_controls"] = custom_controls
            __props__.__dict__["description"] = description
            __props__.__dict__["global_control_actions"] = global_control_actions
            __props__.__dict__["incarnation_number"] = incarnation_number
            __props__.__dict__["name"] = name
            __props__.__dict__["paranoia_level"] = paranoia_level
            __props__.__dict__["predefined_controls"] = predefined_controls
            __props__.__dict__["predefined_controls_version"] = predefined_controls_version
        super(ZPAInspectionProfile, __self__).__init__(
            'zpa:index/zPAInspectionProfile:ZPAInspectionProfile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            associate_all_controls: Optional[pulumi.Input[bool]] = None,
            common_global_override_actions_config: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            controls_infos: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ZPAInspectionProfileControlsInfoArgs']]]]] = None,
            custom_controls: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ZPAInspectionProfileCustomControlArgs']]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            global_control_actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            incarnation_number: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            paranoia_level: Optional[pulumi.Input[str]] = None,
            predefined_controls: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ZPAInspectionProfilePredefinedControlArgs']]]]] = None,
            predefined_controls_version: Optional[pulumi.Input[str]] = None) -> 'ZPAInspectionProfile':
        """
        Get an existing ZPAInspectionProfile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] common_global_override_actions_config: (Optional)
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ZPAInspectionProfileControlsInfoArgs']]]] controls_infos: (Optional) Types for custom controls
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ZPAInspectionProfileCustomControlArgs']]]] custom_controls: (Optional) Types for custom controls
        :param pulumi.Input[str] description: Description of the inspection profile.
        :param pulumi.Input[str] name: The name of the inspection profile.
        :param pulumi.Input[str] paranoia_level: OWASP Predefined Paranoia Level. Range: [1-4], inclusive
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ZPAInspectionProfilePredefinedControlArgs']]]] predefined_controls: The predefined controls. The default predefined control `Preprocessors` are mandatory and injected in the request by default. Individual `predefined_controls` can be set by using the data source `data_source_zpa_predefined_controls` or by group using the data source `get_zpa_inspection_all_predefined_controls`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ZPAInspectionProfileState.__new__(_ZPAInspectionProfileState)

        __props__.__dict__["associate_all_controls"] = associate_all_controls
        __props__.__dict__["common_global_override_actions_config"] = common_global_override_actions_config
        __props__.__dict__["controls_infos"] = controls_infos
        __props__.__dict__["custom_controls"] = custom_controls
        __props__.__dict__["description"] = description
        __props__.__dict__["global_control_actions"] = global_control_actions
        __props__.__dict__["incarnation_number"] = incarnation_number
        __props__.__dict__["name"] = name
        __props__.__dict__["paranoia_level"] = paranoia_level
        __props__.__dict__["predefined_controls"] = predefined_controls
        __props__.__dict__["predefined_controls_version"] = predefined_controls_version
        return ZPAInspectionProfile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="associateAllControls")
    def associate_all_controls(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "associate_all_controls")

    @property
    @pulumi.getter(name="commonGlobalOverrideActionsConfig")
    def common_global_override_actions_config(self) -> pulumi.Output[Mapping[str, str]]:
        """
        (Optional)
        """
        return pulumi.get(self, "common_global_override_actions_config")

    @property
    @pulumi.getter(name="controlsInfos")
    def controls_infos(self) -> pulumi.Output[Sequence['outputs.ZPAInspectionProfileControlsInfo']]:
        """
        (Optional) Types for custom controls
        """
        return pulumi.get(self, "controls_infos")

    @property
    @pulumi.getter(name="customControls")
    def custom_controls(self) -> pulumi.Output[Sequence['outputs.ZPAInspectionProfileCustomControl']]:
        """
        (Optional) Types for custom controls
        """
        return pulumi.get(self, "custom_controls")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Description of the inspection profile.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="globalControlActions")
    def global_control_actions(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "global_control_actions")

    @property
    @pulumi.getter(name="incarnationNumber")
    def incarnation_number(self) -> pulumi.Output[str]:
        return pulumi.get(self, "incarnation_number")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the inspection profile.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="paranoiaLevel")
    def paranoia_level(self) -> pulumi.Output[str]:
        """
        OWASP Predefined Paranoia Level. Range: [1-4], inclusive
        """
        return pulumi.get(self, "paranoia_level")

    @property
    @pulumi.getter(name="predefinedControls")
    def predefined_controls(self) -> pulumi.Output[Sequence['outputs.ZPAInspectionProfilePredefinedControl']]:
        """
        The predefined controls. The default predefined control `Preprocessors` are mandatory and injected in the request by default. Individual `predefined_controls` can be set by using the data source `data_source_zpa_predefined_controls` or by group using the data source `get_zpa_inspection_all_predefined_controls`.
        """
        return pulumi.get(self, "predefined_controls")

    @property
    @pulumi.getter(name="predefinedControlsVersion")
    def predefined_controls_version(self) -> pulumi.Output[str]:
        return pulumi.get(self, "predefined_controls_version")

