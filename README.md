
# dioggy

This is a simple program that automatically pulls a project from GitHub, runs it, restarts it if it exits and updates it if updates are pushed to the repo.

Initially created for the [if-schleife-bot](https://github.com/alexcoder04/if-schleife-bot), runs it by default.

Too lazy to write docs, if you have questions, just open an [issue](https://github.com/alexcoder04/dioggy/issues).

## Usage

```sh
./dioggy
```

## Configuration

### Command-line arguments

 - `-enable-discord-notifications`: sends a Discord message when the application is started/updated

### Environmental variables

 - `GITHUB_CLONE`: which GitHub repo to use, `user/repo` format
 - `PREPARE_COMMAND`: command to run after cloning/pulling
 - `EXEC_COMMAND`: command to run
 - `DISCORD_WEBHOOK_URL`: where to send Discord messages (optional)

