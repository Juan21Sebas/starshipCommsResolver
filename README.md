# starshipCommsResolver
starshipCommsResolver is a microservice developed in GO with the GIN framework, its development is carried out in a hexagonal structure. This API provides a series of ENDPOINT'S to decrypt messages and triangulate the location of a ship with respect to 3 satellites of Han Solo's stellar command. 

## Detailed documentation in Swagger
In the repository there is a folder called swagger, inside there is a file in yaml format that contains the detailed documentation of the API, to access it it is recommended to use the following editor: https://editor.swagger.io/

## Receive information from satellites and decrypt
POST /topsecret/kenobi

This endpoint receives information from the satellites including the distance and the message. Triangulate the distance with respect to the location of the satellites and decipher the message.

### Request example:
json
Copy code
{
    "satellites": [
        {
            "name": "kenobi",
            "distance": 100.0,
            "message": [
                "este",
                "",
                "",
                "mensaje",
                ""
            ]
        },
        {
            "name": "skywalker",
            "distance": 115.5,
            "message": [
                "",
                "es",
                "",
                "",
                "secreto"
            ]
        },
        {
            "name": "sato",
            "distance": 142.7,
            "message": [
                "este",
                "",
                "un",
                "",
                ""
           ]
        }
    ]
}
## Get the stored information of Kenobi - Skywalker - Sato
GET /topsecret_split/kenobi

This endpoint returns previously stored information for the Kenobi satellite, including distance and message.

## Storing Kenobi - Skywalker - Sato information
POST /topsecret_split/kenobi

This endpoint allows information from the Kenobi satellite to be stored for later use in resolving the location and message of the unknown spacecraft.

### Request example:

json
Copy code
{ "distance": 100.0, "message": ["", "este", "", "mensaje", ""] }
Update stored Skywalker information

## Local execution
To run this service locally, follow these steps:

- Clone this repository to your local machine.
- Make sure you have Go installed and configured on your system.
- Run go build to compile the project.
- Run ./your-binary-name to start the server.
- Access http://lcambiar to interact with the service locally.

## Video Presentation



### Autor
#### Juan Sebastian Sanchez Arteta

