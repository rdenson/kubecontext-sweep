package gcptools
import (
  "encoding/json"
  "fmt"
  "os"
  "os/exec"
)

type (
  kubeconfigcontext struct {
    Name string `json:"name"`
    Context interface{} `json:"context"`
  }
  kubeconfiguser struct {
    Name string `json:"name"`
    User interface{} `json:"user"`
  }

  // returned json structure from "kubectl config view"
  kubeconfig struct {
    Kind string `json:"kind"`
    ApiVersion string `json:"apiVersion"`
    Preferences interface{} `json:"preferences"`
    Clusters []kubeconfigcluster `json:"clusters"`
    Contexts []kubeconfigcontext `json:"contexts"`
    Users []kubeconfiguser `json:"users"`
  }
)

func (kc *kubeconfig) SortClustersByActiveness() metadataGrouping {
  clusterCount := len(kc.Clusters)
  statusChan := make(chan *metadata, clusterCount)
  clusterMap := make(metadataGrouping)

  clusterMap["active"] = []*metadata{}
  clusterMap["inactive"] = []*metadata{}
  for _, cluster := range kc.Clusters {
    // can the host system handle the amount of open files this function could generate?
    // examine with ulimit -a
    go determineClusterActiveness(findProjectById, cluster, statusChan)
  }

  for i:=0; i<clusterCount; i++ {
    metadataElement := <- statusChan
    if metadataElement.Active {
      clusterMap["active"] = append(clusterMap["active"], metadataElement)
    } else {
      clusterMap["inactive"] = append(clusterMap["inactive"], metadataElement)
    }
  }

  return clusterMap
}

func determineClusterActiveness(resolve gcloudProjectResolver, cluster kubeconfigcluster, statusComm chan *metadata) {
  kcs := &metadata{ ClusterConfig: &cluster }
  queryResults, resolverErr := resolve(fmt.Sprintf("id = %s", kcs.ClusterConfig.getProjectId()))
  if resolverErr == nil && len(queryResults) > 0 {
    kcs.Active = true
    kcs.ProjectData = &queryResults[0]
  }

  statusComm <- kcs
}

func FetchKubeConfig() *kubeconfig {
  config := new(kubeconfig)
  cmd := exec.Command("bash", "-c", "kubectl config view -o json")
  out, execErr := cmd.CombinedOutput()
  if execErr != nil {
    fmt.Printf("could not run \"%s\"\n%+v\n", cmd, execErr)
    os.Exit(1)
  }

  if unmarshalErr := json.Unmarshal(out, config); unmarshalErr != nil {
    fmt.Printf("%+v\n", unmarshalErr)
    os.Exit(1)
  }

  return config
}
