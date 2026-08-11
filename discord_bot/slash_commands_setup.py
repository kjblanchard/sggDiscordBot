import logging

import discord

from discord_bot.slash_commands.hello_world import add_hello_world_slash_command
from discord_bot.slash_commands.check_repos import add_check_repos_slash_command
from discord_bot.discord_bot import SggBot

log = logging.getLogger(__name__)


def add_all_slash_commands(bot: SggBot, server_id):
    guild = discord.Object(id=int(server_id))
    add_hello_world_slash_command(bot, guild)
    add_check_repos_slash_command(bot, guild)

    @bot.event
    async def on_ready():  # pyright: ignore[reportUnusedFunction]
        await bot.tree.sync(guild=guild)
        log.info("Slash commands synced")
