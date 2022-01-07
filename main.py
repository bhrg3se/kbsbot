# bot.py
import os
import discord
from dotenv import load_dotenv
import colonist


def get_message(resp):
    return 'Username: %s\nKarma: %s\nTotal Games:%d \nWins Percent: %s' % (resp["username"],resp["karma"],resp["totalGames"],resp["winPercent"])


load_dotenv()
TOKEN = os.getenv('DISCORD_TOKEN')


class Client(discord.Client):
    async def on_ready(self):
        for guild in client.guilds:
            print(guild.name)
        # if guild.name == GUILD:
        #     break

        print(
            f'{self.user} is connected to the following guild:\n'
            f'{guild.name}(id: {guild.id})'
        )
        print(f'{client.user} has connected to Discord!')

    async def on_message(self,message):
        if message.content.startswith("--check"):
            name=message.content[7:]
            name=name.strip()
            print(name)
            if name=="":
                name=message.author.name
            resp= colonist.get_profile(name)

            await message.channel.send(get_message(resp))
        if message.content.startswith("Hello KBSBot"):
            await message.channel.send("Hello "+message.author.name)

client = Client()

client.run(TOKEN)


