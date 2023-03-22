const { GoogleAuth, OAuth2Client, Impersonated, IdTokenClient } = require('google-auth-library');
const { Storage } = require('@google-cloud/storage');
  
async function main() {

  const targetPrincipal = "urlsigner@fabled-ray-104117.iam.gserviceaccount.com";
  const bucketName = "some-bucket";
  let projectId = 'fabled-ray-104117'
  fileName = "file.txt"


  const scopes = 'https://www.googleapis.com/auth/cloud-platform'

  // get source credentials
  const auth = new GoogleAuth({
    scopes: scopes
  });
  const client = await auth.getClient();

  // First impersonate
  let targetClient = new Impersonated({
    sourceClient: client,
    targetPrincipal: targetPrincipal,
    lifetime: 30,
    delegates: [],
    targetScopes: [scopes]
  });

  // use the impersonated creds to access gcs
  const storageOptions = {
    projectId,
    authClient: targetClient,
  };

  const storage = new Storage(storageOptions);

  // ####### GCS SignedURL
  // pending https://github.com/googleapis/google-auth-library-nodejs/issues/1443


  const options = {
    version: 'v4',
    action: 'read',
    expires: Date.now() + 15 * 60 * 1000, // 15 minutes
  };

  const  su = await storage
    .bucket(bucketName)
    .file('foo.txt')
    .getSignedUrl(options);

  console.log(su);

}

main().catch(console.error);