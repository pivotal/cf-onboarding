
### What?
If you created any other GCP resources please tear those down as well.

### How?
If you deployed the environment via `bbl`:
- Run `bosh -d <DEPLOYMENT-NAME> delete-deployment` to delete all deployed VMs
- `bbl destroy --gcp-service-account-key=service-account.key.json` to delete the Director and any IaaS resources

If you deployed the environment via Ops Manager:
- Click Username dropdown in top-right > Settings > Advanced > Delete this Installation
- Use `terraform destroy` to destroy any IaaS resources

If you lost the BBL or Terraform state files, you can either:
- Delete the resources manually via the GCP UI
- Use the [leftovers](https://github.com/genevieve/leftovers) CLI to delete IaaS resources which match a given tag
