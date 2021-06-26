package pgsql

const mAddData = `INSERT INTO  goplant.soilmoisture
(value, createdate, recordtime)
VALUES($1, CURRENT_DATE, $2);`

const sTAddData = `INSERT INTO  goplant.soiltemperature
(value, createdate, recordtime)
VALUES($1, CURRENT_DATE, $2);`

const aTAddData = `INSERT INTO  goplant.airtemperature
(value, createdate, recordtime)
VALUES($1, CURRENT_DATE, $2);`

const lAddData = `INSERT INTO  goplant.light
(value, createdate, recordtime)
VALUES($1, CURRENT_DATE, $2);`

const hAddData = `INSERT INTO  goplant.humidity
(value, createdate, recordtime)
VALUES($1, CURRENT_DATE, $2);`

const mGetLastVal = `SELECT t.value
FROM goplant.soilmoisture t
INNER JOIN (
    SELECT id, recordtime, max(recordtime) as MaxDate
    FROM goplant.soilmoisture
    GROUP BY id,recordtime
) tm on t.id = tm.id and t.recordtime = tm.MaxDate`

const stGetLastVal = `SELECT t.value
FROM goplant.soilTemperature t
INNER JOIN (
    SELECT id, recordtime, max(recordtime) as MaxDate
    FROM goplant.soiltemperature
    GROUP BY id,recordtime
) tm on t.id = tm.id and t.recordtime = tm.MaxDate`

const atGetLastVal = `SELECT t.value
FROM goplant.airTemperature t
INNER JOIN (
    SELECT id, recordtime, max(recordtime) as MaxDate
    FROM goplant.airtemperature
    GROUP BY id,recordtime
) tm on t.id = tm.id and t.recordtime = tm.MaxDate`

const lGetLastVal = `SELECT t.value
FROM goplant.light t
INNER JOIN (
    SELECT id, recordtime, max(recordtime) as MaxDate
    FROM goplant.light
    GROUP BY id,recordtime
) tm on t.id = tm.id and t.recordtime = tm.MaxDate`

const hGetLastVal = `SELECT t.value
FROM goplant.humidity t
INNER JOIN (
    SELECT id, recordtime, max(recordtime) as MaxDate
    FROM goplant.humidity
    GROUP BY id,recordtime
) tm on t.id = tm.id and t.recordtime = tm.MaxDate`
