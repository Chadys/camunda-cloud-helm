package zeebe

import (
	"camunda-cloud-helm/charts/ccsm-helm/test/golden"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func TestGoldenContainerSecurityContext(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "ccsm-helm-test",
		Namespace:      "ccsm-helm-" + strings.ToLower(random.UniqueId()),
		GoldenFileName: "containersecuritycontext",
		Templates:      []string{"charts/zeebe/templates/statefulset.yaml"},
		SetValues: map[string]string{
			"zeebe.containerSecurityContext.privileged":          "true",
			"zeebe.containerSecurityContext.capabilities.add[0]": "NET_ADMIN",
		},
	})
}
