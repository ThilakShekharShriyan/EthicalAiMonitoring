FROM python:3.9-alpine
WORKDIR /app
COPY . /app
RUN ls -l .
RUN ls -l ./app
RUN pip install -r requirements.txt
CMD ["python", "main.py"]
