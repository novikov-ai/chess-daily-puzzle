# Chess Daily Puzzle App

## Overview

Welcome to the Chess Daily Puzzle App, a Go-powered application designed to bring the excitement of daily chess puzzles from [Lichess](https://lichess.org/) directly to your Mattermost channels. If you're passionate about chess or simply want to challenge your team or community with daily brain teasers, this app is for you.

![App Screenshot](pics/post.png)

**Key Features:**
- **Stay Engaged:** Solve a new chess puzzle every day to keep your mind sharp.
- **Team Building:** Share chess puzzles with your team, fostering friendly competition and stimulating problem-solving discussions.
- **Continuous Learning:** Enhance your chess skills by practicing puzzles consistently.

## Visualize the Chess Board and Make Your Moves

![Redirect to Board](pics/solve.png)

## React with Emojis and Engage in Puzzle Discussions

![Discuss with Your Team](pics/reply.png)

## Architecture

![Architecture](pics/architecture.png)

## Before You Begin

### Mattermost Webhook URL

Before diving into the chess challenges, ensure you have your Mattermost Webhook URL ready. If you haven't set up an incoming webhook yet, follow these simple steps:

1. Create an incoming webhook. You can find step-by-step instructions [here](https://developers.mattermost.com/integrate/webhooks/incoming/).
2. Once created, copy the generated webhook URL.

### App Configuration

Now, let's configure the app to connect with your Mattermost channel:

1. Begin by cloning this repository to your local machine.
2. Locate the `.env.example` file in the project's root directory and rename it to `.env`.
3. Open the `.env` file and replace the `MATTERMOST_WEBHOOK_URL` value (`your_webhook_url_here`) with the actual Mattermost webhook URL you copied earlier.

### Customize Your App Presentation

The app's appearance can be customized according to your preferences. In the `internal/usecases/presentation/config.go` file, you can configure:

- **Username:** Choose a username that will be displayed when the app sends messages.
- **Message:** Define the message template as you see fit.
- **IconURL:** Set the URL for the app's icon or avatar.

Modify these values to personalize the way the app appears when delivering puzzles to your Mattermost channel.

### Logging to your telegram channel

The app provides built-in logging to your telegram channel.

All you need is:

1. Edit `.env` file and replace `TELEGRAM_BOT_TOKEN` value with your real token, which you can get from [BotFather](https://telegram.me/BotFather).
2. Edit `.env` file and replace `TELEGRAM_CHAT_ID` value with your real chatID, which you can get by doing:
   0. If you have a public channel, just use the ID after @ (next steps are for private channel)
   1. Add your Bot to a private channel.
   2. Send a message to the channel as Bot.
   3. Do a GET Request:
      ~~~
      https://api.telegram.org/bot{TELEGRAM_BOT_TOKEN}/getUpdates
      ~~~
   4. You'll get your ID in the response. 

## How to Install

Getting started is straightforward:

1. Simple installation:

   ```bash
   make install
   ```
   
The App fetches a daily chess puzzle and sends it to the configured Mattermost channel using the provided webhook URL at 11:11 AM on Monday, Wednesday and Friday.

However, you could configure any kind of schedule allowed by the [crontab](https://man7.org/linux/man-pages/man5/crontab.5.html).

## How to Debug

1. Set up your debug URL:

   ```bash
   MATTERMOST_WEBHOOK_URL_DEBUG = "your_webhook_url_here_for_debug"
   ```

2. Build and run the application:

   ```bash
   make run-debug
   ```
   
## How to Uninstall

1. To remove the App:

   ```bash
   make uninstall
   ```
   
---

Enjoy the daily chess challenges and let the strategic thinking begin!