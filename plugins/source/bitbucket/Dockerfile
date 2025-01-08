FROM --platform=$BUILDPLATFORM gradle:8.4-jdk20 as build
ARG GITHUB_ACTOR
ARG GITHUB_TOKEN

WORKDIR /code

COPY . .

RUN gradle jar --no-daemon

FROM eclipse-temurin:21.0.5_11-jre

COPY --from=build /code/app/build/libs/*.jar /app/app.jar

EXPOSE 7777

ENV _JAVA_OPTIONS="--add-opens=java.base/java.nio=ALL-UNNAMED"

ENTRYPOINT ["java", "-jar", "/app/app.jar"]

CMD [ "serve", "--address", "localhost:7777", "--log-format", "json", "--log-level", "info" ]
