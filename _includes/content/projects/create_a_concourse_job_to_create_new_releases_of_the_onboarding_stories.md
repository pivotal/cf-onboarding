
### What?
This is a chance to practice your Concourse skills and help out the Onboarding facilitation team.

We set up the Tracker projects by building a CSV that we then upload to Tracker. We also periodically upload both versions of that CSV as a Github release, and publish it along with notes on changes to the repository.

You can see what these releases look like in [the releases section of the Onboarding Github repository](https://github.com/pivotal-cf/onboarding/releases)

Right now a human has to create the release, then build and upload these CSVs by hand, but we'd rather do that with a Concourse job.


### Expected Result

What we'd like to see is a Concourse pipeline that:

- Takes in the latest `master` of `pivotal-cf/onboarding`
- Builds both CSVs (oss and local)
- Creates a draft release with these CSVs that points to the commit they were built from

Bonus points if the pipeline:

- Runs automatically every time we commit to master but doesn't fill the repository with draft releases
- Gives us a button that lets us publish the release as well
- Shows us all the commits that were included in the release, to make it easier to write release notes
- Handles incrementing semver for us, and lets us choose whether we're making a breaking, feature, or fix release

### Resources

You'll need to ask a facilitator for 
- credentials to automatically create releases in the `pivotal-cf/onboarding` channel
- credentials to update pipelines in the Wings Onboarding team

You may also want to consult
- Instructions on [building stories from source](Import stories to Tracker (from source))
- [Concourse resources](https://concourse-ci.org/included-resources.html)