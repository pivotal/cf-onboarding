
### What?
Cloud Foundry offers a marketplace of services from which users can provision reserved resources on-demand. Think of a service as a factory that delivers service instances. Examples include databases on a shared or dedicated server or accounts on a Software as a Service (SaaS) application.

As an alternative, you can also provide your own service instances in the form of [User-Provided Service Instances](https://docs.cloudfoundry.org/devguide/services/user-provided.html), which enable developers to use services that are not available in their Cloud Foundry marketplace.

Let's create one backed by [RedisLabs](https://redislabs.com/).

### How?
1. Create an account with [RedisLabs](https://redislabs.com/) and generate a New Redis Subscription (select the free 30MB option).
1. As part of setting up that Redis database, set the password (by default there is no password).

### Expected Result
I can see my new user-provided service in `cf services`.

### Resources
[Creating a user-provided service](https://docs.cloudfoundry.org/devguide/services/user-provided.html)

### Relevant Repos and Teams
**CAPI:** https://github.com/cloudfoundry/cloud_controller_ng
