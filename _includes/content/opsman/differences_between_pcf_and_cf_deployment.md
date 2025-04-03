
### CF vs PCF

In the stories earlier in onboarding week, you used `bbl up` to create some GCP infrastructure and BOSH + cf-deployment to deploy a Cloud Foundry.
Here's a quick rundown of the similarities and differences between Ops Manager + TAS for VMs Tile and cf-deployment.

### Technical Similarities and Differences

Similarities:
- Both use BOSH under the hood to deploy software
- Many products can be deployed with either OSS tooling or Ops Manager, e.g. Cloud Foundry or MySQL
- The App Developer CLI experience such as `cf push` is the same, although TPCF adds web apps like [Apps Manager](https://docs.vmware.com/en/VMware-Tanzu-Application-Service/6.0/tas-for-vms/toc-using-apps-manager-index.html) to streamline this experience
- We would encourage both TPCF and OSS users to configure their environments with a Concourse pipeline

Differences:
- Ops Manager uses UI forms or API calls to configure each product while cf-deployment uses a list of ops-files
- Ops Manager Tiles include migrations to automatically handle changes like BOSH property renames when upgrading to newer versions
  - cf-deployment users would need to carefully read release notes to handle any manifest changes
- Ops Manager provides input verifiers to perform input validation prior to deploying, such as ensuring DNS records exist or Database credentials are valid
- PCF provides long-term support (LTS) for Pivotal customers
  - Many customers have complicated change management processes which require new features to be carefully vetted before deploying to production
  - These customers often stay on older versions of PCF products for many months after newer versions have been released
  - LTS means that Tanzu will backport security fixes to previously released versions of TPCF products, while new features are only added to the latest release version
  - cf-deployment users will need to upgrade to the latest release in order to get these security fixes
- Ops Manager provides a consistent experience to install many different products
  - While lots of work has gone into making deploying Cloud Foundry with cf-deployment a nice experience, the quality of documentation on deploying other BOSH releases is highly variable
- Ops Manager attempts to abstract away the BOSH implementation details, although you will still interact with BOSH when debugging
  - This added layer of abstraction can make debugging more difficult
- Clicking Apply Changes in Ops Manager is typically slower than running `bosh deploy`
  - Ops Manager performs many safety checks to ensure everything is in a good state, but this takes additional time
- Tile authors often only expose a subset of BOSH properties via the UI
  - A smaller set of properties to configure makes it easier to get something working initially, but can prevent you from configuring more advanced features
- Ops Manager does not allow you to change the deployment topology (which jobs are located on which VMs)
  - In cf-deployment you can add an ops-file to co-locate a new job on an existing VM, Ops Manager requires the Tile author to make changes to their Tile to allow this
  - Some Tiles like TAS for VMs provide both a full HA version and a "Small Footprint" version to co-locate jobs on a smaller number of VMs

### Ecosystem Similarities and Differences

Similarities:
- Most (if not all) features of core CF components like Diego, UAA, and CAPI are available in both TAS for VMs and cf-deployment
- Some customers use a mix of cf-deployment and Ops Manager across their organization
- Engineers often contribute to both cf-deployment and Ops Manager Tiles
  - e.g. a Diego engineer might submit a PR to add an ops-file to cf-deployment to toggle on a new feature, then submit a PR to the TAS for VMs Tile to add a checkbox to toggle that same feature

Differences:
- cf-deployment and associated tooling (e.g. bbl) are open-source while Ops Manager and most Tiles are closed-source
- The Cloud Foundry Foundation owns cf-deployment (along with all other projects in the `cloudfoundry` GitHub org) while Broadcom owns Ops Manager and many Tiles
- Some Cloud Foundry teams have engineers who are employed by other companies within the CF Foundation like IBM and SAP. These engineers will contribute to open-source components, but generally will not contribute to Broadcom-owned software
- Users download open-source BOSH releases and stemcells from [bosh.io](https://bosh.io) while Broadcom commercial customers will download tiles and stemcells from the Broadcom Portal (RMT).
  - The Broadcom Portal (RMT) also contains Partner Services which are Tiles jointly developed by Broadcom and partner companies like Datadog and New Relic
