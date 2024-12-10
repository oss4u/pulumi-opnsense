import pulumi
import pulumi_opnsense as opnsense

my_random_resource = opnsense.Random("myRandomResource", length=24)
my_random_component = opnsense.RandomComponent("myRandomComponent", length=24)
pulumi.export("output", {
    "value": my_random_resource.result,
})
