
### What?
Your application runs in a container hosted on a [Diego Cell](How the Diego Auction Allocates Jobs). A Cloud Foundry deployment may have many cells distributed among multiple Availability Zones (AZs). Diego will automatically balance the applications you deploy across the defined zones. If an AZ goes down (along with your application), Diego will start a new instance of your application in a healthy cell in a different AZ. Depending on demand, you may want to scale your application horizontally (more instances) and/or vertically (more disk & memory).

You already have two instances of `dora`. Now we'll scale your app _horizontally_ by bumping your instance count up to five and scale it _vertically_ by increasing your per instance memory to 512M.

### How?
1. Run `watch cf app dora` to watch the updates happen.
1. Refer to the **[`cf scale` docs](https://docs.run.pivotal.io/devguide/deploy-apps/cf-scale.html)** to scale your app vertically and horizontally in another terminal tab or buffer.

### Expected Result
`cf app dora` should show five instances and the disk and memory they are each allotted.
`cf env dora` should show new disk and memory limits listed under VCAP_APPLICATION.

### Resources
[Docs: How the Diego Auction Allocates Jobs](https://docs.cloudfoundry.org/concepts/diego/diego-auction.html)
[Docs: Scaling apps on CF](https://docs.run.pivotal.io/devguide/deploy-apps/cf-scale.html)