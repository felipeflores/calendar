/*---------------------------------------------------------
  Programa :
  Autor    :
  Data     :
  ---------------------------------------------------------*/
#include <WiFi.h>      //Biblioteca para utilização do WiFi
#include <ESP32Ping.h>
#include <PubSubClient.h>

#include "FS.h"
#include "SPIFFS.h"

#define FORMAT_SPIFFS_IF_FAILED true

#define CONFIG_FILE "/config.cfg"

#define LED_BUILTIN 2  //Led no GPIO 2 no ESP32
#define LED_BUILTIN_4 4  //Led no GPIO 4 no ESP32

#define CONNECTION_TIMEOUT 10

bool alreadConfigured = false;

char rx_byte;
String stringOne = String('\0');


const char* ssid = "";
const char* password =  "";

//Mqqt
const char* mqttServer = "";
int mqttPort = 0;

WiFiClient espClient;
PubSubClient client(espClient);

int baudRate = 9600;

/*--- SETUP ---*/
void setup() {
  Serial.begin(baudRate); //Inicializa a comunicação serial com taxa de transmissão de 115200 baud rate

  pinMode(LED_BUILTIN, OUTPUT); //Configura o pino do led
  pinMode(LED_BUILTIN_4, OUTPUT); //Configura o pino do led

  Serial.println("Iniciou toda a bagaça");
  if (!SPIFFS.begin(FORMAT_SPIFFS_IF_FAILED)) {
    Serial.println("SPIFFS Mount Failed");
    return;
  }

  Serial.println("Cmecou a parte de verificar se existe");
  if (!existConfigFile(SPIFFS)) {
    alreadConfigured = false;
    Serial.println("Não esta configurado");
    initConfigFile(SPIFFS);
  } else {
    Serial.println("Existe o arquivo");
    readVariables(SPIFFS);
    if (ssid != "" && password != "" &&
        mqttServer != "" && mqttPort > 0) {
      alreadConfigured = true;
    }
  }
  Serial.println(alreadConfigured);
  delay(2000);
  if (alreadConfigured) {
    pinMode(LED_BUILTIN, OUTPUT); //Configura o pino do led
    pinMode(LED_BUILTIN_4, OUTPUT); //Configura o pino do led

    digitalWrite(LED_BUILTIN_4, HIGH);
    delay(1000);
    digitalWrite(LED_BUILTIN_4, LOW);

    //
    //  // Configures static IP address
    //  if (!WiFi.config(local_IP, gateway, subnet, primaryDNS, secondaryDNS)) {
    //    digitalWrite(LED_BUILTIN, HIGH);
    //    delay(1000);
    //    digitalWrite(LED_BUILTIN, LOW);
    //    delay(1000);
    //    digitalWrite(LED_BUILTIN, HIGH);
    //    delay(1000);
    //    digitalWrite(LED_BUILTIN, LOW);
    //    delay(1000);
    //  }

    WiFi.mode(WIFI_STA); //Optional
    WiFi.begin(ssid, password);

    Serial.println("Conectando");
    int timeout_counter = 0;

    while (WiFi.status() != WL_CONNECTED) {
      Serial.print(".");
      delay(200);
      timeout_counter++;
      if (timeout_counter >= CONNECTION_TIMEOUT * 5) {
        digitalWrite(LED_BUILTIN, HIGH);
        delay(1000);
        digitalWrite(LED_BUILTIN, LOW);
        delay(1000);
        digitalWrite(LED_BUILTIN, HIGH);
        delay(1000);
        digitalWrite(LED_BUILTIN, LOW);
        delay(1000);
        digitalWrite(LED_BUILTIN, HIGH);
        delay(1000);
        digitalWrite(LED_BUILTIN, LOW);
        delay(1000);
        ESP.restart();
      }
    }

    digitalWrite(LED_BUILTIN, HIGH);
    Serial.println("\nConnected to the WiFi network");
    Serial.println("IP address: ");
    Serial.println(WiFi.localIP());

    bool success = false;
    while (!success) {
      success = Ping.ping("www.google.com", 3);
      if (!success) {
        Serial.println("Ping failed");
      } else {
        Serial.println("Ping succesful.");
      }
    }

    delay(2000);

    client.setServer(mqttServer, mqttPort);
    client.setCallback(callback);

    timeout_counter = 0;

    String chipId = String((uint32_t)ESP.getEfuseMac(), HEX);
    chipId.toUpperCase();

    while (!client.connected()) {
      Serial.println("Connecting to MQTT...");
      timeout_counter++;
      if (timeout_counter >= CONNECTION_TIMEOUT * 5) {
        break;
      }
      if (client.connect(chipId.c_str())) {
        digitalWrite(LED_BUILTIN_4, HIGH);
        delay(2000);
        digitalWrite(LED_BUILTIN_4, LOW);
        Serial.println("connected");
      } else {
        digitalWrite(LED_BUILTIN_4, HIGH);
        delay(500);
        digitalWrite(LED_BUILTIN_4, LOW);
        delay(500);
        digitalWrite(LED_BUILTIN_4, HIGH);
        delay(500);
        digitalWrite(LED_BUILTIN_4, LOW);
        delay(500);
        digitalWrite(LED_BUILTIN_4, HIGH);
        delay(500);
        digitalWrite(LED_BUILTIN_4, LOW);
        Serial.print("failed with state ");
        Serial.println(client.state());
        delay(2000);
      }
    }
    if (client.connected()) {
      client.subscribe("felipe-casa/calendar");
    }

    delay(1000);
  }
  Serial.println("END");
}

