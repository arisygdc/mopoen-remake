#include <ESP8266WiFi.h>
#include <ESP8266HTTPClient.h>
#include <ArduinoJson.h>
#include <SPI.h>

#define SERVER_IP "192.168.1.73:8080"
#define METHOD_POST "POST"

#ifndef STASSID
#define STASSID "ssid"
#define STAPSK  "pass"
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
  Serial.println(WiFi.localIP());
}

void loop() {
  // put your main code here, to run repeatedly:
  String data = "";
  while(Serial.available()>0){
    data += char(Serial.read());
  }

  if (data.length() > 35) {
    Serial.print("received: ");
    Serial.println(data);

    String alamat = "http://" SERVER_IP "/api/sensor/value"; 
    HTTP_RESULT getResponse = HTTPsend(METHOD_POST, alamat, data);

    Serial.print("[HTTP POST]");
    Serial.print(alamat);
    Serial.printf(" code: %d\n", getResponse.code);
    Serial.print("message: ");
    Serial.println(getResponse.payload);
  }

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
