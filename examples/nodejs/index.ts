import * as pulumi from "@pulumi/pulumi";
import * as opnsense from "@oss4u/opnsense";

const config = new pulumi.Config();
const opnsenseAddress = config.requireObject("opnsense:address");
const opnsenseKey = config.requireObject("opnsense:key");
const opnsenseSecret = config.requireObject("opnsense:secret");
const myHostAliasOverride = new opnsense.unbound.HostAliasOverride("myHostAliasOverride", {
    description: "Pulumi test",
    domain: "example.com",
    enabled: true,
    host: "host",
    hostname: "hostname",
});
export const output = {
    value: myHostAliasOverride.result,
};
