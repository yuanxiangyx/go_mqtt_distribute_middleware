FROM mqtt_brige

WORKDIR /opt/app
RUN rm -rf *
COPY . .

CMD ["sh", "run.sh"]
