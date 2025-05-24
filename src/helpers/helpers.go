package helpers

import (
	"github.com/cert-manager/cert-manager/pkg/issuer/acme/dns/util"
	"io"
	"k8s.io/klog/v2"
	"net/http"
	"strings"
)

const DDNSSSuffix = ".ddnss.de"

func GetDomainName(DNSName string) string {
	domainName := util.UnFqdn(DNSName) // Remove trailing dot (domain.ddnss.de. -> domain.ddnss.de)
	domainName = strings.TrimSuffix(domainName, DDNSSSuffix)

	split := strings.Split(domainName, ".")

	// If it's prefix.domain, return domain
	if len(split) >= 2 {
		return split[len(split)-1] + DDNSSSuffix
	} else {
		return domainName + DDNSSSuffix
	}
}

func GetResponseBody(res *http.Response) (string, error) {
	body, err := io.ReadAll(res.Body)
	if err != nil {
		klog.Errorf("Unable to get body from response")
		return "", err
	}
	result := string(body)
	result = strings.ReplaceAll(result, "\n", "\t")
	return result, err
}
