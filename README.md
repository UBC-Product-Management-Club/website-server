# UBC PMC Website Server
## SetUp
### Go Environemnt
First, make sure that your computer has the latest version of Go (go1.20.4). Follow these [instructions](https://go.dev/doc/install)\
Next, setup a workspace for local project
```
mkdir -p <project-name>/bin
mkdir -p <project-name>/pkg
go env -w GOPATH=/PATH/TO/<project-name>
go env -w GOBIN=/PATH/TO/<project-name>/bin
```
Clone the project and install all the dependencies.
```
cd <project-name>
git clone git@github.com:UBC-Product-Management-Club/productsprint-website.git src
cd src
go mod tidy
```
Anywhere in the project, to run the server
```
go run .
```
## Curl commands
Here are the curl commands to interact with the server.\
**GET Requests**
```bash
curl http://localhost:8080/executive # execs
curl http://localhost:8080/fellow    # fellows
curl http://localhost:8080/project   # projects
```
**POST Requests**
```bash
curl POST http://localhost:8080/executive \
-H "Content-Type: application/json" \
-d '{"img": "url", "name": "string", "title": "string"}'


curl POST http://localhost:8080/fellow \
-H "Content-Type: application/json" \
-d '{"img": "url", "name": "string", "title": "Fellow", "bio_text": "string", "linkedin": "url", "projects": [{"Parent":{"Parent":null,"Path":"projects/pmc-website-bfa1a/databases/(default)/documents/project","ID":"project"},"Path":"projects/pmc-website-bfa1a/databases/(default)/documents/project/aEUFciFC3md5GdkEBAD9","ID":"aEUFciFC3md5GdkEBAD9"}]}]}'

curl POST http://localhost:8080/project \
-H "Content-Type: application/json" \
-d '{"isFinished": true, "title": "string", "text": "string", "image": "url", "link": "url", "detailId": "id"}'
```
## Notes
- To see all the endpoints, view **routes.go** in api/routes.
- More secure auth should be used. CORS should be customized for more security.
- Reading and writing to the database currently doesn't have any timeout. It should be implemented in the future with context.

## Running Docker
There are two Dockerfiles: Production and Development. The development container has more features while production is more lightweight.

Before running Docker, make sure you have Docker and Docker Desktop installed on your device.

Next, build the Docker image by running:
```
docker build -t <Insert Docker Image Name> -f <Insert Dockerfile Name> .
```

In Docker Desktop, run the image and make sure you enable port forwarding to 8080.