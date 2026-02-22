using System.Collections.Generic;
using Pulumi;
using Opnsense = Pulumi.Opnsense;

return await Deployment.RunAsync(() => 
{
    var myHostAliasOverride = new Opnsense.Unbound.HostAliasOverride("myHostAliasOverride", new()
    {
        Description = "Pulumi test",
        Domain = "example.com",
        Enabled = true,
        Host = "host",
        Hostname = "hostname",
    });

    return new Dictionary<string, object?>
    {
        ["output"] = new Dictionary<string, object?>
        {
            ["value"] = myHostAliasOverride.Result,
        },
    };
});

