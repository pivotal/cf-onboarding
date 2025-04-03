
### What?
Behind the scenes the command line depends on several components in Cloud Foundry. If a cf CLI command fails or produces unexpected results, you can re-run it with HTTP tracing enabled to view raw requests and responses between the cf CLI and the other components. The majority of those calls will be to the Cloud Controller REST API, but for logging the CLI hits the Loggregator.

### How?
Enable HTTP tracing by setting the CF_TRACE environment variable (`CF_TRACE=true`) or using the "verbose" flag on any cf command (`-v`).

Try this now by running `cf app dora -v` OR `CF_TRACE=true cf app dora`

### Expected Result
You should see each of the requests and responses that occurs behind the scenes when running those commands. Give it a shot with other commands to learn more about what they're doing under the hood.

### Resources
[Docs: Trace Cloud Controller REST API Calls](https://docs.cloudfoundry.org/devguide/deploy-apps/troubleshoot-app-health.html#trace)