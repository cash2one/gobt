# gobt
A BitTorrent Service based on https://github.com/shiyanhui/dht and https://github.com/btlike/repository .

## Usage
### Installation
```
$ go get -u github.com/xgfone/gobt
$ go install github.com/xgfone/gobt
```

### Deployment
1. Create the mysql database.
2. Import the table into the database. See `mysql.sql`.
3. Download and deploy the elasticsearch. See http://www.elasticsearch.org/.
3. Modify the configuration file.
4. Start the program.

### Run
```
$ gobt /PATH/TO/bt.conf
```

### Notice
1. The logfile is eithor an absolute or a relative path.
2. The level of the logger is one of debug, info, warn, error, crit.
3. It uses GORM as the DB ORM, so for the connection of DB, see [GORM](https://github.com/jinzhu/gorm).
