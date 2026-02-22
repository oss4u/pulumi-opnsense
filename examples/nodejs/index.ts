import * as pulumi from "@pulumi/pulumi";
import * as opnsense from "@pulumi/opnsense";

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
