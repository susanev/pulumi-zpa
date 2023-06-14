// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetServerGroupArgs, GetServerGroupResult, GetServerGroupOutputArgs } from "./getServerGroup";
export const getServerGroup: typeof import("./getServerGroup").getServerGroup = null as any;
export const getServerGroupOutput: typeof import("./getServerGroup").getServerGroupOutput = null as any;
utilities.lazyLoad(exports, ["getServerGroup","getServerGroupOutput"], () => require("./getServerGroup"));

export { ServerGroupArgs, ServerGroupState } from "./serverGroup";
export type ServerGroup = import("./serverGroup").ServerGroup;
export const ServerGroup: typeof import("./serverGroup").ServerGroup = null as any;
utilities.lazyLoad(exports, ["ServerGroup"], () => require("./serverGroup"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "zpa:ServerGroup/serverGroup:ServerGroup":
                return new ServerGroup(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("zpa", "ServerGroup/serverGroup", _module)