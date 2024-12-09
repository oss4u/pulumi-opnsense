// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export { ProviderArgs } from "./provider";
export type Provider = import("./provider").Provider;
export const Provider: typeof import("./provider").Provider = null as any;
utilities.lazyLoad(exports, ["Provider"], () => require("./provider"));

export { RandomArgs } from "./random";
export type Random = import("./random").Random;
export const Random: typeof import("./random").Random = null as any;
utilities.lazyLoad(exports, ["Random"], () => require("./random"));

export { RandomComponentArgs } from "./randomComponent";
export type RandomComponent = import("./randomComponent").RandomComponent;
export const RandomComponent: typeof import("./randomComponent").RandomComponent = null as any;
utilities.lazyLoad(exports, ["RandomComponent"], () => require("./randomComponent"));


// Export sub-modules:
import * as config from "./config";

export {
    config,
};

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "opnsense:index:Random":
                return new Random(name, <any>undefined, { urn })
            case "opnsense:index:RandomComponent":
                return new RandomComponent(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("opnsense", "index", _module)
pulumi.runtime.registerResourcePackage("opnsense", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:opnsense") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
