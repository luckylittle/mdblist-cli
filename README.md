# mdblist-cli

[![GitHub license](https://img.shields.io/github/license/luckylittle/mdblist-cli.svg)](https://github.com/luckylittle/mdblist-cli/blob/master/LICENSE)
[![Version](https://img.shields.io/badge/Version-1.0.0-green.svg)](https://github.com/luckylittle/mdblist-cli/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/luckylittle/mdblist-cli)](https://goreportcard.com/report/github.com/luckylittle/mdblist-cli)

## Overview

### About MDBList

[MDBList.com](https://mdblist.com) is your go-to tool for creating dynamic, auto-updating movie and show lists tailored to your preferences. Seamlessly integrated with Plex, Radarr, Sonarr, Kodi and Stremio, it combines the power of multiple rating platforms like IMDb, TMDb, Letterboxd, Rotten Tomatoes, Metacritic, MyAnimeList, and RogerEbert. Whether you're building a watchlist, tracking ratings, or syncing your library progress, mdblist.com makes it effortless. Don't forget to [become a supporter of MDBList](https://docs.mdblist.com/docs/supporter) and say hi in the [Discord](https://discord.gg/bDWQb3mGkr)!

### About `mdblist-cli`

The idea is to perform basic MDBList operations directly from your command line, regardless of what OS you use. This Golang project leverages it's [native API](https://api.mdblist.com/) - the API documentation is located here: [Apiary](https://mdblist.docs.apiary.io/).

## Download

See the latest [release](https://github.com/luckylittle/mdblist-cli/releases/). We have single binary for MacOSX/ARM64, MacOSX/x86_64, FreeBSD, Linux/ARM, Linux/ARM64, Linux/x86_64, Windows x64.

## Usage

* Set up the API key environment variable in your Shell:

```bash
export MDBLIST_API_KEY=abcdefghijklmnopqrstuvwxy
```

* Help - No arguments

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
  update      Update resources on MDBList

Flags:
  -h, --help            help for mdblist-cli
  -o, --output string   Output format (json, yaml) (default "json")

Use "mdblist-cli [command] --help" for more information about a command.
```

* Help - Get

```bash
$ ./mdblist-cli get --help
Get resources from MDBList

Usage:
  mdblist-cli get [command]

Available Commands:
  last-activities Fetch the last activity timestamps for sync
  list            Fetch a specific list by ID or by username and list name
  list-changes    Fetch changes for a list by its ID
  list-items      Fetch items from a list by ID or by username and list name
  media-info      Fetch information about a media item
  my-limits       Get information about your API limits
  my-lists        Fetch your lists
  top-lists       Fetch the top lists
  user-lists      Fetch a user's lists by ID or username
  watchlist-items Fetch items from the user's watchlist

Flags:
  -h, --help   help for get

Global Flags:
  -o, --output string   Output format (json, yaml) (default "json")

Use "mdblist-cli get [command] --help" for more information about a command.
```

* Help - Search

```bash
$ ./mdblist-cli search --help
Search resources in MDBList

Usage:
  mdblist-cli search [command]

Available Commands:
  lists Search for public lists
  media Search for media

Flags:
  -h, --help   help for search

Global Flags:
  -o, --output string   Output format (json, yaml) (default "json")

Use "mdblist-cli search [command] --help" for more information about a command.
```

* Help - Update

```bash
$ ./mdblist-cli update --help
Update resources on MDBList

Usage:
  mdblist-cli update [command]

Available Commands:
  list-name   Update a list's name by ID or by username and list name

Flags:
  -h, --help   help for update

Global Flags:
  -o, --output string   Output format (json, yaml) (default "json")

Use "mdblist-cli update [command] --help" for more information about a command.
```

## Examples

* `mdblist-cli get my-limits` - Get information about the API key's limits

```json
{
  "api_requests": 1000,
  "api_requests_count": 22,
  "user_id": 12345,
  "patron_status": "former_patron",
  "patreon_pledge": 0
}
```

* `mdblist-cli get top-lists` - Top lists **JSON** (omitted)

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

* `mdblist-cli get top-lists --output=yaml` - Same as above, Top lists **YAML** (omitted)

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

* `mdblist-cli get list --id 2194` - Specific list details

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

* `mdblist-cli get list-items --username garycrawfordgc --listname "latest-tv-shows"` - Get items from the list (omitted)

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

* `mdblist-cli get media-info imdb show tt32159809 --output=yaml` - Get details about the media

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

* `mdblist-cli search media any -q "The Paper"` - Search query

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
  "total": 40
}
```

## Development

### Requirements

Tested with `go version go1.24.3 linux/amd64`.

### Export your API key first

`export MDBLIST_API_KEY="your_actual_api_key_here"`

### Install dependencies

`go mod tidy`

### Get help (test)

`go run . --help`

### Get your API limits (test)

`go run . get my-limits`

### Get your lists (test)

`go run . get my-lists`

## License

GNU General Public License

## Author

Lucian Maly <<lmaly@redhat.com>>

_Last update: Sat 06 Sep 2025 10:02:43 UTC_
