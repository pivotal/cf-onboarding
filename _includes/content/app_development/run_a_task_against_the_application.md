
### What?

When we run `cf push MY_APP`, we're creating a long running process or LRP in
Diego terminology.
However, sometimes you want to run a short-lived script instead like running
`rake db:migrate` after your Ruby process starts up.
A task is an application or script whose code is included as part of a deployed application, but runs independently in its own container.
These tasks execute, print their STDOUT and STDERR to the parent app's log, and
exit with a zero status code for success or non-zero for failure.

### Run a simple task
1. Push the Dora app as described in the "Push a sample app" story
1. In one terminal run `cf logs dora`
1. In another terminal run `cf run-task dora 'sleep 15 && echo "Hello task!"'`
1. Run `cf tasks dora`

### Expected Result
1. You should see "Hello task!" appear in the app logs after ~15 seconds
1. Tasks are asynchronous so the `cf run-task` command returns immediately without
   waiting for the task to finish running. `cf tasks` shows you which tasks are
   currently running.

### Run a failing task
1. Run `cf run-task dora 'cat missing-file.txt'`
1. Run `cf tasks dora`

### Expected Result
1. You should see that the task is soon marked as "FAILED" as running `cat`
   against an invalid filepath will exit non-zero.

### Understanding containers
1. Continue tailing app logs in one terminal
1. Run `cf run-task dora 'touch NEW_FILE.txt && ls'`
1. Run `cf tasks dora` and wait for task to finish
1. Run `cf run-task dora 'ls'`

### Expected Result
1. The `ls` output in the first task should show `NEW_FILE.txt` listed along
   with some other files from the dora app, but the second `ls` will not show
   `NEW_FILE.txt`. Why would that be?
   - A: Each task is run in a new container which contain the contents of the parent
     app's droplet. This is why you see the app files like `Gemfile`. However,
     any new files or modifications you make to the file system while in the
     container are discarded when the task exits. Each new task gets a fresh
     copy of the filesystem containing the droplet files and any writes are only
     viewable from within the current container.

### Understanding container placement
1. Continue tailing app logs in one terminal
1. Run `cf run-task dora 'echo MY INSTANCE IP: $CF_INSTANCE_INTERNAL_IP'` (single quotes
   are important)
1. Run `cf run-task dora 'echo MY INSTANCE IP: $CF_INSTANCE_INTERNAL_IP'` a couple more
   times

### Expected Result
1. If you have more than one Diego cell VM in your deployment (which you
   will not if you used the 'scale-to-one-az.yml' opsfile), you should see
   different IP addresses printed in the logs each time you run the above task.
   This shows that Diego will schedule task containers on whatever cell happens
   to have capacity at that moment. The tasks do not need to be scheduled on the
   same cell as the parent app as tasks are created in a new container
   independent of the parent app container. If you need access to the parent app
   container use `cf ssh` instead.

### Resources
[Running Tasks](https://docs.cloudfoundry.org/devguide/using-tasks.html)
[Tasks API Docs](http://v3-apidocs.cloudfoundry.org/version/3.51.0/index.html#tasks)
