// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetLSSClientTypesResult } from "./getLSSClientTypes";
export const getLSSClientTypes: typeof import("./getLSSClientTypes").getLSSClientTypes = null as any;
utilities.lazyLoad(exports, ["getLSSClientTypes"], () => require("./getLSSClientTypes"));

export { GetLSSConfigControllerArgs, GetLSSConfigControllerResult, GetLSSConfigControllerOutputArgs } from "./getLSSConfigController";
export const getLSSConfigController: typeof import("./getLSSConfigController").getLSSConfigController = null as any;
export const getLSSConfigControllerOutput: typeof import("./getLSSConfigController").getLSSConfigControllerOutput = null as any;
utilities.lazyLoad(exports, ["getLSSConfigController","getLSSConfigControllerOutput"], () => require("./getLSSConfigController"));

export { GetLSSLogTypeFormatsArgs, GetLSSLogTypeFormatsResult, GetLSSLogTypeFormatsOutputArgs } from "./getLSSLogTypeFormats";
export const getLSSLogTypeFormats: typeof import("./getLSSLogTypeFormats").getLSSLogTypeFormats = null as any;
export const getLSSLogTypeFormatsOutput: typeof import("./getLSSLogTypeFormats").getLSSLogTypeFormatsOutput = null as any;
utilities.lazyLoad(exports, ["getLSSLogTypeFormats","getLSSLogTypeFormatsOutput"], () => require("./getLSSLogTypeFormats"));

export { GetLSSStatusCodesResult } from "./getLSSStatusCodes";
export const getLSSStatusCodes: typeof import("./getLSSStatusCodes").getLSSStatusCodes = null as any;
utilities.lazyLoad(exports, ["getLSSStatusCodes"], () => require("./getLSSStatusCodes"));

export { LSSConfigControllerArgs, LSSConfigControllerState } from "./lssconfigController";
export type LSSConfigController = import("./lssconfigController").LSSConfigController;
export const LSSConfigController: typeof import("./lssconfigController").LSSConfigController = null as any;
utilities.lazyLoad(exports, ["LSSConfigController"], () => require("./lssconfigController"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "zpa:LSSConfig/lSSConfigController:LSSConfigController":
                return new LSSConfigController(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("zpa", "LSSConfig/lSSConfigController", _module)
