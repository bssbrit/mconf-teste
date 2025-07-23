FROM python:3.12-slim

WORKDIR /app

RUN apt-get update

RUN apt-get install -y python3

# Install requests library
RUN pip install requests

# Copy your Python script into the container
COPY app.py ./

ENTRYPOINT ["python", "./app.py"]

CMD []
