#!/bin/bash

# Each service account key is tied to a single GCP project.
export BBL_GCP_SERVICE_ACCOUNT_KEY=$(cat $1)
export BBL_GCP_REGION=us-west1
export BBL_IAAS=gcp

bbl cleanup-leftovers --no-confirm
