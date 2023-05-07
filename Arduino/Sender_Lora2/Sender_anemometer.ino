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
#define secret "4923296e0e0937f86e0853f9366c1de635b22a3d4675c08a61057aac95a9d7d8"
float counter = 1;

// anemometer parameters
volatile byte rpmcount; // count signals
volatile unsigned long last_micros;
unsigned long timeold;
unsigned long timemeasure = 25.00; // seconds
int timetoSleep = 1;               // minutes
unsigned long sleepTime = 15;      // minutes
unsigned long timeNow;
int countThing = 0;
int GPIO_pulse = 2; // Arduino pin 2
float rpm, rps;     // frequencies
float radius = 0.1; // meters - measure of the lenght of each the anemometer wing
float velocity_ms;  //m/s
float omega = 0;    // rad/s
float calibration_value = 2.0;
 
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

  // Anemometer setup
  pinMode(GPIO_pulse, INPUT_PULLUP);
  digitalWrite(GPIO_pulse, LOW);

  detachInterrupt(digitalPinToInterrupt(GPIO_pulse));                         // force to initiate Interrupt on zero
  attachInterrupt(digitalPinToInterrupt(GPIO_pulse), rpm_anemometer, RISING); //Initialize the intterrupt pin
  rpmcount = 0;
  rpm = 0;
  timeold = 0;
  timeNow = 0;
}
 
void loop() 
{
  if ((millis() - timeold) >= timemeasure * 1000)
  {
    countThing++;
    detachInterrupt(digitalPinToInterrupt(GPIO_pulse)); // Disable interrupt when calculating
    rps = float(rpmcount) / float(timemeasure);         // rotations per second
    rpm = 60 * rps;                                     // rotations per minute
    omega = 2 * PI * rps;                               // rad/s
    velocity_ms = omega * radius * calibration_value;   // m/s
    Serial.print("rps=");
    Serial.print(rps);
    Serial.print("   rpm=");
    Serial.print(rpm);
    Serial.print("   velocity_ms=");
    Serial.print(velocity_ms);
    Serial.println("   ");
    if (countThing == (60 * 10)) // Send data per 25 seconds
    {
      Serial.println("Send data to server");
      // {\"id\": \"%s\", \"value\": %s, \"secret\": \"%s\"}
      String data = "{\"id\"" + idMonitoring;
      data = data + "\", \"value\":";
      data = data + String(velocity_ms);
      data = ", \"secret\": \"";
      data = data + secret;
      data = data + "\"}";
      LoraSend(LoRa, data);
      countThing = 0;
    }
    timeold = millis();
    rpmcount = 0;
    attachInterrupt(digitalPinToInterrupt(GPIO_pulse), rpm_anemometer, RISING); // enable interrupt
  }
}

void rpm_anemometer()
{
  if (long(micros() - last_micros) >= 5000)
  { // time to debounce measures
    rpmcount++;
    last_micros = micros();
  }
  //   Serial.println("***** detect *****");
}

void LoraSend(LoRaClass l, String data) {
  l.beginPacket();   //Send LoRa packet to receiver
  l.print(data);
  l.endPacket();
}
