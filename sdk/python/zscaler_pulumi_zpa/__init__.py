# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .get_zpa_app_connector_group import *
from .get_zpa_application_segment import *
from .get_zpa_application_segment_browser_access import *
from .get_zpa_application_segment_inspection import *
from .get_zpa_application_segment_pra import *
from .get_zpa_application_server import *
from .get_zpa_ba_certificate import *
from .get_zpa_cloud_connector_group import *
from .get_zpa_enrollment_cert import *
from .get_zpa_id_p_controller import *
from .get_zpa_inspection_all_predefined_controls import *
from .get_zpa_inspection_custom_controls import *
from .get_zpa_inspection_predefined_controls import *
from .get_zpa_inspection_profile import *
from .get_zpa_machine_group import *
from .get_zpa_policy_type import *
from .get_zpa_posture_profile import *
from .get_zpa_provisioning_key import *
from .get_zpa_segment_group import *
from .get_zpa_server_group import *
from .get_zpa_service_edge_group import *
from .get_zpa_trusted_network import *
from .get_zpalss_client_types import *
from .get_zpalss_config_controller import *
from .get_zpalss_log_type_formats import *
from .get_zpalss_status_codes import *
from .get_zpasaml_attribute import *
from .get_zpascim_attribute_header import *
from .get_zpascim_groups import *
from .provider import *
from .zpa_app_connector_group import *
from .zpa_application_segment import *
from .zpa_application_segment_browser_access import *
from .zpa_application_segment_inspection import *
from .zpa_application_segment_pra import *
from .zpa_application_server import *
from .zpa_browser_access import *
from .zpa_inspection_custom_controls import *
from .zpa_inspection_profile import *
from .zpa_policy_access_forwarding_rule import *
from .zpa_policy_access_inspection_rule import *
from .zpa_policy_access_rule import *
from .zpa_policy_access_time_out_rule import *
from .zpa_provisioning_key import *
from .zpa_segment_group import *
from .zpa_server_group import *
from .zpa_service_edge_group import *
from .zpalss_config_controller import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import zscaler_pulumi_zpa.config as __config
    config = __config
else:
    config = _utilities.lazy_import('zscaler_pulumi_zpa.config')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "zpa",
  "mod": "index/zPAAppConnectorGroup",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/zPAAppConnectorGroup:ZPAAppConnectorGroup": "ZPAAppConnectorGroup"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/zPAApplicationSegment",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/zPAApplicationSegment:ZPAApplicationSegment": "ZPAApplicationSegment"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/zPAApplicationSegmentBrowserAccess",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/zPAApplicationSegmentBrowserAccess:ZPAApplicationSegmentBrowserAccess": "ZPAApplicationSegmentBrowserAccess"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/zPAApplicationSegmentInspection",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/zPAApplicationSegmentInspection:ZPAApplicationSegmentInspection": "ZPAApplicationSegmentInspection"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/zPAApplicationSegmentPRA",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/zPAApplicationSegmentPRA:ZPAApplicationSegmentPRA": "ZPAApplicationSegmentPRA"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/zPAApplicationServer",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/zPAApplicationServer:ZPAApplicationServer": "ZPAApplicationServer"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/zPABrowserAccess",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/zPABrowserAccess:ZPABrowserAccess": "ZPABrowserAccess"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/zPAInspectionCustomControls",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/zPAInspectionCustomControls:ZPAInspectionCustomControls": "ZPAInspectionCustomControls"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/zPAInspectionProfile",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/zPAInspectionProfile:ZPAInspectionProfile": "ZPAInspectionProfile"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/zPALSSConfigController",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/zPALSSConfigController:ZPALSSConfigController": "ZPALSSConfigController"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/zPAPolicyAccessForwardingRule",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/zPAPolicyAccessForwardingRule:ZPAPolicyAccessForwardingRule": "ZPAPolicyAccessForwardingRule"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/zPAPolicyAccessInspectionRule",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/zPAPolicyAccessInspectionRule:ZPAPolicyAccessInspectionRule": "ZPAPolicyAccessInspectionRule"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/zPAPolicyAccessRule",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/zPAPolicyAccessRule:ZPAPolicyAccessRule": "ZPAPolicyAccessRule"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/zPAPolicyAccessTimeOutRule",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/zPAPolicyAccessTimeOutRule:ZPAPolicyAccessTimeOutRule": "ZPAPolicyAccessTimeOutRule"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/zPAProvisioningKey",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/zPAProvisioningKey:ZPAProvisioningKey": "ZPAProvisioningKey"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/zPASegmentGroup",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/zPASegmentGroup:ZPASegmentGroup": "ZPASegmentGroup"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/zPAServerGroup",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/zPAServerGroup:ZPAServerGroup": "ZPAServerGroup"
  }
 },
 {
  "pkg": "zpa",
  "mod": "index/zPAServiceEdgeGroup",
  "fqn": "zscaler_pulumi_zpa",
  "classes": {
   "zpa:index/zPAServiceEdgeGroup:ZPAServiceEdgeGroup": "ZPAServiceEdgeGroup"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "zpa",
  "token": "pulumi:providers:zpa",
  "fqn": "zscaler_pulumi_zpa",
  "class": "Provider"
 }
]
"""
)
