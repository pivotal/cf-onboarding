
### What?

The [Cloud Foundry Command Line Interface (cf CLI)](https://docs.cloudfoundry.org/cf-cli/) is a tool you use to deploy and manage your applications, users, and services in a CF environment.

### How?
**[Download and run the latest CLI release installer from GitHub](https://github.com/cloudfoundry/cli/releases)**

### Expected Result
[Open the command line](https://www.davidbaumgold.com/tutorials/command-line/) and run `cf version`. The output should look like...
```
$ cf version
cf version x.xx.x-xxxxxxx
```

Next run `cf help -a` to see all the new commands at your disposal.

### Resources
Aside from the [Official Cloud Foundry CLI Docs](http://docs.cloudfoundry.org/cf-cli/), the output from `cf help -a` is probably your most useful resource. It is, however, a little long. Use [Linux's grep command to find what you're looking for](http://www.thegeekstuff.com/2009/03/15-practical-unix-grep-command-examples) (tip number 6 in that link is especially useful).

And for those uninitiated in the dark darks of the command line, check out [this quick and dirty overview of the command line and how to use it](https://www.davidbaumgold.com/tutorials/command-line/).

### Relevant Repos and Teams
**CLI:** [cloudfoundry/cli](https://github.com/cloudfoundry/cli)