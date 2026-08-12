import logging

import discord
from discord.ext import commands

log = logging.getLogger(__name__)


class SggBot(commands.Bot):
    def __init__(self, token_value: str, server_id: str, app_id: str, **kwargs):
        "Pass all unknown params to the actual discord bot, and then assign the rest to our class"
        super().__init__(**kwargs)
        self.token_value = token_value
        self.server_id = server_id
        self.app_id = app_id

bot_instance: SggBot 

def initialize_discord(token: str, app_id: str, supergoon_server_id: str) -> SggBot:
    intents = discord.Intents.default()
    bot = SggBot(token_value=token, server_id=supergoon_server_id, app_id=app_id, command_prefix="!", intents=intents)
    global bot_instance
    bot_instance = bot
    return bot_instance

def get_current_bot() -> SggBot:
    return bot_instance


async def open_discord_connection(bot: SggBot):
    await bot.login(bot.token_value)
    bot.loop.create_task(bot.connect())


async def close_discord(bot: SggBot):
    await bot.close()
    log.info("Discord connection closed")
