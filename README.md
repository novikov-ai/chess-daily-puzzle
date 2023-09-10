# Chess Daily Puzzle App

## Overview

The Chess Daily Puzzle App is a Go application that fetches daily chess puzzles from [Lichess](https://lichess.org/) and sends a message to a specific Mattermost channel using a webhook. This app allows you to stay engaged with chess puzzles and share them with your team or community on Mattermost.

- Stay engaged and challenged by solving a new chess puzzle every day.
- Team Building: Share chess puzzles with your team, encouraging friendly competition and problem-solving discussions.
- Learning: Improve your chess skills by regularly practicing puzzles.

![App screenshot](pics/post.png)

## Analyze board, make moves and find the solution!

![Redirect to board](pics/solve.png)

## React with an emojis and discuss puzzle in replies

![Discuss with your team](pics/reply.png)

## Before Start

### Mattermost Webhook URL

Make sure that you have your Mattermost Webhook URL. You could follow these steps:
1. Create an incoming webhook. Instructions you could find [here](https://developers.mattermost.com/integrate/webhooks/incoming/).
2. Copy it's url. 

### Configure App
1. Clone this repository to your local machine.
2. Rename an .env.example file to .env in the project root directory.
3. Replace MATTERMOST_WEBHOOK_URL value (`your_webhook_url_here`) in .env file with your actual Mattermost webhook URL (copied before).

### Username, Message, and IconURL
In the internal/usecases/presentation/config.go file, you can customize the app's presentation by configuring:

- Username: The username that appears when the app sends a message.
- Message: The message template.
- IconURL: The URL of the app's icon or avatar.

Modify these values to personalize the way the app appears when sending messages.

## How to Run

1. Build the application:

~~~
make build
~~~

2. Run the application:

~~~
./chess-daily-puzzle
~~~

The app will fetch the daily chess puzzle and send it to the configured Mattermost channel using the provided webhook URL.

Enjoy!