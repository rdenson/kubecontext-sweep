package main
import (
  "fmt"

  "github.com/rdenson/kubecontext-sweep/gcptools"
  "github.com/rdenson/userio"
)

func main() {
  kubeConfig := gcptools.FetchKubeConfig()
  userio.Write(fmt.Sprintf(
    "your local kubectl configuration references...\n  %d clusters\n  %d contexts\n  %d users\n",
    len(kubeConfig.Clusters),
    len(kubeConfig.Contexts),
    len(kubeConfig.Users),
  ))

  /*kcc := &kubeconfigcluster{
    Name: "gke_fr-lwz5gr5nx9p07x8s03ek5guu8m2_us-central1_default",
    Cluster: kubeconfigclusterdata{
      Server: "https://35.188.11.182",
      CAData: "DATA+OMITTED",
    },
  }
  statusChan := make(chan *gcpMetadata, 1)
  determineClusterActiveness(findProject, kcc, statusChan)
  metadata := <- statusChan
  fmt.Printf(
    "%+v\n%+v\n",
    metadata.ClusterConfig,
    metadata.ProjectData,
  )
  */
  sortedClusters := kubeConfig.SortClustersByActiveness()
  userio.WriteInfo("of the clusters known to your local configuration:")
  userio.ListElement(fmt.Sprintf(
    "%d are active",
    len(sortedClusters["active"]),
  ))
  userio.ListElement(fmt.Sprintf(
    "%d are inactive\n",
    len(sortedClusters["inactive"]),
  ))

  for _, elem := range sortedClusters["active"] {
    fmt.Printf(
      "%s >> %s\n",
      elem.ClusterConfig.Name,
      elem.ProjectData.Name,
    )
  }
}
