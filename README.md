# System Design App

I am building this system design app because I did not find any proper free app-based structured resources to study System Design progressively. I found the very useful repository of [System Design Primer](https://github.com/donnemartin/system-design-primer) by Donne Martin, however, there was no way to track progress, see all available topics in a dashboard manner, continue progress where I left off etc.

So I decided to build an app out of the content of the repository.

## Objective

This is going to be a Go-based app. I will use Claude to build the frontend, but as for the backend I will do the entirely from scratch.

## Requirements

The app will use:

- Gin as the web engine
- Postgresql as the persistent store
- Redis for the user session management

I will probably have Claude build the frontend with React.js since that is the framework I am most familiar with still.
