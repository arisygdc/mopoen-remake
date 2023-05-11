void loop() {
  // put your main code here, to run repeatedly:
  String data = "";
  for (int retry = 0; retry < 3; retry++) {
    while(Serial.available()>0){
      data += char(Serial.read());
    }
    if (data.length() > 97) {
      break;
    }
    delay(100);
  }

  if (data.length() > 97) {
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

  delay(100); 
}
