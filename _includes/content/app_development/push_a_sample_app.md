
### What?
This is the "magic" that makes TAS such a compelling product.

Cloud Foundry applications are deployed by `cf push`-ing your code or some compiled artifact. The appropriate [buildpack](https://docs.cloudfoundry.org/buildpacks/) is identified by the buildpack detect scripts, a [staging container](https://docs.cloudfoundry.org/concepts/how-applications-are-staged.html) is created to prepare your app's environment, and a droplet is produced. This droplet will be used by the cells to run your application instances. When you push an app, you can optionally include an [application manifest.yml](https://docs.cloudfoundry.org/devguide/deploy-apps/manifest-attributes.html) that contains properties about your application's deployment.

### How?
1. Clone the **[cloudfoundry/cf-acceptance-tests](https://github.com/cloudfoundry/cf-acceptance-tests)** repo to your local workstation.
1. `cd ./cf-acceptance-tests/assets/dora`
1. **[Create a `manifest.yml` file](https://docs.cloudfoundry.org/devguide/deploy-apps/manifest-attributes.html#minimal-manifest)**
1. In the `manifest.yml`, set "dora" as the app's name
1. In the `manifest.yml`, set the [disk quota](https://docs.cloudfoundry.org/devguide/deploy-apps/manifest-attributes.html#disk-quota) to 1024M
1. In the `manifest.yml`, set the [number of instances](https://docs.cloudfoundry.org/devguide/deploy-apps/manifest-attributes.html#instances) to 2
1. In the `manifest.yml`, [generate a unique route](https://docs.cloudfoundry.org/devguide/deploy-apps/manifest-attributes.html#random-route) for your app the first time you push.
1. `cf push` your app

### Expected Result
`cf app dora` will print the app's status, including the URL. When your visit the URL in your browser you should see the words "Hi, I'm Dora!"

### Resources
[Cf Buildpacks](https://docs.cloudfoundry.org/buildpacks/)
[CF application manifests](https://docs.cloudfoundry.org/devguide/deploy-apps/manifest-attributes.html)
[How Applications are Staged](https://docs.cloudfoundry.org/concepts/how-applications-are-staged.html)
[A beginner-friendly introduction to containers, VMs, and Docker](https://medium.freecodecamp.com/a-beginner-friendly-introduction-to-containers-vms-and-docker-79a9e3e119b)
