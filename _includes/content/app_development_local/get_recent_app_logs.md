
### What?
Cloud Foundry applications are expected to log to [stdout and stderr](https://en.wikipedia.org/wiki/Standard_streams). This is part of [Twelve-Factor app design](https://12factor.net/logs).

A set of Cloud Foundry components (cumulatively named [Loggregator](https://github.com/cloudfoundry/loggregator)) provide a stream of log output from your app and from system components that interact with your app during updates and execution. Logs are gathered and stored in a best-effort manner. If a client is unable to consume log lines quickly enough, logs will be lost. The `--recent` flag will print the logs that are being buffered.

The log output has a [specific format](https://docs.cloudfoundry.org/devguide/deploy-apps/streaming-logs.html). It's important to note that when your application logs, the instance number is prepended to the message (i.e. [APP/*instance_number*]).

System components logs will also appear in relation to your application. Below is a description of a few important ones:
- The stager `STG` is responsible for using a Buildpack to create a droplet.
- The Diego cell `CELL` is responsible for running your application instance.
- The Router `RTR` logs your application has responded to an HTTP request.
- [See more components](https://docs.cloudfoundry.org/devguide/deploy-apps/streaming-logs.html#format)

### How?
`cf logs dora --recent`

### Expected Result
The output will look something like:
```
2017-12-05T21:37:29.96-0500 [STG/0]      OUT Uploaded build artifacts cache (109B)
2017-12-05T21:37:33.77-0500 [STG/0]      OUT Uploaded droplet (75.7M)
2017-12-05T21:37:33.78-0500 [STG/0]      OUT Uploading complete
2017-12-05T21:37:34.11-0500 [CELL/0]     OUT Creating container
2017-12-05T21:37:35.07-0500 [CELL/0]     OUT Successfully created container
2017-12-05T21:37:39.40-0500 [CELL/0]     OUT Starting health monitoring of container
2017-12-05T21:37:42.08-0500 [APP/0]      OUT 2016-12-06 02:37:42.078  INFO 14 --- [           main] pertySourceApplicationContextInitializer : Adding 'cloud' PropertySource to ApplicationContext
2016-12-05T21:45:25.23-0500 [RTR/0]      OUT dora.pcfdev.io - "GET /favicon.ico HTTP/1.1" 200 0 946 "http://dora.local.pcfdev.io/service" ....
```

### Resources
[Docs: Logs for CF components](https://docs.cloudfoundry.org/running/managing-cf/logging.html)
[Docs: Logs for applications](https://docs.cloudfoundry.org/devguide/deploy-apps/streaming-logs.html)