package gcptools
import (
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
)

func (kcc kubeconfigcluster) getProjectId() string {
  return strings.Split(kcc.Name, "_")[1]
}

func (kcc kubeconfigcluster) getRegion() string {
  return strings.Split(kcc.Name, "_")[2]
}

func (kcc kubeconfigcluster) getClusterName() string {
  return strings.Split(kcc.Name, "_")[3]
}
