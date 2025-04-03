
### What?
By omitting the `--recent` from your `cf logs dora` command, Loggregator will stream logs from the application in realtime. You'll use this most often as a diagnosing tool. Let's trigger some app activity to verify that it shows up in the logs.

### How?
1. Run `cf logs dora`
1. In another terminal tab or buffer, run `watch cf app dora`
1. Keep an eye on both, and visit the dora endpoint `http://dora.local.pcfdev.io/sigterm/KILL` in your browser

**Note:** depending on how you installed your PCF Dev environment, the address may actually be `http://dora.dev.cfdev.sh/sigterm/KILL`. You can also find the route value in the output of `cf app dora` and append `sigterm/KILL`.

### Troubleshooting
If you get an error saying that `watch` is not installed on your computer, install it with brew by running `brew install watch`. If you are on a Mac and don't have brew either...I'm confused.

### Expected Result
Cloud Foundry applications are monitored with a constant health check that ensures they are listening on a specific port. When your application stops listening it will be automatically restarted. You can see this in the logs tagged with `CELL` and `API`. If you try to access your app while it is down, the logs will include `RTR` HTTP requests that return a 502 Bad Gateway error.

Refer to the **[dora app README.md](https://github.com/cloudfoundry/cf-acceptance-tests/tree/master/assets/dora)** to find other useful endpoints to experiment with. It's good to know what common logs, warnings, and errors look like so you can recognize them in the wild.