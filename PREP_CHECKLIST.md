## Instructions
- Start a new issue in the onboarding repo. The title should include the date and location of the Onboarding week (e.g. "SF Onboarding Week, July 16 2018").
- Copy the checklist below into the issue.
- Create the issue
- Use the checkboxes to track your work as you get set up.
- When you finish with the checklist, make sure to close the issue.

## Checklist

### Communication
- [ ] Send out an email to each of the participants. You can find a template in FACILITATING.md.

### IaaS Setup
- [ ] Make sure you have clean GCP projects set up. In SF, we use `cf-sf-onboarding-env-1`, `cf-sf-onboarding-env-2`, and `cf-sf-onboarding-env-3`, which already exist. Otherwise, fill out [this form to create new environments](https://docs.google.com/forms/d/e/1FAIpQLSeJ31997Zma1WtLcCtswiysCFWOG5MXNmlYCpJsiYgdG9kKnA/viewform).
- [ ] Add all participants as **owners** to each of the GCP projects.

### Tracker Setup
- [ ] Make sure you have Tracker projects prepared. Again, in SF we re-use the same tracker projects for [Pair 1](https://www.pivotaltracker.com/n/projects/2125981), [Pair 2](https://www.pivotaltracker.com/n/projects/2125982), and [Pair 3](https://www.pivotaltracker.com/n/projects/2089066).
- [ ] Add all participants as members to each of the tracker projects.

### Workstation Setup
- [ ] Make sure you have a workstation for each pair.
- [ ] Re-image each machine. Here are the steps:
  1. [ ] Run `Startup Disk`.
  1. [ ] Unlock with the workstation password (you may need to ask the previous facilitator for the machine's password).
  1. [ ] Select the network disk and click restart. It will likely be the rightmost start disk, and will definitely include the phrase "Network Disk." If you don't see that, try restarting the machine and try again.
  1. [ ] Select deployment option 2, "macOS Sierra" and click the play button. If you look in between the parens of the option name, you'll see the default password for the image. A terminal will open and you’ll see a lot of output, including a bunch of lines that say `installer: PHASE: Running package scripts`. Don’t worry, this is fine. The boot script will search for updates -- you’ll see output about OS X, iTunes, and Safari updates, for example -- and the machine will probably restart a few times as it downloads and installs the updates.
  1. [ ] Log into machine with the default password.
  1. [ ] Open the security preferences and change the password to something sane. Write this down so you can give it to the participants later.
  1. [ ] Open a terminal window and run
     ```
     curl -LOk https://github.com/pivotal-cf-experimental/onboarding-week-construct/archive/master.zip
     unzip master.zip
     cd onboarding-week-construct
     ./construct
     ```
     You'll need to babysit this process, as it will prompt you for the password several times.
  1. Record any issues you have with this process so we can update this checklist.

### Scheduling
- [ ] Schedule a welcome/orientation session for Monday morning. Discuss the goals of the week and prime the participants to consider what they'd like to get out of Onboarding week.
- [ ] Schedule daily standup.
- [ ] Schedule afternoon check in.
- [ ] Schedule retro at the end of the week.
- [ ] Schedule any sessions that you know the group would be interested in -- for example, BOSH boxes and lines.
