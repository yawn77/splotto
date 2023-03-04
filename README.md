# SP Lotto
Bot that plays lotto for you on https://www.spieleplanet.eu/lotto.php.

## Features

## Usage

### Binary
* Download binary from [Github](https://github.com/yawn77/splotto/releases)
* Set environment variables
  * `SP_USERNAME=<your username>`
  * `SP_PASSWORD=<your password>`
  * `TZ=<your timezone>` (e.g., America/Los_Angeles)

### Docker
* Pull latest docker image:
```
docker pull yawn77/splotto:latest
```

* Run docker container:
```
docker run --rm \
  -e SP_USERNAME=<your username> \
  -e SP_PASSWORD=<your password> \
  -e TZ=<your timezone> \
  yawn77/splotto:latest
```

## Development
For testing you have to provide a valid login for https://www.spieleplanet.eu/. If you are using vscode, you can do this by providing a `.env.test` file of the format
```
SP_USERNAME="<username>"
SP_PASSWORD="<password>"
```
In any case you can provide a username and a password as command-line arguments for the `go test` command
```bash
go test -v --race ./... -args -username="<username>" -password="<password>"
```

For execution in vscode provide a `.env` file of the following format
```
SP_USERNAME="<username>"
SP_PASSWORD="<password>"
TZ="<your timezone>"
```
You can also provide the variables above as environment variables and execute with `go run` as usual.

## License
[MIT](https://choosealicense.com/licenses/mit/)
