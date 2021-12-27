from rpi_ws281x import PixelStrip, Color
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

if __name__ == '__main__':
    strip = PixelStrip(LED_COUNT, LED_PIN, LED_FREQ_HZ, LED_DMA, LED_INVERT, LED_BRIGHTNESS, LED_CHANNEL)
    strip.begin()

    try:
        while True:
            colorWipe(strip, Color(0, 0, 255), 24) 
            colorWipe(strip, Color(0, 0, 0), 24) 
    except KeyboardInterrupt:
        colorWipe(strip, Color(0, 0, 0), 10)