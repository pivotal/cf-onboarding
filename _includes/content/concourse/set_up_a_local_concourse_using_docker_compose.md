
### What?
[Concourse](https://concourse-ci.org/) is a VMware-sponsored, pipeline-based, continuous integration and deployment (CI/CD) system. The "pipelines" are a collection of three core concepts: [jobs](https://concourse-ci.org/jobs.html), [tasks](https://concourse-ci.org/tasks.html), and [resources](https://concourse-ci.org/resources.html). You'll learn more about these in upcoming stories. While CI may call to mind test automation, Tanzu teams use it for so much more than that. Take a stroll around the office and check out the jobs up on the CI screens to get a general idea of how broadly we use it to automate all that is automate-able.

In this story, we're going to use a tool called [Docker Compose](https://docs.docker.com/compose/overview/) to set up Concourse in local docker containers. (Most engineering teams at tANU use BOSH and [concourse-bosh-deployment](https://github.com/concourse/concourse-bosh-deployment) to set up their Concourse for CI. However, it's a much longer process and we'd rather jump straight to some cool Concourse concepts.)

### How?
Our friends at Stark & Wayne actually have a pretty good Concourse tutorial. We'll start by using their documentation for setting up your local Concourse.

Open https://concoursetutorial.com/#getting-started in your browser and start there. Just in case you wanna know the outline, here's what youre going to do:
1. Install Docker and Docker Compose:
  ```
  brew cask install docker
  brew install docker-compose
  ```
2. Create a dedicated directory for storing all of the files you will need to download during your time here with us. For example:
  ```
  mkdir my-onboarding-stuff
  cd my-onboarding-stuff
  ```
3. Download the Stark & Wayne's pre-packaged Docker Compose YAML configuration:
  ```
  wget https://raw.githubusercontent.com/starkandwayne/concourse-tutorial/master/docker-compose.yml
  ```
4. Use Docker Compose to start up your local Concourse in docker containers:
  ```
  docker-compose up -d
  ```

**Note:** If `docker-compose` fails to run, make sure you started `docker` and that you typed `docker-compose` and not `docker compose`.

**Open Concourse in your browser**
Now that you've got everything set up, navigate to `http://127.0.0.1:8080` in your browser.

### Expected Result
![Lonely Concourse, no pipelines](https://concoursetutorial.com/images/dashboard-no-pipelines.png)

### Resources
[Concourse Architecture Overview](https://concourse-ci.org/concepts.html)
[All About Concourse for Continuous Integration (video)](https://tanzu.vmware.com/content/blog/all-about-concourse-for-continuous-integration)
[YAML Validator](http://codebeautify.org/yaml-validator)

### Relevant Repos and Teams
**Concourse:** [concourse/concourse](https://github.com/concourse/concourse)