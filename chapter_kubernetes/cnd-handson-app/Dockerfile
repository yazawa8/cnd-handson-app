FROM python:alpine

ADD ./requirements.txt /opt/cnd-web-app/

WORKDIR /opt/cnd-web-app

RUN pip install -r requirements.txt

ADD . /opt/cnd-web-app

EXPOSE 8888

CMD python /opt/cnd-web-app/app.py