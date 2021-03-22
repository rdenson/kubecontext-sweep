package gcptools

type (
  gcloudprojectlabels struct {
    Env string `json:"environment"`
    Svcpool string `json:"service_pool"`
    Subdomain string `json:"tenant_subdomain"`
    Id string `json:"uuid"`
  }
  gcloudprojectparent struct {
    Id string `json:"id"`
    Type string `json:"type"`
  }

  gcoudproject struct {
    Created string `json:"createTime"`
    Labels gcloudprojectlabels `json:"labels"`
    Status string `json:"lifecycleState"`
    Name string `json:"name"`
    Parent gcloudprojectparent `json:"parent"`
    Id string `json:"projectId"`
    Number string `json:"projectNumber"`
  }

  gcloudprojects []gcoudproject
)