/*--- LOOP PRINCIPAL ---*/
void loop() {

  if (Serial.available() > 0) {
    rx_byte = Serial.read();
    
   
    if (rx_byte == EOF || rx_byte == '\r' || rx_byte == '\n') {
      
      Serial.print("you typed: ");
      Serial.println(stringOne);

      if (stringOne == "reset") {
        resetConfig(SPIFFS);
      } else if (stringOne == "clear") {
        Serial.end();
        Serial.begin(baudRate);
        Serial.println("Olha aqui");
      } else if (stringOne == "networks") {
        getNetworks();
        Serial.flush();
      } else {
        if (stringOne != "") {
          String v = stringOne + ";";
          addVariable(SPIFFS, v.c_str());
        }
      }
      stringOne = String('\0');
    } else {

      stringOne.concat(rx_byte);
    }
  }
  if (alreadConfigured) {
    client.loop();
  }
}


void getNetworks() {
  Serial.println("START_WIFI");
  int countWifi = WiFi.scanNetworks(); //Retorna a quantidade de redes wifi disponíveis
  if (countWifi > 0) {
    Serial.print("{ \n");
    Serial.print(" \"count\":");
    Serial.print(countWifi);
    Serial.print(",\n");

    String comma = "";
    Serial.print(" \"networks\": [ \n");
    for (int i = 0; i < countWifi; ++i) {
      Serial.print(comma + "{ \n");

      Serial.print(" \"ssi\": \"");
      Serial.print(WiFi.SSID(i));
      Serial.print("\",");

      Serial.print(" \"rssi\": ");
      Serial.print(WiFi.RSSI(i));

      if (WiFi.encryptionType(i) != WIFI_AUTH_OPEN) {
        Serial.print(",");
        Serial.print(" \"encripted\": true ");
      }
      Serial.print("} \n");
      comma = ",";
    }

    Serial.print("] \n ");
    Serial.print("}\n");
  } else {
    Serial.println("Nenhuma rede encontrada");
  }
  Serial.println("END_WIFI");
}

