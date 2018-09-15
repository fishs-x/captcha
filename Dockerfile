FROM pujielan/golang:1.10.3

ADD ${PWD}/build/output/1.0.0/insur-box-imagecode /usr/local/bin/insur-box-imagecode
RUN chmod +x /usr/local/bin/insur-box-imagecode

EXPOSE "8080"

CMD ["insur-box-imagecode"]