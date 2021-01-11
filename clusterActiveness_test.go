package main
import (
  "testing"

  "github.com/stretchr/testify/suite"
)

var (
  findprojectReturnsSomething findproject = func(s string) gcloudprojects {
    return gcloudprojects{ gcoudproject{} }
  }
  findprojectReturnsNothing findproject = func(s string) gcloudprojects {
    return gcloudprojects{}
  }
)

type (
  ActvienessTestCase struct {
    expectedValue bool
    channel chan *k8sClusterStatus
    findFunction findproject
  }
  ClusterActivenessSuite struct {
    KubectlConfigResponse kubeconfig
    suite.Suite
  }
)

func TestClusterActiveness(t *testing.T) {
  cas := new(ClusterActivenessSuite)

  cas.KubectlConfigResponse = kubeconfig{
    Clusters: []kubeconfigcluster{
      kubeconfigcluster{
        Name: "test_cluster_a",
        Cluster: kubeconfigclusterdata{},
      },
      kubeconfigcluster{
        Name: "test_cluster_b",
        Cluster: kubeconfigclusterdata{},
      },
    },
  }
  suite.Run(t, cas)
}

func (caSuite *ClusterActivenessSuite) TestDetermineClusterActiveness() {
  testCases := map[string]ActvienessTestCase{
    "cluster is active": {
      true,
      make(chan *k8sClusterStatus, 1),
      findprojectReturnsSomething,
    },
    "cluster is not active": {
      false,
      make(chan *k8sClusterStatus, 1),
      findprojectReturnsNothing,
    },
  }

  for _, testCase := range testCases {
    determineClusterActiveness(
      testCase.findFunction,
      caSuite.KubectlConfigResponse.Clusters[0],
      testCase.channel,
    )
    actualClusterStatus := <- testCase.channel
    caSuite.Equal(testCase.expectedValue, actualClusterStatus.Active)
  }
}
