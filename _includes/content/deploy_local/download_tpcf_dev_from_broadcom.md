
### What?
[TPCF Dev](https://tanzu.vmware.com/application-service) is a lightweight [Cloud Foundry](https://docs.cloudfoundry.org/concepts/overview.html) installation that runs on a single virtual machine on your workstation.

With TPCF Dev, you can get the Cloud Foundry developer experience (pushing, scaling, binding, makes it harder, better, faster, stronger) without going through the operator experience... which is admittedly an acquired taste.

It's open source, but from [Tanzu Platform Cloud Foundry (TPCF)](https://tanzu.vmware.com/tanzu) not the [Cloud Foundry Foundation (CFF)](https://www.cloudfoundry.org/foundation/), and the cf CLI plugin that makes it any fun to use is only available through [Pivotal Network](https://network.pivotal.io/).

### How?
**[Download PCF Dev from Pivotal Network](https://network.pivotal.io/products/pcfdev)** (you'll need a [Pivotal Network Account](https://network.pivotal.io/registrations/new) if you don't have one already).

The content of this zip file is not PCF Dev itself, but a script that installs the PCF Dev plugin for the cf CLI. It will only download the PCF Dev image from [S3](https://aws.amazon.com/s3/) when you run `cf dev start` in a later story.

**Note:** Later instructions are written for Linux, but you should choose the files specific to your operating system (hint: Darwin means Mac). Also, make sure to follow the **Usage Instructions** carefully (you should end up with one file, not two).

### Expected Result
You have a zip file and a new found sense of accomplishment.

### Resources
[Pivotal.io: What is PCF Dev?](https://pivotal.io/pcf-dev)
[Doc: What is a virtual machine?](https://azure.microsoft.com/en-us/overview/what-is-a-virtual-machine/) (only incidentally by Microsoft Azure, make nothing of it)

### Relevant Repos and Teams
**PCF Dev:** [pivotal-cf/pcfdev](https://github.com/cloudfoundry-attic/cfdev)
**PCF Dev CLI plugin:** [pivotal-cf/pcfdev-cli](https://github.com/pivotal-cf/pcfdev-cli)