#include <Arduino.h>

#include "../lib/messung.hpp"
#include "../lib/communication.hpp"

const int led1 = 7;
const int led2 = 8;
const int led3 = 9;

const int co2_sensor = 4;
const int ldr = A0;
const int RTH1 = 13;

Messung* messung;
Comms* comm;

void setup()
{
  pinMode(led1, OUTPUT);
  pinMode(led2, OUTPUT);
  pinMode(led3, OUTPUT);

  messung = new Messung(co2_sensor, ldr, RTH1);
  comm = new Comms("Testing", 9600);
}

void loop()
{
  comm->sendJSON(messung->toJSON());
  delay(5000);
}