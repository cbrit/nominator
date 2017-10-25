# NOMinator

This is a remake a Hipchat bot I created at work. I am using it's rebirth as a means to learn the Go programming language. Therefore, I am a Baby Gopher...

[![baby-gopher](https://raw.githubusercontent.com/drnic/babygopher-site/gh-pages/images/babygopher-badge.png)](http://www.babygopher.org)

Since the original NOMinator was created in my companies network, I could not host the code publically. I will be building from the ground up using Slack instead of Hipchat since it's easier to use as a single user. 

## Usage

Create a config.json that contains your Slack API Auth credentials in the following format:
```
{
     "username": "",
     "channel": "",
     "url": "",
     "secrets": {
          "client_id": "",
          "client_secret": "",
          "verification_token": "",
          "api_token": ""
     }
}
```