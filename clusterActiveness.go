package main
import (
  "fmt"
)

type k8sClusterStatus struct {
  Name string
  IsActive bool
}

func determineClusterActiveness(cluster kubeconfigcluster, filterChan chan k8sClusterStatus) {
  kcs := k8sClusterStatus{ Name: cluster.Name }
  matches := findProject(fmt.Sprintf("id = %s", cluster.getProjectId()))
  if len(matches) > 0 {
    kcs.IsActive = true
  }

  filterChan <- kcs
}

func sortClustersByActiveness(clusters []kubeconfigcluster) map[string][]string {
  clusterCount := len(clusters)
  statusChan := make(chan k8sClusterStatus, clusterCount)
  clusterMap := make(map[string][]string)

  clusterMap["active"] = []string{}
  clusterMap["inactive"] = []string{}
  for _, cluster := range clusters {
    // can your system handle the amount of open files this function could generate?
    go determineClusterActiveness(cluster, statusChan)
  }

  for i:=0; i<clusterCount; i++ {
    currentClusterStatus := <- statusChan
    if currentClusterStatus.IsActive {
      clusterMap["active"] = append(clusterMap["active"], currentClusterStatus.Name)
    } else {
      clusterMap["inactive"] = append(clusterMap["inactive"], currentClusterStatus.Name)
    }
  }

  return clusterMap
}
