
### What?
You've already downloaded PCF Dev from Pivotal Network. Now you're going to run it!

### How?
Unzip the file you downloaded in the first story (e.g. `unzip pcfdev-VERSION-linux.zip`)

Run the executable binary file inside (e.g. `./pcfdev-VERSION-linux`)

From the command line, run `cf dev start`. Because you are running it for the first time it will download the image and import it to VirtualBox before starting the VM. In the future, the same command will start PCF Dev without downloading or importing it again.

**Note:** You may need to run `cf install-plugin cfdev` to ensure that you have the `cf dev` command.

FYI ...this process takes awhile. How long of a while will be influenced by your internet connectivity and bandwidth. Go get a snack or read some of the links.

When it wraps up, PCF Dev will have printed the command that you should use to login, followed by credentials for two users, `user` and `admin`.
Run the login command, choose the `admin` user, then select the `cfdev-org` org.
(The users have different permissions. We will need the `admin` permissions.)

### Expected Result
Run `cf target`. You'll see a line that says `API endpoint:   https://api.v3.pcfdev.io`, followed by your user, org, and space info.

### Troubleshooting

If there appears to be no progress for 20-30 minutes, double-check if your system has any disk space left. PCF Dev apparently does not handle this case well.

### Resources
[Tutorial: Getting Started with PCF Dev](https://tanzu.vmware.com/developer/)
[Blog post: Meet Pivotal Cloud Foundry Dev, your Ticket To Running Cloud Foundry Locally](https://tanzu.vmware.com/content/blog/meet-pcf-dev-your-ticket-to-running-cloud-foundry-locally)
[Blog post: A little diddy about binary file formats](https://betterexplained.com/articles/a-little-diddy-about-binary-file-formats/)

### Relevant Repos and Teams
**PCF Dev:** [pivotal-cf/pcfdev](https://github.com/cloudfoundry-attic/cfdev),
**Also PCF Dev:** [pivotal-cf/pcfdev-cli](https://github.com/pivotal-cf/pcfdev-cli)
