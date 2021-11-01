# project

An event listener to store information in the database.

## Requirements

* [Docker 20.10.6+](https://www.docker.com/get-started)
* [Compose 3.3+](https://docs.docker.com/compose/install/)
* [Go 1.15+](https://golang.org/)
* [Postgresql 12.6](https://www.postgresql.org/)

## Running the service
#### For development purposes
1. Modify *config.yaml* file with your needs:

    * Provide a database url where the information will be stored in the form as provided [here](https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING)

      ```sh
      db:
         url: "postgresql://[userspec@][hostspec][/dbname][?paramspec]"
      ```

    * If you are using local blockchain (e.g Ganache), provide other endpoint

      ```sh
      rpc:
         endpoint: "https://localhost"
      ```

2. Open project using IDE (e.g GoLand from JetBrains), open IDE terminal, run

   ```sh
   go mod init
   ```

3. At *bsc-checker-events/assets/main.go* run two migration related scripts:
   ```sh
   //go:generate packr2 clean
   //go:generate packr2
    ```
4. Modify run configuration as follows:
    * KV_VIPER_FILE=config.yaml *(environment variable)*
5. Run service twice with the following command arguments:

   ```sh
   migrate up
   run service
   ```

#### For deployment purposes
1. Navigate to the cloned repository
2. Do the step 1 from development build, except modify config at *configs/spaceship-staking.yaml*, changing contract address and database url (for the contract deployed on the Binance Smart Chain leave the endpoint as it is
3. Build container image:

   ```sh
   docker build -t spaceship-staking .
	```
4. Run using docker-compose
   ```sh
   docker-compose down -v
   docker-compose up -d
	```

### API
To change port, configure
```sh
listener:
  addr: :8010
```
where *8010* is a port to listen on.

#### Endpoints
```sh
/add # add new person to the database
/list #get list of people in the database
/get/{id} # get person by it's id
/info # write message "It's our database"
```

