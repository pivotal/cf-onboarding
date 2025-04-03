
### What?
The **[stemcell](http://bosh.cloudfoundry.org/docs/stemcell.html)** is the foundation of every VM the [BOSH Director](https://bosh.io/docs/bosh-components.html#director) deploys. It is a versioned Operating System [image](https://docs.docker.com/engine/userguide/storagedriver/imagesandcontainers/#images-and-layers) wrapped with [IaaS](https://en.wikipedia.org/wiki/Cloud_computing#Infrastructure_as_a_service_.28IaaS.29)-specific packaging.

A typical stemcell contains a bare bones OS skeleton with a few common utilities pre-installed, some configuration files to securely configure the OS by default, and a BOSH Agent. The [BOSH Agent](https://bosh.io/docs/bosh-components.html#agent) is there to listen for instructions from the Director and to carry them out.

Stemcells are distributed as [tarballs](https://bosh.io/docs/build-stemcell.html#tarball-structure). You need to upload a stemcell for your BOSH Director to use when it deploys your Cloud Foundry VMs.

### How?

Available stemcells **[are listed at bosh.io](http://bosh.io/stemcells)**. Identify the appropriate stemcell for the IaaS  ([Google KVM](https://en.wikipedia.org/wiki/Google_Compute_Engine#Machine_Types)) and OS ([Ubuntu Jammy](https://en.wikipedia.org/wiki/Ubuntu_version_history#Ubuntu_22.04_LTS_(Jammy_Jellyfish))) you'll be using to deploy. You'll also need to pick a specific version of the stemcell; to find the version number, you'll need to look at [cf-deployment.yml](https://github.com/cloudfoundry/cf-deployment/blob/master/cf-deployment.yml) (you'll learn more about this file in the next story), and check in the `stemcells` section. You'll also notice that some stemcells on bosh.io have both a "Light" and "Full" version. A full stemcell includes the entire OS image while a light stemcell contains an ID or URI pointing to a public OS image which has already been uploaded on that IaaS. For public clouds like GCP you'll almost always use the light version.

Now that we've selected which stemcell to upload, we can upload the stemcell tarball to our BOSH Director:

```
bosh upload-stemcell --sha1 $STEMCELL_SHA1 \
  "https://bosh.io/d/stemcells/bosh-google-kvm-ubuntu-jammy-go_agent?v=$STEMCELL_VERSION
```

Note: the SHA1 value should be listed next to the stemcell download link on bosh.io.

### Expected Result

If all went according to plan, the stemcell should be visible via:

```
bosh stemcells
```
**Pro-tip:** since stemcells are just tarballs you can crack them open to take a look. Opening packages.txt will show which packages come with bosh deployed VMS. These packages are sometimes used by BOSH releases

### Resources
[Docs: What is a stemcell?](http://bosh.cloudfoundry.org/docs/stemcell.html)
[Docs: What are light stemcells?](https://bosh.io/docs/build-stemcell.html#light-stemcells)
[Docs: Tarball structure](https://bosh.io/docs/build-stemcell.html#tarball-structure)

### Relevant Repos and Teams
**BOSH:** [cloudfoundry/bosh](https://github.com/cloudfoundry/bosh)
**BOSH:** [cloudfoundry/bosh-agent](https://github.com/cloudfoundry/bosh-agent)