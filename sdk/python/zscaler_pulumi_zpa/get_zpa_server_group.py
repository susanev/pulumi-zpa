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
    'GetZPAServerGroupResult',
    'AwaitableGetZPAServerGroupResult',
    'get_zpa_server_group',
    'get_zpa_server_group_output',
]

@pulumi.output_type
class GetZPAServerGroupResult:
    """
    A collection of values returned by getZPAServerGroup.
    """
    def __init__(__self__, app_connector_groups=None, applications=None, config_space=None, creation_time=None, description=None, dynamic_discovery=None, enabled=None, id=None, ip_anchored=None, modified_time=None, modifiedby=None, name=None, servers=None):
        if app_connector_groups and not isinstance(app_connector_groups, list):
            raise TypeError("Expected argument 'app_connector_groups' to be a list")
        pulumi.set(__self__, "app_connector_groups", app_connector_groups)
        if applications and not isinstance(applications, list):
            raise TypeError("Expected argument 'applications' to be a list")
        pulumi.set(__self__, "applications", applications)
        if config_space and not isinstance(config_space, str):
            raise TypeError("Expected argument 'config_space' to be a str")
        pulumi.set(__self__, "config_space", config_space)
        if creation_time and not isinstance(creation_time, str):
            raise TypeError("Expected argument 'creation_time' to be a str")
        pulumi.set(__self__, "creation_time", creation_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if dynamic_discovery and not isinstance(dynamic_discovery, bool):
            raise TypeError("Expected argument 'dynamic_discovery' to be a bool")
        pulumi.set(__self__, "dynamic_discovery", dynamic_discovery)
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip_anchored and not isinstance(ip_anchored, bool):
            raise TypeError("Expected argument 'ip_anchored' to be a bool")
        pulumi.set(__self__, "ip_anchored", ip_anchored)
        if modified_time and not isinstance(modified_time, str):
            raise TypeError("Expected argument 'modified_time' to be a str")
        pulumi.set(__self__, "modified_time", modified_time)
        if modifiedby and not isinstance(modifiedby, str):
            raise TypeError("Expected argument 'modifiedby' to be a str")
        pulumi.set(__self__, "modifiedby", modifiedby)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if servers and not isinstance(servers, list):
            raise TypeError("Expected argument 'servers' to be a list")
        pulumi.set(__self__, "servers", servers)

    @property
    @pulumi.getter(name="appConnectorGroups")
    def app_connector_groups(self) -> Sequence['outputs.GetZPAServerGroupAppConnectorGroupResult']:
        """
        (string)This field is a json array of app-connector-id only.
        """
        return pulumi.get(self, "app_connector_groups")

    @property
    @pulumi.getter
    def applications(self) -> Sequence['outputs.GetZPAServerGroupApplicationResult']:
        return pulumi.get(self, "applications")

    @property
    @pulumi.getter(name="configSpace")
    def config_space(self) -> str:
        """
        (string)
        """
        return pulumi.get(self, "config_space")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> str:
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        (string) This field is the description of the server group.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="dynamicDiscovery")
    def dynamic_discovery(self) -> bool:
        """
        (bool) This field controls dynamic discovery of the servers.
        """
        return pulumi.get(self, "dynamic_discovery")

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        """
        (bool) This field defines if the server group is enabled or disabled.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ipAnchored")
    def ip_anchored(self) -> bool:
        """
        (bool)
        """
        return pulumi.get(self, "ip_anchored")

    @property
    @pulumi.getter(name="modifiedTime")
    def modified_time(self) -> str:
        return pulumi.get(self, "modified_time")

    @property
    @pulumi.getter
    def modifiedby(self) -> str:
        return pulumi.get(self, "modifiedby")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def servers(self) -> Sequence['outputs.GetZPAServerGroupServerResult']:
        return pulumi.get(self, "servers")


class AwaitableGetZPAServerGroupResult(GetZPAServerGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetZPAServerGroupResult(
            app_connector_groups=self.app_connector_groups,
            applications=self.applications,
            config_space=self.config_space,
            creation_time=self.creation_time,
            description=self.description,
            dynamic_discovery=self.dynamic_discovery,
            enabled=self.enabled,
            id=self.id,
            ip_anchored=self.ip_anchored,
            modified_time=self.modified_time,
            modifiedby=self.modifiedby,
            name=self.name,
            servers=self.servers)


def get_zpa_server_group(id: Optional[str] = None,
                         name: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetZPAServerGroupResult:
    """
    ## Example Usage

    ```python
    import pulumi
    import pulumi_zpa as zpa

    example = zpa.get_zpa_server_group(name="server_group_name")
    ```


    :param str id: The ID of the server group to be exported.
    :param str name: The name of the server group to be exported.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('zpa:index/getZPAServerGroup:getZPAServerGroup', __args__, opts=opts, typ=GetZPAServerGroupResult).value

    return AwaitableGetZPAServerGroupResult(
        app_connector_groups=__ret__.app_connector_groups,
        applications=__ret__.applications,
        config_space=__ret__.config_space,
        creation_time=__ret__.creation_time,
        description=__ret__.description,
        dynamic_discovery=__ret__.dynamic_discovery,
        enabled=__ret__.enabled,
        id=__ret__.id,
        ip_anchored=__ret__.ip_anchored,
        modified_time=__ret__.modified_time,
        modifiedby=__ret__.modifiedby,
        name=__ret__.name,
        servers=__ret__.servers)


@_utilities.lift_output_func(get_zpa_server_group)
def get_zpa_server_group_output(id: Optional[pulumi.Input[Optional[str]]] = None,
                                name: Optional[pulumi.Input[Optional[str]]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetZPAServerGroupResult]:
    """
    ## Example Usage

    ```python
    import pulumi
    import pulumi_zpa as zpa

    example = zpa.get_zpa_server_group(name="server_group_name")
    ```


    :param str id: The ID of the server group to be exported.
    :param str name: The name of the server group to be exported.
    """
    ...
