
Please do **go mod tidy** and enable deep integration in your ide\

commands to run the application
1. docker build -t temp .
2. docker-compose up (if you wants to run in background, dont want to see the logs then you can run docker-compose up -d)

hopefully everything works find at this end but i was not able to test it out using docker because my system is not supporting
docker, everytime when i run docker getting windows is not responding. system starting rebooting

if that not works just use commands to build and run the docker and run commands mentioned in schema.sql in docker container terminal everything will start working perfactly fine.


Curl to hit the api:

curl --request GET \
--url http://localhost:8000/private/v1/fetch/records \
--header 'Authorization: Basic bWV6X2luazpIZWxsbyFAI01lekluaw==' \
--header 'Content-Type: application/json' \
--data '{
"startDate": "2023-05-19",
"endDate": "2023-05-20",
"minCount": 1,
"maxCount": 300
}	'

