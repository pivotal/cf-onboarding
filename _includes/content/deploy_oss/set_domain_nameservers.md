
### What
So far, you've created a TLS certificate, configured your domain, and then used `bbl` to create both Load Balancers (which are configured with that certificate), as well as DNS records for your system domain -- which is itself a subdomain of `cf-onboarding.com`.

If you look back at our earlier description of DNS, you'll recall that you need a series of NS records to guide the resolver to more specific subdomains. The CF R&D Org has an NS record for `cf-app.com`, and the Onboarding Staff has created the NS record for `cf-onboarding.com`. You yourself created a series of A records for `*.my-env.cf-onboarding.com` and such -- all that's left to do is create one last NS record for `my-env.cf-onboarding.com`.

### How
**First, find the nameservers for the DNS records that bbl created for you.**
1. Open the GCP console in your browser.
1. In the hamburger menu (the three-line icon in the top left corner), select `Network services` > `Cloud DNS`.
1. Click on the zone created by `bbl` (it often has `bbl-` prepended to the zone name).
1. The first entry in the table of DNS records should be an NS record for your system domain. In the `Data` column, you should see four values of the form `ns-cloud-b1.googledomains.com.`. Copy that list.

**Next, create an NS record in the Onboarding DNS account to link up your new DNS zone to the exisiting zone for `cf-onboarding.com`.**
1. Next to the Google Cloud Platform logo, you should see the name of a GCP project. Click on the dropdown and select the project called `CF-Greenhouse-Spitfire`.
1. In the hamburger menu (the three-line icon in the top left corner), select `Network services` > `Cloud DNS`.
1. Choose the zone called `cf-onboarding-com`. At this point, you should see NS and SOA records for `cf-onboarding.com.`, and possibly NS records for some subdomains that were never cleaned up (if you see any such records, notice that they've also got four values from `googledomains.com`. We're going to make a new record just like that.)
1. Click `Add record set`
1. Fill out the form. The `DNS Name` should match your system domain, the `Resource Record Type` should be `NS`, and the `Name server` list should have one value for each of the values you copied earlier. (By default, the TTL is set at 5 minutes. The TTL, or Time to Live, indicates how long a DNS client is allowed to cache the DNS lookup result. A longer TTL means you don't need to query the DNS server as much, but if you make a change to your DNS record clients will be slower to pick up the change.)
1. Finish by clicking `Create`

### Expected Result
DNS updates can take a minute to propogate, so go grab a cup of coffee or a snack. When you get back, go to your GCP Cloud DNS entry. Find the IP associated with `*.your-domain.com`. Running `dig api.your-domain.com` should return the same IP address.

