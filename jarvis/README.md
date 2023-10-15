Jarvis is a go bff micro-service used to get gpt2 generated texts.
I chose golang because of its simplicity and scalabilty.
Jarvis calls gpt2 a python module via gRPC to get its results.

Compile and run locally:
Set gpt2 host and port  properly. To run jarvis first compile gpt2.proto by ```./proto_gen.sh```
then ``` go run .```.

Build docker and run:
To create docker image ```docker build . -t IMAGE_MAME``` and run the container ```docker run -p 8080:8080 IMAGE_NAME```

Deployment:
You can deploy both images into some pods using kubernetes. Here for example I have created a ```compose.yaml``` whcih you can run with ```docker compose up```.


# Points of improvement


Add tracing id for user logs.


Add tls/ssl for production use.


Create a session for each user and keep history of their requests.


gpt2 model properties are hardcoded.

I did not do the optional db storing part, in a production env I would use a persistent volume to create and maintain a database or use some database as a service..
