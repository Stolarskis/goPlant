package pgsql

const mAddData = `INSERT INTO  goplant.soilMoisture
(value, createdate, recordtime)
VALUES($1, CURRENT_DATE, $2);`

const sTAddData = `INSERT INTO  goplant.soilTemperature
(value, createdate, recordtime)
VALUES($1, CURRENT_DATE, $2);`

const aTAddData = `INSERT INTO  goplant.airTemperature
(value, createdate, recordtime)
VALUES($1, CURRENT_DATE, $2);`

const lAddData = `INSERT INTO  goplant.light
(value, createdate, recordtime)
VALUES($1, CURRENT_DATE, $2);`

const hAddData = `INSERT INTO  goplant.humidity
(value, createdate, recordtime)
VALUES($1, CURRENT_DATE, $2);`
