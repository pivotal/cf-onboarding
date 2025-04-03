
### What?
[BOSH](http://bosh.io/docs/) is an open source tool for deployment, release engineering, lifecycle management, and monitoring of distributed systems. It is what we use to deploy and manage Cloud Foundry.

While it is under the [Cloud Foundry Foundation umbrella](https://www.cloudfoundry.org/projects/#bosh), it is a distinct product. BOSH can be used to deploy just about any software and Cloud Foundry can (inadvisably) be deployed by something other than BOSH.

Since Cloud Foundry is a maintenance and monitoring system itself, it sometimes confuses people that there would be another maintenance and monitoring system underneath it. Think of it as the "who watches the watchmen" of [high availability](https://docs.cloudfoundry.org/concepts/high-availability.html). Or maybe "turtles all the way down". "Watchmen all the way down." Anyway...

There have been many ways to deploy BOSH over the years, but the latest, greatest, and under active development-est, is  **[bosh-bootloader (a.k.a. bbl)](https://github.com/cloudfoundry/bosh-bootloader)**.

`bbl` stands up the BOSH Director (VM and persistent disk). The BOSH Director deploys and monitors Cloud Foundry VMs and processes. And Cloud Foundry is used to deploy and monitor applications and services in the form of containerized workloads.

You'll use `bbl` to deploy your BOSH Director in a later story, but for now I've told you all of this just so you'll install it (and a few of its dependencies) on your workstation.

This isn't the most exciting first story, but all good things to those who wait, right?

### How?
Install or update the following dependencies on your local machine:
* **Golang** (run `brew install go`)
* **Terraform** is a tool for building, changing, and versioning infrastructure safely and efficiently (run `brew install terraform`)

The following dependencies can be installed through the Cloud Foundry [homebrew tap](https://github.com/cloudfoundry/homebrew-tap)
* **BOSH CLI V2** is a cli to setup and interact with the BOSH director (`brew tap cloudfoundry/tap && brew install cloudfoundry/tap/bosh-cli`)
* **bbl** is a tool for paving infrastructure in your desired IAAS as well as deploying a BOSH director (run `brew install bbl`). It will call through to the BOSH cli.
* **cf cli** is a tool for interacting with a Cloud Foundry deployment (run `brew install cf-cli`)

### Expected Result
* Run `go version` => "go version go1.20.2 darwin/amd64" or above
* `terraform --version` => "Terraform v0.11.7" or above
* `bosh -v` => "version 4.0.1-...." or above
* `bbl version` => "bbl 6.7.2 (darwin/amd64)" or above
* `cf version` => "cf version 6.37.0+...." or above

### Resources
[Docs: BOSH V2 CLI](https://bosh.io/docs/cli-v2.html)
[Docs: CF CLI](http://docs.cloudfoundry.org/cf-cli/)
[Docs: What is Cloud Foundry?](https://www.cloudfoundry.org/platform/)
[Docs: What is Terraform?](https://www.terraform.io/intro/index.html)

### Relevant Teams and Repos
Bosh Ecosystem Team #cf-bosh-core Slack
**Bosh Ecosystem:** [cloudfoundry/bosh](https://github.com/cloudfoundry/bosh)
**Bosh Ecosystem Team CLI repo:** [cloudfoundry/bosh-cli](https://github.com/cloudfoundry/bosh-cli)
**Infrastructure:** [cloudfoundry/bosh-bootloader](https://github.com/cloudfoundry/bosh-bootloader)
**CLI:** [cloudfoundry/cli](https://github.com/cloudfoundry/cli)
