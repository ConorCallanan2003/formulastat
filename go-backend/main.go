package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	// "os"

	_ "github.com/lib/pq"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	fmt.Printf("got / request\n")
	io.WriteString(w, "Invalid Request\n")
}

func main() {

	go generateSqlData()

	http.HandleFunc("/", getRoot)

	//find way of passing data to handleStandingsRequest function...

	for i := 1950; i < 2023; i++ {
		for x := 0; x < 24; x++ {
			year := i
			round := x
			http.HandleFunc("/standings"+strconv.Itoa(i)+"-"+strconv.Itoa(x), func(w http.ResponseWriter, r *http.Request) {
				data := string(mapToJSON(RetrieveSeason(year, round)))
				enableCors(&w)
				fmt.Print("got /" + "standings" + strconv.Itoa(year) + "-" + strconv.Itoa(round) + " request\n")
				io.WriteString(w, data)
			})
		}
	}

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}

}

func generateSqlData() {
	for i := 2022; i > 1949; i-- {

		resp, err := http.Get("http://ergast.com/api/f1/" + strconv.FormatInt(int64(i), 10) + "/driverStandings.json?limit=100")

		if err != nil {
			log.Panic(err)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Panic(err)
		}

		var response Response
		json.Unmarshal([]byte(body), &response)

		numOfRounds, err := strconv.Atoi(response.MRData.StandingsTable.StandingsLists[0].Round)

		if err != nil {
			log.Panic(err)
		}

		if !CheckRoundExists(i, numOfRounds) {
			addSeasonToDB(response.MRData.StandingsTable, numOfRounds)
			fmt.Print(strconv.FormatInt(int64(i), 10) + " Round " + strconv.FormatInt(int64(numOfRounds), 10) + " Added to Database\t")
		} else {
			fmt.Print(strconv.FormatInt(int64(i), 10) + " Round " + strconv.FormatInt(int64(numOfRounds), 10) + " Already Completed\t")
		}

		time.Sleep(30 * time.Second)

		for x := 1; x < numOfRounds; x++ {

			if !CheckRoundExists(i, x) {
				roundWSlash := "/" + strconv.FormatInt(int64(x), 10)
				resp, err := http.Get("http://ergast.com/api/f1/" + strconv.FormatInt(int64(i), 10) + roundWSlash + "/driverStandings.json?limit=100")

				if err != nil {
					log.Panic(err)
				}

				body, err := io.ReadAll(resp.Body)
				if err != nil {
					log.Panic(err)
				}

				var response Response
				json.Unmarshal([]byte(body), &response)

				if response.MRData.Total != "0" {
					addSeasonToDB(response.MRData.StandingsTable, numOfRounds)
					fmt.Print(strconv.FormatInt(int64(i), 10) + " Round " + strconv.FormatInt(int64(x), 10) + " Added to Database\t")
				} else {
					x = 100
				}

				time.Sleep(30 * time.Second)
			} else {
				fmt.Print(strconv.FormatInt(int64(i), 10) + " Round " + strconv.FormatInt(int64(x), 10) + " Already Completed\t")
			}
		}
	}
}

func mapToJSON(drivers []DriverStanding) []byte {

	resultJSON, err := json.Marshal(drivers)
	if err != nil {
		panic(err)
	}

	return resultJSON
}

// func apiCaller() []DriverStanding {

// 	var result []DriverStanding

// 	resp, err := http.Get("http://ergast.com/api/f1/current/driverStandings.json")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	var response Response
// 	json.Unmarshal([]byte(body), &response)

// 	addSeasonToDB(response.MRData.StandingsTable)

// 	var lists = response.MRData.StandingsTable.StandingsLists

// 	for i := 0; i < len(lists); i++ {

// 		result = lists[i].DriverStandings

// 	}

// 	return result

// }

const (
	host     = "db"
	port     = 5433
	user     = "user"
	password = "password"
	dbname   = "cars"
)

