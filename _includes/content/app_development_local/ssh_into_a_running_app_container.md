
### What?
Your applications run on Diego Cells, which create containers for each of your application instances. One of the features of [Diego (the new CF Runtime)](http://www.starkandwayne.com/blog/demystifying-cloud-foundrys-diego/) versus DEAs (the original Runtime) is that you have the ability to [SSH](https://en.wikipedia.org/wiki/Secure_Shell) into the containers for debugging purposes.

### How?
1. Run `cf ssh dora`to begin a secure connection to the SSH server for your app.
1. Run `ls -l` to view files and `cd SOME-DIRECTORY-PATH` to move around the directory tree.
1. Use the `ps aux` command to see what processes are running in the app container.
1. Run `env` to view environment variables. You'll see there are additional environment variables that the cf CLI doesn't print, but that are available to your application. The full [list of variables can be viewed here](https://docs.cloudfoundry.org/devguide/deploy-apps/environment-variable.html).
1. Run `exit`, then SSH back into a specific instance of your app, using the `-i SOME-APP-INDEX` flag (when you run `cf app dora` you'll see the indexes of the instances displayed in turquoise to the left of their information). View the [cf CLI SSH flags](http://cli.cloudfoundry.org/en-US/cf/ssh.html) to see what else is possible.

**TRIVIA:** did you notice that you're SSH-ed in as the `vcap` user? VCAP stands for "VMware's Cloud Application Platform," an old name for Cloud Foundry. Catchy, right?

### Resources
[Docs: Differences Between DEA and Diego Architectures](https://docs.cloudfoundry.org/concepts/diego/dea-vs-diego.html)
[Blog post: Demystifying Cloud Foundry's Diego](http://www.starkandwayne.com/blog/demystifying-cloud-foundrys-diego/)
[Unix/Linux StackExchange: What does aux mean in `ps aux`?](http://unix.stackexchange.com/questions/106847/what-does-aux-mean-in-ps-aux)