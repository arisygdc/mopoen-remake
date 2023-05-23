#include <LoRa.h>
#include <SPI.h>
#include <ArduinoJson.h>
 
#define ss 10
#define rst 9
#define dio0 8

// uuidv4 value 36 character
#define IdMonitoring "f601a393-018b-493c-b929-337ab3693916"
#define Secret "a9a0cd8780b44782g8df720715a74102"

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
  TURBINE = 00;
  sei(); //enable interrupt 
  delay (timeSecond * 1);
  cli(); //disable interrupt
  Calc = (TURBINE * 60 / 7.5); //pulse * 60 / 7.5 

  //units of measurement L / hour
  // convert to m3/s
  float debit = Calc/1000;
  
  DynamicJsonDocument doc(1024);
  doc["id"] = IdMonitoring;
  doc["value"] = debit;
  doc["secret"] = Secret;

  String json;
  serializeJson(doc, json);
  Serial.println(json);

  // String data = idMonitoring + String(Calc, DEC);
  LoraSend(LoRa, json);
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
