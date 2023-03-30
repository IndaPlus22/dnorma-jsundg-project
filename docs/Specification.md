# Project Specification
[Project](https://github.com/IndaPlus22/dnorma-jsundg-project)

## Platformer with Go
We aim to make a platformer with Golang. We'll make use of the [Pixel package](https://pkg.go.dev/github.com/faiface/pixel). The game will be in singleplayer and the playable character will be able to interact with objects and obstacles in different ways. The project is quite feasible since we use a powerful package for the graphics which is the main concern. A full completion of the project might not be as feasible since the team isn't very experienced with game development.

We will be using a simpler version of Git Flow as our Git methodology. There will be a `main`branch which is protected from pushes - only merging pull requests can be done. Under `main` is the `develop` branch which merges all `feature` branches. This branch only merges with main once a major milestones has been reached. There will also be `feature` branches which handle one issue at the time. They all merge in to the `develop` branch when completed.

### Requirements
* There should be a playable character that can move around the map with the WASD keys
    * Player collisions has to be implemented
    * Getting rid of obstacles should be available in certain situations
* There should be obstacles that the player can interact with
    * "Enemies" with simple movement paths should exist
* The player should be able to win or lose
    * Scores and highscores should be kept
* Decent graphics on everything

## Team members
Task areas haven't been decided just yet.
* Daniel Norman aka dnorma
* Johan Sundgren aka jsundg

## Naming conventions
##### Issues and commits
Issues and commit messages are going to be named according to github standards, in other words in future tense. Examples: "Add wall collision", "Fix annoying bug".

##### Pull requests
Pull requests will be named after their associated issue. It will be in the format of "issue/[issue number]-[issue-name], i.e issue #13 "Improve performance" would be "issue/13-improve-performance".

## Weekly milestones
Week 1:     Character attributes(physics properties, position, movement, etc.).
Week 2:     Create map objects such as walls, floors, etc. and collisions.
Week 3:     Graphics for the map and character.
Week 4:     Throwable objects and interactable thingies
Week 5:     More interactable thingies
Week 6:     Extra week for catch up and fun extra features. Optimization.