# What is gitstarseeker do?

1. Read source file to find git repositories.
2. Using github api to fetch repository information, such like description, starts, issues.
3. Make simple goroutine worker pool to fetch mass repositories.
4. Store above data to DB.
5. Set up cron job to update information.
6. provide api or other to show those data





## Setup

1. create ./config/config.yaml, and add following fields
```
mode: DEV
 db:
   username: DB_USERNAME
   password: DB_PASSWORD
   database: DB_NAME

 http:
   workernums: 10
   reqtimeout: 800
   github_header_accept: "application/vnd.github.v3+json"
   github_header_authorization: GITHUB_TOKEN
source:
  awesome-go: "https://raw.githubusercontent.com/avelino/awesome-go/master/README.md"
```
