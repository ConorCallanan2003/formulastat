# formulastat

## This is a webapp built using GoLang and Vue.js for looking through historical F1 statistics.

The backend (written in Go) periodically calls an API with historical F1 data to fill the SQL database that stores the data. This has to be done as the API (http://ergast.com/mrd/) only allows for up to 200 requests per hour.

Using Go's excellent and easy to use implementation of multithreading, this data caching process can run asynchronously while the same Go program can serve data to the front end through HTTP requests.

The Vue.js Frontend is kept as simple and intuitive as possible and uses the Bulma CSS framework for UI components.
