// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetServiceEdgeGroupArgs, GetServiceEdgeGroupResult, GetServiceEdgeGroupOutputArgs } from "./getServiceEdgeGroup";
export const getServiceEdgeGroup: typeof import("./getServiceEdgeGroup").getServiceEdgeGroup = null as any;
export const getServiceEdgeGroupOutput: typeof import("./getServiceEdgeGroup").getServiceEdgeGroupOutput = null as any;
utilities.lazyLoad(exports, ["getServiceEdgeGroup","getServiceEdgeGroupOutput"], () => require("./getServiceEdgeGroup"));

export { ServiceEdgeGroupArgs, ServiceEdgeGroupState } from "./serviceEdgeGroup";
export type ServiceEdgeGroup = import("./serviceEdgeGroup").ServiceEdgeGroup;
export const ServiceEdgeGroup: typeof import("./serviceEdgeGroup").ServiceEdgeGroup = null as any;
utilities.lazyLoad(exports, ["ServiceEdgeGroup"], () => require("./serviceEdgeGroup"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "zpa:ServiceEdge/serviceEdgeGroup:ServiceEdgeGroup":
                return new ServiceEdgeGroup(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("zpa", "ServiceEdge/serviceEdgeGroup", _module)
