
## Decision time!

**Note:** You may have already created a Toolsmiths environment as part of the **Basic BOSH Knowledge** story. If so, feel free to use that one!

By this point you're probably ready to start playing around with an Ops Manager environment.
You have two options for getting one:
- Deploy one from scratch OR
- Claim an already deployed environment

If you want to experience the workflow a customer would go through and want learn a bit more about creating GCP environments, try setting up an environment manually.
The next few stories in the backlog will guide you through this process.
If you've had enough of the IaaS stuff for this week, the following steps will show you how to grab an already deployed Ops Manager with TAS for VMs installed.

## Optional: Claim a Toolsmiths environment

**Warning!** To keep costs down, Toolsmiths environments are automatically destroyed after 24 hours

The CF Toolsmiths team ([#cf-toolsmiths slack](https://pivotal.slack.com/messages/C0563B53F/)) maintains a pool of already deployed Ops Manager + TAS for VMs environments for teams across the org to grab for testing.
You can see a dashboard of all available environments [here](https://environments.toolsmiths.cf-app.com/pooled_gcp_environments).
Note: you must be on the VPN to see this app.

The easiest way to claim an environment to play around with is to use the [smith CLI](https://github.com/pivotal/smith).
The [README](https://github.com/pivotal/smith#smith---the-community-cli-for-toolsmiths) gives instructions for getting the CLI installed and configured.
Your `.smith-token-hook.sh` file should contain:

```
#!/bin/bash
lpass show --notes 'Shared-CF SF Onboarding/Toolsmiths API Token'
```
You should have been given access to this LastPass folder at the start of the week, let your facilitators know if you still don't have access.

You should now be able to claim an environment using the steps described [here](https://github.com/pivotal/smith/blob/master/EXAMPLES.md).
We recommend grabbing the latest TAS for VMs version so you can play around with all the latest and greatest features.
Talk to your facilitators or ask the Toolsmiths team in slack for help if you run into any issues and consider giving the Toolsmiths a Thank You for providing such a valuable service!
