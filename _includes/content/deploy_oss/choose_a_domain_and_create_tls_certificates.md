
### What?
Ok, we're getting closer to deploying. Before you can go prime time with BOSH and CF, you'll need a way to make your platform addressable from the internet -- specifically, by getting a domain name for your Cloud Foundry deployment. In the next story, we're going to use `bbl` to create a load balancer for our soon-to-be Cloud Foundry deployment. This load balancer requires a TLS certificate, which will include your domain in its metadata. So, the order of operations here is:
1. Choose a domain
2. Create a TLS certificate that includes that domain as the "Common Name"
3. Create load balancers and configure them with your TLS certificate

In this story, you'll do steps 1 and 2. The last step will come in the next story.

(If you want some more information on DNS, jump down to the "Side note" below.)

### How?
We're going to use [OpenSSL](https://www.openssl.org/) to generate a self-signed TLS certificate. In real life, you'd get one from a trusted Certificate Authority like [Let's Encrypt](https://letsencrypt.org/), but for this exercise self-signed is sufficient.

So, first, choose a subdomain of `cf-onboarding.com` -- something like `my-env.cf-onboarding.com`, but feel free to be creative and call it whatever you like. From here on out, we're going to call this domain your **system domain**, because this is the domain that the CF system will use for itself, including things like the CF API (AKA Cloud Controller).

Next, we'll make a TLS certificate. To do this, run `openssl req` with a few arguments:

* `-x509` outputs a x509 structure, a standard that defines the format of public key certificates.
* `-newkey rsa:2048` generates a new RSA key of 2048 bits in size.
* `-keyout` and `-out` arguments set output file paths (I generally use `key.pem` and `cert.pem`, respectively)
* The `-nodes` argument sets it to not encrypt private keys, meaning you won't have to enter a PEM passphrase.

**Caution:** when you run OpenSSL, it will ask you to enter values like your Country and State. The only one that really matters (and it matters a great deal) is the "Common Name," which you should fill with your system domain.

### Expected Result
You should have two new `.pem` files, one containing your private key and one containing your cert.

When you run `openssl x509 -noout -subject -in cert.pem` the information returned matches what you entered.

### Resources
[Blog post: Signed vs Self-Signed Certificates](https://www.thoughtco.com/signed-vs-self-signed-certificates-3469534)
[How to: Get Common Name from TLS cert](https://unix.stackexchange.com/questions/103461/get-common-name-cn-from-ssl-certificate)
