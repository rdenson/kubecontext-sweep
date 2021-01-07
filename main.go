package main
import (
  "encoding/json"
  "fmt"
  "os"
  "os/exec"
)

type gcloudprojectlabels struct {
  Env string `json:"environment"`
  Svcpool string `json:"service_pool"`
  Subdomain string `json:"tenant_subdomain"`
  Id string `json:"uuid"`
}
type gcloudprojectparent struct {
  Id string `json:"id"`
  Type string `json:"type"`
}
type gcoudproject struct {
  Created string `json:"createTime"`
  Labels gcloudprojectlabels `json:"labels"`
  Status string `json:"lifecycleState"`
  Name string `json:"name"`
  Parent gcloudprojectparent `json:"parent"`
  Id string `json:"projectId"`
  Number string `json:"projectNumber"`
}
type gcloudprojects []gcoudproject

func findProject(filter string) gcloudprojects {
  var results gcloudprojects
  gcloudProjectList := fmt.Sprintf(
    "gcloud projects list --filter=\"%s\" --format json",
    filter,
  )
  cmd := exec.Command("bash", "-c", gcloudProjectList)
  out, execErr := cmd.CombinedOutput()
  if execErr != nil {
    fmt.Printf("could not run \"%s\"\n%+v\n", cmd, execErr)
    os.Exit(1)
  }

  if unmarshalErr := json.Unmarshal(out, &results); unmarshalErr != nil {
    fmt.Printf("%+v\n", unmarshalErr)
    os.Exit(1)
  }

  return results
}

func main() {
  kubeConfig := getKubeConfig()
  fmt.Printf(
    "your local kubectl configuration references...\n  %d clusters\n  %d contexts\n  %d users\n",
    len(kubeConfig.Clusters),
    len(kubeConfig.Contexts),
    len(kubeConfig.Users),
  )

  sortedClusters := sortClustersByActiveness(kubeConfig.Clusters)
  fmt.Printf("of the clusters known to your local configuration,\n  %d are active\n  %d are inactive\n", len(sortedClusters["active"]), len(sortedClusters["inactive"]))
}
