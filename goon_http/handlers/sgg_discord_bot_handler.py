import asyncio
import hashlib
import hmac
import logging

from flask import Blueprint, request, abort

from discord_bot.discord_bot import get_current_bot
from discord_bot.webhook_reactions.release_event import post_new_release
from discord_bot.webhook_reactions.issue_event import handle_issue_event

log = logging.getLogger(__name__)

sgg_discord_bot_bp = Blueprint("sgg_discord_bot", __name__)

SECRET = "kevinb"


@sgg_discord_bot_bp.route("/api/v1/webhooks/sggDiscordBot", methods=["POST"])
def handle_supergoon_games_discord_bot():
    body = request.get_data()

    signature = request.headers.get("X-Hub-Signature", "")
    if not verify_signature(signature, body):
        log.warning("Invalid signature")
        abort(401)

    event_type = request.headers.get("X-GitHub-Event", "")
    bot = get_current_bot()
    loop = bot.loop

    if event_type == "release":
        payload = request.get_json()
        if payload.get("action") == "published":
            release = payload["release"]
            repository = payload["repository"]
            asyncio.run_coroutine_threadsafe(
                post_new_release(
                    repository["html_url"],
                    release["html_url"],
                    release["name"],
                    release["body"],
                    release["tag_name"],
                ),
                loop,
            ).result(timeout=10)
    elif event_type == "issues":
        payload = request.get_json()
        asyncio.run_coroutine_threadsafe(handle_issue_event(payload), loop).result(timeout=10)

    return "", 200


def verify_signature(signature: str, payload: bytes) -> bool:
    expected_mac = calculate_mac(payload)
    return hmac.compare_digest(signature, expected_mac)


def calculate_mac(payload: bytes) -> str:
    mac = hmac.new(SECRET.encode(), payload, hashlib.sha1)
    return "sha1=" + mac.hexdigest()
