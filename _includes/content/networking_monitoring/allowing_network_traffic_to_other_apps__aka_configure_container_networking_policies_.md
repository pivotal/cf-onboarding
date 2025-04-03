### What?
By default, security groups are configured not to allow app containers to address each other directly.
For example, if `app1` has an overlay IP address 10.255.11.9 and `app2` has an overlay IP address 10.255.11.14,
`app1` would fail to curl 10.255.11.14, and `app2` would fail to curl 10.255.11.9.

[Container networking features](https://github.com/cloudfoundry/cf-networking-release) allow `cf` users
to define network policies between apps that override security groups.

Let's see how can manage the networking security of an app that needs to communicate with other apps.

### Prerequisites

Go through the **Basic BOSH Knowledge** story to gain familiarity with BOSH, specifically BOSH CLI commands.

### How?

1) Deploy service discovery 

Service Discovery is a container networking feature that allows CF apps to communicate with each other via internal routes. 
This way appA doesn't need to know appB's overlay IP address. 
Instead appA can just make a request to an internal route that might look like `appB.apps.internal`
and the service discovery controller will translate that route into the overlay IP. 
Service Discovery is not enabled by default, but can be enabled by using this opsfile, 
[Enable Service Discovery](https://github.com/cloudfoundry/cf-deployment/blob/master/operations/enable-service-discovery.yml). 
This opsfile also seeds the internal domain `apps.internal` for you. 

The following commands download the BOSH manifest file into a tmp file, adds `enable-service-discovery.yml` to the manifest, and updates your BOSH deployment with service discovery enabled. You may want to create an additional directory under your `~/workspace` directory to contain your manifests and other such configurations. Also, ensure that your BOSH environment variables are set (this will depend on how you set up your onboarding environment).

```
  bosh -d cf manifest > /tmp/cf.yml
  bosh deploy /tmp/cf.yml -o ../cf-deployment/operations/enable-service-discovery.yml -d cf
```

To confirm an internal domain was seeded, run `cf domains` and see that there is a domain with the name `apps.internal` and that it is listed as `internal` under `details`.

2) Make sure you have two apps pushed. If don't have any apps pushed, `cf push boots -p <PATH_TO_DORA> && cf push swiper -p <PATH_TO_DORA>`.

3) Map an internal route to swiper and boots
Internal routes can only be used for CF apps to talk to each other via container networking. Internal routes are created the same way other routes are, but they are associated with an internal domain. 

```
cf map-route swiper apps.internal --hostname swiper
cf map-route boots apps.internal --hostname boots
```

Now if you run `cf app swiper` you should see the route `swiper.apps.internal` listed and if you run `cf app boots` you should see the route `boots.apps.internal` listed.

Before applying a network policy
4) In two different terminals, `cf ssh` to boots' and swiper's application containers.
5) On `boots`, run `curl swiper.apps.internal:8080`. See that the connection fails because there isn't a container networking policy. 
6) On `swiper`, run `curl boots.apps.internal:8080`. See that the connection fails because there isn't a container networking policy. 

Apply the network policy to allow access:
7) On your command line, run `cf add-network-policy boots --destination-app swiper --protocol tcp --port 8080`

See if the app containers can communicate
8) In your ssh session on boots, try this again: `curl swiper.apps.internal:8080`.  It works! Hopefully!
9) Now try the reverse, `curl boots.apps.internal:8080` on swiper. There is not policy in this direction, so it should fail!

### Expected Result
Before the network policy is added, the initial curls from both apps should fail. After applying the networking policy, the curl from boots to swiper should succeed.
But because not policy was added in the other direction, from swiper to boots, that curl should never succeed.

Poke around to see what else you can and can't do with the network policy commands.
Can you allow access to mutliple ports?
What happens if you allow access to a port that the app doesn't listen to? Is it give the same error?

### Resources
[Container networking docs](https://github.com/cloudfoundry/cf-networking-release/tree/develop/docs)
[Docs: Understanding CF Networking](https://docs.cloudfoundry.org/concepts/understand-cf-networking.html)
[Docs: Deploy Apps with CF C2C Networking](https://docs.cloudfoundry.org/devguide/deploy-apps/cf-networking.html)
[Video: CF Networking: All Your Packets are Belong to Us](https://www.youtube.com/watch?v=WrTXd42_m10)