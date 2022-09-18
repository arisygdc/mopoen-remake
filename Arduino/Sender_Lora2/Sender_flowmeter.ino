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

// uuidv4 value 36 character
#define idMonitoring "12c76e75-78f4-4cfd-b3cf-aae81373350e"

#define timeSecond 1000
// water flow sensor
int TURBINE;      //incremental signal measure
int HSensor = 2; //Arduino pin 2
int Calc;
 
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

  // water flow sensor setup
  pinMode(HSensor, INPUT);

  attachInterrupt(00, speedrpm, RISING); //interrupt
}
 
void loop() 
{
  delay(timeSecond * (60 * 10));

  TURBINE = 00;
  sei(); //enable interrupt 
  delay (timeSecond * 1);
  cli(); //disable interrupt
  Calc = (TURBINE * 60 / 7.5); //pulse * 60 / 7.5 

  //units of measurement L / hour
  String data = idMonitoring + String(Calc, DEC);
  LoraSend(LoRa, data);
}

void speedrpm () 	 //fungsi penghitungan dan interrupt
{
TURBINE++; //bersifat incrementing (dengan mode falling edge)
}

void LoraSend(LoRaClass l, String data) {
  l.beginPacket();   //Send LoRa packet to receiver
  l.print(data);
  l.endPacket();
}
