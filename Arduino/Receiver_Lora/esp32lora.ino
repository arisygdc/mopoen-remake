#include <Arduino.h>
//Libraries for LoRa
#include <SPI.h>
#include <LoRa.h>

// HTTP client using wifi connection
#include <WiFi.h>
#include <HTTPClient.h>

//define the pins used by the LoRa transceiver module
#define ss 5
#define rst 14
#define dio0 2
 
#define BAND 433E6    //433E6 for Asia, 866E6 for Europe, 915E6 for North America

int sigNoData = 0;

const char* ssid = "REPLACE_WITH_YOUR_SSID";
const char* password = "REPLACE_WITH_YOUR_PASSWORD";

//Your Domain name with URL path or IP address with path
#define SERVER_ADDRESS "192.168.1.73"

struct LORA_RESULT {
  int rssi;
  int read;
  String payload;
};

struct HTTP_RESULT {
  int code;
  String payload;
};

void setup() {
  Serial.begin(115200);
  while(!Serial){delay(100);}
  startLoRA();
  startWiFi();
}

void loop() {
  LORA_RESULT lores = readLoRA();
  if (lores.read > 0) {
    Serial.printf("Read: %d\n", lores.read);
    Serial.print("Received: '");
    Serial.print(lores.payload);
    Serial.printf("' RSSI %d\n", lores.rssi);
    String url = "http://" SERVER_ADDRESS "/api/sensor/value"; 
    sendApi(url, lores.payload);
    sigNoData = 0;
  } else {
    sigNoData += 1;
  }

  if (sigNoData > 50) {
    Serial.println("NO DATA RECEIVE");
  }

  delay(90);
}

void startLoRA()
{
  LoRa.setPins(ss, rst, dio0); //setup LoRa transceiver module
  Serial.println("LoRa initialize");
  while (!LoRa.begin(BAND)) {
    Serial.print(".");
    delay(500);
  }
  LoRa.setSyncWord(0xA5);
  Serial.println("LoRa Initialization OK!");
  delay(2000);
  LoRa.receive();
}

void startWiFi()
{
  WiFi.begin(ssid, password);
  Serial.print("Connecting to ");
  Serial.println(ssid);
  while(WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");
  }
  Serial.print("\nIP address: ");
  Serial.println(WiFi.localIP());
}

LORA_RESULT readLoRA()
{
  LORA_RESULT res = {0,0,""};
  int packetSize = LoRa.parsePacket();    // try to parse packet
  if (packetSize) 
  {
    Serial.print("Received packet ");
    while (LoRa.available())              // read packet
    {
      res.read += 1;
      res.payload += LoRa.readString();
    }
  }
  return res;
}

HTTP_RESULT sendApi(String url, String jsonBody) 
{
  HTTP_RESULT res = {-1, ""};
  if(WiFi.status()== WL_CONNECTED){
    HTTPClient http;
    http.begin(url);
    http.addHeader("Content-Type", "application/json");
    res.code = http.POST(jsonBody);
    res.payload = http.getString();
  }
  return res;
}
