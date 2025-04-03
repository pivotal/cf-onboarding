
### What?
The Resurrector is a plugin to the BOSH Health Monitor that is responsible for automatically recreating VMs that become inaccessible. It continuously cross-references VMs expected to be running against the VMs that are sending heartbeats. When the resurrector does not receive heartbeats for a VM for a certain period of time, it will kick off a task on the Director to try to “resurrect” that VM. The Director may do one of two things:

* create a new VM if the old VM is missing
* replace a VM if the Agent on that VM is not responding to commands

### How?
1. Run `watch bosh vms` so you can keep an eye on the effect you're having on VM state.
1. Open a second terminal buffer and `bosh ssh` into one of the Diego cells.

Killing off a BOSH agent is a little harder than it looks. This is a great thing for CF operators, but less of a good thing when creating exercises to learn about the system. For instance, try killing off an agent process:

1. Run `ps aux | grep bosh-agent` to find the BOSH agent.
1. Kill it mercilessly, `kill -9 <process id>`
1. Run `ps aux` again. Grep for the agent again. Discover that phoenix-like, there is already a new agent process with a new process id. The VM's listed state won't have even flickered. Don't quote me here, but I'm pretty sure [upstart](http://upstart.ubuntu.com/) is responsible for this sorcery.

Looks like we'll have to get creative if we're ever going to see this resurrector at work.

1. While still SSHed into a VM, notice the path to `agent.json` in the output of `ps aux` and throw some un-parseable junk in there.

### Expected Result
Watch the process choke, the VM fail, and the resurrector bring it back! If it doesn't come back within a few minutes, check to see if resurrection is turned on.

### Resources
[Docs: Configuring Health Monitor](https://bosh.io/docs/hm-config.html)
[Docs: BOSH Resurrector](http://bosh.io/docs/resurrector.html)