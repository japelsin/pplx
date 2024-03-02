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
  -f, --frequency_penalty float   Token frequency penalty [0, 1.0]
  -l, --max_tokens int            Token limit per request (default 1000)
  -m, --model string              Model to use (default "sonar-small-online")
  -p, --presence_penalty float    Token presence penalty [-2.0, 2.0]
  -q, --query string              Your query
  -t, --temperature float         Response randomness [0, 2.0]
  -K, --top_k int                 Number of tokens to sample from [0, 2048]
  -P, --top_p float               Probability cutoff for token selection [0, 1.0]
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
