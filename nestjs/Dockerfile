FROM node:18.12
# debian

RUN apt update -y && \
    apt install -y procps && \
    npm install -g @nestjs/cli@10.0.0

WORKDIR /home/node/app

USER node

# Read null linux device so don't let the container die
CMD [ "tail", "-f", "/dev/null"]