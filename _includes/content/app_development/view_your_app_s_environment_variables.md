
### What?
Environment variables are the means by which the Cloud Foundry [runtime](https://www.techopedia.com/definition/5466/runtime-environment-rte) communicates with a deployed application about its environment. You can use them too!

### How?
1. `cf set-env dora MY_VAR hello`
1. `cf env dora`

### Expected Result
The cf CLI will show the application's environment variables classified as System-Provided (i.e. `VCAP_APPLICATION`) and User Provided (`MY_VAR`). Running and Staging variables are provided by operators that apply to all applications.

### Resources
[Cloud Foundry Environment Variables](https://docs.cloudfoundry.org/devguide/deploy-apps/environment-variable.html)
[What is a Runtime Environment (RTE)?](https://docs.cloudfoundry.org/definition/5466/runtime-environment-rte)
[VCAP_APPLICATION](https://docs.cloudfoundry.org/environment-variable.html#VCAP-APPLICATION)
[VCAP_SERVICES](https://docs.cloudfoundry.org/devguide/deploy-apps/environment-variable.html#VCAP-SERVICES)
