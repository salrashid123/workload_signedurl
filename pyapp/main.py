#!/usr/bin/python
from google.auth import credentials, impersonated_credentials

import google.auth
import google.oauth2.credentials
from datetime import datetime, timedelta
from google.cloud import storage

targetPrincipal = 'urlsigner@fabled-ray-104117.iam.gserviceaccount.com'
bucket_name      = 'some-bucket'
file_name='file1.txt'

credentials, project = google.auth.default()   
target_scopes = ['https://www.googleapis.com/auth/devstorage.read_only']
target_credentials = impersonated_credentials.Credentials(
    source_credentials = credentials,
    target_principal=targetPrincipal,
    target_scopes = target_scopes,
    delegates=[],
    lifetime=500)

storage_client = storage.Client(project=project)

data_bucket = storage_client.bucket(bucket_name)
blob = data_bucket.blob(file_name)

expires_at_ms = datetime.now() + timedelta(minutes=2)
signed_url = blob.generate_signed_url(expires_at_ms, credentials=target_credentials)

print(signed_url)