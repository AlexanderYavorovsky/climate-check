#include "DHT.h"

#define DHTPIN 2
#define DHTTYPE DHT11

DHT dht(DHTPIN, DHTTYPE);

void setup() {
  Serial.begin(1000000);
  Serial.println("Starting dht");
  dht.begin();
}

void loop() {
  //wait for 10 seconds
  long waitMs = 10L * 1000L;
  delay(waitMs);

  float h = dht.readHumidity();
  float t = dht.readTemperature();
  if (isnan(h) || isnan(t)) {
    return;
  }

  Serial.print(h);
  Serial.print(" ");
  Serial.print(t);
}
