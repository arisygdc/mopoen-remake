void setup() {
  // put your setup code here, to run once:
  Serial.begin(9600);
  delay(2000);
}

void loop() {
  // put your main code here, to run repeatedly:
  String data = "";
  while(Serial.available()>0){
    data += char(Serial.read());
  }
  Serial.println(data);
  delay(500); 
}
