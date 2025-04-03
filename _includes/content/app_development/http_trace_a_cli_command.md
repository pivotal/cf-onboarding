
### What?
Behind the scenes the command line depends on several components in Cloud Foundry. If a cf CLI command fails or produces unexpected results, re-run it with HTTP tracing enabled to view raw requests and responses between the cf CLI and the other components. The majority of those calls will be to the Cloud Controller REST API, but for logging the CLI hits the Loggregator.

### How?
Enable HTTP tracing with the CF_TRACE environment variable (`CF_TRACE=true`) or the -v flag (`-v`).

Try this now with `cf app` by running `cf app dora -v`.

### Expected Result
You should see each of the requests and responses that resulted in `cf app` delivering information about your app. Now give it a shot with other cf CLI commands!

### Resources
[Trace Cloud Controller REST API Calls](https://docs.cloudfoundry.org/devguide/deploy-apps/troubleshoot-app-health.html#trace)
