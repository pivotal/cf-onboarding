
### What?
Cloud Foundry uses role-based access control, with each role granting permissions in either an **[organization](https://docs.cloudfoundry.org/concepts/roles.html#orgs)** or an **[application space](https://docs.cloudfoundry.org/concepts/roles.html#spaces)**. Let's make a new organization and associated space!

### How?
1. Run `cf orgs` to view a list of organizations
1. Run `cf spaces`  to view a list of spaces within the org you're currently targeting. All spaces are scoped to an org. If you target a different org you will see a different set of spaces.
1. Now, use the `cf create-org` command to create a new organization and the `cf create-space` command to create a new space within it. **Hint:** you'll need to target the org (and also the space after you create it).

### Expected Result
When creating the organization the current user will have been assigned the role of 'Org Manager' for that org. Run `cf org-users NAME-OF-YOUR-ORG` to verify.

When creating the space the current user will have been assigned two new roles for that space—'Space Manager' and 'Space Developer'. Run `cf space-users NAME-OF-YOUR-ORG NAME-OF-YOUR-SPACE` to verify.

**Note:** if the role remains unspecified after creating the org, try the following:

```
cf set-org-role admin NAME_OF_YOUR_ORG OrgManager
cf org-users NAME-OF-YOUR-ORG
```

### Resources
[Docs: Orgs, Spaces, Roles, and Permissions](https://docs.cloudfoundry.org/concepts/roles.html)
[Docs: Creating and managing users with the cf CLI](https://docs.cloudfoundry.org/adminguide/cli-user-management.html)

### Relevant Repos and Teams
**CLI:** [cloudfoundry/cli](https://github.com/cloudfoundry/cli)
**CAPI:** [cloudfoundry/cloud_controller_ng](https://github.com/cloudfoundry/cloud_controller_ng)