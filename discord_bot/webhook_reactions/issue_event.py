import logging
import discord
from discord_bot.discord_bot import get_current_bot

ISSUE_CHANNEL_ID = 1536884393870893137
log = logging.getLogger(__name__)


async def handle_issue_event(payload) -> None:
    _ = payload
    bot = get_current_bot()

    if bot is None:
        log.error("Current bot not set, not posting issue event")
        return
    channel = bot.get_channel(ISSUE_CHANNEL_ID)
    if not isinstance(channel, discord.TextChannel):
        log.error("Could not issue channel")
        return
    issue_dict = payload["issue"]
    url = issue_dict["repository_url"]
    issue_num = issue_dict["number"]
    title = issue_dict["title"]
    created_by = issue_dict["user"]["login"]
    body = issue_dict["body"]
    body_text = body if len(body) < 1200 else "Body too long for discord, view url"

    embed = discord.Embed(
        title="New issue created",
        description=f"#{issue_num}: {title}:\n{body_text}",
        color=0x00FF00,
    )
    embed.add_field(name="Issue Url", value=url, inline=False)
    embed.add_field(name="Created by", value=created_by, inline=False)

    await channel.send(embed=embed)
