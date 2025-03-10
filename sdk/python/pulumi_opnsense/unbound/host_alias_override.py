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

__all__ = ['HostAliasOverrideArgs', 'HostAliasOverride']

@pulumi.input_type
class HostAliasOverrideArgs:
    def __init__(__self__, *,
                 description: pulumi.Input[str],
                 domain: pulumi.Input[str],
                 enabled: pulumi.Input[bool],
                 host: pulumi.Input[str],
                 hostname: pulumi.Input[str]):
        """
        The set of arguments for constructing a HostAliasOverride resource.
        """
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "enabled", enabled)
        pulumi.set(__self__, "host", host)
        pulumi.set(__self__, "hostname", hostname)

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
    def host(self) -> pulumi.Input[str]:
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: pulumi.Input[str]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter
    def hostname(self) -> pulumi.Input[str]:
        return pulumi.get(self, "hostname")

    @hostname.setter
    def hostname(self, value: pulumi.Input[str]):
        pulumi.set(self, "hostname", value)


class HostAliasOverride(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a HostAliasOverride resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: HostAliasOverrideArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a HostAliasOverride resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param HostAliasOverrideArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(HostAliasOverrideArgs, pulumi.ResourceOptions, *args, **kwargs)
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
                 host: Optional[pulumi.Input[str]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = HostAliasOverrideArgs.__new__(HostAliasOverrideArgs)

            if description is None and not opts.urn:
                raise TypeError("Missing required property 'description'")
            __props__.__dict__["description"] = description
            if domain is None and not opts.urn:
                raise TypeError("Missing required property 'domain'")
            __props__.__dict__["domain"] = domain
            if enabled is None and not opts.urn:
                raise TypeError("Missing required property 'enabled'")
            __props__.__dict__["enabled"] = enabled
            if host is None and not opts.urn:
                raise TypeError("Missing required property 'host'")
            __props__.__dict__["host"] = host
            if hostname is None and not opts.urn:
                raise TypeError("Missing required property 'hostname'")
            __props__.__dict__["hostname"] = hostname
            __props__.__dict__["result"] = None
        super(HostAliasOverride, __self__).__init__(
            'opnsense:unbound:HostAliasOverride',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'HostAliasOverride':
        """
        Get an existing HostAliasOverride resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = HostAliasOverrideArgs.__new__(HostAliasOverrideArgs)

        __props__.__dict__["description"] = None
        __props__.__dict__["domain"] = None
        __props__.__dict__["enabled"] = None
        __props__.__dict__["host"] = None
        __props__.__dict__["hostname"] = None
        __props__.__dict__["result"] = None
        return HostAliasOverride(resource_name, opts=opts, __props__=__props__)

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
    def host(self) -> pulumi.Output[str]:
        return pulumi.get(self, "host")

    @property
    @pulumi.getter
    def hostname(self) -> pulumi.Output[str]:
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter
    def result(self) -> pulumi.Output[str]:
        return pulumi.get(self, "result")

