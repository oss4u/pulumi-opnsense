package provider_test

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/blang/semver"
	goopnsense "github.com/oss4u/go-opnsense/opnsense"
	"github.com/oss4u/go-opnsense/opnsense/core/unbound"
	"github.com/oss4u/go-opnsense/opnsense/core/unbound/overrides"
	opnsenseprovider "github.com/oss4u/pulumi-opnsense/provider"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/pulumi/pulumi/sdk/v3/go/property"
	"github.com/stretchr/testify/require"
)

const (
	defaultIntegrationKey    = "R+lLPklKa2QbfPtcpNeWwI9uNaDcd8ZRFJDUDpgH3uKvdyFn9HpOMqmsevTk5RDpk4FUjZFwgu2JHgQ5"
	defaultIntegrationSecret = "WE27qrbOZxTopTFZdsPdEt0rg8Uhqb6tmS44EuAkLsNk3oMY1GKvB4Zjp9S7oqEJeKM03QoH0QxFp76z"
	integrationStackName     = "dev"
)

type integrationConfig struct {
	address       string
	key           string
	secret        string
	pulumiTestDir string
}

func TestProviderE2E_CreateHostOverrideAndAlias(t *testing.T) {
	if os.Getenv("OPNSENSE_E2E") != "1" {
		t.Skip("set OPNSENSE_E2E=1 to run integration test")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	defer cancel()

	cfg := loadIntegrationConfig(t)
	address := ensureOpnSenseRunning(ctx, t, cfg)

	testProvider, err := integration.NewServer(
		context.Background(),
		opnsenseprovider.Name,
		semver.MustParse("1.0.0"),
		integration.WithProvider(opnsenseprovider.Provider()),
	)
	require.NoError(t, err)

	err = testProvider.Configure(p.ConfigureRequest{Args: property.NewMap(map[string]property.Value{
		"address": property.New(address),
		"key":     property.New(cfg.key),
		"secret":  property.New(cfg.secret),
	})})
	require.NoError(t, err)

	api := goopnsense.NewOpnSenseClient(address, cfg.key, cfg.secret)
	api.SetInsecureSkipVerify(true)
	hostsAPI := overrides.GetHostsOverrideApi(api)
	serviceAPI := unbound.New(api).Service

	unique := fmt.Sprintf("%d", time.Now().UnixNano())
	hostname := "pulumi-e2e-host-" + unique
	aliasHostname := "pulumi-e2e-alias-" + unique
	domain := "xyz.de"
	targetIPv4 := "10.201.202.203"
	hostFQDN := hostname + "." + domain

	deleteOverrideByHostDomainIfExists(t, api, hostsAPI, hostname, domain)

	hostInputs := property.NewMap(map[string]property.Value{
		"enabled":     property.New(true),
		"hostname":    property.New(hostname),
		"domain":      property.New(domain),
		"rr":          property.New("A"),
		"description": property.New("pulumi-opnsense provider e2e host"),
		"server":      property.New(targetIPv4),
	})

	hostCreate, err := testProvider.Create(p.CreateRequest{
		Urn:        providerURN("HostOverride", "provider-e2e-host"),
		Properties: hostInputs,
	})
	require.NoError(t, err)
	require.NotEmpty(t, strings.TrimSpace(hostCreate.ID))

	hostID := strings.TrimSpace(hostCreate.ID)
	aliasID := ""

	t.Cleanup(func() {
		if aliasID != "" {
			_ = testProvider.Delete(p.DeleteRequest{
				ID:         aliasID,
				Urn:        providerURN("HostAliasOverride", "provider-e2e-alias"),
				Properties: property.Map{},
				OldInputs:  property.Map{},
			})
		}

		if hostID != "" {
			_ = testProvider.Delete(p.DeleteRequest{
				ID:         hostID,
				Urn:        providerURN("HostOverride", "provider-e2e-host"),
				Properties: property.Map{},
				OldInputs:  property.Map{},
			})
		}

		deleteOverrideByHostDomainIfExists(t, api, hostsAPI, aliasHostname, domain)
		deleteOverrideByHostDomainIfExists(t, api, hostsAPI, hostname, domain)
		_, _ = serviceAPI.Reconfigure()
	})

	_, err = serviceAPI.Reconfigure()
	require.NoError(t, err)

	verifiedHostID := waitForHostOverrideApplied(ctx, t, api, hostsAPI, hostname, domain, targetIPv4, 2*time.Minute)
	if hostID == "" {
		hostID = verifiedHostID
	}

	aliasInputs := property.NewMap(map[string]property.Value{
		"enabled":     property.New(true),
		"host":        property.New(hostID),
		"hostname":    property.New(aliasHostname),
		"domain":      property.New(domain),
		"description": property.New("pulumi-opnsense provider e2e alias"),
	})

	aliasCreate, err := testProvider.Create(p.CreateRequest{
		Urn:        providerURN("HostAliasOverride", "provider-e2e-alias"),
		Properties: aliasInputs,
	})
	require.NoError(t, err)
	require.NotEmpty(t, strings.TrimSpace(aliasCreate.ID))
	aliasID = strings.TrimSpace(aliasCreate.ID)

	_, err = serviceAPI.Reconfigure()
	require.NoError(t, err)

	waitForHostAliasApplied(t, api, aliasID, hostID, hostFQDN, aliasHostname, domain, 2*time.Minute)
}

func providerURN(typ, name string) resource.URN {
	return resource.NewURN(
		"stack",
		"project",
		"",
		tokens.Type("opnsense:unbound:"+typ),
		name,
	)
}

func loadIntegrationConfig(t *testing.T) integrationConfig {
	t.Helper()

	address := strings.TrimSpace(os.Getenv("OPNSENSE_ADDRESS"))
	key := strings.TrimSpace(os.Getenv("OPNSENSE_KEY"))
	if key == "" {
		key = defaultIntegrationKey
	}

	secret := strings.TrimSpace(os.Getenv("OPNSENSE_SECRET"))
	if secret == "" {
		secret = defaultIntegrationSecret
	}

	_, currentFile, _, ok := runtime.Caller(0)
	require.True(t, ok, "cannot determine current file location")

	defaultPulumiDir := filepath.Clean(filepath.Join(filepath.Dir(currentFile), "../tests/test_infra"))
	pulumiDir := strings.TrimSpace(os.Getenv("OPNSENSE_TEST_INFRA_DIR"))
	if pulumiDir == "" {
		pulumiDir = defaultPulumiDir
	}

	return integrationConfig{
		address:       address,
		key:           key,
		secret:        secret,
		pulumiTestDir: pulumiDir,
	}
}

func ensureOpnSenseRunning(ctx context.Context, t *testing.T, cfg integrationConfig) string {
	t.Helper()

	if cfg.address != "" {
		if isOpnSenseReachable(cfg.address, cfg.key, cfg.secret) {
			return cfg.address
		}

		if alternateAddress := alternateSchemeAddress(cfg.address); alternateAddress != "" {
			primaryTransportReachable := isTransportReachable(cfg.address, 4*time.Second)
			alternateTransportReachable := isTransportReachable(alternateAddress, 4*time.Second)

			if !primaryTransportReachable && alternateTransportReachable {
				t.Logf("OPNSENSE_ADDRESS primary scheme transport is not reachable, using %s", alternateAddress)
				return alternateAddress
			}

			if isOpnSenseReachable(alternateAddress, cfg.key, cfg.secret) {
				t.Logf("OPNSENSE_ADDRESS API not reachable with configured scheme, using %s", alternateAddress)
				return alternateAddress
			}

			if isTransportReachable(alternateAddress, 4*time.Second) {
				if strings.EqualFold(strings.TrimSpace(getScheme(cfg.address)), "https") {
					t.Logf("OPNSENSE_ADDRESS over https is not API-reachable, preferring %s", alternateAddress)
					return alternateAddress
				}

				t.Logf("OPNSENSE_ADDRESS not usable on configured scheme, using transport-reachable %s", alternateAddress)
				return alternateAddress
			}

			if strings.EqualFold(strings.TrimSpace(getScheme(cfg.address)), "https") {
				t.Logf("OPNSENSE_ADDRESS over https is not API-reachable, forcing fallback to %s", alternateAddress)
				return alternateAddress
			}
		}

		if isAddressReachable(cfg.address, 443, 4*time.Second) {
			t.Logf("OPNSENSE_ADDRESS is reachable on TCP/443, skipping pulumi up and using existing instance")
			return cfg.address
		}
	}

	if _, err := exec.LookPath("pulumi"); err != nil {
		t.Fatalf("opnsense not reachable and pulumi not available to start infra: %v", err)
	}

	requireDirExists(t, cfg.pulumiTestDir)
	ensurePulumiPythonVenv(ctx, t, cfg.pulumiTestDir)
	ensurePulumiStack(ctx, t, cfg.pulumiTestDir, integrationStackName)
	runPulumi(ctx, t, cfg.pulumiTestDir, "up", "--yes", "--skip-preview", "--non-interactive", "--stack", integrationStackName)

	boundFqdn := runPulumi(ctx, t, cfg.pulumiTestDir, "stack", "output", "boundFqdn", "--stack", integrationStackName)
	serverIPv4 := runPulumi(ctx, t, cfg.pulumiTestDir, "stack", "output", "serverIpv4", "--stack", integrationStackName)

	address := cfg.address
	if address == "" {
		switch {
		case boundFqdn != "":
			address = "https://" + strings.TrimSpace(boundFqdn)
		case serverIPv4 != "":
			address = "https://" + strings.TrimSpace(serverIPv4)
		default:
			t.Fatal("pulumi stack has no boundFqdn/serverIpv4 output and OPNSENSE_ADDRESS is empty")
		}
	}

	deadline := time.Now().Add(5 * time.Minute)
	for time.Now().Before(deadline) {
		if isOpnSenseReachable(address, cfg.key, cfg.secret) {
			return address
		}
		time.Sleep(5 * time.Second)
	}

	t.Fatalf("opnsense did not become reachable at %s after pulumi up", address)
	return ""
}

func isOpnSenseReachable(address, key, secret string) bool {
	if strings.TrimSpace(address) == "" || strings.TrimSpace(key) == "" || strings.TrimSpace(secret) == "" {
		return false
	}

	api := goopnsense.NewOpnSenseClient(address, key, secret)
	api.SetInsecureSkipVerify(true)
	_, statusCode, err := unbound.New(api).Service.Status()
	return err == nil && statusCode == 200
}

func isAddressReachable(address string, defaultPort int, timeout time.Duration) bool {
	if strings.TrimSpace(address) == "" {
		return false
	}

	parsed, err := url.Parse(address)
	if err != nil || parsed.Host == "" {
		return false
	}

	host := parsed.Hostname()
	port := parsed.Port()
	if port == "" {
		port = "443"
		if parsed.Scheme == "http" {
			port = "80"
		}
		if defaultPort > 0 {
			port = fmt.Sprintf("%d", defaultPort)
		}
	}

	conn, dialErr := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	if dialErr != nil {
		return false
	}
	_ = conn.Close()

	return true
}

func alternateSchemeAddress(address string) string {
	parsed, err := url.Parse(address)
	if err != nil {
		return ""
	}

	switch parsed.Scheme {
	case "https":
		parsed.Scheme = "http"
		return parsed.String()
	case "http":
		parsed.Scheme = "https"
		return parsed.String()
	default:
		return ""
	}
}

func getScheme(address string) string {
	parsed, err := url.Parse(address)
	if err != nil {
		return ""
	}

	return parsed.Scheme
}

func isTransportReachable(address string, timeout time.Duration) bool {
	if strings.TrimSpace(address) == "" {
		return false
	}

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Timeout:   timeout,
		Transport: transport,
	}

	resp, err := client.Get(address)
	if err != nil {
		return false
	}
	_ = resp.Body.Close()

	return true
}

