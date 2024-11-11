# reporting/app/main.py

# To run this
# curl http://localhost:5000/reporting


from flask import Flask, jsonify

app = Flask(__name__)

@app.route('/reporting', methods=['GET'])
def generate_report():
    # Mock report data
    report_data = {
        "bias_analysis": {"total_texts": 10, "bias_detected": 3},
        "toxicity_analysis": {"total_texts": 10, "toxicity_detected": 2}
    }
    return jsonify({"report": report_data})

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
