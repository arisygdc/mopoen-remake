/**
   PostHTTPClient.ino

    Created on: 21.11.2016

*/

#include <ESP8266WiFi.h>
#include <ESP8266HTTPClient.h>
#include <SPI.h>
#include <LoRa.h>

#define SERVER_IP "192.168.0.4"
#define METHOD_POST "POST"

#ifndef STASSID
#define STASSID "ssid"
#define STAPSK  "44432100"
#endif

int id_sensor = 0;

struct HTTP_RESULT {
  int code;
  String payload;
};

void setup() {
  Serial.begin(9600);
  Serial.println();
  Serial.println();
  Serial.println();
  // Serial Connected

  WiFi.begin(STASSID, STAPSK);

  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");
  }
  Serial.println("");
  Serial.print("Connected! IP address: ");
  Serial.println(WiFi.localIP());
  // WiFi Connected
  
  if (!LoRa.begin(433E6)) {
    Serial.println("Starting LoRa failed!");
    while (1);
  }
  // LoRa Connected 
  
  delay(2000);
}

void loop() {
  // Receive data from Lora
  int packetSize = LoRa.parsePacket();
  if (packetSize) {
    // received a packet
    Serial.print("Received packet '");

    // read packet
    while (LoRa.available()) {
      Serial.print((char)LoRa.read());
    }

    // print RSSI of packet
    Serial.print("' with RSSI ");
    Serial.println(LoRa.packetRssi());
  }
  
  //   wait for WiFi connection
  HTTP_RESULT getResponse;
  if ((WiFi.status() == WL_CONNECTED)) {
    String valueSensorJson = (String)LoRa.read();
    getResponse = HTTPsend(METHOD_POST, "http://" SERVER_IP "/api/sensor/value", valueSensorJson);
    Serial.printf("[HTTP] POST SETUP DEVICE... code: %d\n", getResponse.code);
  }

  delay(10000);
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
