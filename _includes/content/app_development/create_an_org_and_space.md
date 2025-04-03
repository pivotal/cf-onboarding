### What?
Cloud Foundry uses role-based access control, with each role granting permissions in either an **[organization](https://docs.cloudfoundry.org/concepts/roles.html#orgs)** or an **[application space](https://docs.cloudfoundry.org/concepts/roles.html#spaces)**. Let's make a new organization and associated space!

Customer Operators use this to assign different orgs and spaces to different teams and can assign quotas to manage load across their company.

### How?
1. View orgs `cf orgs` and spaces `cf spaces` (each space is scoped to an org)
1. Target the system org `cf target -o system` and view spaces `cf spaces` and apps `cf apps`
1. Create a new `onboarding-org` organization with an `onboarding-space` space

### Expected Result
When creating the organization the current user will be assigned the role of 'Org Manager'.

Run `cf org-users onboarding-org` to verify.

When creating the space the current user will be assigned two new roles in the space - 'Space Manager' and 'Space Developer'

Run `cf space-users onboarding-org onboarding-space` to verify.

### Resources
[Orgs, Spaces, Roles, and Permissions](https://docs.cloudfoundry.org/concepts/roles.html)
[Creating and managing users with the cf CLI](https://docs.cloudfoundry.org/adminguide/cli-user-management.html)

### Relevant Repos and Teams
**CLI:** [cloudfoundry/cli](https://github.com/cloudfoundry/cli)
**CAPI:** [cloudfoundry/cloud_controller_ng](https://github.com/cloudfoundry/cloud_controller_ng)
