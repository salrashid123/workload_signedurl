
package com.test;

import java.net.URL;
import java.util.Arrays;
import java.util.concurrent.TimeUnit;

import com.google.auth.oauth2.GoogleCredentials;
import com.google.auth.oauth2.ImpersonatedCredentials;
import com.google.cloud.storage.BlobInfo;
import com.google.cloud.storage.HttpMethod;
import com.google.cloud.storage.Storage;
import com.google.cloud.storage.StorageException;
import com.google.cloud.storage.StorageOptions;

// mvn  clean install exec:java -q

public class TestApp {
	public static void main(String[] args) {
		TestApp tc = new TestApp();
	}

	public TestApp() {
		try {

			String targetServiceAccount = "urlsigner@fabled-ray-104117.iam.gserviceaccount.com";
			String BUCKET_NAME1 = "some-bucket";

			GoogleCredentials credentials = GoogleCredentials.getApplicationDefault();
			ImpersonatedCredentials targetCredentials = ImpersonatedCredentials.create(credentials,
					targetServiceAccount, null, Arrays.asList("https://www.googleapis.com/auth/devstorage.read_only"),
					5);

			Storage storage_service = StorageOptions.newBuilder().setCredentials(targetCredentials)
					.build()
					.getService();			

			String BLOB_NAME1 = "file1.txt";
			BlobInfo BLOB_INFO1 = BlobInfo.newBuilder(BUCKET_NAME1, BLOB_NAME1).build();

			URL url = storage_service.signUrl(
					BLOB_INFO1,
					60,
					TimeUnit.DAYS,
					Storage.SignUrlOption.httpMethod(HttpMethod.GET));
			System.out.println(url);

		} catch (StorageException ex) {
			ex.printStackTrace();
			System.out.println("Error:  " + ex);
		} catch (Exception ex) {
			ex.printStackTrace();
			System.out.println("Error:  " + ex);
		}
	}

}
