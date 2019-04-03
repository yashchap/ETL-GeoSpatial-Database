# GO_Internship_Assignment

## Setup Postgres SQL:
  Create database NYCBuildingFootprint

## Running the project: 
1) clone respository
2) open terminal from the folder and run command: go run index.go
3) open url : 127.0.0.1:8000

### About project:

Developed an API to ETL the NYC Open data of building footprint. <br/>
API endpoint 1: https://data.cityofnewyork.us/resource/9ey5-eyh6.json <br/>
API endpoint 2: https://data.cityofnewyork.us/resource/mtik-6c5q.json

Developed an API to extract data from Postgres buildingfootprint table generated fromn above ETL setup and applied following filters:
1) Filter by Construction date.
2) Filter by FeetCode.
3) Select whether you want to avg heightroof or not.
