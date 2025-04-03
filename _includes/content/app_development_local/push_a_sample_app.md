
### What?
Cloud Foundry applications are deployed by `cf push`-ing your code or some compiled artifact. You can specify the appropriate [buildpack](https://docs.cloudfoundry.org/buildpacks/) for your app in the command line or in your app manifest, but you can also leave it blank and let the buildpack detect scripts determine which is the most appropriate.

The system will prepare your app's environment in a [staging container](https://docs.cloudfoundry.org/concepts/how-applications-are-staged.html), and from that produce an app "droplet" (e.g. tarball). This droplet will be used as the content of each new application instance the container orchestrator spins up.

When you push an app, you can optionally include an [application manifest.yml](https://docs.cloudfoundry.org/devguide/deploy-apps/manifest.html) that contains properties about your application's deployment. Having a `manifest.yml` means you don't have to type as much in your `cf push` command, and it also makes pushing an easily replicable process (which is what we're all about here).

### How?
1. Clone the **[cloudfoundry/cf-acceptance-tests](https://github.com/cloudfoundry/cf-acceptance-tests)** repo to your local workstation by running `git clone https://github.com/cloudfoundry/cf-acceptance-tests.git` in your command line.
1. Run `cd ./cf-acceptance-tests/assets/dora` to change directories.
1. **[Create a `manifest.yml` file](https://docs.cloudfoundry.org/devguide/deploy-apps/manifest.html#minimal-manifest)**
1. In the `manifest.yml`, set "dora" as the app's name
1. In the `manifest.yml`, set the [disk quota](https://docs.cloudfoundry.org/devguide/deploy-apps/manifest.html#disk-quota) to 1024M
1. In the `manifest.yml`, set the [number of instances](https://docs.cloudfoundry.org/devguide/deploy-apps/manifest.html#instances) to 2
1. In the `manifest.yml`, [generate a unique random route](https://docs.cloudfoundry.org/devguide/deploy-apps/manifest-attributes.html#random-route) for your app the first time you push.
1. `cf push` your app

**Note:** If you face problems with Ruby version, just install the latest version of Ruby and change the version in the Gemfile.

### Expected Result
`cf app dora` will print the app's status, including the URL. When your visit the URL in your browser you should see the words "Hi, I'm Dora!" What an exciting app, right? Total game-changer.

### Resources
[Docs: CF Buildpacks](https://docs.cloudfoundry.org/buildpacks/)
[Docs: CF application manifests](https://docs.cloudfoundry.org/devguide/deploy-apps/manifest.html)
[Docs: How Applications are Staged](https://docs.cloudfoundry.org/concepts/how-applications-are-staged.html)
[Blog post: A beginner-friendly introduction to containers, VMs, and Docker](https://medium.freecodecamp.com/a-beginner-friendly-introduction-to-containers-vms-and-docker-79a9e3e119b)