void callback(char* topic, byte* payload, unsigned int length) {
  Serial.print("Message arrived in topic: ");
  Serial.println(topic);
  String messageTemp;
  for (int i = 0; i < length; i++) {
    messageTemp += (char)payload[i];
  }
  Serial.println("Message: " + messageTemp);

  // If a message is received on the topic esp32/output, you check if the message is either "on" or "off".
  // Changes the output state according to the message
  if (String(topic) == "felipe-casa/calendar") {
    Serial.print("Changing output to ");
    if (messageTemp == "on") {
      Serial.println("Relay ON");
      digitalWrite(LED_BUILTIN_4, HIGH);    //pins on board seams to be working in "reverse"
    }
    else if (messageTemp == "off") {
      Serial.println("Relay OFF");
      digitalWrite(LED_BUILTIN_4, LOW);   //pins on board seams to be working in "reverse"
    }
  }
}


boolean existConfigFile(fs::FS &fs) {
  Serial.printf("Reading file: %s\r\n", CONFIG_FILE);

  File file = fs.open(CONFIG_FILE);
  if (!file || file.isDirectory()) {
    Serial.println("− failed to open file for reading");
    return false;
  }
  return true;
}
void initConfigFile(fs::FS &fs) {
  Serial.printf("Init file: %s\r\n", CONFIG_FILE);

  File file = fs.open(CONFIG_FILE, FILE_WRITE);
  if (!file) {
    Serial.println("− failed to open file for initialize");
    return;
  }
  if (file.print("")) {
    Serial.println("− inicializado");
  } else {
    Serial.println("− init failed");
  }
}

void resetConfig(fs::FS &fs) {
  Serial.printf("Deleting file: %s\r\n", CONFIG_FILE);
  if (fs.remove(CONFIG_FILE)) {
    Serial.println("− file deleted");
  } else {
    Serial.println("− delete failed");
  }
  initConfigFile(fs);
}

void addVariable(fs::FS &fs, const char * variable) {
  Serial.printf("Appending to file: %s\r\n", CONFIG_FILE);
  File file = fs.open(CONFIG_FILE, FILE_APPEND);
  if (file.print(variable)) {
    Serial.println("− message appended");
  } else {
    Serial.println("− append failed");
  }
}

void readVariables(fs::FS &fs) {
  File file = fs.open(CONFIG_FILE);
  if (!file) {
    Serial.println("− failed to open file for initialize");
    return;
  }
  bool existContent = false;

  char value [512] = {'\0'};
  uint16_t i = 0;
  while (file.available()) {
    existContent = true;
    value [i] = file.read();
    i++;
  }
  value [i] = '\0';
  file.close();

  if (existContent) {
    int arraySize = sizeByKey(value, ';');
    String * myArray = new String[arraySize];
    split(arraySize, value, ';', myArray);

    for (int i = 0; i < arraySize; ++i) {
      String newValue = myArray[i] + "=";
      int as = sizeByKey(newValue, '=');
      String * valueLine = new String[as];
      split(as, newValue, '=', valueLine);

      if (valueLine[0] == "SSID") {
        ssid = valueLine[1].c_str();
      } else if (valueLine[0] == "PASSWORD") {
        password = valueLine[1].c_str();
      } else if (valueLine[0] == "MQTT_SERVER") {
        mqttServer = valueLine[1].c_str();
      } else if (valueLine[0] == "MQTT_PORT") {
        mqttPort = valueLine[1].toInt();
      }
    }
  }
}

int sizeByKey(String str, char key) {
  int count = 0;
  for (int i = 0; i < str.length(); i++) {
    if (str[i] == key) {
      count++;
    }
  }
  return count;
}

void split(int arraySize, String str, char key, String* strs) {
  int StringCount = 0;
  // Split the string into substrings
  for (uint8_t i = 0; i < arraySize; i++) {
    int index = str.indexOf(key);
    if (index == -1) {// No space found
      strs[StringCount++] = str;
      break;
    } else {
      strs[StringCount] = str.substring(0, index);
      StringCount++;
      str = str.substring(index + 1);
    }
  }
}

void deleteFile(fs::FS &fs, const char * path) {
  Serial.printf("Deleting file: %s\r\n", path);
  if (fs.remove(path)) {
    Serial.println("− file deleted");
  } else {
    Serial.println("− delete failed");
  }
}
