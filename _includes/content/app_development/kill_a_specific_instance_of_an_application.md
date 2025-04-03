
### What?
Sometimes it's clear that a particular instance of your app is having a problem, but it isn't being identified and culled by the [HealthCheck](https://docs.cloudfoundry.org/devguide/deploy-apps/healthchecks.html). This is a good way to take it out yourself.

### How?
Run `watch cf app dora`.

In another buffer, restart a specific instance of dora using `cf restart-app-instance APP_NAME INSTANCE_INDEX`.

Alternatively, you can also kill a specific instance of dora using `cf curl /v2/apps/APP_GUID/instances/0 -d '' -X DELETE` (reading the `cf curl` help content and [API Docs: Terminate the running app instance at the given index](http://v2-apidocs.cloudfoundry.org/apps/terminate_the_running_app_instance_at_the_given_index.html) may clarify some things).

### Expected Result
You should see the correct instance of your app dying and recovering.

### Explanation
You might notice that `cf restart-app-instance` is way more straightforward. The alternative way of terminating an app instance using `cf curl` is part of Cloud Controller API v2.

### Resources
[Cloud Foundry CLI Reference Guide: restart-app-instance](http://cli.cloudfoundry.org/en-US/cf/restart-app-instance.html)
