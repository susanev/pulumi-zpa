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
    'GetZPAIdPControllerResult',
    'AwaitableGetZPAIdPControllerResult',
    'get_zpa_id_p_controller',
    'get_zpa_id_p_controller_output',
]

@pulumi.output_type
class GetZPAIdPControllerResult:
    """
    A collection of values returned by getZPAIdPController.
    """
    def __init__(__self__, admin_metadatas=None, admin_sp_signing_cert_id=None, auto_provision=None, creation_time=None, description=None, disable_saml_based_policy=None, domain_lists=None, enable_scim_based_policy=None, enabled=None, id=None, idp_entity_id=None, login_name_attribute=None, login_url=None, modified_time=None, modifiedby=None, name=None, reauth_on_user_update=None, redirect_binding=None, scim_enabled=None, scim_service_provider_endpoint=None, scim_shared_secret_exists=None, sign_saml_request=None, sso_types=None, use_custom_sp_metadata=None, user_metadatas=None, user_sp_signing_cert_id=None):
        if admin_metadatas and not isinstance(admin_metadatas, list):
            raise TypeError("Expected argument 'admin_metadatas' to be a list")
        pulumi.set(__self__, "admin_metadatas", admin_metadatas)
        if admin_sp_signing_cert_id and not isinstance(admin_sp_signing_cert_id, str):
            raise TypeError("Expected argument 'admin_sp_signing_cert_id' to be a str")
        pulumi.set(__self__, "admin_sp_signing_cert_id", admin_sp_signing_cert_id)
        if auto_provision and not isinstance(auto_provision, str):
            raise TypeError("Expected argument 'auto_provision' to be a str")
        pulumi.set(__self__, "auto_provision", auto_provision)
        if creation_time and not isinstance(creation_time, str):
            raise TypeError("Expected argument 'creation_time' to be a str")
        pulumi.set(__self__, "creation_time", creation_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if disable_saml_based_policy and not isinstance(disable_saml_based_policy, bool):
            raise TypeError("Expected argument 'disable_saml_based_policy' to be a bool")
        pulumi.set(__self__, "disable_saml_based_policy", disable_saml_based_policy)
        if domain_lists and not isinstance(domain_lists, list):
            raise TypeError("Expected argument 'domain_lists' to be a list")
        pulumi.set(__self__, "domain_lists", domain_lists)
        if enable_scim_based_policy and not isinstance(enable_scim_based_policy, bool):
            raise TypeError("Expected argument 'enable_scim_based_policy' to be a bool")
        pulumi.set(__self__, "enable_scim_based_policy", enable_scim_based_policy)
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if idp_entity_id and not isinstance(idp_entity_id, str):
            raise TypeError("Expected argument 'idp_entity_id' to be a str")
        pulumi.set(__self__, "idp_entity_id", idp_entity_id)
        if login_name_attribute and not isinstance(login_name_attribute, str):
            raise TypeError("Expected argument 'login_name_attribute' to be a str")
        pulumi.set(__self__, "login_name_attribute", login_name_attribute)
        if login_url and not isinstance(login_url, str):
            raise TypeError("Expected argument 'login_url' to be a str")
        pulumi.set(__self__, "login_url", login_url)
        if modified_time and not isinstance(modified_time, str):
            raise TypeError("Expected argument 'modified_time' to be a str")
        pulumi.set(__self__, "modified_time", modified_time)
        if modifiedby and not isinstance(modifiedby, str):
            raise TypeError("Expected argument 'modifiedby' to be a str")
        pulumi.set(__self__, "modifiedby", modifiedby)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if reauth_on_user_update and not isinstance(reauth_on_user_update, bool):
            raise TypeError("Expected argument 'reauth_on_user_update' to be a bool")
        pulumi.set(__self__, "reauth_on_user_update", reauth_on_user_update)
        if redirect_binding and not isinstance(redirect_binding, bool):
            raise TypeError("Expected argument 'redirect_binding' to be a bool")
        pulumi.set(__self__, "redirect_binding", redirect_binding)
        if scim_enabled and not isinstance(scim_enabled, bool):
            raise TypeError("Expected argument 'scim_enabled' to be a bool")
        pulumi.set(__self__, "scim_enabled", scim_enabled)
        if scim_service_provider_endpoint and not isinstance(scim_service_provider_endpoint, str):
            raise TypeError("Expected argument 'scim_service_provider_endpoint' to be a str")
        pulumi.set(__self__, "scim_service_provider_endpoint", scim_service_provider_endpoint)
        if scim_shared_secret_exists and not isinstance(scim_shared_secret_exists, bool):
            raise TypeError("Expected argument 'scim_shared_secret_exists' to be a bool")
        pulumi.set(__self__, "scim_shared_secret_exists", scim_shared_secret_exists)
        if sign_saml_request and not isinstance(sign_saml_request, str):
            raise TypeError("Expected argument 'sign_saml_request' to be a str")
        pulumi.set(__self__, "sign_saml_request", sign_saml_request)
        if sso_types and not isinstance(sso_types, list):
            raise TypeError("Expected argument 'sso_types' to be a list")
        pulumi.set(__self__, "sso_types", sso_types)
        if use_custom_sp_metadata and not isinstance(use_custom_sp_metadata, bool):
            raise TypeError("Expected argument 'use_custom_sp_metadata' to be a bool")
        pulumi.set(__self__, "use_custom_sp_metadata", use_custom_sp_metadata)
        if user_metadatas and not isinstance(user_metadatas, list):
            raise TypeError("Expected argument 'user_metadatas' to be a list")
        pulumi.set(__self__, "user_metadatas", user_metadatas)
        if user_sp_signing_cert_id and not isinstance(user_sp_signing_cert_id, str):
            raise TypeError("Expected argument 'user_sp_signing_cert_id' to be a str")
        pulumi.set(__self__, "user_sp_signing_cert_id", user_sp_signing_cert_id)

    @property
    @pulumi.getter(name="adminMetadatas")
    def admin_metadatas(self) -> Sequence['outputs.GetZPAIdPControllerAdminMetadataResult']:
        return pulumi.get(self, "admin_metadatas")

    @property
    @pulumi.getter(name="adminSpSigningCertId")
    def admin_sp_signing_cert_id(self) -> str:
        return pulumi.get(self, "admin_sp_signing_cert_id")

    @property
    @pulumi.getter(name="autoProvision")
    def auto_provision(self) -> str:
        return pulumi.get(self, "auto_provision")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> str:
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="disableSamlBasedPolicy")
    def disable_saml_based_policy(self) -> bool:
        return pulumi.get(self, "disable_saml_based_policy")

    @property
    @pulumi.getter(name="domainLists")
    def domain_lists(self) -> Sequence[str]:
        return pulumi.get(self, "domain_lists")

    @property
    @pulumi.getter(name="enableScimBasedPolicy")
    def enable_scim_based_policy(self) -> bool:
        return pulumi.get(self, "enable_scim_based_policy")

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="idpEntityId")
    def idp_entity_id(self) -> str:
        return pulumi.get(self, "idp_entity_id")

    @property
    @pulumi.getter(name="loginNameAttribute")
    def login_name_attribute(self) -> str:
        return pulumi.get(self, "login_name_attribute")

    @property
    @pulumi.getter(name="loginUrl")
    def login_url(self) -> str:
        return pulumi.get(self, "login_url")

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
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="reauthOnUserUpdate")
    def reauth_on_user_update(self) -> bool:
        return pulumi.get(self, "reauth_on_user_update")

    @property
    @pulumi.getter(name="redirectBinding")
    def redirect_binding(self) -> bool:
        return pulumi.get(self, "redirect_binding")

    @property
    @pulumi.getter(name="scimEnabled")
    def scim_enabled(self) -> bool:
        return pulumi.get(self, "scim_enabled")

    @property
    @pulumi.getter(name="scimServiceProviderEndpoint")
    def scim_service_provider_endpoint(self) -> str:
        return pulumi.get(self, "scim_service_provider_endpoint")

    @property
    @pulumi.getter(name="scimSharedSecretExists")
    def scim_shared_secret_exists(self) -> bool:
        return pulumi.get(self, "scim_shared_secret_exists")

    @property
    @pulumi.getter(name="signSamlRequest")
    def sign_saml_request(self) -> str:
        return pulumi.get(self, "sign_saml_request")

    @property
    @pulumi.getter(name="ssoTypes")
    def sso_types(self) -> Sequence[str]:
        return pulumi.get(self, "sso_types")

    @property
    @pulumi.getter(name="useCustomSpMetadata")
    def use_custom_sp_metadata(self) -> bool:
        return pulumi.get(self, "use_custom_sp_metadata")

    @property
    @pulumi.getter(name="userMetadatas")
    def user_metadatas(self) -> Sequence['outputs.GetZPAIdPControllerUserMetadataResult']:
        return pulumi.get(self, "user_metadatas")

    @property
    @pulumi.getter(name="userSpSigningCertId")
    def user_sp_signing_cert_id(self) -> str:
        return pulumi.get(self, "user_sp_signing_cert_id")


