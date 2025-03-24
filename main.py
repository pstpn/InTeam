from flask import Flask, request, send_file, Response
import base64
import os
import uuid

app = Flask(__name__)

UPLOAD_FOLDER = 'uploads'
os.makedirs(UPLOAD_FOLDER, exist_ok=True)

@app.after_request
def after_request(response):
    response.headers.add('Access-Control-Allow-Origin', '*')
    response.headers.add('Access-Control-Allow-Headers', 'Content-Type,Authorization')
    response.headers.add('Access-Control-Allow-Methods', 'GET,PUT,POST,DELETE,OPTIONS')
    return response

@app.route('/convert', methods=['POST', 'OPTIONS'])
def base64_to_image():
    if request.method == 'OPTIONS':
        return Response()
    
    try:
        base64_str = request.json.get('image_data')
        if not base64_str:
            return Response("No image_data provided", status=400)
        
        image_data = base64.b64decode(base64_str)
        
        filename = f"{uuid.uuid4()}.jpg"
        filepath = os.path.join(UPLOAD_FOLDER, filename)
        
        with open(filepath, 'wb') as f:
            f.write(image_data)
        
        return send_file(filepath, mimetype='image/jpeg')
    
    except Exception as e:
        return Response(f"Error processing image: {str(e)}", status=500)

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)