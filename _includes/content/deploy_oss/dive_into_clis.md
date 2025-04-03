
### What?
A CLI , or **Command-line interface** , is a way to issue commands via the terminal. Many Cloud Foundry products don't have any GUIs (Graphical User Interfaces, a website for example), instead the way users interact with our products is via many (many) CLIs.  You've been copying and pasting a lot of CLI commands as part of this onboarding, but for those of us who haven't used CLIs before, let us dig in a little more.

Not all CLIs are the same, but after you start using them you will find a lot of commonality.

Here are some of the Cloud Foundry CLI's you will be using in this onboarding:
- [gcloud](https://cloud.google.com/pubsub/docs/quickstart-cli)
- [bbl](https://github.com/cloudfoundry/bosh-bootloader)
- [credhub](https://github.com/cloudfoundry-incubator/credhub-cli)
- [bosh](https://bosh.io/docs/cli-v2/)
- [cf](https://github.com/cloudfoundry/cli)

Now that you have all of these CLIs set up already from prior stories, let's dig into them.

**How to get help!**

There are many common ways to get help text about the CLI you want to use. Unfortunately, nothing is completely standardized, so you might have to try a few options first.

In your terminal type `bbl --help` and see the results. It will show everything that's available via the bbl cli.

Typically you can append `--help`, `-help`, `-h`, or `help` and get help text, but nothing is standardized, so you might need to try each to find out which one works.

For example all of the following are valid commands to get help for the `bbl` cli.
```
bbl --help
bbl -h
bbl help
```

The `help` results will typically return a list of commands you can run with the CLI. If there is a specific CLI command you want to find out more about you can append the same help words at the end of the command. For example `bbl env-id --help` will give help text about the `env-id` command.

**Reading CLI usage syntax**
One of the common features of help text, is showing the CLI usage. The usage is written in a common syntax.

For example here is the usage for the `bosh deploy` command: `bosh [OPTIONS] deploy [DEPLOY-OPTIONS] PATH`

- `[Square brackets]` indicate an optional parameter
- `CAPITALS` indicate a value (or values) that the user needs to input and change for their needs. When you replace the text it should not be capitalized.
- `--words-prepended-with-dashes` indicates a flag to change how the command is run. The flags usage should be explained in the help text.
-  `--words-prepended-with-dashes-with-equal-sign-after=VALUE` indicates an option that the user assigns a value to. Its usage should be explained in the help text.

Let's try using a flag. Run `bosh instances`.  `this command runs fine when you have only 1 deployment`

You will get the helpful error message `expected argument for flag -d, --deployment`. This is telling you that you must use the `-d` flag to specify a bosh deployment, in order to return the bosh instances.
Given this, let's try running `bosh -d YOUR-DEPLOYMENT-NAME instances`. Success! (If you're having trouble getting your deployment name try `bosh deployments --help`.

**The Linux CLI**
If you've been using the terminal, you've been using a CLI. Commands like `cd` (change directory), `ls` (list), and `cp` (copy) are all apart of the linux and mac CLIs.
In order to get help for the linux CLI you need to prepend the command with `man`, which stands for manual.

For example `man cp` results in the following manual for `cp`
```
NAME
     cp -- copy files

SYNOPSIS
     cp [-R [-H | -L | -P]] [-fi | -n] [-apvX] source_file target_file
     cp [-R [-H | -L | -P]] [-fi | -n] [-apvX] source_file ... target_directory

DESCRIPTION
     In the first synopsis form, the cp utility copies the contents of the source_file to the
     target_file.  In the second synopsis form, the contents of each named source_file is copied
     to the destination target_directory.  The names of the files themselves are not changed.  If
     cp detects an attempt to copy a file to itself, the copy will fail.
```
Manual pages can be a little bit dense, so always feel free to google things like `linux terminal copy one file to a different folder` and you should be able to find some help from stack overflow.

### How?
**Let's practice using CLIs**

0. Look at the help text for credhub
0. Add a `value` credential type  with any contents you want to credhub
___
0. Look at the help text for `bosh` and for `bosh logs`
0. Use `bosh logs` to download the logs for all diego-cell instances
___
0. Use `mkdir` to make a new temp file `learning-clis` in the `/tmp` directory.
0. Use `vim` to make a new file `hello-world.sh` in your new directory. In the file write `echo "hello world"`
0. Use `chmod` to make the file `hello-world.sh` executable
0. Run your new script!

### Expected Result
You should be able to add a credential to credhub. You should see something like the following output:
```
id: 27b47c13-d427-4758-8857-8c176a754490
name: /my-super-cool-value
type: value
value: <redacted>
version_created_at: "2019-03-22T23:51:13Z"
```
You should be able to get bosh logs and see them downloaded on your computer

```
Task 3678 | 23:52:24 | Fetching logs for diego-cell/a21fbda8-0a3d-41cb-9f16-6ff3a3286cd2 (0): Finding and packing log files (00:00:05)

Task 3678 Started  Fri Mar 22 23:52:24 UTC 2019
Task 3678 Finished Fri Mar 22 23:52:29 UTC 2019
Task 3678 Duration 00:00:05
Task 3678 done

Downloading resource '070ca5d5-699b-4b1c-55e8-81efe931fc68' to '/tmp/cf.diego-cell.0-20190322-165231-227811.tgz'...

#############################################################  99.94% 8.54 MiB/s
Succeeded
```

Your hello-world.sh script should print out `hello world`.

### Resources
[linux command line primer](https://lifehacker.com/a-command-line-primer-for-beginners-5633909)
[how to make a new directory](https://www.techwalla.com/articles/how-to-make-a-new-directory-in-linux)
[how to make a file executable and run it](https://askubuntu.com/questions/229589/how-to-make-a-file-e-g-a-sh-script-executable-so-it-can-be-run-from-termina)