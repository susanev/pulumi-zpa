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
    'GetZPAInspectionPredefinedControlsResult',
    'AwaitableGetZPAInspectionPredefinedControlsResult',
    'get_zpa_inspection_predefined_controls',
    'get_zpa_inspection_predefined_controls_output',
]

@pulumi.output_type
class GetZPAInspectionPredefinedControlsResult:
    """
    A collection of values returned by getZPAInspectionPredefinedControls.
    """
    def __init__(__self__, action=None, action_value=None, associated_inspection_profile_names=None, attachment=None, control_group=None, control_number=None, creation_time=None, default_action=None, default_action_value=None, description=None, id=None, modified_time=None, modifiedby=None, name=None, paranoia_level=None, severity=None, version=None):
        if action and not isinstance(action, str):
            raise TypeError("Expected argument 'action' to be a str")
        pulumi.set(__self__, "action", action)
        if action_value and not isinstance(action_value, str):
            raise TypeError("Expected argument 'action_value' to be a str")
        pulumi.set(__self__, "action_value", action_value)
        if associated_inspection_profile_names and not isinstance(associated_inspection_profile_names, list):
            raise TypeError("Expected argument 'associated_inspection_profile_names' to be a list")
        pulumi.set(__self__, "associated_inspection_profile_names", associated_inspection_profile_names)
        if attachment and not isinstance(attachment, str):
            raise TypeError("Expected argument 'attachment' to be a str")
        pulumi.set(__self__, "attachment", attachment)
        if control_group and not isinstance(control_group, str):
            raise TypeError("Expected argument 'control_group' to be a str")
        pulumi.set(__self__, "control_group", control_group)
        if control_number and not isinstance(control_number, str):
            raise TypeError("Expected argument 'control_number' to be a str")
        pulumi.set(__self__, "control_number", control_number)
        if creation_time and not isinstance(creation_time, str):
            raise TypeError("Expected argument 'creation_time' to be a str")
        pulumi.set(__self__, "creation_time", creation_time)
        if default_action and not isinstance(default_action, str):
            raise TypeError("Expected argument 'default_action' to be a str")
        pulumi.set(__self__, "default_action", default_action)
        if default_action_value and not isinstance(default_action_value, str):
            raise TypeError("Expected argument 'default_action_value' to be a str")
        pulumi.set(__self__, "default_action_value", default_action_value)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if modified_time and not isinstance(modified_time, str):
            raise TypeError("Expected argument 'modified_time' to be a str")
        pulumi.set(__self__, "modified_time", modified_time)
        if modifiedby and not isinstance(modifiedby, str):
            raise TypeError("Expected argument 'modifiedby' to be a str")
        pulumi.set(__self__, "modifiedby", modifiedby)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if paranoia_level and not isinstance(paranoia_level, str):
            raise TypeError("Expected argument 'paranoia_level' to be a str")
        pulumi.set(__self__, "paranoia_level", paranoia_level)
        if severity and not isinstance(severity, str):
            raise TypeError("Expected argument 'severity' to be a str")
        pulumi.set(__self__, "severity", severity)
        if version and not isinstance(version, str):
            raise TypeError("Expected argument 'version' to be a str")
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def action(self) -> str:
        """
        (Computed)
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="actionValue")
    def action_value(self) -> str:
        """
        (Computed)
        """
        return pulumi.get(self, "action_value")

    @property
    @pulumi.getter(name="associatedInspectionProfileNames")
    def associated_inspection_profile_names(self) -> Sequence['outputs.GetZPAInspectionPredefinedControlsAssociatedInspectionProfileNameResult']:
        """
        (Computed)
        * `id`- (Computed)
        * `name`- (Computed)
        """
        return pulumi.get(self, "associated_inspection_profile_names")

    @property
    @pulumi.getter
    def attachment(self) -> str:
        """
        (Computed)
        """
        return pulumi.get(self, "attachment")

    @property
    @pulumi.getter(name="controlGroup")
    def control_group(self) -> str:
        """
        (Computed)
        """
        return pulumi.get(self, "control_group")

    @property
    @pulumi.getter(name="controlNumber")
    def control_number(self) -> str:
        """
        (Computed)
        """
        return pulumi.get(self, "control_number")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> str:
        """
        (Computed)
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="defaultAction")
    def default_action(self) -> str:
        """
        (Computed)
        """
        return pulumi.get(self, "default_action")

    @property
    @pulumi.getter(name="defaultActionValue")
    def default_action_value(self) -> str:
        """
        (Computed)
        """
        return pulumi.get(self, "default_action_value")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        (Computed)
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        (Computed)
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="modifiedTime")
    def modified_time(self) -> str:
        """
        (Computed)
        """
        return pulumi.get(self, "modified_time")

    @property
    @pulumi.getter
    def modifiedby(self) -> str:
        return pulumi.get(self, "modifiedby")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="paranoiaLevel")
    def paranoia_level(self) -> str:
        """
        (Computed)
        """
        return pulumi.get(self, "paranoia_level")

    @property
    @pulumi.getter
    def severity(self) -> str:
        """
        (Computed)
        """
        return pulumi.get(self, "severity")

    @property
    @pulumi.getter
    def version(self) -> Optional[str]:
        return pulumi.get(self, "version")


