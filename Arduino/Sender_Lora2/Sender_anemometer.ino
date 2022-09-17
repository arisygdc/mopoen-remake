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

// uuidv4 value 36 character
#define idMonitoring "12c76e75-78f4-4cfd-b3cf-aae81373350e"
float counter = 1;

// anemometer parameters
volatile byte rpmcount; // count signals
volatile unsigned long last_micros;
unsigned long timeold;
unsigned long timemeasure = 25.00; // seconds
int timetoSleep = 1;               // minutes
unsigned long sleepTime = 15;      // minutes
unsigned long timeNow;
int GPIO_pulse = 2; // Arduino = D2
float rpm, rps;     // frequencies
float radius = 0.1; // meters - measure of the lenght of each the anemometer wing
float velocity_kmh; // km/h
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
  float wind_speed = anemometer_read();
  String data = idMonitoring + String(wind_speed);
  LoraSend(LoRa, data);
  counter += 1.5;
  delay(timeSecond * 5);
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

float anemometer_read() {
  detachInterrupt(digitalPinToInterrupt(GPIO_pulse)); // Disable interrupt when calculating
  rps = float(rpmcount) / float(timemeasure);         // rotations per second
  rpm = 60 * rps;                                     // rotations per minute
  rpmcount = 0;
  attachInterrupt(digitalPinToInterrupt(GPIO_pulse), rpm_anemometer, RISING); // enable interrupt
  return rpm;
}

void LoraSend(LoRaClass l, String data) {
  l.beginPacket();   //Send LoRa packet to receiver
  l.print(data);
  l.endPacket();
}
