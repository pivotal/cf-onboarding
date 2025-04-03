
### What?
When a BOSH VM is healthy, it is listed with the status "running". Let's use **[Monit](https://mmonit.com/monit/)** to trigger a state change.

### How?
1. Run `watch bosh vms`. (If you don't have `watch` installed, run `brew install watch`)
1. In another buffer or tab, `bosh ssh` into one of your Diego cells
1. Run `sudo -i` to run as root
1. Run `monit summary` (need root access to do this)
1. Run `monit stop all`
1. Observe the state of the bosh VM in question.
1. Run `monit start all`

### Expected Result
The Diego job for the cell you SSHed to should be listed as `failing` as soon as you stop the Monit jobs. When you run `monit start all` it should return to the `running` state.

### Resources
[Docs: BOSH CLI health commands](https://bosh.io/docs/sysadmin-commands.html#health)
[Docs: BOSH Job Lifecycle](http://bosh.io/docs/job-lifecycle.html)
[Docs: Monit](https://mmonit.com/monit/)