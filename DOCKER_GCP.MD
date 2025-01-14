## **DOCKER BUILD & RUN**

- ATAU 
- docker-compose up --build ( kalau sudah ada docker-compose.yml )
- docker build -t gc2-p3-xiowel:latest .
- docker run -p 8080:8080 -v "$PWD\.env:/.env" gc2-p3-xiowel:latest
- ATAU 
- docker run -p 8080:8080 -v "E:\learn-code\HACKTIV8\PHASE 3\WEEK2\graded-challange-2-XioweL\.env:/.env" gc2-p3-xiowel:latest
- docker run -p 8080:8080 -v "${pwd}\.env:/.env" gc2-p3-xiowel:latest


## **DOCKER BUILD & RUN**
- gcloud auth login
- gcloud config set project gradedc2(PROJECT_ID)
- docker gc2-p3-xiowel gcr.io/gradedc2/gc2-p3-xiowel:latest  / docker tag(latest) gc2-p3-xiowel gcr.io/gradedc2/gc2-p3-xiowel
- docker push gcr.io/gradedc2/gc2-p3-xiowel / docker push gcr.io/PROJECT_ID/my-go-app
- gcloud run deploy xiowel-service --image gcr.io/gradedc2/gc2-p3-xiowel:latest --
  platform managed --region asia-southeast2 --allow-unauthenticated --port 8080

## **IF ERROR** 
- denied: Unauthenticated request. Unauthenticated requests do not have permission "artifactregistry.repositories.uploadArtifacts" on resource "projects/gradedc2/locations/us/repositories/gcr.io" (or it may not exist)
- gcloud auth configure-docker
