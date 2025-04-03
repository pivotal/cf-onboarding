
### What?
As previously discussed, **[bbl](https://github.com/cloudfoundry/bosh-bootloader)** is a command line utility for standing up a BOSH Director (the VM and persistent disk) on an IAAS.
In the past, you would have used Terraform (or on AWS, a CloudFormation template) to create the networks and VMs and [bosh create-env](https://bosh.io/docs/init-google.html) to set up your Director.
Under the hood that's what `bbl` uses too.

#### Digging in: What does `bbl` do, actually?
`bbl`'s responsbilities can be broken down into the following things:
1. IaaS paving
2. Bootstrapping a BOSH director and jumpbox
3. Uploading a cloud-config to the BOSH director

[Let's break it down.](https://giphy.com/gifs/L253rIFQkyXSg/tile)

**IaaS Paving**
Before you deploy a BOSH director or CF (or anything else), you'll need to set up some back IaaS resources. This varies a bit by IaaS (AWS, GCP, etc.), but most of the time that includes the following things:
- A private network for deploying your VMs. The network will typically include a subnet for each **availability zone** (or AZ) you're hoping to use. Each subnet is defined by a specific range of IP addresses that it's allowed to use. For example, if you want to use 3 AZs, you'll end up with three subnets defined by IP ranges 10.0.32.0/24, 10.0.48.0/24, and 10.0.64.0/24. Each subnet would live in a different AZ, so that if one AZ were to experience issues, software deployed to the other AZs would continue to function.
- Routing rules. These are closely related to your network and subnet definitions, but they express how traffic **from your VMs** gets to the public internet, or even to other VMs. For example, you'll see rules that say "Traffic to 10.0.0.0/16" -- meaning, all private IP addresses -- "should get routed to the gateway for my private network."
- Firewall rules. These allow or disallow network traffic to your VMs from certain IP addresses. Most commonly, we want to block all traffic from public IP addresses, and then allow a few exceptions so that users can, you know, use the platform and apps deployed there.
- External IP addresses. Most VMs will use **private IPs** from your `10.0.X.Y` ranges, but a few VMs need to be available on the public internet. In our example, these will eventually be used by load balancers and the BOSH jumpbox.
You might see a few other things that are specific to a certain IaaS -- instance groups on GCP and security groups on AWS, for example -- the resources listed above are the most common across platforms.

**Bootstraping a BOSH Director and Jumpbox**
All that IaaS-paving was done for the purpose of getting a BOSH director running. HOWEVER, for super security fun time reasons, we want to deploy the director a private IP -- that means that any IP address outside of our private network will fail to send traffic to it.

Well, in that case, how do _I_ access the BOSH director from my local machine? That's where the jumpbox becomes your friend. The jumpbox is a pretty lazy VM: all it does is hang out on both a public and private IP address, and it acts as the bridge between the outside world (i.e. you) and your private network. Any time you want to communicate with the BOSH director, you have to go through the jumpbox. If that sounds like a lot of work, you wouldn't be wrong, but have no fear: `bbl` is on the case.

In the case of TPCF, customers deploy a jumpbox that gives access to the Operations Manager VM which deploys and manages the BOSH Director VM.

After running `bbl up` below, if you run `bbl print-env`, you'll see a lot of output. One of the most important lines you'll see will say something like:
```
export BOSH_ALL_PROXY=ssh+socks5://jumpbox@<some IP address>:22?private-key=/path/to/some/bosh_jumpbox_private.key
```
This line of `bash` would set environment variables that the BOSH CLI can use to discover the location and credentials for your jumpbox, so that you can tunnel through the jumpbox and issue CLI commands against the director.

So, to **actually** set all of the variables, you'd need to run:
```
eval "$(bbl print-env)"
```
from the bbl state directory.

**Uploading a cloud-config to the BOSH director**
The point of a [cloud-config](https://bosh.io/docs/cloud-config) is to hold IaaS- and environment-specific information about your deployment, so that your deployment manifest -- in this case, cf-deployment, which you'll use soon -- doesn't have to know about things like your IaaS credentials or identifiers for resources on your IaaS.

If you were setting up your environment without bbl, after paving your IaaS with Terraform, and [bootstrapping a bosh director with the `create-env` command](https://bosh.io/docs/init-google.html), you'd need to build a cloud-config by hand based on the output from your Terraform work. `bbl` bails you out of this work by constructing that cloud-config for you, and uploading it to your BOSH director.

### How?
Make sure you're in the directory that you created in the last step, with the `service-account.key.json` file generated by `gcloud`.

Run

_Note: if you specified a different default region for `gcloud init`, change "us-east1" below. `gcloud config list` will list your previously-entered defaults._

```
bbl up \
  --name FRIENDLY_ENV_NAME \
  --gcp-region us-east1 \
  --gcp-service-account-key service-account.key.json \
  --iaas gcp
```

> If you don't supply a friendly name, bbl will pick a random name for you.

...and...*<<drumroll, please>>*...it might fail. Sad panda. But if it does, the error is straightforward, so follow its instructions to fix the problem.

If you are on a Mac M1/M2 laptop the above instructions wont work on your local laptop due to the arch change. You'll need to run the commands from a docker container.

You can use the following Dockerfile:

```
FROM golang

RUN apt-get update
RUN apt-get install -y build-essential zlib1g-dev ruby ruby-dev openssl libxslt1-dev libxml2-dev libssl-dev libreadline-dev libyaml-dev libsqlite3-dev sqlite3 apt-transport-https ca-certificates gnupg
RUN echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] https://packages.cloud.google.com/apt cloud-sdk main" | tee -a /etc/apt/sources.list.d/google-cloud-sdk.list
RUN curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key --keyring /usr/share/keyrings/cloud.google.gpg add -
RUN apt-get update && apt-get install -y google-cloud-cli
RUN wget https://github.com/cloudfoundry/bosh-cli/releases/download/v7.6.2/bosh-cli-7.6.2-linux-amd64 -O /usr/local/bin/bosh && chmod +x /usr/local/bin/bosh
RUN wget https://github.com/cloudfoundry/bosh-bootloader/releases/download/v9.0.23/bbl-v9.0.23_linux_amd64 -O /usr/local/bin/bbl && chmod +x /usr/local/bin/bbl
```

Build the docker image using the platform option, for example:

```
docker build .  --platform linux/amd64 -t bbl
```

and run the container also with the platform option, for example
```
docker run --platform linux/amd64 -v $(pwd):/workspace -w /workspace -it bbl
```

**Pro-tip:** if you're using iTerm, hold down the command-⌘ key and click on the link to open it directly from the prompt.

### Expected Result
If your `bbl up` succeeds, it will have printed the list of steps it has taken to spin up your brand-spanking new BOSH Director and generated a whole slew of files and directories that represent bbl's internal state in the current working directory. Those directories and files correspond, more or less, to the different steps outlined above. Just as importantly, they all have some interface for customizing your BOSH setup. No worries, we won't make you do that here.
- `bbl-state.json`. This file is very important as it contains credentials and other metadata related to your BOSH director and infrastructure.
- `terraform`. This directory includes the terraform templates that bbl used to pave your IaaS.
- `jumpbox-deployment`, `create-jumpbox.sh`, and `delete-jumpbox.sh` -- for bootstrapping a jumpbox. The shell scripts actually contain the `create-env` commands that bbl used to set up your jumpbox.
- `bosh-deployment`, `create-director.sh`, and `delete-director.sh` -- for bootstrapping a BOSH director. The shell scripts actually contain the `create-env` commands that bbl used to set up your director.
- `cloud-config`, which includes the cloud-config bbl generates for you.

**Financially important:** Please do not expose your GCP credentials by pushing `bbl-state.json` to Github. You'd be surprised how quickly bitcoin mining VMs pop up. This costs the company beaucoup moolah.

### Resources
[Docs: Using bosh create-env](https://bosh.io/docs/cli-v2.html#create-env)
[Docs: cloud-configs](https://bosh.io/docs/cloud-config)
[Docs: What is cf-deployment?](https://github.com/cloudfoundry/cf-deployment/blob/master/texts/deployment-guide.md)
[Video: Creating Cloud Environments](https://drive.google.com/file/d/1ltNggSltyLdRs8orh-qCzZTM_EAx3x7h/view?usp=sharing)
  - Note: your onboarding week facilitator may schedule this talk in-person if you're in SF",bbl
