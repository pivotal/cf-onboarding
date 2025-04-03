
### What?
To SSH into a BOSH job, you'll use the `bosh` cli.

### How?
bosh-bootloader sets up a jumpbox that allows `bosh` commands to be run from your workstation
* `eval "$(bbl print-env)"` to target the director
* Now you will only need the VM name to use `bosh ssh`.

### Expected Result
Running `bosh -d cf ssh VM/GUID` opens a shell in your targeted machine.

### Resources
[Forum question: What's the distinction between an HTTP proxy, tunnel, and gateway?](http://stackoverflow.com/questions/10377679/whats-distinction-of-http-proxy-tunnel-gateway)