func requireDirExists(t *testing.T, dir string) {
	t.Helper()

	info, err := os.Stat(dir)
	require.NoError(t, err)
	require.True(t, info.IsDir(), "%s is not a directory", dir)
}

func ensurePulumiPythonVenv(ctx context.Context, t *testing.T, dir string) {
	t.Helper()

	venvPython := filepath.Join(dir, ".venv", "bin", "python")
	if _, err := os.Stat(venvPython); err == nil {
		return
	}

	createCmd := exec.CommandContext(ctx, "python3", "-m", "venv", ".venv")
	createCmd.Dir = dir
	out, err := createCmd.CombinedOutput()
	require.NoErrorf(t, err, "failed to create virtualenv: %s", string(out))

	pipCmd := exec.CommandContext(ctx, venvPython, "-m", "pip", "install", "-r", "requirements.txt")
	pipCmd.Dir = dir
	out, err = pipCmd.CombinedOutput()
	require.NoErrorf(t, err, "failed to install Pulumi test_infra requirements: %s", string(out))
}

func ensurePulumiStack(ctx context.Context, t *testing.T, dir, stackName string) {
	t.Helper()

	selectCmd := exec.CommandContext(ctx, "pulumi", "stack", "select", stackName, "--non-interactive")
	selectCmd.Dir = dir
	if out, err := selectCmd.CombinedOutput(); err == nil {
		_ = out
		return
	}

	initCmd := exec.CommandContext(ctx, "pulumi", "stack", "init", stackName, "--non-interactive")
	initCmd.Dir = dir
	out, err := initCmd.CombinedOutput()
	require.NoErrorf(t, err, "failed to init Pulumi stack %s: %s", stackName, string(out))
}

