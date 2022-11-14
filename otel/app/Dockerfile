# syntax=docker/dockerfile:1

FROM python:3.8

WORKDIR /app

RUN pip3 install opentelemetry-distro opentelemetry-exporter-otlp flask

RUN opentelemetry-bootstrap --action=install

COPY ./app/server.py .

CMD opentelemetry-instrument \
    python \
    server.py