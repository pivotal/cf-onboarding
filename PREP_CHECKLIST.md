## Instructions
- Start a [new issue](https://github.com/pivotal/cf-onboarding/issues/new) in the onboarding repo. The title should include the date and location of the Onboarding week (e.g. "SF Onboarding Week, July 16 2018").
- Copy the checklist below into the issue.
- Create the issue
- Use the checkboxes to track your work as you get set up.
- When you finish with the checklist, make sure to close the issue.

## Checklist

### Communication
- [ ] Several weeks before: send out an email to the participants reminding them that they are signed up and to clear their calendars.  You can find a template in [FACILITATING.md](https://github.com/pivotal/cf-onboarding/blob/master/FACILITATING.md).
  - [ ] You will be communicating with / authing these folks a lot. I find it helpful to save a copy/pasteable list of their names <emails> here:
    - `name <email>.name <email>`
- [ ] The week before: send out an email to each of the participants. You can find a template in [FACILITATING.md](https://github.com/pivotal/cf-onboarding/blob/master/FACILITATING.md).
- [ ] Get a [Zoom](https://zoom.us/) room setup for remote participants. {<can save link here>}  (Make sure to include this on your calendar invites.)

### Scheduling
- [ ] Schedule a welcome/orientation session (1hr at 9:30 Monday morning). Discuss the goals of the week and prime the participants to consider what they'd like to get out of Onboarding week.
- [ ] Schedule daily standup (15min at 9:15, starting Tuesday).
- [ ] Schedule afternoon check in (30min at 1:30).
- [ ] Schedule retro at the end of the week (1hr at 4:00), include link to retro board: https://retros.cfapps.io/retros/onboarding-sf (pw: onboarding-sf (pls don't change it))
- [ ] Schedule any sessions that you know the group would be interested in -- for example, BOSH boxes and lines.
- [ ] In the past, non Engineers have appreciated it when we blocked out the entire days of onbaording. This is especially helpful if done a few weeks in advance; so, they have time to move meetings and avoid having new ones scheduled in the first place.
---
These are core CF topics, other topics are a bit more "extra credit".
- [ ] BOSH {}
- [ ] CF Overview {}
- [ ] Concourse {}

As a helpful reminder, you can save the name of the speaker w/ checkbox. Tips to
share with speakers can be found in [BOXES_AND_LINES.md](https://github.com/pivotal/cf-onboarding/blob/master/BOXES_AND_LINES.md).

### IaaS Setup
- [ ] Make sure you have [GCP projects](https://console.cloud.google.com) set up.
  - [ ] In SF, we use:
    - [cf-sf-onboarding-env-1](https://console.cloud.google.com/billing/0076DC-766E1F-EBDCB8/reports?project=cf-sf-onboarding-env-1)
    - [cf-sf-onboarding-env-2](https://console.cloud.google.com/billing/0076DC-766E1F-EBDCB8/reports?project=cf-sf-onboarding-env-2)
    - [cf-sf-onboarding-env-3](https://console.cloud.google.com/billing/0076DC-766E1F-EBDCB8/reports?project=cf-sf-onboarding-env-3)
    - [cf-sf-onboarding-env-4](https://console.cloud.google.com/billing/0076DC-766E1F-EBDCB8/reports?project=cf-sf-onboarding-env-4)
    - [cf-sf-onboarding-env-5](https://console.cloud.google.com/billing/0076DC-766E1F-EBDCB8/reports?project=cf-sf-onboarding-env-5)
  - [ ] Otherwise, Fill out [this form to create new environments](https://docs.google.com/forms/d/e/1FAIpQLSeJ31997Zma1WtLcCtswiysCFWOG5MXNmlYCpJsiYgdG9kKnA/viewform).
- [ ] Run [this script](scripts/clean_up_gcp_project.sh) to clean up any VMs, networks, and DNS records that may have been leftover from the previous session.
- [ ] Add all [participants](https://docs.google.com/spreadsheets/d/1eW23iJRD56CE859o0j6ArSaGtLfP0JlHXSCdXFhbCdI/edit?usp=sharing) as **owners** to each of the GCP projects.
  - [ ] In SF:
    - [IAM for env-1](https://console.cloud.google.com/iam-admin/iam?project=cf-sf-onboarding-env-1)
    - [IAM for env-2](https://console.cloud.google.com/iam-admin/iam?project=cf-sf-onboarding-env-2)
    - [IAM for env-3](https://console.cloud.google.com/iam-admin/iam?project=cf-sf-onboarding-env-3)
    - [IAM for env-4](https://console.cloud.google.com/iam-admin/iam?project=cf-sf-onboarding-env-4)
    - [IAM for env-5](https://console.cloud.google.com/iam-admin/iam?project=cf-sf-onboarding-env-5)
- [ ] If your CF on-boarding will include the GCP track (`gcp` tracker label / the `oss` build) then ensure participants have access to the `CF-Onboarding-dns` project in GCP (for DNS propagation). [IAM for DNS](https://console.cloud.google.com/iam-admin/iam?project=cf-onboarding-dns)

### Tracker Setup
- [ ] Make sure you have Tracker projects prepared.
  - [ ] In SF, we use the following trackers projects:
    - [Pair 1](https://www.pivotaltracker.com/n/projects/2125981)
    - [Pair 2](https://www.pivotaltracker.com/n/projects/2125982)
    - [Pair 3](https://www.pivotaltracker.com/n/projects/2089066)
    - [Pair 4](https://www.pivotaltracker.com/n/projects/2361806)
    - [Pair 5](https://www.pivotaltracker.com/n/projects/2361807)
- [ ] Add all participants as members to each of the tracker projects. (A list of emails works here.)
- [ ] Delete everything from the tracker projects.
- [ ] Build the tracker back log csv
  - [ ] Run `./build oss` from the onboarding project, this will generate `onboarding-tracker.csv`
       - Note: you need docker installed `brew cask install docker`
         - To get the command line tool to show up you may have to first start docker via Spotlight Search
         - If docker gives you this error: `failed to dial gRPC: unable to upgrade to h2c, received 502 context canceled`
           - In docker settings: enable experimental, restart docker, disable experiment and restart again.
- [ ] Import the CSV into each Tracker project (Taskbar > More > Import CSV > Choose File)

### LastPass Setup
- Skip this step for participants who are not authorized to access closed-source content (e.g. Pivotal Labs client participants.)
- [ ] Give participants access to the shared LastPass folder:
  - LastPass > Sharing Center > Shared-CF SF Onboarding > Manage > Invite Users (A list of emails works here.)
### Workstation Setup
Note: local IT departments can sometimes help with workstation setup. (Helpful for remote facilitation.)

- [ ] Make sure you have a workstation for each pair.
- [ ] Ensure ethernet cable is plugged into iMac, not monitor.
- [ ] Re-image each machine. [Here](WORKSTATION_IMAGING.md) are steps:
- [ ] Open a terminal window and run
 ```
 curl -LOk https://github.com/pivotal-cf-experimental/onboarding-week-construct/archive/master.zip
 unzip master.zip
 cd onboarding-week-construct-master
 ./construct
 ```
 You'll need to babysit this process, as it will prompt you for the password several times.

### Misc

- [ ] Record any issues you have with this process so we can update this checklist.
