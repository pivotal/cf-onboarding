
### What?

Deploy the Pivotal flavored Kubernetes with the aide of BOSH and Ops Manager

### Caveat
You will need to prepare Ops Manager on GCP differently than you prepare it for TAS for VMs.

### Visualizing the Architecture
![pks-architecture](https://docs.pivotal.io/runtimes/pks/1-3/images/pks-overview.png)

Summary of steps
1. Setup Ops Manager and deploy BOSH Director.
2. Partially setup PKS API Load Balancer.
3. Upload PKS tile, configure the tile, and deploy PKS Control Plane VM with some chosen API wildcard domain.
4. After PKS tile is deployed, update the PKS API load balancer with the IP address of the control plane VM that BOSH deployed on GCP.
5. Use `pks` CLI to deploy cluster
6. Use `kubectl` to deploy workload on cluster

You can find very detailed steps to setup everything on GCP [here](https://docs.pivotal.io/runtimes/pks/1-3/gcp-index.html).

Alternatively, if you can't be bothered, toolsmiths provides a [Deploy PKS on GCP](https://environments.toolsmiths.cf-app.com/deploy) pipeline.

### Expected Results
- Setup Ops Manager on GCP
- Configure and deploy PKS tile on Ops Manager
- Use `pks` CLI to create cluster
- Use `kubectl` CLI to deploy workload on cluster

### Resources
- https://pivotal.io/platform/pivotal-container-service
- PCF documentation on PKS: https://docs.pivotal.io/runtimes/pks/1-3/
- Installing pks CLI: https://docs.pivotal.io/runtimes/pks/1-3/installing-pks-cli.html
- Using pks CLI: https://docs.pivotal.io/runtimes/pks/1-3/configure-api.html
- Installing kubectl CLI: https://docs.pivotal.io/runtimes/pks/1-3/installing-kubectl-cli.html