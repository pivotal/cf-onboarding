#!/bin/bash

#### USAGE
# ./clean_up_gcp_project.sh <PATH TO SERVICE ACCOUNT KEY>

#### HOW TO FIND FILE PATH TO SERVICE ACCOUNT KEY
# 1. Navigate to the list of service accounts for the GCP Project you want to delete.
# 2. Click on one and choose "+ Create Key" to create a JSON key. This will
# download the file containing the key to your workstation.
#    - There may be multiple service accounts for a single GCP Project. 
#    - The [default serviceaccount](https://cloud.google.com/iam/docs/service-accounts#default) 
#      has a large number of permissions and may be a good choice.
#     - Consider updating this document with instructions on how to identify a default service account.
# 3. Copy the path to the file.

display_usage() {
  echo "Please provide the one required argument: the file path to a service account key for the target CF Onboarding Env's GCP Project."
}

# if less than one argument supplied, display usage
if [  $# -lt 1 ]
then
  display_usage
  exit 1
fi

export BBL_GCP_SERVICE_ACCOUNT_KEY=$(cat $1)
export BBL_GCP_REGION=us-west1
export BBL_IAAS=gcp

read -p "Are you sure? " -n 1 -r
if [[ $REPLY =~ ^[Yy]$ ]]
then
  eacho "Starting to clean up GCP resources ..."
  bbl cleanup-leftovers --no-confirm
  echo "Finished cleaning up GCP resources."
fi