class AwaitableGetZPAIdPControllerResult(GetZPAIdPControllerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetZPAIdPControllerResult(
            admin_metadatas=self.admin_metadatas,
            admin_sp_signing_cert_id=self.admin_sp_signing_cert_id,
            auto_provision=self.auto_provision,
            creation_time=self.creation_time,
            description=self.description,
            disable_saml_based_policy=self.disable_saml_based_policy,
            domain_lists=self.domain_lists,
            enable_scim_based_policy=self.enable_scim_based_policy,
            enabled=self.enabled,
            id=self.id,
            idp_entity_id=self.idp_entity_id,
            login_name_attribute=self.login_name_attribute,
            login_url=self.login_url,
            modified_time=self.modified_time,
            modifiedby=self.modifiedby,
            name=self.name,
            reauth_on_user_update=self.reauth_on_user_update,
            redirect_binding=self.redirect_binding,
            scim_enabled=self.scim_enabled,
            scim_service_provider_endpoint=self.scim_service_provider_endpoint,
            scim_shared_secret_exists=self.scim_shared_secret_exists,
            sign_saml_request=self.sign_saml_request,
            sso_types=self.sso_types,
            use_custom_sp_metadata=self.use_custom_sp_metadata,
            user_metadatas=self.user_metadatas,
            user_sp_signing_cert_id=self.user_sp_signing_cert_id)


def get_zpa_id_p_controller(id: Optional[str] = None,
                            name: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetZPAIdPControllerResult:
    """
    Use the **zpa_idp_controller** data source to get information about an Identity Provider created in the Zscaler Private Access cloud. This data source is required when creating:

    1. Access policy Rules
    2. Access policy timeout rules
    3. Access policy forwarding rules
    4. Access policy inspection rules
    5. Access policy isolation rules

    ## Example Usage

    ```python
    import pulumi
    import pulumi_zpa as zpa

    example = zpa.get_zpa_id_p_controller(name="idp_name")
    ```

    ```python
    import pulumi
    import pulumi_zpa as zpa

    example = zpa.get_zpa_id_p_controller(id="1234567890")
    ```


    :param str id: The name of the Identity Provider (IdP) to be exported.
    :param str name: The name of the Identity Provider (IdP) to be exported.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('zpa:index/getZPAIdPController:getZPAIdPController', __args__, opts=opts, typ=GetZPAIdPControllerResult).value

    return AwaitableGetZPAIdPControllerResult(
        admin_metadatas=__ret__.admin_metadatas,
        admin_sp_signing_cert_id=__ret__.admin_sp_signing_cert_id,
        auto_provision=__ret__.auto_provision,
        creation_time=__ret__.creation_time,
        description=__ret__.description,
        disable_saml_based_policy=__ret__.disable_saml_based_policy,
        domain_lists=__ret__.domain_lists,
        enable_scim_based_policy=__ret__.enable_scim_based_policy,
        enabled=__ret__.enabled,
        id=__ret__.id,
        idp_entity_id=__ret__.idp_entity_id,
        login_name_attribute=__ret__.login_name_attribute,
        login_url=__ret__.login_url,
        modified_time=__ret__.modified_time,
        modifiedby=__ret__.modifiedby,
        name=__ret__.name,
        reauth_on_user_update=__ret__.reauth_on_user_update,
        redirect_binding=__ret__.redirect_binding,
        scim_enabled=__ret__.scim_enabled,
        scim_service_provider_endpoint=__ret__.scim_service_provider_endpoint,
        scim_shared_secret_exists=__ret__.scim_shared_secret_exists,
        sign_saml_request=__ret__.sign_saml_request,
        sso_types=__ret__.sso_types,
        use_custom_sp_metadata=__ret__.use_custom_sp_metadata,
        user_metadatas=__ret__.user_metadatas,
        user_sp_signing_cert_id=__ret__.user_sp_signing_cert_id)


@_utilities.lift_output_func(get_zpa_id_p_controller)
def get_zpa_id_p_controller_output(id: Optional[pulumi.Input[Optional[str]]] = None,
                                   name: Optional[pulumi.Input[Optional[str]]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetZPAIdPControllerResult]:
    """
    Use the **zpa_idp_controller** data source to get information about an Identity Provider created in the Zscaler Private Access cloud. This data source is required when creating:

    1. Access policy Rules
    2. Access policy timeout rules
    3. Access policy forwarding rules
    4. Access policy inspection rules
    5. Access policy isolation rules

    ## Example Usage

    ```python
    import pulumi
    import pulumi_zpa as zpa

    example = zpa.get_zpa_id_p_controller(name="idp_name")
    ```

    ```python
    import pulumi
    import pulumi_zpa as zpa

    example = zpa.get_zpa_id_p_controller(id="1234567890")
    ```


    :param str id: The name of the Identity Provider (IdP) to be exported.
    :param str name: The name of the Identity Provider (IdP) to be exported.
    """
    ...
