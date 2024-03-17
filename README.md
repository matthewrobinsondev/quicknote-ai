# What is quick note ai?
Quick note ai is for all the termites out there who refuse to leave there terminals even for a quick second!

It's for when you are trying to create a quick note on a thought that has popped into your head without leaving your state of flow. I consistently think of these things each day but either choose to stay in my state of flow or to exit that flow and write down my thought. However I then run into the issue of taking too long on the note or never coming back to finishing the note.

Quick note takes the original concept of this & shoots it off into the AI mothership & spits back out a markdown file with:
- Notes
- Examples
- Resources
into your directory of choice.

# Installation
## Requirements
`go version 1.22.0`

## Clone the directory
`git clone https://github.com/matthewrobinsondev/quicknote-ai`

## Create a configuration file in your root config
for example on linux it's
`$HOME/.config/quicknote-ai/config.toml`

You can copy the example `config.toml` or just create your own file but only two keys are needed
```toml
openai_api_key="xyz"
notes_directory="Documents/Notes"
```

## Build the cli tool
```go
go build -o quicknote-ai

```

## Move the executable to your bin location
for example
`sudo mv quicknote-ai /usr/local/bin/`

# Usage
*Currently only one type of note is accepted. I'm currently looking at adding both a flag for open ai model & note templae. However in the current version they are both hard coded.*

`quicknote-ai note -t "Your quick note prompt"`

![](https://github.com/matthewrobinsondev/quicknote-ai/blob/master/example-usage.gif)
