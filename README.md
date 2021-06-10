 No newline at end of file
# goPlant

[![Stolarskis](https://circleci.com/gh/Stolarskis/goPlant.svg?style=shield)](https://app.circleci.com/pipelines/github/Stolarskis)

Api component for storing plant sensor data.

Endpoints for api represent each sensor type.

#### Endpoints

 
- `/data/moisture`
- `/data/soilTemp`
- `/data/airTemp`
- `/data/humidity`
- `/data/light`

Api currently runs in a docker container, running on a raspberry pi. In the process of setting up api and database on local kubernetes cluster!

## Inspiration for GoPlant 
---
There are many arduino based auto-planters out there. While following a guide, I realized that cable management for all these sensors was going to be a real issue. And what if I wanted to monitor multiple "gardens" that were far apart from each other. The easiest solution to this would be to store the data wirelessly and then perform actions based on that data.

## Main Components of Plant Monitoring System
---
- GoPlant Api - Stores data sent by sensor components to a sql database
- Sensor Components - List of parts used below - Takes a reading from a sensor and makes a upload request to the goPlant api
- Action Components - (Very much a work in progess) A similar setup to a sensor, but receives requests and acts on them accordingly. An example would be a module that distributes water based on the values in body of request.
---
## Sensor Component Parts!
---
### Sensor Types and Sensors Used
- Light Intensity - PhotoResistor (Not setup)
- Soil Moisture - [Vegetronix VB400](https://www.vegetronix.com/Products/VH400/)
- Soil Temperature - [Vegetronix THERM200](https://vegetronix.com/Products/THERM200/)
- Air Humidity - [DHT11](https://vegetronix.com/Products/THERM200/)
- Air Temperature - [DHT11](https://vegetronix.com/Products/THERM200/)
---
### Develoment Boards
[ESP8266](https://www.amazon.com/Organizer-ESP8266-Internet-Development-Compatible/dp/B081PX9YFV)

Tiny little board that is able to read sensors and make requests to api via built in wifi module. Its very cool.
