FROM hseeberger/scala-sbt:graalvm-ce-21.3.0-java17_1.6.2_3.1.1

WORKDIR /app
COPY . .

RUN sbt clean compile

EXPOSE 9000
CMD ["sbt", "run"]
