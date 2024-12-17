using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Opnsense = Pulumi.Opnsense;

return await Deployment.RunAsync(() => 
{
    var myRandomResource = new Opnsense.Random("myRandomResource", new()
    {
        Length = 24,
    });

    var myRandomComponent = new Opnsense.RandomComponent("myRandomComponent", new()
    {
        Length = 24,
    });

    return new Dictionary<string, object?>
    {
        ["output"] = 
        {
            { "value", myRandomResource.Result },
        },
    };
});

