
### What?
Cloud Foundry stores a history of events about your application. You'll generally use `cf logs` when debugging your app's behavior, but every once and awhile the select, well-formatted information contained in `cf events` is more helpful than a firehose of data.

### How?
1. `cf stop dora`
1. `cf start dora`
1. `cf events dora`

### Expected Result
You should see the list of events including the recent stop & start.

### Resources
[Docs: cf CLI](http://docs.cloudfoundry.org/cf-cli/cf-help.html)
[StackOverflow: When to use cf events?](http://stackoverflow.com/questions/24545838/when-to-use-cf-events)