
### What?
As we've seen, a [Concourse resource](https://concourse-ci.org/resources.html) can be checked for new versions, used as a trigger when new versions are found, pulled down at a specific version, and pushed up idempotently to create new versions.

There's a whole host of [system resource types](https://concourse-ci.org/resource-types.html) available, but they don't cover every potential use case. It won't be too long till you'll need to create one of your very own.

### How?
Use the **[Concourse documentation](https://concourse-ci.org/implementing-resource-types.html)** to create your own resource. Make liberal use of resource repos on Github to guide you on your path. Your steps will broadly be:

1. Write check / in / out scripts that make sense for your resource.
1. Package in a Docker container and publish to the docker registry.
1. Configure your pipeline to include your new resource type, create an instance of that resource type, and consume that resource.
1. Lots of debugging, with poor feedback mechanisms.

**Debugging Note:** write errors to `stderr` because `stdout` must be only a JSON object.

### Expected Result
A resource that can be used as a trigger for a job in your pipeline.

Some ideas:
* weather alerts (https://openweathermap.org/api)
* stock watch

To date, my favorite combinations of resources/tasks have been:
* an XKCD notifier
* a bot that spammed a Slack channel with vocabulary from a Danish thesaurus

Let's get creative!

### Resources
[Docs / Github : List of resource types and links to their implementation](https://concourse-ci.org/resource-types.html)