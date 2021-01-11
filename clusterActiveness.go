package main
import (
  "fmt"
)

type (
  findproject func(string) gcloudprojects
  k8sClusterStatus struct {
    //kubectl cluster config data
    Config kubeconfigcluster
    //google cloud api data
    //Data interface{}
    Active bool
  }
)

func determineClusterActiveness(find findproject, cluster kubeconfigcluster, statusComm chan *k8sClusterStatus) {
  kcs := &k8sClusterStatus{ Config: cluster }
  if len(find(fmt.Sprintf("id = %s", kcs.Config.getProjectId()))) > 0 {
    kcs.Active = true
  }

  statusComm <- kcs
}

func sortClustersByActiveness(clusters []kubeconfigcluster) map[string][]string {
  clusterCount := len(clusters)
  statusChan := make(chan *k8sClusterStatus, clusterCount)
  clusterMap := make(map[string][]string)

  clusterMap["active"] = []string{}
  clusterMap["inactive"] = []string{}
  for _, cluster := range clusters {
    // can your system handle the amount of open files this function could generate?
    go determineClusterActiveness(findProject, cluster, statusChan)
  }

  for i:=0; i<clusterCount; i++ {
    currentClusterStatus := <- statusChan
    if currentClusterStatus.Active {
      clusterMap["active"] = append(clusterMap["active"], currentClusterStatus.Config.Name)
    } else {
      clusterMap["inactive"] = append(clusterMap["inactive"], currentClusterStatus.Config.Name)
    }
  }

  return clusterMap
}
