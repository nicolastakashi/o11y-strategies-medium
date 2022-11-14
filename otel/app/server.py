from flask import Flask

app = Flask(__name__)

@app.route("/status")
def server_request():
    return "Ok"


if __name__ == "__main__":
    app.run("0.0.0.0", port=9090)
