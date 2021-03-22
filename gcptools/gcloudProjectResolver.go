package gcptools
import (
  "encoding/json"
  "fmt"
  "os/exec"
)

type gcloudProjectResolver func(string) (gcloudprojects, error)

func findProjectById(filter string) (gcloudprojects, error) {
  var results gcloudprojects
  gcloudProjectList := fmt.Sprintf(
    "gcloud projects list --filter=\"%s\" --format json",
    filter,
  )
  cmd := exec.Command("bash", "-c", gcloudProjectList)
  out, execErr := cmd.CombinedOutput()
  if execErr != nil {
    return nil, fmt.Errorf("could not execute: %s\n%+v\n", gcloudProjectList, execErr)
  }

  if unmarshalErr := json.Unmarshal(out, &results); unmarshalErr != nil {
    return nil, fmt.Errorf("error unmarshaling data from: %s\n%+v\n", gcloudProjectList, execErr)
  }

  return results, nil
}
