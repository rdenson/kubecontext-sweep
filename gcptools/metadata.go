package gcptools

type (
  metadata struct {
    //kubectl cluster config data
    ClusterConfig *kubeconfigcluster
    //google cloud api data
    ProjectData *gcoudproject
    Active bool
  }

  metadataGrouping map[string][]*metadata
)
