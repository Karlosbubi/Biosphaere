#ifndef Bio_Messung
#define Bio_Messung

#include <Arduino.h>
#include "DHT.h"
#include "ArduinoJson.h"


class Messung
{
	int ppm;
	int lux;

	float temp;
	float hum;


	int co2_sensor;
	int ldr;
	int RTH1;
	DHT *dht;

	int readCO2PWM(int pwmpin, int range)
	{

		unsigned long th;
		int ppm_pwm = 0;

		do
		{
			th = pulseIn(pwmpin, HIGH, 2500000) / 1000;

			float pulsepercent = th / 1004.0;

			ppm_pwm = range * pulsepercent;
		} while (th == 0);

		return ppm_pwm;
	}

	void messen()
	{
		ppm = readCO2PWM(co2_sensor, 5000);
		lux = analogRead(ldr);

		temp = dht->readTemperature();
		hum = dht->readHumidity();

	}

public:
	Messung(int CO2_Pin, int ldr_Pin, int RTH1_Pin, int baud)
	{
		co2_sensor = CO2_Pin;
		ldr = ldr_Pin;
		RTH1 = RTH1_Pin;

		pinMode(co2_sensor, INPUT);
		pinMode(ldr, INPUT);

		dht = new DHT(RTH1, DHT22);

		dht->begin();
		Serial.begin(baud);
	}

	StaticJsonDocument<32> toJSON(){
		messen();

	StaticJsonDocument<32> doc;

	doc["ppm"] = ppm;
	doc["lux"] = lux;
	doc["humid"] = hum;
	doc["temp"] = temp;

	return doc;
	}

	void sendMessung()
	{
		Serial.print(ppm);
		Serial.print(",");
		Serial.print(hum);
		Serial.print(",");
		Serial.print(temp);
		Serial.print(",");
		Serial.println(lux);
	}
};

#endif