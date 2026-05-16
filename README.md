# Everdream Bot — YAGPDB Fork

Yet Another General Purpose Discord Bot, customized for Everdream.

## Quick Start (Local Dev)

1. **Copy env files**
   ```bash
   cp yagpdb_docker/app.env.example yagpdb_docker/app.env
   cp yagpdb_docker/db.env.example yagpdb_docker/db.env
   ```

2. **Fill in `yagpdb_docker/app.env`** with your Discord bot token, client ID, secret, etc.

3. **Set a strong password in `yagpdb_docker/db.env`**

4. **Build and run**
   ```bash
   docker compose -f yagpdb_docker/docker-compose.yml up -d --build
   ```

5. **Initialize the database**
   ```bash
   docker compose -f yagpdb_docker/docker-compose.yml exec db psql -U yagpdb -c "CREATE DATABASE yagpdb;"
   ```

6. **Access the web panel** at `http://localhost`

## Deploy to Coolify

1. Create a new service in Coolify → select "Docker Compose"
2. Point to this repo (`xaelistic/yagpdb`)
3. Set the compose file path: `yagpdb_docker/docker-compose.yml`
4. Add environment variables from `app.env` and `db.env`
5. Deploy

## Features

- **Auto-moderation** — spam, invite links, bad words, caps, emoji spam
- **Reaction roles** — users react to get roles
- **Custom commands** — `!dream`, `!profile`, etc.
- **Logging** — message edits/deletes, mod actions, joins/leaves
- **Auto-role** — assign roles on join
- **Verification** — rules acceptance gate
- **Web dashboard** — configure everything from a browser

## Configuration

After first setup, go to `http://your-host` and:
1. Log in with Discord
2. Select your server
3. Configure:
   - General settings (prefix, etc.)
   - Auto-mod rules
   - Reaction roles
   - Custom commands
   - Logging channels
   - Verification

## Updating

```bash
docker compose -f yagpdb_docker/docker-compose.yml pull
docker compose -f yagpdb_docker/docker-compose.yml up -d --build
```

## License

GPL-3.0 (same as upstream YAGPDB)
