
### What?
Ops Manager is a web application that you use to streamline deploying and managing Tanzu Platform Cloud Foundry (TPCF) deployments.

Ops Manager can deploy many different Tanzu software products such as Cloud Foundry, Gemfire, and Kubernetes (TKGI) as well as partner products like Jenkins. These products are represented as Tiles in Ops Manager.
A Tile is a zip file with a `.pivotal` extension containing:
- one or more BOSH release tarballs containing the product's source code and dependencies
- a metadata file used to generate a BOSH manifest and Ops Manager form elements for that product
- migration files to automatically handle property changes when upgrading to a newer Tile

## UI Walkthrough

The Ops Manager Dashboard looks like the following:

Visit https://pcf.dhaka.cf-app.com/ to login, click the "Broadcom Okta" link at the bottom ans sign in as usual. You should have read-only access to the Opsman UI.

The Ops Manager shown here has many Tiles installed:
- The BOSH Director tile comes pre-installed with your Ops Manager. Ops Manager will deploy a BOSH Director on your chosen IaaS to manage your deployments as it's first task on install and upgrade.
- The Small Footprint VMware Tanzu Application Service Tile is the Cloud Foundry deployment. This is roughly equivalent to the Cloud Foundry you deployed in earlier stories with cf-deployment. Small Footprint is used to reduce IAAS costs as it puts all the different CF vms on a smaller set of VMs as containers. It should not be use in production.
- Many of these tiles provide services (like a Mysql DB) for CF apps to use.

More information about the Dashboard UI can be found [here](hhttps://docs.vmware.com/en/VMware-Tanzu-Operations-Manager/3.0/vmware-tanzu-ops-manager/pcf-interface.html).
Tiles can be downloaded from Broadcom. We won't have entitlements to see the same customer portal but going through ERP Release Management should allow you to download tiles.

Clicking on the "BOSH Director for GCP" tile, it looks like the following:

![Tile UI](https://github.com/pivotal/cf-onboarding/raw/master/images/director-tile.png)

Tiles provide multiple configuration panes that let you select configuration values needed for the product.
For example, "Google Config", "Director Config", and "Create Available Zones" are all forms to select your Director and IaaS configuration.
After saving these forms and clicking Apply Changes on the Dashboard, Ops Manager will deploy a BOSH Director using the provided values.
All other Tiles will be deployed as a separate BOSH deployment using that Director.
For example, filling out the forms on the TAS for VMs Tile will generate a BOSH manifest and clicking Apply Changes will then deploy that manifest using the BOSH Director.
After filling out the Director forms and clicking Apply Changes, Ops Manager will run `bosh upload-release`, `bosh deploy`, and other BOSH commands you've run in previous stories and stream the output to the UI.

In addition to the UI, Ops Manager provides an [API](https://docs.vmware.com/en/VMware-Tanzu-Operations-Manager/3.0/vmware-tanzu-ops-manager/install-ops-man-api.html) to allow users to configure Tiles via scripts.
The UI provides a nice onboarding experience, but we still encourage customers to setup Concourse pipelines to automatically configure their tiles.
This ensures configuration is checked into source control and is reproducible in case something goes wrong later on.
The [om CLI](https://github.com/pivotal-cf/om) can be used to programmatically perform operations in Ops Manager rather than clicking around the UI. For example, you can get current product config, make some changes, and update Ops Manager to reflect those changes using the following commands:

```
# This will output a STAGED-CONFIG.yml
om -u USERNAME -p PASSWORD -k -t https://OPS-MAN-URL-OR-IP/uaa staged-config -p TILE-NAME

# Whatever the file ended up being named in the previous command
vim STAGED-CONFIG.yml

# Re-upload it!
om -u USERNAME -p PASSWORD -k -t https://OPS-MAN-URL-OR-IP/uaa configure-product -c STAGED-CONFIG.yml
```

## One Platform

TPCF refers to your entire IT platform, not just the Cloud Foundry deployment.
A company might require some apps to run on CF and others on Kubernetes.
These apps might need Databases, logging, and other services to be deployed as well.
Ops Manager provides a single Dashboard to manage this wide array of components.

Links:
- [TPCF Docs](https://docs.vmware.com/en/VMware-Tanzu-Application-Service/index.html)
