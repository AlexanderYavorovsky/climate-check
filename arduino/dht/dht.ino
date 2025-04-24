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
  delay(5000);

  float h = dht.readHumidity();
  float t = dht.readTemperature();
  if (isnan(h) || isnan(t)) {
    return;
  }

  Serial.print(h);
  Serial.print(" ");
  Serial.print(t);
}
