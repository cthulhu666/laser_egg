FROM python:3.8-slim

RUN mkdir /app
WORKDIR /app

ADD Pipfile .
ADD Pipfile.lock .
RUN pip install --upgrade pip && \
    pip install pipenv && \
    pipenv install --system --deploy --ignore-pipfile
COPY . .

CMD python main.py
