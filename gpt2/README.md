gpt2 is a service which computes generated text with hugging face model. It is used by jarvis as a backend. You can see its api in protos folder.

You can create container with ```docker build . -t gpt2 ``` and then run it with ```docker run -p 50050:50050 gpt2 ```. ALso there is docker compose in jarvis, which you can use to tun both services.