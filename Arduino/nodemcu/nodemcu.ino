#include <ESP8266WiFi.h>
#include <ESP8266HTTPClient.h>
#include <SPI.h>

#define SERVER_IP "192.168.1.2:8080"
#define METHOD_POST "POST"

#ifndef STASSID
#define STASSID "SSID"
#define STAPSK  "pwd"
#endif

struct HTTP_RESULT {
  int code;
  String payload;
};

void setup() {
  // put your setup code here, to run once:
  Serial.begin(9600);
  delay(2000);
  WiFi.begin(STASSID, STAPSK);

  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");
  }
}

void loop() {
  // put your main code here, to run repeatedly:
  String data = "";
  String id_sensor = "";
  String valStr = "";
  while(Serial.available()>0){
    data += char(Serial.read());
    id_sensor = data.substring(0, 35);
    valStr = data.substring(37, data.length());
  }

  if ((WiFi.status() == WL_CONNECTED)) {
    String valueSensorJson = "{\"id\": \"cf3eff4c-e693-4540-bbf9-624750e8bfa5 123\", \"value\": 32, \"secret\": \"c9a9cd5750b44752b8df720715a74108\"}"; //-> test with static value
    HTTP_RESULT getResponse = HTTPsend(METHOD_POST, "http://" SERVER_IP "/api/v1/sensor/data", valueSensorJson);
    Serial.printf("[HTTP] POST SETUP DEVICE... code: %d\n", getResponse.code);
  }

  Serial.println(data);
  delay(500); 
}

HTTP_RESULT HTTPsend(String http_method, String url, String jsonBody) {
  WiFiClient client;
  HTTPClient http;

  // configure traged server and url
  http.begin(client, url); //HTTP
  http.addHeader("Content-Type", "application/json");
  int httpCode;
  if (http_method == METHOD_POST) {
    // start connection and send HTTP header and body
    httpCode = http.POST(jsonBody);
  }
  // httpCode will be negative on error
  HTTP_RESULT getResp = {httpCode, http.getString()};
  http.end();
  return getResp;
}
