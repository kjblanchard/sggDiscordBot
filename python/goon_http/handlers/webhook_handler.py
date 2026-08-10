from flask import Blueprint

webhook_bp = Blueprint("webhook", __name__)


@webhook_bp.route("/api/v1/webhooks", methods=["GET"])
def handle_webhook():
    return "Welcome User! !"
