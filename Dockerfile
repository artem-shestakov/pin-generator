FROM python
WORKDIR /app
ADD . /app
RUN pip install -r requirements.txt
EXPOSE 5000
ENV WEBAPP Prod
ENTRYPOINT ["python3", "gserver.py"]