#ifndef Bio_Comms
#define Bio_Comms

#include <Arduino.h>
#include "ArduinoJson.h"

class Comms{
    int baud;
    String identifier;

public:
    Comms(String id, int baud){
        this->baud = baud;
        identifier = id;

        Serial.begin(baud);
    }

    void sendJSON(DynamicJsonDocument data){

        DynamicJsonDocument doc(1024);
        doc["id"] = identifier;
        doc["data"] = data;

        serializeJson(doc, Serial);
        Serial.println("");
    }

};

#endif