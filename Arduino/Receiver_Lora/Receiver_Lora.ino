
//BOARD MANAGER ESP32 1.0.2
//LIBRARY MANAGER LORA SANDEEP MISTRY 0.8.0
//GND   GND
//3.3V  VCC
//D5    NSS
//D23   MOSI
//D19   MISO
//D18   SCK
//D14   RST
//D2    DIO0


#include <LoRa.h>
#include <SPI.h>
#include <SoftwareSerial.h>
SoftwareSerial espSerial(5, 6);

 
#define ss 10
#define rst 9
#define dio0 8
 
void setup() 
{
  Serial.begin(9600);
  while (!Serial);
  Serial.println("LoRa Receiver");
 
  //LoRa.setPins(ss, rst, dio0);    //setup LoRa transceiver module
 
  while (!LoRa.begin(433E6))     //433E6 - Asia, 866E6 - Europe, 915E6 - North America
  {
    Serial.println(".");
    delay(500);
  }
  espSerial.begin(9600);
  LoRa.setSyncWord(0xA5);
  Serial.println("LoRa Initializing OK!");
}
 
void loop() 
{
  int packetSize = LoRa.parsePacket();    // try to parse packet
  if (packetSize) 
  {
    
    Serial.print("Received packet '");
 
    while (LoRa.available())              // read packet
    {
      String LoRaData = LoRa.readString();
      Serial.print(LoRaData); 
      espSerial.print(LoRaData); 
    }
    Serial.print("' with RSSI ");         // print RSSI of packet
    Serial.println(LoRa.packetRssi());
  }
}
