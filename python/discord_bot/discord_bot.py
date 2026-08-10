import logging

import discord
from discord import app_commands

log = logging.getLogger(__name__)

discord_application_id = ""
supergoon_games_server_id = ""


def initialize_discord(token: str, app_id: str, supergoon_server_id: str) -> discord.Client:
    global discord_application_id, supergoon_games_server_id
    discord_application_id = app_id
    supergoon_games_server_id = supergoon_server_id

    intents = discord.Intents.default()
    bot = discord.Client(intents=intents)
    bot.tree = app_commands.CommandTree(bot)
    bot.token_value = token
    return bot


async def open_discord_connection(bot: discord.Client):
    await bot.login(bot.token_value)
    bot.loop_task = bot.loop.create_task(bot.connect())


async def close_discord(bot: discord.Client):
    await bot.close()
    log.info("Discord connection closed")
