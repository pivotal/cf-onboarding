
**Warning:** This story is not relevant if you already opted for the Toolsmiths environment instead of manual creation. Okay, now that you've been warned...

Configure the BOSH GCP tile using this [guide](https://docs.vmware.com/en/VMware-Tanzu-Operations-Manager/3.0/vmware-tanzu-ops-manager/install-ops-man-api.html). Some tips to keep you from getting stuck:
1. When you get to "Create Availability Zones", you'll want to use the same AZs that you specified in your terraform.
2. Make sure to read in the instructions carefully when you configure the networks. Especially be careful when setting the network names: they have the form `<network-name>/<subnet-name>/<region>`

After you've configured the BOSH Director tile (it should be green to indicate that you're finished configuring the tile), you can visit the **Review Pending Changes** page and click **Apply Changes**. You'll see the same logs as if you used `bbl` or the `bosh` CLI.
Once the deployment is done, you can visit `/change_log` to review the logs of past deployment.

After the deploy succeeds, click around and see if you can find information about the BOSH Director's health. Can you find the BOSH Director's IP address? How much memory/disk/CPU is it using?

### Expected Results
OpsMan successfully deploys a BOSH Director.

### Resources
- [Manually installing Ops Manager on GCP](https://docs.vmware.com/en/VMware-Tanzu-Operations-Manager/3.0/vmware-tanzu-ops-manager/gcp-index.html)
