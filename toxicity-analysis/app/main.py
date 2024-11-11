# toxicity-analysis/app/main.py

# To run this 
# python3 toxicity-analysis/main.py
# curl -X POST http://localhost:5000/toxicity-analysis -d '{"text": "This is a toxic example."}' -H "Content-Type: application/json"



from flask import Flask, request, jsonify

app = Flask(__name__)

@app.route('/toxicity-analysis', methods=['POST'])
def analyze_toxicity():
    data = request.get_json()
    text = data.get("text", "")
    # Mock processing: Let's assume any text with "toxic" has toxicity
    is_toxic = "toxic" in text.lower()
    return jsonify({"text": text, "toxicity_detected": is_toxic})

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
