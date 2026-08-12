import logging
import discord

from discord_bot.webhook_reactions.reactions import RPG_NOTIFICATIONS_CHANNEL_ID

log = logging.getLogger(__name__)

bot_instance: discord.Client | None = None


def set_bot_instance(bot: discord.Client):
    global bot_instance
    bot_instance = bot


async def post_new_release(url: str, release_url: str, release_name: str, release_body: str, tag_name: str):
    if bot_instance is None:
        log.error("Bot instance not set")
        return
    channel = bot_instance.get_channel(int(RPG_NOTIFICATIONS_CHANNEL_ID))
    if not isinstance(channel, discord.TextChannel):
        log.error("Could not find RPG notifications channel")
        return

    embed = discord.Embed(
        title="A new release has just been posted",
        description=f"Check out the latest release for Supergoon RPG with tag {tag_name}\nPlay the emscripten build here https://escapethefate.supergoon.com or the dev build here https://escapethefate-dev.supergoon.com",
        color=0x00FF00,
    )
    embed.add_field(name="Release URL (downloads and notes)", value=release_url, inline=True)
    embed.add_field(name="Name", value=release_name, inline=True)
    embed.add_field(name="Release Body", value=release_body, inline=False)
    embed.add_field(name="Repository Url", value=url, inline=False)
    embed.add_field(name="Post issues here.", value=f"{url}/issues", inline=False)

    await channel.send(embed=embed)
