FROM ubuntu:latest
LABEL authors="hnaderi268"

RUN apt-get update && apt-get install -y \
    python3.9 \
    python3-pip

ENV PATH="$PATH:/usr/local/bin/python:/usr/local/bin/python3"

COPY requirements.txt /app/requirements.txt
WORKDIR /app

RUN pip install -r requirements.txt
COPY . /app

EXPOSE 50050

ENTRYPOINT [ "python3" , "main.py" ]
