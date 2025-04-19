<h2 align="center">TrackIt</h2>

<h4 align="center">a tool for keeping tabs on website changes and getting instant alerts when something's different.</h4>

<br/>
<p align="center">
  <a href="#installation">Install</a> •
  <a href="#usage">Usage</a> •
  <a href="#contributing">Contributing</a> •
  <a href="https://t.me/youngafru">Hit me up</a> •
  <a href="https://twitter.com/nimagholamy">Follow me</a>
</p>

---

TrackIt is a slick little tool for hackers, builders, and anyone who wants to keep an eye on their websites. It checks your URLs, compares the content, and pings you when something changes. Whether you're watching for defacements, sneaky edits, or just wanna know when your site updates, TrackIt’s got your back. Discord and Telegram alerts? Yeah, we got those too.

<br/>

## Why TrackIt?
- **Spot Changes Fast**: Keep an eye on your URLs and know when something’s off.
- **Get Pinged Instantly**: Alerts hit your Discord or Telegram like a boss.
- **Speedy Checks**: Handles multiple URLs at once, no sweat.
- **Customizable AF**: Tweak the config to fit your flow.
- **Secure by Default**: SSL checks and redirect handling baked in.

<br/>

## Installation
Run the following command to install the latest version. Easy peasy.

```bash
go install -v github.com/nimaism/trackit/cmd/trackit@latest
```

<br/>

## Usage
```console
$ ./trackit -h

████████╗██████╗░░█████╗░░█████╗░██╗░░██╗██╗████████╗
╚══██╔══╝██╔══██╗██╔══██╗██╔══██╗██║░██╔╝██║╚══██╔══╝
░░░██║░░░██████╔╝███████║██║░░╚═╝█████═╝░██║░░░██║░░░
░░░██║░░░██╔══██╗██╔══██║██║░░██╗██╔═██╗░██║░░░██║░░░
░░░██║░░░██║░░██║██║░░██║╚█████╔╝██║░╚██╗██║░░░██║░░░
░░░╚═╝░░░╚═╝░░╚═╝╚═╝░░╚═╝░╚════╝░╚═╝░░╚═╝╚═╝░░░╚═╝░░░
                 v1.0.0

TrackIt is your homie for tracking website changes.

Usage:
  ./trackit [flags]

Flags:
   -config string   Path to the configuration file (default "config.yaml")
   -duc, -disable-update-check  Disable automatic update check
```

### Setup Your Config
TrackIt runs on vibes and YAML. Here’s how you set it up:

```yaml
urls_file: "urls.txt" # List of URLs to watch
interval: 10         # How often to check (in minutes)
storage_file: "data.json" # Where to save the hashes
concurrency: 10      # How many URLs to check at once
notifier:
  discord:
    enabled: true
    webhook_url: "https://discord.com/api/webhooks/your-webhook-url"
  telegram:
    enabled: true
    bot_token: "your-telegram-bot-token"
    chat_id: "your-telegram-chat-id"
network:
  timeout_sec: 10
  verify_ssl: true
  disable_redirect: false
```

### Let’s Roll
1. Create a file (`urls.txt`) with the URLs you wanna track:
   ```txt
   http://example.com
   http://another-example.com
   ```

2. Create a `config.yaml` file and set it up like the example above. Customize it to fit your flow.

3. Fire it up:
   ```bash
   ./trackit -config config.yaml
   ```

4. Sit back and relax. If something changes, TrackIt will slide into your DMs (Discord/Telegram).

<br/>

## Contributing
Got ideas? Found a bug? Wanna make this thing even cooler? Pull requests and issues are always welcome. Let’s build something awesome together.

Oh, and if you’re vibing with TrackIt, drop a star ⭐ or buy me a coffee. Every little bit helps keep the hustle alive.

<br/>

## License

**TrackIt** is free as in freedom. Do whatever you want with it. Distributed under the **MIT** License. See the LICENSE file for the boring legal stuff.

---

`haPpY HaCkiNg.`
