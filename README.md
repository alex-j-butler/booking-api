# Booking API

A Golang web API that handles booking of TF2 servers.

Note: This repository does not contain the server daemon, only the client-side API for interacting with the API.

## Dependencies

`unbuffer` is required for running the SRCDS servers.
`tar` is required for automatically downloading SteamCMD.

### Windows

#### Note: Windows is not supported.

## API

The application exposes an API over both HTTP and RPC.

### HTTP

* `GET /api/v1/servers/listall`
    
    Retrieves the list of all configured servers.

    #### Example response

    ```json
    {
        "servers": [
            {
                "uuid": "f95c0bb8-4c01-4ef4-8e1b-70ccf0aa0d19",
                "name": "Qixalite Bookable - #1",
                "tags": [
                    "bookable"
                ],
                "ip_address": "168.1.12.98",
                "port": 60015,
                "stv_port": 60016,
                "server_password": "example",
                "rcon_password": "example",
                "executable": "/home/tf2/bookable/serverfiles/srcds_run",
                "options": [
                    "-console",
                    "+map cp_badlands"
                ],
                "running": false
            }
        ]
    }
    ```

* `GET /api/v1/servers/list?uuid=<uuid>`

    Retrieves information for a single configured server, specified by the UUID.

    #### Example response

    ```json
    {
        "uuid": "f95c0bb8-4c01-4ef4-8e1b-70ccf0aa0d19",
        "name": "Qixalite Bookable - #1",
        "tags": [
            "bookable"
        ],
        "ip_address": "168.1.12.98",
        "port": 60015,
        "stv_port": 60016,
        "server_password": "example",
        "rcon_password": "example",
        "executable": "/home/tf2/bookable/serverfiles/srcds_run",
        "options": [
            "-console",
            "+map cp_badlands"
        ],
        "running": false
    }
    ```

* `GET /api/v1/servers/console?uuid=<uuid>&lines=<lines>`

    Retrieves the latest console log lines from the server specified by the UUID.

    `lines` specifies how many lines to return. **Defaults to 20 if not specified. Maximum is 100.**

* `POST /api/v1/servers/start`

    Starts the specified server.

    #### Request parameters:
    ```json
    {
        "uuid": "<uuid>"
    }
    ```

    #### Example response:

    ```json
    {
        "code": 200,
        "message": "Server started"
    }
    ```

* `POST /api/v1/servers/stop`

    Stops the specified server.

    Note that currently this endpoint will block for a few seconds while the server shuts down.

    #### Request parameters:
    ```json
    {
        "uuid": "<uuid>"
    }
    ```

    #### Example response:

    ```json
    {
        "code": 200,
        "message": "Server stopped"
    }
    ```

