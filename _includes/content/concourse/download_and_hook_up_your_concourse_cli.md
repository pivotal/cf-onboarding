
### What?
Now that you have Concourse set up, the first thing you'll want to do is download the `fly` CLI and authenticate with your target. This is done with the `fly login` command. The login command serves double duty: it authenticates with a given endpoint  and saves it under a more convenient name. The name and token are stored in `~/.flyrc` (though you shouldn't really edit the file manually).

### How?
1. Download the `fly` CLI by clicking the button in the **cli:** section on the bottom bar that corresponds to your OS or by running `brew cask install fly` on MacOS (if you have homebrew installed).
1. Move it to your computer's $PATH and make it executable.
1. Run `fly login --help` for instructions to login to Concourse and save a new target (**hint:** you should have seen the URL you'll need for this in a previous story). Our concourse is configured to allow users to login without any credentials (take a look at the `docker-compose.yml` and you should see how that was achieved).

### Expected Result
The CLI prints "target saved" and you can run `fly -t ...` with other commands.

**Note:** If your system prevents `fly` from running, try going to System Preferences -> Security and Privacy and see if you can figure out how to appease Mac and allow `fly`.

### Resources
[Fly login documentation](https://concourse-ci.org/fly.html#fly-login)

### Relevant Repos and Teams
[concourse/fly](https://github.com/concourse/fly)