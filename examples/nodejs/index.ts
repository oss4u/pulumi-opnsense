import * as pulumi from "@pulumi/pulumi";
import * as opnsense from "@pulumi/opnsense";

const myRandomResource = new opnsense.Random("myRandomResource", {length: 24});
const myRandomComponent = new opnsense.RandomComponent("myRandomComponent", {length: 24});
export const output = {
    value: myRandomResource.result,
};
