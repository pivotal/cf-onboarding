
### What?
Each BOSH job can specify customizable properties in a **[spec file](http://bosh.io/docs/jobs.html#spec)**. Some will be required, others will have defaults. We're going to fiddle with the values and see what happens.

### How?
1. Run `bosh releases`
1. Select a job, your pick. Ooh, roulette. Let's get risky. Or not. This is a good time to revisit the [diego-design-notes](https://github.com/cloudfoundry/diego-design-notes) if you'd like to be incisive about this.
1. Find the spec file for the job on Github. A spec file resides at `some-release/jobs/some-job-name/spec`. For instance, the spec path of the cf-release uaa job is `https://github.com/cloudfoundry-attic/cf-release/blob/master/jobs/uaa` which points to `https://github.com/cloudfoundry/uaa-release/blob/master/jobs/uaa/spec`.
1. Choose a spec property where you expect to be able to tell if the value changes. An example in uaa might be `login.logout.redirect.url` (read the property description in the spec to learn more). A more straight-forward option might be a port. Whatever it is, establish how you plan to test the change before implementing it.
1. Update your Cloud Foundry manifest to reflect a change in that job's property. This will be under jobs: some-job-name: properties: some-property. If you don't see the property you're looking for it's because it has been left as its default value.
1. `bosh deploy` !!

### Expected Result
A few things could happen at this point. The deployment could run successfully (woot!) and you are able to demonstrate that the change took effect (double woot). The update could fail, but now you have something to diagnose (hint: it might have something to do with the user permissions of the user responsible for making changes to that job). Orrrrr you could break something! That one is a super exciting option, chalk full of investigative opportunities.

### Resources
[Docs: What is a BOSH spec file?](http://bosh.io/docs/jobs.html#spec)
[Tool: Chaos Lemur](https://github.com/strepsirrhini-army/chaos-lemur)
[YAML Validator](http://codebeautify.org/yaml-validator)
BOSH's Wikipedia entry is also good, but it has parentheses in the URL and that's well past the limits of my markdown skills at this hour of night