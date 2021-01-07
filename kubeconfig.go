package main
import (
  "encoding/json"
  "fmt"
  "os"
  "os/exec"
  "strings"
)

type (
  kubeconfigclusterdata struct {
    Server string `json:"server"`
    CAData string `json:"certificate-authority-data"`
  }
  kubeconfigcluster struct {
    Name string `json:"name"`
    Cluster kubeconfigclusterdata `json:"cluster"`
  }
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

func (kcc kubeconfigcluster) getProjectId() string {
  return strings.Split(kcc.Name, "_")[1]
}

func getKubeConfig() kubeconfig {
  var config kubeconfig
  cmd := exec.Command("bash", "-c", "kubectl config view -o json")
  out, execErr := cmd.CombinedOutput()
  if execErr != nil {
    fmt.Printf("could not run \"%s\"\n%+v\n", cmd, execErr)
    os.Exit(1)
  }

  if unmarshalErr := json.Unmarshal(out, &config); unmarshalErr != nil {
    fmt.Printf("%+v\n", unmarshalErr)
    os.Exit(1)
  }

  return config
}
