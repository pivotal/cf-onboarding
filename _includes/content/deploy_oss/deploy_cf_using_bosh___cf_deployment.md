
### What?
[Cf-deployment](https://github.com/cloudfoundry/cf-deployment/blob/master/README.md#purpose) is the supported way to deploy CF in the OSS world. It consists primarily of a deployment manifest, as well as **operations files**, or ops-files, that transform the manifest to create a slightly different deployment.

The standard cf-deployment is fairly hefty, so we're going to scale ours down to a more reasonable size using an [operations file](https://bosh.io/docs/cli-ops-files.html) that reduces the number of VMs we deploy. Using an operations file to make small adjustments like this means we can edit our deployment without changing the core `cf-deployment.yml` manifest.

For TPCF, customers would upload the TAS tile to Operations Manager to install CF

### How?
Clone the **[cf-deployment repo](https://github.com/cloudfoundry/cf-deployment)** to your workstation (we assume you put it in `~/workspace`, next to your onboarding directory).
```
bosh -d cf deploy ../cf-deployment/cf-deployment.yml -v system_domain=<YOUR_DOMAIN> --non-interactive -o ../cf-deployment/operations/use-compiled-releases.yml -o ../cf-deployment/operations/scale-to-one-az.yml
```

This one step is part of a step-by-step procedure listed in [these instructions](https://github.com/cloudfoundry/cf-deployment/blob/master/texts/deployment-guide.md), but this backlog has already set you up with everything you need.

Let's discuss the options we used:
- We added `--non-interactive` flag. Normally, after initial preprocessing steps, BOSH will present the "manifest changes", differences between the currently-deployed manifest (currently nothing) and the deploying manifest. Without this flag, you'll be prompted ~5 minutes after the `bosh deploy` command to approve the manifest changes.
- We added `-o operations/use-compiled-releases.yml`, to skip compiling all of your releases.
- We added `-o operations/scale-to-one-az.yml` to the deploy command. This ops-file will cut your deployment size in half, so you'll deploy faster (and your deployment will cost less money).

### Expected Result
Wait for the BOSH deploy to complete. It will take a while...like 20 minutes or so...you might want to check out the next "intermission" story, or go grab a snack.
Once the deploy is finished, run `bosh vms`. All of them (currently 15) should have a status of "running".

### Resources
[Docs: What is a BOSH release?](https://bosh.io/docs/release.html)
[cf-deployment Github README](https://github.com/cloudfoundry/cf-deployment/blob/master/README.md)
