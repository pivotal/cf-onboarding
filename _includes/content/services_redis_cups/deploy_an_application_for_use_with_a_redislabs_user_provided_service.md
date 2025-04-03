
### What?
To play around with the RedisLabs user-provided service instance we just created, we're going to need an app to use it. Let's deploy and bind one now!

### How?
Use this [sample app](https://github.com/pivotal-cf/cf-redis-example-app) and the [instructions](https://github.com/pivotal-cf/cf-redis-example-app#using-redis-labs-cups) provided in the `README` to deploy the sample application and bind your service instance to it.

**Note:** The instructions above actually have a slight error; the command `cf cups redis -p` should accept JSON format, not a comma-separated list. For example, `cf cups redis -p '{"username":"..."}'`. A PR has been created to fix this, so disregard this message and consider contributing back to the onboarding repo to remove this message when that PR is accepted.

**P.S.:** Don't forget to register for an account on https://redislabs.com/ and create a DB!

### Expected Result
You can see the sample app listed under `cf apps`.
You can see the service listed under `cf services`, and see that it is bound to the sample app.

From your terminal, run:
```
export APP=APP_NAME.YOUR_CF_DOMAIN.com
curl -X PUT $APP/foo -d 'data=some-value'
```
This should return `success`. Next run:
```
curl -X GET $APP/foo
```
This should return `some-value`.

The value should also be present in your RedisLabs Dashboard.

### Resources
[User Provided Services (CUPS) Documentation](https://docs.cloudfoundry.org/devguide/services/user-provided.html)
[Bind a Service Instance Documentation](https://docs.cloudfoundry.org/devguide/services/application-binding.html)