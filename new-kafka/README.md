## Deployment (Kubernetes):

first you have to deploy the cr & operator:

```sh
    helm install my-strimzi-kafka-operator oci://quay.io/strimzi-helm/strimzi-kafka-operator --version 0.42.0 
```
the you have to deploy the `Kafka` deployment.
file is given in this direcoty name `kafka-deploy.yaml`


for authentication we will create a user for it. to create a user we will use `KafkaUser` . to create user, deploy `kafka-user.yaml` . make sure you enable the user operator in `kafka-deployment.yaml`,. this will create a secret with the name of user.

note: if you want to provide custom poassword then you have to create a secret (see `custom-pass-secret.yaml`). then you have set the password in `kafka-user.yaml` on ref of `custom-pass-secret.yaml`

```yaml
  entityOperator:
    userOperator: {}
```