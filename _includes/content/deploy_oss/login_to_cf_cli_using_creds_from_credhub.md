
### What?
We're almost there. The final step is to target your newly deployed Cloud Foundry.

### How?
First, point the CF CLI at your deployed cloudfoundry. To do this, run `cf api api.YOUR_DOMAIN.com --skip-ssl-validation`.

Next, before you can login to the CF CLI, you'll need to get your admin password, which is stored in Credhub. When you used `bbl` to create your BOSH setup, it colocated a Credhub server with the BOSH director, so we can use `bbl` to get a handful of environment variables that will help you target the Credhub that was created as part of your BOSH setup.

Make sure you've got your bbl env set up (`eval "$(bbl print-env)"`). If you run `env | grep CREDHUB`, you can look at all of the environment variables that get used to access Credhub.

Next, make sure you've got the credhub CLI installed. You can do that with `brew install cloudfoundry/tap/credhub-cli`.

Next, you can use the credhub CLI to search for the "path" that the variable is stored at:
```
credhub find -n cf_admin_password
```

The result should include a path that looks like `/bosh-bbl-env-baikal-2018-06-04t22-02z/cf/cf_admin_password`. Copy that path, and use the credhub CLI one more time to fetch the value of the credential:
```
credhub get -n /bosh-bbl-env-baikal-2018-06-04t22-02z/cf/cf_admin_password
```

The output should include a field called `value` -- that's your CF admin password.

Finally, provide your credentials via `cf login` (hint: if it is not already apparent, your username is "admin").

### Expected Result
`cf login` works.

### Troubleshooting
* Do the domain names associated with your TLS cert, your load balancer, and your system_domain all match the domain you registered on GCP?
* Are the nameservers associated with your domain the same as those associated with your load balancer? (run `bbl lbs` to check)
* When you `dig` or `ping` your CC api endpoint, is the IP you hit what you expect it to be?
* If you were redeploying, did you run the `bosh -n interpolate` step before running `bosh -d cf deploy`?

### Resources
[Forum question: Get Common Name from TLS cert](https://unix.stackexchange.com/questions/103461/get-common-name-cn-from-ssl-certificate)
