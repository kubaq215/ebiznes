# Use Ubuntu 22.04 as the base image
FROM ubuntu:22.04

ENV TZ=Europe/Warsaw
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# Install dependencies
RUN apt-get update \
    && apt-get install -y \
    wget \
    curl \
    unzip \
    zip \
    openjdk-8-jdk \
    software-properties-common

# Install Python 3.8
RUN add-apt-repository ppa:deadsnakes/ppa -y \
    && apt-get update \
    && apt-get install -y python3.8 \
    && rm -rf /var/lib/apt/lists/*

# Install Kotlin and Gradle
RUN curl -s https://get.sdkman.io | bash
SHELL ["/bin/bash", "-c"]  
RUN source "$HOME/.sdkman/bin/sdkman-init.sh" \
    && sdk install kotlin \
    && sdk install gradle \
    && sdk flush archives

# Install ktlint
RUN curl -sSLO https://github.com/pinterest/ktlint/releases/download/0.49.0/ktlint \
&& chmod a+x ktlint \
&& mv ktlint /usr/local/bin/

RUN apt purge python3 -y \
    && rm -rf /usr/bin/python3 \
    && rm -rf /usr/bin/python3.10 \
    && ln -s /usr/bin/python3.8 /usr/bin/python3

RUN rm -rf /usr/share/icons

# Set environment variables
ENV PATH="/root/.sdkman/candidates/kotlin/current/bin:/root/.sdkman/candidates/gradle/current/bin:$PATH"
ENV JAVA_HOME=/usr/lib/jvm/java-8-openjdk-amd64
ENV PATH=$PATH:$JAVA_HOME/bin

# Set working directory
WORKDIR /app

# Copy project files
COPY demo .
COPY db ./db/

# Build the project (if necessary)
RUN gradle build

# Define default command to run when the container starts
CMD ["gradle", "run"]