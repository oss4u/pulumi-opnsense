package unbound

import (
	goopnsense "github.com/oss4u/go-opnsense/opnsense/core/unbound/overrides"
	"github.com/oss4u/go-opnsense/opnsense/types"
)

func OverridesHostToHostOverrideArgs(host *goopnsense.OverridesHost) *HostOverrideArgs {
	enabled := host.Host.Enabled.Bool()
	rr := host.Host.Rr.String()
	mxprio := host.Host.Mxprio.Int()
	args := HostOverrideArgs{
		Enabled:     &enabled,
		Hostname:    &host.Host.Hostname,
		Domain:      &host.Host.Domain,
		Rr:          &rr,
		Description: &host.Host.Description,
	}
	if host.Host.Rr == "A" || host.Host.Rr == "AAAA" {
		args.Server = &host.Host.Server
	} else if host.Host.Rr == "MX" {
		args.Mx = &host.Host.Mx
		args.MxPrio = &mxprio
	}
	return &args
}

//func setRROnArgs(r *goopnsense.Rr) *string {
//	var result string
//	result = fmt.Sprintf("%s", *r)
//	return &result
//}
//
//func setMxPrioOnArgs(m *goopnsense.MxPrio) *int {
//	var result int
//	result, _ = strconv.Atoi(fmt.Sprintf("%d", *m))
//	return &result
//}
//
//func setEnabledOnArgs(t *types.Bool) *bool {
//	var b bool
//	if *t == true {
//		b = true
//	} else {
//		b = false
//	}
//	return &b
//
//}

func HostOverrideArgsToOverridesHost(args *HostOverrideArgs) *goopnsense.OverridesHost {
	host := goopnsense.OverridesHostDetails{
		Enabled:     types.Bool(*args.Enabled),
		Hostname:    *args.Hostname,
		Domain:      *args.Domain,
		Rr:          goopnsense.Rr(*args.Rr),
		Description: *args.Description,
	}
	if *args.Rr == "A" || *args.Rr == "AAAA" {
		host.Server = *args.Server
	} else if *args.Rr == "MX" {
		host.Mx = *args.Mx
		host.Mxprio = goopnsense.MxPrio(*args.MxPrio)
	}
	return &goopnsense.OverridesHost{Host: host}
}
