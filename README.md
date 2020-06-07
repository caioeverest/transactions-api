# Transactions API

This is a basic experimental testing API.

##Testing:

The API has unity tests on the service and handler packages. To simply run the tests you can execute the prompt `make test`. It will print the basic test report with the coverage from the service and handler packages. You can also get a more complete report by executing the `make test-report` command. It will open a report on your default web browser.

##Building:

You can build a binary file from the app by executing the prompt `make build`. This will create a folder bin on the current directory and, inside it, there will be a transaction-api file that you can execute.

> The executable expects an application.yml file to find its configurations.

##How to run it?

###Environment variables:

The app uses the application.yml file to map its variables. By default, it has the values:

| Variable reference | Environment variable | Default value  |
|--------------------|----------------------|----------------|
|         ENV        |          ENV         |   development  |
|      HTTP.Port     |       HTTP_PORT      |      8080      |
|   HTTP.Greetings   |    HTTP_GREETINGS    |   hello there  |
|    Database.User   |        DB_USER       |      admin     |
|  Database.Password |      DB_PASSWORD     |      admin     |
|    Database.Host   |        DB_HOST       |    localhost   |
|    Database.Port   |        DB_PORT       |      5432      |
|   Database.DbName  |        DB_NAME       | transactionsdb |

###Developer mode:

The API uses a Postgres database to store its information. By default, it expects the data source to be in the host machine. If you don't have a postgres ready, you can execute the prompt to make the postgres. It will run a postgres container for you.

To run the API you can simply use the command `make run`. It will execute the main file on the cmd folder and start the application.

###Container:

To run all you can simply use `make docker-run`. It will automatically start the Postgres and the Transactions API. You can check it by accessing `http://localhost:8080/health`

###API Reference:

The full API documentation can be found at http://localhost:8080/docs

| Variable reference | Environment variable | Default value                                    |
|--------------------|----------------------|--------------------------------------------------|
|       Method       |         Path         |                    Description                   |
|         GET        |        /health       |                Return health check               |
|         GET        |         /docs        |               Open the swagger docs              |
|        POST        |       /accounts      |               Create a new account               |
|         GET        |       /accounts      |          List all account on repository          |
|         GET        | /accounts/:accountID | Return the account corresponding that account ID |
|        POST        |     /transactions    |             Create a new transaction             |
|         GET        |     /transactions    |               List all transactions              |
