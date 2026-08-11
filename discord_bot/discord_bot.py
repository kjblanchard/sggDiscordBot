import logging

import discord
from discord.ext import commands

log = logging.getLogger(__name__)

discord_application_id = ""
supergoon_games_server_id = ""


class SggBot(commands.Bot):
    def __init__(self, token_value: str, **kwargs):
        super().__init__(**kwargs)
        self.token_value = token_value


def initialize_discord(token: str, app_id: str, supergoon_server_id: str) -> SggBot:
    global discord_application_id, supergoon_games_server_id
    discord_application_id = app_id
    supergoon_games_server_id = supergoon_server_id
    intents = discord.Intents.default()
    bot = SggBot(token_value=token, command_prefix="!", intents=intents)
    return bot

def get_supergoon_games_server_id():
    return supergoon_games_server_id


async def open_discord_connection(bot: SggBot):
    await bot.login(bot.token_value)
    bot.loop.create_task(bot.connect())


async def close_discord(bot: SggBot):
    await bot.close()
    log.info("Discord connection closed")
