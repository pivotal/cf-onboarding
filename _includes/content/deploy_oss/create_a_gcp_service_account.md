
### What?
We're going to create a **[service account](https://cloud.google.com/compute/docs/access/service-accounts)** with `gcloud`, which `bbl` can use to set up infrastructure on GCP. Service accounts authenticate applications running on your virtual machine instances to other Google Cloud Platform services, like APIs.

### How?
Create a new directory in your workspace and `cd` into it. You're going to be generating a few files over the next couple stories and it'll be good to put them somewhere logical.

Then, follow these steps to set up your GCP service account, replacing ACCOUNT_NAME with a name of your choice and PROJECT_ID with your assigned project's name:
1. Run the following command to create a new GCP service account:
```
gcloud iam service-accounts create ACCOUNT_NAME
```
1. Create a key for that account to access your project:
```
gcloud iam service-accounts keys create "service-account.key.json" --iam-account "ACCOUNT_NAME@PROJECT_ID.iam.gserviceaccount.com"
```
1. Add the "editor" [Identity and Access Management (IAM) role](https://cloud.google.com/compute/docs/access/iam) to your service account for your project:
```
gcloud projects add-iam-policy-binding PROJECT_ID --member 'serviceAccount:ACCOUNT_NAME@PROJECT_ID.iam.gserviceaccount.com' --role 'roles/editor'
```

### Expected Result
You should now have a `service-account.key.json` in your workspace.

Verify that the account is listed in your GCP dashboard under Menu > IAM & Admin > Service Accounts.

### Resources
[Docs: GCP Service Accounts](https://cloud.google.com/compute/docs/access/service-accounts)
[Docs: Identity and Access Management (IAM) Roles](https://cloud.google.com/compute/docs/access/iam)
[Docs: Comparison GCP <> AWS](https://cloud.google.com/docs/compare/aws/)
