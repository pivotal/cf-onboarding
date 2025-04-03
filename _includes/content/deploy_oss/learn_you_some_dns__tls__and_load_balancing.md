
I'll admit, this story is a little goofy, since there isn't a ton to do. Instead, we're just going to introduce a handful of interrelaled concepts, so that you'll feel a bit more informed as you go through the process of setting up everything needed to make your Cloud Foundry addressable from the public internet.

Let's start of with some vocabulary:

* An **IP address** is number associated with a computer, and allows other computers to find them. We typically think of two kinds of IP addresses: **public** and **private** IPs. Public IPs are unique, and anybody on the internet can try to communicate with the computer associated to a public IP. Private IPs are not accessible by just any computer -- you have to be on the same private network to communicate with a computer on a private IP. In most Cloud Foundry deployments, the majority of VM instances use private IPs.
* **Load balancers** are computers with the sole job of distributing requests across a set of other computers (sometimes referred to as "backends" or "backend servers"). In CF deployments, we typically create load balancers with public IPs, and then configure them to route traffic to components that have private IPs.
* A **domain** is a series of strings separated by the `.` character that serves as an alias for an IP address, or set of IP addresses. For example, `tanzu.vmware.com` is a domain that maps to the IP address(s) of the load balancers on PWS. We use domains for a few reasons. One is that IP addresses are hard for humans to remember, but `google.com` is pretty easy. We also use domains to group computers together and even do some load balancing. For example, `google.com` does not resolve to a single IP address (one poor computer could never handle that kind of load!).
  * A **subdomain** is a term you'll see as well. It's a relative term that describes the relationship of one domain to another -- for example, `my-site.example.com` is a subdomain of `example.com`.
* **DNS** stands for "Domain Name System," which translates domains IP addresses. We'll explain how it works in a second.
* **TLS** stands for "Transport Layer Security". You'll often also hear it referred to as **SSL**, which stands for "Secure Socket Layer," even though the term "SSL" is deprecated. Anywho, TLS is a system for encrypting your communications and validating that you're communicating with the right computer on the internet. To prove its identity, a computer will present a **certificate** to the machine making the request.

Why introduce all this now? Because you'll work on these bit-by-bit over the next few stories. Here's how.

Most components in CF are deployed to private IPs. As such, for a user to access the platform, or any app deployed to it, we need some way of getting traffic into the private network for CF. For that, we typically set up load balancers; as mentioned before, load balancers are mostly deployed to public IPs, so they're accessible from the internet. But how does a user know where to find the load balancer? We could just share the public IP around, but that's hard for people to remember, and looks janky on a marketing webpage. So, we create a domain for our system and have it (and some of its subdomains) resolve to our load balancers.

Ok, so we've got domains that resolve to load balancers, which are deployed to public IPs and route traffic to our private IPs. But, how does a computer verify the identity of the load balancers once it sends them some traffic using its domain? The load balancer presents all clients with its TLS certificate, which essentially states "I am in fact `example.com`, and I have the certificate to prove it." This is important: the certificate includes the domain in its "metadata."

### Side note: How does DNS work exactly?
As part of these stories, you'll pick a subdomain of `cf-onboarding.com` -- like `my-env.cf-onboarding.com` -- and use it for your environment. That means that if anyone tries to make an HTTP request to `my-env.cf-onboarding.com`, a process on their computer, called a **DNS resolver**, will make a series of queries to computers on the internet, called **DNS servers**, essentially asking, "Do you know this guy, `my-env.cf-onboarding.com`?"

The simplest way to manage such a request is to create an **A record** (Address record).
These records map a domain, like `onboarding.example.com`, to a specific IP address like `1.2.3.4`.
In this setup, when a user runs `curl onboarding.example.com`, their DNS resolver first queries its nearest DNS server for the `.com` TLD (Top Level Domain).
The DNS server redirects your computer to _another_ DNS server that knows about `.com` domains.
Then your computer asks that server about `example.com`.
The DNS server redirects your computer to _yet another_ DNS server that knows about the `example.com` domain.
Then your computer asks that server about the `onboarding.example.com` domain.
That name server has the A record, and returns an answer: `1.2.3.4`

The following table is an example request and response flow to resolve `onboarding.example.com` to the A record `1.2.3.4` by your computer:

  | # | DNS Request  | DNS Response  |
  |---|---|---|
  | 1 | `.com`  to local DNS | `NS Record 1.1.1.1`  |
  | 2 | `example.com` to `1.1.1.1`  | `NS Record 2.2.2.2`  |
  | 3 | `onboarding.example.com` to `2.2.2.2`  | `A Record 1.2.3.4`   |

However, Cloud Foundry requires a slightly more complex setup, because we want to resolve all queries that _end_ with `onboarding.example.com` -- for example, `api.onboarding.example.com` or `my-app.onboarding.example.com`. We also often want to reserve certain subdomains like `bosh.onboarding.example.com` or `ssh.onboarding.example.com`.

To achieve that, we first set up a different kind of record, an **NS record** (Nameserver record), which maps a domain to other DNS servers.
It essentially says, "If you're looking for information about domains that end with `onboarding.example.com`, go ask this other set of DNS servers." Your DNS resolver will then query those servers instead. This is a useful way to hook up different DNS providers or piggy-pack on existing domains. You're going to leverage this functionality in order to register your subdomain of `cf-onboarding.com` on GCP.

In a few stories, you'll use bbl to create load balancers that have IP addresses associated with them. The same  bbl command will _also_ create a few DNS records:
- An NS record that defines the Google DNS servers. You'll take that list of servers to our "master" DNS account on GCP to configure it to defer to your DNS records.
- An A record for a few reserved subdomains like `ssh.my-env.cf-onboarding.com`. This DNS record maps the `ssh.` subdomain to a load balancer for SSH connections to CF Apps (you'll get to explore this feature more later).
- An A record for `*.my-env.cf-onboarding.com`. This record maps **all** other subdomains of `my-env.cf-onboarding.com` to a load balancer that delegates traffic to the router. This **wildcard domain** enables apps that get pushed to the platform to piggyback on the `my-env.cf-onboarding.com` domain, so that you can make subdomains on the fly like `my-app.my-env.cf-onboarding.com`.

### Resources
[Blog post: What is a domain name?](https://www.lifewire.com/what-is-a-domain-name-2483189)
[Blog post: SSL vs TLS, what's the difference?](https://luxsci.com/blog/ssl-versus-tls-whats-the-difference.html)
[Tutorial: How SSL and TLS encryption works](http://computer.howstuffworks.com/encryption4.htm)
[Blog post: What is a Domain Name Server?](https://www.lifewire.com/what-is-a-dns-server-817513)
[Video: How DNS works](https://www.youtube.com/watch?v=GlZC4Jwf3xQ)
[Video: A Tour of TLS](https://www.youtube.com/watch?v=yzz3bcnWf7M)"
