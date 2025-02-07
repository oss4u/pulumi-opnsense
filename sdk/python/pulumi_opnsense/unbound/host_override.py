# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities

__all__ = ['HostOverrideArgs', 'HostOverride']

@pulumi.input_type
class HostOverrideArgs:
    def __init__(__self__, *,
                 description: pulumi.Input[str],
                 domain: pulumi.Input[str],
                 enabled: pulumi.Input[bool],
                 hostname: pulumi.Input[str],
                 rr: pulumi.Input[str],
                 mx: Optional[pulumi.Input[str]] = None,
                 mx_prio: Optional[pulumi.Input[int]] = None,
                 server: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a HostOverride resource.
        """
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "enabled", enabled)
        pulumi.set(__self__, "hostname", hostname)
        pulumi.set(__self__, "rr", rr)
        if mx is not None:
            pulumi.set(__self__, "mx", mx)
        if mx_prio is not None:
            pulumi.set(__self__, "mx_prio", mx_prio)
        if server is not None:
            pulumi.set(__self__, "server", server)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Input[str]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: pulumi.Input[str]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Input[str]:
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def hostname(self) -> pulumi.Input[str]:
        return pulumi.get(self, "hostname")

    @hostname.setter
    def hostname(self, value: pulumi.Input[str]):
        pulumi.set(self, "hostname", value)

    @property
    @pulumi.getter
    def rr(self) -> pulumi.Input[str]:
        return pulumi.get(self, "rr")

    @rr.setter
    def rr(self, value: pulumi.Input[str]):
        pulumi.set(self, "rr", value)

    @property
    @pulumi.getter
    def mx(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "mx")

    @mx.setter
    def mx(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mx", value)

    @property
    @pulumi.getter
    def mx_prio(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "mx_prio")

    @mx_prio.setter
    def mx_prio(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "mx_prio", value)

    @property
    @pulumi.getter
    def server(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "server")

    @server.setter
    def server(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server", value)


class HostOverride(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 mx: Optional[pulumi.Input[str]] = None,
                 mx_prio: Optional[pulumi.Input[int]] = None,
                 rr: Optional[pulumi.Input[str]] = None,
                 server: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a HostOverride resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: HostOverrideArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a HostOverride resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param HostOverrideArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(HostOverrideArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 mx: Optional[pulumi.Input[str]] = None,
                 mx_prio: Optional[pulumi.Input[int]] = None,
                 rr: Optional[pulumi.Input[str]] = None,
                 server: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = HostOverrideArgs.__new__(HostOverrideArgs)

            if description is None and not opts.urn:
                raise TypeError("Missing required property 'description'")
            __props__.__dict__["description"] = description
            if domain is None and not opts.urn:
                raise TypeError("Missing required property 'domain'")
            __props__.__dict__["domain"] = domain
            if enabled is None and not opts.urn:
                raise TypeError("Missing required property 'enabled'")
            __props__.__dict__["enabled"] = enabled
            if hostname is None and not opts.urn:
                raise TypeError("Missing required property 'hostname'")
            __props__.__dict__["hostname"] = hostname
            __props__.__dict__["mx"] = mx
            __props__.__dict__["mx_prio"] = mx_prio
            if rr is None and not opts.urn:
                raise TypeError("Missing required property 'rr'")
            __props__.__dict__["rr"] = rr
            __props__.__dict__["server"] = server
            __props__.__dict__["result"] = None
        super(HostOverride, __self__).__init__(
            'opnsense:unbound:HostOverride',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'HostOverride':
        """
        Get an existing HostOverride resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = HostOverrideArgs.__new__(HostOverrideArgs)

        __props__.__dict__["description"] = None
        __props__.__dict__["domain"] = None
        __props__.__dict__["enabled"] = None
        __props__.__dict__["hostname"] = None
        __props__.__dict__["mx"] = None
        __props__.__dict__["mx_prio"] = None
        __props__.__dict__["result"] = None
        __props__.__dict__["rr"] = None
        __props__.__dict__["server"] = None
        return HostOverride(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Output[str]:
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def hostname(self) -> pulumi.Output[str]:
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter
    def mx(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "mx")

    @property
    @pulumi.getter
    def mx_prio(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "mx_prio")

    @property
    @pulumi.getter
    def result(self) -> pulumi.Output[str]:
        return pulumi.get(self, "result")

    @property
    @pulumi.getter
    def rr(self) -> pulumi.Output[str]:
        return pulumi.get(self, "rr")

    @property
    @pulumi.getter
    def server(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "server")

