from rpi_ws281x import PixelStrip, Color
import flask
import time

LED_COUNT = 60        
LED_PIN = 18          
LED_FREQ_HZ = 800000  
LED_DMA = 10          
LED_BRIGHTNESS = 255  
LED_INVERT = False    
LED_CHANNEL = 0       

def colorWipe(strip, color, wait_ms=50):
    for i in range(strip.numPixels() ):
        strip.setPixelColor(i, color)
        time.sleep(wait_ms / 1000.0)
        strip.show()

def setColor(strip, color ):
    for i in range(strip.numPixels()):
        strip.setPixelColor(i, color)
        time.sleep(20 / 1000.0)
        strip.show()
    strip.show()

strip = PixelStrip(LED_COUNT, LED_PIN, LED_FREQ_HZ, LED_DMA, LED_INVERT, LED_BRIGHTNESS, LED_CHANNEL)
strip.begin()

app = flask.Flask(__name__)
app.config["DEBUG"] = True

@app.route('/setStaticColor', methods=['POST'])
def setStaticColor():
    jsonBody = flask.request.get_json()
    setColor(strip, Color(jsonBody["r"], jsonBody["g"], jsonBody["b"]))
    return { "success": True } 

try:
    app.run(host="192.168.0.164")   
except KeyboardInterrupt:
    setColor(strip, Color(0, 0, 0))





