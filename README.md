# Perplexity CLI

[![Go](https://github.com/japelsin/pplx/actions/workflows/release.yml/badge.svg)](https://github.com/japelsin/pplx/actions/workflows/release.yml)
[![License](https://img.shields.io/badge/license-MIT-blue)](https://github.com/japelsin/pplx/blob/main/LICENSE)

CLI for searching with [Perplexity](https://www.perplexity.ai/)'s API. Can also be used as a chatbot.

## Prerequisites

- Perplexity account and API key. You'll be prompted for the API key the first time you run `pplx`.

## Installation

### With [Homebrew](https://brew.sh)

```bash
brew install japelsin/tap/pplx
```

### From source

If you have [go](https://go.dev/) installed:

```bash
go install github.com/japelsin/pplx@latest
```

You could also grab the appropriate executable from [releases](https://github.com/japelsin/pplx/releases).

## Usage

### Search

Search command. Most parameters allowed by `pplx-api` are available as options. The model is set through the config (see below).

```
Usage:
  pplx search [flags]

Flags:
  -m, --max_tokens int          Maximum number of tokens to be used per request. Defaults to config value. (default 1000)
  -f, --frequency_penalty int   How much to penalize token frequency.
  -p, --presence_penalty int    How much to penalize token presence. Between -2 and 2.
  -t, --temperature int         The amount of randomness in the response. Between 0 and 2.
  -K, --top_k int               Number of tokens to consider when generating tokens. Between 0 and 2048.
  -P, --top_p int               Nucleus sampling. Probability cutoff for token selection. Between 0 and 1.
```

The API reference can be found [here](https://docs.perplexity.ai/reference/post_chat_completions).

### Config

```
Usage:
  pplx config [command]

Available Commands:
  path        Get configuration file path
  reset       Reset config
  set         Set config value
```
