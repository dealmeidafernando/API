# API Job Documentation 

The API's main purpose is to provide resources for the Job:
- Create a Job
- Edit a Job
- Search by ID
- Search all
- Delete a job.

## Pre-requisites

`mysql version 5+` See how to download and install in [Mysql site.](https://dev.mysql.com/downloads/repo/apt/)

`golang version 1.11+`  See how to download and install in [Golang site.](https://golang.org/doc/install)

`postman version 6.4.2+`  See how to download and install in [Postman site.](https://www.getpostman.com/apps)


## Development

First of all, download the project to your machine.

After download, execute script of 

After download, enter in the root project directory:
```
cd api-golang
```

Create a `local.env` file (see `local.env.sample`):
```
cp local.env.sample local.env
```

And now you can up the application:
```bash
docker-compose up
```

Note: *Everytime you change some code file, you will must run*  ```docker-compose up --build``` *to compile the binary and run the application again.*

Then you will be able to make API calls like so:
```
curl -v -X GET "http://localhost:8080/health"
```

## API

Check the folder `docs` to download and import the Postman collection `API-Job.postman_collection.json`


## Building

JobAd-Inspector has some external dependencies, such as JobAd Id Generator, which is hosted on private repositories on GitHub. 

If `make build` command fails and returns the following error, `could not read Username for 'https://github.com': terminal prompts disabled`, you may need to execute one of those steps bellow:

* if you don't use ssh-key, set `GIT_TERMINAL_PROMPT=1` environment variable:

* if you use ssh-key, check if your git global configs (`git config --global --list`) has the configuration bellow:

    `url.git@github.com:.insteadof=https://github.com/`
    
if your git-client doesn't have this config, execute the command bellow:
 ```sh
 git config --global --add url."git@github.com:".insteadOf "https://github.com/"
 ```


