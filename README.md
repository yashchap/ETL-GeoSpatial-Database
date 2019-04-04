# GO_Internship_Assignment

### Setup Postgres SQL:
  Create database NYCBuildingFootprint

### Running the project: 
1) clone respository
2) open terminal from the folder and run command: go run index.go
3) open url : 127.0.0.1:8000

### About project:

Developed an API to ETL the NYC Open data of building footprint. <br/>
API endpoint 1: https://data.cityofnewyork.us/resource/9ey5-eyh6.json <br/>
API endpoint 2: https://data.cityofnewyork.us/resource/mtik-6c5q.json

### Implementation Details:

1) When Page loads, it will call "index_handler" function of Go which will load html tempate.
2) Ajax call to /extract url will run "extract_handler" function of Go which will fetch data from API endpoint and store it in PostgreSQL.
3) Whenever filter is applied to fetched data, "last_modi_handler" function will get executed which will filter the data from database and send the data to html template.

Developed an API to extract data from Postgres buildingfootprint table generated fromn above ETL setup and applied following filters:
1) Filter by Construction date.
2) Filter by FeetCode.
3) Select whether you want to avg heightroof or not.

### APIs build:
1) To extract and load API data to PostgreSQL.
2) To filter data and fetch data from database.




![Alt text](https://github.com/yashchap/GO_Internship_Assignment/blob/master/s3.png)
![Alt text](https://github.com/yashchap/GO_Internship_Assignment/blob/master/s2.png)
![Alt text](https://github.com/yashchap/GO_Internship_Assignment/blob/master/s1.png)

