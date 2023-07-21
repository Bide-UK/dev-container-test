FROM node:14-bullseye

WORKDIR /workspace/app

COPY package*.json ./
COPY db.json ./

RUN npm install
RUN npm install -g json-server

EXPOSE 30000

CMD ["npm", "start"]
