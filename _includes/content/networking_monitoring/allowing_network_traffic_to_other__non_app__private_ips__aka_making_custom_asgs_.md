### What?
In the previous story, you used container networking policies to allow traffic between app containers. Cool beans.
But, what if you want your app to be able to establish connections with other private IP addresses that aren't apps?

If you need help thinking of a reason you'd want to do this, imagine trying to connect your app to anything that was deployed by BOSH, which mostly deploys to private IPs (which, as we've pointed out in previous stories, are blocked from app containers by default).
The most common example of this are BOSH-deployed services like [cf-mysql](https://github.com/cloudfoundry/cf-mysql-deployment).
However, rather than force you to set up another BOSH deployment, let's instead use our existing CF deployment, and see if we can't think of some CF component as a "service" that an app may want to leverage.
For the sake of this story, let's try to query the router for its list of routes (it's a thing you can do!). We'll create our own ASG to allow our app containers to open a connection to the router's private IP.

### Prerequisites

Go through the **Basic BOSH Knowledge** story to gain familiarity with BOSH, specifically BOSH CLI commands.

### How?
The basic outline of this workflow is this:
1. Curl the router's `/routes` endpoint from an app container, and see that the connection isn't allowed.
1. Create and bind an ASG that allows your app container to connect to the IP address of the router.
1. Try curling the `/routes` endpoint again, and see that you get successful response from the router.

Construct the URL for querying the router
In order to curl the router's `/routes` endpoint, you'll to get the router's IP address, port, as well as special credentials (username and password) for making this request. Don't worry, you can find all of this in your manifest or vars-store/Credhub.
- Router IP Address: Run `bosh instances` and look for an instance group called `router` and copy the IP address.
- Router Status Port: routing-release [sets port 8080 as the default](https://github.com/cloudfoundry/routing-release/blob/b63e07857d90df0dc58339416bb09e860e32cf46/jobs/gorouter/spec#L38-L40) -- unless you wrote an ops-file to override the default, you should use this port.
- Router Status Username: Again, [routing-release has provided a default](https://github.com/cloudfoundry/routing-release/blob/b63e07857d90df0dc58339416bb09e860e32cf46/jobs/gorouter/spec#L41-L43). Go with that, most likely.
- Router Status Password: cf-deployment uses [a BOSH variable called `router_status_password` to configure this password](https://github.com/cloudfoundry/cf-deployment/blob/2118b5d8ac09feed153f6eea552fd561dd6aa999/cf-deployment.yml#L865). The way to fetch this value depends on whether you used Credhub or a vars-store to maintain your credentials. 
  For Credhub, make sure you've run `eval "$(bbl print-env)"` (don't forget the double quotes!). Then run `credhub find -n router_status_password` to discover the full name of the variable (it will look something like `/bosh-bbl-env-something/cf/router_status_password`). Copy the full name, and run `credhub get -n $FULL_NAME` to get find the password.
  For vars-store, just grep `router_status_password` from your vars-store

Before creating the ASG:
1. Make sure you have an app pushed. If don't have one running, `cf push dora -p <PATH_TO_DORA>`.
1. `cf ssh` to dora's application container, and then run `curl router-status:$ROUTER_PASSWORD@$ROUTER_IP:8080/routes`.

Create and bind an ASG:
1. Write a file with contents like this:
  ```
  [
    {
      "protocol": "tcp",
      "destination": "$ROUTER_IP/32",
      "ports": "8080",
      "description": "Allow apps to query router-status"
    }
  ]
  ```
1. On your command line, run `cf create-security-group router-asg /PATH/TO/ASG_FILE.json`
1. Bind the ASG to the space that includes your pushed app: `cf bind-security-group router-asg $MY_ORG $MY_SPACE`
1. Restart dora to allow the ASG to take effect: `cf restart dora`

See if the app container can communicate with the router:
1. SSH back onto dora, and re-run `curl router-status:$ROUTER_PASSWORD@$ROUTER_IP:8080/routes`

### Expected Result
The initial curl should fail. After applying the security group, the curl should succeed.

Poke around to see what else you can and can't do with the ASG commands.
Can you allow access to mutliple ports?
What happens if you allow access to a port that the router doesn't listen to? Does it give the same error?


---
Interested in learning more about networking?

🤓 If you are interested in learning even more about networking, check out these [CF Networking Onboarding stories](https://github.com/pivotal/cf-networking-program-onboarding).

🐞 These stories are intended to help engineers to learn how container networking and routing components work under-the-hood and how to debug them.

➕ To add these stories to this tracker: clone the [repo](https://github.com/pivotal/cf-networking-program-onboarding), run the build script, and then import the created CSV to your tracker.

🙋 Also, the networking program is experimenting with running CF _Networking_ Onboarding weeks, much like the CF Onboarding week that you are apart of right now.
So instead of doing these extra stories now (there are a lot of them), you could sign up for a Networking Onboarding week. Reach out to Amelia Downs if this interests you.

📅 Or if you don't have yet another week free to learn, you could complete the stories in your flex time.