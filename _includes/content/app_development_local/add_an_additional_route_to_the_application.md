
### What?
CF creates a route for your app when it is first pushed, but you can create additional routes using the cf CLI.

### How?
**[Create and map a new route for dora](https://docs.cloudfoundry.org/devguide/deploy-apps/routes-domains.html#map-route)**

### Expected Result
1. Run `cf routes` to verify your route has been created and correctly mapped to dora
1. Visit route and see dora content

### Additional notes
Combinations of route and domain names must be unique. Even though you may not see a particular route listed in your space, that doesn't mean it's available—it may be taken by another application in a different org or space.

If you had created the new route using `cf create-route` without mapping it to the dora app, it would have existed un-mapped and the router would have returned 404 Not Found. While at first glance this might seem like a useless feature, it allows you to claim a route before you're ready to use it.

### Resources
[CF Routes and Domains](https://docs.cloudfoundry.org/devguide/deploy-apps/routes-domains.html)
[Basic Linux Networking reference](http://www.penguintutor.com/linux/basic-network-reference)