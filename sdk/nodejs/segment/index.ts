// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { ApplicationSegmentArgs, ApplicationSegmentState } from "./applicationSegment";
export type ApplicationSegment = import("./applicationSegment").ApplicationSegment;
export const ApplicationSegment: typeof import("./applicationSegment").ApplicationSegment = null as any;
utilities.lazyLoad(exports, ["ApplicationSegment"], () => require("./applicationSegment"));

export { ApplicationSegmentBrowserAccessArgs, ApplicationSegmentBrowserAccessState } from "./applicationSegmentBrowserAccess";
export type ApplicationSegmentBrowserAccess = import("./applicationSegmentBrowserAccess").ApplicationSegmentBrowserAccess;
export const ApplicationSegmentBrowserAccess: typeof import("./applicationSegmentBrowserAccess").ApplicationSegmentBrowserAccess = null as any;
utilities.lazyLoad(exports, ["ApplicationSegmentBrowserAccess"], () => require("./applicationSegmentBrowserAccess"));

export { ApplicationSegmentInspectionArgs, ApplicationSegmentInspectionState } from "./applicationSegmentInspection";
export type ApplicationSegmentInspection = import("./applicationSegmentInspection").ApplicationSegmentInspection;
export const ApplicationSegmentInspection: typeof import("./applicationSegmentInspection").ApplicationSegmentInspection = null as any;
utilities.lazyLoad(exports, ["ApplicationSegmentInspection"], () => require("./applicationSegmentInspection"));

export { ApplicationSegmentPRAArgs, ApplicationSegmentPRAState } from "./applicationSegmentPRA";
export type ApplicationSegmentPRA = import("./applicationSegmentPRA").ApplicationSegmentPRA;
export const ApplicationSegmentPRA: typeof import("./applicationSegmentPRA").ApplicationSegmentPRA = null as any;
utilities.lazyLoad(exports, ["ApplicationSegmentPRA"], () => require("./applicationSegmentPRA"));

export { BrowserAccessArgs, BrowserAccessState } from "./browserAccess";
export type BrowserAccess = import("./browserAccess").BrowserAccess;
export const BrowserAccess: typeof import("./browserAccess").BrowserAccess = null as any;
utilities.lazyLoad(exports, ["BrowserAccess"], () => require("./browserAccess"));

export { GetApplicationSegmentArgs, GetApplicationSegmentResult, GetApplicationSegmentOutputArgs } from "./getApplicationSegment";
export const getApplicationSegment: typeof import("./getApplicationSegment").getApplicationSegment = null as any;
export const getApplicationSegmentOutput: typeof import("./getApplicationSegment").getApplicationSegmentOutput = null as any;
utilities.lazyLoad(exports, ["getApplicationSegment","getApplicationSegmentOutput"], () => require("./getApplicationSegment"));

export { GetApplicationSegmentBrowserAccessArgs, GetApplicationSegmentBrowserAccessResult, GetApplicationSegmentBrowserAccessOutputArgs } from "./getApplicationSegmentBrowserAccess";
export const getApplicationSegmentBrowserAccess: typeof import("./getApplicationSegmentBrowserAccess").getApplicationSegmentBrowserAccess = null as any;
export const getApplicationSegmentBrowserAccessOutput: typeof import("./getApplicationSegmentBrowserAccess").getApplicationSegmentBrowserAccessOutput = null as any;
utilities.lazyLoad(exports, ["getApplicationSegmentBrowserAccess","getApplicationSegmentBrowserAccessOutput"], () => require("./getApplicationSegmentBrowserAccess"));

export { GetApplicationSegmentInspectionArgs, GetApplicationSegmentInspectionResult, GetApplicationSegmentInspectionOutputArgs } from "./getApplicationSegmentInspection";
export const getApplicationSegmentInspection: typeof import("./getApplicationSegmentInspection").getApplicationSegmentInspection = null as any;
export const getApplicationSegmentInspectionOutput: typeof import("./getApplicationSegmentInspection").getApplicationSegmentInspectionOutput = null as any;
utilities.lazyLoad(exports, ["getApplicationSegmentInspection","getApplicationSegmentInspectionOutput"], () => require("./getApplicationSegmentInspection"));

export { GetApplicationSegmentPRAArgs, GetApplicationSegmentPRAResult, GetApplicationSegmentPRAOutputArgs } from "./getApplicationSegmentPRA";
export const getApplicationSegmentPRA: typeof import("./getApplicationSegmentPRA").getApplicationSegmentPRA = null as any;
export const getApplicationSegmentPRAOutput: typeof import("./getApplicationSegmentPRA").getApplicationSegmentPRAOutput = null as any;
utilities.lazyLoad(exports, ["getApplicationSegmentPRA","getApplicationSegmentPRAOutput"], () => require("./getApplicationSegmentPRA"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "zpa:Segment/applicationSegment:ApplicationSegment":
                return new ApplicationSegment(name, <any>undefined, { urn })
            case "zpa:Segment/applicationSegmentBrowserAccess:ApplicationSegmentBrowserAccess":
                return new ApplicationSegmentBrowserAccess(name, <any>undefined, { urn })
            case "zpa:Segment/applicationSegmentInspection:ApplicationSegmentInspection":
                return new ApplicationSegmentInspection(name, <any>undefined, { urn })
            case "zpa:Segment/applicationSegmentPRA:ApplicationSegmentPRA":
                return new ApplicationSegmentPRA(name, <any>undefined, { urn })
            case "zpa:Segment/browserAccess:BrowserAccess":
                return new BrowserAccess(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("zpa", "Segment/applicationSegment", _module)
pulumi.runtime.registerResourceModule("zpa", "Segment/applicationSegmentBrowserAccess", _module)
pulumi.runtime.registerResourceModule("zpa", "Segment/applicationSegmentInspection", _module)
pulumi.runtime.registerResourceModule("zpa", "Segment/applicationSegmentPRA", _module)
pulumi.runtime.registerResourceModule("zpa", "Segment/browserAccess", _module)
