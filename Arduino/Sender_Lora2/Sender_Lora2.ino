
//#BOARD MANAGER ESP32 1.0.2
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
 
#define ss 10
#define rst 9
#define dio0 8
#define timeSecond 1000

float counter = 0;
 
void setup() 
{
  Serial.begin(9600); 
  while (!Serial);
  Serial.println("LoRa Sender");
 
  //Ra.setPins(ss, rst, dio0);    //setup LoRa transceiver module
  
  while (!LoRa.begin(433E6))     //433E6 - Asia, 866E6 - Europe, 915E6 - North America
  {
    Serial.println(".");
    delay(500);
  }
  LoRa.setSyncWord(0xA5);
  Serial.println("LoRa Initializing OK!");
}
 
void loop() 
{
  Serial.print("Sending packet: ");
  Serial.println(counter);
 
  LoRa.beginPacket();   //Send LoRa packet to receiver
  LoRa.print("cf0633d290c88dd347265ca490e9ae39: " + counter);
  // LoRa.print(counter);
  LoRa.endPacket();
 
  counter+=1.5;
 
  delay(timeSecond * );
}
