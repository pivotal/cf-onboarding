
### What?
Environment variables are the means by which the Cloud Foundry [runtime](https://www.techopedia.com/definition/5466/runtime-environment-rte) communicates with a deployed application about its environment. You can use them too!

### How?
1. Run `cf set-env dora SOME_VAR some-var-value`
1. `cf env dora`

### Expected Result
The cf CLI will show the application's environment variables classified as System-Provided (i.e. `VCAP_APPLICATION`) and User-Provided (including `SOME_VAR`). Running and Staging variables are provided by operators that apply to all applications.

### Resources
[Docs: Cloud Foundry Environment Variables](https://docs.run.pivotal.io/devguide/deploy-apps/environment-variable.html)
[Docs: What is a Runtime Environment (RTE)?](https://www.techopedia.com/definition/5466/runtime-environment-rte)
[Docs: VCAP_APPLICATION](https://docs.run.pivotal.io/devguide/deploy-apps/environment-variable.html#VCAP-APPLICATION)
[Docs: VCAP_SERVICES](https://docs.run.pivotal.io/devguide/deploy-apps/environment-variable.html#VCAP-SERVICES)
[Wikipedia: What is an environment variable?](https://en.wikipedia.org/wiki/Environment_variable)