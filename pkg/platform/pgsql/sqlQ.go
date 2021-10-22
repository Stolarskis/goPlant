package pgsql

const mAddData = `INSERT INTO  goplant.soilmoisture
(value, recordtime)
VALUES($1, $2);`

const sTAddData = `INSERT INTO  goplant.soiltemperature
(value, recordtime)
VALUES($1, $2);`

const aTAddData = `INSERT INTO  goplant.airtemperature
(value, recordtime)
VALUES($1, $2);`

const lAddData = `INSERT INTO  goplant.light
(value, recordtime)
VALUES($1, $2);`

const hAddData = `INSERT INTO  goplant.humidity
(value, recordtime)
VALUES($1, $2);`

const mGetLastVal = `SELECT value 
FROM goplant.soilmoisture
WHERE recordtime = (SELECT MAX(recordtime) FROM goplant.soilmoisture)`

const stGetLastVal = `SELECT value
FROM goplant.soiltemperature
WHERE recordtime = (SELECT MAX(recordtime) FROM goplant.soiltemperature)`

const atGetLastVal = `SELECT value
FROM goplant.airtemperature
WHERE recordtime = (SELECT MAX(recordtime) FROM goplant.airtemperature)`

const lGetLastVal = `SELECT value
FROM goplant.light
WHERE recordtime = (SELECT MAX(recordtime) FROM goplant.light)`

const hGetLastVal = `SELECT value
FROM goplant.humidity
WHERE recordtime = (SELECT MAX(recordtime) FROM goplant.humidity)`
