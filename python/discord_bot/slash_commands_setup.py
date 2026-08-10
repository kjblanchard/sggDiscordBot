import logging

import discord

from discord_bot.slash_commands.hello_world import add_hello_world_slash_command
from discord_bot.slash_commands.check_repos import add_check_repos_slash_command
from discord_bot.discord_bot import supergoon_games_server_id

log = logging.getLogger(__name__)


def add_all_slash_commands(bot: discord.Client):
    guild = discord.Object(id=int(supergoon_games_server_id))
    add_hello_world_slash_command(bot, guild)
    add_check_repos_slash_command(bot, guild)

    @bot.event
    async def on_ready():
        await bot.tree.sync(guild=guild)
        log.info("Slash commands synced")
