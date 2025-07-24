FROM python:3.12-slim

WORKDIR /app

RUN apt-get update

RUN apt-get install -y python3


RUN pip install requests


COPY . ./

ENTRYPOINT ["python", "./app.py"]

CMD []
