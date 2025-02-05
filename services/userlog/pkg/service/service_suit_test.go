package service_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	mRegistry "go-micro.dev/v4/registry"

	"github.com/owncloud/ocis/v2/ocis-pkg/registry"
)

func init() {
	r := registry.GetRegistry(registry.Inmemory())
	service := registry.BuildGRPCService("com.owncloud.api.gateway", "", "", "")
	service.Nodes = []*mRegistry.Node{{
		Address: "any",
	}}

	_ = r.Register(service)
}
func TestSearch(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Userlog service Suite")
}
