import json
import requests
from flask import Flask, request, jsonify
from flask_sqlalchemy import SQLAlchemy

db = SQLAlchemy()
app = Flask(__name__)
model = "phi-custom1"  # TODO: update this for whatever model you wish to use

# Configure the SQLite database
app.config['SQLALCHEMY_DATABASE_URI'] = 'sqlite:///chat_history.db'
app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False

db.init_app(app)

# Define the MessageHistory model
class MessageHistory(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    user_id = db.Column(db.String(100), nullable=False)
    role = db.Column(db.String(10), nullable=False)
    content = db.Column(db.Text, nullable=False)

# Create the database and tables
with app.app_context():
    db.create_all()

def chat(messages):
    r = requests.post(
        "http://0.0.0.0:11434/api/chat",
        json={"model": model, "messages": messages, "stream": True},
        stream=True
    )
    r.raise_for_status()
    output = ""

    for line in r.iter_lines():
        body = json.loads(line)
        if "error" in body:
            raise Exception(body["error"])
        if body.get("done") is False:
            message = body.get("message", "")
            content = message.get("content", "")
            output += content
            # the response streams one token at a time, print that as we receive it
            # print(content, end="", flush=True)

        if body.get("done", False):
            message["content"] = output
            return message

@app.route('/chat', methods=['POST'])
def handle_chat():
    data = request.json
    user_id = data.get("user_id")
    user_input = data.get("message")

    if not user_id or not user_input:
        return jsonify({"error": "user_id and message are required"}), 400

    # Retrieve the user's message history from the database
    messages = MessageHistory.query.filter_by(user_id=user_id).all()
    message_list = [{"role": msg.role, "content": msg.content} for msg in messages]

    # Print the received request
    print(f"Received request from user_id: {user_id} with message: {user_input}")

    # Append the new user message to the history
    new_message = MessageHistory(user_id=user_id, role="user", content=user_input)
    db.session.add(new_message)
    db.session.commit()
    message_list.append({"role": "user", "content": user_input})

    # Get the model's response
    response_message = chat(message_list)

    # Save the model's response to the history
    new_response = MessageHistory(user_id=user_id, role="assistant", content=response_message["content"])
    db.session.add(new_response)
    db.session.commit()

    return jsonify({"response": response_message["content"]})

@app.route('/new', methods=['POST'])
def new_session():
    data = request.json
    user_id = data.get("user_id")

    if not user_id:
        return jsonify({"error": "user_id is required"}), 400

    # Delete all messages for the user
    MessageHistory.query.filter_by(user_id=user_id).delete()
    db.session.commit()

    return jsonify({"message": "All previous messages deleted for user_id: {}".format(user_id)}), 200

if __name__ == "__main__":
    app.run(host='0.0.0.0', port=5000)
