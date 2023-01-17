export default class Driver {

    position;
    points;
    wins;
    id;
    number;
    lastName;
    nationality;
    totalrounds;
    team;

    constructor(position, points, wins, id, number, lastName, nationality, totalrounds, team) {
        this.position = position;
        this.points = points;
        this.wins = wins;
        this.id = id;
        this.number = number;
        this.lastName = lastName;
        this.nationality = nationality;
        this.totalrounds = totalrounds;
        this.team = team;
    }


    static retrieveDrivers(year, round) {

        var drivers = [];

        var xmlHttp = new XMLHttpRequest();

        var yearString = year?.toString() || "2022"
        var roundString = round?.toString() || "22"

        xmlHttp.open( "GET", "http://3.252.163.138:3333/standings" + yearString + "-" + roundString, false ); // false for synchronous request
        xmlHttp.send( null );

        const obj = JSON.parse(xmlHttp.responseText);

        console.log(obj)

        obj.forEach(e => {
            const newDriver = new Driver(e.PositionText, e.Points, e.Wins, e.Driver.DriverId, e.Driver.PermanentNumber, e.Driver.FamilyName, e.Driver.Nationality, parseInt(e.TotalRounds), e.Constructors[0].Name)
            console.log("TEAM: " + newDriver.team)
            drivers.push(newDriver);
        });

        return drivers;
    }
}
