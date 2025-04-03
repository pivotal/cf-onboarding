
### What?
Your application runs in a container hosted on a Diego Cell. A Cloud Foundry deployment may have many cells distributed among multiple Availability Zones. Diego will automatically balance the applications you deploy across the defined AZs. If an AZ went down (along with your application), Diego would start a new instance of your application in a healthy cell in a different AZ. Depending on demand, you may want to scale your application horizontally (more instances) and/or vertically (more disk & memory).

You already have two instances of `dora`. Now we'll scale your app _horizontally_ by bumping your instance count up to five and scale _vertically_ by increasing your memory to 512M.

### How?
1. Refer to the **[`cf scale` docs](https://docs.cloudfoundry.org/devguide/deploy-apps/cf-scale.html)** to scale your app vertically and horizontally.
1. In another terminal buffer, run `watch cf app dora` to watch the updates happen.

### Expected Result
`cf app dora` should show five instances and the disk and memory they are each allotted.
`cf env dora` should show my new disk and memory limits listed under VCAP_APPLICATION.

### Resources
[Scaling apps on CF](https://docs.cloudfoundry.org/devguide/deploy-apps/cf-scale.html)
