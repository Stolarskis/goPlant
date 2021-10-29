#include <ArduinoJson.h>
#include <ESP8266WiFi.h>
#include <ESP8266mDNS.h>
#include <WiFiUdp.h>
#include <ArduinoOTA.h>

// This script is used as the main template for all analog sensor firmware.
// The only difference between them is what the endpoint is set to.

#define WIFI_ID ""
#define WIFI_PASSWORD ""
#define HOST_IP ""
#define HOST_PORT
#define ENDPOINT_LIGHT "/data/light"
#define ENDPOINT_SOIL_MOISTURE "/data/moisture"
#define ENDPOINT_SOIL_TEMP "/data/soilTemp"

const char *ssid = WIFI_ID;
const char *password = WIFI_PASSWORD;
const char *host = HOST_IP;
const int port = HOST_PORT;
const char *url = Enter_Endpoint_here;

void setup()
{
  Serial.begin(115200);
  Serial.println("Booting");
  WiFi.mode(WIFI_STA);
  WiFi.begin(ssid, password);
  while (WiFi.waitForConnectResult() != WL_CONNECTED)
  {
    Serial.println("Connection Failed! Rebooting...");
    delay(5000);
    ESP.restart();
  }

  ArduinoOTA.onStart([]()
  {
    String type;
    if (ArduinoOTA.getCommand() == U_FLASH)
    {
      type = "sketch";
    }
    else
    { // U_FS
      type = "filesystem";
    }

    // NOTE: if updating FS this would be the place to unmount FS using FS.end()
    Serial.println("Start updating " + type);
  });
  ArduinoOTA.onEnd([]()
  {
    Serial.println("\nEnd");
  });
  ArduinoOTA.onProgress([](unsigned int progress, unsigned int total)
  {
    Serial.printf("Progress: %u%%\r", (progress / (total / 100)));
  });
  ArduinoOTA.onError([](ota_error_t error)
  {
    Serial.printf("Error[%u]: ", error);
    if (error == OTA_AUTH_ERROR)
    {
      Serial.println("Auth Failed");
    }
    else if (error == OTA_BEGIN_ERROR)
    {
      Serial.println("Begin Failed");
    }
    else if (error == OTA_CONNECT_ERROR)
    {
      Serial.println("Connect Failed");
    }
    else if (error == OTA_RECEIVE_ERROR)
    {
      Serial.println("Receive Failed");
    }
    else if (error == OTA_END_ERROR)
    {
      Serial.println("End Failed");
    }
  });
  ArduinoOTA.begin();
  Serial.println("Ready");
  Serial.print("IP address: ");
  Serial.println(WiFi.localIP());
}

/**
   Main Firmware Loop

   1. Checks for Over The Air (OTA) firmware updates
   2. Attempts to connect to Wifi and if successful, every 10s,
      grabs latest sensor reading and attempts POST request
      to GoPlant Server.
*/
void loop()
{
  ArduinoOTA.handle();

  //Connect to Server
  WiFiClient client;
  const int httpPort = port;
  if (!client.connect(host, httpPort))
  {
    Serial.println("connection failed");
    delay(10000);
    return;
  }

  //Arduino maps analog output to range between 0-1023. Each unit represents is 0.0049 volts.
  int sensorOutput = analogRead(0);

  //Add sensor output to sensorData
  StaticJsonDocument<200> sensorData;
  sensorData["value"] = String(sensorOutput);
  String sensorDataJson = "";
  serializeJson(sensorData, sensorDataJson);

  //Send POST Request to server
  client.println(String("POST ") + url + " HTTP/1.1");
  client.println(String("Host: ") + host + ":" + port);
  client.println("Content-Type: application/json");
  client.println("User-Agent: Arduino/1.0");
  client.println("Connection: close");
  client.print("Content-Length: ");
  client.println(sensorDataJson.length());
  client.println();
  client.println(sensorDataJson);

  delay(60000);
}
