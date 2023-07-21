FROM python:3.9-bullseye

WORKDIR /workspace/app

COPY app.py .

RUN pip install flask

EXPOSE 50000

CMD ["python", "./app.py","po"]
