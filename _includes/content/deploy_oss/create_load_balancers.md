
### What
We've been alluding to load balancers for a while, but now we'll actually get around to creating them with bbl.

A load balancer is some machine -- or set of machines -- that serve as the entry point to the system and delegate traffic to other components. Importantly, one of their jobs is to make sure that traffic is distributed evenly among VMs that do similar work. The easiest load balancer to understand is the HTTP load balancer, which balances traffic to the [CF routers](http://docs.cloudfoundry.org/concepts/architecture/router.html).

Most public IaaSes provide load balancing solutions, but some do not. For example, if you were deploying to vSphere, you'd need to come up with your own load balancing solution: most people use [HAProxy](https://github.com/cloudfoundry-incubator/haproxy-boshrelease) or hardware load balancers for this.

In this story, you'll update bbl's "plan" with the `--lb-type` flag. When set to `cf`, this flag will cause bbl to make a few different load balancers for you:
- HTTP and Websocket load balancers that delegate to the HTTP router, commonly known as GoRouter
- SSH load balancer the delegates to the SSH proxy
- TCP load balancer that delegates to the TCP routers
- Credhub load balancer that delegates directly to Credhub

The other thing you'll do in this command is configure the load balancers with the TLS certificate you made in the previous story.

Finally, if you provide the `--lb-domain` flag, bbl will also create DNS records for each of the domains associated with your load balancers, like A records for `*.my-env.cf-onboarding.com` and `ssh.my-env.cf-onboarding.com`

### How
To set up load balancers with bbl, we're going to leverage bbl's "plan" feature. Remember how you saw all those directories and files in your working directory after running `bbl up`? Under the hood, bbl first laid out a "plan" of all the stuff it was going to do, and then it did them. We're going to ask bbl to update its plan to include load balancers, and then we're going to run `bbl up` again to effect those changes on GCP.

Run
```
bbl plan \
  --lb-type cf \
  --lb-cert PATH_TO_TLS_CERT --lb-key PATH_TO_TLS_KEY \
  --lb-domain YOUR_DOMAIN \
  --gcp-service-account-key service-account.key.json
```
You should be able to see that a handful of files in your bbl state directory have been updates -- specifically, your terraform and cloud-configs directory.

Next, run
```
bbl up --gcp-service-account-key service-account.key.json
```

### Expected Results
Network Services > Load balancing to see the load balancer that `bbl` generated on GCP.

Network Services > Cloud DNS to see the [Cloud DNS Zone](https://cloud.google.com/dns/overview) and [record sets](https://cloud.google.com/dns/overview#supported_dns_record_types) that `bbl` generated on GCP.

### Resources
[Docs: Types of load balancing](https://cloud.google.com/compute/docs/load-balancing/#types_of_load_balancing)
[Docs: HAProxy](http://www.haproxy.org/)"
