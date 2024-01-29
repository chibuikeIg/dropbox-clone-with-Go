# Dropbox/Google Drive Clone With Go

Dropbox/Google drive clone with Go is a side project inspired by Nikhil Gupta article titled [Design Dropbox/Google Drive
](https://nikhilgupta1.medium.com/design-dropbox-google-drive-81cd343571a8). Following the design architecture described in the article and the Hexagonal architecture, I learnt how to develop and design micro services and cloud native applications using Go. Everything was not covered in this project and a lot more challenges and improvements will be tackled as I continue working on this project.

## Challenges & Solutions

- Generating random unique number between 1 - 1000 per multi upload parts: Using recursion I traversed through a map containing previously generated/stored values to check if the current generated number already exists and if it does I regenerate another to ensure the numbers are not generated twice.
- Running Database Migration in CMD without running application server: Using Go command-line flags package, I programmed arguments and options which enables me to check if a user is trying to run the application server or DB migration via the command line.
- Service Discovery: Using Docker-compose and consul, I ensured each service is running on the same network and each has the `CONSUL_HTTP_ADDR` environment variable created during containerization. Each service has a service registration function which registers them in consul for other services to be able to send requests and receive responses. This pattern of communication is only being used by the API Gateway service as I aim to use gRPC for service to service communication.

## Technologies/Tools

- Languages: Golang
- Web Frameworks: Gin
- Databases: DynamoDB
- Tools: Git, Docker, Consul, Redis
- Cloud Services: AWS (S3)
- Frontend Technologies: HTML, CSS, JavaScript (Vue.js)
- Problem-solving and algorithmic skills

### Running The Application

To run this application clone this repository down to your local machine and follow the below instructions. This application requires you to have docker setup and running in your machine

- CD into working directory
- In each service folder create a `.env` file following the `.env-example` and ensure you have all the values setup in it
- In the root directory of the application where you have the `go.work` file run `go mod tidy` to download all required packages
- Run migrations in each service folder using `go run . create -table_name='name' -primary_key='primary_key' -range_key='range_key'`. Add the `-mt=update` flag to update table.
- Use `docker-compose up --build` to run and build all services once.
- Test endpoints using postman or any client service of your choice

#### Required Tables

- API_GATEWAY: `table_name=users, primary_key=ID, range_key=email, secondary_Index=email`
- FILE_META_DATA_SERVICE: `table_name=files, primary_key=id, range_key=userid,folderid, secondary_Index=userid,folderid; table_name=folders, primary_key=id, range_key=userid, secondary_Index=userid`
- FILE_UPLOAD_SERVICE: `table_name=upload_parts -primary_key=uploadid -range_key=userid`
- USER_SERVICE: `table_name=users, primary_key=ID, range_key=email, secondary_Index=email`

### Design Screenshot

![Service Design Screenshot](https://miro.medium.com/v2/resize:fit:828/format:webp/1*XlJby2-ltVG3b7lNm35haA.png)