func runPulumi(ctx context.Context, t *testing.T, dir string, args ...string) string {
	t.Helper()

	cmd := exec.CommandContext(ctx, "pulumi", args...)
	cmd.Dir = dir
	cmd.Env = append(os.Environ(), "PULUMI_SKIP_UPDATE_CHECK=1")

	out, err := cmd.CombinedOutput()
	require.NoErrorf(t, err, "pulumi %s failed: %s", strings.Join(args, " "), string(out))

	return strings.TrimSpace(string(out))
}

func deleteOverrideByHostDomainIfExists(t *testing.T, api *goopnsense.OpnSenseApi, hosts overrides.OverridesHostsApi, hostname, domain string) {
	t.Helper()

	searchPayload := map[string]any{
		"current":  1,
		"rowCount": 500,
	}

	payloadRaw, err := json.Marshal(searchPayload)
	require.NoError(t, err)

	raw, err := api.ModifyingRequest("unbound", "settings", "search_host_override", string(payloadRaw), []string{})
	require.NoError(t, err)

	var decoded map[string]any
	if err := json.Unmarshal([]byte(raw), &decoded); err != nil {
		return
	}

	rows, ok := decoded["rows"].([]any)
	if !ok {
		return
	}

	for _, row := range rows {
		rowMap, ok := row.(map[string]any)
		if !ok {
			continue
		}

		candidateHost, _ := rowMap["hostname"].(string)
		candidateDomain, _ := rowMap["domain"].(string)
		uuid, _ := rowMap["uuid"].(string)

		if strings.EqualFold(candidateHost, hostname) && strings.EqualFold(candidateDomain, domain) && uuid != "" {
			_ = hosts.DeleteByID(uuid)
		}
	}
}

