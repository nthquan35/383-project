FROM node:14-buster

#RUN npm install --build-from-source zeromq@6.0.0-beta.5 

WORKDIR /app
COPY ./package.json ./
RUN npm install
COPY . .

CMD ["npm", "run", "start"]
