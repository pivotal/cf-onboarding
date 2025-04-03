
### What?
With BOSH it is easy to scale deployments. All you need to do is modify the number of instances of that [job](https://bosh.io/docs/jobs.html) in the deployment manifest and redeploy.

What's the best way to achive that? You actually have a few options, but in this story, we're going to focus on using **[ops files](http://bosh.io/docs/cli-ops-files/#ops)**.

Ops files are a way for operators to make edits to a deployment manifest in a reproducible way. Using a special syntax ([here are some examples](https://github.com/cppforlife/go-patch/blob/master/docs/examples.md)), they include a list of mutations to make to any YAML file (although you'll mostly use them for your manifest). You can use ops-files to do the following:
- **Remove** a section of YAML
- **Replace** one section of YAML with your own YAML. Tweaks to a "replace" command can also allow you to:
  - **Add** a new section of YAML
  - **Append** to a list

### How?
1. Get your current bosh manifest by running `bosh -d cf manifest > /tmp/cf.yml`
1. Write an ops-file called `scale-diego-cells.yml` that will scale up the number of Diego cells to 3 instances. See if you can use the ops-file documentation and examples to figure out how. (Some hints: think about which spot in the manifest you want to update. What operation do you need to carry out?)
1. Vet your changes by running `bosh interpolate`:
  ```
  bosh interpolate /tmp/cf.yml \
  -o ... [any other ops-files you used to deploy orginally] ... \
  -o scale-diego-cells.yml \
  -v system_domain=$SYSTEM_DOMAIN
  ```
  Dig around in the output, and see if it matches what you expected. Specifically, is the instance count for the `diego-cell` instance group set to what you wanted?
1. Deploy:
  ```
  bosh deploy /tmp/cf.yml \
  -o ... [any other ops-files you used to deploy originally] ... \
  -o scale-diego-cells.yml \
  -v system_domain=$SYSTEM_DOMAIN
  ```
  (Also, notice the parallels between the `interpolate` and `deploy` commands.)

###  Expected Result
You should easily be able to scale the number of Diego Cells up or down. What happens to your apps at that point? Are they redistributed as soon as there is a new cell or do you have to scale the app to trigger a relocation?
cf-deployment installs [`cfdot`](https://github.com/cloudfoundry/cfdot) on the diego cells, which you can use to interrogate Diego. The `watch` and `tail` bash commands will also be your friends during this investigation.

### Resources
[Docs: BOSH Deploying, a step-by-step walk through](http://bosh.io/docs/deploying-step-by-step.html)
[Docs: BOSH Deployment Manifest docs](http://bosh.io/docs/manifest-v2/)
[Docs: Creating a new BOSH VM](http://bosh.io/docs/bosh-components.html#create-vm)
[YAML Validator](http://codebeautify.org/yaml-validator)
[JSON Patch, the inspriation for ops-files](JSON Patch](http://jsonpatch.com/)
[cf-deployment ops-files](https://github.com/cloudfoundry/cf-deployment/tree/master/operations)