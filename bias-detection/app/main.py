# bias-detection/app/main.py
# To run and test this you need to just run 
# python3 main.py
# curl -X POST http://localhost:5000/bias-detection -d '{"text": "Sample text with bias"}' -H "Content-Type: application/json"

from flask import Flask, request, jsonify

app = Flask(__name__)

@app.route('/bias-detection', methods=['POST'])
def detect_bias():
    data = request.get_json()
    text = data.get("text", "")
    # Mock processing: Let's assume any text with "bias" has bias
    has_bias = "bias" in text.lower()
    return jsonify({"text": text, "bias_detected": has_bias})

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