func addSeasonToDB(standingsTable StandingsTable, totalRounds int) {

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)

	CheckError(err)

	//defer delays the execution of a function until the surrounding function returns
	defer db.Close()

	insertStmt := `CREATE TABLE IF NOT EXISTS "` + standingsTable.Season + `-` + standingsTable.StandingsLists[0].Round + `"(
		id SERIAL PRIMARY KEY,
		"name" VARCHAR( 50 ),
		"position" int,
		points REAL,
		wins int,
		number VARCHAR( 3 ),
		drivercode VARCHAR( 4 ),
		nationality VARCHAR ( 80 ),
		constructorid VARCHAR ( 50 ),
		constructorname VARCHAR ( 80 ),
		totalrounds int
		);`
	_, errr := db.Exec(insertStmt)
	CheckError(errr)

	driverStandings := standingsTable.StandingsLists[0].DriverStandings

	for i := 0; i < len(driverStandings); i++ {

		driver := driverStandings[i]

		insertDynStmt := `insert into "` + standingsTable.Season + "-" + standingsTable.StandingsLists[0].Round + `"("name", "position", points, wins, number, drivercode, nationality, constructorid, constructorname, totalrounds) values('` + driver.Driver.FamilyName + `', '` + driver.PositionText + `', '` + driver.Points + `', '` + driver.Wins + `', '` + driver.Driver.PermanentNumber + `', '` + driver.Driver.Code + `', '` + driver.Driver.Nationality + `', '` + driver.Constructors[0].ConstructorId + `', '` + driver.Constructors[0].Name + `', '` + strconv.Itoa(totalRounds) + `')`
		_, err = db.Exec(insertDynStmt)
		CheckError(err)
	}

}

func RetrieveSeason(season int, round int) []DriverStanding {

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)

	CheckError(err)

	//defer delays the execution of a function until the surrounding function returns
	defer db.Close()

	insertStmt := `SELECT * FROM "` + strconv.FormatInt(int64(season), 10) + `-` + strconv.FormatInt(int64(round), 10) + `"`

	rows, err := db.Query(insertStmt)

	if err != nil {
		log.Print(err)
	}

	defer rows.Close()

	//parse json

	drivers := []DriverStanding{}

	for rows.Next() {
		driver := DriverStanding{}
		id := 0
		driver.Constructors = []Constructor{}
		driver.Constructors = append(driver.Constructors, Constructor{})
		rows.Scan(&id, &driver.Driver.FamilyName, &driver.PositionText, &driver.Points, &driver.Wins, &driver.Driver.PermanentNumber, &driver.Driver.Code, &driver.Driver.Nationality, &driver.Constructors[0].ConstructorId, &driver.Constructors[0].Name, &driver.TotalRounds)
		drivers = append(drivers, driver)
	}

	return drivers
}

func CheckRoundExists(season int, round int) bool {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)

	CheckError(err)

	//defer delays the execution of a function until the surrounding function returns
	defer db.Close()

	insertStmt := `SELECT EXISTS (
		SELECT FROM 
			pg_tables
		WHERE 
			schemaname = 'public' AND 
			tablename  = '` + strconv.FormatInt(int64(season), 10) + `-` + strconv.FormatInt(int64(round), 10) + `'
		);`
	rows, errr := db.Query(insertStmt)

	var res string

	defer rows.Close()

	result := true

	for rows.Next() {
		rows.Scan(&res)
		if res == "false" {
			result = false
		}
	}

	CheckError(errr)

	return result
}

func CheckError(err error) {
	if err != nil {
		fmt.Print(err)
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

type Response struct {
	MRData MRData
}

type MRData struct {
	StandingsTable StandingsTable
	Total          string
}

type StandingsTable struct {
	StandingsLists []StandingsList
	Season         string
}

type StandingsList struct {
	DriverStandings []DriverStanding
	Season          string
	Round           string
}

type DriverStanding struct {
	PositionText string
	Points       string
	Wins         string
	Driver       Driver
	Constructors []Constructor
	TotalRounds  string
}

type Driver struct {
	DriverId        string
	PermanentNumber string
	Code            string
	Url             string
	DateOfBirth     string
	FamilyName      string
	Nationality     string
}

type Constructor struct {
	ConstructorId string
	Url           string
	Name          string
	Nationality   string
}
