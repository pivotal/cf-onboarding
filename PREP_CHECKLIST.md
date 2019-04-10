## Instructions
- Start a new issue in the onboarding repo. The title should include the date and location of the Onboarding week (e.g. "SF Onboarding Week, July 16 2018").
- Copy the checklist below into the issue.
- Create the issue
- Use the checkboxes to track your work as you get set up.
- When you finish with the checklist, make sure to close the issue.

## Checklist

### Communication
- [ ] Several weeks before: send out an email to the participants reminding them that they are signed up and to clear their calendars.  You can find a template in [FACILITATING.md](FACILITATING.md).
  - [ ] You will be communicating with these folks a lot. I find it helpful to save a copy/pasteable list of their names <emails> here:
    - `name <email>.name <email>`
- [ ] The week before: send out an email to each of the participants. You can find a template in [FACILITATING.md](FACILITATING.md).
- [ ] Get a [Zoom](https://zoom.us/) room setup for remote participants. `<can save link here>`  (Make sure to include this on your calendar invites.)

### Scheduling
- [ ] Schedule a welcome/orientation session (1hr at 9:30 Monday morning). Discuss the goals of the week and prime the participants to consider what they'd like to get out of Onboarding week.
- [ ] Schedule daily standup (15min at 9:15, starting Tuesday).
- [ ] Schedule afternoon check in (30min at 1:30).
- [ ] Schedule retro at the end of the week (1hr at 4:00), include link to retro board: https://retros.cfapps.io/retros/onboarding-sf
- [ ] Schedule any sessions that you know the group would be interested in -- for example, BOSH boxes and lines.
- [ ] In the past, non Engineers have appreciated it when we blocked out their whole day.

### IaaS Setup
- [ ] Make sure you have [GCP projects](https://console.cloud.google.com) set up. In SF, we use `cf-sf-onboarding-env-1`, `cf-sf-onboarding-env-2`, and `cf-sf-onboarding-env-3`, which already exist. Otherwise, fill out [this form to create new environments](https://docs.google.com/forms/d/e/1FAIpQLSeJ31997Zma1WtLcCtswiysCFWOG5MXNmlYCpJsiYgdG9kKnA/viewform).
- [ ] Clean up any VMs, networks, and DNS records that may have been leftover from the previous week
- [ ] Add all [participants](https://docs.google.com/spreadsheets/d/1eW23iJRD56CE859o0j6ArSaGtLfP0JlHXSCdXFhbCdI/edit?usp=sharing) as **owners** to each of the GCP projects.
- [ ] If your CF on-boarding will include the GCP track (`gcp` tracker label / the `oss` build) then ensure participants have access to the `CF-Onboarding-dns` project in GCP (for DNS propagation).

### Tracker Setup
- [ ] Make sure you have Tracker projects prepared. Again, in SF we re-use the same tracker projects for [Pair 1](https://www.pivotaltracker.com/n/projects/2125981), [Pair 2](https://www.pivotaltracker.com/n/projects/2125982), and [Pair 3](https://www.pivotaltracker.com/n/projects/2089066).
  1. [ ] SF FACILITATORS: Run `./build oss` from the onboarding project, this
     will generate `onboarding-tracker.csv`
       - Note: you need docker installed `brew cask install docker`
         - if this fails: `failed to dial gRPC: unable to upgrade to h2c, received 502 context canceled` - enable experimental, restart docker, disable experiment and restart again. 
  1. [ ] SF FACILITATORS: Import the CSV into each Tracker project (Taskbar >
     More > Import CSV > Choose File)
  1. [ ] Move all stories from icebox to backlog.
- [ ] Add all participants as members to each of the tracker projects.

### LastPass Setup
- [ ] Give all participants access to the shared LastPass folder:
  - LastPass > Sharing Center > Shared-CF SF Onboarding > Manage > Invite Users

### Workstation Setup
- [ ] Make sure you have a workstation for each pair.
- [ ] Ensure ethernet cable is plugged into iMac, not monitor.
- [ ] Re-image each machine. Here are the steps:
  - [ ] Run `Startup Disk`.
  - [ ] Unlock with the workstation password (you may need to ask the previous facilitator for the machine's password).
  - [ ] Select the network disk and click restart. It will likely be the rightmost start disk, and will definitely include the phrase "Network Disk." If you don't see that, try restarting the machine and try again.
  - [ ] Select deployment option 2, "macOS Sierra" and click the play button. If you look in between the parens of the option name, you'll see the default password for the image. A terminal will open and you’ll see a lot of output, including a bunch of lines that say `installer: PHASE: Running package scripts`. Don’t worry, this is fine. The boot script will search for updates -- you’ll see output about OS X, iTunes, and Safari updates, for example -- and the machine will probably restart a few times as it downloads and installs the updates.
  - [ ] Log into machine with the default password.
  - [ ] Open the security preferences and change the password to something sane. Write this down so you can give it to the participants later.
  - [ ] Open a terminal window and run
     ```
     curl -LOk https://github.com/pivotal-cf-experimental/onboarding-week-construct/archive/master.zip
     unzip master.zip
     cd onboarding-week-construct-master
     ./construct
     ```
     You'll need to babysit this process, as it will prompt you for the password several times.
- [ ] Record any issues you have with this process so we can update this checklist.
