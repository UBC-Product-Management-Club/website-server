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
## Notes
- To see all the endpoints, view **routes.go** in api/routes.
- The database still only has a basic template.
- More secure auth should be used. CORS should be customized for more security.