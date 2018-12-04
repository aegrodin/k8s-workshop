USER="aegrodin"
TOKEN="6UGpkXjOkMww3N_DlXXz0WQ9LAPwEYUpIVgtFPh2_GaXakNYgppSF3qFjIf3sqS1QhjV_rVNw2NI6NY2aABhNVCtcQy0JRfw3uEnsESiQv7YcCOpPD4aTXIlxVNzpg"
CA_PATH="/home/a.shvalov/k8scomm.ca.crt"

kubectl config set-cluster ws-${USER} \
    --embed-certs=true \
    --server=https://35.240.115.206 \
    --certificate-authority=${CA_PATH}
kubectl config set-credentials ws-${USER} --token=${TOKEN}
kubectl config set-context ws-${USER} \
    --cluster=ws-${USER} \
    --user=ws-${USER} \
    --namespace=${USER}
kubectl config use-context ws-${USER}