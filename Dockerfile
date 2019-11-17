FROM python:3.8

RUN mkdir /app
WORKDIR /app
ADD . .
RUN pip install -r requirements.txt

CMD python main.py
