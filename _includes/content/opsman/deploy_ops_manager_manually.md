
### Before You Go Any Further

**This story is optional.** There are easier ways to get an Ops Manager, and you should only do this if you are terribly interested in it. If you already opted to use a Toolsmiths environment, you especially needn't trouble yourself with this and should enjoy the easy life of on-demand environment provisioning.

Okay, now that you've been warned...

### Ops Manager Interface
[Understand how an Operator uses Ops Manager](https://docs.vmware.com/en/VMware-Tanzu-Operations-Manager/3.0/vmware-tanzu-ops-manager/pcf-interface.html)

## Deploying on GCP
While we use BOSH to deploy a director and many of our services, Ops Manager is currently not deployed with BOSH.

The first step to installing Ops Manager is to set up your infrastructure.
There are some terraform scripts that will do this for you that you can download from Pivotal Network.
Go to [network.pivotal.io](https://network.pivotal.io) and make sure you're signed in to PivNet (if you don't have an account yet, click "Register" to create one.).
Next, search for `VMware Tanzu Application Service for VMs`. Once you open the product, you should see GCP Terraform templates in "Release Download Files".

Unzip the file to get the terraform templates. You can take a look at the README to see how the whole repo works, but the important bit is in the section titled "Var File". Use those instructions to build a var file for terraform. **Take care to notice some of the following gotcha's.**
1. The `project` variable needs to equal the ID of your GCP project. That's not the same thing as the _name_ of your GCP project. To find the ID, click on the project dropdown to the left of the search bar. You should say a table mapping project names to project IDs.
2. The `opsman_image_url` bears some explanation. As it's name implies, it's a URL to an OpsMan image. In the GCP case specifically, that means tha the OpsManager team has already uploaded an image in a global GCP bucket, so all you need to do is point your terraform templates at the right image. To construct the URL, you'll first need to know that all buckets have a URL of the form `https://storage.googleapis.com/<bucket name>/<file name>`. To find the bucket and file name, you'll want to find the "Pivotal Cloud Foundry Operations Manager" product on PivNet, and then download "Ops Manager for GCP". Surprise! What you actually got was a PDF that includes the bucket and file name. Note that there are different buckets for different regions -- pick the one that makes the most sense. You can test your url by running `curl --head URL` -- if you get a 200, you're in good shape.
3. `sevice_account_key` needs to contain GCP credentials that have the "owner" [Identity and Access Management (IAM) role](https://cloud.google.com/compute/docs/access/iam). If you are reusing the service account key from the OSS GCP onboarding epic then you will need to add the "owner" role to it. This can be done by adding another iam-policy-binding to it using `gcloud projects add-iam-policy-binding PROJECT_ID --member 'serviceAccount:ACCOUNT_NAME@PROJECT_ID.iam.gserviceaccount.com' --role 'roles/owner'`. You'll then paste the contents of the service account key file into the var file. Make sure to keep the `<<SERVICE_ACCOUNT_KEY` and `SERVICE_ACCOUNT_KEY` tokens, as those designate the beginning and end of the creds.

4. `dns_suffix` is like the `system_domain` from the OSS GCP track -- it should be some subdomain of `cf-onboarding.com` and is used to represent your particular environment.
5. For SSL configuration, let's use self-signed certificates. That means you'll set values for `ssl_cert` and `ssl_private_key`. To generate SSL keys, run `openssl req -x509 -newkey rsa:2048 -keyout key.pem -out cert.pem -nodes`. When you generate the certificates, make sure that the "Common Name" field matches your `dns_suffix`.
6. Leave `buckets_location`, `opsman_storage_bucket_count`, `create_iam_service_account_members` blank.

Finally, you'll run
```
terraform init
terraform apply
```

Terraform will not only pave your IaaS for you -- it will also deploy the OpsMan VM for you. However, before you can use the provided domain to access OpsMan, you need to do one last manual DNS hookup:
1. In the "Cloud DNS" tab in the GCP console, click on the zone that matches your full OpsMan domain.
2. You should see an `NS` record with four domains in the `Data` field. Copy them.
3. Go back to the "Cloud DNS" and navigate to the zone that matches your DNS suffix (it should look like `<env-name>.<project-name>.cf-app.com`).
4. Add a record set. It should be an `NS` record, you'll need to add each of the four domains you copied as a new name server.
5. Test that it works by running `nslookup pcf.<env_name>.<dns_suffix>`. If it returns an IP address, you've successfully set up DNS.

Congrats! You've got OpsManager up and running.

Before you can use it, you need to configure authentication. Your best bet is to use "Internal Authentication." Create a user with whichever useername, password, and decryption passphrase you want (you can leave the http proxy information blank).

### Expected Results
OpsManager is running and you can log in.

### Relevant Teams
- #ops-manager Slack
- #pas-releng Slack

### Resources
- https://docs.vmware.com/en/VMware-Tanzu-Operations-Manager/3.0/vmware-tanzu-ops-manager/gcp-index.html
