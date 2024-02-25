# Perplexity CLI

CLI for interfacing with [Perplexity](https://www.perplexity.ai/)'s API. Can also be used as a chatbot.

## Prerequisites

- Perplexity account and API key. You'll be prompted for the API key the first time you run `pplx`.

## Installation

### With Homebrew

If you're using [Homebrew](https://brew.sh/):

```bash
brew install japelsin/tap/pplx
```

### From source

If you have [go](https://go.dev/) installed:

```bash
go install github.com/japelsin/pplx@latest
```

Otherwise you could grab the appropriate executable from releases.

## Usage

### Search

The response is always streamed, all other [parameters](https://docs.perplexity.ai/reference/post_chat_completions) are available. The model is set through the config (see below).

```
Usage:
  pplx search [flags]

Flags:
  -f, --frequency_penalty int   How much to penalize token reuse. 1 is no penalty. Between 0 and 1.
  -m, --max_tokens int          Maximum number of tokens to be used per request. Defaults to config value. (default 1000)
  -p, --presence_penalty int    How much to penalize existing tokens. Between -2 and 2.
  -t, --temperature int         The amount of randomness in the response. Between 0 and 2.
  -K, --top_k int               Number of tokens to consider when generating tokens, lower values result in higher probability tokens being used. Between 0 and 2048.
  -P, --top_p int               Nucleus sampling. Probability cutoff for token selection, lower values result in higher probability tokens being used. Between 0 and 1.
```

### Config

Configure `pplx`.

```
Usage:
pplx config [command]

Available Commands:
path        Get configuration file path
reset       Reset config
set         Set config value
```

#### Config values

- **Additional instructions**: Instructions appended to queries. If you intend to use `pplx` as a search engine it's recommended to instruct it to always provide its sources.

- **Model**: Available models are: `pplx-7b-chat`, `pplx-70b-chat`, `pplx-7b-online`, `pplx-70b-online`, `llama-2-70b-chat`, `codellama-34b-instruct`, `codellama-70b-instruct`, `mistral-7b-instruct`, and `mixtral-8x7b-instruct`.

- **API Key**: Self explanatory
