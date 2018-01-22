# Setting Up Your Environment
Keep in mind that gRPC requires golang v1.6 or greater, to install run

```
go get -u google.golang.org/grpc
```

and also run

```
go get -u github.com/golang/protobuf/protoc-gen-go
```

Finally install latest version of protobuf from release page [https://github.com/google/protobuf/releases](https://github.com/google/protobuf/releases) which
contains the `protoc` binary that must be copied/moved to your `PATH`, e.g., `/usr/local/bin`. 



# Generating Golang API
To build the api (generate go code from proto file) you will want to run a command like
```
protoc -I api/ api/hello.proto --go_out=plugins=grpc:api
```

where in general you would run

```
protoc -I <input_directory> <path_to_file> --go_out=plugins=grpc:<output_directory>
```

# Generated Documentation
We can follow along [https://github.com/pseudomuto/protoc-gen-doc](https://github.com/pseudomuto/protoc-gen-doc). You can
follow along the docker commands (recommended) or run locally (assuming you have golang set up properly)
via 

```
go get -u github.com/pseudomuto/protoc-gen-doc/cmd/...
```

and, if you did the above, you can simply run

```
protoc --doc_out docs --doc_opt=markdown,api.md  api/hello.proto
```

where what the above comes from is

```
protoc --doc_out <output_folder> --doc_opt=<format, file_output> <api_proto_input_files>
```
# REST Gateway

* generate gRPC go stub
```
protoc -I api/ api/hello.proto -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway --go_out=plugins=grpc:api
```

* generate gateway got stub
```
protoc -I /usr/local/include/ -I api/ api/hello.proto -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway --grpc-gateway_out=logtostderr=true:api
```

* swagger output
```
protoc -I /usr/local/include/ -I api/ api/hello.proto -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway --swagger_out=logtostderr=true:swagger
```

## Example API calls

* get swagger docs
```
curl http://localhost:8502/swagger/hello.swagger.json

{
  "swagger": "2.0",
  "info": {
    "title": "hello.proto",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/v1/door": {
      "post": {
        "summary": "*\nDetermines if you knocked the door and the appropriate\nreply if you didnt.",
        "operationId": "GetHello",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/helloReply"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/helloKnock"
            }
          }
        ],
        "tags": [
          "HelloService"
        ]
      }
    }
  },
  "definitions": {
    "helloKnock": {
      "type": "object",
      "properties": {
        "knockDoor": {
          "type": "boolean",
          "format": "boolean"
        }
      },
      "title": "*\nKnock the door or not",
      "externalDocs": {
        "description": "Find out more about this grpc example",
        "url": "https://github.com/alejandroEsc/grpc-example"
      }
    },
    "helloReply": {
      "type": "object",
      "properties": {
        "reply": {
          "type": "boolean",
          "format": "boolean"
        },
        "replyMessage": {
          "type": "string"
        }
      },
      "description": "*\nReply based on whether a door was knocked or not."
    }
  }
}
```

* knock the door
```
curl -X POST http://localhost:8502/v1/door/knock/true

{"reply":true,"replyMessage":"Hello!"}
```

* do not knock the door
```
curl -X POST http://localhost:8502/v1/door/knock/false
{"replyMessage":"You should try and knock"}
```

* failure case

```
 curl -X POST http://localhost:8502/v1/door/knock/
{"error":"type mismatch, parameter: knockDoor, error: strconv.ParseBool: parsing \"\": invalid syntax","code":3}
```