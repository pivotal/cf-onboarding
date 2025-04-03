
### What?
By omitting the `--recent` from your `cf logs dora` command, Loggregator will stream logs from the application in realtime. You'll use this most often as a diagnosing tool. Let's trigger some app activity to verify that it shows up in the logs.

### How?
1. Run `cf logs dora`
1. In another terminal buffer, run `watch cf app dora`. (If you don't have `watch` installed, you can get it by running `brew install watch`).
1. Visit the dora endpoint `/sigterm/KILL`

### Expected Result
Cloud Foundry applications are monitored with a constant health check that ensure they are listening on a specific port. When your application stops listening it'll be automatically restarted.

You can see this in the logs tagged with `CELL` and `API`

When your application is down, in the `RTR` logs, you will see that HTTP requests will return a 502 Bad Gateway.

Refer to the **[dora app README.md](https://github.com/cloudfoundry/cf-acceptance-tests/tree/master/assets/dora)** to find other useful endpoints to experiment with. It's good to know what common logs, warnings, and errors look like so you can recognize them in the wild.