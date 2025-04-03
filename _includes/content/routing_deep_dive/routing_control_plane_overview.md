
### What?
In previous stories you pushed apps and created routes. In this set of stories you are going to learn the "behind-the-scenes" on how routes are propagated.

Here is the entire overview. Each step marked with a ✨will be explained in more detail in a story in this track.

**When an app dev maps a route**
1. The app dev maps a route to an app using CAPI.
1. CAPI sends this information to Diego BBS.
1. The Route Emitter on each Diego Cell polls Diego BBS for updated information about the apps on its Cell.
1. ✨ The Route Emitter continually sends the route registration information it gets from the Diego BBS and the IP and port of the app to the GoRouter via NATS.
1. ✨ The GoRouter keeps a mapping of routes -> ip:ports in a routes table, which is consulted when someone curls the route.

### How?
The following stories will look at how many components (CAPI, Diego BBS, Route Emitter, Nats, GoRouter) work together to make routes work.

### Expected Result
You can talk about route propagation at a high level.

### Resources
[CF Routes and Domains](https://docs.cloudfoundry.org/devguide/deploy-apps/routes-domains.html)

