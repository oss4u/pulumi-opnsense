using System.Collections.Generic;
using System.Linq;
using Pulumi;
using opnsense = Pulumi.opnsense;

return await Deployment.RunAsync(() => 
{
    var myRandomResource = new opnsense.Random("myRandomResource", new()
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

