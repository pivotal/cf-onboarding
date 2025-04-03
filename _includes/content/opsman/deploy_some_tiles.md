
## What?
OpsManager uses "tiles" to deploy different products. You can find tiles by searching on [PivNet](https://network.pivotal.io). Then you can download them, upload them to your OpsManager, and finally configure and deploy them.

Some tiles have dependencies on other tiles. You should find information about that in the product description on PivNet.

To understand the workflow of an actual TPCF/OpsMan customer, choose a tile or two to install and follow the appropriate docs.

## How?
Let's start by installing the mother of them all: TAS for VMs. Go to ERP, search for "VMware Tanzu Application Service for VMs", and download the "Small Footprint TAS for VMs" -- we'll use this to save some time during the deploy. Your best bet is to choose the latest stable version (avoid releases with `alpha` or `rc` in it).

And yes, the file is rather large (as of this writing, it's 13 GB), so feel free to take a ping pong break while it's downloading.

Use the official docs as guide to configure and deploy TAS for VMs. You can the docs for TAS for VMs 2.3 [here](https://docs.pivotal.io/pivotalcf/customizing/gcp-er-config.html), but if you're using a new version than that, you can use the dropdown next to "Pivotal Cloud Foundry" to find docs for the appropriate version.

## Expected Result
You've got a running TAS for VMs.
