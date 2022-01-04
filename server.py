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

state = {
    "color": Color(0, 0, 0), 
    "mode": "static"
}

def wheel(pos):
    """Generate rainbow colors across 0-255 positions."""
    if pos < 85:
        return Color(pos * 3, 255 - pos * 3, 0)
    elif pos < 170:
        pos -= 85
        return Color(255 - pos * 3, 0, pos * 3)
    else:
        pos -= 170
        return Color(0, pos * 3, 255 - pos * 3)

def colorWipe(strip, color, wait_ms=50):
    for i in range(strip.numPixels() ):
        strip.setPixelColor(i, color)
        time.sleep(wait_ms / 1000.0)
        strip.show()

def setColor(strip, color ):
    for i in range(strip.numPixels()):
        strip.setPixelColor(i, color)
        # time.sleep(20 / 1000.0)
        # strip.show()
    strip.show()

def rainbow(strip, wait_ms=10, iterations=5):
    for j in range(256 * iterations):
        for i in range(strip.numPixels()):
            strip.setPixelColor(i, wheel((i + j) & 255))
            # strip.setPixelColor(i, wheel((int(i * 256 / strip.numPixels()) + j) & 255))
        strip.show()
        time.sleep(wait_ms / 1000.0)

strip = PixelStrip(LED_COUNT, LED_PIN, LED_FREQ_HZ, LED_DMA, LED_INVERT, LED_BRIGHTNESS, LED_CHANNEL)
strip.begin()

app = flask.Flask(__name__)
app.config["DEBUG"] = True

@app.route('/setStaticColor', methods=['POST'])
def setStaticColor():
    jsonBody = flask.request.get_json()
    global state
    state["mode"] = "static"
    state["color"] = Color(jsonBody["r"], jsonBody["g"], jsonBody["b"])
    setColor(strip, state["color"])
    return { "success": True } 

@app.route('/on', methods=['POST'])
def setOn():
    global state
    if state["mode"] == "static":
        print(state["color"])
        setColor(strip, state["color"])
    elif state.mode == "rainbow":
        setRainbow(strip)

    return { "success": True } 

@app.route('/off', methods=['POST'])
def setOff():
    setColor(strip, Color(0, 0, 0))
    return { "success": True } 

@app.route('/setMode', methods=['POST'])
def setMode():
    jsonBody = flask.request.get_json()
    global state
    state["mode"] = jsonBody["mode"]

    if state["mode"] == "rainbow":
        rainbow(strip)

    return { "success": True } 

try:
    app.run(host="192.168.0.164")   
except KeyboardInterrupt:
    setColor(strip, Color(0, 0, 0))
