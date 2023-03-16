from flask import Flask, abort
import requests
import signal
import random
import time

app = Flask(__name__)

@app.route("/status")
def server_request():
    return "Ok"

@app.route('/payments')
def payments():
    # num = random.randint(0, 100)
    # if num % 2 == 0:
    #     time.sleep(10)  

    return "Payed"

@app.route('/checkouts')
def checkout():
    # num = random.randint(0, 100)
    # if num % 2 == 0:
    #     response = requests.get('http://paymentapi:9090/payments')
    #     return response.text
    # else:
    #     abort(500)
        
    response = requests.get('http://paymentapi:9090/payments')
    return response.text

def sigterm_handler(signum, frame):
    print('Received SIGTERM signal. Exiting gracefully...')
    server.shutdown()
    exit(0)

if __name__ == "__main__":
    signal.signal(signal.SIGTERM, sigterm_handler)
    server = app.run(host='0.0.0.0', port=9090)
