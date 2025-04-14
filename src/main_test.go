package main

import (
	"github.com/moritzj29/cert-manager-ddnss-de-webhook/src/ddnss"
	"os"
	"testing"

	acmetest "github.com/cert-manager/cert-manager/test/acme"
)

var (
	domain = os.Getenv("DDNSS_DOMAIN_URL")
	zone   = domain + "."
)

func TestRunsSuite(t *testing.T) {
	solver := ddnss.NewSolver(nil)
	fixture := acmetest.NewFixture(solver,
		acmetest.SetResolvedZone(zone),
		acmetest.SetDNSName(domain),
		acmetest.SetAllowAmbientCredentials(false),
		acmetest.SetManifestPath("../testdata/ddnss-solver"),
		//		acmetest.SetConfig("{}"), // needed only on initial run (fails)
	)

	fixture.RunBasic(t)
}
