
### What?

It can be unclear what is happening underneath bbl. Taking a look at Terraform can clear up some of the mystery.

### How?

You can use [this](https://learn.hashicorp.com/terraform/getting-started/build) to make a simple compute instance.
However, it is specific to AWS, so if you are using GCP, you will have to translate it.
[This](https://www.terraform.io/docs/providers/google/r/compute_instance.html) will be helpful when translating.

### Resources

For more info on how bbl and terraform work together, check out these docs and the rest of the repo:
https://github.com/cloudfoundry/bosh-bootloader/blob/master/docs/index.md