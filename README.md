# Test MS SQL Connection

A simple Golang script for testing MS SQL Connections.

## How to use the script

1. Download the the latest release.

2. Add execution permissions:

    ```bash
    chmod +x test-mssql-connection
    ```

3. Set the `DATABASE_URL`:

    ```bash
    export DATABASE_URL="sqlserver://username:password@host:port?database=dbName&param1=value"
    ```

4. Execute the script:

    ```bash
    ./test-mssql-connection
    ```

### Using Docker

```bash
export DATABASE_URL="sqlserver://username:password@host:port?database=dbName&param1=value"

docker run --rm -e DATABASE_URL=$DATABASE_URL -it charles1888/test-mssql-connection:latest
```

### Using Kubernetes

Deploy the following Kubernetes Pod:

```bash
cat << EOF | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: test-mssql-connection
spec:
  containers:
  - name: test-mssql-connection
    image: charles1888/test-mssql-connection:latest
    env:
    - name: DATABASE_URL
      value: "sqlserver://username:password@host:port?database=dbName&param1=value"
  restartPolicy: OnFailure
  nodeSelector:
    kubernetes.io/os: linux
EOF
```

Please, note the above approach is not recommended because the `DATABASE_URL` content is exposed in the Pod's manifest, but can be fine for a quick test. The recommended approach is to use Kubernetes Secrets instead:

```bash
kubectl create secret generic test-mssql-connection \
  --from-literal=DATABASE_URL='sqlserver://username:password@host:port?database=dbName&param1=value'
```

Then reference the environment variable `DATABASE_URL` from the already created secret:

```bash
cat << EOF | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: test-mssql-connection
spec:
  containers:
  - name: test-mssql-connection
    image: charles1888/test-mssql-connection:latest
    env:
    - name: DATABASE_URL
      valueFrom:
        secretKeyRef:
          name: test-mssql-connection
          key: DATABASE_URL
  restartPolicy: OnFailure
  nodeSelector:
    kubernetes.io/os: linux
EOF
```

To get the container output:

```bash
kubectl logs -f test-mssql-connection
```

Cleaning the stuff:

```bash
kubectl delete pod test-mssql-connection
kubectl delete secret test-mssql-connection
```

## Who to build the script

1. Clone this repository.

2. Install the dependencies:

    ```bash
    go mod download
    ```

3. Build the script:

    ```bash
    CGO_ENABLED=0 GOOS=linux go build -a -o test-mssql-connection
    ```