func findOverrideUUIDByHostDomain(t *testing.T, api *goopnsense.OpnSenseApi, hostname, domain string) (string, error) {
	t.Helper()

	searchPayload := map[string]any{
		"current":  1,
		"rowCount": 500,
	}

	payloadRaw, err := json.Marshal(searchPayload)
	if err != nil {
		return "", err
	}

	raw, err := api.ModifyingRequest("unbound", "settings", "search_host_override", string(payloadRaw), []string{})
	if err != nil {
		return "", err
	}

	var decoded map[string]any
	if err := json.Unmarshal([]byte(raw), &decoded); err != nil {
		return "", err
	}

	rows, ok := decoded["rows"].([]any)
	if !ok {
		return "", nil
	}

	for _, row := range rows {
		rowMap, ok := row.(map[string]any)
		if !ok {
			continue
		}

		candidateHost, _ := rowMap["hostname"].(string)
		candidateDomain, _ := rowMap["domain"].(string)
		uuid, _ := rowMap["uuid"].(string)

		if strings.EqualFold(candidateHost, hostname) && strings.EqualFold(candidateDomain, domain) {
			return strings.TrimSpace(uuid), nil
		}
	}

	return "", nil
}

func waitForHostOverrideApplied(
	ctx context.Context,
	t *testing.T,
	api *goopnsense.OpnSenseApi,
	hosts overrides.OverridesHostsApi,
	hostname,
	domain,
	targetIPv4 string,
	timeout time.Duration,
) string {
	t.Helper()

	deadline := time.Now().Add(timeout)
	var lastErr error

	for time.Now().Before(deadline) {
		if err := ctx.Err(); err != nil {
			lastErr = err
			break
		}

		uuid, err := findOverrideUUIDByHostDomain(t, api, hostname, domain)
		if err != nil {
			lastErr = err
			time.Sleep(3 * time.Second)
			continue
		}

		if uuid == "" {
			lastErr = fmt.Errorf("override uuid for %s.%s not found", hostname, domain)
			time.Sleep(3 * time.Second)
			continue
		}

		readBack, readErr := hosts.Read(uuid)
		if readErr != nil {
			lastErr = readErr
			time.Sleep(3 * time.Second)
			continue
		}

		if readBack == nil {
			lastErr = fmt.Errorf("override uuid %s returned nil", uuid)
			time.Sleep(3 * time.Second)
			continue
		}

		if strings.EqualFold(strings.TrimSpace(readBack.Host.Hostname), hostname) &&
			strings.EqualFold(strings.TrimSpace(readBack.Host.Domain), domain) &&
			strings.TrimSpace(readBack.Host.Server) == targetIPv4 {
			return uuid
		}

		lastErr = fmt.Errorf(
			"override mismatch got hostname=%s domain=%s server=%s expected hostname=%s domain=%s server=%s",
			readBack.Host.Hostname,
			readBack.Host.Domain,
			readBack.Host.Server,
			hostname,
			domain,
			targetIPv4,
		)
		time.Sleep(3 * time.Second)
	}

	t.Fatalf("override verification for %s.%s failed: %v", hostname, domain, lastErr)
	return ""
}

