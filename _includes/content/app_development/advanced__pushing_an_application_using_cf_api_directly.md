
### What?
The `cf push` command suffices common use cases of running an app on Cloud Foundry, with a lot of tasks automated and running under the hood. But for more complicated workflows, as an advanced user, you might want to make requests to the Cloud Controller(CC) API directly. Are you up for the challenge?

### How?
1. If dora is still running, `cf delete dora` to start fresh.
1. ** Follow instructions on [How to Create an App Using V3 of the CC API](https://github.com/cloudfoundry/cloud_controller_ng/wiki/How-to-Create-an-App-Using-V3-of-the-CC-API)**

### Expected Result
1. `cf app dora` will print the app's status, including the URL. When your visit the URL in your browser you should see the words "Hi, I'm Dora!"

### Additional notes
Using the API directly lets users do as much or as little of `cf push` as they like. Some example workflows:
- Rolling back to a previous droplet
- Running a task without having to run the app first

### Resources
[V3 API docs](http://v3-apidocs.cloudfoundry.org)
[V2 API docs](https://v2-apidocs.cloudfoundry.org/)
[Component: Cloud Controller](https://docs.cloudfoundry.org/concepts/architecture/cloud-controller.html)
[Using Experimental of CLI Commands](https://docs.cloudfoundry.org/devguide/v3-commands.html)
