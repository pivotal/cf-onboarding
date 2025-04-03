
### What?
You're going to deploy your BOSH Director to [Google Cloud Platform (GCP)](https://cloud.google.com). GCP is significantly cheaper than AWS given Broadcom's contracts.

The `gcloud` CLI manages authentication, local configuration, developer workflow, and interactions with the Cloud Platform APIs. Download and install the [Google SDK](https://cloud.google.com/sdk) to get `gcloud`, as you'll need it to `bbl up`.

This is another "fun-sized" story (you know you'll miss them once you hit the 6-hour ones....)

### How?
Follow **[these instructions](https://cloud.google.com/sdk/docs/downloads-interactive)** to install Google Cloud SDK on your workstation.

Once you've done that, run `gcloud init` in your terminal.

You should have received an invitation to join a GCP project from your facilitator. Use your @vmware.com email address and the name of your assigned project to configure the tool (the project name is in the GCP dashboard menu bar and the URL).

### Expected Result
Run `gcloud version` => "Google Cloud SDK 203.0.0 ..." or higher

Run `gcloud config list` =>
```
[core]
account = your-email@vmware.com
disable_usage_reporting = False
project = PROJECT_ID

Your active configuration is: [default]
```

### Resources
[Docs: Comparison GCP <> AWS](https://cloud.google.com/docs/compare/aws/)
[Internal Doc: The Great GCP Migration of 2016](https://docs.google.com/document/d/1ze6znVK32UlpsmGXHmMn4ZAAAvuUCK0yVPXhXH74-n0/edit#heading=h.x5ivrrqjuddi)
[Internal Doc: The Great GCP Migration of 2016](https://docs.google.com/document/d/1ze6znVK32UlpsmGXHmMn4ZAAAvuUCK0yVPXhXH74-n0/edit#heading=h.x5ivrrqjuddi)
