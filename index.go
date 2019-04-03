package main

import (
  "database/sql"
  "fmt"
  "encoding/json"
  "io/ioutil"
  "net/http"
  "html/template"
  _ "github.com/lib/pq"
)

var db *sql.DB

const (
  host     = "localhost"
  port     = 5432
  user     = "yash"
  password = "101296"
  dbname   = "NYCBuildingFootprint"
)

type PublicKeyBuildP struct {
    Base_bbl string
	Bin string
	Cnstrct_yr string
	Doitt_id string
	Feat_code string
	Geomsource string
	Groundelev string
	Heightroof string
	Lstmoddate string
	Lststatype string
	Mpluto_bbl string
	The_geom struct{
		Type string
		Coordinates []float64
	}
}

type PublicKeyBuild struct {
    Base_bbl string
	Bin string
	Cnstrct_yr string
	Doitt_id string
	Feat_code string
	Geomsource string
	Groundelev string
	Heightroof string
	Lstmoddate string
	Lststatype string
	Mpluto_bbl string
	Shape_area string
	Shape_len string
	
}

func index_handler(w http.ResponseWriter, r *http.Request) {


	/*
		'/'=> default page load
		'/extract' => extract data from go api using javascript: 127.0.0.1:8000/extract
		'/analyze' => Make API's to getch data from db and visualize
	*/
	t, _ := template.ParseFiles("index.html")
    t.Execute(w,nil)

}

func last_modi_handler(w http.ResponseWriter, r *http.Request){
	
	var data PublicKeyBuild
	isAvg, ok := r.URL.Query()["isheightavg"]
	selectSql := ``
    if !ok || len(isAvg[0]) < 1 {
		selectSql = `SELECT * FROM buildingfootprint`
	} else {
		selectSql = `SELECT AVG(heightroof) FROM buildingfootprint`
	}
	flag := false
	
	constYears, ok := r.URL.Query()["constyear"]
    if !ok || len(constYears[0]) < 1 {
        //log.Println("Url Param 'key' is missing")
	} else {
		if(flag){
			selectSql = selectSql + ` AND cnstrct_yr > `+constYears[0]+``
			flag = true;
		} else {
			selectSql = selectSql + ` WHERE cnstrct_yr > `+constYears[0]+``
			flag = true;
		}
	}
	featCode, ok := r.URL.Query()["featcode"]
    if !ok || len(featCode[0]) < 1 {
        //log.Println("Url Param 'key' is missing")
	} else {
		if(flag){
			selectSql = selectSql + ` AND feat_code = `+featCode[0]+``
			flag = true;
		} else {
			selectSql = selectSql + ` WHERE feat_code = `+featCode[0]+``
			flag = true;
		}
	}
	isavg, ok := r.URL.Query()["isheightavg"]
    if !ok || len(isavg[0]) < 1 {
		fmt.Println(selectSql)
		row := db.QueryRow(selectSql)
		switch err := row.Scan(&data.Base_bbl, &data.Bin, &data.Cnstrct_yr, &data.Doitt_id, &data.Feat_code, &data.Geomsource, &data.Groundelev, &data.Heightroof, &data.Lstmoddate, &data.Lststatype, &data.Mpluto_bbl, &data.Shape_area, &data.Shape_len); err {
		case sql.ErrNoRows:
		  fmt.Println("No rows were returned!")
		case nil:
		  fmt.Println(data)
		default:
		  panic(err)
		}
		
	} else{
		fmt.Println(selectSql)
		row := db.QueryRow(selectSql)
		selectSql = selectSql + ` GROUP BY heightroof`
		switch err := row.Scan(&data.Heightroof); err {
		case sql.ErrNoRows:
		  fmt.Println("No rows were returned!")
		case nil:
		  fmt.Println(data)
		default:
		  panic(err)
		}
	}
	
	b, err := json.Marshal(data)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Fprintln(w,string(b))

	
	

	



	
}


func extract_handler(w http.ResponseWriter, r *http.Request){
	createSql := `CREATE TABLE buildingfootprint (base_bbl bigserial, bin bigserial, cnstrct_yr INT, doitt_id bigserial, feat_code bigserial, geomsource TEXT, groundelev INT, heightroof FLOAT(3), lstmoddate DATE, lststatype TEXT, mpluto_bbl bigserial, shape_area FLOAT(4), shape_len FLOAT(4));`
	var err error
	_, err = db.Exec(createSql)
	if err != nil {
		//panic(err)
		fmt.Fprintf(w,"table exists")

	} else {
		/* Building API */
		build_resp, _ := http.Get("https://data.cityofnewyork.us/resource/mtik-6c5q.json")
		bytes_build, _ := ioutil.ReadAll(build_resp.Body)
		keys_build := make([]PublicKeyBuild,0)
		json.Unmarshal(bytes_build, &keys_build)

		/* Building_p API */
		
		build_p_resp, _ := http.Get("https://data.cityofnewyork.us/resource/9ey5-eyh6.json")
		bytes_build_p, _ := ioutil.ReadAll(build_p_resp.Body)
		keys_build_p := make([]PublicKeyBuildP,0)
		json.Unmarshal(bytes_build_p, &keys_build_p)

		for i:=0; i<len(keys_build); i++{
			insertsql := `INSERT INTO Buildingfootprint(base_bbl, bin, cnstrct_yr, doitt_id, feat_code, geomsource, groundelev, heightroof, lstmoddate, lststatype, mpluto_bbl, shape_area, shape_len) values 
			(`+keys_build[i].Base_bbl+`,`+keys_build[i].Bin+`,`+keys_build[i].Cnstrct_yr+`,`+keys_build[i].Doitt_id+`,`+keys_build[i].Feat_code+`,'`+keys_build[i].Geomsource+`',`+keys_build[i].Groundelev+`,`+keys_build[i].Heightroof+`,'`+keys_build[i].Lstmoddate+`','`+keys_build[i].Lststatype+`',`+keys_build[i].Mpluto_bbl+`,`+keys_build[i].Shape_area+`,`+keys_build[i].Shape_len+`);`
			fmt.Println(insertsql)
			_, err = db.Exec(insertsql)
			if err != nil {
				panic(err)
			}
		}
	}


}
func initDb() {
	
    var err error
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

    db, err = sql.Open("postgres", psqlInfo)
    if err != nil {
        panic(err)
	}
	err = db.Ping()
    if err != nil {
        panic(err)
    }
    fmt.Println("Successfully connected!")

}

func main() {

	initDb()
	defer db.Close()
	fmt.Println("Successfully connected!")
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/extract", extract_handler)
	http.HandleFunc("/filter", last_modi_handler)
	
	http.ListenAndServe(":8000", nil) 

	

}