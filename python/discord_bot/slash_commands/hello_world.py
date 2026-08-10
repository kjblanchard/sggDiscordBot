import discord
from discord import app_commands


def add_hello_world_slash_command(bot: discord.Client, guild: discord.Object):
    @bot.tree.command(name="hello-world", description="Showcase of a basic slash command", guild=guild)
    async def hello_world(interaction: discord.Interaction):
        await interaction.response.send_message("Hello world!")