class AwaitableGetZPAInspectionPredefinedControlsResult(GetZPAInspectionPredefinedControlsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetZPAInspectionPredefinedControlsResult(
            action=self.action,
            action_value=self.action_value,
            associated_inspection_profile_names=self.associated_inspection_profile_names,
            attachment=self.attachment,
            control_group=self.control_group,
            control_number=self.control_number,
            creation_time=self.creation_time,
            default_action=self.default_action,
            default_action_value=self.default_action_value,
            description=self.description,
            id=self.id,
            modified_time=self.modified_time,
            modifiedby=self.modifiedby,
            name=self.name,
            paranoia_level=self.paranoia_level,
            severity=self.severity,
            version=self.version)


def get_zpa_inspection_predefined_controls(id: Optional[str] = None,
                                           name: Optional[str] = None,
                                           version: Optional[str] = None,
                                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetZPAInspectionPredefinedControlsResult:
    """
    Use the **zpa_inspection_predefined_controls** data source to get information about an OWASP predefined control and prefedined control version. This data source is required when creating an inspection profile.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_zpa as zpa

    example = zpa.get_zpa_inspection_predefined_controls(name="Failed to parse request body",
        version="OWASP_CRS/3.3.0")
    pulumi.export("zpaInspectionPredefinedControls", example)
    ```


    :param str id: (Computed)
    :param str name: The name of the predefined control.
    :param str version: The version of the predefined control, the default is: `OWASP_CRS/3.3.0`
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    __args__['version'] = version
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('zpa:index/getZPAInspectionPredefinedControls:getZPAInspectionPredefinedControls', __args__, opts=opts, typ=GetZPAInspectionPredefinedControlsResult).value

    return AwaitableGetZPAInspectionPredefinedControlsResult(
        action=__ret__.action,
        action_value=__ret__.action_value,
        associated_inspection_profile_names=__ret__.associated_inspection_profile_names,
        attachment=__ret__.attachment,
        control_group=__ret__.control_group,
        control_number=__ret__.control_number,
        creation_time=__ret__.creation_time,
        default_action=__ret__.default_action,
        default_action_value=__ret__.default_action_value,
        description=__ret__.description,
        id=__ret__.id,
        modified_time=__ret__.modified_time,
        modifiedby=__ret__.modifiedby,
        name=__ret__.name,
        paranoia_level=__ret__.paranoia_level,
        severity=__ret__.severity,
        version=__ret__.version)


@_utilities.lift_output_func(get_zpa_inspection_predefined_controls)
def get_zpa_inspection_predefined_controls_output(id: Optional[pulumi.Input[Optional[str]]] = None,
                                                  name: Optional[pulumi.Input[Optional[str]]] = None,
                                                  version: Optional[pulumi.Input[Optional[str]]] = None,
                                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetZPAInspectionPredefinedControlsResult]:
    """
    Use the **zpa_inspection_predefined_controls** data source to get information about an OWASP predefined control and prefedined control version. This data source is required when creating an inspection profile.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_zpa as zpa

    example = zpa.get_zpa_inspection_predefined_controls(name="Failed to parse request body",
        version="OWASP_CRS/3.3.0")
    pulumi.export("zpaInspectionPredefinedControls", example)
    ```


    :param str id: (Computed)
    :param str name: The name of the predefined control.
    :param str version: The version of the predefined control, the default is: `OWASP_CRS/3.3.0`
    """
    ...
