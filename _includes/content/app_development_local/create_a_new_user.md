
### What?
Cloud Foundry uses role-based access control, with each role granting permissions in either an [organization](https://docs.cloudfoundry.org/concepts/roles.html#orgs) or an [application space](https://docs.cloudfoundry.org/concepts/roles.html#spaces). **[View the types of roles and their abilities](https://docs.cloudfoundry.org/concepts/roles.html#roles)**. Without a certain role in an org and space you won't be able to perform specific actions.

### How?
1. Create a new user
`cf create-user some-user some-password`
1. Login as that user
`cf auth some-user some-password`
1.  And view orgs
`cf orgs`

Not much there, huh? By default `some-user` cannot view or perform any actions, so as an admin, you'll have to assign some roles. Log back in as the admin user and update `some-user`'s roles so they can push an application in the onboarding-org / onboarding-space org and space.
1. Login back in as admin
1. And give some-user the SpaceDeveloper role for your org and space
`cf set-space-role some-user NAME-OF-YOUR-ORG NAME-OF-YOUR-SPACE SpaceDeveloper`

### Expected Result
When logged in as `some-user`, you can see your org's name in your list of orgs and your space's name in your list of spaces.

### Resources
[Docs: Role permissions](https://docs.cloudfoundry.org/concepts/roles.html)
[Docs: Setting user roles](https://docs.cloudfoundry.org/concepts/roles.html#roles)

### Relevant Repos and Teams
**CLI:** [cloudfoundry/cli](https://github.com/cloudfoundry/cli)
**CAPI:** [cloudfoundry/cloud_controller_ng](https://github.com/cloudfoundry/cloud_controller_ng)