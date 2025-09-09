# mdblist-cli

[![GitHub license](https://img.shields.io/github/license/luckylittle/mdblist-cli.svg)](https://github.com/luckylittle/mdblist-cli/blob/master/LICENSE)
[![Version](https://img.shields.io/badge/Version-0.0.6-green.svg)](https://github.com/luckylittle/mdblist-cli/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/luckylittle/mdblist-cli)](https://goreportcard.com/report/github.com/luckylittle/mdblist-cli)

## Overview

### About MDBList

[MDBList.com](https://mdblist.com) is your go-to tool for creating dynamic, auto-updating movie and show lists tailored to your preferences. Seamlessly integrated with Plex, Radarr, Sonarr, Kodi and Stremio, it combines the power of multiple rating platforms like IMDb, TMDb, Letterboxd, Rotten Tomatoes, Metacritic, MyAnimeList, and RogerEbert. Whether you're building a watchlist, tracking ratings, or syncing your library progress, mdblist.com makes it effortless. Don't forget to [become a supporter of MDBList](https://docs.mdblist.com/docs/supporter) and say hi in the [Discord](https://discord.gg/bDWQb3mGkr)!

### About `mdblist-cli`

The idea is to perform basic MDBList operations directly from your command line, regardless of what OS you use. This Golang project leverages it's [native API](https://api.mdblist.com/) - the API documentation is located on [Apiary](https://mdblist.docs.apiary.io/).

## Download

See the latest [release](https://github.com/luckylittle/mdblist-cli/releases/). We have single binary for MacOSX/ARM64, MacOSX/x86_64, FreeBSD, Linux/ARM, Linux/ARM64, Linux/x86_64, Windows x64.

## Usage

* :warning: Set up the API key environment variable in your Shell first :warning:

```bash
export MDBLIST_API_KEY=abcdefghijklmnopqrstuvwxy
```

* No arguments - available commands

<details>

```bash
$ ./mdblist-cli
A command-line interface to perform various actions against the MDBList RESTful API.

Usage:
  mdblist-cli [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  get         Get resources from MDBList
  help        Help about any command
  search      Search resources in MDBList
  update      Update resources in MDBList

Flags:
  -h, --help            help for mdblist-cli
  -o, --output string   Output format (json, yaml) (default "json")

Use "mdblist-cli [command] --help" for more information about a command.
```

</details>

* Help - Get

<details>

```bash
$ ./mdblist-cli get --help
Get resources from MDBList.

Usage:
  mdblist-cli get [command]

Available Commands:
  last-activities Fetch the last activity timestamps for sync.
  list            Retrieves details of a list.
  list-changes    Returns Trakt IDs for items changed after the last list update.
  list-items      Fetches items from a specified list.
  media-info      Fetch information about a media item
  my-limits       Show information about user limits.
  my-lists        Fetches users lists.
  top-lists       Outputs the top lists sorted by Trakt likes.
  user-lists      Fetch a user's lists.
  watchlist-items Fetches watchlist items, they are sorted by date added.

Flags:
  -h, --help   help for get

Global Flags:
  -o, --output string   Output format (json, yaml) (default "json")

Use "mdblist-cli get [command] --help" for more information about a command.
```

</details>

* Help - Search

<details>

```bash
$ ./mdblist-cli search --help
Search resources in MDBList.

Usage:
  mdblist-cli search [command]

Available Commands:
  lists       Search public lists by title.
  media       Search for movie, show or both (any).

Flags:
  -h, --help   help for search

Global Flags:
  -o, --output string   Output format (json, yaml) (default "json")

Use "mdblist-cli search [command] --help" for more information about a command.
```

</details>

* Help - Update

<details>

```bash
$ ./mdblist-cli update --help
Update resources in MDBList.

Usage:
  mdblist-cli update [command]

Available Commands:
  list-items  You can modify static list by adding or removing items.
  list-name   Updates the name of a list.

Flags:
  -h, --help   help for update

Global Flags:
  -o, --output string   Output format (json, yaml) (default "json")

Use "mdblist-cli update [command] --help" for more information about a command.
```

</details>

## Examples

* `mdblist-cli get my-limits` - Get information about the API key's limits

<details>

```json
{
  "api_requests": 1000,
  "api_requests_count": 22,
  "user_id": 12345,
  "patron_status": "former_patron",
  "patreon_pledge": 0
}
```

</details>

* `mdblist-cli get top-lists` - Top lists **JSON** (omitted)

<details>

```json
[
  {
    "id": 2194,
    "name": "Latest TV Shows",
    "slug": "latest-tv-shows",
    "description": "",
    "mediatype": "show",
    "items": 300,
    "likes": 477,
    "user_id": 1230,
    "user_name": "garycrawfordgc",
    "dynamic": true
  },
  ...
```

</details>

* `mdblist-cli get top-lists --output=yaml` - Top lists again, but **YAML** (omitted)

<details>

```yaml
- id: 2194
  name: Latest TV Shows
  slug: latest-tv-shows
  description: ""
  mediatype: show
  items: 300
  likes: 477
  userid: 1230
  username: garycrawfordgc
  dynamic: true
  private: false
  ...
```

</details>

* `mdblist-cli get list --id 2194` - Specific list details

<details>

```json
[
  {
    "id": 2194,
    "name": "Latest TV Shows",
    "slug": "latest-tv-shows",
    "description": "",
    "mediatype": "show",
    "items": 300,
    "likes": 477,
    "user_id": 1230,
    "user_name": "garycrawfordgc",
    "dynamic": true
  }
]
```

</details>

* `mdblist-cli get list-items --username garycrawfordgc --listname "latest-tv-shows"` - Get items from the list (omitted)

<details>

```json
{
  "movies": [],
  "shows": [
    {
      "id": 253941,
      "rank": 1,
      "adult": 0,
      "title": "The Paper",
      "imdb_id": "tt32159809",
      "tvdb_id": 449872,
      "language": "en",
      "mediatype": "show",
      "release_year": 2025,
      "spoken_language": "en"
    },
    ...
```

</details>

* `mdblist-cli get media-info imdb show tt32159809 --output=yaml` - Get details about the media, **YAML** (omitted)

<details>

```yaml
title: The Paper
year: 2025
released: "2025-09-04"
releaseddigital: ""
description: The documentary crew that immortalized Dunder Mifflin's Scranton branch is in search of a new subject when they discover a historic Toledo newspaper, The Truth Teller, and the eager publisher trying to revive it.
runtime: 297
score: 0
scoreaverage: 72
ids:
    imdb: tt32159809
    trakt: 239158
    tmdb: 253941
    tvdb: 449872
    mal: null
type: show
ratings:
    - source: imdb
      value: 6.9
      score: 69
      votes: 981
      url: 67
      ...
```

</details>

* `mdblist-cli search media any -q "The Paper"` - Search query, first 100 items (omitted)

<details>

```json
    ...
    {
      "title": "The Paper",
      "year": 2025,
      "score": 0,
      "score_average": 72,
      "type": "show",
      "ids": {
        "imdbid": "tt32159809",
        "tmdbid": 253941,
        "traktid": 239158,
        "malid": null,
        "tvdbid": 449872
      }
    }
  ],
  "total": 100
}
```

</details>

* `mdblist-cli update list-items -a add -i 113124 --movie-imdb tt26581740` - Add item to the static list

<details>

`List items updated successfully (action: add).`

```json
{
  "added": {
    "episodes": 0,
    "movies": 1,
    "seasons": 0,
    "shows": 0
  },
  "existing": {
    "episodes": 0,
    "movies": 0,
    "seasons": 0,
    "shows": 0
  },
  "not_found": {
    "episodes": 0,
    "movies": 0,
    "seasons": 0,
    "shows": 0
  }
}
```

</details>

* `mdblist-cli update list-items -a remove -i 113124 --movie-imdb tt26581740 --output yaml` - Remove item from the static list, **YAML**

<details>

`List items updated successfully (action: remove).`

```yaml
added: {}
existing: {}
notfound:
    episodes: 0
    movies: 0
    seasons: 0
    shows: 0
```

</details>

## Development

### Requirements

Tested with `go version go1.24.3 linux/amd64`.

### Export your API key first

`export MDBLIST_API_KEY="your_actual_api_key_here"`

### Install dependencies

`go mod tidy`

### Get help (test before build)

`go run . --help`

### Get your API limits (test before build)

`go run . get my-limits`

### Get your lists (test before build)

`go run . get my-lists`

## License

GNU General Public License

## Author

Lucian Maly <<lmaly@redhat.com>>

_Last update: Tue 09 Sep 2025 07:04:27 UTC_