func waitForHostAliasApplied(
	t *testing.T,
	api *goopnsense.OpnSenseApi,
	aliasUUID,
	hostUUID,
	hostFQDN,
	hostname,
	domain string,
	timeout time.Duration,
) {
	t.Helper()

	deadline := time.Now().Add(timeout)
	var lastErr error

	for time.Now().Before(deadline) {
		resultRaw, statusCode, readErr := api.NonModifyingRequest("unbound", "settings", "get_host_alias", []string{aliasUUID})
		if readErr != nil {
			lastErr = readErr
			time.Sleep(3 * time.Second)
			continue
		}

		if statusCode != 200 || strings.TrimSpace(resultRaw) == "[]" || strings.TrimSpace(resultRaw) == "" {
			lastErr = fmt.Errorf("alias uuid %s not available yet: status=%d body=%q", aliasUUID, statusCode, resultRaw)
			time.Sleep(3 * time.Second)
			continue
		}

		var decoded map[string]any
		if err := json.Unmarshal([]byte(resultRaw), &decoded); err != nil {
			lastErr = fmt.Errorf("alias uuid %s decode failed: %w", aliasUUID, err)
			time.Sleep(3 * time.Second)
			continue
		}

		aliasBlock, ok := decoded["alias"].(map[string]any)
		if !ok {
			lastErr = fmt.Errorf("alias uuid %s payload has no alias object: %s", aliasUUID, resultRaw)
			time.Sleep(3 * time.Second)
			continue
		}

		gotHost := normalizeAliasHost(aliasBlock["host"])
		gotHostname := strings.TrimSpace(anyToString(aliasBlock["hostname"]))
		gotDomain := strings.TrimSpace(anyToString(aliasBlock["domain"]))

		if (strings.EqualFold(gotHost, hostUUID) || strings.EqualFold(gotHost, hostFQDN)) &&
			strings.EqualFold(gotHostname, hostname) &&
			strings.EqualFold(gotDomain, domain) {
			return
		}

		lastErr = fmt.Errorf(
			"alias mismatch got host=%s hostname=%s domain=%s expected host=%s or %s hostname=%s domain=%s",
			gotHost,
			gotHostname,
			gotDomain,
			hostUUID,
			hostFQDN,
			hostname,
			domain,
		)
		time.Sleep(3 * time.Second)
	}

	t.Fatalf("alias verification for %s.%s failed: %v", hostname, domain, lastErr)
}

func normalizeAliasHost(value any) string {
	switch v := value.(type) {
	case string:
		return strings.TrimSpace(v)
	case map[string]any:
		if selected, ok := v["selected"].(string); ok {
			return strings.TrimSpace(selected)
		}

		for _, candidate := range v {
			if candidateMap, ok := candidate.(map[string]any); ok {
				if selectedFlag, ok := candidateMap["selected"].(float64); ok && selectedFlag == 1 {
					if candidateValue, ok := candidateMap["value"].(string); ok {
						return strings.TrimSpace(candidateValue)
					}
				}
			}
		}
	}

	return ""
}

func anyToString(value any) string {
	switch v := value.(type) {
	case string:
		return v
	case fmt.Stringer:
		return v.String()
	default:
		if value == nil {
			return ""
		}
		return fmt.Sprintf("%v", value)
	}
}
