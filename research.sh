# get project ids that kubectl config knows about; path: clusters[*]
# formatted as below ↓
# cluster:
#   certificate-authority-data: DATA+OMITTED
#   server: https://x.x.X.x
# name: gke_fr-02zy4wy6v5bx7aaaw7ie175sv52_us-west99_default
#
# output for name in the following format:
#   ENGINE_ID_REGION_CLUSTER-NAME (delimited by "_")
kubectl config view | yq r - 'clusters[*].name' | cut -d'_' -f2
# in the case of GCP, ID is the project id; the above command hews this value
# could also use (because no yaml?):
kubectl config view -o json | jq '.clusters | map(.name)'


# OPTION
# for each project id, see if we can find it using gcloud
gcloud projects list \
  --filter="id = ${PROJECT_ID}" \
  --format json \
| jq '. | length'
# using jq, we can determine if project was found; looking for 0 (not found)


# OPTION
# prereq: yq r - 'clusters[*].cluster.server' → https://0.0.0.0
# check using the IP attached to the dictionary in the clusters list
curl -k --no-keepalive --connect-timeout 10 ${K8S_MASTER_ADDRESS}
# this should not return 0 if the cluster is not in use
# could the IP be reused? is this reliable?


# deletion commands
kubectl config delete-cluster ${NAME}
kubectl config delete-context ${NAME}
kubectl config unset name.${NAME}
